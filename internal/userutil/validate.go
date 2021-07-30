package userutil

type Validator struct {
	text string
}

var (
	zero   byte = 48
	nine   byte = 57
	bigA   byte = 65
	bigZ   byte = 90
	smallA byte = 97
	smallZ byte = 122
)

func NewValidator(text string) *Validator {
	x := new(Validator)
	x.text = text
	return x
}

func (v *Validator) ValidateLen() bool {
	if len(v.text) < 4 || len(v.text) > 32 {
		return false
	}
	return true
}

func (v *Validator) ValidateByEqual() bool {
	for i := 0; i < len(v.text); i++ {
		if !validateByte(v.text[i]) {
			return false
		}
	}
	return true
}
func (v *Validator) GetRightVar() string {
	if v.ValidateByEqual() && v.ValidateLen() {
		return v.text
	}
	return ""
}

func validateByte(b byte) bool {
	if b >= zero && b <= nine {
		return true
	}
	if b >= bigA && b <= bigZ {
		return true
	}
	if b >= smallA && b <= smallZ {
		return true
	}
	return false
}
