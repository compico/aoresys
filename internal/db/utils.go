package db

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/compico/aoresys/internal/userutil"
	"github.com/google/uuid"
)

func generateUUID(model bool) (string, error) {
	var buuid string
	for {
		uuuid, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		buuid = uuuid.String()
		suuid := strings.ReplaceAll(uuuid.String(), "-", "")
		fmt.Println(suuid)
		h, err := hex.DecodeString(suuid)
		if err != nil {
			return "", err
		}
		u := []byte{}
		for i := 3; i < 16; i += 4 {
			h[i] &^= 0b11110000
			u = append(u, h[i])
		}
		r := ((u[0] ^ u[1]) ^ (u[2] ^ u[3])) % 2
		if r == 0 && userutil.STEVE {
			break
		}
		if r == 1 && userutil.ALEX {
			break
		}
	}
	return buuid, nil
}
