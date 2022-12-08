package facade

import "testing"

func TestAPI(t *testing.T){
	api := NewAPI()
	apiStr := api.Test()
	if apiStr != "AModelBModel"{
		t.Fatalf("testing error")
	}
}