package presenter

import (
	"net/http"

	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/useroutput"
)

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
	Create(out *useroutput.CreateUserOutput)
}

func (p *userPresent) UserList(out *useroutput.UserList) {
	p.delivery.JSON(http.StatusOK, out)
}

func (p *userPresent) UserByID(out *useroutput.UserByID) {
	p.delivery.JSON(http.StatusOK, out)
}
func (p *userPresent) Create(out *useroutput.CreateUserOutput) {
	p.delivery.JSON(http.StatusCreated, out)
}
