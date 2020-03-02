package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestGetDetailsForCatalogItemsReturnProperData(t *testing.T) {
	var (
		request  *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := MakeTestServer()

	productID := "THINGAMAJIG12"
	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/catalog/"+productID, nil)
	server.ServeHTTP(recorder, request)

	var detail catalogItem

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}

	payload, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}
	err = json.Unmarshal(payload, &detail)
	if err != nil {
		t.Errorf("Error unmarshaling respone to catalog: %v", err)
	}
	if detail.QuantityInStock != 1000 {
		t.Errorf("Expected 1000 qty in stock, got %d", detail.QuantityInStock)
	}
	if detail.ShipsWithin != 99 {
		t.Errorf("Expected shipWithin 14 days, got %d", detail.ShipsWithin)
	}

	if detail.ProductID != "THINGAMAJIG12" {
		t.Errorf("Expected SKU THINGAMAJIG12, got %s", detail.ProductID)
	}

	if detail.ListingID != 1 {
		t.Errorf("Expected product ID of 1, got %d", detail.ListingID)
	}
}

func MakeTestServer() *negroni.Negroni {
	fakeClient := fakeWebClient{}
	return NewServerFromClient(fakeClient)
}

type fakeWebClient struct{}

func (client fakeWebClient) getFulfillmentStatus(productID string) (status fulfillmentStatus, err error) {
	status = fulfillmentStatus{
		ProductID:       productID,
		ShipsWithin:     99,
		QuantityInStock: 1000,
	}
	return status, err
}
