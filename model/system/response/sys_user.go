package response

import (
	"project/model/system"
)

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}

type LoginResponse struct {
	User      system.SysUser `json:"user"`
	Token     string         `json:"token"`
	ExpiresAt int64          `json:"expiresAt"`
}

type DeptUser struct {
	Id        uint   `json:"user_id"`
	HeaderImg string `json:"url"`
	NickName  string `json:"nick_name"`
}
