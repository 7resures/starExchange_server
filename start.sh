#!/bin/sh

# 确保 Redis 和 MySQL 服务可用
echo "Waiting for Redis and MySQL to be ready..."
until nc -z redis 6379; do
    echo "Waiting for Redis..."
    sleep 2
done

until nc -z mysql 3306; do
    echo "Waiting for MySQL..."
    sleep 2
done

# 执行数据库迁移
./my-go-app -db

# 启动应用程序
./my-go-app
