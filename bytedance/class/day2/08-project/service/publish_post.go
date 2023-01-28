package service

import (
	"errors"
	"main/repository"
	"time"
	"unicode/utf8"
)

func PublishPost(topicId int64, content string) (int64, error) {
	return NewPublishPostFlow(topicId, content).Do()
}

type PublishPostFlow struct {
	content string
	topicId int64

	postId int64
}

func NewPublishPostFlow(topicId int64, content string) *PublishPostFlow {
	return &PublishPostFlow{
		content: content,
		topicId: topicId,
	}

}

func (f PublishPostFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	return 0, nil
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

func (f PublishPostFlow) checkParam() error {
	if utf8.RuneCountInString(f.content) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f PublishPostFlow) publish() error {
	post := &repository.Post{
		ParentId:   f.topicId,
		Content:    f.content,
		CreateTime: time.Now(),
	}
	if err := repository.NewPostDaoInstance().CreatePost(post); err != nil {
		return err
	}
	f.postId = post.Id
	return nil
}
