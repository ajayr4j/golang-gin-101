# golang-gin-101

## Installation

	- https://tecadmin.net/install-go-on-ubuntu/
	- go get -u github.com/gin-gonic/gin

## Initial Setup

	- go mod init your-project-name
	- cd your-project-name
  
## Sample Go Program
  - vi example.go
  - Paste the following code:
 
        package main
        import "github.com/gin-gonic/gin"

        func main() {
          r := gin.Default()
          r.GET("/ping", func(c *gin.Context) {
            c.JSON(200, gin.H{
              "message": "pong",
            })
          })
          r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
        }
 
	- go mod tidy
  - go run example.go
 
 Done! 
 
 Visit localhost:8080/ping
 
 ### If you want to learn the API building further, I followed the below mentioned playlist and updated the same as versions here, check it out.
 	- https://www.youtube.com/watch?v=qR0WnWL2o1Q&list=PL3eAkoh7fypr8zrkiygiY1e9osoqjoV9w
 
 
