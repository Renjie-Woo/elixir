# 进度条：progress bar

## 演示
### 代码块
```golang
func main()  {
    fmt.Println("this is a progress bar")
	var title = "demo"
	var current = 12
	var total = 100
	var unit = "Mib"
	var newBar = progressBar.NewProgressBar(title,current, total)
	newBar.SetUnit(unit)
	newBar.SetGraph(">")
	for i := current; i <= total; i ++ {
		newBar.Run(i)
		time.Sleep(time.Second / 100)
	}
}
```
### 效果
> this is a progress bar    
> demo: [>>>>>>>>>>>>>>>>]  100.00%  100.0 Mib/100.0 Mib