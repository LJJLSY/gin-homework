**<h1>项目结构</h1>** 
<!-- ------------------------------------------------------- --> 
```text
gin-homework/  
├── main.go              # 程序入口
├── config/              # 配置管理
│   └── config.go
│   └── config.yaml
├── handles/             # 业务逻辑层
│   └── handles.go
├── jwt/                 # token逻辑层
│   ├── jwt.go
├── middleware/          # 中间件
│   ├── AuthMiddleware.go
│   ├── logger.go
├── models/              # 数据模型
│   └── models.go
├── response/            # 请求和响应
    └── response.go
