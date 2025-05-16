# GridUpdatesDownloadMemberConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to **string** | The name of the updates download member. | [optional] 
**Interface** | Pointer to **string** | The source interface for updates download requests. | [optional] 
**IsOnline** | Pointer to **bool** | Determines if the updates download member is online or not. | [optional] [readonly] 

## Methods

### NewGridUpdatesDownloadMemberConfig

`func NewGridUpdatesDownloadMemberConfig() *GridUpdatesDownloadMemberConfig`

NewGridUpdatesDownloadMemberConfig instantiates a new GridUpdatesDownloadMemberConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridUpdatesDownloadMemberConfigWithDefaults

`func NewGridUpdatesDownloadMemberConfigWithDefaults() *GridUpdatesDownloadMemberConfig`

NewGridUpdatesDownloadMemberConfigWithDefaults instantiates a new GridUpdatesDownloadMemberConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *GridUpdatesDownloadMemberConfig) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GridUpdatesDownloadMemberConfig) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GridUpdatesDownloadMemberConfig) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GridUpdatesDownloadMemberConfig) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInterface

`func (o *GridUpdatesDownloadMemberConfig) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *GridUpdatesDownloadMemberConfig) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *GridUpdatesDownloadMemberConfig) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *GridUpdatesDownloadMemberConfig) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIsOnline

`func (o *GridUpdatesDownloadMemberConfig) GetIsOnline() bool`

GetIsOnline returns the IsOnline field if non-nil, zero value otherwise.

### GetIsOnlineOk

`func (o *GridUpdatesDownloadMemberConfig) GetIsOnlineOk() (*bool, bool)`

GetIsOnlineOk returns a tuple with the IsOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnline

`func (o *GridUpdatesDownloadMemberConfig) SetIsOnline(v bool)`

SetIsOnline sets IsOnline field to given value.

### HasIsOnline

`func (o *GridUpdatesDownloadMemberConfig) HasIsOnline() bool`

HasIsOnline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


