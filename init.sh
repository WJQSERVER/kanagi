#!/bin/sh

APPLICATION=kanagi

# 检查并复制 Caddyfile
if [ ! -f /data/caddy/config/Caddyfile ]; then
    cp /data/caddy/Caddyfile /data/caddy/config/Caddyfile
fi

# 检查并复制 config.yaml
if [ ! -f /data/${APPLICATION}/config/config.yaml ]; then
    cp /data/${APPLICATION}/config.yaml /data/${APPLICATION}/config/config.yaml
fi

# 启动 Caddy
/data/caddy/caddy run --config /data/caddy/config/Caddyfile > /data/${APPLICATION}/log/caddy.log 2>&1 &

# 启动 Go 应用
/data/${APPLICATION}/${APPLICATION} > /data/${APPLICATION}/log/run.log 2>&1 &

# 保持脚本运行
while true; do
    sleep 1
done