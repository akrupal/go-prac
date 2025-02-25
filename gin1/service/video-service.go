package service

import "gin-poc/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (v *videoService) Save(vid entity.Video) entity.Video {
	v.videos = append(v.videos, vid)
	return vid

}

func (v *videoService) FindAll() []entity.Video {
	return v.videos
}
