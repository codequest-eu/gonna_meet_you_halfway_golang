package util

import "strings"

func RemoveCharacter(character rune, str string) string {
	return strings.Map(func(r rune) rune {
		if r == character {
			return -1
		}
		return r
	}, str)
}
