package pkg

// TODO Fix here
type Response struct {
	RequestID     string `json:"request_id,omitempty"`
	LeaseID       string `json:"lease_id,omitempty"`
	Renewable     string `json:"renewable,omitempty"`
	LeaseDuration string `json:"lease_duration,omitempty"`
	Data          ADData `json:"data,omitempty"`
	WrapInfo      string `json:"wrap_info,omitempty"`
	Warnings      string `json:"warnings,omitempty"`
	Auth          string `json:"auth,omitempty"`
	// 	"wrap_info": null,
	// 	"warnings": null,
	// 	"auth": null
	// }
}

type ADConfigResponse struct {
	Response
	Data ADData `json:"data,omitempty"`
}

type ADRoleResponse struct {
	LastVaultRotation  string `json:"last_vault_rotation,omitempty"`
	PasswordLastSet    string `json:"password_last_set,omitempty"`
	ServiceNameAccount string `json:"service_account_name,omitempty"`
	TTL                int    `json:"ttl,omitempty"`
}

type ADCredResponse struct {
	CurrentPassword string `json:"current_password,omitempty"`
	LastPassword    string `json:"last_password,omitempty"`
	Username        string `json:"username,omitempty"`
}

type ADCheckOutData struct {
	Password           string `json:"password,omitempty"`
	ServiceNameAccount string `json:"service_account_name,omitempty"`
}

type ADCheckOut struct {
	Response
	Data ADCheckOutData `json:"data,omitempty"`
}

type ADCheckInData struct {
	CheckIns []string `json:"check_ins,omitempty"`
}

type ADCheckIn struct {
	Response
	Data ADCheckInData `json:"data,omitempty"`
}

// TODO Fix here
type ADCheckStatusData struct {
	// "data": {
	// 	"buzz@example.com": {
	// 	  "available": true
	// 	},
	// 	"fizz@example.com": {
	// 	  "available": false,
	// 	  "borrower_client_token": "4c653e473bf7e27c6759fccc3def20c44d776279",
	// 	  "borrower_entity_id": "631256b1-8523-9838-5501-d0a1e2cdad9c"
	// 	}
	//   },
}

type ADCheckStatus struct {
	Response
	Data ADCheckStatusData `json:"data,omitempty"`
}
