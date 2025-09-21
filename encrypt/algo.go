package encrypt

func Nimbus(str string) string {
	encryptedString := ""
	for _, c := range str {
		char := string(int(c) + 2)
		encryptedString += char
	}

	return encryptedString
}
