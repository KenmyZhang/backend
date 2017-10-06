package models
import ()
type StringMap map[string]string
type SessionValue struct {
	Id             string        `bson:"_id" json:"id"`
	Token          string        `bson:"token" json:"token"`
	CreateAt       int64         `bson:"createAt" json:"create_at"`
	ExpiresAt      int64         `bson:"expiresAt" json:"expires_at"`
	LastActivityAt int64         `bson:"lastActivityAt" json:"last_activity_at"`
	UserId         string        `bson:"userId" json:"user_id"`
	DeviceId       string        `bson:"deviceId" json:"device_id"`
	Roles          string        `bson:"roles" json:"roles"`
	IsOAuth        bool          `bson:"isOAuth" json:"is_oauth"`
	Props          StringMap     `bson:"props" json:"props"`
}