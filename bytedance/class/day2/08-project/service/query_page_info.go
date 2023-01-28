package service

import (
	"errors"
	"main/repository"
	"sync"
)

type PageInfo struct {
	TopicInfo *repository.Topic
	PostList  []*repository.Post
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

func NewQueryPageInfoFlow(topicId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topicId,
	}
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic *repository.Topic
	posts []*repository.Post
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	// 获取topic信息
	var wg sync.WaitGroup
	wg.Add(2)
	var topicErr, postErr error
	go func() {
		defer wg.Done()
		topic, err := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		if err != nil {
			topicErr = err
			return
		}
		f.topic = topic
	}()
	// 获取post信息
	go func() {
		defer wg.Done()
		posts, err := repository.NewPostDaoInstance().QueryPostById(f.topicId)
		if err != nil {
			postErr = err
			return
		}
		f.posts = posts
	}()
	wg.Wait()
	if topicErr != nil {
		return topicErr
	}
	if postErr != nil {
		return postErr
	}

	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	postList := make([]*repository.Post, 0)
	for _, post := range f.posts {
		postList = append(postList, post)
	}
	f.pageInfo = &PageInfo{
		TopicInfo: f.topic,
		PostList:  f.posts,
	}
	return nil
}
