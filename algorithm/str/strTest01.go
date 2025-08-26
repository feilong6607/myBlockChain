package str

import (
	"strings"
)

func STRMETHOD(str string) bool {
	for strings.Contains(str, "{}") || strings.Contains(str, "[]") || strings.Contains(str, "()") {

		if strings.Contains(str, "{}") {
			str = strings.Replace(str, "{}", "", -1)
		} else if strings.Contains(str, "[]") {
			str = strings.Replace(str, "[]", "", -1)
		} else {
			str = strings.Replace(str, "()", "", -1)
		}

	}
	return str == ""

}
