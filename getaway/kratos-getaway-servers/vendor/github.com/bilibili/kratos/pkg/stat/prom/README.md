# prom

## 项目简介

封装prometheus类。TODO：补充grafana通用面板json文件！！！
指标类别

Prometheus中主要使用的四类指标类型，如下所示
- Counter (累加指标)
- Gauge (测量指标)
- Summary (概略图)
- Histogram (直方图)


定义指标

这里我们需要引入另一个依赖库

go get github.com/prometheus/client_golang/prometheus


下面先来定义了两个指标数据，一个是Guage类型， 一个是Counter类型。分别代表了CPU温度和磁盘失败次数统计，使用上面的定义进行分类。

    cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "cpu_temperature_celsius",
        Help: "Current temperature of the CPU.",
    })
    hdFailures = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "hd_errors_total",
            Help: "Number of hard-disk errors.",
        },
        []string{"device"},
    )
    
注册指标

func init() {
    // Metrics have to be registered to be exposed:
    prometheus.MustRegister(cpuTemp)
    prometheus.MustRegister(hdFailures)
}


使用prometheus.MustRegister是将数据直接注册到Default Registry，就像上面的运行的例子一样，这个Default Registry不需要额外的任何代码就可以将指标传递出去。注册后既可以在程序层面上去使用该指标了，这里我们使用之前定义的指标提供的API（Set和With().Inc）去改变指标的数据内容

func main() {
    cpuTemp.Set(65.3)
    hdFailures.With(prometheus.Labels{"device":"/dev/sda"}).Inc()

    // The Handler function provides a default handler to expose metrics
    // via an HTTP server. "/metrics" is the usual endpoint for that.
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8080", nil))
}


其中With函数是传递到之前定义的label=”device”上的值，也就是生成指标类似于

    cpu_temperature_celsius 65.3
    hd_errors_total{"device"="/dev/sda"} 1

当然我们写在main函数中的方式是有问题的，这样这个指标仅仅改变了一次，不会随着我们下次采集数据的时候发生任何变化，我们希望的是每次执行采集的时候，程序都去自动的抓取指标并将数据通过http的方式传递给我们。

Counter数据采集实例

下面是一个采集Counter类型数据的实例，这个例子中实现了一个自定义的，满足采集器(Collector)接口的结构体，并手动注册该结构体后，使其每次查询的时候自动执行采集任务。

我们先来看下采集器Collector接口的实现


type Collector interface {
    // 用于传递所有可能的指标的定义描述符
    // 可以在程序运行期间添加新的描述，收集新的指标信息
    // 重复的描述符将被忽略。两个不同的Collector不要设置相同的描述符
    Describe(chan<- *Desc)

    // Prometheus的注册器调用Collect执行实际的抓取参数的工作，
    // 并将收集的数据传递到Channel中返回
    // 收集的指标信息来自于Describe中传递，可以并发的执行抓取工作，但是必须要保证线程的安全。
    Collect(chan<- Metric)
}



了解了接口的实现后，我们就可以写自己的实现了，先定义结构体，这是一个集群的指标采集器，每个集群都有自己的Zone,代表集群的名称。另外两个是保存的采集的指标。

type ClusterManager struct {
    Zone         string
    OOMCountDesc *prometheus.Desc
    RAMUsageDesc *prometheus.Desc
}


我们来实现一个采集工作,放到了ReallyExpensiveAssessmentOfTheSystemState函数中实现，每次执行的时候，返回一个按照主机名作为键采集到的数据，两个返回值分别代表了OOM错误计数，和RAM使用指标信息。


func (c *ClusterManager) ReallyExpensiveAssessmentOfTheSystemState() (
    oomCountByHost map[string]int, ramUsageByHost map[string]float64,
) {
    oomCountByHost = map[string]int{
        "foo.example.org": int(rand.Int31n(1000)),
        "bar.example.org": int(rand.Int31n(1000)),
    }
    ramUsageByHost = map[string]float64{
        "foo.example.org": rand.Float64() * 100,
        "bar.example.org": rand.Float64() * 100,
    }
    return
}

 

实现Describe接口，传递指标描述符到channel


// Describe simply sends the two Descs in the struct to the channel.
func (c *ClusterManager) Describe(ch chan<- *prometheus.Desc) {
    ch <- c.OOMCountDesc
    ch <- c.RAMUsageDesc
}



Collect函数将执行抓取函数并返回数据，返回的数据传递到channel中，并且传递的同时绑定原先的指标描述符。以及指标的类型（一个Counter和一个Guage）


