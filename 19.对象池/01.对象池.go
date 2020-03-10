package _9_对象池

import (
	"errors"
	"time"
)

//资源对象
type ResourceObj struct {
	ID int
}
//定义一个对象池结构
type ObjPool struct {
	items chan *ResourceObj
}
//实例一个池子并初始多少个资源对象
func NewPool(numObj int) *ObjPool {
	pool := ObjPool{}
	pool.items = make(chan *ResourceObj, numObj)
	for i := 0; i < numObj; i ++ {
		pool.items <- &ResourceObj{ID:i+1}
	}
	return &pool
}
//获取资源,需要注意获取时超时,不会然一直阻塞,导致其它使用者一直阻塞
func (p *ObjPool) GetObj(timeout time.Duration) (obj *ResourceObj, err error) {
	select {
	case obj = <- p.items:
		return
	case <-time.After(timeout):
		err = errors.New("timeout")
		return
	}
}
//还回资源,放回对象池子里.供其它人使用,但是要判断无法放入时报错.一般情况对象为空或池子满啦.
func (p *ObjPool) ReleaseObj(v *ResourceObj) error {
	select {
	case p.items <- v:
		return nil
	default:
		return errors.New("overflow")
	}
}
