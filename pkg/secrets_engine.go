package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	adConfigEndpoint  = "ad/config"
	adRoleEndpoint    = "ad/roles"
	adCredEndpoint    = "ad/creds"
	adLibraryEndpoint = "ad/library"
)

func (sdk vaultSDK) CreateADConfig(cfg ADConfig) (string, error) {
	if err := cfg.Validate(); err != nil {
		return "", err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, adConfigEndpoint)
	_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "created", nil
}

func (sdk vaultSDK) ViewADConfig() (ADConfigResponse, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, adConfigEndpoint)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return ADConfigResponse{}, err
	}

	var adr ADConfigResponse
	if err := json.Unmarshal(body, &adr); err != nil {
		return ADConfigResponse{}, err
	}

	return adr, nil
}

func (sdk vaultSDK) DeleteADConfig() (string, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, adConfigEndpoint)
	_, _, err := sdk.processRequest(http.MethodDelete, url, nil, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "deleted", nil
}

func (sdk vaultSDK) CreateADRole(role ADRole, name string) (string, error) {
	if err := role.Validate(); err != nil {
		return "", err
	}

	data, err := json.Marshal(role)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, adRoleEndpoint, name)
	_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "created", nil
}

func (sdk vaultSDK) ViewADRoles() (ADRoleResponse, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, adRoleEndpoint)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return ADRoleResponse{}, err
	}

	var adrr ADRoleResponse
	if err := json.Unmarshal(body, &adrr); err != nil {
		return ADRoleResponse{}, err
	}

	return adrr, nil
}

func (sdk vaultSDK) ViewADRole(name string) (ADRoleResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, adRoleEndpoint, name)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return ADRoleResponse{}, err
	}

	var adrr ADRoleResponse
	if err := json.Unmarshal(body, &adrr); err != nil {
		return ADRoleResponse{}, err
	}

	return adrr, nil
}

func (sdk vaultSDK) DeleteADRole(name string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, adRoleEndpoint, name)
	_, _, err := sdk.processRequest(http.MethodDelete, url, nil, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "deleted", nil
}

func (sdk vaultSDK) ViewADCreds(name string) (ADCredResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, adCredEndpoint, name)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return ADCredResponse{}, err
	}

	var adcr ADCredResponse
	if err := json.Unmarshal(body, &adcr); err != nil {
		return ADCredResponse{}, err
	}

	return adcr, nil
}

func (sdk vaultSDK) CreateADLibrary(adl ADLibrary, name string) (string, error) {
	if err := adl.Validate(); err != nil {
		return "", err
	}

	data, err := json.Marshal(adl)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, adLibraryEndpoint, name)
	_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "created", nil
}

func (sdk vaultSDK) ViewADLibrary(name string) (ADLibrary, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, adConfigEndpoint)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return ADLibrary{}, err
	}

	var adl ADLibrary
	if err := json.Unmarshal(body, &adl); err != nil {
		return ADLibrary{}, err
	}

	return adl, nil
}

func (sdk vaultSDK) ViewADLibraries() (ADLibrary, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, adLibraryEndpoint)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return ADLibrary{}, err
	}

	var adl ADLibrary
	if err := json.Unmarshal(body, &adl); err != nil {
		return ADLibrary{}, err
	}

	return adl, nil
}

func (sdk vaultSDK) DeleteADLibrary(name string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, adConfigEndpoint, name)
	_, _, err := sdk.processRequest(http.MethodDelete, url, nil, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "deleted", nil
}

func (sdk vaultSDK) ADCheckOut(name, ttl string) (ADCheckOut, error) {
	req := map[string]interface{}{"ttl": ttl}
	data, err := json.Marshal(req)
	if err != nil {
		return ADCheckOut{}, err
	}

	url := fmt.Sprintf("%s/%s/%s/check-out", sdk.baseURL, adLibraryEndpoint, name)
	_, body, err := sdk.processRequest(http.MethodPost, url, data, http.StatusOK)
	if err != nil {
		return ADCheckOut{}, err
	}

	var adc ADCheckOut
	if err := json.Unmarshal(body, &adc); err != nil {
		return ADCheckOut{}, err
	}

	return adc, nil
}

func (sdk vaultSDK) ADCheckIn(name string, serviceAccountNames []string, manage bool) (ADCheckIn, error) {
	req := map[string]interface{}{"service_account_names": serviceAccountNames}
	data, err := json.Marshal(req)
	if err != nil {
		return ADCheckIn{}, err
	}
	var manageURL string
	if manage {
		manageURL = "manage/"
	}
	url := fmt.Sprintf("%s/%s/%s%s/check-in", sdk.baseURL, adLibraryEndpoint, manageURL, name)

	_, body, err := sdk.processRequest(http.MethodPost, url, data, http.StatusOK)
	if err != nil {
		return ADCheckIn{}, err
	}

	var adc ADCheckIn
	if err := json.Unmarshal(body, &adc); err != nil {
		return ADCheckIn{}, err
	}

	return adc, nil
}

func (sdk vaultSDK) ADCheckStatus(name string) (ADCheckStatus, error) {
	url := fmt.Sprintf("%s/%s/%s/status", sdk.baseURL, adLibraryEndpoint, name)

	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return ADCheckStatus{}, err
	}

	var adc ADCheckStatus
	if err := json.Unmarshal(body, &adc); err != nil {
		return ADCheckStatus{}, err
	}

	return adc, nil
}
