# 变量逃逸


## 小总
- 逃逸分析。
    1. 通俗来讲，当一个对象的指针被多个方法或线程引用时，我们称这个指针发生了逃逸。
    1. 逃逸分析决定一个变量是分配在堆上还是分配在栈上。
- 为什么要逃逸分析
    1. 堆:
        1. 会形成内存碎片
        1. 堆适合不可预知大小的内存
        1. 分配需要Go频繁地进行垃圾回收 
        1. 堆分配内存首先需要去找到一块大小合适的内存块，之后要通过垃圾回收才能释放
        1. 垃圾回收会占用比较大的系统开销（占用CPU容量的25%）
    1. 栈:
        1. 栈可以自动清理
        1. 栈内存分配则会非常快
        1. 栈分配内存只需要两个CPU指令：“PUSH”和“RELEASSE”，分配和释放
    1. 变量如何分配
        1. 如果发现一个变量退出函数没有用了,就丢到栈上
        1. 如果发现一个变量退出函数还有其它地方在引用,则分配到堆上.
- 逃逸分析是怎么完成的
    1. Go逃逸分析最基本的原则是：
        1. 如果一个函数返回对一个变量的引用，那么它就会发生逃逸。
    1. 决定分配到哪儿?
        1. 如果函数外部没有引用，则优先放到栈中；
            1. 可能放到堆上的情形：定义了一个很大的数组，需要申请的内存过大，超过了栈的存储能力。
        1. 如果函数外部存在引用，则必定放到堆中；
      
> 参考
1. https://www.cnblogs.com/qcrao-2018/p/10453260.html
1. 【逃逸是怎么发生的？很赞 结尾有很多参考资料】https://www.do1618.com/archives/1328/go-%E5%86%85%E5%AD%98%E9%80%83%E9%80%B8%E8%AF%A6%E7%BB%86%E5%88%86%E6%9E%90/
1. 【Go的变量到底在堆还是栈中分配】https://github.com/developer-learning/night-reading-go/blob/master/content/discuss/2018-07-09-make-new-in-go.md
1. 【Golang堆栈的理解】https://segmentfault.com/a/1190000017498101
1. 【逃逸分析 编写栈分配内存建议】https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/
1. 【逃逸分析 比较简洁】https://studygolang.com/articles/17584
1. 【逃逸分析定义】https://cloud.tencent.com/developer/article/1117410
1. 【逃逸分析例子】https://my.oschina.net/renhc/blog/2222104
1. https://gocn.vip/article/355
1. 【汇编代码 传参】https://github.com/maniafish/about_go/blob/master/heap_stack.md
1. 【逃逸分析的缺陷】https://studygolang.com/articles/12396
1. 【比较好的逃逸分析的例子】http://www.agardner.me/golang/garbage/collection/gc/escape/analysis/2015/10/18/go-escape-analysis.html
  