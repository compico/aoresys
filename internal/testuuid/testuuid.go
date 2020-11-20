package testuuid

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func StartTest() {
	x := []string{}
	var (
		zero int
		one  int
	)
	for i := 0; i < 1000; i++ {
		x = append(x, uuid.New().String())
	}
	for i := 0; i < len(x); i++ {
		u := []byte{}
		z := strings.ReplaceAll(x[i], "-", "")

		y, err := hex.DecodeString(string(z))
		if err != nil {
			fmt.Println(err.Error())
		}
		for l := 3; l < 16; l += 4 {
			y[l] &^= 0b11110000
			u = append(u, y[l])
		}
		r := ((u[0] ^ u[1]) ^ (u[2] ^ u[3])) % 2
		if r == 0 {
			zero++
		}
		if r == 1 {
			one++
		}
	}
	fmt.Printf("\tZero == %v\tOne == %v", zero, one)
}
