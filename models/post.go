package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type Post struct {
	Id                string `orm:"pk" json:"id"`
	CreateAt          int64  `orm:"null" json:"create_at"`
	UpdateAt          int64  `orm:"null" json:"update_at"`
	DeleteAt          int64  `orm:"null" json:"delete_at"`
	UserId            string `orm:"size(50)" json:"user_id"`
	ChannelId         string `orm:"size(50)" json:"channel_id"`
	Message           string `orm:"size(50)" json:"message"`
	Type              string `orm:"size(50)" json:"type"`
	LastPictureUpdate int64  `orm:"null" json:"last_picture_update"`
}

func init() {
	orm.RegisterModel(new(Post))
}

func (u *Post) ToJson() string {
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

func GetPosts(channelId string, offset, limit int) (*[]Post, error) {
	posts := &[]Post{}
	o := orm.NewOrm()
	o.Using("default")
    if _, err := o.Raw("SELECT * FROM post where channel_id = ? limit ?, ?",channelId, offset, limit).QueryRows(posts); err != nil {
    	return nil, err
    }
	return posts, nil
}
