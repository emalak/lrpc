package lrpc

import (
	"context"
	"github.com/oppositemc/lrpc/rpc/feed"
	landmark_storage "github.com/oppositemc/lrpc/rpc/storage"
)

type gateway interface {
	// storage (recommendation) service calls

	GetLandmark(ctx context.Context, landmarkId, userId string) (*LandmarkPreview, error)
	AddLandmark(ctx context.Context, id string) error
	LikeLandmark(ctx context.Context, userId, landmarkId string) error
	DislikeLandmark(ctx context.Context, userId, landmarkId string) error
	GetLikes(ctx context.Context, landmarkId string) (int, error)
	ViewLandmark(ctx context.Context, userId, landmarkId string) error
	RecommendLandmarks(ctx context.Context, userId string, amount int) ([]string, error)
	GetRandomFeed(ctx context.Context, amount int) ([]string, error)
	AddUser(ctx context.Context, userId string) error
	CreateComment(ctx context.Context, parentId, authorId, text string, attachments []string, rating int) error
	GetComments(ctx context.Context, landmarkId string) ([]*Comment, error)
	GetProfileComments(ctx context.Context, userId string, limit int) ([]*Comment, error)

	// feed service calls

	GetFeed(ctx context.Context, userId string, amount int) ([]string, error)
}

func (c *Client) GetLandmark(ctx context.Context, landmarkId, userId string) (*LandmarkPreview, error) {
	res, err := c.Storage.Client.GetLandmark(ctx, &landmark_storage.GetLandmarkRequest{
		LandmarkId: landmarkId,
		UserId:     userId,
	})
	if err != nil {
		return nil, err
	}
	return &LandmarkPreview{
		Id:     res.Id,
		Liked:  res.Liked,
		Rating: res.Rating,
	}, nil
}

func (c *Client) AddLandmark(ctx context.Context, id string) error {
	_, err := c.Storage.Client.AddLandmark(ctx, &landmark_storage.AddLandmarkRequest{Id: id})
	return err
}

func (c *Client) LikeLandmark(ctx context.Context, userId, landmarkId string) error {
	_, err := c.Storage.Client.LikeLandmark(ctx, &landmark_storage.LikeLandmarkRequest{
		UserId:     userId,
		LandmarkId: landmarkId,
	})
	return err
}

func (c *Client) DislikeLandmark(ctx context.Context, userId, landmarkId string) error {
	_, err := c.Storage.Client.DislikeLandmark(ctx, &landmark_storage.DislikeLandmarkRequest{
		UserId:     userId,
		LandmarkId: landmarkId,
	})
	return err
}

func (c *Client) GetLikes(ctx context.Context, landmarkId string) (int, error) {
	res, err := c.Storage.Client.GetLikes(ctx, &landmark_storage.GetLikesRequest{LandmarkId: landmarkId})
	if err != nil {
		return 0, err
	}
	return int(res.Likes), nil
}

func (c *Client) ViewLandmark(ctx context.Context, userId, landmarkId string) error {
	_, err := c.Storage.Client.ViewLandmark(ctx, &landmark_storage.ViewLandmarkRequest{
		UserId:     userId,
		LandmarkId: landmarkId,
	})
	return err
}

func (c *Client) RecommendLandmarks(ctx context.Context, userId string, amount int) ([]string, error) {
	res, err := c.Storage.Client.RecommendLandmarks(ctx, &landmark_storage.RecommendLandmarksRequest{
		UserId: userId,
		Amount: int64(amount),
	})
	if err != nil {
		return nil, err
	}
	return res.Ids, nil
}

func (c *Client) GetRandomFeed(ctx context.Context, amount int) ([]string, error) {
	res, err := c.Storage.Client.GetRandomFeed(ctx, &landmark_storage.GetRandomFeedRequest{Count: int64(amount)})
	if err != nil {
		return nil, err
	}
	return res.Ids, nil
}

func (c *Client) AddUser(ctx context.Context, userId string) error {
	_, err := c.Storage.Client.AddUser(ctx, &landmark_storage.AddUserRequest{UserId: userId})
	return err
}

func (c *Client) CreateComment(ctx context.Context, parentId, authorId, text string, attachments []string, rating int) error {
	_, err := c.Storage.Client.CreateComment(ctx, &landmark_storage.CreateCommentRequest{
		ParentId:    parentId,
		AuthorId:    authorId,
		Text:        text,
		Attachments: attachments,
		Rating:      int32(rating),
	})
	return err
}

func (c *Client) GetComments(ctx context.Context, landmarkId string) ([]*Comment, error) {
	res, err := c.Storage.Client.GetComments(ctx, &landmark_storage.GetCommentsRequest{LandmarkId: landmarkId})
	if err != nil {
		return nil, err
	}
	if res.Comments == nil || len(res.Comments) == 0 {
		return nil, nil
	}
	comments := make([]*Comment, len(res.Comments))
	for i, v := range res.Comments {
		comment := &Comment{
			Id:          v.Id,
			UserId:      v.UserId,
			Grade:       int(v.Grade),
			Attachments: v.Attachments,
			Text:        v.Text,
			ReplyId:     v.ReplyId,
			Timestamp:   int(v.Timestamp),
		}
		comments[i] = comment
	}
	return comments, nil
}

func (c *Client) GetProfileComments(ctx context.Context, userId string, limit int) ([]*Comment, error) {
	res, err := c.Storage.Client.GetProfileComments(ctx, &landmark_storage.GetProfileCommentsRequest{
		UserId: userId,
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, err
	}
	if res.Comments == nil || len(res.Comments) == 0 {
		return nil, nil
	}
	comments := make([]*Comment, len(res.Comments))
	for i, v := range res.Comments {
		comment := &Comment{
			Id:          v.Id,
			UserId:      v.UserId,
			Grade:       int(v.Grade),
			Attachments: v.Attachments,
			Text:        v.Text,
			ReplyId:     v.ReplyId,
			Timestamp:   int(v.Timestamp),
		}
		comments[i] = comment
	}
	return comments, nil
}

func (c *Client) GetFeed(ctx context.Context, userId string, amount int) ([]string, error) {
	res, err := c.Feed.Client.GetFeed(ctx, &feed.GetFeedRequest{
		UserId: userId,
		Amount: int32(amount),
	})
	if err != nil {
		return nil, err
	}
	if res.LandmarkIds != nil || len(res.LandmarkIds) == 0 {
		return nil, err
	}
	return res.LandmarkIds, nil
}
