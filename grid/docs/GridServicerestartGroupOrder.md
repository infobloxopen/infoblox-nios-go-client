# GridServicerestartGroupOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Groups** | Pointer to **[]string** | The ordered list of the Service Restart Group. | [optional] 

## Methods

### NewGridServicerestartGroupOrder

`func NewGridServicerestartGroupOrder() *GridServicerestartGroupOrder`

NewGridServicerestartGroupOrder instantiates a new GridServicerestartGroupOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridServicerestartGroupOrderWithDefaults

`func NewGridServicerestartGroupOrderWithDefaults() *GridServicerestartGroupOrder`

NewGridServicerestartGroupOrderWithDefaults instantiates a new GridServicerestartGroupOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridServicerestartGroupOrder) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridServicerestartGroupOrder) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridServicerestartGroupOrder) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridServicerestartGroupOrder) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetGroups

`func (o *GridServicerestartGroupOrder) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GridServicerestartGroupOrder) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GridServicerestartGroupOrder) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GridServicerestartGroupOrder) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


