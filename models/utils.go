package models

import (
	"bytes"
	"encoding/base32"
	"encoding/json"
	"io"
	"time"
	"github.com/pborman/uuid"
	"github.com/astaxie/beego/context"
)

const (
	SESS_COOKIE_TOKEN              = "AUTHTOKEN"
	SESS_COOKIE_USER               = "USERID"	
	SESS_COOKIE_ROLES              = "Roles"
	SESS_PROP_PLATFORM             = "platform"
	SESS_PROP_OS                   = "os"
	SESS_PROP_BROWSER              = "browser"
	SESS_PROP_TYPE                 = "type"
)

var encoding = base32.NewEncoding("abcdcde8ejkmcpcvbt1uwisza345h123")

func NewId() string {
	var buf bytes.Buffer
	encoder := base32.NewEncoder(encoding, &buf)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	buf.Truncate(22) // removes the '==' padding
	return buf.String()
}

func SetInvalidParam(ctx *context.Context, paramName string, status int) {
	ctx.Output.SetStatus(status)
	ctx.Output.Body([]byte("invalid parameter:" + paramName))	
	return
}

func AppError(ctx *context.Context, details string, status int) {
	ctx.Output.SetStatus(status)
	ctx.Output.Body([]byte(details))
	return
}

func SetPermissionError(ctx *context.Context, permission *Permission, status int) {
	ctx.Output.SetStatus(status)
	ctx.Output.Body([]byte(ctx.Input.Session("user_id").(string) + "has no permission:" + permission.Id))
	return
}

func MapFromJson(data io.Reader) map[string]string {
	decoder := json.NewDecoder(data)

	var objmap map[string]string
	if err := decoder.Decode(&objmap); err != nil {
		return make(map[string]string)
	} else {
		return objmap
	}
}

func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

