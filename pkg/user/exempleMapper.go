package user

type IUserMapper interface {
	DtoToUser(dto UsersDto) Users
	UserToDto(user Users) UsersDto
}

type mapper struct{}

func UserMapper() IUserMapper {
	return &mapper{}
}

func (m *mapper) UserToDto(user Users) UsersDto {
	//TODO implement me
	panic("implement me")
}

func (m *mapper) DtoToUser(dto UsersDto) Users {
	//TODO implement me
	panic("implement me")
}
