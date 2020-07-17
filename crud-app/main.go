package main
import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

// global variables inside package "main"
var db *gorm.DB = nil
var err error // err is the var name, error is the var type

func main() {
	r := gin.Default()

	// connection to the DB
	db, err = gorm.Open("mysql", "root:123@/go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Error Connecting Database")
	}
	defer db.Close()

	// auto migrate
	db.AutoMigrate(&Post{})

	// enabling watching the queries in the terminal
	db.LogMode(true)

	// routes
	r.GET("/posts", getAllPosts)
	r.GET("/posts/:id", getSinglePost)
	r.POST("/posts", createPost)
	r.PATCH("/posts/:id", updatePost)
	r.DELETE("/posts/:id", deletePost)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

