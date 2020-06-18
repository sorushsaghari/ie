package user

type Dto struct {
	Email string `form:"email"`
	Password string `form:"password"`
}
type CreateDto struct {
	Dto
	Name string `form:"name"`
	Family string `form:"family"`
	Username string `form:"username"`
	Avatar string `form:"avatar"`
}


func (dto *Dto) Parse() *User{
	return & User{
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func (dto *CreateDto) Parse() *User{
	return & User{
		Name:     dto.Name,
		Family:   dto.Family,
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		Avatar:   dto.Avatar,
	}
}
