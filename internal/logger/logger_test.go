package logger_test

import (
	"log/slog"
	"testing"

	"github.com/tshiba06/account_backend/internal/logger"
)

func TestReplaceAttr(t *testing.T) {
	type inputType struct {
		groups []string
		a      slog.Attr
	}

	testCases := []struct {
		name       string
		input      inputType
		assertFunc func(a slog.Attr)
	}{
		{
			name: "secret sensitive fields",
			input: inputType{
				groups: []string{"attr"},
				a:      slog.Any("email", "test@example.com"),
			},
			assertFunc: func(a slog.Attr) {
				if !a.Equal(slog.Any("email", "**********")) {
					t.Errorf("email field must be secret")
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := logger.ReplaceAttr(tc.input.groups, tc.input.a)

			tc.assertFunc(a)
		})
	}
}

func TestGetDiff(t *testing.T) {
	type inputType struct {
		before any
		after  any
	}

	testCases := []struct {
		name         string
		preparedFunc func() inputType
		assertFunc   func(before, after map[string]any)
	}{
		{
			name: "case: struct",
			preparedFunc: func() inputType {
				type testStruct struct {
					Name string
					Age  int
				}
				return inputType{
					before: testStruct{
						Name: "test",
						Age:  20,
					},
					after: testStruct{
						Name: "test",
						Age:  30,
					},
				}
			},
			assertFunc: func(before, after map[string]any) {
				if before["Age"] != 20 || after["Age"] != 30 {
					t.Errorf("cannot get the diff of Age")
				}
				// check the other fields
				if len(before) != 1 || len(after) != 1 {
					t.Errorf("cannot get the diff: before %+v\tafter %+v", before, after)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := tc.preparedFunc()
			before, after := logger.GetDiff(input.before, input.after)

			tc.assertFunc(before, after)
		})
	}
}
