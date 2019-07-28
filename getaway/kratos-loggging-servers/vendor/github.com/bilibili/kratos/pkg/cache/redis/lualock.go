package redis
/**
 redis lua lock
 */
import (
	"time"
	"sync"
	"errors"
	"context"
	//xredis "github.com/garyburd/redigo/redis"

)

/**
1,获取当前时间戳，单位是毫秒
2,跟上面类似，轮流尝试在每个master节点上创建锁，过期时间较短，一般就几十毫秒
3,尝试在大多是节点上创建锁，比如5个节点要求就是3个节点n/2+1
4,客户端计算建立好锁的时间，如果建立锁的时间小于超时时间，就算建立成功了
5,要是建立锁失败了，那么就依次之前建立过的锁删除
6,要是别人建立了锁，就不断轮询去尝试取锁

 */

const (
	//控制log输出，默认输出
	isLogOutput = false

	// DefaultExpiry is used when Mutex Duration is 0
	DefaultExpiry = 2 * time.Second
	// DefaultTries is used when Mutex Duration is 0
	DefaultTries = 16
	// DefaultDelay is used when Mutex Delay is 0
	DefaultDelay = 130 * time.Millisecond
	// DefaultFactor is used when Mutex Factor is 0
	DefaultFactor = 0.01
	//DefaultQuorum mutex lock number is cluster%2+1 = DefaultQuorum
	DefaultQuorum =1

	luaDel =`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("del", KEYS[1]) else return 0 end`
	luaQuery =`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("set", KEYS[1], ARGV[1], "xx", "px", ARGV[2]) else return "ERR" end`
	luaAdd=`if redis.call("set", KEYS[1], ARGV[1], "nx", "px", ARGV[2]) then  return "OK"  else return "ERR" end`
	//命令中的NX表示如果key不存在就添加，存在则直接返回。PX表示以毫秒为单位设置key的过期时间，这里是30000ms。
	// 设置过期时间是防止获得锁的客户端突然崩溃掉或其他异常情况，导致redis中的对象锁一直无法释放，造成死锁。
)

// Locker interface with Lock returning an error when lock cannot be aquired
type RedisLocker interface {
	Lock() error
	Touch() bool
	Unlock() error
}


type Mutex struct {
	Name   string        // Resouce name key
	value string //value
	Expiry time.Duration //最多可以获取锁的时间，超过自动解锁
	Tries int            //失败最多获取锁的次数
	Delay time.Duration //获取锁失败后等待多少时间后重试

	Quorum int // Quorum for the lock, set to len(addrs)/2+1 by NewMutex()
	Factor float64 // Drift factor, DefaultFactor if 0
	until time.Time

	ctx  context.Context
	nodes []*Pool //redis connect
	nodem sync.Mutex
}

/**
创建redislock
 */
func NewMutex(ctx context.Context,name string,nodes *Pool)(mMutex *Mutex,err error){
	if len(name)==0 {
		err =errors.New("name is null ");
		return
	}
	mMutex= &Mutex{}
	mMutex.ctx = ctx
	mMutex.Name = name
	mMutex.value ="shanghai"
	mMutex.Tries =DefaultTries
	mMutex.Delay =DefaultDelay
	mMutex.Factor =DefaultFactor
	mMutex.Expiry =DefaultExpiry

	mMutex.nodes =make([]*Pool,2)
	mMutex.nodes = append(mMutex.nodes,nodes)
	mMutex.Quorum =mMutex.QuorumUnmber()
	return
}

/**
设置redis锁的主机数
cluster%2+1 = DefaultQuorum
 */
func(m *Mutex)QuorumUnmber()(quorum int){
	quorum =len(m.nodes)/2+1
	return
}


/**
add lock
`if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("set", KEYS[1], ARGV[1], "xx", "px", ARGV[2])
	else
		return "ERR"
	end`
//conn.Do("EVAL", lua, 1, key,num)
	KEYS[1])=key ARGV[1]=num

 */
