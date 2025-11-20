# 支付宝付款 功能使用说明

## 功能概述

完成了使用支付宝官方v3版本进行下单、查询订单、退款的功能。

## 新增文件

### API接口定义

- `api/alipay/v1/alipay.go` - 支付宝付款API接口定义

### 控制器

- `internal/controller/alipay/alipay.go` - 支付宝付款功能控制器

### 业务逻辑

- `internal/logic/alipay/alipay.go` - 支付宝付款功能业务逻辑
- `internal/logic/logic.go` - 用于在项目启动时完成支付宝客户端初始化
- `internal/service/alipay.go`  - 定义支付宝功能接口

## API接口

### 1. 创建支付订单

创建支付宝支付订单，支持电脑网站、手机网站、APP三种支付方式。

```http
POST /v1/alipay/create
Content-Type: application/json

{
    "subject": "购买会员服务",
    "out_trade_no": "20231001123456789",
    "total_amount": 99.00,
    "body": "VIP会员 - 1个月",
    "pay_type": "web"
}
```

**支付类型pay_type说明**：

- `web`：电脑网站支付，返回 HTML 表单
- `wap`：手机网站支付，返回 URL
- `app`：APP 支付，返回订单字符串

**成功响应示例**：

```http
200 OK
Content-Type: application/json

{
    "code": 0,
    "message": "success",
    "data": {
        "order_string": "https://openapi.alipay.com/gateway.do?...",
        "out_trade_no": "20231001123456789"
    }
}
```

------

### 2. 支付回调通知

支付宝异步通知接口，用户支付成功后支付宝会调用此接口通知商户。

```http
POST /v1/alipay/notify
Content-Type: application/x-www-form-urlencoded

notify_time=2023-10-01 12:00:00&
notify_type=trade_status_sync&
notify_id=2023100100222001234567890123456789&
app_id=2021001122334455&
charset=utf-8&
version=1.0&
sign_type=RSA2&
sign=xxxxx&
trade_no=2023100122001234567890123456&
out_trade_no=20231001123456789&
trade_status=TRADE_SUCCESS&
total_amount=99.00&
buyer_id=2088102177111111
```

**说明**：

- **供支付宝服务器调用，不需要前端调用**
- 必须返回 `success` 才能停止重复通知
- 需要验证签名，确保通知来源可靠

------

### 3. 查询订单状态

根据商户订单号查询支付宝交易状态。

```http
GET /v1/alipay/query?out_trade_no=20231001123456789
```

**成功响应示例（支付成功）**：

```http
200 OK
Content-Type: application/json

{
    "code": 0,
    "message": "success",
    "data": {
        "trade_no": "2023100122001234567890123456",
        "out_trade_no": "20231001123456789",
        "trade_status": "TRADE_SUCCESS",
        "total_amount": "99.00",
        "buyer_user_id": "2088102177111111"
    }
}
```

**交易状态说明**：

- `WAIT_BUYER_PAY`：等待买家付款
- `TRADE_CLOSED`：交易关闭或已退款
- `TRADE_SUCCESS`：支付成功
- `TRADE_FINISHED`：交易结束，不可退款

------

### 4. 申请退款

对已支付的订单申请退款。

```http
POST /v1/alipay/refund
Content-Type: application/json

{
    "out_trade_no": "20231001123456789",
    "refund_amount": 99.00,
    "refund_reason": "用户申请退款"
}
```

**成功响应示例**：

```http
200 OK
Content-Type: application/json

{
    "code": 0,
    "message": "success",
    "data": {
        "trade_no": "2023100122001234567890123456",
        "out_trade_no": "20231001123456789",
        "refund_amount": "99.00",
        "refund_status": "SUCCESS"
    }
}
```

## 配置说明

在 `manifest/config/config.yaml` 中添加支付宝相关配置：

```yaml
alipay:
  appId: "9021000156612338" # 配置支付宝开放平台内的自己应用的APPID
  privateKey: "" # 配置支付宝开放平台内的自己应用的私钥
  alipayPublicKey: "" # 配置支付宝开放平台内的支付宝的公钥
  notifyUrl: "https://xxxx.com/v1/alipay/notify" # 支付成功后支付宝调用的回调地址，必须是公网能访问的地址
  returnUrl: "https://xxxx.com/pay/result" # 支付成功后的跳转地址
  isSandbox: true  # 是否开启沙箱环境。开发时设置为true以进行调试，若正式上线需要修改为false
```

## 依赖

项目已添加以下依赖：

- `github.com/smartwalle/alipay/v3` - 支付宝官方库

## 常见错误码说明

| 错误码 | 描述                             |
| ------ | -------------------------------- |
| 40001  | 参数错误                         |
| 40002  | 金额必须大于 0.01                |
| 40003  | 支付类型错误，仅支持 web/wap/app |
| 40004  | 订单号不存在或已关闭             |
| 50001  | 支付服务异常                     |
| 50002  | 签名验证失败                     |

## 使用流程

