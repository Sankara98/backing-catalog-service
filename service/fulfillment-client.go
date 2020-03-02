package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type fulfillmentClient interface {
	getFulfillmentStatus(productID string) (status fulfillmentStatus, err error)
}

type fulfillmentWebClient struct {
	rootURL string
}

func (client fulfillmentWebClient) getFulfillmentStatus(productID string) (status fulfillmentStatus, err error) {
	httpclient := &http.Client{}

	productURL := fmt.Sprintf("%s/%s", client.rootURL, productID)
	fmt.Printf("About to request SKU details from backing service: %s\n", productURL)
	req, _ := http.NewRequest("GET", productURL, nil)

	resp, err := httpclient.Do(req)
	if err != nil {
		fmt.Printf("Errored when sending request to the server: %s\n",
			err.Error())
		return
	}

	defer resp.Body.Close()
	payload, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(payload, &status)
	if err != nil {
		fmt.Println("Failed to unmarshall server response")
		return
	}

	return status, err

}
