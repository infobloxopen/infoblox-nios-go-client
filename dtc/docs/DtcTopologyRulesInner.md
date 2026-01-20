# DtcTopologyRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | Reference to the topology rule | [optional] 
**DestType** | Pointer to **string** | The type of the destination for this rule. | [optional] 
**DestinationLink** | Pointer to **string** | The reference to the destination object. | [optional] 
**ReturnType** | Pointer to **string** | The type of the return value for this source. | [optional] 
**Topology** | Pointer to **string** | The topology for this rule. | [optional] [readonly] 
**Valid** | Pointer to **bool** | Indicates whether the rule is valid. | [optional] [readonly] 
**Sources** | Pointer to [**[]DtcTopologyRulesInnerOneOf1SourcesInner**](DtcTopologyRulesInnerOneOf1SourcesInner.md) | Conditions for matching sources. | [optional] 

## Methods

### NewDtcTopologyRulesInner

`func NewDtcTopologyRulesInner() *DtcTopologyRulesInner`

NewDtcTopologyRulesInner instantiates a new DtcTopologyRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcTopologyRulesInnerWithDefaults

`func NewDtcTopologyRulesInnerWithDefaults() *DtcTopologyRulesInner`

NewDtcTopologyRulesInnerWithDefaults instantiates a new DtcTopologyRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcTopologyRulesInner) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcTopologyRulesInner) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcTopologyRulesInner) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcTopologyRulesInner) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDestType

`func (o *DtcTopologyRulesInner) GetDestType() string`

GetDestType returns the DestType field if non-nil, zero value otherwise.

### GetDestTypeOk

`func (o *DtcTopologyRulesInner) GetDestTypeOk() (*string, bool)`

GetDestTypeOk returns a tuple with the DestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestType

`func (o *DtcTopologyRulesInner) SetDestType(v string)`

SetDestType sets DestType field to given value.

### HasDestType

`func (o *DtcTopologyRulesInner) HasDestType() bool`

HasDestType returns a boolean if a field has been set.

### GetDestinationLink

`func (o *DtcTopologyRulesInner) GetDestinationLink() string`

GetDestinationLink returns the DestinationLink field if non-nil, zero value otherwise.

### GetDestinationLinkOk

`func (o *DtcTopologyRulesInner) GetDestinationLinkOk() (*string, bool)`

GetDestinationLinkOk returns a tuple with the DestinationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLink

`func (o *DtcTopologyRulesInner) SetDestinationLink(v string)`

SetDestinationLink sets DestinationLink field to given value.

### HasDestinationLink

`func (o *DtcTopologyRulesInner) HasDestinationLink() bool`

HasDestinationLink returns a boolean if a field has been set.

### GetReturnType

`func (o *DtcTopologyRulesInner) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *DtcTopologyRulesInner) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *DtcTopologyRulesInner) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *DtcTopologyRulesInner) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetTopology

`func (o *DtcTopologyRulesInner) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *DtcTopologyRulesInner) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *DtcTopologyRulesInner) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *DtcTopologyRulesInner) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetValid

`func (o *DtcTopologyRulesInner) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DtcTopologyRulesInner) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DtcTopologyRulesInner) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *DtcTopologyRulesInner) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetSources

`func (o *DtcTopologyRulesInner) GetSources() []DtcTopologyRulesInnerOneOf1SourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *DtcTopologyRulesInner) GetSourcesOk() (*[]DtcTopologyRulesInnerOneOf1SourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *DtcTopologyRulesInner) SetSources(v []DtcTopologyRulesInnerOneOf1SourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *DtcTopologyRulesInner) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


