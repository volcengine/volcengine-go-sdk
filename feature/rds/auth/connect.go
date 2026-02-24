package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/client/metadata"
	"github.com/volcengine/volcengine-go-sdk/volcengine/endpoints"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/signer/volc"
)

const (
	defaultService    = "rds_mysql"
	defaultAPIVersion = "2022-01-01"
	defaultAPI        = "ConnectDatabase"
	defaultExpires    = 900
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

	// Use StandardEndpointResolver to get regional endpoint
	config.EndpointResolver = endpoints.NewStandardEndpointResolver()

	// Create session and resolve endpoint via standard path
	sess, err := session.NewSession(config)
	if err != nil {
		return "", fmt.Errorf("unable to create session: %w", err)
	}
	clientCfg := sess.ClientConfig(defaultService)

	// Set up handlers - only need Sign
	var handlers request.Handlers
	handlers.Sign.PushBackNamed(volc.SignRequestHandler)

	// Create request
	op := &request.Operation{
		Name:       defaultAPI,
		HTTPMethod: http.MethodGet,
		HTTPPath:   "",
	}

	req := request.New(*config, metadata.ClientInfo{
		ServiceName:   defaultService,
		ServiceID:     defaultService,
		SigningName:   defaultService,
		SigningRegion: *config.Region,
		Endpoint:      clientCfg.Endpoint,
		APIVersion:    defaultAPIVersion,
	}, handlers, nil, op, nil, nil)

	req.SetContext(ctx)

	// Set query params
	q := req.HTTPRequest.URL.Query()
	q.Set("Action", defaultAPI)
	q.Set("Version", defaultAPIVersion)
	q.Set("DBUser", dbUser)
	q.Set("InstanceId", instanceId)
	if expires > 0 {
		q.Set("X-Expires", strconv.Itoa(expires))
	} else {
		expires = defaultExpires
		q.Set("X-Expires", strconv.Itoa(defaultExpires))
	}
	req.HTTPRequest.URL.RawQuery = q.Encode()

	// Presign the request
	expireDuration := time.Duration(expires) * time.Second
	signedUrl, _, err := req.PresignRequest(expireDuration)
	if err != nil {
		return "", fmt.Errorf("unable to presign request: %w", err)
	}

	return signedUrl, nil
}
