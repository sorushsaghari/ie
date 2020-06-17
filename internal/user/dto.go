package user

type Dto struct {
	Name string `form:"name"`
	Family string `form:"family"`
	Username string `form:"username"`
	Email string `form:"email"`
	Password string `form:"password"`
	Avatar string `form:"avatar"`
}


func (dto *Dto) Parse() *User{
	return & User{
		Name:     dto.Name,
		Family:   dto.Family,
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		Avatar:   dto.Avatar,
	}
}
