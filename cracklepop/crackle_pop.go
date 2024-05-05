package cracklepop

import (
	"fmt"
	"strings"
)

func CracklePop(start, end int) (string, error) {
	if start < 0 {
		return "", ErrNegativeInput(start)
	}
	if end < 0 {
		return "", ErrNegativeInput(end)
	}
	if start > end {
		return "", ErrInvalidInput{start, end}
	}

	var builder strings.Builder
	for i := start; i <= end; i++ {
		if i%15 == 0 {
			// i is divisible by both 3 and 5
			builder.WriteString("CracklePop")
		} else if i%3 == 0 {
			builder.WriteString("Crackle")
		} else if i%5 == 0 {
			builder.WriteString("Pop")
		} else {
			builder.WriteString(fmt.Sprint(i))
		}
		builder.WriteString("\n")
	}

	return builder.String(), nil
}
