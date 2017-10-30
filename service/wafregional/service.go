// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// WAFRegional provides the API operation methods for making requests to
// AWS WAF Regional. See this package's package overview docs
// for details on the service.
//
// WAFRegional methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type WAFRegional struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*aws.Client)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// Service information constants
const (
	ServiceName = "waf-regional" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName    // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the WAFRegional client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a WAFRegional client from just a config.
//     svc := wafregional.New(myConfig)
//
//     // Create a WAFRegional client with additional configuration
//     svc := wafregional.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *WAFRegional {
	var signingName string
	signingRegion := config.Region

	svc := &WAFRegional{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2016-11-28",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSWAF_Regional_20161128",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a WAFRegional operation and runs any
// custom request initialization.
func (c *WAFRegional) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
