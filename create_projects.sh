#!/bin/bash

# 获取token
TOKEN=$(curl -s -X POST http://localhost:8000/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"username": "hsy199895", "password": "123456"}' | \
  grep -o '"token":"[^"]*"' | cut -d'"' -f4)

echo "Token: $TOKEN"

# 项目名称数组
PROJECTS=("测试项目1" "测试项目2" "测试项目3" "测试项目4" "测试项目5")

# 并行创建项目
for project in "${PROJECTS[@]}"; do
  curl -X POST http://localhost:8000/v1/projects \
    -H "Content-Type: application/json" \
    -H "Authorization: $TOKEN" \
    -d "{\"projectName\": \"$project\", \"graphId\": null}" &
done

# 等待所有后台任务完成
wait

echo "所有项目创建完成"
