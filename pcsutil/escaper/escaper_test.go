package escaper_test

import (
	"fmt"
	"testing"
	"xpan/pcsutil/escaper"
)

func TestEscape(t *testing.T) {
	fmt.Println(escaper.Escape(`asdf'asdfasd[]a[\[][sdf\[d]`, []rune{'[', '\''}))
}
