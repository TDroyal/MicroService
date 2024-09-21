# 微服务（暂时服务于Mistore）

### 1.功能介绍

完成如下两个微服务：

- `captcha`验证码微服务（生成验证码、验证验证码）     [生成验证码官网](https://github.com/mojocn/base64Captcha)

- 后台角色管理微服务（增删改查）



### 2.项目创建-启动流程

1.首先使用`go-micro`的脚手架创建`captcha`的服务端：  

```bash
go-micro new service captcha
```

![image-20240921165036284](./../../study/notes/images/image-20240921165036284.png)

2.定义好`.proto`文件后，执行（参考生成的Makefile文件）：  

```bash
protoc --proto_path=. --micro_out=. --go_out=:. proto/captcha.proto
```

3.再依次下载好需要依赖的包`go mod tidy`，然后再去`handler`文件夹下实现具体的方法，实现接口。

4.然后集成consul服务发现即可：

tips:在开发阶段，使用`consul agent -dev`命令即可开启一个`consul`服务发现。

集成consul服务代码如下：

```go
import(
	"github.com/go-micro/plugins/v4/registry/consul"
)
consulRegistry := consul.NewRegistry(
	registry.Addrs(fmt.Sprintf("%s:%s", ip, port)),
)

srv := micro.NewService()
srv.Init(
	micro.Name(service),
	micro.Version(version),
	micro.Registry(consulRegistry), //自己配置consul
)
```

5.实现完服务端后，再用`go-micro new client captcha`生成调用此微服务的客户端，将客户端集成到`mistore`中即可，同时将`proto`生成的`go`文件放到`mistore`上面去。

6.最后，启动微服务的服务端即可，打开`127.0.0.1:8500`即可查看微服务的状态。