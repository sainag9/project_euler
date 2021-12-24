package project_euler

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func NamesScores() int64 {

	data, err := ioutil.ReadFile("./p022_names.txt")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	var sb strings.Builder
	_, _ = sb.Write(data)
	s := strings.Split(sb.String(), ",")
	sort.Strings(s)
	scores := int64(0)
	for i := 0; i < len(s); i++ {
		sum := int64(0)
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == '"' {
				continue
			}
			sum += int64(s[i][j] - 'A' + 1)
		}
		scores += int64(i+1) * sum
	}
	return scores
}
