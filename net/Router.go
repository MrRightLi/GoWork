package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type User struct {
	id      int
	User_id string
	Mobile  string
	name    string
	card_id string
}

func main() {
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)


	// app run mode
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// http://localhost:2824/api/v1/userinfo/
	v1 := router.Group("/api/v1/userinfo/")
	{
		v1.GET("/:id", FetchSingleUser)
	}

	test := router.Group("/api/test/")
	{
		test.GET("/:id", FetchSingleUser)
	}

	router.Run(":2824")
}

func FetchSingleUser(c *gin.Context) {
	id := c.Param("id")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/tt?charset=utf8")
	checkErr(err)
	defer db.Close()

	checkErr(err)

	var (
		user   User
		result gin.H
	)
	row := db.QueryRow("SELECT user_id, mobile FROM user WHERE id = ?;", id)
	err = row.Scan(&user.User_id, &user.Mobile)

	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
