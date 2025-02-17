package service

type Service struct {
	UserService *UserService
	AuthService *AuthService
}

func New() *Service {
	return &Service{UserService: newUserService()}
}
