/*
* @Author: starine
* @Date:   2022/7/24 10:31
 */

package main

func repeatedCharacter(s string) byte {
	oc := [26]int{}
	for i := 0; i < len(s)-1; i++ {
		if oc[s[i]-'a'] == 1 {
			return s[i]
		}
		oc[s[i]-'a'] = 1
	}
	return 'a'
}
