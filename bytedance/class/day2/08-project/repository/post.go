package repository

import (
	"sync"
	"time"
)

type Post struct {
	Id         int64     `gorm:"column:id"`
	ParentId   int64     `gorm:"parent_id"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}

type PostDao struct {
}

var (
	postDao  *PostDao
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostById(id int64) ([]*Post, error) {
	return postIndexMap[id], nil
}

func (d *PostDao) CreatePost(*Post) error {
	return nil
}
