package checks

import (
	"regexp"
)

func Username(s ...string) (string, bool) {
	for i := range s {
		switch {
		case len(s) == 0:
			return "用户名或者密码 为空值!", false
		case f(s[i]):
			return "用户名或者密码 不合法!", false
		}
	}
	return "", true
}

func f(s string) bool {
	ok, _ := regexp.MatchString("^[a-zA-Z0-9-\\.!@$%#+-=~]{4,32}$", s)
	return ok
}
