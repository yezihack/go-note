# go笔记

[时光脚步的博客](https://www.sgfoot.com)

## 学习资料

#### [《Go入门指南》](https://github.com/yezihack/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)


## 命令
1. go build 编译并安装自身包和依赖包
1. go install 安装自身包和依赖包
1. [gofmt](https://golang.org/cmd/gofmt/)格式化
1. [godoc](golang.org/cmd/godoc/) 生成代码文档, godoc -http=:6060

## Go 命名规范
1. 文件名小写
1. 函数,方法,变量名采用驼峰命名法
1. 接口以"er"结束

## 学习目录
1. [理解go执行顺序](理解go执行顺序)
1. [理解defer执行顺序](理解defer执行顺序)
1. [循环获取的坑和解决及性能深入分析](demo_slice/s1)

## 特殊
```
执行优先级 init() 比 main()高,它不能够被人为调用.自动执行,一般用于初使数据.

```