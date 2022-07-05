FROM alpine
LABEL maintainer="zhouboyi<1144188685@qq.com>"

WORKDIR /go/note-gin
COPY ./main /go/note-gin
COPY ./application-docker.yaml /go/note-gin

# 设置环境变量
ENV ENVCONFIG docker

EXPOSE 18091
ENTRYPOINT ["./main"]
