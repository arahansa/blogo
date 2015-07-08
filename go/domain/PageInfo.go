// PageInfo
package domain

import "fmt"

type PageInfo struct {
	BeginPage      int
	EndPage        int
	TotalPageCount int
}

func (b PageInfo) PrevBeginPage() int {
	return b.BeginPage - int
}
func (b PageInfo) NextEndPage() int {
	return b.EndPage + int
}

func (b PageInfo) Pagenation() []int {
	var pageRow = make([]int, b.EndPage-b.BeginPage+1)
	for i := 0; b.BeginPage+i <= b.EndPage; i++ {
		pageRow[i] = b.BeginPage + i
	}
	return pageRow
}

func (b PageInfo) String() string {
	return fmt.Sprintf("PageInfo(%d,%d, %d)", b.BeginPage, b.EndPage, b.TotalPageCount)
}
