package models_test

import (
	"testing"
	"time"

	"github.com/joshraphael/go-retroachievements/models"
	"github.com/stretchr/testify/require"
)

func TestDateTimeUnmarshalJSON(tt *testing.T) {
	tests := []struct {
		name   string
		input  string
		assert func(t *testing.T, date *models.DateTime, err error)
	}{
		{
			name:  "empty string default",
			input: "",
			assert: func(t *testing.T, date *models.DateTime, err error) {
				require.NotNil(t, date)
				require.True(t, date.Time.IsZero())
				require.NoError(t, err)
			},
		},
		{
			name:  "unknown bytes",
			input: "?>?>>L:",
			assert: func(t *testing.T, date *models.DateTime, err error) {
				require.NotNil(t, date)
				require.True(t, date.Time.IsZero())
				require.EqualError(t, err, "parsing time \"?>?>>L:\" as \"2006-01-02 15:04:05\": cannot parse \"?>?>>L:\" as \"2006\"")
			},
		},
		{
			name:  "successfully unmarshal",
			input: "2024-03-02 17:27:03",
			assert: func(t *testing.T, date *models.DateTime, err error) {
				ts, tErr := time.Parse(time.DateTime, "2024-03-02 17:27:03")
				require.NoError(t, tErr)
				require.NotNil(t, date)
				require.Equal(t, ts, date.Time)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			d := &models.DateTime{}
			err := d.UnmarshalJSON([]byte(test.input))
			test.assert(t, d, err)
		})
	}
}

func TestDateTimeString(tt *testing.T) {
	expectedString := "2024-03-02 17:27:03"
	t, err := time.Parse(time.DateTime, expectedString)
	require.NoError(tt, err)
	d := &models.DateTime{t}
	require.Equal(tt, `"`+expectedString+`"`, d.String())
}

func TestLongMonthDateUnmarshalJSON(tt *testing.T) {
	tests := []struct {
		name   string
		input  string
		assert func(t *testing.T, date *models.LongMonthDate, err error)
	}{
		{
			name:  "empty string default",
			input: "",
			assert: func(t *testing.T, date *models.LongMonthDate, err error) {
				require.NotNil(t, date)
				require.True(t, date.Time.IsZero())
				require.NoError(t, err)
			},
		},
		{
			name:  "unknown bytes",
			input: "?>?>>L:",
			assert: func(t *testing.T, date *models.LongMonthDate, err error) {
				require.NotNil(t, date)
				require.True(t, date.Time.IsZero())
				require.EqualError(t, err, "parsing time \"?>?>>L:\" as \"January 02, 2006\": cannot parse \"?>?>>L:\" as \"January\"")
			},
		},
		{
			name:  "successfully unmarshal",
			input: "March 02, 2024",
			assert: func(t *testing.T, date *models.LongMonthDate, err error) {
				ts, tErr := time.Parse(models.LongMonthDateFormat, "March 02, 2024")
				require.NoError(t, tErr)
				require.NotNil(t, date)
				require.Equal(t, ts, date.Time)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			d := &models.LongMonthDate{}
			err := d.UnmarshalJSON([]byte(test.input))
			test.assert(t, d, err)
		})
	}
}

func TestLongMonthDateString(tt *testing.T) {
	expectedString := "March 02, 2024"
	t, err := time.Parse(models.LongMonthDateFormat, expectedString)
	require.NoError(tt, err)
	d := &models.LongMonthDate{t}
	require.Equal(tt, `"`+expectedString+`"`, d.String())
}
