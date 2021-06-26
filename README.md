# golang-gin-101

## Installation

	- Install: https://tecadmin.net/install-go-on-ubuntu/

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
 
 Checkout localhost:8080/ping
 