func (c *ClusterManager) Collect(ch chan<- prometheus.Metric) {
    oomCountByHost, ramUsageByHost := c.ReallyExpensiveAssessmentOfTheSystemState()
    for host, oomCount := range oomCountByHost {
        ch <- prometheus.MustNewConstMetric(
            c.OOMCountDesc,
            prometheus.CounterValue,
            float64(oomCount),
            host,
        )
    }
    for host, ramUsage := range ramUsageByHost {
        ch <- prometheus.MustNewConstMetric(
            c.RAMUsageDesc,
            prometheus.GaugeValue,
            ramUsage,
            host,
        )
    }
}


创建结构体及对应的指标信息,NewDesc参数第一个为指标的名称，第二个为帮助信息，显示在指标的上面作为注释，第三个是定义的label名称数组，第四个是定义的Labels

func NewClusterManager(zone string) *ClusterManager {
    return &ClusterManager{
        Zone: zone,
        OOMCountDesc: prometheus.NewDesc(
            "clustermanager_oom_crashes_total",
            "Number of OOM crashes.",
            []string{"host"},
            prometheus.Labels{"zone": zone},
        ),
        RAMUsageDesc: prometheus.NewDesc(
            "clustermanager_ram_usage_bytes",
            "RAM usage as reported to the cluster manager.",
            []string{"host"},
            prometheus.Labels{"zone": zone},
        ),
    }
}

 

执行主程序


func main() {
    workerDB := NewClusterManager("db")
    workerCA := NewClusterManager("ca")

    // Since we are dealing with custom Collector implementations, it might
    // be a good idea to try it out with a pedantic registry.
    reg := prometheus.NewPedanticRegistry()
    reg.MustRegister(workerDB)
    reg.MustRegister(workerCA)
}



如果直接执行上面的参数的话，不会获取任何的参数，因为程序将自动推出，我们并未定义http接口去暴露数据出来，因此数据在执行的时候还需要定义一个httphandler来处理http请求。

添加下面的代码到main函数后面，即可实现数据传递到http接口上：


    gatherers := prometheus.Gatherers{
        prometheus.DefaultGatherer,
        reg,
    }

    h := promhttp.HandlerFor(gatherers,
        promhttp.HandlerOpts{
            ErrorLog:      log.NewErrorLogger(),
            ErrorHandling: promhttp.ContinueOnError,
        })
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        h.ServeHTTP(w, r)
    })
    log.Infoln("Start server at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Errorf("Error occur when start server %v", err)
        os.Exit(1)
    }

   

其中prometheus.Gatherers用来定义一个采集数据的收集器集合，可以merge多个不同的采集数据到一个结果集合，这里我们传递了缺省的DefaultGatherer，所以他在输出中也会包含go运行时指标信息。同时包含reg是我们之前生成的一个注册对象，用来自定义采集数据。

promhttp.HandlerFor()函数传递之前的Gatherers对象，并返回一个httpHandler对象，这个httpHandler对象可以调用其自身的ServHTTP函数来接手http请求，并返回响应。其中promhttp.HandlerOpts定义了采集过程中如果发生错误时，继续采集其他的数据。

尝试刷新几次浏览器获取最新的指标信息


clustermanager_oom_crashes_total{host="bar.example.org",zone="ca"} 364
clustermanager_oom_crashes_total{host="bar.example.org",zone="db"} 90
clustermanager_oom_crashes_total{host="foo.example.org",zone="ca"} 844
clustermanager_oom_crashes_total{host="foo.example.org",zone="db"} 801
# HELP clustermanager_ram_usage_bytes RAM usage as reported to the cluster manager.
# TYPE clustermanager_ram_usage_bytes gauge
clustermanager_ram_usage_bytes{host="bar.example.org",zone="ca"} 10.738111282075208
clustermanager_ram_usage_bytes{host="bar.example.org",zone="db"} 19.003276633920805
clustermanager_ram_usage_bytes{host="foo.example.org",zone="ca"} 79.72085409108028
clustermanager_ram_usage_bytes{host="foo.example.org",zone="db"} 13.041384617379178


每次刷新的时候，我们都会获得不同的数据，类似于实现了一个数值不断改变的采集器。当然，具体的指标和采集函数还需要按照需求进行修改，满足实际的业务需求。

   
--------------------- 
作者：mingkai_beijing 
来源：CSDN 
原文：https://blog.csdn.net/u014029783/article/details/80001251 
版权声明：本文为博主原创文章，转载请附上博文链接！