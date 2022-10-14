package dtos

// UserPasswordDto 修改密码dto
type UserPasswordDto struct {
	OldPassword string // 旧密码
	Password    string // 新密码
}
