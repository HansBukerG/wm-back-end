package utils

import (
	"strconv"
	"strings"

	model "github.com/HansBukerG/wm-back-end/src/models"
)

func LookForPalindromes(product *model.Product) bool {
	return (IsPalindrome(strconv.Itoa(product.Id)) ||
		CheckFilter(product.Brand) ||
		CheckFilter(product.Description))
}

func IsPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func CheckFilter(text string) bool {
	descriptionFields := strings.Fields(text)
	for _, value := range descriptionFields {
		if IsPalindrome(value) {
			return true
		}
	}
	return false
}

func CheckValue(search string) int {
	_, err := strconv.Atoi(search)
	search = strings.Trim(search, " ")
	if err == nil { //ITS A NUMBER
		return 1
	} else { // NOT A NUMBER
		if len(search) >= 3 {
			return 2
		} else { // Dont accomplish the requeriments
			return 0
		}
	}
}
