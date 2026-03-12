# DtcTopologyRuleDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationLink** | Pointer to **string** | The reference to the destination DTC pool or DTC server. | [optional] 
**Priority** | Pointer to **int64** | Priority. | [optional] 

## Methods

### NewDtcTopologyRuleDestination

`func NewDtcTopologyRuleDestination() *DtcTopologyRuleDestination`

NewDtcTopologyRuleDestination instantiates a new DtcTopologyRuleDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcTopologyRuleDestinationWithDefaults

`func NewDtcTopologyRuleDestinationWithDefaults() *DtcTopologyRuleDestination`

NewDtcTopologyRuleDestinationWithDefaults instantiates a new DtcTopologyRuleDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationLink

`func (o *DtcTopologyRuleDestination) GetDestinationLink() string`

GetDestinationLink returns the DestinationLink field if non-nil, zero value otherwise.

### GetDestinationLinkOk

`func (o *DtcTopologyRuleDestination) GetDestinationLinkOk() (*string, bool)`

GetDestinationLinkOk returns a tuple with the DestinationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLink

`func (o *DtcTopologyRuleDestination) SetDestinationLink(v string)`

SetDestinationLink sets DestinationLink field to given value.

### HasDestinationLink

`func (o *DtcTopologyRuleDestination) HasDestinationLink() bool`

HasDestinationLink returns a boolean if a field has been set.

### GetPriority

`func (o *DtcTopologyRuleDestination) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DtcTopologyRuleDestination) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DtcTopologyRuleDestination) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DtcTopologyRuleDestination) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


