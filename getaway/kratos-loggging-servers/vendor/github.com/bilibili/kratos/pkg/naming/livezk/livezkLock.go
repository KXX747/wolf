package livezk
/**
分布式锁
 */

import (
	"time"
	"context"
	"sync"
	"github.com/samuel/go-zookeeper/zk"
	"fmt"
	"sort"
	"strings"
	"path"
)

const(
	// DefaultExpiry is used when Mutex Duration is 0
	DefaultExpiry = 2 * time.Second
	// DefaultTries is used when Mutex Duration is 0
	DefaultTries = 16
	// DefaultDelay is used when Mutex Delay is 0
	DefaultDelay = 130 * time.Millisecond
	// DefaultFactor is used when Mutex Factor is 0
	DefaultFactor = 0.01
)

// Locker interface with Lock returning an error when lock cannot be aquired
type ZKLocker interface {
	Lock() error
	Touch() bool
	Unlock() error
}


type ZkMutex struct {
	Name   string        // Resouce name key
	value []byte //value
	flags int32 //path类型
	acl []zk.ACL //zk权限
	root string //zk的root节点
	path string //zk添加的路径 =root+name
	nodepathzk string //zk添加之后的路径

	Expiry time.Duration //最多可以获取锁的时间，超过自动解锁
	Tries int            //失败最多获取锁的次数
	Delay time.Duration //获取锁失败后等待多少时间后重试

	Quorum int // Quorum for the lock, set to len(addrs)/2+1 by NewMutex()
	Factor float64 // Drift factor, DefaultFactor if 0
	until time.Time //时间，根据该参数控制是否结束

	ctx  context.Context
	nodem sync.Mutex //
	zlivezk *livezk

}


/**
初始化zklock conn
 */
func NewZkMutex(zconn *livezk,mMutex *ZkMutex)(err error){

	mMutex.zlivezk=zconn
	mMutex.nodepathzk =mMutex.path
	return
}


//加锁
func(m *ZkMutex)lockZk()  {
	m.nodem.Lock()
	defer m.nodem.Unlock()
	fmt.Println()
	fmt.Println("进入了 。。。。。")

	/**
 	 在zookeeper指定节点（locks）下创建临时顺序节点node_n
    获取locks下所有子节点children
    对子节点按节点自增序号从小到大排序
    判断本节点是不是第一个子节点，若是，则获取锁；若不是，则监听比该节点小的那个节点的删除事件
    若监听事件生效，则回到第二步重新进行判断，直到获取到锁
 */

	//创建本次lock节点
	err :=m.CreatePathZk()
	if err!=nil {
		fmt.Println("m.CreatePathZk err =",err)
		return
	}

	//获取节点的上一个节点，监听上一个节点是否被删除，删除表示本次的节点为最小节点,即拿到了锁
	//循环获取lock
	paths,_,_,err:=m.zlivezk.zkConn.ChildrenW(m.root)
	if err!=nil {
		fmt.Println("zconn.zkConn.ChildrenW err =",err)
		return
	}
	//排序数组
	sort.Strings(paths)
	//截取地址的最后一节
	spArr:=strings.Split(m.nodepathzk,"/")
	spath := spArr[len(spArr)-1]

	if spath==paths[0] {
		//表示自己已经获取到了锁,并且正在执行任务
		fmt.Println("相等。。。。")
		time.Sleep(200*time.Millisecond)

		err=m.deleteNodepath()
		return
	}else {
		//未拿到锁，需要监听上一个的node删除事件

		//获取当前节点在数组的排序位置
		currentNodePosition:=sort.SearchStrings(paths, spath)
		//获取当前节点的上一个节点
		position:=currentNodePosition-1
		previousNodePostion:=paths[position]
		//上一个节点的路径
		previousNode :=m.root+"/"+previousNodePostion
		fmt.Println("previousNode=",previousNode)
		//监听指定的node节点
		var events <-chan zk.Event
		_,_,events,err=m.zlivezk.zkConn.GetW(previousNode)
		if err!=nil {
			fmt.Println("zlivezk.zkConn.GetW err=",err )
			return
		}
		fmt.Println("zlivezk.zkConn.GetW events=",events )

		for  {
			select {
			case event :=<-events:
				fmt.Println("events =",events)
				if event.Type == zk.EventNodeCreated {
					fmt.Println("has node EventNodeCreated\n", event.Path)

				} else if event.Type == zk.EventNodeDeleted {
					fmt.Println("has new node EventNodeDeleted\n", event.Path)
					//ch := make(chan zk.Event)
					//close(ch)
					m.unlockZk()
				} else if event.Type == zk.EventNodeDataChanged {
					fmt.Println("has node data EventNodeDataChanged", event.Path)

				}else if event.Type == zk.EventNodeChildrenChanged {

					fmt.Println("has node EventNodeChildrenChanged", event.Path)
				}

			default:

			}
		}
	}
}

//解锁
func(m *ZkMutex) unlockZk()(err error){
	m.nodem.Lock()
	defer m.nodem.Unlock()
	err=m.deleteNodepath()
	return
}

func(m *ZkMutex) deleteNodepath() (err error) {
	defer m.nodem.Unlock()
	//删除并释放lock
	err=m.zlivezk.zkConn.Delete(m.nodepathzk,-1)
	fmt.Println("删除节点。。。。。nodepathzk=",m.nodepathzk)
	return
}


//创建path在zk
//ls /microservice/live/lock/kratos_server
func(m *ZkMutex) CreatePathZk()(err error){
	//1,检查key是否在zk中存在
	var isExists bool
	var  stat *zk.Stat
	isExists, stat, err=m.checkIsExistsByzk()
	if err!=nil {
		fmt.Println("m.checkIsExistsByzk() err=",err)
		return
	}
	fmt.Println("isExists=",isExists,"  stat=",stat)
	if isExists {
		return
	}

	//2,节点不存在创建节点,并保存节点
	err=m.createPathEphemralSequentialZk()
	if err!=nil {
		fmt.Println("m.createPathEphemralSequentialZk() err=",err)
		return
	}
	fmt.Println("ZkMutex = ",m.nodepathzk)

	//获取自己的前一个编号
	return
}


//检查zk中是否存在
func(m *ZkMutex) checkIsExistsByzk() (bool, *zk.Stat,error){
	//	path,err=zconn.zkConn.Create(m.Name,m.value,m.flags, m.acl)
	isExists, stat, err:=m.zlivezk.zkConn.Exists(m.nodepathzk)
	if err!=nil {
		return isExists,nil,err
	}
	return isExists,stat,nil
}


//创建临时路径
func (m *ZkMutex) createPathEphemralSequentialZk() (err error){
	if err =m.createAll(m.root,m.flags); err != nil {
		fmt.Println("l.createAll err=",err)
		return
	}
	m.nodepathzk,err=m.zlivezk.zkConn.Create(m.path,m.value,m.flags, m.acl)
	return
}

//创建root节点　
func(m *ZkMutex) createAll(root string,flags int32) (err error) {
	seps := strings.Split(root, "/")
	lastPath := "/"
	ok := false
	for _, part := range seps {
		if part == "" {
			continue
		}
		lastPath = path.Join(lastPath, part)
		if ok, _, err = m.zlivezk.zkConn.Exists(lastPath); err != nil {
			return err
		} else if ok {
			continue
		}
		if _, err = m.zlivezk.zkConn.Create(lastPath, nil, flags, zk.WorldACL(zk.PermAll)); err != nil {
			return
		}
	}
	return
}



