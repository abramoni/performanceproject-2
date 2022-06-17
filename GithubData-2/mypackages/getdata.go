package mypackages

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var responsedata metric

type postRequestData struct {
	AccessToken string `json:"accessToken"`
}

func PostRequestForAccessToken() (message string) {

	postBody, _ := json.Marshal(map[string]string{
		"username": "admin",
		"password": "customerPerf",
	})
	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	resp, err := http.Post("https://10.14.19.226/irisservices/api/v1/public/accessTokens", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var postRequestResponse postRequestData
	checkValid := json.Valid(body)

	if checkValid {
		//fmt.Println("json is valid")
		json.Unmarshal(body, &postRequestResponse)
		//fmt.Println("Access Token = ", postRequestResponse.AccessTokens)

	} else {
		fmt.Println("json invalid")
	}
	message = postRequestResponse.AccessToken
	//sb := string(body)
	//log.Printf(sb)
	return
}

func GetRequestForData(accessToken string, url string) (data []byte) {

	//url := "https://10.14.19.226/irisservices/api/v1/public/statistics/timeSeriesStats?endTimeMsecs=1654671600000&entityId=2790138600742128&metricName=kSystemUsageBytes&metricUnitType=0&range=day&rollupFunction=rate&rollupIntervalSecs=1800&schemaName=kBridgeClusterStats&startTimeMsecs=1654585200000"
	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + accessToken

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	//var responsedata metric

	checkValid := json.Valid(body)

	if checkValid {
		return body
	} else {
		fmt.Println("json invalid")
		return nil
	}

	/*if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var responsedata metric

	checkValid := json.Valid(body)

	if checkValid {
		//fmt.Println("json is valid")
		json.Unmarshal(body, &responsedata)
		fmt.Println("metric name = ", responsedata.Name)
		fmt.Println("instances = ", len(responsedata.Datavec))
		for i := 0; i < len(responsedata.Datavec); i++ {
			fmt.Print("time = ", responsedata.Datavec[i].Time)
			fmt.Println("    rate = ", responsedata.Datavec[i].Value.Data)
		}
	} else {
		fmt.Println("json invalid")
	} */
}
