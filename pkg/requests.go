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

type AlibabaConfig struct {
	AccessKey string `json:"access_key,omitempty"`
	SecretKey string `json:"secret_key,omitempty"`
}

type AlibabaRole struct {
	Name           string   `json:"name,omitempty"`            // Specifies the name of the role to generate credentials against. This is part of the request URL.
	RemotePolicies []string `json:"remote_policies,omitempty"` // The names and types of a pre-existing policies to be applied to the generate access token. Example: "name:AliyunOSSReadOnlyAccess,type:System".
	InlinePolicies []string `json:"inline_policies,omitempty"` // The policy document JSON to be generated and attached to the access token.
	RoleArn        string   `json:"role_arn,omitempty"`        // The ARN of a role that will be assumed to obtain STS credentials. See Vault AliCloud documentation regarding trusted actors.
	TTL            uint64   `json:"ttl,omitempty"`             // The duration in seconds after which the issued token should expire. Defaults to 0, in which case the value will fallback to the system/mount defaults.
	MaxTTL         uint64   `json:"max_ttl,omitempty"`         // The maximum allowed lifetime of tokens issued using this role.
}

type AWSRootIAMCreds struct {
	MaxRetries       int64  `json:"max_retries,omitempty"`       // Number of max retries the client should use for recoverable errors. The default (-1) falls back to the AWS SDK's default behavior.
	AccessKey        string `json:"access_key,omitempty"`        // Specifies the AWS access key ID.
	SecretKey        string `json:"secret_key,omitempty"`        // Specifies the AWS secret access key.
	Region           string `json:"region,omitempty"`            // Specifies the AWS region. If not set it will use the AWS_REGION env var, AWS_DEFAULT_REGION env var, or us-east-1 in that order.
	IAMEndpoint      string `json:"iam_endpoint,omitempty"`      // Specifies a custom HTTP IAM endpoint to use.
	STSEndpoint      string `json:"sts_endpoint,omitempty"`      // Specifies a custom HTTP STS endpoint to use.
	UsernameTemplate string `json:"username_template,omitempty"` // Template describing how dynamic usernames are generated.
}

type AWSLease struct {
	Lease    string `json:"lease,omitempty"`     // Specifies the lease value provided as a string duration with time suffix. "h" (hour) is the largest suffix.
	LeaseMax string `json:"lease_max,omitempty"` // Specifies the lease value provided as a string duration with time suffix. "h" (hour) is the largest suffix.
}

type AWSRole struct {
	Name                   string   `json:"name,omitempty"`                     // Specifies the name of the role to create. This is part of the request URL.
	CredentialsType        string   `json:"credential_type,omitempty"`          // Specifies the type of credential to be used when retrieving credentials from the role. Must be one of iam_user, assumed_role, or federation_token.
	RoleArn                []string `json:"role_arns,omitempty"`                // Specifies the ARNs of the AWS roles this Vault role is allowed to assume. Required when credential_type is assumed_role and prohibited otherwise. This is a comma-separated string or JSON array.
	PolicyArn              []string `json:"policy_arns,omitempty"`              // Specifies a list of AWS managed policy ARN. The behavior depends on the credential type. With iam_user, the policies will be attached to IAM users when they are requested. With assumed_role and federation_token, the policy ARNs will act as a filter on what the credentials can do, similar to policy_document. When credential_type is iam_user or federation_token, at least one of policy_arns or policy_document must be specified. This is a comma-separated string or JSON array.
	PolicyDocument         string   `json:"policy_document,omitempty"`          // The IAM policy document for the role. The behavior depends on the credential type. With iam_user, the policy document will be attached to the IAM user generated and augment the permissions the IAM user has. With assumed_role and federation_token, the policy document will act as a filter on what the credentials can do, similar to policy_arns.
	IAMGroups              []string `json:"iam_groups,omitempty"`               // A list of IAM group names. IAM users generated against this vault role will be added to these IAM Groups. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters.
	IAMTags                []string `json:"iam_tags,omitempty"`                 // A list of strings representing a key/value pair to be used as a tag for any iam_user user that is created by this role. Format is a key and value separated by an = (e.g. test_key=value). Note: when using the CLI multiple tags can be specified in the role configuration by adding another iam_tags assignment in the same command.
	DefaultSTSTTL          string   `json:"default_sts_ttl,omitempty"`          // The default TTL for STS credentials. When a TTL is not specified when STS credentials are requested, and a default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of assumed_role or federation_token.
	MaxSTSTTL              string   `json:"max_sts_ttl,omitempty"`              // The max allowed TTL for STS credentials (credentials TTL are capped to max_sts_ttl). Valid only when credential_type is one of assumed_role or federation_token.
	UserPath               string   `json:"user_path,omitempty"`                // The path for the user name. Valid only when credential_type is iam_user. Default is /
	PermissionsBoundaryArn string   `json:"permissions_boundary_arn,omitempty"` // The ARN of the AWS Permissions Boundary to attach to IAM users created in the role. Valid only when credential_type is iam_user. If not specified, then no permissions boundary policy will be attached.
}

type AWSCreds struct {
	Name            string `json:"name,omitempty"`              // Specifies the name of the role to generate credentials against. This is part of the request URL.
	RoleArn         string `json:"role_arn,omitempty"`          // The ARN of the role to assume if credential_type on the Vault role is assumed_role. Must match one of the allowed role ARNs in the Vault role. Optional if the Vault role only allows a single AWS role ARN; required otherwise.
	RoleSessionName string `json:"role_session_name,omitempty"` // The role session name to attach to the assumed role ARN. role_session_name is limited to 64 characters; if exceeded, the role_session_name in the assumed role ARN will be truncated to 64 characters. If role_session_name is not provided, then it will be generated dynamically by default.
	TTL             string `json:"ttl,omitempty"`               // Specifies the TTL for the use of the STS token. This is specified as a string with a duration suffix. Valid only when credential_type is assumed_role or federation_token. When not specified, the default_sts_ttl set for the role will be used. If that is also not set, then the default value of 3600s will be used. AWS places limits on the maximum TTL allowed. See the AWS documentation on the DurationSeconds parameter for AssumeRole (for assumed_role credential types) and GetFederationToken (for federation_token credential types) for more details.

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

func (ac AlibabaConfig) Validate() error {
	if ac.AccessKey == "" {
		return ErrMissingAccessKey
	}
	if ac.SecretKey == "" {
		return ErrMissingSecretKey
	}
	return nil
}

func (ar AlibabaRole) Validate() error {
	if ar.Name == "" {
		return ErrMissingName
	}
	return nil
}

func (aric AWSRootIAMCreds) Validate() error {
	if aric.AccessKey == "" {
		return ErrMissingAccessKey
	}
	if aric.SecretKey == "" {
		return ErrMissingSecretKey
	}
	return nil
}

func (al AWSLease) Validate() error {
	if al.Lease == "" {
		return ErrMissingLease
	}
	if al.LeaseMax == "" {
		return ErrMissingLeaseMax
	}
	return nil
}

func (ar AWSRole) Validate() error {
	if ar.Name == "" {
		return ErrMissingName
	}
	if ar.CredentialsType == "" {
		return ErrMissingCredentialsType
	}
	return nil
}

func (ac AWSCreds) Validate() error {
	if ac.Name == "" {
		return ErrMissingName
	}
	return nil
}
