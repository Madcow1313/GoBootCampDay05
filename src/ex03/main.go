package main

import "fmt"

type Present struct {
	Value int
	Size  int
}

type Backpack struct {
	presents []Present
	value    int
}

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func fillWeightTable(presents []Present, capacity int) [][]Backpack {
	bp := make([][]Backpack, len(presents)+1)
	for i := 0; i < len(presents)+1; i++ {
		bp[i] = make([]Backpack, capacity+1)
	}
	for i := 0; i < len(presents)+1; i++ {
		for j := 0; j < capacity+1; j++ {
			if i == 0 || j == 0 {
				bp[i][j] = Backpack{presents: []Present{}, value: 0}
			} else if i == 1 {
				if presents[0].Size <= j {
					var pr []Present
					pr = append(pr, presents[0])
					bp[1][j] = Backpack{presents: pr, value: presents[0].Value}
				} else {
					bp[1][j] = Backpack{presents: []Present{}, value: 0}
				}
			} else {
				if presents[i-1].Size > j {
					bp[i][j] = bp[i-1][j]
				} else {
					newValue := presents[i-1].Value + bp[i-1][j-presents[i-1].Size].value
					if bp[i-1][j].value > newValue {
						bp[i][j] = bp[i-1][j]
					} else {
						newPresents := bp[i-1][j-presents[i-1].Size].presents
						newPresents = append(newPresents, presents[i-1])
						bp[i][j] = Backpack{presents: newPresents, value: newValue}
					}
				}
			}
		}
	}
	return bp
}

func findMostValue(bp [][]Backpack, capacity int, l int) []Present {
	biggestBp := Backpack{presents: []Present{}, value: 0}
	for i := 0; i < l+1; i++ {
		for j := 0; j < capacity+1; j++ {
			if bp[i][j].value > biggestBp.value {
				biggestBp = bp[i][j]
			}
		}
	}
	mostValue := biggestBp.presents
	return mostValue
}

func grabPresents(presents []Present, capacity int) []Present {

	bp := fillWeightTable(presents, capacity)
	// for i := 0; i < len(presents)+1; i++ {
	// 	for j := 0; j < capacity+1; j++ {
	// 		fmt.Println(temp[i][j].presents, temp[i][j].value)
	// 	}
	// }
	mostValue := findMostValue(bp, capacity, len(presents))
	return mostValue
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

	phslice := make([]Present, 0)
	phslice = append(phslice, p1, p2, p3, p4, p5, p6)
	mostValuePresents := grabPresents(phslice, 2)
	fmt.Println(mostValuePresents)

	/*empty*/
	fmt.Println("empty slice", grabPresents([]Present{}, 1))

	/*all sizes > capacity*/
	phslice2 := make([]Present, 0)
	phslice2 = append(phslice2, p2)
	fmt.Println("presents size > capacity", grabPresents(phslice2, 1))
}
