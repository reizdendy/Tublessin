package login

import (
	"context"
	"tublessin/common/model"
)

type LoginServiceApi struct {
	LoginService model.LoginClient
}

type LoginServiceApiInterface interface {
	HandleLoginMontir(montirAccount *model.MontirLoginForm) (*model.LoginResponeMessage, error)
}

func NewLoginServiceApi(loginService model.LoginClient) LoginServiceApiInterface {
	return LoginServiceApi{LoginService: loginService}
}

func (s LoginServiceApi) HandleLoginMontir(montirAccount *model.MontirLoginForm) (*model.LoginResponeMessage, error) {
	result, err := s.LoginService.MontirLogin(context.Background(), montirAccount)
	if err != nil {
		return nil, err
	}

	return result, nil
}
