package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func romanToArabic(romanNum string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	arabicNum := 0
	prev := 0
	for _, curr := range romanNum {
		if val, ok := romanMap[curr]; ok {
			if val > prev {
				arabicNum += val - 2*prev
			} else {
				arabicNum += val
			}
			prev = val
		} else {
			fmt.Printf("В римском числе неверный символ: %c\n", curr)
			return 0
		}
	}
	return arabicNum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите римское число: ")
	romanNum, _ := reader.ReadString('\n')
	romanNum = strings.TrimSpace(romanNum) 
	arabicNum := romanToArabic(romanNum)
	fmt.Printf("%s = %d\n", romanNum, arabicNum)
}