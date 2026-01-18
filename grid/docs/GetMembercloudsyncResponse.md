# GetMembercloudsyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CloudSyncEnabled** | Pointer to **bool** | Option to enable/disable Cloud Sync. | [optional] 
**HostName** | Pointer to **string** | Host name of the parent Member | [optional] [readonly] 
**Result** | Pointer to [**Membercloudsync**](Membercloudsync.md) |  | [optional] 

## Methods

### NewGetMembercloudsyncResponse

`func NewGetMembercloudsyncResponse() *GetMembercloudsyncResponse`

NewGetMembercloudsyncResponse instantiates a new GetMembercloudsyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMembercloudsyncResponseWithDefaults

`func NewGetMembercloudsyncResponseWithDefaults() *GetMembercloudsyncResponse`

NewGetMembercloudsyncResponseWithDefaults instantiates a new GetMembercloudsyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMembercloudsyncResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMembercloudsyncResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMembercloudsyncResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMembercloudsyncResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudSyncEnabled

`func (o *GetMembercloudsyncResponse) GetCloudSyncEnabled() bool`

GetCloudSyncEnabled returns the CloudSyncEnabled field if non-nil, zero value otherwise.

### GetCloudSyncEnabledOk

`func (o *GetMembercloudsyncResponse) GetCloudSyncEnabledOk() (*bool, bool)`

GetCloudSyncEnabledOk returns a tuple with the CloudSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSyncEnabled

`func (o *GetMembercloudsyncResponse) SetCloudSyncEnabled(v bool)`

SetCloudSyncEnabled sets CloudSyncEnabled field to given value.

### HasCloudSyncEnabled

`func (o *GetMembercloudsyncResponse) HasCloudSyncEnabled() bool`

HasCloudSyncEnabled returns a boolean if a field has been set.

### GetHostName

`func (o *GetMembercloudsyncResponse) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetMembercloudsyncResponse) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetMembercloudsyncResponse) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *GetMembercloudsyncResponse) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetResult

`func (o *GetMembercloudsyncResponse) GetResult() Membercloudsync`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMembercloudsyncResponse) GetResultOk() (*Membercloudsync, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMembercloudsyncResponse) SetResult(v Membercloudsync)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMembercloudsyncResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


