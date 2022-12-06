package practical_go_programming

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
)

type wrappedError struct {
	msg string
	err error
}

func (e *wrappedError) Error() string {
	return fmt.Sprintf("msg: %s, err: %s", e.msg, e.err.Error())
}

func (e *wrappedError) Unwrap() error {
	return e.err
}

func Test_5_1_5_エラーの色々な比較(t *testing.T) {
	err := errors.New("original error")
	werr := &wrappedError{"wrapped error", err}

	// エラーを値として比較する
	// wrappedErrorはUnwrap()を実装しているので、errors.Is()で再帰的にUnwrapして評価される
	assert.True(t, errors.Is(werr, err))

	var pathError *fs.PathError
	werr2 := &wrappedError{"wrapped path error", &fs.PathError{}}
	// エラーを型として比較する
	assert.True(t, errors.As(werr2, &pathError))
}
