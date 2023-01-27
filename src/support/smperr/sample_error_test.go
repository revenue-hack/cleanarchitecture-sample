package smperr_test

import (
	"testing"

	"github.com/revenue-hack/cleanarchitecture-sample/src/support/smperr"
)

func TestBadRequest(t *testing.T) {
	err := func() error { return smperr.BadRequestf("sampleerror: %d", 1) }

	switch err().(type) {
	case smperr.AppError:
	default:
		t.Error("failed to cast")
	}
}
