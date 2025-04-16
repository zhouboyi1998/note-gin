#### stage.1 compile
FROM golang:1.18 AS builder

# 设置工作目录
WORKDIR /go/note-gin

# 拷贝文件
COPY . .

# 设置环境变量 (设置代理, 禁用VCS)
ENV GOPROXY https://goproxy.cn,direct
ENV GOFLAGS -buildvcs=false

# 下载依赖
RUN go mod download

# 编译 (禁用CGO, 设置操作系统, 设置CPU架构)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

#### stage.2 package
FROM alpine
LABEL maintainer="zhouboyi<1144188685@qq.com>"

# 设置工作目录
WORKDIR /go/note-gin

# 拷贝编译好的可执行文件、配置文件
COPY --from=builder /go/note-gin/main /go/note-gin
COPY ./application-docker.yaml /go/note-gin

# 设置环境变量
ENV ENVCONFIG docker

EXPOSE 18091
ENTRYPOINT ["./main"]
