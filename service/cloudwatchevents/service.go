// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// CloudWatchEvents provides the API operation methods for making requests to
// Amazon CloudWatch Events. See this package's package overview docs
// for details on the service.
//
// CloudWatchEvents methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type CloudWatchEvents struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*aws.Client)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// Service information constants
const (
	ServiceName = "events"    // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the CloudWatchEvents client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CloudWatchEvents client from just a config.
//     svc := cloudwatchevents.New(myConfig)
//
//     // Create a CloudWatchEvents client with additional configuration
//     svc := cloudwatchevents.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *CloudWatchEvents {
	var signingName string
	signingRegion := config.Region

	svc := &CloudWatchEvents{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2015-10-07",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSEvents",
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

// newRequest creates a new request for a CloudWatchEvents operation and runs any
// custom request initialization.
func (c *CloudWatchEvents) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
