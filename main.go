package main

import (
	"fmt"

	"github.com/henrikkorsgaard/go-aak-birth/CKAN"
)

func main() {
	id := "2f981f6a-33a9-4ce6-a977-d3a88b68d856"
	fmt.Println(id)
	data, err := CKAN.DatasetFromID(id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	/*

		count := ckan.GetTotal(id)
		for {
			time.Sleep(2 * time.Second)
			numberOfBirths := ckan.GetTotal(id)
			if numberOfBirths > count {
				count = numberOfBirths
				response, err := ckan.GetData(id)
				if err != nil {
					fmt.Println(err)
				}

				if response.Success {
					fmt.Println(response.Result.Records[response.Result.Total-1])
				}
			}
		}*/
}
