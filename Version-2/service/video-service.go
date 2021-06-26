package service

import "go-projects/entity"

//Interface for the services
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

//Allows to group/combine items of possibly different types into a single type
type videoService struct {
	videos []entity.Video
}

//Constructor used to get the new instance of the videoService
func New() VideoService {
	return &videoService{}
}

//Service Implementations
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}