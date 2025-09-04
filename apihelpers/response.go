package apihelpers

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"mv/mvto-do/constants"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type APIRes struct {
	Status    bool        `json:"status"`
	Message   string      `json:"message"`
	Code string      `json:"code"`
	Data      interface{} `json:"data"`
}

func SendInternalServerError() (int, APIRes) {
	var apiRes APIRes
	apiRes.Status = false
	apiRes.Message = constants.ErrorCodeMap[constants.InternalServerError]
	apiRes.Code = constants.InternalServerError
	return http.StatusInternalServerError, apiRes
}

func CustomResponse(c *gin.Context, code int, data interface{}, optionalParams ...interface{}) {
	// Format optional parameters into a single string

	cRH, _ := c.Get("reqH")
	requestId, clientId, clientVersion, platform := extractReqHeaderWithReflection(cRH)

	var optionalParamsStr string
	if len(optionalParams) > 0 {
		optionalParamsStr = fmt.Sprintf(" optionalParams: %v", optionalParams)
	}

	logrus.Info("CustomResponse ", optionalParamsStr, " code: ", code, " | requestId:", requestId, " | clientId:", clientId, " | platform:", platform, " | clientVersion:", clientVersion)

	// Send the JSON response
	c.JSON(code, data)
}

func extractReqHeaderWithReflection(val interface{}) (requestId, clientId, clientVersion, platform string) {
	v := reflect.ValueOf(val)

	// Handle pointer types
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Make sure it's a struct
	if v.Kind() != reflect.Struct {
		return "", "", "", ""
	}

	// Extract RequestId field
	if requestIdField := v.FieldByName("RequestId"); requestIdField.IsValid() && requestIdField.Kind() == reflect.String {
		requestId = requestIdField.String()
	}

	// Extract ClientId field
	if clientIdField := v.FieldByName("ClientId"); clientIdField.IsValid() && clientIdField.Kind() == reflect.String {
		clientId = clientIdField.String()
	}

	//Extract ClientVersion Field
	if clientVersionField := v.FieldByName("ClientVersion"); clientVersionField.IsValid() && clientVersionField.Kind() == reflect.String {
		clientVersion = clientVersionField.String()
	}

	//Extract Platform Field
	if platformField := v.FieldByName("Platform"); platformField.IsValid() && platformField.Kind() == reflect.String {
		platform = platformField.String()
	}

	return requestId, clientId, clientVersion, platform
}
