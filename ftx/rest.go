package ftx

import (
	"bytes"
	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

const URL = "https://ftx.com/api/"


func (client *FtxClient) signRequest(method string, path string, body []byte) *http.Request {
	timestamp := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)
	signaturePayload := timestamp + method + "/api" + path + string(body)
	signature := client.sign(signaturePayload)
	req, _ := http.NewRequest(method, URL+path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", client.ApiKey)
	req.Header.Set("FTX-TS", timestamp)
	req.Header.Set("FTX-SIGN", signature)
	if client.Subaccount != "" {
		req.Header.Set("FTX-SUBACCOUNT", client.Subaccount)
	}
	return req
}


func (client *FtxClient) _get(path string, body []byte) (*http.Response, error) {
	signedRequest := client.signRequest("GET", path, body)
	resp, err := client.Client.Do(signedRequest)
	return resp, err
}

func (client *FtxClient) _processResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error processing response: ", err)
		return err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		log.Printf("Error processing response: ", err)
		return err
	}
	return nil
}
