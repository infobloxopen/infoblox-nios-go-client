# GetSmartfolderPersonalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The personal Smart Folder descriptive comment. | [optional] 
**GroupBys** | Pointer to [**[]SmartfolderPersonalGroupBys**](SmartfolderPersonalGroupBys.md) | The personal Smart Folder groupping rules. | [optional] 
**IsShortcut** | Pointer to **bool** | Determines whether the personal Smart Folder is a shortcut. | [optional] [readonly] 
**Name** | Pointer to **string** | The personal Smart Folder name. | [optional] 
**QueryItems** | Pointer to [**[]SmartfolderPersonalQueryItems**](SmartfolderPersonalQueryItems.md) | The personal Smart Folder filter queries. | [optional] 
**Result** | Pointer to [**SmartfolderPersonal**](SmartfolderPersonal.md) |  | [optional] 

## Methods

### NewGetSmartfolderPersonalResponse

`func NewGetSmartfolderPersonalResponse() *GetSmartfolderPersonalResponse`

NewGetSmartfolderPersonalResponse instantiates a new GetSmartfolderPersonalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSmartfolderPersonalResponseWithDefaults

`func NewGetSmartfolderPersonalResponseWithDefaults() *GetSmartfolderPersonalResponse`

NewGetSmartfolderPersonalResponseWithDefaults instantiates a new GetSmartfolderPersonalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSmartfolderPersonalResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSmartfolderPersonalResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSmartfolderPersonalResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSmartfolderPersonalResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetSmartfolderPersonalResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSmartfolderPersonalResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSmartfolderPersonalResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSmartfolderPersonalResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroupBys

`func (o *GetSmartfolderPersonalResponse) GetGroupBys() []SmartfolderPersonalGroupBys`

GetGroupBys returns the GroupBys field if non-nil, zero value otherwise.

### GetGroupBysOk

`func (o *GetSmartfolderPersonalResponse) GetGroupBysOk() (*[]SmartfolderPersonalGroupBys, bool)`

GetGroupBysOk returns a tuple with the GroupBys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBys

`func (o *GetSmartfolderPersonalResponse) SetGroupBys(v []SmartfolderPersonalGroupBys)`

SetGroupBys sets GroupBys field to given value.

### HasGroupBys

`func (o *GetSmartfolderPersonalResponse) HasGroupBys() bool`

HasGroupBys returns a boolean if a field has been set.

### GetIsShortcut

`func (o *GetSmartfolderPersonalResponse) GetIsShortcut() bool`

GetIsShortcut returns the IsShortcut field if non-nil, zero value otherwise.

### GetIsShortcutOk

`func (o *GetSmartfolderPersonalResponse) GetIsShortcutOk() (*bool, bool)`

GetIsShortcutOk returns a tuple with the IsShortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShortcut

`func (o *GetSmartfolderPersonalResponse) SetIsShortcut(v bool)`

SetIsShortcut sets IsShortcut field to given value.

### HasIsShortcut

`func (o *GetSmartfolderPersonalResponse) HasIsShortcut() bool`

HasIsShortcut returns a boolean if a field has been set.

### GetName

`func (o *GetSmartfolderPersonalResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSmartfolderPersonalResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSmartfolderPersonalResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSmartfolderPersonalResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueryItems

`func (o *GetSmartfolderPersonalResponse) GetQueryItems() []SmartfolderPersonalQueryItems`

GetQueryItems returns the QueryItems field if non-nil, zero value otherwise.

### GetQueryItemsOk

`func (o *GetSmartfolderPersonalResponse) GetQueryItemsOk() (*[]SmartfolderPersonalQueryItems, bool)`

GetQueryItemsOk returns a tuple with the QueryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryItems

`func (o *GetSmartfolderPersonalResponse) SetQueryItems(v []SmartfolderPersonalQueryItems)`

SetQueryItems sets QueryItems field to given value.

### HasQueryItems

`func (o *GetSmartfolderPersonalResponse) HasQueryItems() bool`

HasQueryItems returns a boolean if a field has been set.

### GetResult

`func (o *GetSmartfolderPersonalResponse) GetResult() SmartfolderPersonal`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSmartfolderPersonalResponse) GetResultOk() (*SmartfolderPersonal, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSmartfolderPersonalResponse) SetResult(v SmartfolderPersonal)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSmartfolderPersonalResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


