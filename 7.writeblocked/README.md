实现[文中](https://time.geekbang.org/column/article/481908)“写阻塞”和“写入部分数据”的情况

“写入部分数据”需要在分别启动服务端和客户端后，当客户端写入完第一批数据阻塞后，杀掉服务端进程，客户端会报 `write: broken pipe`的错误