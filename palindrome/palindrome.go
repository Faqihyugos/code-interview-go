package palindrome

func palindromeRecursive(value string, i int) bool {
	if i < len(value)/2 {
		if string(value[i]) != string(value[(len(value)-i-1)]) {
			return false
		} else {
			return palindromeRecursive(value, i+1)
		}
	} else {
		return true
	}
}

func IsPalindrome(value string) bool {
	// recursive
	return palindromeRecursive(value, 0)

	// // Membandingkan char dengan cara cek char
	// for i := 0; i < len(value)/2; i++ { //Optimasi dengan panjang karakter dibagi 2
	// 	indexAwal := i
	// 	indexAkhir := (len(value) - i - 1)

	// 	if string(value[indexAwal]) != string(value[indexAkhir]) {
	// 		return false
	// 	}
	// }
	// return true

	// // menggunakan perulangan for untuk membalikan string dan menggunakan string
	// var temp string
	// for i := len(value) - 1; i >= 0; i-- {
	// 	temp += string(value[i])
	// }

	// return temp == value
}
