package main

import (
	"regexp"
	"strconv"
	"strings"
)

func solve(IP string) string {
	if validIPv4(IP) {
		return "IPv4"
	} else if validIPv6(IP) {
		return "IPv6"
	} else {
		return "Neither"
	}
}
func validIPv4(IP string) bool {
	strs := strings.Split(IP, ".")
	if len(strs) != 4 {
		return false
	}
	for _, str := range strs {
		if len(str) > 3 {
			return false
		}
		if len(str) > 1 && str[0] == '0' {
			return false
		}
		val, err := strconv.Atoi(str)
		if err != nil || val < 0 || val > 255 {
			return false
		}
	}
	return true
}
func validIPv6(IP string) bool {
	strs := strings.Split(IP, ":")
	if len(strs) != 8 {
		return false
	}
	for _, str := range strs {
		if len(str) > 4 || len(str) == 0 {
			return false
		}
		re := regexp.MustCompile(`^([0-9]|[a-f]|[A-F])+$`)
		if !re.MatchString(str) {
			return false
		}
	}
	return true
}

func main() {

}
