package main

import (
	"fmt"
	"strconv"
	"strings"
)

var code string

func main() {

	country := map[string]string{"TH": "Thailand", "HK": "Hongkong", "JP": "Japan", "CN": "China", "US": "America"}
	for {
		fmt.Println("Enter country code (TH,HK,JP,CN,US):")
		fmt.Scanln(&code)
		if _, err := strconv.ParseFloat(code, 64); err == nil {
			fmt.Printf("%q looks like a number.\n", code)
		} else {
			break
		}
	}
	code2 := strings.ToUpper(code)
	fmt.Println(code2, "is country code of", country[code2])

}
