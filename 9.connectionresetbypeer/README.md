模拟 `read: connection reset by peer` 异常

异常原因：client 的接收缓冲区还有未读数据，这时当 client 关闭，server 会报此异常。与 EOF 异常的区别是，当 client 把接收缓冲区的数据都读取完毕再关闭，就是 EOF 了（参见 4.readdata）。

参考文章：[difference between socket failing with EOF and connection reset](https://stackoverflow.com/questions/15816352/difference-between-socket-failing-with-eof-and-connection-reset)

> And if the peer closes the connection while it has pending unread data in its socket receive buffer the connection will also be reset. And it is also possible for the reset to originate locally.