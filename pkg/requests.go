package pkg

type ADConfig struct {
	// Password parameters
	TTL            int    `json:"ttl,omitempty"`             // The default password time-to-live in seconds. Once the ttl has passed, a password will be rotated the next time it's requested.
	MaxTTL         int    `json:"max_ttl,omitempty"`         // The maximum password time-to-live in seconds. No role will be allowed to set a custom ttl greater than the max_ttl.
	PasswordPolicy string `json:"password_policy,omitempty"` // Name of the password policy to use to generate passwords from. Mutually exclusive with length and formatter.
	// Connection parameters
	URL            string `json:"url,omitempty"`             // The LDAP server to connect to. Examples: ldaps://ldap.myorg.com, ldaps://ldap.myorg.com:636. This can also be a comma-delineated list of URLs, e.g. ldaps://ldap.myorg.com,ldaps://ldap.myorg.com:636, in which case the servers will be tried in-order if there are errors during the connection process. Default is ldap://127.0.0.1.
	RequestTimeout string `json:"request_timeout,omitempty"` // Timeout, in seconds, for the connection when making requests against the server before returning back an error.
	StartTLS       bool   `json:"starttls,omitempty"`        // If true, issues a StartTLS command after establishing an unencrypted connection.
	InsecureTLS    bool   `json:"insecure_tls,omitempty"`    // If true, skips LDAP server SSL certificate verification - insecure, use with caution!
	Certificate    string `json:"certificate,omitempty"`     // CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded.
	// Binding parameters
	BindDB    string `json:"binddn"`              // Distinguished name of object to bind when performing user and group search. Example: cn=vault,ou=Users,dc=example,dc=com
	BindPass  string `json:"bindpass"`            // Password to use along with binddn when performing user search.
	UserDn    string `json:"userdn,omitempty"`    // Base DN under which to perform user search. Example: ou=Users,dc=example,dc=com
	UPNDomain string `json:"upndomain,omitempty"` // The domain (userPrincipalDomain) used to construct a UPN string for authentication. The constructed UPN will appear as [binddn]@UPNDomain. Example: if upndomain=example.com and binddn=admin, the UPN string admin@example.com will be used to login to Active Directory.
	// Other parameters
	LastRotationTolerance string `json:"last_rotation_tolerance,omitempty"` // Tolerance duration to use when checking the last rotation time. Active Directory often shows a "pwdLastSet" time after Vault's because it takes a while for password updates to be propagated across a large cluster. By default, if Active Directory's last rotation time is within 5 seconds of Vault's, Vault considers itself to have been the last entity that rotated the password. However, if it's been more than 5 seconds, Vault thinks that something rotated the password out-of-band, and re-rotates it so it will "know" it and be able to continue returning it. This may be too low for larger Active Directory clusters, and too high for smaller ones.
}

type ADData struct {
	ADConfig
	Formatter               string `json:"formatter,omitempty"`
	Length                  int    `json:"length,omitempty"`
	TLSMaxVersion           string `json:"tls_max_version,omitempty"`
	TLSMinVersion           string `json:"tls_min_version,omitempty"`
	UsePre111GroupBehaviour bool   `json:"use_pre111_group_cn_behavior,omitempty"`
}

type ADRole struct {
	ServiceNameAccount string `json:"service_account_name"` // The name of a pre-existing service account in Active Directory that maps to this role.
	TTL                string `json:"ttl,omitempty"`        // The password time-to-live in seconds. Defaults to the configuration ttl if not provided.
}

type ADLibrary struct {
	ServiceAccountNames       []string `json:"service_account_names,omitempty"`
	TTL                       string   `json:"ttl,omitempty"`
	MaxTTL                    string   `json:"max_ttl,omitempty"`
	DisableCheckInEnforcement bool     `json:"disable_check_in_enforcement,omitempty"`
}

func (cfg ADConfig) Validate() error {
	if cfg.BindDB == "" {
		return ErrMissingBindDB
	}
	if cfg.BindPass == "" {
		return ErrMissingBindPass
	}
	return nil
}

func (adr ADRole) Validate() error {
	if adr.ServiceNameAccount == "" {
		return ErrMissingServiceName
	}
	return nil
}

func (adl ADLibrary) Validate() error {
	if len(adl.ServiceAccountNames) == 0 {
		return ErrMissingServiceName
	}
	return nil
}
