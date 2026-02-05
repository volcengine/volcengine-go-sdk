package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/endpoints"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const (
	defaultService = "rds_mysql"
)

// BuildAuthToken will return an authorization token used as the password for a DB connection.
//
// * config - volcengine.Config containing Credentials, Region, and DisableSSL settings
// * dbUser - User account within the database to sign in with
// * instanceId - The database instance identifier
// * expires - Expiration time in seconds (if <= 0, defaults to 900 seconds/15 minutes)
func BuildAuthToken(ctx context.Context, config *volcengine.Config, dbUser, instanceId string, expires int) (string, error) {
	if config == nil {
		return "", fmt.Errorf("config must not be nil")
	}
	if config.Credentials == nil {
		return "", fmt.Errorf("credentials must not be nil")
	}
	if config.Region == nil || *config.Region == "" {
		return "", fmt.Errorf("region must not be empty")
	}
	if dbUser == "" || instanceId == "" {
		return "", fmt.Errorf("dbUser or instanceId must not be empty")
	}

	region := *config.Region

	// Get base credentials for signing
	baseCreds, err := config.Credentials.GetBase(region, defaultService)
	if err != nil {
		return "", fmt.Errorf("unable to get credentials: %w", err)
	}

	if baseCreds.AccessKeyID == "" || baseCreds.SecretAccessKey == "" {
		return "", fmt.Errorf("accessKeyID or secretAccessKey must not be empty")
	}

	// Build regional endpoint
	host := volcengineutil.NewEndpoint().GetRegionalEndpoint(defaultService, region)

	// Add scheme based on DisableSSL config
	disableSSL := false
	if config.DisableSSL != nil {
		disableSSL = *config.DisableSSL
	}
	url := endpoints.AddScheme(host, disableSSL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("unable to build request: %w", err)
	}

	q := req.URL.Query()
	q.Set("Action", "ConnectDatabase")
	q.Set("Version", "2022-01-01")
	if expires > 0 {
		q.Set("X-Expires", strconv.Itoa(expires))
	} else {
		// Expire Time: 15 minute
		q.Set("X-Expires", "900")
	}
	q.Set("DBUser", dbUser)
	q.Set("InstanceId", instanceId)
	req.URL.RawQuery = q.Encode()

	signedUrl := volc.SignUrl(req, &baseCreds)
	return signedUrl, nil
}
