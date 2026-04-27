package main

import (
	"fmt"
	"os"
	"reflect_and_interfaces/log"
	"reflect_and_interfaces/richerror"
	"reflect_and_interfaces/simpleerror"
	"strconv"
	"time"
)

type User struct {
	ID   uint
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User{id: %d, name: %s}", u.ID, u.Name)
}

func main() {
	logger := log.Log{}

	u := User{ID: 123, Name: "Amirhossein"}
	fmt.Println(u)
	// fmt.Stringer

	_, OErr := os.OpenFile("./storage/data.txt", os.O_RDWR, 0777)
	if OErr != nil {
		logger.Append(OErr)
		// fmt.Println(OErr.Error())
		fmt.Println(OErr.Error())

		fmt.Println(OErr)
	}

	user, gErr := getUserByID(0)
	// gErr.Error()
	if gErr != nil {
		logger.Append(gErr)
		fmt.Println(gErr.Error())
		// // type assertion => برای اینکه به تایپ واقعی برسیم ، به اون کانکریت تایپه ، کانکریت ولیو که پاس داده شده به عنوان اینترفیس برسیم
		// rErr, ok := gErr.(*richerror.RichError)
		// if ok {
		// 	logger.Append(rErr)
		// } else {
		// 	logger.Append(&richerror.RichError{
		// 		Message:   gErr.Error(),
		// 		MetaData:  nil,
		// 		Operation: "unknown",
		// 	})
		// }
		// fmt.Println(err.Error())
		// logger.Errors = append(logger.Errors, gErr)
		// logger.Append(gErr)
	}

	_, g2Err := getUserByIDTwo(0)
	fmt.Println("operation", g2Err.Operation, g2Err.Message, g2Err.MetaData)

	if g2Err != nil {
		// fmt.Println(err.Error())
		// logger.Errors = append(logger.Errors, gErr)
		logger.Append(g2Err)
		fmt.Println(g2Err.Error())

	}

	_, g3Err := getUserByIDThree(0)
	if g3Err != nil {
		logger.Append(g3Err)
		fmt.Println(g3Err.Error())

	}

	logger.Save()

	fmt.Println("user", user)
}

// abstraction type
func getUserByID(id int) (User, error) {
	if id == 0 {
		return User{}, &richerror.RichError{
			Message: "id is not valid",
			MetaData: map[string]string{
				"id": strconv.Itoa(id),
			},
			Operation: "getUserByID",
			Time:      time.Now(),
		}
	}

	return User{}, nil
}

// concrete type تایپ واقعی
func getUserByIDTwo(id int) (User, *richerror.RichError) {
	if id == 0 {
		return User{}, &richerror.RichError{
			Message: "id is not valid",
			MetaData: map[string]string{
				"id": strconv.Itoa(id),
			},
			Operation: "getUserByID",
			Time:      time.Now(),
		}
	}

	return User{}, nil
}

func getUserByIDThree(id int) (User, error) {
	if id == 0 {
		return User{}, &simpleerror.SimpleError{
			Output:    "id is 0",
			Operation: "getUserByIDThree",
		}
	}

	return User{}, nil
}
