package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

// BuildAuthToken will return an authorization token used as the password for a DB connection.
//
// * credentials - Credentials containing AccessKeyID, SecretAccessKey and Region for signing
// * dbUser - User account within the database to sign in with
// * instanceId - The database instance identifier
// * expires - Expiration time in seconds (if <= 0, defaults to 900 seconds/15 minutes)
func BuildAuthToken(ctx context.Context, credentials *base.Credentials, dbUser, instanceId string, expires int) (string, error) {
	if credentials == nil || credentials.AccessKeyID == "" ||
		credentials.SecretAccessKey == "" || credentials.Region == "" {
		return "", fmt.Errorf("credentials provider must not be nil")
	}
	if dbUser == "" || instanceId == "" {
		return "", fmt.Errorf("dbUser or instanceId must not be empty")
	}
	credentials.Service = "rds_mysql"
	endpoint := volcengineutil.NewEndpoint().GetEndpoint()
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
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
	url := volc.SignUrl(req, credentials)
	return url, nil
}
