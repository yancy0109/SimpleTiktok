package repository

import (
	"os"
	"testing"
)

func TestVideoDao_PublishVideo(t *testing.T) {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	vedio := &Video{
		AuthorId: 123,
		Title:    "213213",
		PlayUrl:  "dwadw",
		CoverUrl: "sda",
	}
	vedioDao := NewVedioDaoInstance()
	vedioDao.PublishVideo(vedio)
}
