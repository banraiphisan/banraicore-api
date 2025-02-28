package utils

import (
	"encoding/json"
	"fmt"
)

func PrintToJSON(v interface{}) {
	fmt.Println("PrintToJSON")

	s, _ := json.MarshalIndent(v, "", "\t")
	fmt.Println(string(s))
}
func IsEqueThen(input, target, success, failure string) string {
	if input == target {
		return success
	}

	return failure
}
