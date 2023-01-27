package userusecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/id"
	"github.com/revenue-hack/cleanarchitecture-sample/src/domain/user"
	"github.com/revenue-hack/cleanarchitecture-sample/src/mock/mock_user"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase"
	"github.com/revenue-hack/cleanarchitecture-sample/src/usecase/userusecase/input"
)

var (
	in  = input.GetUserByIDInput{ID: "userid1"}
	ctx = context.TODO()
)

func Test_GetUserByIDExec(t *testing.T) {
	ctrl := gomock.NewController(t)

	inputUserID, err := id.NewUserIDByVal(in.ID)
	if err != nil {
		t.Error(err)
		return
	}
	user1, _ := user.NewUser(inputUserID, "firstName", "lastName")

	userRepo := mock_user.NewMockUserRepository(ctrl)
	userRepo.EXPECT().FindByID(ctx, inputUserID).Return(user1, nil)
	usecase := userusecase.NewGetUserByID(userRepo)
	out, err := usecase.Exec(ctx, &in)
	if err != nil {
		t.Error(err)
	}
	if out.ID != user1.ID().String() {
		t.Errorf("id is not match. result is %s, expected is %s", out.ID, user1.ID())
	}
	if out.FirstName != user1.FirstName() {
		t.Errorf("first name is not match. result is %s, expected is %s", out.FirstName, user1.FirstName())
	}
	if out.LastName != user1.LastName() {
		t.Errorf("last name is not match. result is %s, expected is %s", out.LastName, user1.LastName())
	}
}