1. 设置好商品内容和价格，创建支付订单
2. 用户访问支付界面，完成支付
3. 支付宝回调相关接口，继续进行相关业务处理
4. 业务处理完毕，返回`sucess`给到支付宝端，整个业务处理完毕

## 测试流程

### 环境准备

1. 拉取项目的`zgp_dev`分支。
2. **获取沙箱账号**
   - 访问：https://open.alipay.com/develop/sandbox/app
   - 登录支付宝账号
   - 获取沙箱AppID和密钥
3. **配置应用**

```yaml
   alipay:
     appId: "9021000135687177"
     privateKey: "你的应用私钥"
     alipayPublicKey: "支付宝公钥"
     notifyUrl: "https://yourdomain.com/v1/alipay/notify"
     returnUrl: "https://yourdomain.com/pay/result"
     isSandbox: true
```

3. **沙箱买家账号**（需要访问(https://open.alipay.com/develop/sandbox/account查看）

```
   账号: jmllpm6277@sandbox.com
   登录密码: 111111
   支付密码: 111111
```

4. **内网穿透工具**

   1. 官网下载：[cpolar官网-安全的内网穿透工具 | 无需公网ip | 远程访问 | 搭建网站](https://www.cpolar.com/)

   2. 登录，登录完毕后“连接您的帐户”部分的`./cpolar authtoken`命令给出了配置说明，执行这个命令，完成认证。

   3. 运行命令：`./cpolar http 8000`，命令行界面会跳转至`cpolar`窗口，`cpolar`会给出一个映射地址，例如：

    ```bash
    https://4d759dd7.r8.vip.cpolar.cn -> http://localhost:8000
    ```

   4. 添加`https://4d759dd7.r8.vip.cpolar.cn`到`config.yaml`里（**注意需要https的域名**），修改`notifyUrl`和`returnUrl`的域名即可。
   4. 修改ApiFox中的环境为上述域名。

### 测试流程

1. 调用“创建支付订单”接口，修改`out_trade_no`，得到返回的`data/order_string`的url链接，访问此链接，使用沙箱提供的账号密码完成支付
2. 在支付的任何过程中都可以调用“查询订单状态”接口，查看交易状态变动
3. 订单完成后，查看go服务日志，如果正常完成，会有`支付宝回调参数:`这样的日志内容
4. 尝试退款

## 技术支持

- 支付宝开放平台：https://open.alipay.com
- 沙箱环境：https://open.alipay.com/develop/sandbox/app
- 沙箱账号：https://open.alipay.com/develop/sandbox/account
- 错误码查询：https://open.alipay.com/api/errCheck
- API文档：https://opendocs.alipay.com/open

## 正式上线所需材料
1. 企业支付宝账号或个人支付宝账号，企业支付宝开通教程：https://opendocs.alipay.com/b/065tsj
2. 创建一个支付宝应用，网址这里需要当前支付宝账号注册的公司主体和网站备案公司主体是一致的。
3. 提交应用审核，审核通过后得到AppID，然后申请公私钥，即可。

类似流程：[网页支付功能完整流程和实现支付宝手机网站支付 官方指引 https://opendocs.alipay.com/open - 掘金](https://juejin.cn/post/7289370200561762364)


## 支付功能实现
根据需求，本产品需要具备支付功能的主要有以下3个场景：
1. 用户购买专业版/团队版套餐时，完成支付
2. 每月出账，完成超额内容和续费的支付
3. 购买图谱

### 套餐支付
#### 创建套餐购买支付订单
创建用户套餐购买的支付宝支付订单，支持电脑网站、手机网站、APP三种支付方式。
```http request
POST /v1/alipay/create-subscription-pay
Content-Type: application/json

{
    "user_plan": "professional",
    "member_count": 5,
    "pay_type": "web"
}
```
**参数说明：**
- `user_plan`：套餐类型，必填
  - `professional`：专业版套餐
  - `team`：团队版套餐
- `member_count`：团队人数，当user_plan为team时必填，最小为1
- `pay_type`：支付类型，必填
  - `web`：电脑网站支付，返回HTML表单
  - `wap`：手机网站支付，返回URL
- `app`：APP支付，返回订单字符串

**套餐价格计算逻辑：**
- 专业版：直接从系统配置中获取价格
- 团队版：根据团队人数乘以单人价格计算

价格配置位于系统配置文件中，可根据实际需求进行调整。

**订单号生成规则：**

套餐支付订单号采用以下格式：

```http request
CHIDU_SUB_{时间戳}_{用户ID}
```

例如：`CHIDU_SUB_1696123456_1001`

**支付成功后的流程：**
1. 支付宝回调系统的支付通知接口 
2. 系统验证支付状态和签名
3. 更新billing_payments表，将支付状态设置为"paid"，记录支付时间和交易ID
4. 更新billing_records表，将订单状态设置为"paid"
5. 更新user_subscriptions表，设置用户的套餐类型（user_plan）
6. 返回"success"字符串给支付宝，结束回调通知