# DtcTopologyRule

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

## Methods

### NewDtcTopologyRule

`func NewDtcTopologyRule() *DtcTopologyRule`

NewDtcTopologyRule instantiates a new DtcTopologyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcTopologyRuleWithDefaults

`func NewDtcTopologyRuleWithDefaults() *DtcTopologyRule`

NewDtcTopologyRuleWithDefaults instantiates a new DtcTopologyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcTopologyRule) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcTopologyRule) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcTopologyRule) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcTopologyRule) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDestType

`func (o *DtcTopologyRule) GetDestType() string`

GetDestType returns the DestType field if non-nil, zero value otherwise.

### GetDestTypeOk

`func (o *DtcTopologyRule) GetDestTypeOk() (*string, bool)`

GetDestTypeOk returns a tuple with the DestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestType

`func (o *DtcTopologyRule) SetDestType(v string)`

SetDestType sets DestType field to given value.

### HasDestType

`func (o *DtcTopologyRule) HasDestType() bool`

HasDestType returns a boolean if a field has been set.

### GetDestinationLink

`func (o *DtcTopologyRule) GetDestinationLink() DtcTopologyRuleDestinationLink`

GetDestinationLink returns the DestinationLink field if non-nil, zero value otherwise.

### GetDestinationLinkOk

`func (o *DtcTopologyRule) GetDestinationLinkOk() (*DtcTopologyRuleDestinationLink, bool)`

GetDestinationLinkOk returns a tuple with the DestinationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLink

`func (o *DtcTopologyRule) SetDestinationLink(v DtcTopologyRuleDestinationLink)`

SetDestinationLink sets DestinationLink field to given value.

### HasDestinationLink

`func (o *DtcTopologyRule) HasDestinationLink() bool`

HasDestinationLink returns a boolean if a field has been set.

### GetReturnType

`func (o *DtcTopologyRule) GetReturnType() string`

GetReturnType returns the ReturnType field if non-nil, zero value otherwise.

### GetReturnTypeOk

`func (o *DtcTopologyRule) GetReturnTypeOk() (*string, bool)`

GetReturnTypeOk returns a tuple with the ReturnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnType

`func (o *DtcTopologyRule) SetReturnType(v string)`

SetReturnType sets ReturnType field to given value.

### HasReturnType

`func (o *DtcTopologyRule) HasReturnType() bool`

HasReturnType returns a boolean if a field has been set.

### GetSources

`func (o *DtcTopologyRule) GetSources() []DtcTopologyRuleSources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *DtcTopologyRule) GetSourcesOk() (*[]DtcTopologyRuleSources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *DtcTopologyRule) SetSources(v []DtcTopologyRuleSources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *DtcTopologyRule) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTopology

`func (o *DtcTopologyRule) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *DtcTopologyRule) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *DtcTopologyRule) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *DtcTopologyRule) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetValid

`func (o *DtcTopologyRule) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DtcTopologyRule) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DtcTopologyRule) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *DtcTopologyRule) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


