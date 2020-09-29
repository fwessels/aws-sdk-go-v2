// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EventCategory string

// Enum values for EventCategory
const (
	EventCategoryInsight EventCategory = "insight"
)

type InsightType string

// Enum values for InsightType
const (
	InsightTypeApicallrateinsight InsightType = "ApiCallRateInsight"
)

type LookupAttributeKey string

// Enum values for LookupAttributeKey
const (
	LookupAttributeKeyEvent_id      LookupAttributeKey = "EventId"
	LookupAttributeKeyEvent_name    LookupAttributeKey = "EventName"
	LookupAttributeKeyRead_only     LookupAttributeKey = "ReadOnly"
	LookupAttributeKeyUsername      LookupAttributeKey = "Username"
	LookupAttributeKeyResource_type LookupAttributeKey = "ResourceType"
	LookupAttributeKeyResource_name LookupAttributeKey = "ResourceName"
	LookupAttributeKeyEvent_source  LookupAttributeKey = "EventSource"
	LookupAttributeKeyAccess_key_id LookupAttributeKey = "AccessKeyId"
)

type ReadWriteType string

// Enum values for ReadWriteType
const (
	ReadWriteTypeReadonly  ReadWriteType = "ReadOnly"
	ReadWriteTypeWriteonly ReadWriteType = "WriteOnly"
	ReadWriteTypeAll       ReadWriteType = "All"
)