package main

import (
	"fmt"
	"sort"
)

type FileNames []string

func (p FileNames) Len() int {
	return len(p)
}

func (p FileNames) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p FileNames) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	fn := FileNames{"Samanta", "Nick", "Ruslan", "Aragorn"}
	sort.Sort(fn)
	fmt.Println("fn", fn)
	names := []string{"Samanta", "Nick", "Ruslan", "Aragorn"}
	sort.Sort(sort.StringSlice(names))
	fmt.Println("names", names)
	namesTwo := []string{"Samanta", "Nick", "Ruslan", "Aragorn"}
	sort.Strings(namesTwo)
	fmt.Println("namesTwo", namesTwo)
}
