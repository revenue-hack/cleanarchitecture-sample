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
	UserList(out *output.UserList)
	UserByID(out *output.UserByID)
}

func (p *userPresent) UserList(out *output.UserList) {
	p.delivery.JSON(200, out)
}

func (p *userPresent) UserByID(out *output.UserByID) {
	p.delivery.JSON(200, out)
}
