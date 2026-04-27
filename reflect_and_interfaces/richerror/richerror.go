package richerror

import "time"

type RichError struct {
	Message   string
	MetaData  map[string]string
	Operation string
	Time      time.Time
}

func (r RichError) Error() string {
	// return fmt.Sprintf("message: %s, operation: %s, metaData: %+v",
	// 	r.Message, r.Operation, r.MetaData)
	return r.Message
}
