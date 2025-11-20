# 重置用户配额和出账功能说明文档

此文档是对用户配置每月重置和每月出账功能实现的说明

## 方案设计

使用Go Frame的`gcron`创建一个定时任务`service/subscription.go`，设置每天2点重置用户配额和出账；

## 涉及表单
1. 重置：`user_subscriptions`，每月重置还会刷新`quota_reset_time`字段
2. 记录：`billing_records`
3. 超额记录：`billing_overages`
