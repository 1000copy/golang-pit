异常处理，在golang内到底行不行？
----------------

今天测试了下，发现尽管golang作者和社区一再啰嗦，其实 panic /recover 可以实现try catch finally。就是说，语言层面，虽然写法不太一样，但是golang支持异常处理。

只不过，标准库不这么干，因此要全盘使用
exception handle模式，需要包装标准库再使用。

golang推荐用old style编码，一定基于他们的某种信仰（或者恐惧：），我希望可以用现代的try catch，也是一样的道理。信仰在于 80386 就已经被纳入CPU指令集支持的异常，说明即使使用指令周期来衡量性能的领域，也不那么在乎异常可能引发的性能拖累，我们做app的担心个鸟。


What is cstyle error handler?
------------------
 类似C的错误处理代码
（check，check，check style 的）,吓尿了
http://weibo.com/1644105187/Bz9fQjKv6

如何封装标准库？
--------------

 babysitteraway_test.go 展示了一种做法