/*
适配器就是将不兼容的接口也让新接口能调用. 中间插入一层,也就是适配器层.
1. 创建一个新接口
2. 实现新接口的方法
3. 创建一个适配器,并实现新接口里的所有方法.
4. 适配器里持有一个旧接口.并现在新接口与旧接口的对接工作.
*/
package main

import "fmt"

//SD读卡接口
type SDCarder interface {
	ReadSD()          //读卡能力
	WriteSD(s string) //写卡能力
}

//实现SD读卡能力
type SDCard struct {
	Data string
}

func (s *SDCard) ReadSD() {
	fmt.Println("读取卡内容:", s.Data)
}
func (s *SDCard) WriteSD(data string) {
	s.Data = data
}
func NewSDCard() *SDCard {
	return &SDCard{}
}

type ComputerI interface {
	ReadSD(sd SDCarder)
}
type Computer struct {
}

func (c *Computer) ReadSD(sd SDCarder) {
	sd.ReadSD()
}
func NewComputer() *Computer {
	return &Computer{}
}

//----------------FTCard-------------
//旧结构, 也就是需要实现适配的结构
type TFCard struct {
	data string
}

//实例一个TF对象
func NewTFCard() *TFCard {
	return &TFCard{}
}

func (f *TFCard) ReadTF() {
	fmt.Println("FT读取内容", f.data)
}
func (f *TFCard) WriteTF(data string) {
	f.data = data
}

//创建一个适配器,让SD能读取TF卡
type SDAdapterTF struct {
	//持有一个旧接口.
	tf *TFCard
}

//实例一个适配器
func NewSDAdapterTF(tf *TFCard) *SDAdapterTF {
	return &SDAdapterTF{
		tf: tf,
	}
}

//必须实现新接口的方法1
func (a *SDAdapterTF) ReadSD() {
	a.tf.ReadTF()
}

//必须实现新接口的方法2
func (a *SDAdapterTF) WriteSD(data string) {
	a.tf.WriteTF(data)
}

func main() {
	var computer ComputerI
	computer = NewComputer()
	sd := NewSDCard()
	sd.WriteSD("智零小姐姐")
	computer.ReadSD(sd)

	tf := NewTFCard()
	tf.WriteTF("我是一张旧照片")
	ad := NewSDAdapterTF(tf)
	ad.ReadSD()

}
