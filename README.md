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
  
获取文章列表  
```text
GET http://localhost:8080/api/v1/users/postlist
```
<img width="792" height="1077" alt="image" src="https://github.com/user-attachments/assets/afc0b1b8-62ff-4e27-9001-68e943df012e" />  

获取单个文章  
```text
GET http://localhost:8080/api/v1/users/post?id=1
```
<img width="783" height="831" alt="image" src="https://github.com/user-attachments/assets/f08054f3-7ba9-422e-8241-ed20152a6a8d" />  

更新文章  
```text
PUT http://localhost:8080/api/v1/users/updatepost
token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImFkbWluMiIsImV4cCI6MTc3MjEyNzk4NywibmJmIjoxNzcyMDQxNTg3LCJpYXQiOjE3NzIwNDE1ODd9.UCvYq66vLLfuW_I987lbRVUqCObRBMTvohlyU3Y-Rcg
{
    "id":1,
    "title":"第一篇文章",
    "content":"这是用户2的第一篇文章（首次更改）"
}
```
<img width="942" height="779" alt="image" src="https://github.com/user-attachments/assets/6f18b9a2-5839-4f22-a8cd-c1aa801894a4" />  

删除文章  
```text
DELETE http://localhost:8080/api/v1/users/deletepost?id=1
token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImFkbWluMiIsImV4cCI6MTc3MjEyNzk4NywibmJmIjoxNzcyMDQxNTg3LCJpYXQiOjE3NzIwNDE1ODd9.UCvYq66vLLfuW_I987lbRVUqCObRBMTvohlyU3Y-Rcg
```
<img width="864" height="666" alt="image" src="https://github.com/user-attachments/assets/b56d9827-7002-4129-a3f0-3f20fedce4ae" />  

创建评论  
```text
POST http://localhost:8080/api/v1/users/createcomment
token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6ImFkbWluMiIsImV4cCI6MTc3MjEyNzk4NywibmJmIjoxNzcyMDQxNTg3LCJpYXQiOjE3NzIwNDE1ODd9.UCvYq66vLLfuW_I987lbRVUqCObRBMTvohlyU3Y-Rcg
{
    "postid":"2",
    "content":"写得很好"
}
```
<img width="939" height="756" alt="image" src="https://github.com/user-attachments/assets/6aa24527-4234-4f8e-9edb-67f610343609" />
  


  
