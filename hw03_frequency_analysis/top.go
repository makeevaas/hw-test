package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var NotWord = "-"

func Top10(s string) []string {
	// Place your code here.
	s = strings.ToLower(s)
	var lstWords []string
	wordLsts := make([][]string, 0, len(s))
	// top 10
	newSlice := strings.Fields(s)
	sort.Strings(sort.StringSlice(newSlice))
	var word string
	re := regexp.MustCompile(`^[!\,\.\d_]|[!\,\.\d_]$|^` + regexp.QuoteMeta(NotWord) + `$`)
	for _, r := range newSlice {
		if r == NotWord {
			continue
		}
		r = re.ReplaceAllString(r, "")
		if r != word {
			word = r
		} else {
			continue
		}
		lstWords = nil
		for _, r := range newSlice {
			r = re.ReplaceAllString(r, "")
			if r == word {
				lstWords = append(lstWords, r)
			}
		}
		wordLsts = append(wordLsts, lstWords)
	}
	// определить большинство
	sort.SliceStable(wordLsts, func(i, j int) bool {
		return len(wordLsts[i]) > len(wordLsts[j])
	})
	// вывести топ 10
	top10Str := make([]string, 0, 9)
	top := 10
	for i, r := range wordLsts {
		if i == top {
			break
		}
		top10Str = append(top10Str, r[0])
	}
	return top10Str
}
