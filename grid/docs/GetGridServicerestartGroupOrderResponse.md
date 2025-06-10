# GetGridServicerestartGroupOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Groups** | Pointer to **[]string** | The ordered list of the Service Restart Group. | [optional] 
**Result** | Pointer to [**GridServicerestartGroupOrder**](GridServicerestartGroupOrder.md) |  | [optional] 

## Methods

### NewGetGridServicerestartGroupOrderResponse

`func NewGetGridServicerestartGroupOrderResponse() *GetGridServicerestartGroupOrderResponse`

NewGetGridServicerestartGroupOrderResponse instantiates a new GetGridServicerestartGroupOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridServicerestartGroupOrderResponseWithDefaults

`func NewGetGridServicerestartGroupOrderResponseWithDefaults() *GetGridServicerestartGroupOrderResponse`

NewGetGridServicerestartGroupOrderResponseWithDefaults instantiates a new GetGridServicerestartGroupOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridServicerestartGroupOrderResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridServicerestartGroupOrderResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridServicerestartGroupOrderResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridServicerestartGroupOrderResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetGroups

`func (o *GetGridServicerestartGroupOrderResponse) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GetGridServicerestartGroupOrderResponse) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GetGridServicerestartGroupOrderResponse) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GetGridServicerestartGroupOrderResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetResult

`func (o *GetGridServicerestartGroupOrderResponse) GetResult() GridServicerestartGroupOrder`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridServicerestartGroupOrderResponse) GetResultOk() (*GridServicerestartGroupOrder, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridServicerestartGroupOrderResponse) SetResult(v GridServicerestartGroupOrder)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridServicerestartGroupOrderResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


