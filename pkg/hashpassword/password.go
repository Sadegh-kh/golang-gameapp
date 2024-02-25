package hashpassword

import "encoding/base64"

// TODO - better to use bcrypt library
func EncodePasword(p string) string {
	p = base64.StdEncoding.EncodeToString([]byte(p))
	return p
}
func DecodePasword(p string) (string, error) {
	var data = make([]byte, 1024)
	data, err := base64.StdEncoding.DecodeString(p)
	if err != nil {
		return "", err
	}
	p = string(data)
	return p, nil
}
