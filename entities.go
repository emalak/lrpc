package lrpc

import (
	"github.com/valyala/fastjson"
)

type dto interface {
	JSON() []byte
}

type Comment struct {
	Id          string
	UserId      string
	Grade       int
	Attachments []string
	Text        string
	ReplyId     string
	Timestamp   int
}

func (c Comment) JSON() []byte {
	var p fastjson.Arena
	obj := p.NewObject()

	obj.Set("Id", p.NewString(c.Id))
	obj.Set("UserId", p.NewString(c.UserId))
	obj.Set("Text", p.NewString(c.Text))
	obj.Set("ReplyId", p.NewString(c.ReplyId))

	obj.Set("Grade", p.NewNumberInt(c.Grade))
	obj.Set("Timestamp", p.NewNumberInt(c.Timestamp))

	attachments := p.NewArray()
	for _, attachment := range c.Attachments {
		attachments.SetArrayItem(len(attachments.GetArray()), p.NewString(attachment))
	}
	obj.Set("Attachments", attachments)

	return obj.MarshalTo(nil)
}

type LandmarkPreview struct {
	Id     string
	Liked  bool
	Rating float32
}

func (l LandmarkPreview) JSON() []byte {
	var p fastjson.Arena
	obj := p.NewObject()

	obj.Set("Id", p.NewString(l.Id))
	if l.Liked {
		obj.Set("Liked", p.NewTrue())
	} else {
		obj.Set("Liked", p.NewFalse())
	}
	obj.Set("Rating", p.NewNumberFloat64(float64(l.Rating)))

	return obj.MarshalTo(nil)
}
