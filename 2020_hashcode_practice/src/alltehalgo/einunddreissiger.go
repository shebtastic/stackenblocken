package alltehalgo

import (
	"fmt"
)

func HelpSlices(x []int) {
	for i := 0; i > len(x); i++ {
		fmt.Printf("%+v", x[i])
	}
}