package apifootball

import "net/http"

type APIClient struct {
	token      string
	baseUrl    string
	httpClient *http.Client
}
