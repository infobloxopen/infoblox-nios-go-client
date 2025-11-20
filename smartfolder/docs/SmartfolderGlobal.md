# SmartfolderGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The global Smart Folder descriptive comment. | [optional] 
**GroupBys** | Pointer to [**[]SmartfolderGlobalGroupBys**](SmartfolderGlobalGroupBys.md) | Global Smart Folder grouping rules. | [optional] 
**Name** | Pointer to **string** | The global Smart Folder name. | [optional] 
**QueryItems** | Pointer to [**[]SmartfolderGlobalQueryItems**](SmartfolderGlobalQueryItems.md) | The global Smart Folder filter queries. | [optional] 

## Methods

### NewSmartfolderGlobal

`func NewSmartfolderGlobal() *SmartfolderGlobal`

NewSmartfolderGlobal instantiates a new SmartfolderGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartfolderGlobalWithDefaults

`func NewSmartfolderGlobalWithDefaults() *SmartfolderGlobal`

NewSmartfolderGlobalWithDefaults instantiates a new SmartfolderGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *SmartfolderGlobal) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SmartfolderGlobal) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SmartfolderGlobal) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SmartfolderGlobal) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *SmartfolderGlobal) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SmartfolderGlobal) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SmartfolderGlobal) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SmartfolderGlobal) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroupBys

`func (o *SmartfolderGlobal) GetGroupBys() []SmartfolderGlobalGroupBys`

GetGroupBys returns the GroupBys field if non-nil, zero value otherwise.

### GetGroupBysOk

`func (o *SmartfolderGlobal) GetGroupBysOk() (*[]SmartfolderGlobalGroupBys, bool)`

GetGroupBysOk returns a tuple with the GroupBys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBys

`func (o *SmartfolderGlobal) SetGroupBys(v []SmartfolderGlobalGroupBys)`

SetGroupBys sets GroupBys field to given value.

### HasGroupBys

`func (o *SmartfolderGlobal) HasGroupBys() bool`

HasGroupBys returns a boolean if a field has been set.

### GetName

`func (o *SmartfolderGlobal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmartfolderGlobal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmartfolderGlobal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmartfolderGlobal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueryItems

`func (o *SmartfolderGlobal) GetQueryItems() []SmartfolderGlobalQueryItems`

GetQueryItems returns the QueryItems field if non-nil, zero value otherwise.

### GetQueryItemsOk

`func (o *SmartfolderGlobal) GetQueryItemsOk() (*[]SmartfolderGlobalQueryItems, bool)`

GetQueryItemsOk returns a tuple with the QueryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryItems

`func (o *SmartfolderGlobal) SetQueryItems(v []SmartfolderGlobalQueryItems)`

SetQueryItems sets QueryItems field to given value.

### HasQueryItems

`func (o *SmartfolderGlobal) HasQueryItems() bool`

HasQueryItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


