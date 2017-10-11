package models

import (
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type Post struct {
	Id                 string                 `orm:"pk"` 
	CreateAt           int64                  `orm:"null"`
	UpdateAt           int64                  `orm:"null"`
	EditAt             int64                  `orm:"null"`
	DeleteAt           int64                  `orm:"null"`
	UserId             string                 `orm:"size(50)"`
	ChannelId          string                 `orm:"size(50)"`
	Message            string                 `orm:"size(50)"`
	Type               string                 `orm:"size(50)"`
	PendingPostId      string                 `orm:"size(50)"`
	LastPictureUpdate  int64                  `orm:"null"`
}

func init() {
    orm.RegisterModel(new(Post))
}

func PostToJson(u *Post) string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func PostsToJson(u *[]Post) string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}



func (o *Post) PreSave() {
	if o.Id == "" {
		o.Id = NewId()
	}

	if o.CreateAt == 0 {
		o.CreateAt = GetMillis()
	}

	o.UpdateAt = o.CreateAt
}

func CreatePost(post *Post) (*Post, error) {
	post.PreSave()
    o := orm.NewOrm()
    o.Using("default")
	if _, err := o.Insert(post); err != nil {
		return nil, err
	}
	return post, nil
}

func GetPost(channelId string, page, perpage int) ( *[]Post, error) {
	posts := &[]Post{}
    o := orm.NewOrm()
    o.Using("default")

	return posts,nil
}