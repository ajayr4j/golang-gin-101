package controller

import (
	"net/http"
	"go-projects/entity"
	"go-projects/service"
	"go-projects/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//Interface to Implement the Controllers
type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

//Allows to group/combine items of possibly different types into a single type
type controller struct {
	service service.VideoService
}

var validate *validator.Validate

//Constructor used to get the new instance of the controller
func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool",validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON((&video))
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":" Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK,"index.html", data)
}