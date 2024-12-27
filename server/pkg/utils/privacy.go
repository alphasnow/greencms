// @author AlphaSnow

package utils

import "strings"

func PrivacyEmail(str string) string {
	if str == "" {
		return ""
	}
	parts := strings.Split(str, "@")
	name := parts[0]
	maskedName := name[:3] + "****"
	return maskedName + "@" + parts[1]
}

func PrivacyPhone(str string) string {
	if str == "" {
		return ""
	}
	return str[:3] + "****" + str[7:]
}

func PrivacyUsername(str string) string {
	if str == "" {
		return ""
	}
	sl := len(str)
	return str[:2] + "****" + str[sl-2:]
}
