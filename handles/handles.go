package handles

import (
	"errors"
	"gin-project/jwt"
	"gin-project/models"
	"gin-project/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	db        *gorm.DB
	jwtSecret []byte
}

func NewUserHandler(db *gorm.DB, jwtSecret []byte) *UserHandler {
	return &UserHandler{
		db:        db,
		jwtSecret: jwtSecret}
}

// 注册
func (u *UserHandler) Register(c *gin.Context) {
	//接收注册请求
	var req response.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	if err := u.db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		response.ErrorResponse(c, http.StatusConflict, "Username already exists")
		return
	}

	// 检查邮箱是否已存在
	if err := u.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		response.ErrorResponse(c, http.StatusConflict, "Email already exists")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// 创建用户
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := u.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	response.Success(c, "success to register "+req.Username, gin.H{
		"username": req.Username,
		"email":    req.Email,
	})
}

// 登录
func (u *UserHandler) Login(c *gin.Context) {
	//接收登录请求
	var req response.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//检查用户名是否正确
	var user models.User
	if err := u.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, 401, "Invalid credentials")
			return
		}
		response.ErrorResponse(c, 401, err.Error())
		return
	}

	//检查密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		response.ErrorResponse(c, 401, "Invalid credentials")
		return
	}

	//生成token
	token, err := jwt.GenerateToken(u.jwtSecret, user.ID, user.Username)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response.Success(c, "success to login "+user.Username, gin.H{
		"token":    token,
		"userID":   user.ID,
		"username": user.Username,
		"email":    user.Email,
	})
}

// 创建文章
func (u *UserHandler) CreatePost(c *gin.Context) {
	//接收创建文章请求
	var req response.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("userID")
	// 创建文章
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID.(uint),
	}

	if err := u.db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	response.Success(c, "success to create post", gin.H{
		"title":   req.Title,
		"content": req.Content,
	})
}

// 获取文章列表
func (u *UserHandler) Postlist(c *gin.Context) {
	var posts []models.Post
	if err := u.db.Find(&posts).Error; err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//只显示其中几个信息
	postslist := make([]response.Postslist, len(posts))
	for i, p := range posts {
		postslist[i] = response.Postslist{
			ID:      p.ID,
			Title:   p.Title,
			Content: p.Content,
			UserID:  p.UserID,
		}
	}

	response.Success(c, "success to get posts", gin.H{
		"posts": postslist,
	})
}

// 获取单个文章
func (u *UserHandler) Post(c *gin.Context) {
	id := c.Query("id")

	var post models.Post
	if err := u.db.Preload("Comments").Where("id = ?", id).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, 401, "Record Not Found")
			return
		}
		response.ErrorResponse(c, 401, err.Error())
		return
	}

	response.Success(c, "success to get post", gin.H{
		"post": post,
	})
}

// 更新文章
func (u *UserHandler) Updatepost(c *gin.Context) {
	//获取更新信息
	var upd response.UpdatePostRequest
	if err := c.ShouldBindJSON(&upd); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//检查文章是否存在
	var post models.Post
	if err := u.db.Where("id = ?", upd.ID).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, 401, "Record Not Found")
			return
		}
		response.ErrorResponse(c, 401, err.Error())
		return
	}

	userID, _ := c.Get("userID")
	//检查更改人是否文章作者
	if post.UserID != userID.(uint) {
		response.ErrorResponse(c, 403, "You must be the author to edit this")
		return
	}

	//更改标题
	if upd.Title != "" {
		if err := u.db.Model(&models.Post{}).
			Where("id = ?", upd.ID).
			Update("title", upd.Title).Error; err != nil {
			response.ErrorResponse(c, 401, "Failed to update")
			return
		}
	}

	//更改内容
	if upd.Content != "" {
		if err := u.db.Model(&models.Post{}).
			Where("id = ?", upd.ID).
			Update("content", upd.Content).Error; err != nil {
			response.ErrorResponse(c, 401, "Failed to update")
			return
		}
	}

	if upd.Title == "" && upd.Content == "" {
		response.Success(c, "No changes were made", nil)
	}

	if upd.Title != "" || upd.Content != "" {
		response.Success(c, "success to update", gin.H{
			"postid":  upd.ID,
			"title":   upd.Title,
			"content": upd.Content,
		})
	}
}

// 删除文章
func (u *UserHandler) DeletePost(c *gin.Context) {
	//获取文章id，检查文章是否存在
	id := c.Query("id")

	var post models.Post
	if err := u.db.Where("id = ?", id).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, 401, "Record Not Found")
			return
		}
		response.ErrorResponse(c, 401, err.Error())
		return
	}

	userID, _ := c.Get("userID")
	//检查删除人是否文章作者
	if post.UserID != userID.(uint) {
		response.ErrorResponse(c, 403, "You must be the author to delete this")
		return
	}

	//删除文章
	if err := u.db.Where("id = ?", id).
		Delete(&models.Post{}).Error; err != nil {
		response.ErrorResponse(c, 401, "Failed to delete")
		return
	}

	response.Success(c, "success to delete", gin.H{
		"postid": id,
	})
}

// 创建评论
func (u *UserHandler) CreateComment(c *gin.Context) {
	//获取文章评论
	var req response.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//检查文章是否存在
	var post models.Post
	if err := u.db.Where("id = ?", req.PostID).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, 401, "Record Not Found")
			return
		}
		response.ErrorResponse(c, 401, err.Error())
		return
	}

	userID, _ := c.Get("userID")
	// 创建评论
	comment := models.Comment{
		Content: req.Content,
		PostID:  post.ID,
		UserID:  userID.(uint),
	}

	if err := u.db.Create(&comment).Error; err != nil {
		response.ErrorResponse(c, 500, "Failed to create comment")
		return
	}

	response.Success(c, "success to create comment", gin.H{
		"postid":  req.PostID,
		"content": req.Content,
	})
}

// 获取某篇文章的所有评论列表
func (u *UserHandler) Commentlist(c *gin.Context) {
	postid := c.Query("postid")

	//检查文章是否存在
	var post models.Post
	if err := u.db.Where("id = ?", postid).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, 401, "Post Not Found")
			return
		}
		response.ErrorResponse(c, 401, err.Error())
		return
	}

	//检查文章是否有评论
	var comment models.Comment
	if err := u.db.Where("post_id = ?", postid).First(&comment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorResponse(c, 401, "Comment Not Found")
			return
		}
		response.ErrorResponse(c, 401, err.Error())
		return
	}

	//获取所有评论
	var comments []models.Comment
	if err := u.db.Where("post_id = ?", postid).Find(&comments).Error; err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, "success to get comments of postid "+postid, gin.H{
		"comments": comments,
	})
}
