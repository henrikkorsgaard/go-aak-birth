package CKAN

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	server = "https://www.odaa.dk"
)

type response struct {
	Success bool   `json:"success"`
	Result  result `json:"result"`
}

type result struct {
	Total   int16         `json:"total"`
	ID      string        `json:"resource_id"`
	Fields  []interface{} `json:"fields"`
	Records []interface{} `json:"records"`
}

type Dataset struct {
	Name    string
	ID      string
	Fields  map[string]interface{}
	Records map[string]interface{}
	Size    int16
}

func DatasetFromID(id string) (d Dataset, err error) {
	var raw []byte
	raw, err = getData(id)

	var iface interface{}
	err = json.Unmarshal(raw, &iface)
	if m["result"]
	m := m["result"].iface.(map[string]interface{})
	r :=
		fmt.Println(m["help"])

	return
}

func getTotalNumberOfRecords(id string) int16 {
	url := fmt.Sprintf("%s/api/3/action/datastore_search?resource_id=%s", server, id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var data = response{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}

	return data.Result.Total
}

func getData(id string) (b []byte, err error) {

	limit := getTotalNumberOfRecords(id)

	url := fmt.Sprintf("%s/api/3/action/datastore_search?resource_id=%s&limit=%d", server, id, limit)
	var resp *http.Response
	resp, err = http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		b, err = ioutil.ReadAll(resp.Body)
	}
	return
}
