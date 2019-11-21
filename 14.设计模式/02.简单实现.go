package main

func main() {

}

//场景:万能播器实现并兼容所有的音乐电影格式, 如mp3, avi, mp4等
//1. 创建一个新接口,定义输出的方法
//2. 实现新接口里所有的方法
//3. 创建一个适配器, 实现新接口里的所有的方法.
//4. 适配里持有各种格式的实例

type Commoner interface {
	Play()
}
