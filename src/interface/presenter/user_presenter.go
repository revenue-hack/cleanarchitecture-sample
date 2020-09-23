package presenter

import "github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/output"

type userPresent struct {
	delivery Presenter
}

func NewUserPresenter(p Presenter) UserPresenter {
	return &userPresent{
		delivery: p,
	}
}

type UserPresenter interface {
	UserList(out *output.UserList) error
	UserByID(out *output.UserByID) error
}

func (p *userPresent) UserList(out *output.UserList) error {
	return p.delivery.WriteJson(out)
}

func (p *userPresent) UserByID(out *output.UserByID) error {
	return p.delivery.WriteJson(out)
}
