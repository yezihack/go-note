package _3_限流设计

import (
	"sync"
	"time"
)

type CountLimitService struct {
	Interval time.Duration //设置计数器时间间隔
	MaxCount uint //计数器最大值
	Lock sync.Mutex //锁
	ReqCount uint //计数器
}
func NewCountLimitService(interval time.Duration, maxCount uint) *CountLimitService {
	reqLimit := CountLimitService{
		Interval:interval,
		MaxCount:maxCount,
	}
	go func() {
		ticker := time.NewTicker(interval)
		for {
			<-ticker.C
			reqLimit.Lock.Lock()
			reqLimit.ReqCount = 0
			reqLimit.Lock.Unlock()
		}
	}()
	return &reqLimit
}
//获取token
func (l *CountLimitService) GetToken() bool {
	l.Lock.Lock()
	defer l.Lock.Unlock()
	if l.ReqCount < l.MaxCount {
		l.ReqCount += 1
		return true
	}
	return false
}