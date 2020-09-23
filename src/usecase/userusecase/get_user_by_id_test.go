package userusecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/revenue-hack/cleanarchitecture-sample/src/entity"
	"github.com/revenue-hack/cleanarchitecture-sample/src/mock/mock_userusecase"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
)

var (
	in       = input.GetUserByIDInput{ID: "userid1"}
	user1, _ = entity.NewUser(in.ID, "firstName", "lastName")
)

func Test_GetUserByIDExec(t *testing.T) {
	ctrl := gomock.NewController(t)
	userRepo := mock_userusecase.NewMockUserRepository(ctrl)
	userRepo.EXPECT().FindByID(in.ID).Return(user1, nil)

	usecase := userusecase.NewGetUserByID(userRepo)
	out, err := usecase.Exec(&in)
	if err != nil {
		t.Error(err)
	}
	if out.ID != user1.ID() {
		t.Errorf("id is not match. result is %s, expected is %s", out.ID, user1.ID())
	}
	if out.FirstName != user1.FirstName() {
		t.Errorf("first name is not match. result is %s, expected is %s", out.FirstName, user1.FirstName())
	}
	if out.LastName != user1.LastName() {
		t.Errorf("last name is not match. result is %s, expected is %s", out.LastName, user1.LastName())
	}
}
