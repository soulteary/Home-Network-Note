package fn

import (
	"os"
	"regexp"
	"strconv"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type NaturalSort []os.DirEntry

func (n NaturalSort) Len() int {
	return len(n)
}

func (n NaturalSort) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NaturalSort) Less(i, j int) bool {
	re := regexp.MustCompile(`(\d+)`)
	si := re.FindAllStringSubmatch(n[i].Name(), -1)
	sj := re.FindAllStringSubmatch(n[j].Name(), -1)

	minLen := len(si)
	if len(sj) < minLen {
		minLen = len(sj)
	}

	for k := 0; k < minLen; k++ {
		ii := atoi(si[k][1])
		jj := atoi(sj[k][1])
		if ii != jj {
			return ii < jj
		}
	}

	return n[i].Name() < n[j].Name()
}
