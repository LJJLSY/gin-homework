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

**3.测试API**  
用户注册  
```text
POST http://localhost:8080/api/v1/users/register
{
    "username":"admin2",
    "password":"admin2123",
    "email":"admin2123@139.com"
}
```
<img width="945" height="765" alt="image" src="https://github.com/user-attachments/assets/bbfe1ad7-78ea-4a5f-be47-295ab42b5c0b" />  

用户登录  
```text
POST http://localhost:8080/api/v1/users/login
{
    "username":"admin2",
    "password":"admin2123"
}
```
<img width="1902" height="834" alt="image" src="https://github.com/user-attachments/assets/de67c1f6-8d2f-4f00-adb1-9bc750128ca0" />  

创建文章  
```text
POST http://localhost:8080/api/v1/users/createpost
token：eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImFkbWluMiIsImV4cCI6MTc3MjEyNjEyOSwibmJmIjoxNzcyMDM5NzI5LCJpYXQiOjE3NzIwMzk3Mjl9.ZYecjyD54jPElnDf8mvrg1cc4dpETi2S3o3lF3mWUKM
{
    "title": "第一篇文章",
    "content": "这是用户2的第一篇文章"
}
```
<img width="1335" height="287" alt="image" src="https://github.com/user-attachments/assets/2263fe0b-2308-45df-9d8f-01e8984a9f7b" />  
<img width="930" height="753" alt="image" src="https://github.com/user-attachments/assets/6e0cb5e5-4714-4f09-b3dd-9e9baa9e38f2" />


