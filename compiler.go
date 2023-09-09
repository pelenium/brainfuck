package main

import (
	"fmt"
)

func fuckBrain(source string) {
	arr := make([]byte, 30000)

	ptr := 0

	s := []int{}
	index := 0
	for index < len(source) {
		i := source[index]
		switch string(i) {
		case ".":
			print(string(arr[ptr]))
		case ",":
			var c string
			fmt.Scan(&c)
			arr[ptr] = c[0]
		case ">":
			if ptr < 30000 {
				ptr++
			} else {
				println("Error: bounds out of range")
				return
			}
		case "<":
			if ptr > 0 {
				ptr--
			} else {
				println("Error: bounds out of range")
				return
			}
		case "+":
			arr[ptr]++
		case "-":
			arr[ptr]--
		case "[":
			s = append(s, index)
		case "]":
			if arr[ptr] != 0 {
				index = s[len(s)-1] - 1
				s = s[:len(s)-1]
			}
		}
		index++
	}
}
