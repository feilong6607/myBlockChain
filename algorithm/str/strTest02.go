package str

import "strings"

func STRMETHOD2(strArr []string) string {

	if len(strArr) == 0 {

		return ""
	}

	str := strArr[0]
	if len(str) == 0 {
		return ""
	}

	for i := 1; i < len(strArr); i++ {

		for !strings.Contains(strArr[i], str) {

			str = str[0 : len(str)-1]
			if len(str) == 0 {
				return ""
			}

		}

	}
	return str

}
