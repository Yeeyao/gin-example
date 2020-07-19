# go-gin-example 小结

- [参考](https://eddycjy.com/go-categories/)

## 环境

- OS: Manjaro 20.0.3 Lysia

- Kernel： x86_64 Linux 5.6.16-1-MANJARO

- CPU: Intel Core i5-4210M @ 4x 3.2GHz 

## usage

- go.mod 中 module 将被默认设置成后面的部分

### MariaDB

- 当建立数据库然后 root 默认没有密码，使用普通 linux 账户 登录 root 则需要先设置密码

- create database blog default charset utf8 COLLATE utf8_general_ci;

- ALTER USER 'root'@'localhost' IDENTIFIED BY 'password';

- UPDATE user SET password=PASSWORD('password') WHERE User='root' AND Host = 'localhost';

- systemctl start mariadb

- 这里我用 root 账户来开启 mysql 服务的

## step 2

- 调用 routers.InitRouter()

    - gin.New() 创建一个空 Engine r

    - r.Use 将一个全局的中间件绑定到 router 中

        - 中间件可以是 logger, error management 

        - 这里代码绑定了 logger, recovery

    - gin.SetMode

        - 根据输入的字符串设定 gin mode

            - 这里从 ini 中读取 mode 的字符串

    - 最后调用 r.GET 一个 router handler

        - 处理 http GET 请求，然后设置好返回的 200 以及头部

- 初始化一个 http.Server s

    - 使用 setting 的设置来设置端口

    - 使用上面创建的 Engine r 作为 router 

    - 调用 s.ListenAndServe 来监听

## step 3

- 注册开头为 /api/v1/tags 的请求头，根据不同方法执行不同处理

- 部分处理需要 orm 数据库的支持

- 需要在 创建，更新，删除，查询 后进行 orm 处理，可以自己实现 gorm 的 callbacks

- 可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm 将停止未来操作并回滚所有更改。

- 最后添加修改 tag 以及删除 tag 功能 还是通过 routers 的 tag.go 先进行验证，然后调用 models 的 tag.go 实际处理

## step 4

- 博客文章和类接口定义和编写

- gorm

    - article TagID 外键，查找

    - preload 联表查找

- 文章的相关操作

## step 5

- 添加 JWT 进行身认证

### 目前模块交互和作用

####　main

- 程序入口

	- routers.InitRouter()
	
#### routers

- 初始化　r := gin.New()

- r.Use 绑定　gin.Logger() gin.Recovery() 中间件

- r.GET　api.GetAuth 获取 jwt 数据

- r.Use(jwt.JWT()) 绑定　jwt 中间件

- 最后定义好 http 方法	

#### v1

- 分别实现 tag 以及 article 的方法

- 先利用 *gin.Context 进行数据查询验证

- 调用 models 方法进行数据库信息处理

####　model

- tag article auth 

	- 具体模块的数据库数据处理
	
- models

	- 数据库相关初始化

#### Auth

- 使用 jwt 进行账户验证

#### JWT

- 在 router.go 中被调用验证

####　http 调用

- 匹配 router.go 中定义的方法，然后调用 v1 的对应方法 

## step 6

- 添加日志

    - 发生错误等直接通过 logging 模块打印

## step 7

- golang >= 1.8 可以使用 http.Server Shutdown()

### 重启服务

- 不关闭现有链接，新的进程启动并替代旧进程，新的进程接管旧进程，

- 使用 endless 热更采用创建子进程后，将原进程退出

## step 9

- docker 需要编写 Dockerfile

- 需要注意，之前安装 Maraidb 将会导致 linux 开机启动 mysqld.service 占用 3306 端口，导致这里 docker 创建 mysql 镜像时提示端口被占用 

- 可以通过 ss -tnpl 查看当前监听的端口

- 这里提示找不到 blog数据库库 

- 需要进入 docker 然后创建数据库

- 删除过程中，如果 image 被 container 使用，则不能成功删除，只是删除 tag

- [删除参考](http://yaxin-cn.github.io/Docker/how-to-delete-a-docker-image.html)

```bash
# 启动服务
systemctl start docker
# 查看 containes
docker ps -a
# 查看 images
docker images
# 启动
docker run
# 停止后启动
docker start <container_id>
# 帮助
docker <command> --help
# 进入 bash shell 中输入 exit 离开
docker exec it <container_id> bash
# link
docker run --link <image> <image>
# 停止容器
docker stop <container_name>
# 删除容器
docker rm <container_id>
# 删除镜像
docker rmi <image_name>
```

- docker 使用 scratch

## mysql 挂载数据卷

- 数据卷被设计用来持久化数据的，其生命周期独立于容器

- 可以提供给一个或者多个容器使用

    - 数据卷可以在容器间共享和重用

    - 对数据卷的修改会马上生效

    - 数据卷的更新不会影响镜像

## step 10

- gorm callbacks

- 修改之前在 model 模块下 article.go 以及 tag.go 中两个操作函数，替换为使用回调

- 添加 gorm 的更新和删除的 callbacks

## step 11

- cron 定时

- 使用 cron 定时硬删除无效数据