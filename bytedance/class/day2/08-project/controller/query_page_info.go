package controller

import (
	"main/service"
	"strconv"
)

func PublishPost(topicIdStr, content string) *PageData {
	topic, _ := strconv.ParseInt(topicIdStr, 10, 64)
	postId, err := service.PublishPost(topic, content)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"post_id": postId,
		},
	}
}
