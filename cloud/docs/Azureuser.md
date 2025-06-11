# Azureuser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ClientId** | Pointer to **string** | The unique Client ID of this Azure user. Maximum 255 characters. | [optional] 
**ClientSecretKey** | Pointer to **string** | The Client Secret Key for the Client ID of this user. Maximum 255 characters. | [optional] 
**LastUsed** | Pointer to **int64** | The timestamp when this Azure user credentials was last used. | [optional] [readonly] 
**Name** | Pointer to **string** | The Azure user name. Maximum 64 characters. | [optional] 
**Status** | Pointer to **string** | Indicate the validity status of this Azure user. | [optional] [readonly] 
**TenantId** | Pointer to **string** | The Azure Tenant ID of this Azure user. Maximum 64 characters. | [optional] 

## Methods

### NewAzureuser

`func NewAzureuser() *Azureuser`

NewAzureuser instantiates a new Azureuser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureuserWithDefaults

`func NewAzureuserWithDefaults() *Azureuser`

NewAzureuserWithDefaults instantiates a new Azureuser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Azureuser) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Azureuser) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Azureuser) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Azureuser) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetClientId

`func (o *Azureuser) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Azureuser) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Azureuser) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Azureuser) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecretKey

`func (o *Azureuser) GetClientSecretKey() string`

GetClientSecretKey returns the ClientSecretKey field if non-nil, zero value otherwise.

### GetClientSecretKeyOk

`func (o *Azureuser) GetClientSecretKeyOk() (*string, bool)`

GetClientSecretKeyOk returns a tuple with the ClientSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretKey

`func (o *Azureuser) SetClientSecretKey(v string)`

SetClientSecretKey sets ClientSecretKey field to given value.

### HasClientSecretKey

`func (o *Azureuser) HasClientSecretKey() bool`

HasClientSecretKey returns a boolean if a field has been set.

### GetLastUsed

`func (o *Azureuser) GetLastUsed() int64`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *Azureuser) GetLastUsedOk() (*int64, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *Azureuser) SetLastUsed(v int64)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *Azureuser) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetName

`func (o *Azureuser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Azureuser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Azureuser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Azureuser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Azureuser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Azureuser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Azureuser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Azureuser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTenantId

`func (o *Azureuser) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Azureuser) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Azureuser) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Azureuser) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


