package prompt

import (
	"fmt"
	"testing"
)

func TestPrompt_Render(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("%#v\n", a[:0])
	fmt.Printf("%#v\n", a[:2])
	fmt.Printf("%#v\n", a[2:])
	fmt.Printf("%#v\n", a[:5])
	fmt.Printf("%#v\n", a[5:])
}
