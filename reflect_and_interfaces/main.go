package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	s1 := simpleData{
		ID:    10,
		Name:  "amirhossein",
		Email: "malekian",
	}

	s2 := simpleDataTwo{
		ID:    10,
		Name:  "amirhossein",
		Email: "malekian",
	}

	fmt.Println(Json(s1))
	fmt.Println(JsonTwo(s2))
	b1, _ := Json(s1)
	fmt.Printf("%s\n", b1)

	b2, _ := Json(s1)
	fmt.Printf("%s\n", b2)
}

func String(err error) string {
	return fmt.Sprintln(err.Error())
}

func StringTwo(err error) string {
	return fmt.Sprintln(err)
}

type simpleData struct {
	ID    uint
	Name  string
	Email string
}

func (s simpleData) MarshalJSON() ([]byte, error) {
	//return []byte(fmt.Sprintf(`{"id": %d, "name": "%s", "email": "%s"}`, s.ID, s.Name, s.Email)), nil
	return []byte(`{"id":` + strconv.Itoa(int(s.ID)) + `, "name": "` + s.Name + `", "email": "` + s.Email + `"}`), nil
	//instead of strconv.Itoa(int(s.ID)) => strconv.FormatUint(uint64(s.ID), 10)

}

type simpleDataTwo struct {
	ID    uint
	Name  string
	Email string
}

func Json(data simpleData) ([]byte, error) {
	return json.Marshal(data)
}

func JsonTwo(data simpleDataTwo) ([]byte, error) {
	return json.Marshal(data)
}
