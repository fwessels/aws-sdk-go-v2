// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// MTurk provides the API operation methods for making requests to
// Amazon Mechanical Turk. See this package's package overview docs
// for details on the service.
//
// MTurk methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type MTurk struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*aws.Client)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// Service information constants
const (
	ServiceName = "mturk-requester" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName       // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the MTurk client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a MTurk client from just a config.
//     svc := mturk.New(myConfig)
//
//     // Create a MTurk client with additional configuration
//     svc := mturk.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *MTurk {
	var signingName string
	signingRegion := config.Region

	svc := &MTurk{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-01-17",
				JSONVersion:   "1.1",
				TargetPrefix:  "MTurkRequesterServiceV20170117",
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

// newRequest creates a new request for a MTurk operation and runs any
// custom request initialization.
func (c *MTurk) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
