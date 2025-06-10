# GetThreatprotectionStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Member** | Pointer to **string** | The Grid member name to get threat protection statistics. If nothing is specified then event statistics is returned for the Grid. | [optional] [readonly] 
**StatInfos** | Pointer to [**[]ThreatprotectionStatisticsStatInfos**](ThreatprotectionStatisticsStatInfos.md) | The list of event statistical information for the Grid or particular members. | [optional] [readonly] 
**Result** | Pointer to [**ThreatprotectionStatistics**](ThreatprotectionStatistics.md) |  | [optional] 

## Methods

### NewGetThreatprotectionStatisticsResponse

`func NewGetThreatprotectionStatisticsResponse() *GetThreatprotectionStatisticsResponse`

NewGetThreatprotectionStatisticsResponse instantiates a new GetThreatprotectionStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatprotectionStatisticsResponseWithDefaults

`func NewGetThreatprotectionStatisticsResponseWithDefaults() *GetThreatprotectionStatisticsResponse`

NewGetThreatprotectionStatisticsResponseWithDefaults instantiates a new GetThreatprotectionStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatprotectionStatisticsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatprotectionStatisticsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatprotectionStatisticsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatprotectionStatisticsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetMember

`func (o *GetThreatprotectionStatisticsResponse) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetThreatprotectionStatisticsResponse) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetThreatprotectionStatisticsResponse) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetThreatprotectionStatisticsResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetStatInfos

`func (o *GetThreatprotectionStatisticsResponse) GetStatInfos() []ThreatprotectionStatisticsStatInfos`

GetStatInfos returns the StatInfos field if non-nil, zero value otherwise.

### GetStatInfosOk

`func (o *GetThreatprotectionStatisticsResponse) GetStatInfosOk() (*[]ThreatprotectionStatisticsStatInfos, bool)`

GetStatInfosOk returns a tuple with the StatInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatInfos

`func (o *GetThreatprotectionStatisticsResponse) SetStatInfos(v []ThreatprotectionStatisticsStatInfos)`

SetStatInfos sets StatInfos field to given value.

### HasStatInfos

`func (o *GetThreatprotectionStatisticsResponse) HasStatInfos() bool`

HasStatInfos returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatprotectionStatisticsResponse) GetResult() ThreatprotectionStatistics`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatprotectionStatisticsResponse) GetResultOk() (*ThreatprotectionStatistics, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatprotectionStatisticsResponse) SetResult(v ThreatprotectionStatistics)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatprotectionStatisticsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


