// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeGameSessionsInput
type DescribeGameSessionsInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for an alias associated with the fleet to retrieve all
	// game sessions for.
	AliasId *string `type:"string"`

	// Unique identifier for a fleet to retrieve all game sessions for.
	FleetId *string `type:"string"`

	// Unique identifier for the game session to retrieve. You can use either a
	// GameSessionId or GameSessionArn value.
	GameSessionId *string `min:"1" type:"string"`

	// Maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int64 `min:"1" type:"integer"`

	// Token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this action. To start
	// at the beginning of the result set, do not specify a value.
	NextToken *string `min:"1" type:"string"`

	// Game session status to filter results on. Possible game session statuses
	// include ACTIVE, TERMINATED, ACTIVATING, and TERMINATING (the last two are
	// transitory).
	StatusFilter *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeGameSessionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeGameSessionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeGameSessionsInput"}
	if s.GameSessionId != nil && len(*s.GameSessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionId", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.StatusFilter != nil && len(*s.StatusFilter) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StatusFilter", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeGameSessionsOutput
type DescribeGameSessionsOutput struct {
	_ struct{} `type:"structure"`

	// Collection of objects containing game session properties for each session
	// matching the request.
	GameSessions []GameSession `type:"list"`

	// Token that indicates where to resume retrieving results on the next call
	// to this action. If no token is returned, these results represent the end
	// of the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeGameSessionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeGameSessions = "DescribeGameSessions"

// DescribeGameSessionsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves a set of one or more game sessions. Request a specific game session
// or request all game sessions on a fleet. Alternatively, use SearchGameSessions
// to request a set of active game sessions that are filtered by certain criteria.
// To retrieve protection policy settings for game sessions, use DescribeGameSessionDetails.
//
// To get game sessions, specify one of the following: game session ID, fleet
// ID, or alias ID. You can filter this request by game session status. Use
// the pagination parameters to retrieve results as a set of sequential pages.
// If successful, a GameSession object is returned for each game session matching
// the request.
//
// Available in Amazon GameLift Local.
//
//    * CreateGameSession
//
//    * DescribeGameSessions
//
//    * DescribeGameSessionDetails
//
//    * SearchGameSessions
//
//    * UpdateGameSession
//
//    * GetGameSessionLogUrl
//
//    * Game session placements StartGameSessionPlacement DescribeGameSessionPlacement
//    StopGameSessionPlacement
//
//    // Example sending a request using DescribeGameSessionsRequest.
//    req := client.DescribeGameSessionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeGameSessions
func (c *Client) DescribeGameSessionsRequest(input *DescribeGameSessionsInput) DescribeGameSessionsRequest {
	op := &aws.Operation{
		Name:       opDescribeGameSessions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeGameSessionsInput{}
	}

	req := c.newRequest(op, input, &DescribeGameSessionsOutput{})
	return DescribeGameSessionsRequest{Request: req, Input: input, Copy: c.DescribeGameSessionsRequest}
}

// DescribeGameSessionsRequest is the request type for the
// DescribeGameSessions API operation.
type DescribeGameSessionsRequest struct {
	*aws.Request
	Input *DescribeGameSessionsInput
	Copy  func(*DescribeGameSessionsInput) DescribeGameSessionsRequest
}

// Send marshals and sends the DescribeGameSessions API request.
func (r DescribeGameSessionsRequest) Send(ctx context.Context) (*DescribeGameSessionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGameSessionsResponse{
		DescribeGameSessionsOutput: r.Request.Data.(*DescribeGameSessionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeGameSessionsResponse is the response type for the
// DescribeGameSessions API operation.
type DescribeGameSessionsResponse struct {
	*DescribeGameSessionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGameSessions request.
func (r *DescribeGameSessionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}