package main

var (
	dataPeople = people{
		&person{1, "Frank Ocean", 47, people{}},
		&person{2, "Paul Gascoigne", 51, people{}},
		&person{3, "Uwe Seeler", 81, people{}},
	}

	dataFriendship = friendship{
		1: []int{2, 3},
		2: []int{1},
		3: []int{3, 1},
	}
)
