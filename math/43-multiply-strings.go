package main

import (
	"strconv"
	"strings"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	var add func(s1, s2 string) string
	add = func(s1, s2 string) string {
		sb := strings.Builder{}
		carry := 0
		i, j := len(s1)-1, len(s2)-1
		for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
			n1 := int(s1[i] - '0')
			n2 := int(s2[j] - '0')
			m := n1 + n2 + carry
			sb.WriteString(strconv.Itoa((m % 10)))
			carry = m / 10
		}
		if j >= 0 {
			i = j
			s1 = s2
		}
		for i >= 0 {
			n := int(s1[i]-'0') + carry
			if n > 9 {
				carry = 1
				sb.WriteString(strconv.Itoa(n % 10))
			} else {
				sb.WriteString(strconv.Itoa(n))
				carry = 0
			}
			i--
		}
		if carry > 0 {
			sb.WriteString(strconv.Itoa(carry))
		}
		runes := []rune(sb.String())
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	res := "0"
	for i := len(num2) - 1; i >= 0; i-- {
		n2 := int(num2[i] - '0')
		carry := 0
		s := strings.Repeat("0", len(num2)-i-1)
		for j := len(num1) - 1; j >= 0; j-- {
			n1 := int(num1[j] - '0')
			m := n1*n2 + carry
			s += strconv.Itoa(m % 10)
			carry = m / 10
		}
		if carry > 0 {
			s += strconv.Itoa(carry)
		}
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		res = add(res, string(runes))
	}

	return res
}
