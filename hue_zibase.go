package main

/*
* This script is the first implementation of a bridge between the Zibase Home automation box, and Philips Hue lights
* As it is not possible to send PUT request from the Zibase box, only GET requests are possible, we need a bridge that will convert a request to the 
* corresponding Philips Hue API
*/

import (
    "net/http"
    "fmt"
    "bytes"
    "io/ioutil"
   "encoding/json"
)

type light struct {
    State state
    Type string
    ModelId string
    UniqueId string
    SwVersion string
    PointSymbol PointSymbol
}
type state struct {
	On bool
	Bri uint16
	Hue uint16
	Sat uint16
	XY []float32
	Alert string
	Effect string
	ColorMode string
	Reachable bool
}

type PointSymbol struct {
	p1 string
	p2 string
	p3 string
	p4 string
	p5 string
	p6 string
	p7 string
	p8 string
}

/*
* This method is executed when a request is sent to /light
* This used to manage lights. Right now, we can only switch on or off specific Philips Hue lights. 
* GET ip:3000/light?id=<id of the light>&switch=on : Will switch on the corresponding light
* GET ip:3000/light?id=<id of the light>&switch=off : Will switch off the corresponding light
* GET ip:3000/light?id=<id of the light : According to the current state of the light, will switch on or off. 
*/
func response(rw http.ResponseWriter, request *http.Request) {
	baseUrl := "http://<ip of your philips hue module>/api/<philips hue user>"
	url := fmt.Sprintf("%s/lights/%s/state", baseUrl, request.URL.Query().Get("id"))

    data := ""
	if request.URL.Query().Get("state") == "on" {
		data = `{"on":true}`	
	} else if request.URL.Query().Get("state") == "off"{
		data = `{"on":false}`
	} else {
		resp, _ := http.Get(fmt.Sprintf("%s/lights/%s", baseUrl, request.URL.Query().Get("id")))
		body, _ := ioutil.ReadAll(resp.Body)
		
		defer resp.Body.Close()
		var lgt light
		json.Unmarshal(body, &lgt)
		if lgt.State.On == false {
			data = `{"on":true}`
		} else {
			data = `{"on":false}`
		}
	}
    
    var jsonStr = []byte(data)
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}

func main() {
    http.HandleFunc("/light", response)
    http.ListenAndServe(":3000", nil)
}
