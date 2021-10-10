package main

type SummaryRanges struct {
	arr [10001]int
}


func Constructor() SummaryRanges {
	return SummaryRanges{}
}


func (this *SummaryRanges) AddNum(val int)  {
	this.arr[val] = 1
}


func (this *SummaryRanges) GetIntervals() [][]int {
	result := make([][]int, 0)
	var start, end int
	for i := 1; i < len(this.arr) ; i++ {
		switch {
		case this.arr[i-1] == 0 && this.arr[i] == 0 :
		case this.arr[i-1] == 0 && this.arr[i] == 1 :
			start, end = i, i
		case this.arr[i-1] == 1 && this.arr[i] == 1 :
			end = i
		case this.arr[i-1] == 1 && this.arr[i] == 0 :
			result = append(result, []int{start, end})
		}
	}
	return result
}
