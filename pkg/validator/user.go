package validator

import "strconv"

func IsPhoneNumberValid(phoneNumber string) bool {
	// TODO - use regular expertion to support +98
	if len(phoneNumber) != 11 {
		return false
	}
	if phoneNumber[:2] != "09" {
		return false
	}
	// 09ahfjkdshjl
	if _,err:=strconv.Atoi(phoneNumber[2:]); err!=nil {
		return false
	}

	return true
}