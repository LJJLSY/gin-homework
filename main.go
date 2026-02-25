package main

import (
	"gin-project/config"
	"gin-project/handles"
	"gin-project/middleware"
	"gin-project/models"
	"gin-project/response"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strconv"
)

func main() {
	//获取配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	//连接数据库
	dsn := cfg.Database.Username + ":" +
		cfg.Database.Password + "@tcp(" +
		cfg.Database.Host + ":" +
		strconv.Itoa(cfg.Database.Port) + ")/" +
		cfg.Database.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// 测试连接
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get the database: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database connection test failed: %v", err)
	}
	log.Println("Database connection successful")

	//自动迁移表
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 初始化handler
	userhandler := handles.NewUserHandler(db, []byte(cfg.JWT.Secret))

	//创建GIN引擎
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.Logger())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		response.Success(c, "healthy", gin.H{
			"status": "ok",
		})
	})

	// 公开路由
	public := r.Group("/api/v1")
	{
		public.POST("/users/register", userhandler.Register)      //注册
		public.POST("/users/login", userhandler.Login)            //登录
		public.GET("/users/postlist", userhandler.Postlist)       //获取文章列表
		public.GET("/users/post", userhandler.Post)               //获取单个文章
		public.GET("/users/commentlist", userhandler.Commentlist) //获取评论列表
	}

	// 需要认证的路由
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware([]byte(cfg.JWT.Secret)))
	{
		protected.POST("/users/createpost", userhandler.CreatePost)       //创建文章
		protected.PUT("/users/updatepost", userhandler.Updatepost)        //更新文章
		protected.DELETE("/users/deletepost", userhandler.DeletePost)     //删除文章
		protected.POST("/users/createcomment", userhandler.CreateComment) //创建评论
	}

	// 启动服务器
	addr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
