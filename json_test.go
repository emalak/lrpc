package lrpc

import (
	"fmt"
	"github.com/brianvoe/gofakeit"
	"testing"
)

func TestComment_JSON(t *testing.T) {
	c := Comment{
		Id:          gofakeit.UUID(),
		UserId:      gofakeit.UUID(),
		Grade:       3,
		Attachments: []string{gofakeit.UUID(), gofakeit.UUID()},
		Text:        gofakeit.HipsterSentence(10),
		ReplyId:     gofakeit.UUID(),
		Timestamp:   213123,
	}
	fmt.Println(string(c.JSON()))
}

func TestLandmarkPreview_JSON(t *testing.T) {
	l := LandmarkPreview{
		Id:     gofakeit.UUID(),
		Liked:  gofakeit.Bool(),
		Rating: 3,
	}
	fmt.Println(string(l.JSON()))
}
