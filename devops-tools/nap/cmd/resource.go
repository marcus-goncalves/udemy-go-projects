package nap

import (
	"bytes"
	"io"
	"log"
	"text/template"
)

type RestResource struct {
	Endpoint string
	Method   string
	Router   *CBRouter
}

func NewResource(endpoint string, method string, router *CBRouter) *RestResource {
	return &RestResource{
		Endpoint: endpoint,
		Method:   method,
		Router:   router,
	}
}

func (r *RestResource) RenderEndpoint(params map[string]string) string {
	if params == nil {
		return r.Endpoint
	}

	t, err := template.New("resource").Parse(r.Endpoint)
	if err != nil {
		log.Fatalln("unable to parse endpoint")
	}

	buffer := &bytes.Buffer{}
	t.Execute(buffer, params)
	endpoint, err := io.ReadAll(buffer)
	if err != nil {
		log.Fatalln("unable to read endpoint")
	}

	return string(endpoint)
}
