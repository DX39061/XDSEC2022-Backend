package utility

import "regexp"

func VerifyEmailFormat(email string) bool {
	pattern := "^[a-zA-Z\\d.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z\\d](?:[a-zA-Z\\d-]{0,61}[a-zA-Z\\d])?(?:\\.[a-zA-Z\\d](?:[a-zA-Z\\d-]{0,61}[a-zA-Z\\d])?)*$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func VerifyTelephoneFormat(telephone string) bool {
	pattern := `^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(telephone)
}

func VerifyStudentIDFormat(studentID string) bool {
	if studentID == "" {
		// could be empty
		return true
	}
	pattern := `^\d{11}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(studentID)
}

func VerifyUsernameFormat(username string) bool {
	pattern := `^.{2,16}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)
}

func VerifyPasswordFormat(password string) bool {
	pattern := `^.{8,40}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(password)
}
