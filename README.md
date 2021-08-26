# cdnParse

### 说明
基于Cname判断CDN

### 运行
##### 使用cdn
```
./main -u www.baidu.com    
```
```
map[link:https://lab.skk.moe/cdn name:百度旗下业务地域负载均衡系统]
true
```
##### 未使用cdn
```
./main -u www.tuya.com        
```
```
false
```

### 编译
##### 打包linux可执行文件
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

##### 打包windows可执行文件
```
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

##### 打包mac可执行文件
```
$ go build main.go
```
