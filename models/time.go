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
	if s == "" {
		*dt = DateTime{time.Time{}}
		return nil
	}
	nt, err := time.Parse(time.DateTime, s)
	if err != nil {
		return err
	}
	*dt = DateTime{nt}
	return nil
}

func (dt DateTime) MarshalJSON() ([]byte, error) {
	return []byte(dt.String()), nil
}

func (dt *DateTime) String() string {
	return fmt.Sprintf("%q", dt.Format(time.DateTime))
}

// DateOnly is a time data structure that can be used for string timestamps formatted as "2006-01-02 15:04:05"
type DateOnly struct {
	time.Time
}

func (do *DateOnly) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		*do = DateOnly{time.Time{}}
		return nil
	}
	nt, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return err
	}
	*do = DateOnly{nt}
	return nil
}

func (do *DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(do.String()), nil
}

func (do *DateOnly) String() string {
	return fmt.Sprintf("%q", do.Format(time.DateOnly))
}

// RFC3339NumColonTZ is a time data structure that can be used for string dates formatted as "2006-01-02T15:04:05-07:00"
type RFC3339NumColonTZ struct {
	time.Time
}

const (
	RFC3339NumColonTZFormat = "2006-01-02T15:04:05-07:00"
)

func (r *RFC3339NumColonTZ) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		*r = RFC3339NumColonTZ{time.Time{}}
		return nil
	}
	nt, err := time.Parse(RFC3339NumColonTZFormat, s)
	if err != nil {
		return err
	}
	*r = RFC3339NumColonTZ{nt}
	return nil
}

func (r RFC3339NumColonTZ) MarshalJSON() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *RFC3339NumColonTZ) String() string {
	return fmt.Sprintf("%q", r.Format(RFC3339NumColonTZFormat))
}
