package utils

import "regexp"

func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	suc := reg.MatchString(email)
	return suc
}

func CheckPhoneAndRegionValid(region, phone string) bool {
	return region == "+86" && len(phone) == 11 && VerifyPhoneFormat(phone)
}

func VerifyPhoneFormat(phone string) bool {
	regular := "^((13[0-9])|(14[5,7,9])|(15([0-3]|[5-9]))|(166)|(17[0,1,3,5,6,7,8])|(18[0-9])|(19[8|9]))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}
