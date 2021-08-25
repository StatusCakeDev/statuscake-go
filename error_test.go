package statuscake_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/StatusCakeDev/statuscake-go"
)

func TestAPIError(t *testing.T) {
	t.Run("it conforms to the error interface", func(t *testing.T) {
		err := statuscake.NewAPIError(
			"error message",
			errors.New("parent error message"),
		)

		got := err.Error()
		expectEqual(t, got, "parent error message: error message")
	})
}

func TestUnwrap(t *testing.T) {
	t.Run("it returns the wrapped error if exists", func(t *testing.T) {
		parent := errors.New("parent error message")

		err := statuscake.NewAPIError(
			"error message",
			parent,
		)

		got := err.Unwrap()
		if got != parent {
			t.Errorf("expected %+v, got %+v", parent, got)
		}
	})

	t.Run("it returns nil if no wrapped error exists", func(t *testing.T) {
		err := statuscake.NewAPIError(
			"error message",
			nil,
		)

		got := err.Unwrap()
		if got != nil {
			t.Errorf("expected <nil>, got %+v", got)
		}
	})
}

func TestErrors(t *testing.T) {
	t.Run("returns error messages contained within the error", func(t *testing.T) {
		errors := map[string][]string{
			"field": []string{
				"is required",
				"should be numeric",
			},
		}

		err := statuscake.APIError{
			Errors: errors,
		}

		got := statuscake.Errors(err)
		expectEqual(t, got, errors)
	})

	t.Run("returns an empty map if the error is of an unexpected type", func(t *testing.T) {
		got := statuscake.Errors(errors.New("unexpected error"))
		expectEqual(t, got, map[string][]string{})
	})
}

func expectEqual(t *testing.T, got, expected interface{}) {
	if diff := cmp.Diff(got, expected, cmp.AllowUnexported(statuscake.APIError{})); diff != "" {
		fmt.Println(diff)
		t.Fail()
	}
}
