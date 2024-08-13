package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	raHttp "github.com/joshraphael/go-retroachievements/http"
	"github.com/joshraphael/go-retroachievements/models"
	"github.com/stretchr/testify/require"
)

type testObj struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestResponseObject(tt *testing.T) {
	tests := []struct {
		name        string
		code        int
		objBody     testObj
		errBody     models.ErrorResponse
		readerInput func(inputBytes []byte, errorBytes []byte) string
		assert      func(t *testing.T, obj *testObj, err error)
	}{
		{
			name: "fail to decode response",
			code: http.StatusOK,
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return "?"
			},
			assert: func(t *testing.T, obj *testObj, err error) {
				require.Nil(t, obj)
				require.EqualError(t, err, "invalid character '?' looking for beginning of value")
			},
		},
		{
			name: "success",
			code: http.StatusOK,
			objBody: testObj{
				ID:   "8710298370",
				Name: "test",
			},
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return string(inputBytes)
			},
			assert: func(t *testing.T, obj *testObj, err error) {
				require.Equal(t, "8710298370", obj.ID)
				require.Equal(t, "test", obj.Name)
				require.NoError(t, err)
			},
		},
		{
			name: "not found",
			code: http.StatusNotFound,
			objBody: testObj{
				ID:   "8710298370",
				Name: "test",
			},
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return string(inputBytes)
			},
			assert: func(t *testing.T, obj *testObj, err error) {
				require.Nil(t, obj)
				require.NoError(t, err)
			},
		},
		{
			name: "not authorized - error marshalling response",
			code: http.StatusUnauthorized,
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return "?"
			},
			assert: func(t *testing.T, obj *testObj, err error) {
				require.Nil(t, obj)
				require.EqualError(t, err, "invalid character '?' looking for beginning of value")
			},
		},
		{
			name: "not authorized",
			code: http.StatusUnauthorized,
			errBody: models.ErrorResponse{
				Message: "test",
				Errors: []models.ErrorDetail{
					{
						Status: http.StatusUnauthorized,
						Code:   "unauthorized",
						Title:  "Not Authorized",
					},
				},
			},
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return string(errorBytes)
			},
			assert: func(t *testing.T, obj *testObj, err error) {
				require.Nil(t, obj)
				require.EqualError(t, err, "error responses: [401] Not Authorized")
			},
		},
		{
			name: "unknown error",
			code: http.StatusInternalServerError,
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return string(errorBytes)
			},
			assert: func(t *testing.T, obj *testObj, err error) {
				require.Nil(t, obj)
				require.EqualError(t, err, "unknown error returned: 500")
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			objBytes, err := json.Marshal(test.objBody)
			require.NoError(t, err)
			errBytes, err := json.Marshal(test.errBody)
			require.NoError(t, err)
			reader := bytes.NewBufferString(test.readerInput(objBytes, errBytes))
			readerCloser := io.NopCloser(reader)
			r := &http.Response{
				StatusCode: test.code,
				Body:       readerCloser,
			}
			obj, err := raHttp.ResponseObject[testObj](r)
			test.assert(t, obj, err)
		})
	}
}

func TestResponseList(tt *testing.T) {
	tests := []struct {
		name        string
		code        int
		listBody    []testObj
		errBody     models.ErrorResponse
		readerInput func(inputBytes []byte, errorBytes []byte) string
		assert      func(t *testing.T, list []testObj, err error)
	}{
		{
			name: "fail to decode response",
			code: http.StatusOK,
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return "?"
			},
			assert: func(t *testing.T, list []testObj, err error) {
				require.Nil(t, list)
				require.EqualError(t, err, "invalid character '?' looking for beginning of value")
			},
		},
		{
			name: "success",
			code: http.StatusOK,
			listBody: []testObj{
				{
					ID:   "8710298370",
					Name: "test",
				},
				{
					ID:   "1212121",
					Name: "newName",
				},
			},
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return string(inputBytes)
			},
			assert: func(t *testing.T, list []testObj, err error) {
				require.Len(t, list, 2)
				require.Equal(t, "8710298370", list[0].ID)
				require.Equal(t, "test", list[0].Name)
				require.Equal(t, "1212121", list[1].ID)
				require.Equal(t, "newName", list[1].Name)
				require.NoError(t, err)
			},
		},
		{
			name: "not authorized - error marshalling response",
			code: http.StatusUnauthorized,
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return "?"
			},
			assert: func(t *testing.T, list []testObj, err error) {
				require.Nil(t, list)
				require.EqualError(t, err, "invalid character '?' looking for beginning of value")
			},
		},
		{
			name: "not authorized",
			code: http.StatusUnauthorized,
			errBody: models.ErrorResponse{
				Message: "test",
				Errors: []models.ErrorDetail{
					{
						Status: http.StatusUnauthorized,
						Code:   "unauthorized",
						Title:  "Not Authorized",
					},
				},
			},
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return string(errorBytes)
			},
			assert: func(t *testing.T, list []testObj, err error) {
				require.Nil(t, list)
				require.EqualError(t, err, "error responses: [401] Not Authorized")
			},
		},
		{
			name: "unknown error",
			code: http.StatusInternalServerError,
			readerInput: func(inputBytes []byte, errorBytes []byte) string {
				return string(errorBytes)
			},
			assert: func(t *testing.T, list []testObj, err error) {
				require.Nil(t, list)
				require.EqualError(t, err, "unknown error returned: 500")
			},
		},
	}
	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			listBytes, err := json.Marshal(test.listBody)
			require.NoError(t, err)
			errBytes, err := json.Marshal(test.errBody)
			require.NoError(t, err)
			reader := bytes.NewBufferString(test.readerInput(listBytes, errBytes))
			readerCloser := io.NopCloser(reader)
			r := &http.Response{
				StatusCode: test.code,
				Body:       readerCloser,
			}
			list, err := raHttp.ResponseList[testObj](r)
			test.assert(t, list, err)
		})
	}
}
