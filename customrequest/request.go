package customrequest

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/albrow/forms"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

const DumpBody = true

func Get(c *gin.Context, url string, target interface{}) error {

	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	response, err := httpClient.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return err
	}
	reader := response.Body
	defer reader.Close()
	return jsoniter.NewDecoder(reader).Decode(target)
}

func GetApi(c *gin.Context, url string) (interface{}, error) {
	var target interface{}
	httpClient := &http.Client{
		Timeout: 25 * time.Second,
	}

	response, err := httpClient.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return nil, err
	}
	reader := response.Body
	defer reader.Close()
	jsoniter.NewDecoder(reader).Decode(target)
	return target, err
}

func PostAll(c *gin.Context, values *forms.Data, Url string) map[string]interface{} {
	var res map[string]interface{}

	jsonValue, err := json.Marshal(values)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}

	bytejson := bytes.NewBuffer(jsonValue)

	req, err := http.Post(Url, "application/json", bytejson)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}

	//// if len(key) > 0 {
	//// 	SetHeader(req, key, value)
	//// }

	reader := req.Body
	defer reader.Close()

	jsoniter.NewDecoder(reader).Decode(&res)

	return res
}

func PostStruct(c *gin.Context, values *forms.Data, Url string, target interface{}) {

	jsonValue, err := json.Marshal(values)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}

	bytejson := bytes.NewBuffer(jsonValue)

	req, err := http.Post(Url, "application/json", bytejson)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}

	//// if len(key) > 0 {
	//// 	SetHeader(req, key, value)
	//// }

	reader := req.Body
	defer reader.Close()

	jsoniter.NewDecoder(reader).Decode(target)
}

func PostStructInterface(c *gin.Context, values map[string]interface{}, Url string, nameheader string, token string, target interface{}) {

	jsonValue, err := json.Marshal(values)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}

	bytejson := bytes.NewBuffer(jsonValue)

	req, err := http.Post(Url, "application/json", bytejson)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}
	req.Header.Add(nameheader, token)

	reader := req.Body
	defer reader.Close()

	jsoniter.NewDecoder(reader).Decode(target)

}

func PostStructInterfaceWioContext(values map[string]interface{}, Url string, nameheader string, token string, target interface{}) {

	jsonValue, _ := json.Marshal(values)

	bytejson := bytes.NewBuffer(jsonValue)

	req, _ := http.Post(Url, "application/json", bytejson)

	req.Header.Add(nameheader, token)

	reader := req.Body
	defer reader.Close()

	jsoniter.NewDecoder(reader).Decode(target)

}

func PostStructFile(c *gin.Context, values map[string]interface{}, Url string, target interface{}) {

	jsonValue, err := json.Marshal(values)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}

	bytejson := bytes.NewBuffer(jsonValue)

	req, err := http.Post(Url, "multipart/form-data", bytejson)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}

	reader := req.Body
	defer reader.Close()

	jsoniter.NewDecoder(reader).Decode(target)

}

func Post(c *gin.Context, values map[string]interface{}, Url string) map[string]interface{} {
	var res map[string]interface{}

	jsonValue, err := json.Marshal(values)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}

	bytejson := bytes.NewBuffer(jsonValue)

	req, err := http.Post(Url, "application/json", bytejson)
	if err != nil {
		c.AsciiJSON(500, err.Error())
	}

	defer req.Body.Close()

	reader := req.Body
	defer reader.Close()

	jsoniter.NewDecoder(reader).Decode(&res)

	return res
}

func PostwithArray(c *gin.Context, urltarget string, key []string, val []string, nameheader string, token string, target interface{}) {

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	for i := 0; i < len(key); i++ {
		_ = writer.WriteField(key[i], val[i])
	}
	err := writer.Close()
	if err != nil {
		c.AsciiJSON(500, err.Error())
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", urltarget, payload)

	if err != nil {
		c.AsciiJSON(500, err.Error())
		return
	}
	req.Header.Add(nameheader, token)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		c.AsciiJSON(500, err.Error())
		return
	}
	defer res.Body.Close()

	reader := res.Body
	defer reader.Close()

	jsoniter.NewDecoder(reader).Decode(target)

}

func SetHeader(req *http.Request, key []string, value []string) *http.Request {

	for i := 0; i < len(key); i++ {
		req.Header.Set(key[i], value[i])
	}
	req.Header.Set("Content-Type", "application/json")

	return req
}

func GetHeader(c *gin.Context, key []string) (string, string) {
	var apptoken, usertoken string

	for i := 0; i < len(key); i++ {
		if key[i] == "APPTOKEN" {
			apptoken = c.Request.Header.Get(key[i])
		}
		if key[i] == "USERTOKEN" {
			usertoken = c.Request.Header.Get(key[i])
		}
	}
	return apptoken, usertoken
}
