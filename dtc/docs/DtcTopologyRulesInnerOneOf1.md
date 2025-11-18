# DtcTopologyRulesInnerOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestType** | Pointer to **string** | The type of the destination for this rule. | [optional] 
**DestinationLink** | Pointer to **string** | The reference to the destination object. | [optional] 
**ReturnType** | Pointer to **string** | The type of the return value for this source. | [optional] 
**Topology** | Pointer to **string** | The topology for this rule. | [optional] [readonly] 
**Valid** | Pointer to **bool** | Indicates whether the rule is valid. | [optional] [readonly] 
**Sources** | Pointer to [**[]DtcTopologyRulesInnerOneOf1SourcesInner**](DtcTopologyRulesInnerOneOf1SourcesInner.md) | Conditions for matching sources. | [optional] 

## Methods

### NewDtcTopologyRulesInnerOneOf1

`func NewDtcTopologyRulesInnerOneOf1() *DtcTopologyRulesInnerOneOf1`

NewDtcTopologyRulesInnerOneOf1 instantiates a new DtcTopologyRulesInnerOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcTopologyRulesInnerOneOf1WithDefaults

`func NewDtcTopologyRulesInnerOneOf1WithDefaults() *DtcTopologyRulesInnerOneOf1`

NewDtcTopologyRulesInnerOneOf1WithDefaults instantiates a new DtcTopologyRulesInnerOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestType

`func (o *DtcTopologyRulesInnerOneOf1) GetDestType() string`

GetDestType returns the DestType field if non-nil, zero value otherwise.

### GetDestTypeOk

`func (o *DtcTopologyRulesInnerOneOf1) GetDestTypeOk() (*string, bool)`

GetDestTypeOk returns a tuple with the DestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestType

`func (o *DtcTopologyRulesInnerOneOf1) SetDestType(v string)`

SetDestType sets DestType field to given value.

### HasDestType

`func (o *DtcTopologyRulesInnerOneOf1) HasDestType() bool`

HasDestType returns a boolean if a field has been set.

### GetDestinationLink

`func (o *DtcTopologyRulesInnerOneOf1) GetDestinationLink() string`

GetDestinationLink returns the DestinationLink field if non-nil, zero value otherwise.

### GetDestinationLinkOk

`func (o *DtcTopologyRulesInnerOneOf1) GetDestinationLinkOk() (*string, bool)`

GetDestinationLinkOk returns a tuple with the DestinationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLink

`func (o *DtcTopologyRulesInnerOneOf1) SetDestinationLink(v string)`

SetDestinationLink sets DestinationLink field to given value.

### HasDestinationLink

`func (o *DtcTopologyRulesInnerOneOf1) HasDestinationLink() bool`

HasDestinationLink returns a boolean if a field has been set.

### GetReturnType

`func (o *DtcTopologyRulesInnerOneOf1) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *DtcTopologyRulesInnerOneOf1) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *DtcTopologyRulesInnerOneOf1) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *DtcTopologyRulesInnerOneOf1) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetTopology

`func (o *DtcTopologyRulesInnerOneOf1) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *DtcTopologyRulesInnerOneOf1) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *DtcTopologyRulesInnerOneOf1) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *DtcTopologyRulesInnerOneOf1) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetValid

`func (o *DtcTopologyRulesInnerOneOf1) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DtcTopologyRulesInnerOneOf1) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DtcTopologyRulesInnerOneOf1) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *DtcTopologyRulesInnerOneOf1) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetSources

`func (o *DtcTopologyRulesInnerOneOf1) GetSources() []DtcTopologyRulesInnerOneOf1SourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *DtcTopologyRulesInnerOneOf1) GetSourcesOk() (*[]DtcTopologyRulesInnerOneOf1SourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *DtcTopologyRulesInnerOneOf1) SetSources(v []DtcTopologyRulesInnerOneOf1SourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *DtcTopologyRulesInnerOneOf1) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


