package ctfflag

func FetchFlag(password string) string {
	pass := []byte{96, 47, 54, 54, 97, 47, 53, 51, 52, 100, 56, 51, 54, 47, 53, 51, 49, 50, 52, 53, 48, 48, 98, 98, 52, 53, 96, 96, 97, 97, 48, 55}
	flag := []byte{101, 107, 96, 102, 122, 102, 110, 94, 102, 110, 94, 102, 96, 99, 102, 100, 115, 94, 102, 104, 115, 103, 116, 97, 94, 106, 100, 120, 124}

	for i, c := range pass {
		pass[i] = c + 1
	}

	for i, c := range flag {
		flag[i] = c + 1
	}

	if password == string(pass) {
		return string(flag)
	}

	return ""
}
