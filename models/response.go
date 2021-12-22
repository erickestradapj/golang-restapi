package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* ===== STRUCT RESPONSE ===== */
type Response struct {
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	ContentType string
	respWrite   http.ResponseWriter
}

/* ===== DEFAULT RESPONSE ===== */
func CreateDefaultResponse(res http.ResponseWriter) Response {
	return Response{
		Status:      http.StatusOK,
		respWrite:   res,
		ContentType: "application/json",
	}
}

/* ===== NOT FOUND ===== */
func (resp *Response) NotFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource Not Found"
}

/* ===== RESPONSE TO CLIENT ===== */
func (resp *Response) Send() {
	resp.respWrite.Header().Set("Content-Type", resp.ContentType)
	resp.respWrite.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWrite, string(output))
}

/* ===== RESPONSE DATA TO CLIENT ===== */
func SendData(res http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(res)
	response.Data = data
	response.Send()
}

/* ===== RESPONSE ERROR TO CLIENT ===== */
func SendNotFound(res http.ResponseWriter) {
	response := CreateDefaultResponse(res)
	response.NotFound()
	response.Send()
}

/* ===== UNPROCESSABLE ENTITY ===== */
func (resp *Response) UnprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "UnprocessableEntity Not Found"
}

/* ===== SEND UNPROCESSABLE ENTITY ===== */
func SendUnprocessableEntity(res http.ResponseWriter) {
	response := CreateDefaultResponse(res)
	response.UnprocessableEntity()
	response.Send()
}
