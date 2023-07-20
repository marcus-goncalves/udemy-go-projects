package pork

import (
	"github.com/spf13/viper"
	nap "udemy-projects.com/devops-tools/nap/cmd"
)

var api *nap.API

func GithubAPI() *nap.API {
	if api == nil {
		api = nap.NewApi("http://github.com")
		token := viper.GetString("token")
		api.SetAuth(nap.NewAuthToken(token))
		api.AddResource("fork", GetForkResouce())
		api.AddResource("search", GetSearchResource())
		api.AddResource("docs", GetReadmeResource())
		api.AddResource("pullrequest", GetPullRequestResource())
	}

	return api
}
