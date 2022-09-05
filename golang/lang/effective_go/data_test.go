/**
 * Effective Go
 * https://go.dev/doc/effective_go#data
 */
package effective_go

import (
	"testing"
	"uryoya/code_by_example/test"
)

func TestMaps(t *testing.T) {
	var m = map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	test.Expect(t, m["3"] == 3)
}
