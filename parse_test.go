package xsd_datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestValidDateTime(t *testing.T) {
	plus7, _ := time.LoadLocation("Asia/Bangkok")
	for idx, test := range []struct {
		value    string
		datetime time.Time
	}{
		{
			value: "2006-01-02T15:04:05",
			datetime: time.Date(
				2006, 01, 02,
				15, 04, 05, 0,
				time.UTC,
			),
		},
		{
			value: "2006-01-02T15:04:05+07:00",
			datetime: time.Date(
				2006, 01, 02,
				15, 04, 05, 0,
				plus7,
			),
		},

		{
			value: "2006-01-02T15:04:05Z",
			datetime: time.Date(
				2006, 01, 02,
				15, 04, 05, 0,
				time.UTC,
			),
		},
		{
			value: "2006-01-02T15:04:05-00:00",
			datetime: time.Date(
				2006, 01, 02,
				15, 04, 05, 0,
				time.UTC,
			),
		},
		{
			value: "-2006-01-02T15:04:05",
			datetime: time.Date(
				-2006, 01, 02,
				15, 04, 05, 0,
				time.UTC,
			),
		},
		{
			value: "2006-01-02T15:04:05.99999",
			datetime: time.Date(
				2006, 01, 02,
				15, 04, 05, 999990000,
				time.UTC,
			),
		},
		{
			value: "2006-01-02T15:04:05.999999999",
			datetime: time.Date(
				2006, 01, 02,
				15, 04, 05, 999999999,
				time.UTC,
			),
		},
		{
			value: "2006-01-02T15:04:05.99+07:00",
			datetime: time.Date(
				2006, 01, 02,
				15, 04, 05, 990000000,
				plus7,
			),
		},
	} {
		t.Run(fmt.Sprintf("Valid%d", idx), func(t *testing.T) {
			datetime, err := Parse(test.value)
			if err != nil {
				t.Error(err)
			}
			if !datetime.Equal(test.datetime) {
				t.Error()
			}
		})
	}
}

func TestInvalidDateTime(t *testing.T) {
	for idx, value := range []string{
		"2006-02-29T15:04:05",              // the days part, 29, is out of range
		"2006-01-02",                       // all the parts must be specified
		"2006-01-02T15:04",                 // all the parts must be specified
		"2006-01-02T15:04:05+02:0",         // all the parts must be specified
		"2006-01-02T25:04:05+02:00",        // the hours part, 25, is out of range
		"2006-01-02T15:04:05+02:000",       // specified value is too long
		"06-02-29T15:04",                   // all the parts must be specified
		"2006-01-02T15:04:05.999999999999", // too many fractional seconds specified
	} {
		t.Run(fmt.Sprintf("Invalid%d", idx), func(t *testing.T) {
			_, err := Parse(value)
			if err == nil {
				t.Error(value)
			} else {
				fmt.Println(err)
			}
		})
	}
}
