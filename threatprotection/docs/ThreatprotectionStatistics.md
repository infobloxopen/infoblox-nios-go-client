# ThreatprotectionStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Member** | Pointer to **string** | The Grid member name to get threat protection statistics. If nothing is specified then event statistics is returned for the Grid. | [optional] [readonly] 
**StatInfos** | Pointer to [**[]ThreatprotectionStatisticsStatInfos**](ThreatprotectionStatisticsStatInfos.md) | The list of event statistical information for the Grid or particular members. | [optional] [readonly] 

## Methods

### NewThreatprotectionStatistics

`func NewThreatprotectionStatistics() *ThreatprotectionStatistics`

NewThreatprotectionStatistics instantiates a new ThreatprotectionStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionStatisticsWithDefaults

`func NewThreatprotectionStatisticsWithDefaults() *ThreatprotectionStatistics`

NewThreatprotectionStatisticsWithDefaults instantiates a new ThreatprotectionStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ThreatprotectionStatistics) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ThreatprotectionStatistics) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ThreatprotectionStatistics) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ThreatprotectionStatistics) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetMember

`func (o *ThreatprotectionStatistics) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ThreatprotectionStatistics) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ThreatprotectionStatistics) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *ThreatprotectionStatistics) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetStatInfos

`func (o *ThreatprotectionStatistics) GetStatInfos() []ThreatprotectionStatisticsStatInfos`

GetStatInfos returns the StatInfos field if non-nil, zero value otherwise.

### GetStatInfosOk

`func (o *ThreatprotectionStatistics) GetStatInfosOk() (*[]ThreatprotectionStatisticsStatInfos, bool)`

GetStatInfosOk returns a tuple with the StatInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatInfos

`func (o *ThreatprotectionStatistics) SetStatInfos(v []ThreatprotectionStatisticsStatInfos)`

SetStatInfos sets StatInfos field to given value.

### HasStatInfos

`func (o *ThreatprotectionStatistics) HasStatInfos() bool`

HasStatInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


