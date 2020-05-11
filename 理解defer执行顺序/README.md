**目录导航**

[TOC]

# go 加载顺序
- 每一个执行程序都是从main.go做为入口
- main.go 使用import引用包, 按照此用顺序依次加载
- 已经加载过的包不会重复加载,只会加载一次.
- 如果A包引用了B包, B又引用了A包, 这样是不允许的

# defer 加载顺序
- 基于go 加载顺序基础上, defer只在当前函数里有自己的顺序, 采用FILO 先进后出 栈形式


# 通过代码进一步解理

- 引用包的顺序
```
import (
	"fmt"
	"github.com/yezihack/studyGo/理解defer执行顺序/demo1"
	"github.com/yezihack/studyGo/理解defer执行顺序/demo2"
	"time"
	"github.com/yezihack/studyGo/理解defer执行顺序/common"
	"flag"
)
```
- 文件夹的顺序

```
- 理解gdefer执行顺序
	- common
			- tools.go
	- demo1
			- demo1.go
	- demo2
			- demo2.go
			- demo2_1.go
    - main.go
```

### 执行结果并分析
- 先按main.go的import包顺序加载
- 第一行是demo1, 发现demo1/demo1.go import common包, common/tools.go有个init函数,先执行
- 然后执行demo1/demo1.go里的init函数
- 发现没有其它包了,又回到main.go的import,准备加载第二个的 demo2 包
- 程序发现demo2包下有个demo2_1.go里有个init函数,并执行它
- 又回到main.go 发现有个common包, 但是这个包之前已经加载,将不会再加载一次.
- 然后再执行 main.go 里的init函数
- 然后继续执行 main函数里的代码, 发现有关键字defer,
- 而且有两个, main函数里的代码全执行完后,再执行写在最后的defer, 依次向上找defer关键并执行

```
2019-03-01 14:59:19.061115871 +0800 CST 我是common文件夹里的tools.go文件中的init
2019-03-01 14:59:19.06119396 +0800 CST 我是demo1文件夹里的demo1.go文件里的init
2019-03-01 14:59:19.061224977 +0800 CST 我是demo2文件夹里的demo2_1.go文件里的init
2019-03-01 14:59:19.061258156 +0800 CST 我是main.go里的init
Normal执行中....
2019-03-01 14:59:19.061308476 +0800 CST 我是demo1文件夹里的demo1.go文件里的show函数, number :1
2019-03-01 14:59:19.061352368 +0800 CST 我是demo2文件夹下的demo2.go文件里的show方法, number: 2
2019-03-01 14:59:19.565080217 +0800 CST 我是main函数 END
2019-03-01 14:59:19.565219345 +0800 CST 我是main函数 START
```

```
2019-03-01 15:00:32.664563857 +0800 CST 我是common文件夹里的tools.go文件中的init
2019-03-01 15:00:32.664635553 +0800 CST 我是demo1文件夹里的demo1.go文件里的init
2019-03-01 15:00:32.66466453 +0800 CST 我是demo2文件夹里的demo2_1.go文件里的init
2019-03-01 15:00:32.664696325 +0800 CST 我是main.go里的init
GoChannel执行中....
2019-03-01 15:00:32.664829948 +0800 CST 我是demo1文件夹里的demo1.go文件里的show函数, number :1
2019-03-01 15:00:32.664883194 +0800 CST 我是demo2文件夹下的demo2.go文件里的show方法, number: 2
2019-03-01 15:00:33.168909883 +0800 CST 我是main函数 END
2019-03-01 15:00:33.169044697 +0800 CST 我是main函数 START

```

## 来源博客: [https://www.sgfoot.com/view/224](https://www.sgfoot.com/view/224)