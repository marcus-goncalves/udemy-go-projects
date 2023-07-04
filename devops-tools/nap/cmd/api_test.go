package nap

import (
	"fmt"
	"net/http"
	"testing"
)

func TestApiCall(t *testing.T) {
	api := NewApi("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		return nil
	})

	res := NewResource("get", "GET", router)
	api.AddResource("get", res)
	if err := api.Call("get", nil); err != nil {
		t.Fail()
	}

	resources := api.ResourceNames()
	if len(resources) != 1 || resources[0] != "get" {
		t.Fail()
	}
}

func TestApiAuth(t *testing.T) {
	api := NewApi("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response, _ interface{}) error {
		return nil
	})

	res := NewResource("/basic-auth/{{.user}}/{{.pass}}", "GET", router)
	api.AddResource("basic-auth", res)
	api.SetAuth(&AuthBasic{
		Username: "user",
		Password: "passw0rd",
	})
	if err := api.Call("basic-auth", map[string]string{
		"user": "user",
		"pass": "passw0rd",
	}); err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
