// https://leetcode.com/problems/defanging-an-ip-address/
package defangipaddr

import "strings"

// defangIPaddr defangs a valid IPv4 IP address, replacing "." with "[.]"
func defangIPaddr(address string) string {
	var strBuilder strings.Builder
	var ch string // Holds the character to add to the builder

	for _, c := range address {
		if c == '.' {
			ch = "[.]"
		} else {
			ch = string(c)
		}
		strBuilder.WriteString(ch)
	}
	defanged := strBuilder.String()
	return defanged
}
