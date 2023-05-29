package user

type ExampleMapper interface {
	UserToDto(user Users)
	UsersToDtos(user []Users)
}
