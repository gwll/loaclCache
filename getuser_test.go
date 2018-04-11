// getuser_test
package loaclCache

import (
	"fmt"
	"testing"
	"time"
)

func Test_getuser(t *testing.T) {

	Set("asdf", 5, 2)

	time.Sleep(time.Second * 1)
	num, ok := Get("asdf")

	fmt.Println(num, ok)
}
