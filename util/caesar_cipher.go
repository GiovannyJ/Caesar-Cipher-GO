package caesarciphergo
import ("strings")

func Enc(text string, shift int) string {
	shift %= 26 

	text = strings.ToUpper(text)
	
	result := ""
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			encrypted := ((int(char) - int('A') + shift) % 26) + int('A')
			result += string(rune(encrypted))
		} else {
			result += string(char)
		}
	}
	
	return result
}

func Dec(text string, shift int) string {
	shift %= 26 
	
	text = strings.ToUpper(text)
	
	result := ""
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			decrypted := ((int(char) - int('A') - shift + 26) % 26) + int('A')
			result += string(rune(decrypted))
		} else {
			result += string(char)
		}
	}
	
	return result
}