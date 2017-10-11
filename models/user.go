package models 

import (
    "encoding/json"
 	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation" 	
 	"golang.org/x/crypto/bcrypt"
 	"errors"
 	"fmt"
 	"strings"
)

type User struct {
	Id                 string    `orm:"pk"`
	CreateAt           int64     `orm:"null"`
	UpdateAt           int64     `orm:"null"`
	DeleteAt           int64     `orm:"null"` 
	Username           string    `orm:"size(50); unique" valid:"Required"`
	Password           string    `orm:"size(100)"`
	AuthData           *string   `orm:"null"`
	AuthService        string    `orm:"null"`
	Email              string    `orm:"null" valid:"Email; MaxSize(100)"`
	EmailVerified      bool      `orm:"null"`
	Age    			   int       `orm:"null" valid:"Range(1, 140)"`
	Nickname           string    `orm:"size(50)"`
	Position           string    `orm:"size(50)"`
	Roles              string    `orm:"size(50)"`
	PhoneNum		   string    `orm:"size(50)" valid:"Mobile"`
	LastPasswordUpdate int64     `orm:"null"`
	LastPictureUpdate  int64     `orm:"null"`
	FailedAttempts     int       `orm:"null"`
}

const (
	SESSION_PROP_PLATFORM             = "platform"
	SESSION_PROP_OS                   = "os"
	SESSION_PROP_BROWSER              = "browser"
)

func init() {
    orm.RegisterModel(new(User))
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *User) Valid(v *validation.Validation) {
    if strings.Index(u.Username, "admin") != -1 {
        // 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
        v.SetError("Name", "名称里不能含有 admin")
    }
}

func (u *User) GetRoles() []string {
	return strings.Fields(u.Roles)
}

func (u *User) ToJson() string {
	b, err := json.Marshal(u)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func CreateUser(user *User) (*User, error) {
	user.PreSave()
	valid := validation.Validation{}
    b, err := valid.Valid(user)
    if err != nil {
		return nil, err
    }
    if !b {
        // validation does not pass
        // blabla...
        validError := ""
        for _, err := range valid.Errors {
            tmp := fmt.Sprintln(err.Key, err.Message)
            validError = validError + tmp
        }
        return nil, errors.New(validError)
    }	
    o := orm.NewOrm()
    o.Using("default")
    var count int
    err = o.Raw("SELECT COUNT(*) FROM user").QueryRow(&count)
	if count == 0 {
		user.Roles = "system_admin"
	} else {
		user.Roles = "normal_user"
	}

	if _, err := o.Insert(user); err != nil {
		return nil, err
	}
	return user, nil
}


func GetUser(userId string) (*User, error) {
	user := &User{}
	o := orm.NewOrm()
	
	if err := o.QueryTable("user").Filter("Id", userId).One(user); err != nil {
		return nil, errors.New("read user failed:" + err.Error())
	}

	return user, nil
}

func AuthenticateUserForLogin(username , password string) (*User, error) {
	user := &User{}
	o := orm.NewOrm()
	
	if err := o.QueryTable("user").Filter("Username", username).One(user); err != nil {
		return nil, errors.New("read user failed:" + err.Error())
	}

	if !ComparePassword(user.Password, password) {
		user.FailedAttempts = user.FailedAttempts+1
		if _, err := o.Update(user); err != nil {
			return nil, errors.New("Update password attempts failed" + err.Error())
		}
		return nil, errors.New("Wrong password")
	} else {
		user.FailedAttempts = 0
		if _, err := o.Update(user); err != nil {
			return nil, errors.New("Update password attempts failed" + err.Error())
		}
		return user, nil
	}
}

func ComparePassword(hash string, password string) bool {

	if len(password) == 0 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}


func (u *User) PreSave() {
	if u.Id == "" {
		u.Id = NewId()
	}

	if u.Username == "" {
		u.Username = NewId()
	}

	if u.AuthData != nil && *u.AuthData == "" {
		u.AuthData = nil
	}

	u.Username = strings.ToLower(u.Username)
	u.Email = strings.ToLower(u.Email)

	u.CreateAt = GetMillis()
	u.UpdateAt = u.CreateAt

	u.LastPasswordUpdate = u.CreateAt

	if len(u.Password) > 0 {
		u.Password = HashPassword(u.Password)
	}
}