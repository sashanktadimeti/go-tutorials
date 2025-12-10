package main

import "sort"

type Person struct {
	name string
	age  int
}
type PersonList []Person

func (pl PersonList) Len() int {
	return len(pl)
}
func (pl PersonList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}
func (pl PersonList) Less(i, j int) bool {
	return pl[i].age < pl[j].age
}



func main() {
	people := PersonList{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}
	sort.Sort(people)
	for _, person := range people {
		println(person.name, person.age)
	}
}

