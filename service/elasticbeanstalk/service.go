// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// ElasticBeanstalk provides the API operation methods for making requests to
// AWS Elastic Beanstalk. See this package's package overview docs
// for details on the service.
//
// ElasticBeanstalk methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ElasticBeanstalk struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*aws.Client)

// Used for custom request initialization logic
var initRequest func(*aws.Request)

// Service information constants
const (
	ServiceName = "elasticbeanstalk" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName        // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ElasticBeanstalk client with a config.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ElasticBeanstalk client from just a config.
//     svc := elasticbeanstalk.New(myConfig)
//
//     // Create a ElasticBeanstalk client with additional configuration
//     svc := elasticbeanstalk.New(myConfig, aws.NewConfig().WithRegion("us-west-2"))
func New(config aws.Config) *ElasticBeanstalk {
	var signingName string
	signingRegion := config.Region

	svc := &ElasticBeanstalk{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2010-12-01",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ElasticBeanstalk operation and runs any
// custom request initialization.
func (c *ElasticBeanstalk) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
