package decrypt

func NimbusDecrypt(str string) string {
	decryptString := ""
	for _, c := range str {
		char := string(int(c) - 2)
		decryptString += char
	}

	return decryptString
}