func (m *Mutex) Lock() (err error) {
	Info_log(" 等待进入。。。。。%s",m.Name)
	m.nodem.Lock()
	defer m.nodem.Unlock()
	Info_log("")
	Info_log(" 进入。。。。。%s",m.Name)
	retries := m.Tries
	for  i := 0; i < retries; i++  {
		n:=len(m.nodes)
		start := time.Now()
		//在集群2/3的机器中设置锁，防止其中的某个挂掉导致锁失败
		for _,node:= range m.nodes {
			if node==nil {
				continue
			}
			//数据保存到redis
			var reply interface{}
			rconn:=node.Get(m.ctx)
			reply,err=rconn.Do("EVAL",luaAdd,1,m.Name, m.value,int(m.Expiry/time.Millisecond))
			//reply, err = xredis.String(rconn.Do("set", m.Name, m.value, "nx", "px", int(m.Expiry/time.Millisecond)))
			Info_log("Lock 添加锁 reply=%s"," err=%s",string(reply.([]byte)),err)
			rconn.Close()
			if err != nil {
				continue
			}
			if string(reply.([]byte)) != "OK" {
				continue
			}
		}
		Info_log("n=%  Quorum =%s",n,m.Quorum)

		//对比
		until := time.Now().Add(m.Expiry - time.Now().Sub(start) - time.Duration(int64(float64(m.Expiry)*m.Factor)) + 2*time.Millisecond)
		Info_log("time.Now()=%s",time.Now().Before(until))

		//判断是否主机上的redis都已经执行并且是当前时间之前
		if n>m.Quorum && time.Now().Before(until){
			Info_log("m.Quorum=%d",m.Quorum)
			m.until = until
			//fmt.Println("n>m.Quorum return " , n>m.Quorum && time.Now().Before(until))
			time.Sleep(DefaultDelay)
			break
		}
		Info_log("time.Now().Before")


		/**
		未实现过期，重试获取锁
		*/
		for _,node:= range m.nodes {
			if node == nil {
				continue
			}
			//过期删除redis的值
			rconn:=node.Get(m.ctx)
			var status interface{}
			status, err =rconn.Do("EVAL",luaDel,1,m.Name,m.value)
			rconn.Close()
			Info_log("Unlock 删除的值 status=%s err=%s",status,err)

			if err!=nil {
				continue
			}
		}

		// Have no delay on the last try so we can return ErrFailed sooner.
		if i == retries-1 {
			continue
		}

		time.Sleep(DefaultDelay)
	}

	return
}



/**
Unlock

`if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end`

	//conn.Do("EVAL", lua, 1, key,num,time)
	//参数详解
	KEYS[1]=key
	ARGV[1]= num
	ARGV[2]=time
 */
func (m *Mutex) Unlock() (err error) {
	m.nodem.Lock()
	defer m.nodem.Unlock()
	m.value = ""
	//重置为当前时间以前
	m.until = time.Unix(0, 0)

	//删除锁
	for _,node:=range m.nodes {
		if node ==nil {
			continue
		}
		//获取conn
		var status interface{}
		rconn:=node.Get(m.ctx)
		status, err =rconn.Do("EVAL",luaDel,1,m.Name,m.value)
		rconn.Close()
		Info_log("Unlock 删除的值 status=%s err=%s",status,err)
		if err != nil {
			Info_log("status err=%s",err)
		}
		if status == 0 {
			Info_log("status=%s",status)
		}
	}


	return
}

/**

// Touch resets m's expiry to the expiry value.
// It is a run-time error if m is not locked on entry to Touch.
// It returns the status of the touch
func (m *Mutex) Touch() bool {
	m.nodem.Lock()
	defer m.nodem.Unlock()

	value := m.value
	if value == "" {
		panic("redsync: touch of unlocked mutex")
	}

	expiry := m.Expiry
	if expiry == 0 {
		expiry = DefaultExpiry
	}
	reset := int(expiry / time.Millisecond)

	n := 0
	for _, node := range m.nodes {
		if node == nil {
			continue
		}

		conn := node.Get()
		reply, err := touchScript.Do(conn, m.Name, value, reset)
		conn.Close()
		fmt.Println("unlock delScript.Do(conn, m.Name, value)")
		if err != nil {
			continue
		}
		if reply != "OK" {
			continue
		}
		n++
	}
	if n >= m.Quorum {
		return true
	}
	return false
}
 */

/**
log 输出
 */
func Info_log(format string, args ...interface{}) {
	if !isLogOutput {
		return
	}
	//ft:=fmt.Sprintf(format,args)
	//fmt.Println(ft)
}



