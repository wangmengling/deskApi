# 得到最新的 golang docker 镜像
FROM golang:latest
# 在容器内部创建一个目录来存储我们的 web 应用，接着使它成为工作目录。

# RUN mkdir -p /go/src/deskapi
# WORKDIR /go/src/deskapi
# # 复制 deskapi 目录到容器中
# COPY . /go/src/deskapi
# # 下载并安装第三方依赖到容器中
# RUN go-wrapper download
# RUN go-wrapper install
# # RUN go build .
# # 设置 PORT 环境变量
# ENV PORT 8080
# # 给主机暴露 8080 端口，这样外部网络可以访问你的应用
# EXPOSE 5000
# # 告诉 Docker 启动容器运行的命令
# CMD ["go-wrapper", "run"]
# # CMD [ "run"]

RUN mkdir -p /go/src/deskapi
WORKDIR /go/src/deskapi
COPY . /go/src/deskapi
CMD ["go-wrapper", "run"]
ONBUILD COPY . /go/src/deskapi
ONBUILD RUN go-wrapper download
ONBUILD RUN go-wrapper install
RUN go get github.com/gin-gonic/gin
RUN go get -d -v
RUN go install -v