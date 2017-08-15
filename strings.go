/**
 * Created by 93201 on 2017/6/20.
 */
package utils

import "strings"

func HideMobile(mobile string) string {
	return HideString(mobile, 3, 8)
}

func HideEmail(email string) string {
	tmp := strings.Split(email, "@")
	if len(tmp) < 2 || tmp[0] == "" {
		return email
	}
	start, end := 0, len(tmp[0])
	if len(tmp[0]) > 10 {
		start, end = 3, end-4
	} else if len(tmp[0]) > 6 {
		start, end = 2, 6
	} else if len(tmp[0]) > 3 {
		start = 1
	}
	return HideString(tmp[0], start, end) + "@" + tmp[1]
}

func HideIDCard(idCard string) string {
	if len(idCard) == 18 {
		return strings.Repeat("*", 6) + idCard[6:12] + strings.Repeat("*", 6)
	}
	return strings.Repeat("*", len(idCard))
}

func HideString(s string, start, end int) string {
	if end > len(s) {
		end = len(s)
	}
	if start > end || start < 0 {
		return strings.Repeat("*", len(s))
	}
	return s[:start] + strings.Repeat("*", end-start) + s[end:]
}

func ZipString(str string) string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(
				strings.Replace(str, " ", "", -1),
				"\r", "", -1),
			"\n", "", -1),
		"\t", "", -1)
}
