package template2

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

//直接输出到屏幕
func TxtTemplate() error {
	tpl, err := template.New("t1").Parse("我的第一个模板,名称叫:{{.}}")
	name := "文本模板"
	if err != nil {
		return err
	}
	err = tpl.Execute(os.Stdout, name)
	if err != nil {
		return err
	}
	return nil
}

//解析模板内容定向到变量里
func TxtTemplateToParam() error {
	tpl, err := template.New("t2").Parse("我的第2个模板,名称叫:{{.}}")
	name := "文本模板"
	if err != nil {
		return err
	}
	//定义一个buffer,将模板输出的内容定向此变量里
	tplContent := bytes.NewBuffer([]byte{})
	err = tpl.Execute(tplContent, name)
	if err != nil {
		return err
	}
	fmt.Println("输出的模板内容:", tplContent)
	return nil
}

type MyBuffer struct {
	Content []byte
}

func (b *MyBuffer) Write(p []byte) (n int, err error) {
	b.Content = append(b.Content, p...)
	return len(b.Content), nil
}
func TxtTemplateByStruct() error {
	tpl, err := template.New("t2").Parse("我的第3个模板,名称叫:{{.}}")
	name := "文本模板"
	if err != nil {
		return err
	}
	//定义一个buffer,将模板输出的内容定向此变量里
	tplContent := new(MyBuffer)
	tplContent.Content = make([]byte, 0)
	err = tpl.Execute(tplContent, name)
	if err != nil {
		return err
	}
	fmt.Println("输出的模板内容:", string(tplContent.Content))
	return nil
}
