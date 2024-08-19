package errors_test

import (
	"fmt"
	"testing"

	"github.com/tshiba06/account_backend/internal/errors"
)

func TestError(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		preparedFunc func() (int, string)
		assertFunc   func(string)
	}{
		{
			name: "normal",
			preparedFunc: func() (int, string) {
				return 400, "parameters are invalid."
			},
			assertFunc: func(m string) {
				want := fmt.Sprintf("Error code: %d, Message: %s", 400, "parameters are invalid.")
				if m != want {
					t.Errorf("expected is %s, but actual is %s", want, m)
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			code, m := tc.preparedFunc()

			err := errors.New(code, m)

			tc.assertFunc(err.Error())
		})
	}
}
