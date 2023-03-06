package presenter

import "github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/useroutput"

type userPresent struct {
	delivery Presenter
}

func NewUserPresenter(p Presenter) UserPresenter {
	return &userPresent{
		delivery: p,
	}
}

type UserPresenter interface {
	UserList(out *useroutput.UserList)
	UserByID(out *useroutput.UserByID)
}

func (p *userPresent) UserList(out *useroutput.UserList) {
	p.delivery.JSON(200, out)
}

func (p *userPresent) UserByID(out *useroutput.UserByID) {
	p.delivery.JSON(200, out)
}
