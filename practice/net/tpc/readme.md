## How to use

// demo1
```bash
# 启动 server
go run server/TCPserver.go 8001

# 多个终端分别运行
go run client/TCPclient.go localhost:8001

# 查看运行情况
netstat -anp TCP | grep 8001
```

//demo2
```bash

# start server
go run kvTCP.go 9000

# connect
nc localhost 9000

```
**input and output example**
PRINT
LOOKUP 1
Did not find key!
ADD 1 2 3 4
Add operation successful!
LOOKUP 1
{2 3 4}
ADD 4 -1 -2 -3
Add operation successful!
PRINT
key: 1 value: {2 3 4}
key: 4 value: {-1 -2 -3}
STOP


# Reference

* https://wskdsgcf.gitbook.io/mastering-go-zh-cn/13.0/13.6/13.6.1