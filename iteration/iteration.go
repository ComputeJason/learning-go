package iteration

import "strings"

func Repeat(toRepeat string, iterationLength int) string {
	
	var finalString strings.Builder;
	

	for i := 0; i < iterationLength; i++ {
		finalString.WriteString(toRepeat)
	}

	return finalString.String()
}

