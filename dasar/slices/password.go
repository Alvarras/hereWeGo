package dasar

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	upperCase := false
	oneDigit := false

	for _, c := range password {
		if c >= 'A' && c <= 'Z' {
			upperCase = true
		}
		if c >= '0' && c <= '9' {
			oneDigit = true

		}
	}

	return upperCase && oneDigit

}
