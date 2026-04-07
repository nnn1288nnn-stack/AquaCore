# Nginx Dockerfile
FROM nginx:1.25-alpine

# 移除預設配置
RUN rm /etc/nginx/conf.d/default.conf

# 複製自定義配置
COPY conf/nginx.conf /etc/nginx/nginx.conf
COPY html/ /usr/share/nginx/html/

# 暴露端口
EXPOSE 80 443

# 健康檢查
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=2 \
    CMD wget --no-verbose --tries=1 --spider http://localhost/health || exit 1

# 啟動 Nginx
CMD ["nginx", "-g", "daemon off;"]
