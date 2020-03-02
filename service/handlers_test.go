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

func TestGetDetailsForCatalogItemsReturnPropeData(t *testing.T) {
	var (
		request  *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := MakeTestServer()

	targetSKU := "THINGAMAJIG12"
	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/catalog/"+targetSKU, nil)
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
		t.Errorf("Expected 1000 qty in stock, got &d", detail.QuantityInStock)
	}
	if detail.ShipsWithin != 99 {
		t.Errorf("Expected shipWithin 14 days, got %d", detail.ShipsWithin)
	}

	if detail.SKU != "THINGAMAJIG12" {
		t.Errorf("Expected SKU THINGAMAJIG12, got %s", detail.SKU)
	}

	if detail.ProductID != 1 {
		t.Errorf("Expected product ID of 1, got %d", detail.ProductID)
	}
}

func MakeTestServer() *negroni.Negroni {
	fakeClient := fakeWebClient{}
	return NewServerFromClient(fakeClient)
}

type fakeWebClient struct{}

func (client fakeWebClient) getFulfillmentStatus(sku string) (status fulfillmentStatus, err error) {
	status = fulfillmentStatus{
		SKU:             sku,
		ShipsWithin:     99,
		QuantityInStock: 1000,
	}
	return status, err
}
