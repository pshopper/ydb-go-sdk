package errors

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	for _, test := range []struct {
		error error
		text  string
	}{
		{
			error: Error(fmt.Errorf("TestError")),
			text:  "TestError at `errors.TestError(github.com/ydb-platform/ydb-go-sdk/v3/internal/errors/errors_test.go:14)`",
		},
		{
			error: Errorf("TestError%s", "Printf"),
			// nolint:lll
			text: "TestErrorPrintf at `errors.TestError(github.com/ydb-platform/ydb-go-sdk/v3/internal/errors/errors_test.go:18)`",
		},
		{
			error: ErrorfSkip(0, "TestError%s", "Printf"),
			// nolint:lll
			text: "TestErrorPrintf at `errors.TestError(github.com/ydb-platform/ydb-go-sdk/v3/internal/errors/errors_test.go:23)`",
		},
	} {
		t.Run(test.text, func(t *testing.T) {
			if test.error.Error() != test.text {
				t.Fatalf("unexpected text of error: \"%s\", exp: \"%s\"", test.error.Error(), test.text)
			}
		})
	}
}