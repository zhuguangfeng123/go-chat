#基础镜像
FROM ubuntu

#把编译后的打包进来这个镜像 放到工作目录app下
COPY gochat /app/gochat
RUN chmod +x /app/gochat
WORKDIR /app

#CMD是执行命令
ENTRYPOINT ["/app/gochat"]