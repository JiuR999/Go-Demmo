# 使用官方基础镜像
FROM golang:1.22 
# 设置工作目录
WORKDIR /app

# 复制当前目录内容到工作目录
COPY . .

# 编译Go程序
RUN go build -o myapp

# 暴露端口
EXPOSE 8080

# 运行编译后的二进制文件
CMD ["./myapp"]
