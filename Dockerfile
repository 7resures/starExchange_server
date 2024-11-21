# 使用 Golang 镜像作为构建环境
FROM golang:1.23-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 Go 模块文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制 Go 项目代码
COPY . .

# 编译 Go 应用
RUN go build -o estar main.go

# 使用更小的镜像来运行应用程序
FROM alpine:latest

# 安装 ca-certificates 以支持 HTTPS
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /app

# 复制构建阶段的二进制文件
COPY --from=builder /app/estar .

# 复制配置文件（如 setting.yaml）
COPY --from=builder /app/setting.yaml .

# 暴露应用服务的端口
EXPOSE 8080

# 启动应用
CMD ["./estar"]
