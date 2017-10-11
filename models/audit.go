package models

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"io"
)

type Audit struct {
	Id        string `bson:"_id" json:"id"`
	CreateAt  int64  `bson:"createAt" json:"create_at"`
	UserId    string `bson:"userId" json:"user_id"`
	Action    string `bson:"action" json:"action"`
	ExtraInfo string `bson:"extraInfo" json:"extra_info"`
	IpAddress string `bson:"ipAddress" json:"ip_address"`
}

func (o *Audit) ToJson() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func AuditFromJson(data io.Reader) *Audit {
	decoder := json.NewDecoder(data)
	var o Audit
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func LoginAuditWithUserId(c *context.Context, userId, extraInfo string) {

	audit := &Audit{UserId: userId, IpAddress: c.Input.IP(), Action: c.Input.URL(), ExtraInfo: extraInfo}
    o := orm.NewOrm()
    o.Using("default")
    o.Insert(audit)
}