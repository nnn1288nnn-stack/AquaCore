# 構建階段 (Build Stage)
FROM golang:1.21-alpine AS builder

# 安裝 git
RUN apk add --no-cache git

WORKDIR /build

# 複製 go.mod
COPY go.mod ./

# 下載依賴並生成 go.sum
RUN go mod download && go mod tidy

# 複製所有源代碼
COPY . .

# 編譯 Go 應用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app main.go

# 來自 Alpine 的極簡運行階段
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# 複製編譯好的二進制文件
COPY --from=builder /build/app .

# 暴露端口
EXPOSE 8080

# 健康檢查
HEALTHCHECK --interval=30s --timeout=30s --start-period=10s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# 啟動應用
CMD ["./app"]
