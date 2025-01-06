# 使用官方 Go 镜像作为构建阶段
FROM golang:1.23.2-alpine3.19 AS builder

# 接收版本参数，默认为当前时间戳
ARG VERSION

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用 
# 使用传入的VERSION参数
RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-X admin/internal/constant.VERSION=${VERSION} -s -w" \
    -o /admin ./cmd

# 使用轻量级的 Alpine 作为运行时镜像
FROM alpine:3.19

# 安装必要的证书和时区
RUN apk --no-cache add ca-certificates tzdata

# 设置时区 
ENV TZ=Asia/Shanghai

# 创建非 root 用户
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# 切换到非 root 用户
USER appuser

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder --chown=appuser:appgroup /admin /app/admin

# 暴露端口
EXPOSE 8000

# 设置入口点
ENTRYPOINT ["/app/admin"]

CMD ["init"]
# 默认命令
CMD ["run"]