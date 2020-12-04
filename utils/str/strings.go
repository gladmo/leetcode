package str

import (
	"sort"
)

type myByte []byte

func (m myByte) Len() int {
	return len(m)
}

func (m myByte) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m myByte) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func Sort(in string) string {
	my := myByte(in)
	sort.Sort(my)
	return in
}
