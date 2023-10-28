package main

import (
	"fmt"
)

func fuckBrain(source string) {
	arr := make([]byte, 300000)

	ptr := 15000

	stack := []int{}
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
					fmt.Println("Error: pointer decrementing error: bounds out of range")
					return
				}
			case "+":
				arr[ptr]++
			case "-":
				arr[ptr]--
			case "[":
				if arr[ptr] == 0 {
					loopCounter := 1
					for loopCounter > 0 {
						index++
						if string(source[index]) == "[" {
							loopCounter++
						} else if string(source[index]) == "]" {
							loopCounter--
						}
					}
				} else {
					stack = append(stack, index)
				}
			case "]":
				if arr[ptr] != 0 {
					index = stack[len(stack)-1]
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
		index++
	}
}
