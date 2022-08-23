package main

import "fmt"

func main() {

	//Arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	//Slices
	slice := []int{1, 2, 3}

	fmt.Println(slice)

	slice = append(slice, 4, 42, 27)

	fmt.Println(slice)

	s2 := slice[1:]
	s3 := slice[:2]
	s4 := slice[1:2]
	fmt.Println(s2, s3, s4)

	//Maps
	m := map[string]int{"foo": 42}

	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	//Structs

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 7
	u.FirstName = "Naruto"
	u.LastName = "Uzumaki"
	fmt.Println(u.FirstName)

	u2 := user{ID: 8,
		FirstName: "Sasuke",
		LastName:  "Uchiha", //end with comma so that compilers doesn't think is the end
	}
	fmt.Println(u2)

	u3 := user{9, "Sakura", "Haruno"}
	fmt.Println(u3)
}
