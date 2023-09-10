package main

import (
	"fmt"
)

func fuckBrain(source string) {
	arr := make([]byte, 300000)

	ptr := 0

	s := []int{}
	index := 0
	for index < len(source) {
		i := source[index]
		if i > 32 {
			switch string(i) {
			case ".":
				fmt.Print(string(arr[ptr]))
			case ",":
				var c string
				fmt.Scan(&c)
				arr[ptr] = c[0]
			case ">":
				if ptr < 30000 {
					ptr++
				} else {
					fmt.Println("Error: pointer incrementing error: bounds out of range")
					return
				}
			case "<":
				if ptr > 0 {
					ptr--
				} else {
					fmt.Println("Error: pointer incrementing error: bounds out of range")
					return
				}
			case "+":
				arr[ptr]++
			case "-":
				arr[ptr]--
			case "[":
				s = append(s, index-1)
			case "]":
				if arr[ptr] != 0 {
					index = s[len(s)-1]
				}
				s = s[:len(s)-1]
			}
		}
		index++
	}
}
