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
			input: "\"\"",
			assert: func(t *testing.T, date *models.DateTime, err error) {
				require.NotNil(t, date)
				require.True(t, date.Time.IsZero())
				require.NoError(t, err)
			},
		},
		{
			name:  "unknown bytes",
			input: "\"?>?>>L:\"",
			assert: func(t *testing.T, date *models.DateTime, err error) {
				require.NotNil(t, date)
				require.True(t, date.Time.IsZero())
				require.EqualError(t, err, "parsing time \"?>?>>L:\" as \"2006-01-02 15:04:05\": cannot parse \"?>?>>L:\" as \"2006\"")
			},
		},
		{
			name:  "successfully unmarshal",
			input: "\"2024-03-02 17:27:03\"",
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

func TestRFC3339NumColonTZUnmarshalJSON(tt *testing.T) {
	tests := []struct {
		name   string
		input  string
		assert func(t *testing.T, date *models.RFC3339NumColonTZ, err error)
	}{
		{
			name:  "empty string default",
			input: "\"\"",
			assert: func(t *testing.T, date *models.RFC3339NumColonTZ, err error) {
				require.NotNil(t, date)
				require.True(t, date.Time.IsZero())
				require.NoError(t, err)
			},
		},
		{
			name:  "unknown bytes",
			input: "\"?>?>>L:\"",
			assert: func(t *testing.T, date *models.RFC3339NumColonTZ, err error) {
				require.NotNil(t, date)
				require.True(t, date.Time.IsZero())
				require.EqualError(t, err, "parsing time \"?>?>>L:\" as \"2006-01-02T15:04:05-07:00\": cannot parse \"?>?>>L:\" as \"2006\"")
			},
		},
		{
			name:  "successfully unmarshal",
			input: "\"2024-05-07T08:48:54+00:00\"",
			assert: func(t *testing.T, date *models.RFC3339NumColonTZ, err error) {
				ts, tErr := time.Parse(models.RFC3339NumColonTZFormat, "2024-05-07T08:48:54+00:00")
				require.NoError(t, tErr)
				require.NotNil(t, date)
				require.Equal(t, ts, date.Time)
				require.NoError(t, err)
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			d := &models.RFC3339NumColonTZ{}
			err := d.UnmarshalJSON([]byte(test.input))
			test.assert(t, d, err)
		})
	}
}

func TestRFC3339NumColonTZString(tt *testing.T) {
	expectedString := "2024-05-07T08:48:54+00:00"
	t, err := time.Parse(models.RFC3339NumColonTZFormat, expectedString)
	require.NoError(tt, err)
	d := &models.RFC3339NumColonTZ{t}
	require.Equal(tt, `"`+expectedString+`"`, d.String())
}
