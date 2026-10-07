package generator

import (
	"crypto/rand"
	"strings"
)

// принимает:
// bool:
// цифры, специальные символы
// возвр готовый пароль
func GenPassword(settings map[string]*bool, leng int) string {
	var (
		lower = "abcdefghijklmnopqrstuvwxyz"
		upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		num   = "0123456789"
		spec  = `!@#$%^&*()_+-=[]{}|;:'",.<>?/~` + "`"
	)
	var charset strings.Builder
	charset.Grow(leng)
	if *settings["digits"] {
		charset.WriteString(num)
	}
	if *settings["special"] {
		charset.WriteString(spec)
	}
	if *settings["lowercase"] {
		charset.WriteString(lower)
	}
	if *settings["uppercase"] {
		charset.WriteString(upper)
	}
	available := charset.String()
	availableLen := len(available)
	password := make([]byte, leng)
	randomBytes := make([]byte, leng)
	rand.Read(randomBytes)
	for i := 0; i < leng; i++ {
		password[i] = available[int(randomBytes[i])%availableLen]
	}
	return string(password)
}
