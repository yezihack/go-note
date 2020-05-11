package Patterm
//什么是工厂模式
//1.定义一个创建对象类, 让子类自已决定实现哪个工厂类
//优点:屏蔽了实现类的过程与细节.
//缺点:每增加一个产品时,都要增加一个具体类和对象实现工厂

//我们要现实一个造飞机的工厂

//定义枚举类型, 避免实现时误传参数
type flyType int
const (
	WARPLANE flyType = iota //战斗机
	COPTER//直升机
)

//我们先定义一个飞机工厂接口
type Fly interface {
	Effect() string //飞机的作用
	Wheel() int //飞机的轮子
}
//战斗机
type Warplane struct {
	Name string
	Number int
}
func (w *Warplane) Effect() string {
	return w.Name
}
func (w *Warplane) Wheel() int {
	return w.Number
}
//直升机
type Copter struct {
	Name string
	Number int
}
func (w *Copter) Effect() string {
	return w.Name
}
func (w *Copter) Wheel() int {
	return w.Number
}

//工厂生成各种飞机
func NewFly(whatFly flyType) Fly {
	var f Fly
	switch whatFly {
	case WARPLANE: //生产一个战斗机
		f = &Warplane{
			Name:"战斗机,维护世界和平",
			Number:3,
		}
	case COPTER://生产一架直升机.
		f = &Copter{
			Name:"直升机,救援",
			Number:0,
		}
	}
	return f
}

