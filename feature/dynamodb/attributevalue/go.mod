module github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.31.1-0.20210108204630-4822f3195720
	github.com/aws/aws-sdk-go-v2/service/dynamodb v0.31.0
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v0.31.0
	github.com/google/go-cmp v0.5.4
)

replace github.com/aws/aws-sdk-go-v2/service/dynamodb => ../../../service/dynamodb/

replace github.com/aws/aws-sdk-go-v2 => ../../../

replace github.com/aws/aws-sdk-go-v2/service/dynamodbstreams => ../../../service/dynamodbstreams/

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../../service/internal/accept-encoding/
