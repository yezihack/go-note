package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)
var (
	//接受一个参数,也就一个自定义的文件名,默认:default.prof
	cpuProfile = flag.String("cpu_profile", "cpu.prof", "输入一个记录cpu信息的文件名称")
	memProfile = flag.String("mem_profile", "mem.prof", "输入一个记录内存信息的文件名称")
)
func main() {
	flag.Parse()
	if *cpuProfile != "" {//判断值是否为空
		f, err := os.Create(*cpuProfile) //创建一个文件.
		if err != nil {
			log.Fatal("权限不足,无法创建文件", err)
		}
		err = pprof.StartCPUProfile(f) //开启收集cpu,性能指标信息
		if err != nil {
			log.Fatal("收集cpu信息失败", err)
		}
		defer f.Close() //注意这里的顺序, defer是FILO先进后出,必须要在关闭收集句柄之后再关闭,否则无法收集
		defer pprof.StopCPUProfile() //关闭收集的句柄
	}
	//实际的应用程序,一个简单的非波纳契数列
	for i := 1; i < 20; i ++ {
		ret := fab(i)
		fmt.Print(ret, ",")
	}
	//收集内存信息
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal("权限不足,无法创建文件", err)
		}
		runtime.GC()
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("收集内存信息失败", err)
		}
		defer f.Close()
	}
}
func fab(n int) int {
	if n <= 2 {
		return 1
	}
	return fab(n - 1) + fab(n -2)
}
