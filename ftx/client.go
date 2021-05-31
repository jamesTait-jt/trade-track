package ftx

import (
	"net/http"
	"net/url"
)

// The FtxClient contains all the information required to get information for a specific subaccount
type FtxClient struct {
	Client *http.Client
	ApiKey string
	Secret []byte
	Subaccount string
}

func New(apiKey string, secret string, subaccount string) *FtxClient {
	httpClient := http.Client{}
	return &FtxClient{
		Client: &httpClient,
		ApiKey: apiKey,
		Secret: []byte(secret),
		Subaccount: url.PathEscape(subaccount),
	}
}