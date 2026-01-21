# GetAzureuserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**ClientId** | Pointer to **string** | The unique Client ID of this Azure user. Maximum 255 characters. | [optional] 
**ClientSecretKey** | Pointer to **string** | The Client Secret Key for the Client ID of this user. Maximum 255 characters. | [optional] 
**LastUsed** | Pointer to **int64** | The timestamp when this Azure user credentials was last used. | [optional] [readonly] 
**Name** | Pointer to **string** | The Azure user name. Maximum 64 characters. | [optional] 
**Status** | Pointer to **string** | Indicate the validity status of this Azure user. | [optional] [readonly] 
**TenantId** | Pointer to **string** | The Azure Tenant ID of this Azure user. Maximum 64 characters. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Azureuser**](Azureuser.md) |  | [optional] 

## Methods

### NewGetAzureuserResponse

`func NewGetAzureuserResponse() *GetAzureuserResponse`

NewGetAzureuserResponse instantiates a new GetAzureuserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAzureuserResponseWithDefaults

`func NewGetAzureuserResponseWithDefaults() *GetAzureuserResponse`

NewGetAzureuserResponseWithDefaults instantiates a new GetAzureuserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAzureuserResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAzureuserResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAzureuserResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAzureuserResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetClientId

`func (o *GetAzureuserResponse) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetAzureuserResponse) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetAzureuserResponse) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetAzureuserResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecretKey

`func (o *GetAzureuserResponse) GetClientSecretKey() string`

GetClientSecretKey returns the ClientSecretKey field if non-nil, zero value otherwise.

### GetClientSecretKeyOk

`func (o *GetAzureuserResponse) GetClientSecretKeyOk() (*string, bool)`

GetClientSecretKeyOk returns a tuple with the ClientSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretKey

`func (o *GetAzureuserResponse) SetClientSecretKey(v string)`

SetClientSecretKey sets ClientSecretKey field to given value.

### HasClientSecretKey

`func (o *GetAzureuserResponse) HasClientSecretKey() bool`

HasClientSecretKey returns a boolean if a field has been set.

### GetLastUsed

`func (o *GetAzureuserResponse) GetLastUsed() int64`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *GetAzureuserResponse) GetLastUsedOk() (*int64, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *GetAzureuserResponse) SetLastUsed(v int64)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *GetAzureuserResponse) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetName

`func (o *GetAzureuserResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAzureuserResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAzureuserResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAzureuserResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *GetAzureuserResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAzureuserResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAzureuserResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAzureuserResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *GetAzureuserResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *GetAzureuserResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *GetAzureuserResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *GetAzureuserResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUuid

`func (o *GetAzureuserResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetAzureuserResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetAzureuserResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetAzureuserResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetAzureuserResponse) GetResult() Azureuser`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAzureuserResponse) GetResultOk() (*Azureuser, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAzureuserResponse) SetResult(v Azureuser)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAzureuserResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


