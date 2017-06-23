# 压测 ReadLine

## 结果

![](http://bigfile.b0.upaiyun.com/snapshots/ba5622a8-7fd5-42bf-9961-05f25d624015.png)

## 特殊说明

- ReadLineByDIY 的 buffer 大小是 1024，bufio 库内函数为 4096，如果一行数据过大，会出现问题。

## 一些建议

- 按行读取，如果可以确保每行字符数小于 4096 的话，可以尝试调用 ReadSlice；否则的话，调用 ReadLine，根据 isPrefix 判断。
- 自己实现 ReadLine 代码量太大，而且性能方便跟 ReadLine／ReadSlice 没有啥优化。
