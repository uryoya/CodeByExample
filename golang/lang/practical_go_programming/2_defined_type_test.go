package practical_go_programming

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 2.2.1 メソッドを追加して組み込み型を拡張する
type HTTPStatus int

const (
	StatusOK              HTTPStatus = 200
	StatusUnauthorized    HTTPStatus = 401
	StatusPaymentRequired HTTPStatus = 402
	StatusForbidden       HTTPStatus = 403
)

func (s HTTPStatus) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	default:
		return fmt.Sprintf("HTTPStatus(%d)", s)
	}
}

func Test2_2_1_メソッドを追加して組み込み型を拡張する(t *testing.T) {
	assert.Equal(t, StatusOK.String(), "OK")
	assert.Equal(t, StatusForbidden.String(), "Forbidden")
	s := HTTPStatus(500)
	assert.Equal(t, s.String(), "HTTPStatus(500)")
}

// 2.5 機密情報を扱うフィールドを定義して出力書式をカスタマイズする
type ConfidentialCustomer struct {
	CustomerID int64
	CreditCard CreditCard
}

type CreditCard string

func (c CreditCard) String() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func (c CreditCard) GoString() string {
	return "xxxx-xxxx-xxxx-xxxx"
}

func Test2_5_機密情報を扱うフィールドを定義して出力書式をカスタマイズする(t *testing.T) {
	c := ConfidentialCustomer{
		CustomerID: 1,
		CreditCard: "4111-1111-1111-1111",
	}

	assert.Equal(t, fmt.Sprint(c), "{1 xxxx-xxxx-xxxx-xxxx}")
	assert.Equal(t, fmt.Sprintf("%v", c), "{1 xxxx-xxxx-xxxx-xxxx}")
	assert.Equal(t, fmt.Sprintf("%+v", c), "{CustomerID:1 CreditCard:xxxx-xxxx-xxxx-xxxx}")
	assert.Equal(t, fmt.Sprintf("%#v", c), "practical_go_programming.ConfidentialCustomer{CustomerID:1, CreditCard:xxxx-xxxx-xxxx-xxxx}")
	bytes, _ := json.Marshal(c)
	assert.Equal(t, fmt.Sprint(string(bytes)), `{"CustomerID":1,"CreditCard":"4111-1111-1111-1111"}`) // 元通り利用可能
}
