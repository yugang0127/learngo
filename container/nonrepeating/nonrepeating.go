package main

import "fmt"

func lengthOfNonRepeatingStr(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingStr("abcdadc"),
		lengthOfNonRepeatingStr("aaa"),
		lengthOfNonRepeatingStr("cdefgbkaie"),
		lengthOfNonRepeatingStr("123aaa"),
		lengthOfNonRepeatingStr("我爱鱼丸粗面我爱鱼丸粗面"),
		lengthOfNonRepeatingStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"),
	)
}
