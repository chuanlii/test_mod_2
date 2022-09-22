package testmod2

import (
	"fmt"

	base "github.com/chuanlii/test_mod_base"
)

func init() {
	fmt.Println("testmod2")
}
func F2() string {
	base.Base()
	return "f2"
}
