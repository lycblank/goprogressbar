# goprogressbar
golang实现的控制台进度条

#安装
``` shell
go get -u github.com/lycblank/goprogressbar
```

#使用
```go
	bar := NewProgressBar(50)
	for i:=1;i<=50;i++{
		bar.Play(int64(i))
		time.Sleep(100*time.Millisecond)
	}
	bar.Finish()
```

