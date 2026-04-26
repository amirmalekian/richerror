package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func main() {

	name := "amirhossein"
	stringReader := strings.NewReader(name)

	// scanner := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(stringReader)

	scanner.Scan()
	fmt.Println("output")
	fmt.Println(scanner.Text())

	var scores = Int{6, 9, 12, 59, 3, 103, 56}

	fmt.Println("before sort:", scores)

	sort.Sort(scores)

	fmt.Println("after sort:", scores)

	usersMap := userStore{
		3: {ID: 3, Name: "Ali"},
		1: {ID: 1, Name: "Zahra"},
		4: {ID: 4, Name: "Mahsa"},
		2: {ID: 2, Name: "Reza"},
	}

	fmt.Println("\nBefore sort (map iteration order is random):")
	for id, u := range usersMap {
		fmt.Printf("ID:%d Name:%s\n", id, u.Name)
	}

	us := userSorter{data: usersMap}

	for k := range usersMap {
		us.keys = append(us.keys, k)
	}

	sort.Sort(us)

	fmt.Println("\nAfter sort by Name (map view):")
	for _, k := range us.keys {
		u := usersMap[k]
		fmt.Printf("ID:%d Name:%s\n", u.ID, u.Name)
	}

	var keys []uint
	for k := range usersMap {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		// return keys[i] < keys[j]
		return usersMap[keys[i]].Name < usersMap[keys[j]].Name
	})

	fmt.Println("\nSort with sort.Slice() without using less, len, swap fn:")
	for _, k := range keys {
		// fmt.Println(usersMap[k])
		fmt.Printf("ID:%d Name:%s\n", usersMap[k].ID, usersMap[k].Name)

	}
}

type Int []int

func (in Int) Len() int {
	return len(in)
}

func (in Int) Less(i, j int) bool {
	return in[i] < in[j]
}

func (in Int) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

// func (user userStore) Len() int {
// 	return len(user)
// }

// func (user userStore) Less(i, j int) bool {

// }

// func (user userStore) Swap(i, j int) {

// }

type User struct {
	ID   uint
	Name string
}

type userStore map[uint]User

type userSorter struct {
	data userStore
	keys []uint
}

func (us userSorter) Len() int {
	return len(us.keys)
}

func (us userSorter) Less(i, j int) bool {
	u1 := us.data[us.keys[i]]
	u2 := us.data[us.keys[j]]

	return u1.Name < u2.Name
	// return u1.ID < u2.ID
}

func (us userSorter) Swap(i, j int) {
	us.keys[i], us.keys[j] = us.keys[j], us.keys[i]

}
