# GetSmartfolderGlobalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The global Smart Folder descriptive comment. | [optional] 
**GroupBys** | Pointer to [**[]SmartfolderGlobalGroupBys**](SmartfolderGlobalGroupBys.md) | Global Smart Folder grouping rules. | [optional] 
**Name** | Pointer to **string** | The global Smart Folder name. | [optional] 
**QueryItems** | Pointer to [**[]SmartfolderGlobalQueryItems**](SmartfolderGlobalQueryItems.md) | The global Smart Folder filter queries. | [optional] 
**Result** | Pointer to [**SmartfolderGlobal**](SmartfolderGlobal.md) |  | [optional] 

## Methods

### NewGetSmartfolderGlobalResponse

`func NewGetSmartfolderGlobalResponse() *GetSmartfolderGlobalResponse`

NewGetSmartfolderGlobalResponse instantiates a new GetSmartfolderGlobalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSmartfolderGlobalResponseWithDefaults

`func NewGetSmartfolderGlobalResponseWithDefaults() *GetSmartfolderGlobalResponse`

NewGetSmartfolderGlobalResponseWithDefaults instantiates a new GetSmartfolderGlobalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSmartfolderGlobalResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSmartfolderGlobalResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSmartfolderGlobalResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSmartfolderGlobalResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetSmartfolderGlobalResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSmartfolderGlobalResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSmartfolderGlobalResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSmartfolderGlobalResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroupBys

`func (o *GetSmartfolderGlobalResponse) GetGroupBys() []SmartfolderGlobalGroupBys`

GetGroupBys returns the GroupBys field if non-nil, zero value otherwise.

### GetGroupBysOk

`func (o *GetSmartfolderGlobalResponse) GetGroupBysOk() (*[]SmartfolderGlobalGroupBys, bool)`

GetGroupBysOk returns a tuple with the GroupBys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBys

`func (o *GetSmartfolderGlobalResponse) SetGroupBys(v []SmartfolderGlobalGroupBys)`

SetGroupBys sets GroupBys field to given value.

### HasGroupBys

`func (o *GetSmartfolderGlobalResponse) HasGroupBys() bool`

HasGroupBys returns a boolean if a field has been set.

### GetName

`func (o *GetSmartfolderGlobalResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSmartfolderGlobalResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSmartfolderGlobalResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSmartfolderGlobalResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueryItems

`func (o *GetSmartfolderGlobalResponse) GetQueryItems() []SmartfolderGlobalQueryItems`

GetQueryItems returns the QueryItems field if non-nil, zero value otherwise.

### GetQueryItemsOk

`func (o *GetSmartfolderGlobalResponse) GetQueryItemsOk() (*[]SmartfolderGlobalQueryItems, bool)`

GetQueryItemsOk returns a tuple with the QueryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryItems

`func (o *GetSmartfolderGlobalResponse) SetQueryItems(v []SmartfolderGlobalQueryItems)`

SetQueryItems sets QueryItems field to given value.

### HasQueryItems

`func (o *GetSmartfolderGlobalResponse) HasQueryItems() bool`

HasQueryItems returns a boolean if a field has been set.

### GetResult

`func (o *GetSmartfolderGlobalResponse) GetResult() SmartfolderGlobal`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSmartfolderGlobalResponse) GetResultOk() (*SmartfolderGlobal, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSmartfolderGlobalResponse) SetResult(v SmartfolderGlobal)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSmartfolderGlobalResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


