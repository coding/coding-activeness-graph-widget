# Coding 活跃图挂件

### 开发环境
- Go 1.4.2

### 运行方法
```
go run web.go
```

### 挂件用法
在需要显示挂件的页面里插入如下 HTML 代码：
```
<iframe src="http://localhost:3000/graph/coding" frameborder="0" scrolling="no" width="1200"></iframe>
```
将其中的 `http://localhost:3000` 换成应用实际所在的 URL。