### ucenter demo

#说明

使用`micro`开发的用户中心微服务


#功能

提供基本的注册,登陆,第三方登陆,多app登陆
提供多账号绑定


#项目结构

根目录的`main.go`是`server`端入口, 对应的`tenno.ucenter` 执行文件
`client` 包对应的是 `client`端代码,项目使用`gin`做网关路由


#部署

1.启动服务发现
` ./consul agent -dev > /out.log 2>&1 &`
2.启动server端
`./tenno.ucenter --registry_address=127.0.0.1:8500`
3.启动网关
`micro api --handler=http`
4.启动client代理
进入client目录 `./client`

