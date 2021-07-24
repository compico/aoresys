package userutil_test

import (
	"testing"

	"github.com/compico/aoresys/internal/userutil"
)

var usernames = map[string]bool{
	"Compico":   true,
	"c0mpico":   true,
	"comp|co":   false,
	"com&*%co":  false,
	"cooompico": true,
	"c*oM*2043": false,
}

func TestValidateByEqual(t *testing.T) {
	for k, v := range usernames {
		x := userutil.NewValidator(k)
		if r := x.ValidateByEqual(); v != r {
			t.Errorf("%v\t%v != %v \n", k, v, r)
		}
	}
}
