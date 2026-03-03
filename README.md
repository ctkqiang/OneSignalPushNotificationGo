# 推送通知服务

## 项目简介

推送通知服务是一个基于 Go 语言开发的后端服务，用于发送各种类型的推送通知。该服务使用 OneSignal 作为推送服务提供商，支持发送文本通知和文本+图片通知。

## 功能特性

- 支持发送文本推送通知
- 支持发送文本和图片推送通知
- 内置安全中间件，包括：
  - 速率限制
  - 恶意关键词检测（防止 SQL 注入等攻击）
  - 请求体大小限制
  - 安全 HTTP 头部设置
- 自动生成 Swagger API 文档
- 详细的序列图文档

## 技术栈

- **后端框架**：Gin
- **推送服务**：OneSignal
- **API 文档**：Swagger
- **文档工具**：PlantUML
- **语言**：Go

## 项目结构

```
PushNoification/
├── docs/               # 文档目录
│   ├── docs.go         # Swagger 文档生成文件
│   ├── swagger.json    # Swagger JSON 文档
│   ├── swagger.yaml    # Swagger YAML 文档
│   └── sequence.puml   # 序列图文档
├── internal/           # 内部代码
│   ├── api/            # API 相关代码
│   │   ├── handler/    # 请求处理器
│   │   ├── middleware/ # 中间件
│   │   └── routes/     # 路由定义
│   ├── config/         # 配置文件
│   ├── structure/      # 数据结构
│   └── utilities/      # 工具函数
├── main.go             # 主入口文件
├── go.mod              # Go 模块文件
└── README.md           # 项目说明文档
```

## 安装和运行

### 前提条件

- Go 1.18+ 环境
- OneSignal 账号和应用凭据

### 安装步骤

1. 克隆项目代码

2. 安装依赖
   ```bash
   go mod tidy
   ```

3. 配置环境变量
   - 设置 OneSignal 应用 ID 和 API Key

4. 运行服务
   ```bash
   go run main.go
   ```

## API 文档

服务启动后，可以通过以下地址访问 Swagger API 文档：

```
http://localhost:8080/swagger/index.html
```

### 主要 API 端点

#### 发送文本推送通知
- **URL**: `/push/text`
- **方法**: POST
- **请求体**:
  ```json
  {
    "title": "通知标题",
    "message": "通知内容"
  }
  ```
- **响应**:
  ```json
  {
    "status": "success",
    "message": "通知发送成功",
    "data": { /* OneSignal 响应数据 */ }
  }
  ```

#### 发送文本和图片推送通知
- **URL**: `/push/text-image`
- **方法**: POST
- **请求体**:
  ```json
  {
    "title": "通知标题",
    "message": "通知内容",
    "image_url": "https://example.com/image.jpg"
  }
  ```
- **响应**:
  ```json
  {
    "status": "success",
    "message": "通知发送成功",
    "data": { /* OneSignal 响应数据 */ }
  }
  ```

## 安全特性

1. **速率限制**：每 IP 每分钟最多 60 个请求
2. **恶意关键词检测**：防止 SQL 注入等攻击
3. **请求体大小限制**：最大 1MB
4. **安全 HTTP 头部**：设置了以下安全头部
   - X-Content-Type-Options: nosniff
   - X-Frame-Options: DENY
   - X-XSS-Protection: 1; mode=block
   - Strict-Transport-Security: max-age=31536000; includeSubDomains

## 示例使用

### 发送文本通知

```bash
curl -X POST http://localhost:8080/push/text \
  -H "Content-Type: application/json" \
  -d '{"title": "测试通知", "message": "这是一条测试通知"}'
```

### 发送文本和图片通知

```bash
curl -X POST http://localhost:8080/push/text-image \
  -H "Content-Type: application/json" \
  -d '{"title": "测试通知", "message": "这是一条包含图片的测试通知", "image_url": "https://example.com/image.jpg"}'
```

## 序列图

项目包含以下序列图：

- **发送文本推送通知流程**：展示了从客户端发起请求到通知发送完成的完整流程
- **发送文本和图片推送通知流程**：展示了包含图片的通知发送流程

序列图文件位于 `docs/sequence.puml`，可以使用 PlantUML 工具查看。

## 贡献指南

1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

本项目使用 MIT 许可证。详见 LICENSE 文件。
