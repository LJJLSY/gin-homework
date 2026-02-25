# 项目结构
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
```
# 运行环境
mysql  
# 快速开始
<!-- ------------------------------------------------------- -->
**1.安装依赖**  
go mod tidy  
go get xxx  

**2.运行项目**  
go run main.go  

服务器将在 http://0.0.0.0:8080 启动。
