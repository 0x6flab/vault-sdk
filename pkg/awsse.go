package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	awsRootEndpoint       = "aws/config/root"
	awsRotateRootEndpoint = "aws/config/rotate-root"
	awsLeaseEndpoint      = "aws/config/lease"
	awsRoleEndpoint       = "aws/roles/"
)

func (sdk vaultSDK) CreateAWSRootIAMCreds(cfg AWSRootIAMCreds) (string, error) {
	if err := cfg.Validate(); err != nil {
		return "", err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, awsRootEndpoint)
	_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "created", nil
}

func (sdk vaultSDK) ViewAWSRootConfig() (AWSRootIAMCreds, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, awsRootEndpoint)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return AWSRootIAMCreds{}, err
	}

	var adr AWSRootIAMCreds
	if err := json.Unmarshal(body, &adr); err != nil {
		return AWSRootIAMCreds{}, err
	}

	return adr, nil
}

func (sdk vaultSDK) AWSRotataRootCreds(creds AWSRootIAMCreds) (string, error) {
	data, err := json.Marshal(creds)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, awsRotateRootEndpoint)
	_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusOK)
	if err != nil {
		return "", err
	}

	return "rotated", nil
}

func (sdk vaultSDK) ConfigureAWSLease(lease AWSLease) (string, error) {
	if err := lease.Validate(); err != nil {
		return "", err
	}

	data, err := json.Marshal(lease)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s", sdk.baseURL, awsLeaseEndpoint)
	_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "configured", nil
}

func (sdk vaultSDK) ViewAWSLease() (AWSLeaseResponse, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, awsLeaseEndpoint)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return AWSLeaseResponse{}, err
	}

	var alr AWSLeaseResponse
	if err := json.Unmarshal(body, &alr); err != nil {
		return AWSLeaseResponse{}, err
	}

	return alr, nil
}

func (sdk vaultSDK) CreateAWSRole(role AWSRole, name string) (string, error) {
	if err := role.Validate(); err != nil {
		return "", err
	}

	data, err := json.Marshal(role)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, awsRoleEndpoint, name)
	_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "created", nil
}

// TODO Fix here
func (sdk vaultSDK) ViewAWSRoles() (AWSRoleResponse, error) {
	url := fmt.Sprintf("%s/%s", sdk.baseURL, awsRoleEndpoint)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return AWSRoleResponse{}, err
	}

	var arr AWSRoleResponse
	if err := json.Unmarshal(body, &arr); err != nil {
		return AWSRoleResponse{}, err
	}

	return arr, nil
}

func (sdk vaultSDK) ViewAWSRole(name string) (AWSRoleResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, awsRoleEndpoint, name)
	_, body, err := sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
	if err != nil {
		return AWSRoleResponse{}, err
	}

	var arr AWSRoleResponse
	if err := json.Unmarshal(body, &arr); err != nil {
		return AWSRoleResponse{}, err
	}

	return arr, nil
}

func (sdk vaultSDK) DeleteAWSRole(name string) (string, error) {
	url := fmt.Sprintf("%s/%s/%s", sdk.baseURL, awsRoleEndpoint, name)
	_, _, err := sdk.processRequest(http.MethodDelete, url, nil, http.StatusNoContent)
	if err != nil {
		return "", err
	}

	return "deleted", nil
}

func (sdk vaultSDK) CreateAWSCreds(cred AWSCreds, credType, name string) (string, error) {
	if err := cred.Validate(); err != nil {
		return "", err
	}

	data, err := json.Marshal(cred)
	if err != nil {
		return "", err
	}
	switch credType {
	case "creds":
		url := fmt.Sprintf("%s/aws/%s/%s", sdk.baseURL, credType, name)
		_, _, err = sdk.processRequest(http.MethodGet, url, nil, http.StatusOK)
		if err != nil {
			return "", err
		}

		return "created", nil
	case "sts":
		url := fmt.Sprintf("%s/aws/%s/%s", sdk.baseURL, credType, name)
		_, _, err = sdk.processRequest(http.MethodPost, url, data, http.StatusNoContent)
		if err != nil {
			return "", err
		}

		return "created", nil
	default:
		return "", ErrInvalidAWSCredType
	}

}
