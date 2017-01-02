package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

var VelMap map[string]string

func (v *Vel) New(s string) *Vel {
	if strings.HasPrefix(s, "http") {
		var c = &http.Client{
			Timeout: time.Second * 10,
		}
		req, err := http.NewRequest("GET", s, nil)
		req.Header.Add("User-Agent", "joomla-velcheck 0.1 author: keith@keithp.net")
		resp, err := c.Do(req)
		if err != nil {
			log.Fatalf("Could not load %s: %s", s, err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		v.Raw = body
		return v
	}
	raw, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatalf("Could not load %s: %s", s, err.Error())
	}
	v.Raw = raw
	return v
}

func (v *Vel) Marshall() *Vel {
	mutex.Lock()
	// Actually marshall the data from JSON
	j := &JsonVel{}
	err2 := json.Unmarshal(v.Raw, j)
	if err2 != nil {
		log.Fatal(err2)
	}
	v.Json = j

	// Populate the component list
	for i := range v.Json.Data {
		if v.Json.Data[i].Com_whatever != "" {
			if VelMap[v.Json.Data[i].Com_whatever] != "" {
				isLess, err := lessThan(v.Json.Data[i].Version_effected, VelMap[v.Json.Data[i].Com_whatever])
				if err != nil {
					log.Printf("Error in comparing: %s", err)
				}
				if !isLess {
					VelMap[strings.ToLower(strings.Replace(v.Json.Data[i].Com_whatever, " ", "", 1))] = v.Json.Data[i].Version_effected
				}
			}
			if v.Json.Data[i].Version_effected != "" {
				VelMap[strings.ToLower(strings.Replace(v.Json.Data[i].Com_whatever, " ", "", 1))] = v.Json.Data[i].Version_effected
			} else {
				VelMap[strings.ToLower(strings.Replace(v.Json.Data[i].Com_whatever, " ", "", 1))] = "9999"
			}
		}
	}
	mutex.Unlock()
	return v
}

func init() {
	VelMap = make(map[string]string)
}
