package log

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect_and_interfaces/richerror"
	"reflect_and_interfaces/simpleerror"
	"time"
)

type Log struct {
	Errors []richerror.RichError
}

func (l *Log) Print() {
	for i, e := range l.Errors {
		fmt.Println("i", i, "error", e)
	}
}

func (l *Log) Append(err error) {
	var finalError richerror.RichError
	// type assertion
	rErr, ok := err.(*richerror.RichError)
	if ok {
		finalError = *rErr
	} else {
		sErr, ok := err.(*simpleerror.SimpleError)
		if ok {
			finalError = richerror.RichError{
				Message:   sErr.Output,
				MetaData:  nil,
				Operation: sErr.Operation,
				Time:      time.Now(),
			}
		} else {
			finalError = richerror.RichError{
				Message:   err.Error(),
				MetaData:  nil,
				Operation: "unknown",
				Time:      time.Now(),
			}
		}
	}
	l.Errors = append(l.Errors, finalError)
}

func (l *Log) Save() {
	// for i, e := range l.Errors {
	// 	fmt.Printf("i: %d, operation: %s, message: %s, meta-data: %+v\n",
	// 		i, e.Operation, e.Message, e.MetaData)
	// }
	var fileHandler *os.File
	if f, oErr := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777); oErr != nil {
		fmt.Println("Error opening file:", oErr)

		return
	} else {
		fileHandler = f
		defer fileHandler.Close()
	}

	data, _ := json.Marshal(l.Errors)
	fileHandler.Write(data)

}
