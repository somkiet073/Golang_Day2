package validate

import "fmt"

func CardNumber(number string) string {
	var result string
	switch number {
	case "1":
		result = "11111111"
	case "2":
		result = "22222222"
	case "3":
		result = "33333333"
	case "4":
		result = "44444444"
	case "5":
		result = "55555555"
	default:
		fmt.Println("No information available for that day.")
	}

	return result
}
