# blogserver

### 简介

简单的博客系统服务端程序

### 使用Makefile构建

等待完善

### 使用 Dockerfile构建docker镜像 ( 目前已经使用cobra 改造入口文件兼容 多命令，等待完成后将重新构建镜像 )

```shell
# 构建镜像
docker build -t hayuzi/blogserver:v1.0.0 .

# 删除镜像
docker rmi hayuzi/blogserver:v1.0.0  

# 执行（执行的时候需要挂载配置文件进去替换）
docker run --name=blogserver -p 8081:8081 hayuzi/blogserver:v1.0.0

# 挂载相关目录到容器中并运行
docker run --name=blogserver -p 8081:8081 -v /your/config/path:/data/blog/config  -v /your/storage/path:/data/blog/storage hayuzi/blogserver:v1.0.0

```