# GetDtcTopologyRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**DestType** | Pointer to **string** | The type of the destination for this DTC Topology rule. | [optional] 
**DestinationLink** | Pointer to [**DtcTopologyRuleDestinationLink**](DtcTopologyRuleDestinationLink.md) |  | [optional] 
**ReturnType** | Pointer to **string** | Type of the DNS response for rule. | [optional] 
**Sources** | Pointer to [**[]DtcTopologyRuleSources**](DtcTopologyRuleSources.md) | The conditions for matching sources. Should be empty to set rule as default destination. | [optional] 
**Topology** | Pointer to **string** | The DTC Topology the rule belongs to. | [optional] [readonly] 
**Valid** | Pointer to **bool** | True if the label in the rule exists in the current Topology DB. Always true for SUBNET rules. Rules with non-existent labels may be configured but will never match. | [optional] [readonly] 
**Result** | Pointer to [**DtcTopologyRule**](DtcTopologyRule.md) |  | [optional] 

## Methods

### NewGetDtcTopologyRuleResponse

`func NewGetDtcTopologyRuleResponse() *GetDtcTopologyRuleResponse`

NewGetDtcTopologyRuleResponse instantiates a new GetDtcTopologyRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcTopologyRuleResponseWithDefaults

`func NewGetDtcTopologyRuleResponseWithDefaults() *GetDtcTopologyRuleResponse`

NewGetDtcTopologyRuleResponseWithDefaults instantiates a new GetDtcTopologyRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcTopologyRuleResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcTopologyRuleResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcTopologyRuleResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcTopologyRuleResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDestType

`func (o *GetDtcTopologyRuleResponse) GetDestType() string`

GetDestType returns the DestType field if non-nil, zero value otherwise.

### GetDestTypeOk

`func (o *GetDtcTopologyRuleResponse) GetDestTypeOk() (*string, bool)`

GetDestTypeOk returns a tuple with the DestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestType

`func (o *GetDtcTopologyRuleResponse) SetDestType(v string)`

SetDestType sets DestType field to given value.

### HasDestType

`func (o *GetDtcTopologyRuleResponse) HasDestType() bool`

HasDestType returns a boolean if a field has been set.

### GetDestinationLink

`func (o *GetDtcTopologyRuleResponse) GetDestinationLink() DtcTopologyRuleDestinationLink`

GetDestinationLink returns the DestinationLink field if non-nil, zero value otherwise.

### GetDestinationLinkOk

`func (o *GetDtcTopologyRuleResponse) GetDestinationLinkOk() (*DtcTopologyRuleDestinationLink, bool)`

GetDestinationLinkOk returns a tuple with the DestinationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLink

`func (o *GetDtcTopologyRuleResponse) SetDestinationLink(v DtcTopologyRuleDestinationLink)`

SetDestinationLink sets DestinationLink field to given value.

### HasDestinationLink

`func (o *GetDtcTopologyRuleResponse) HasDestinationLink() bool`

HasDestinationLink returns a boolean if a field has been set.

### GetReturnType

`func (o *GetDtcTopologyRuleResponse) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *GetDtcTopologyRuleResponse) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *GetDtcTopologyRuleResponse) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *GetDtcTopologyRuleResponse) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetSources

`func (o *GetDtcTopologyRuleResponse) GetSources() []DtcTopologyRuleSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetDtcTopologyRuleResponse) GetSourcesOk() (*[]DtcTopologyRuleSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetDtcTopologyRuleResponse) SetSources(v []DtcTopologyRuleSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetDtcTopologyRuleResponse) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTopology

`func (o *GetDtcTopologyRuleResponse) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *GetDtcTopologyRuleResponse) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *GetDtcTopologyRuleResponse) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *GetDtcTopologyRuleResponse) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetValid

`func (o *GetDtcTopologyRuleResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *GetDtcTopologyRuleResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *GetDtcTopologyRuleResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *GetDtcTopologyRuleResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcTopologyRuleResponse) GetResult() DtcTopologyRule`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcTopologyRuleResponse) GetResultOk() (*DtcTopologyRule, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcTopologyRuleResponse) SetResult(v DtcTopologyRule)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcTopologyRuleResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


