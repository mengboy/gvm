## 一个简单的go版本控制工具
### 支持os
- darwin
- linux

### 本地没有安装go
执行 ```./install.sh [version] [os] [arch] [local]``` 安装，需要提前安装wget/curl
- version go版本
- os 操作系统 linux/darwin
- arch 架构 amd64/386
- local 如果有这个参数，会从https://studygolang.com/下载
```
./install.sh 1.15.2 darwin amd64 local

[✔] wget -P ~/.gvm/ https://studygolang.com/dl/golang/go1.15.2.darwin-amd64.tar.gz
[✔] sudo tar -zxf ~/.gvm/go1.15.2.darwin-amd64.tar.gz -C /usr/local
please add 'export PATH=$PATH:/usr/local/go/bin' in your shell setting file. e.g.: .bashrc\.zshrc.
```
- 添加环境变量 'export PATH=$PATH:/usr/local/go/bin'
- 设置GOPATH

### 安装 
```
go get github.com/mengboy/gvm
```

### 使用
- 获取go官网版本列表 
```
gvm remote
```
- 获取本地已下载go列表
```
gvm list
```
- 安装某个版本
```
gvm install [version]
// 如 gvm install 1.15.2 
```

- 切换到某个版本
```
gvm use [version]
// 如 gvm use 1.15.2 
```
