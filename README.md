# Wlog

> golang格式化日志打印。
>
------
### 使用步骤
1. 获取
```go
go get github.com/wms3001/wlog
```
2. 控制台输出日志
```go
var wlog = Wlog{}
wlog.Show("This is Test!!")
```
控制台打印
```go
2023/03/29 22:08:24.221888 D:/GolandProjects/yanwen/yanwen_test.go:26: Message:1680098904221
2023/03/29 22:08:24.231198 D:/GolandProjects/yanwen/yanwen_test.go:27: Message:2914e982da11d2587a74333f4be179bf
```
