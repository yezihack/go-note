package main

import (
	"code.qschou.com/golang/errors"
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

//对象池,就是池子里预定多少个资源,需要就取资源,取完则还回去.供其它人使用.
type ReusableObj struct {
	//对象池里的资源
	ID int
}
//定义一个池子.
type ObjPool struct {
	bufChan chan *ReusableObj //池子里只放资源对象
}
//实例一个池子.有大小控制
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	//声明chan的缓冲大小
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i ++{ //创建多少个对象放在池子里
		objPool.bufChan <- &ReusableObj{
			ID:i,
		}
	}
	return &objPool
}
//从池子里获取一个资源
func (p *ObjPool) GetObj(timeout time.Duration) (obj *ReusableObj, err error) {
	select {
	case obj = <-p.bufChan:
		fmt.Printf("获取资源ID:%d\n", obj.ID)
		return
	case <-time.After(timeout):
		s := fmt.Sprintf("获取资源ID:%d,超时\n", obj.ID)
		err = errors.New("timeout," +s)
		return
	}
}
//还回资源
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error{
	select {
	case p.bufChan <- obj: //如果缓冲未满则成功,如果池子满了chan会阻塞
		fmt.Printf("还回资源ID:%d\n", obj.ID)
		return nil
	default://如果缓冲满了,无法放回池子.
		return errors.New("overflow")
	}
}
func main() {
	pool := NewObjPool(5)
	//if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	//	log.Fatal(err)
	//}
	for i := 0; i < 10; i ++ {
		obj, err := pool.GetObj(time.Millisecond * 100)
		if err != nil {
			log.Fatal(err)
		} else {
			time.Sleep(time.Millisecond * 200)
			if err := pool.ReleaseObj(obj); err != nil {
				log.Fatal(err)
			}
		}
	}
	fmt.Println("Done ")
}
