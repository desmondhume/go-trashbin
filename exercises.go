package main

import "fmt"

func main() {
	var arr [5]int
	fmt.Println("emp:", arr)

	arr[4] = 3
	fmt.Println("arr:", arr)

	b := [3]int{1, 2, 3}
	fmt.Println("b:", b)

	alphabet := [3]int{4, 5, 6}
	var c [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = alphabet[j]
		}
	}
	fmt.Println("c:", c)

	// Slices

	s := []string{"a", "b", "c"}
	fmt.Println("slice:", s)
	s = append(s, "ciao", "lol", "eheh", "ciaobello")
	fmt.Println("slice:", s)

	sc := make([]string, len(s))
	copy(sc, s)
	fmt.Println("slice copy:", sc)
	fmt.Println("slice copy:", sc[3:5])

	// Maps
	m := make(map[string]interface{})
	m["first_name"] = "Donald"
	m["last_name"] = "Draper"
	m["age"] = 35
	fmt.Println("map:", m)

	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 1000)
	fmt.Println("Total:", total)

	donald := make(map[string]interface{})
	donald["first_name"] = "Donald"
	donald["last_name"] = "Draper"
	donald["age"] = 35

	roger := make(map[string]interface{})
	roger["first_name"] = "Roger"
	roger["last_name"] = "Sterling"
	roger["age"] = 60
	pp(donald, roger)
}

func pp(rows ...map[string]interface{}) {
	for i, row := range rows {
		fmt.Println("Reading row", i)
		for k, v := range row {
			fmt.Println(k, "is", v)
		}
	}
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
