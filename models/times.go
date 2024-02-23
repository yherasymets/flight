package models

import (
	"fmt"
	"strconv"
	"time"
)

type CustomTime time.Time

func (t *CustomTime) UnmarshalJSON(jsonValue []byte) error {
	s := string(jsonValue)
	s = s[1 : len(s)-1] // Remove the quotes from the string
	time, err := time.Parse(time.TimeOnly, s)
	if err != nil {
		return err
	}
	*t = CustomTime(time)
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	object := fmt.Sprint(ct)
	quotedJson := strconv.Quote(object)
	return []byte(quotedJson), nil
}

// String returns the time in the custom format
func (ct CustomTime) String() string {
	t := time.Time(CustomTime(ct))
	return fmt.Sprint(t.Format(time.TimeOnly))
}
