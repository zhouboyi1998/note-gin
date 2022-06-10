FROM golang:1.18
LABEL maintainer="zhouboyi<1144188685@qq.com>"

WORKDIR /go/note-gin
COPY . /go/note-gin

# 设置环境变量
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOFLAGS=-buildvcs=false
ENV ENVCONFIG docker

# 构建项目
RUN go mod tidy
RUN go build main.go

EXPOSE 18091
ENTRYPOINT ["./main"]
