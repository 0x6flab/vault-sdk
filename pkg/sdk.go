package pkg

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

type SDK interface {
}

type vaultSDK struct {
	baseURL string
	token   string
	client  *http.Client
}

type Config struct {
	BaseURL         string
	APIToken        string
	TLSVerification bool
}

func NewSDK(conf Config) SDK {
	return &vaultSDK{
		baseURL: conf.BaseURL,
		token:   conf.APIToken,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: !conf.TLSVerification,
				},
			},
		},
	}
}

func (sdk vaultSDK) processRequest(method, url string, data []byte, expectedRespCodes int) (http.Header, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return make(http.Header), []byte{}, err
	}

	req.Header.Set("Authorization", sdk.token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := sdk.client.Do(req)
	if err != nil {
		return make(http.Header), []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return make(http.Header), []byte{}, err
	}

	if resp.StatusCode != expectedRespCodes {
		return make(http.Header), []byte{}, nil
	}

	return resp.Header, body, nil
}
