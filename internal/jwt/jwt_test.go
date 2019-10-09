package jwt

import (
	"fmt"
	"strings"
)

func ExampleEncode() {
	s, err := Encode(Msg{1, "meng"})
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	ss := strings.Split(s, ".")
	fmt.Println(ss[0])
	r2 := Msg{}
	r2, err = Decode(s)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(r2.UserID)
	fmt.Println(r2.UserName)
	// OutPut:
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
	// 1
	// meng
}
