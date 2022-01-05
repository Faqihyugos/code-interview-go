package vokal

import "strconv"

func CountVokal(str string) string {
	count := 0
	for _, ch := range str {
		switch ch {
		case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
			count++
		}
	}
	countstr := strconv.Itoa(count)
	return "kalimat " + str + "mengandung vokal sebanyak: " + countstr
}

// cek huruf vokal
// karakter vokal => a, i, u, e, o
// func VokalSwitch(value rune) string {
// 	switch value {
// 	case 'a', 'i', 'u', 'e', 'o', 'A', 'I', 'U', 'E', 'O':
// 		valuecon1 := strconv.QuoteRune(value)
// 		return valuecon1 +" adalah vokal"
// 	default:
// 		valuecon := strconv.QuoteRune(value)
// 		return valuecon + " adalah konsonan"
// 	}
// }


// func IsVokal(value rune) string{
// 	if value == 'a' || value == 'i' || value == 'u' || value == 'e' || value == 'o' {
// 		valuecon1 := strconv.QuoteRune(value)
// 		return valuecon1 +" adalah vokal"
// 	}else{
// 		valuecon := strconv.QuoteRune(value)
// 		return valuecon + " adalah konsonan"
// 	}
// }