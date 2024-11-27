package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tests int
	fmt.Fscanf(reader, "%d\n", &tests)

	for i := 0; i < tests; i++ {
		var lejak int
		fmt.Fscanf(reader, "%d\n", &lejak)

		line, _ := reader.ReadString('\n')
		strNums := strings.Fields(line)
		mas := make([]int, lejak)

		for j, str := range strNums {
			mas[j], _ = strconv.Atoi(str)
		}

		slices.Sort(mas)

		min := 1 << 30
		for j := 0; j < lejak-1; j++ {
			xor := mas[j] ^ mas[j+1]
			if xor < min {
				min = xor
			}
		}

		fmt.Fprintln(writer, min)
	}
}
