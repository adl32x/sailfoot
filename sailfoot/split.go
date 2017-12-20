package sailfoot

import "strings"

func SplitLine(line string) []string {

	if len(line) == 0 {
		return nil
	}

	var result []string

	i := 0
	n := 0

	quoteStarted := false

	for i < len(line) {
		var nextChar string
		var currentChar string

		currentChar = string([]rune(line)[i])
		if (i + 1) != len(line) {
			nextChar = string([]rune(line)[i+1])
		}
		// fmt.Println(currentChar)

		if quoteStarted && currentChar == "\\" && nextChar == "'" {
			i = i + 2
			continue
		}

		if !quoteStarted && currentChar == "'" {
			quoteStarted = true
			i++
			n = i
			continue
		} else if quoteStarted && currentChar == "'" {
			quoteStarted = false

			appendS := line[n:i]
			appendS = strings.Replace(appendS, "\\'", "'", -1)

			result = append(result, appendS)

			if nextChar == "" {
				break
			}

			if nextChar == " " {
				i++
				i++
				n = i
				continue
			}
		} else if !quoteStarted && currentChar == " " {
			result = append(result, line[n:i])
			i++
			n = i
			continue
		} else if nextChar == "" {
			result = append(result, line[n:i+1])
			break
		} else {
			i++
			continue
		}

	}
	return result
}
