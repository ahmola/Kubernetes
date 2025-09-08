package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 모델
type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Docker로 띄운 Postgres 접속 정보
	dsn := "host=postgres user=user password=pass dbname=db port=5432 sslmode=disable"

	// DB 연결
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 테이블 자동 생성
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	r := gin.Default()

	// 헬스체크
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	})

	// 유저 생성
	r.POST("/users", func(c *gin.Context) {
		var u User
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
			return
		}
		if err := db.Create(&u).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "db create failed"})
			return
		}
		c.JSON(http.StatusCreated, u)
	})

	// 유저 목록
	r.GET("/users", func(c *gin.Context) {
		var users []User
		if err := db.Order("id DESC").Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "db query failed"})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	// 단건 조회
	r.GET("/users/:id", func(c *gin.Context) {
		var u User
		if err := db.First(&u, c.Param("id")).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusOK, u)
	})

	// 서버 실행 (기본 8080)
	r.Run(":8080")
}
