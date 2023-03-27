package api

import (
	"log"
	"net/http"
	"server/common"
	"server/models"
	"server/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) { // ctx:上下文
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "phone number should be 11 digits"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "password should be at least 6 digits"})
		return
	}
	if len(name) == 0 {
		name = utils.RandomName(10)
	}

	// 判断手机号是否存在
	if isTelephoneExists(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "this user already exists"})
		return
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := new(models.User)
	newUser.Name = name
	newUser.Telephone = telephone
	newUser.Password = string(hashedPassword)
	DB.Create(&newUser)
	DB.Save(&newUser)

	// 返回结果
	models.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		models.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "phone number should be 11 digits")
		return
	}
	if len(password) < 6 {
		models.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "password should be at least 6 digits")
		return
	}

	// 判断手机号是否存在
	var user models.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		models.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "this user does not exists")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		models.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "this user does not exists")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generating error: %v", err)
		return
	}
	// 返回结果
	models.Success(ctx, gin.H{"token": token}, "登陆成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}

func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user *models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
