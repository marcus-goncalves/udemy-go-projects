package nap

import "fmt"

type API struct {
	BaseUrl       string
	Resources     map[string]*RestResource
	DefaultRouter *CBRouter
	Client        *Client
}

func NewApi(baseUrl string) *API {
	return &API{
		BaseUrl:       baseUrl,
		Resources:     make(map[string]*RestResource),
		DefaultRouter: NewRouter(),
		Client:        NewClient(),
	}
}

func (a *API) SetAuth(auth Authentication) {
	a.Client.SetAuth(auth)
}

func (a *API) AddResource(name string, res *RestResource) {
	a.Resources[name] = res
}

func (a *API) Call(name string, params map[string]string, payload interface{}) error {
	res, ok := a.Resources[name]
	if !ok {
		return fmt.Errorf("resource does not exist: %s", name)
	}

	if err := a.Client.ProcessRequest(a.BaseUrl, res, params, payload); err != nil {
		return err
	}

	return nil
}

func (a *API) ResourceNames() []string {
	resources := []string{}
	for k := range a.Resources {
		resources = append(resources, k)
	}

	return resources
}
