package design

//  IUserinfo 用户接口
type IUserInfo interface {
	Login(phone, code int) (*UserInfo, error)
	Register(phone, code int) (*UserInfo, error)
}

type UserInfo struct {
	Name string
}

// IUserInfoFacade门面模式
type IUserInfoFacade interface {
	LoginOrRegister(phone, code int) (*UserInfo, error)
}

type UserService struct{}

func (u UserService) Login(phone, code int) (*UserInfo, error) {
	return &UserInfo{Name: "login"}, nil
}

func (u UserService) Register(phone, code int) (*UserInfo, error) {
	return &UserInfo{Name: "register"}, nil
}

func (u UserService) LoginOrRegister(phone, code int) (*UserInfo, error) {
	user, err := u.Login(phone, code)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	return u.Register(phone, code)
}
