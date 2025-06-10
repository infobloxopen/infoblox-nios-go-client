# SmartfolderPersonal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The personal Smart Folder descriptive comment. | [optional] 
**GroupBys** | Pointer to [**[]SmartfolderPersonalGroupBys**](SmartfolderPersonalGroupBys.md) | The personal Smart Folder groupping rules. | [optional] 
**IsShortcut** | Pointer to **bool** | Determines whether the personal Smart Folder is a shortcut. | [optional] [readonly] 
**Name** | Pointer to **string** | The personal Smart Folder name. | [optional] 
**QueryItems** | Pointer to [**[]SmartfolderPersonalQueryItems**](SmartfolderPersonalQueryItems.md) | The personal Smart Folder filter queries. | [optional] 

## Methods

### NewSmartfolderPersonal

`func NewSmartfolderPersonal() *SmartfolderPersonal`

NewSmartfolderPersonal instantiates a new SmartfolderPersonal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartfolderPersonalWithDefaults

`func NewSmartfolderPersonalWithDefaults() *SmartfolderPersonal`

NewSmartfolderPersonalWithDefaults instantiates a new SmartfolderPersonal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *SmartfolderPersonal) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SmartfolderPersonal) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SmartfolderPersonal) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SmartfolderPersonal) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *SmartfolderPersonal) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SmartfolderPersonal) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SmartfolderPersonal) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SmartfolderPersonal) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroupBys

`func (o *SmartfolderPersonal) GetGroupBys() []SmartfolderPersonalGroupBys`

GetGroupBys returns the GroupBys field if non-nil, zero value otherwise.

### GetGroupBysOk

`func (o *SmartfolderPersonal) GetGroupBysOk() (*[]SmartfolderPersonalGroupBys, bool)`

GetGroupBysOk returns a tuple with the GroupBys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBys

`func (o *SmartfolderPersonal) SetGroupBys(v []SmartfolderPersonalGroupBys)`

SetGroupBys sets GroupBys field to given value.

### HasGroupBys

`func (o *SmartfolderPersonal) HasGroupBys() bool`

HasGroupBys returns a boolean if a field has been set.

### GetIsShortcut

`func (o *SmartfolderPersonal) GetIsShortcut() bool`

GetIsShortcut returns the IsShortcut field if non-nil, zero value otherwise.

### GetIsShortcutOk

`func (o *SmartfolderPersonal) GetIsShortcutOk() (*bool, bool)`

GetIsShortcutOk returns a tuple with the IsShortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShortcut

`func (o *SmartfolderPersonal) SetIsShortcut(v bool)`

SetIsShortcut sets IsShortcut field to given value.

### HasIsShortcut

`func (o *SmartfolderPersonal) HasIsShortcut() bool`

HasIsShortcut returns a boolean if a field has been set.

### GetName

`func (o *SmartfolderPersonal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SmartfolderPersonal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SmartfolderPersonal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SmartfolderPersonal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueryItems

`func (o *SmartfolderPersonal) GetQueryItems() []SmartfolderPersonalQueryItems`

GetQueryItems returns the QueryItems field if non-nil, zero value otherwise.

### GetQueryItemsOk

`func (o *SmartfolderPersonal) GetQueryItemsOk() (*[]SmartfolderPersonalQueryItems, bool)`

GetQueryItemsOk returns a tuple with the QueryItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryItems

`func (o *SmartfolderPersonal) SetQueryItems(v []SmartfolderPersonalQueryItems)`

SetQueryItems sets QueryItems field to given value.

### HasQueryItems

`func (o *SmartfolderPersonal) HasQueryItems() bool`

HasQueryItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


