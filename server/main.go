package main

import (
	// "net/http"
	// "server/controllers"
	"fmt"
	"math/rand"
	"net/http"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := InitDB()

	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) { // ctx:上下文
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
			name = RandomName(10)
		}

		// 判断手机号是否存在
		if isTelephoneExists(db, telephone) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "this user already exists"})
			return
		}

		// 创建用户
		newUser := new(models.User)
		newUser.Name = name
		newUser.Telephone = telephone
		newUser.Password = password
		db.Create(&newUser)
		db.Save(&newUser)

		// 返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func RandomName(n int) string {
	var name = []byte("diuqwuibcunoewncmenwphfiwehonnCIOEWNOICNWIEH")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())

	for i := range result {
		result[i] = name[rand.Intn(len(name))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	// driverName := "mysql"
	host := "localhost"
	port := "3307"
	database := "ginessential"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database, err: " + err.Error())
	}
	db.AutoMigrate(&models.User{})
	return db
}

func isTelephoneExists(db *gorm.DB, telephone string) bool {
	var user *models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

// 程序入口
// func main() {
// 	controllers.Router()
// 	http.ListenAndServe(":8080", nil)
// }
