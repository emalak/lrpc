package lrpc

import (
	"context"
	"github.com/google/uuid"
	feed "github.com/oppositemc/lrpc/rpc/feed"
	storage "github.com/oppositemc/lrpc/rpc/storage"
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
	GetComments(ctx context.Context, landmarkId string, limit int) ([]*Comment, error)
	GetProfileComments(ctx context.Context, userId string, limit int) ([]*Comment, error)
	GetFavouriteLandmarks(ctx context.Context, userId string) ([]uuid.UUID, error)
	GetLikesAmount(ctx context.Context, userId string) (int, error)
	IsLiked(ctx context.Context, landmarkId string, userId string) (bool, error)
	// feed service calls

	GetFeed(ctx context.Context, userId string, amount int) ([]string, error)
}

func (c *Client) GetLandmark(ctx context.Context, landmarkId, userId string) (*LandmarkPreview, error) {
	res, err := c.Storage.Client.GetLandmark(ctx, &storage.GetLandmarkRequest{
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

func (c *Client) AddLandmark(ctx context.Context, id string, score float32) error {
	_, err := c.Storage.Client.AddLandmark(ctx, &storage.AddLandmarkRequest{Id: id, Score: score})
	return err
}

func (c *Client) LikeLandmark(ctx context.Context, userId, landmarkId string) error {
	_, err := c.Storage.Client.LikeLandmark(ctx, &storage.LikeLandmarkRequest{
		UserId:     userId,
		LandmarkId: landmarkId,
	})
	return err
}

func (c *Client) DislikeLandmark(ctx context.Context, userId, landmarkId string) error {
	_, err := c.Storage.Client.DislikeLandmark(ctx, &storage.DislikeLandmarkRequest{
		UserId:     userId,
		LandmarkId: landmarkId,
	})
	return err
}

func (c *Client) GetLikes(ctx context.Context, landmarkId string) (int, error) {
	res, err := c.Storage.Client.GetLikes(ctx, &storage.GetLikesRequest{LandmarkId: landmarkId})
	if err != nil {
		return 0, err
	}
	return int(res.Likes), nil
}

func (c *Client) ViewLandmark(ctx context.Context, userId, landmarkId string) error {
	_, err := c.Storage.Client.ViewLandmark(ctx, &storage.ViewLandmarkRequest{
		UserId:     userId,
		LandmarkId: landmarkId,
	})
	return err
}

func (c *Client) RecommendLandmarks(ctx context.Context, userId string, amount int) ([]string, error) {
	res, err := c.Storage.Client.RecommendLandmarks(ctx, &storage.RecommendLandmarksRequest{
		UserId: userId,
		Amount: int64(amount),
	})
	if err != nil {
		return nil, err
	}
	return res.Ids, nil
}

func (c *Client) GetRandomFeed(ctx context.Context, amount int) ([]string, error) {
	res, err := c.Storage.Client.GetRandomFeed(ctx, &storage.GetRandomFeedRequest{Count: int64(amount)})
	if err != nil {
		return nil, err
	}
	return res.Ids, nil
}

func (c *Client) AddUser(ctx context.Context, userId string) error {
	_, err := c.Storage.Client.AddUser(ctx, &storage.AddUserRequest{UserId: userId})
	return err
}

func (c *Client) CreateComment(ctx context.Context, parentId, authorId, text string, attachments []string, rating int) error {
	_, err := c.Storage.Client.CreateComment(ctx, &storage.CreateCommentRequest{
		ParentId:    parentId,
		AuthorId:    authorId,
		Text:        text,
		Attachments: attachments,
		Rating:      int32(rating),
	})
	return err
}

func (c *Client) GetComments(ctx context.Context, landmarkId string, limit, offset int) ([]*Comment, error) {
	res, err := c.Storage.Client.GetComments(ctx, &storage.GetCommentsRequest{
		LandmarkId: landmarkId,
		Limit:      int32(limit),
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
			ParentId:    v.ParentId,
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

func (c *Client) GetProfileComments(ctx context.Context, userId string, limit, offset int) ([]*Comment, error) {
	res, err := c.Storage.Client.GetProfileComments(ctx, &storage.GetProfileCommentsRequest{
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
			ParentId:    v.ParentId,
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
	return res.LandmarkIds, nil
}

func (c *Client) GetFavouriteLandmarks(ctx context.Context, userId string, limit, offset int) ([]uuid.UUID, error) {
	res, err := c.Storage.Client.GetFavouriteLandmarks(ctx, &storage.GetFavouriteLandmarksRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return uuidArray(res.Ids)
}

func (c *Client) GetLikesAmount(ctx context.Context, userId string) (int, error) {
	res, err := c.Storage.Client.GetLikesAmount(ctx, &storage.GetLikesAmountRequest{UserId: userId})
	if err != nil {
		return 0, err
	}
	return int(res.Count), nil
}

func (c *Client) IsLiked(ctx context.Context, landmarkId string, userId string) (bool, error) {
	res, err := c.GetLandmark(ctx, landmarkId, userId)
	if err != nil {
		return false, err
	}
	return res.Liked, nil
}

func (c *Client) GetUserTags(ctx context.Context, userId string) ([]string, error) {
	res, err := c.Storage.Client.GetUserTags(ctx, &storage.GetUserTagsRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return res.Ids, nil
}

func (c *Client) GetLandmarkTags(ctx context.Context, landmarkId string) ([]string, error) {
	res, err := c.Storage.Client.GetLandmarkTags(ctx, &storage.GetLandmarkTagsRequest{LandmarkId: landmarkId})
	if err != nil {
		return nil, err
	}
	return res.Ids, nil
}

func (c *Client) CountReviews(ctx context.Context, userId string) (int, error) {
	res, err := c.Storage.Client.CountReviews(ctx, &storage.CountReviewsRequest{UserId: userId})
	if err != nil {
		return 0, err
	}
	return int(res.Count), nil
}

func (c *Client) ConnectTags(ctx context.Context, id1, id2 string, score float64) error {
	_, err := c.Storage.Client.ConnectTags(ctx, &storage.ConnectTagsRequest{
		Id1:   id1,
		Id2:   id2,
		Score: float32(score),
	})
	return err
}

func (c *Client) DisconnectTags(ctx context.Context, id1, id2 string) error {
	_, err := c.Storage.Client.DisconnectTags(ctx, &storage.DisconnectTagsRequest{
		Id1: id1,
		Id2: id2,
	})
	return err
}

func (c *Client) DeleteTag(ctx context.Context, id string) error {
	_, err := c.Storage.Client.DeleteTag(ctx, &storage.DeleteTagRequest{Id: id})
	return err
}

func (c *Client) AddLandmarkTag(ctx context.Context, landmarkId, tagId string) error {
	_, err := c.Storage.Client.AddLandmarkTag(ctx, &storage.AddLandmarkTagRequest{
		LandmarkId: landmarkId,
		TagId:      tagId,
	})
	return err
}

func (c *Client) DeleteLandmarkTag(ctx context.Context, landmarkId, tagId string) error {
	_, err := c.Storage.Client.RemoveLandmarkTag(ctx, &storage.RemoveLandmarkTagRequest{
		LandmarkId: landmarkId,
		TagId:      tagId,
	})
	return err
}

func (c *Client) GetConnectedTags(ctx context.Context, tagId string) ([]TagWithScore, error) {
	res, err := c.Storage.Client.GetConnectedTags(ctx, &storage.GetConnectedTagsRequest{TagId: tagId})
	if err != nil {
		return nil, err
	}
	tags := make([]TagWithScore, len(res.Tags))
	for i, v := range res.Tags {
		tags[i] = TagWithScore{
			Id:    v.Id,
			Score: float64(v.Score),
		}
	}
	return tags, nil
}

func (c *Client) CountFriends(ctx context.Context, userId string) (int, error) {
	res, err := c.Storage.Client.CountFriends(ctx, &storage.CountFriendsRequest{UserId: userId})
	return int(res.Count), err
}

func (c *Client) CreateTag(ctx context.Context, id string) error {
	_, err := c.Storage.Client.CreateTag(ctx, &storage.CreateTagRequest{Id: id})
	return err
}

func (c *Client) SetUserTag(ctx context.Context, userId string, tagId string) error {
	_, err := c.Storage.Client.SetUserTag(ctx, &storage.SetUserTagRequest{
		UserId: userId,
		TagId:  tagId,
	})
	return err
}

func (c *Client) DeleteUserTag(ctx context.Context, userId string, tagId string) error {
	_, err := c.Storage.Client.DeleteUserTag(ctx, &storage.DeleteUserTagRequest{
		UserId: userId,
		TagId:  tagId,
	})
	return err
}

func (c *Client) DeleteComment(ctx context.Context, userId, commentId string) error {
	_, err := c.Storage.Client.DeleteComment(ctx, &storage.DeleteCommentRequest{
		UserId:    userId,
		CommentId: commentId,
	})
	return err
}

func (c *Client) EditComment(ctx context.Context, userId, commentId, text string) error {
	_, err := c.Storage.Client.EditComment(ctx, &storage.EditCommentRequest{
		UserId:    userId,
		CommentId: commentId,
		Text:      text,
	})
	return err
}

func (c *Client) GetFriends(ctx context.Context, userId string) ([]string, error) {
	res, err := c.Storage.Client.GetFriends(ctx, &storage.GetFriendsRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return res.Ids, err
}

func (c *Client) AddFriend(ctx context.Context, sender, receiver string) error {
	_, err := c.Storage.Client.AddFriend(ctx, &storage.AddFriendRequest{Sender: sender, Receiver: receiver})
	return err
}

func (c *Client) DeleteFriend(ctx context.Context, sender, receiver string) error {
	_, err := c.Storage.Client.DeleteFriend(ctx, &storage.DeleteFriendRequest{
		Sender:   sender,
		Receiver: receiver,
	})
	return err
}

func (c *Client) IsFriend(ctx context.Context, user1, user2 string) (bool, error) {
	res, err := c.Storage.Client.IsFriend(ctx, &storage.IsFriendRequest{
		User1: user1,
		User2: user2,
	})
	if err != nil {
		return false, err
	}
	return res.IsFriend, err
}

func (c *Client) GetLandmarksByTag(ctx context.Context, tagId string, limit, offset int) ([]string, error) {
	res, err := c.Storage.Client.GetLandmarksByTag(ctx, &storage.GetLandmarksByTagRequest{
		TagId:  tagId,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}
	return res.Ids, nil
}

func (c *Client) GetLandmarksFiltered(ctx context.Context, include, exclude []string, limit, offset int) ([]string, error) {
	res, err := c.Storage.Client.GetLandmarksFiltered(ctx, &storage.GetLandmarksFilteredRequest{
		Include: include,
		Exclude: exclude,
		Offset:  int32(offset),
		Limit:   int32(limit),
	})
	if err != nil {
		return nil, err
	}
	return res.Ids, nil
}
