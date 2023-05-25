package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"receipt-processor-challeng/src/model"
	"testing"
)

func Test_ReceiptPoints_BadId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := GetRouter()
	req, _ := http.NewRequest("GET", "/receipts/bad_id/points", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
}

func Test_ProcessReceipt_BadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := GetRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	body, err := os.Open("bad-json.json")
	if err != nil {
		t.Error()
		return
	}
	defer func(body *os.File) {
		err := body.Close()
		if err != nil {

		}
	}(body)

	resp, err := http.Post(server.URL+"/receipts/process", "application/json", body)
	if err != nil {
		t.Error()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	assert.Equal(t, "400 Bad Request", resp.Status)
}

func Test_ProcessReceipt_ThenPointReceipt(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := GetRouter()
	server := httptest.NewServer(router)
	defer server.Close()

	body, err := os.Open("../../examples/example-receipt-1.json")
	if err != nil {
		t.Error()
		return
	}
	defer func(body *os.File) {
		err := body.Close()
		if err != nil {

		}
	}(body)

	resp, err := http.Post(server.URL+"/receipts/process", "application/json", body)
	if err != nil {
		t.Error()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	assert.Equal(t, "200 OK", resp.Status)

	var id *model.ProcessReceiptReturn
	err = json.NewDecoder(resp.Body).Decode(&id)
	if err != nil {
		return
	}

	resp, err = http.Get(server.URL + "/receipts/" + id.Id + "/points")
	if err != nil {
		t.Error()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error()
		}
	}(resp.Body)

	assert.Equal(t, "200 OK", resp.Status)

	var points *model.GetPointsReturn
	err = json.NewDecoder(resp.Body).Decode(&points)
	if err != nil {
		return
	}

	assert.Equal(t, 28, points.Points)

}
