package main

import (
	"container/heap"
	"errors"
	"fmt"
	"os"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []*Present

func (ph PresentHeap) Len() int { return len(ph) }

func (ph PresentHeap) Less(i, j int) bool {
	if ph[i].Value != ph[j].Value {
		return ph[i].Value >= ph[j].Value
	}
	return ph[i].Size <= ph[j].Size
}

func (ph PresentHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}

func (ph *PresentHeap) Push(x any) {
	present := x.(*Present)
	*ph = append(*ph, present)
}

func (ph *PresentHeap) Pop() any {
	old := *ph
	n := len(old)
	present := old[n-1]
	old[n-1] = nil
	*ph = old[0 : n-1]
	return present
}

func (ph *PresentHeap) update(present *Present, value int, size int) {
	present.Value = value
	present.Size = size
	heap.Fix(ph, present.Value)
}

func getNCoolestPresents(phslice []Present, n int) ([]Present, error) {
	if n < 0 || n > len(phslice) {
		return nil, errors.New("n should be positive and less than size of slice")
	}
	coolest := make([]Present, 0)
	//ph := make(PresentHeap, len(phslice))
	var ph PresentHeap
	for i := 0; i < len(phslice); i++ {
		//ph[i] = &phslice[i]
		temp := &phslice[i]
		ph.Push(temp)
	}
	heap.Init(&ph)
	l := ph.Len()
	for i := 0; i < n && i < l; i++ {
		popped := heap.Pop(&ph).(*Present)
		coolest = append(coolest, *popped)
	}
	return coolest, nil
}

func main() {
	var p1 Present
	p1.Value = 5
	p1.Size = 1

	var p2 Present
	p2.Value = 4
	p2.Size = 5

	var p3 Present
	p3.Value = 3
	p3.Size = 1

	var p4 Present
	p4.Value = 5
	p4.Size = 2

	var p5 Present
	p5.Value = 3
	p5.Size = 2

	var p6 Present
	p6.Value = 2
	p6.Size = 2

	var p7 Present
	p7.Value = 3
	p7.Size = 1

	var p8 Present
	p8.Value = 3
	p8.Size = 3

	var p9 Present
	p9.Value = -3
	p9.Size = 3

	phslice := make([]Present, 0)
	phslice = append(phslice, p1, p2, p3, p4, p5, p6, p7, p8, p9)

	for _, value := range phslice {
		fmt.Printf("Original (%d,%d)\n", value.Value, value.Size)
	}

	coolest, err := getNCoolestPresents(phslice, len(phslice))
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Coolest")
	for _, value := range coolest {
		fmt.Printf("(%d,%d)\n", value.Value, value.Size)
	}
}
