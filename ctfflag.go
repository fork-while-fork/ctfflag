package ctfflag

import "fmt"

func FetchFlag(password string) {
	pass := []byte{97, 48, 55, 55, 98, 48, 54, 52, 53, 101, 57, 52, 55, 48, 54, 52, 50, 51, 53, 54, 49, 49, 99, 99, 53, 54, 97, 97, 98, 98, 49, 56}
	flag := []byte{102, 108, 97, 103, 123, 103, 111, 95, 103, 111, 95, 103, 97, 100, 103, 101, 116, 95, 103, 105, 116, 104, 117, 98, 95, 107, 101, 121, 125}
	if password == string(pass) {
		fmt.Println(string(flag))
	}
}
