package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

var (
	fileName = "modify.txt"
)

func main() {
	done := make(chan bool)
	fileName = GetRootDir() + fileName
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	Fsnotify(done)
	<-time.After(5 * time.Second)
	f.WriteString("aa")
	<-time.After(5 * time.Second)
	f.WriteString("bb")

	defer f.Close()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill, os.Interrupt)
	<-sigChan
	done <- true
}

func Fsnotify(done chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				//判断事件发生的类型，如下5种
				// Create 创建
				// Write 写入
				// Remove 删除
				// Rename 重命名
				// Chmod 修改权限
				fmt.Println("op", event.Op)
				fmt.Println("name", event.Name)
				if event.Op&event.Op == fsnotify.Create {
					fmt.Println("Create")
				}
				if event.Op&event.Op == fsnotify.Write {
					fmt.Println("Write")
				}
				if event.Op&event.Op == fsnotify.Remove {
					fmt.Println("remove")
				}
				if event.Op&event.Op == fsnotify.Chmod {

				}
				watcher.Remove(event.Name)
				watcher.Add(event.Name)
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()
	f := GetRootDir()
	err = watcher.Add(f)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		<-done
		watcher.Close()
	}()
}

// GetRootDir 获取执行路径
func GetRootDir() string {
	// 文件不存在获取执行路径
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file = fmt.Sprintf("%s%s", file, string(os.PathSeparator))
	}
	return file
}
