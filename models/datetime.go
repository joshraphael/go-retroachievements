package models

import (
	"fmt"
	"strings"
	"time"
)

// DateTime is a time data structure that can be used for string timestamps formatted as "2006-01-02 15:04:05"
type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(time.DateTime, s)
	*dt = DateTime{nt}
	return
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	return []byte(dt.String()), nil
}

func (dt *DateTime) String() string {
	return fmt.Sprintf("%q", dt.Format(time.DateTime))
}
