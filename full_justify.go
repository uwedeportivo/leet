package leet

import (
	"fmt"
	"strings"
)

func justifyLine(words []string, maxWidth int) (string, int, error) {
	var index int
	ll := 0
	pieces := make([]string, 0, 2 * len(words) - 1)
	first := true

	for index = 0; index < len(words); index++ {
		if first {
			first = false
			if ll + len(words[0]) <= maxWidth {
				ll = ll + len(words[0])
				pieces = append(pieces, words[0])
			} else {
				return "", 0, fmt.Errorf("maxWidth smaller than word")
			}
		} else {
			if ll+len(words[index])+1 <= maxWidth {
				ll = ll + len(words[index]) + 1
				pieces = append(pieces, " ")
				pieces = append(pieces, words[index])
			} else {
				break
			}
		}
	}

	if index == len(words) {
		// last line is left justified
		padding := maxWidth - ll
		if padding > 0 {
			pieces = append(pieces, strings.Repeat(" ", padding))
		}
	} else  if ll < maxWidth {
		padding := maxWidth - ll

		if len(pieces) == 1 {
			pieces[0] = pieces[0] + strings.Repeat(" ", padding)
		} else {
			numSepPieces := (len(pieces) - 1) / 2
			padSize := padding / numSepPieces
			padStr := strings.Repeat(" ", padSize)
			padRemains := padding % numSepPieces

			for i := 1; i < len(pieces); i += 2 {
				pieces[i] = pieces[i] + padStr
			}

			for i, j := 1, 0; j < padRemains; i, j = i + 2, j + 1 {
				pieces[i] = pieces[i] + " "
			}
		}
	}

	return strings.Join(pieces, ""), index, nil
}

func FullJustify(words []string, maxWidth int) []string {
	var res []string

	rest := words
	for len(rest) > 0 {
		justifiedLine, idx, err := justifyLine(rest, maxWidth)
		if err != nil {
			return nil
		}
		res = append(res, justifiedLine)
		rest = rest[idx:]
	}
	return res
}
