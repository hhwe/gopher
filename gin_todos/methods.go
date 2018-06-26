package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
	"os"
	"io"
)

func main() {
	// Creates a r without any middleware by default
	//r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	//r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	//r.Use(gin.Recovery())

	//// Per route middleware, you can add as many as you desire.
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	//
	//// Authorization group
	//// authorized := r.Group("/", AuthRequired())
	//// exactly the same as:
	//authorized := r.Group("/")
	//// per group middleware! in this case we use the custom created
	//// AuthRequired() middleware just in the "authorized" group.
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//	// nested group
	//	testing := authorized.Group("testing")
	//	testing.GET("/analytics", analyticsEndpoint)
	//}

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other rs match /user/john, it will redirect to /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	r.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		// c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	//// Simple group: v1
	//v1 := r.Group("/v1")
	//{
	//	v1.POST("/login", loginEndpoint)
	//	v1.POST("/submit", submitEndpoint)
	//	v1.POST("/read", readEndpoint)
	//}
	//
	//// Simple group: v2
	//v2 := r.Group("/v2")
	//{
	//	v2.POST("/login", loginEndpoint)
	//	v2.POST("/submit", submitEndpoint)
	//	v2.POST("/read", readEndpoint)
	//}


	r.Run(":8080")
}
