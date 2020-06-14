package tokenbucket

import (
	"sync"
	"time"
)

//用户配置平均发送速率，每隔r秒，发送进去k个
//假设桶最多存放max个令牌，令牌到达时令牌桶已经满了，就丢弃
//每个请求到达，就从令牌桶中删除一个，然后就有了通行的资格

type Bucket struct {
	cap   int          //容量
	ch    chan bool    //存放令牌桶容器
	timer *time.Ticker //时间间隔
	n     int          //每隔时间间隔发送的量
	mu    sync.Mutex
}

func NewBucket(cap int, num int, interval time.Duration) *Bucket {
	bucket := &Bucket{
		cap:   cap,
		ch:    make(chan bool, cap),
		timer: time.NewTicker(interval),
		n:     num,
	}
	go bucket.startTicker()
	return bucket
}

func (bucket *Bucket) startTicker() {
	//因为既有容量限制，又有单次发送数据量的限制
	for i := 0; i < bucket.cap && i < bucket.n; i++ {
		bucket.ch <- true
	}
	for {
		select {
		case <-bucket.timer.C:
			for i := 0; i < bucket.n; i++ {
				bucket.Add()
			}
		}
	}
}

func (bucket *Bucket) Add() {
	bucket.mu.Lock()
	defer bucket.mu.Unlock()
	if len(bucket.ch) < bucket.cap {
		bucket.ch <- true
	}
}

func (bucket *Bucket) Get() bool {
	select {
	case <-bucket.ch:
		return true
	default:
		return false
	}
}
