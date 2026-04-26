package main

import (
	"fmt"
	"os"
	"reflect_and_interfaces/log"
	"reflect_and_interfaces/richerror"
	"strconv"
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

		fmt.Println(OErr.Error())
	}

	user, gErr := getUserByID(0)
	// gErr.Error()
	if gErr != nil {
		// type assertion => برای اینکه به تایپ واقعی برسیم ، به اون کانکریت تایپه ، کانکریت ولیو که پاس داده شده به عنوان اینترفیس برسیم
		rErr, ok := gErr.(*richerror.RichError)
		if ok {
			logger.Append(rErr)
		} else {
			logger.Append(&richerror.RichError{
				Message:   gErr.Error(),
				MetaData:  nil,
				Operation: "unknown",
			})
		}
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

	}

	_, g3Err := getUserByIDThree(0)
	if g3Err != nil {
		logger.Append(g3Err)
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
		}
	}

	return User{}, nil
}

type simpleError struct {
	Output    string
	Operation string
}

func (s simpleError) Error() string {
	return "output: " + s.Output + ", operation: " + s.Operation
}

func getUserByIDThree(id int) (User, error) {
	if id == 0 {
		return User{}, &simpleError{
			Output: "id is 0",
			Operation: "getUserByIDThree",
		}
	}

	return User{}, nil
}
