package unicode

import (
	"fmt"
)

const unicode = ("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")

func SkrivUT() {
	for i := 0; i < len(unicode); i++ {
		fmt.Printf("%x\n", unicode[i])
	}

}
