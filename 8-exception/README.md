今天测试了下，发现尽管golang作者和社区一再啰嗦，其实 panic recover
可以实现try catch finally。只不过，标准库不这么干，因此要全盘使用
exception handle模式，需要包装标准库再使用。

 类似C的错误处理代码
（check，check，check style 的）,吓尿了
http://weibo.com/1644105187/Bz9fQjKv6
actually ...

panic/recover 就是try catch finally !