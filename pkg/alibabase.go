package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	alibabaConfigEndpoint = "alicloud/config"
	alibabaRoleEndpoint   = "alicloud/role"
)

func (sdk vaultSDK) CreateAlibabaConfig(cfg AlibabaConfig) (string, error) {
	if err := cfg.Validate(); err != nil {
		return "", err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, alibabaConfigEndpoint)
	_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "created", nil
}

func (sdk vaultSDK) ViewAlibabaConfig() (AlibabaConfig, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, alibabaConfigEndpoint)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return AlibabaConfig{}, err
	}

	var ac AlibabaConfig
	if err := json.Unmarshal(body, &ac); err != nil {
		return AlibabaConfig{}, err
	}

	return ac, nil
}

func (sdk vaultSDK) CreateAlibabaRole(role AlibabaRole, name string) (string, error) {
	if err := role.Validate(); err != nil {
		return "", err
	}

	data, err := json.Marshal(role)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, alibabaRoleEndpoint, name)
	_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "created", nil
}

// TODO Fix here
func (sdk vaultSDK) ViewAlibabaRoles() (AlibabaRoleResponse, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, alibabaRoleEndpoint)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return AlibabaRoleResponse{}, err
	}

	var adrr AlibabaRoleResponse
	if err := json.Unmarshal(body, &adrr); err != nil {
		return AlibabaRoleResponse{}, err
	}

	return adrr, nil
}

func (sdk vaultSDK) ViewAlibabaRole(name string) (AlibabaRoleResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, alibabaRoleEndpoint, name)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return AlibabaRoleResponse{}, err
	}

	var arr AlibabaRoleResponse
	if err := json.Unmarshal(body, &arr); err != nil {
		return AlibabaRoleResponse{}, err
	}

	return arr, nil
}

func (sdk vaultSDK) DeleteAlibabaRole(name string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, alibabaRoleEndpoint, name)
	_, _, err := sdk.processRequest(http.MethodDelete, url, nil, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "deleted", nil
}

func (sdk vaultSDK) CreateAlibabaRAMCreds(name string) (AlibabaRAMCreds, error) {
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, alibabaRoleEndpoint, name)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusNoContent)
	if err != nil {
		return AlibabaRAMCreds{}, err
	}

	var arc AlibabaRAMCreds
	if err := json.Unmarshal(body, &arc); err != nil {
		return AlibabaRAMCreds{}, err
	}

	return arc, nil
}
