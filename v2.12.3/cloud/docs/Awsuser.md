# Awsuser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AccessKeyId** | Pointer to **string** | The unique Access Key ID of this AWS user. Maximum 255 characters. | [optional] 
**AccountId** | Pointer to **string** | The AWS Account ID of this AWS user. Maximum 64 characters. | [optional] 
**GovcloudEnabled** | Pointer to **bool** | Indicates if gov cloud is enabled or disabled. | [optional] 
**LastUsed** | Pointer to **int64** | The timestamp when this AWS user credentials was last used. | [optional] [readonly] 
**Name** | Pointer to **string** | The AWS user name. Maximum 64 characters. | [optional] 
**NiosUserName** | Pointer to **string** | The NIOS user name mapped to this AWS user. Maximum 64 characters. | [optional] 
**SecretAccessKey** | Pointer to **string** | The Secret Access Key for the Access Key ID of this user. Maximum 255 characters. | [optional] 
**Status** | Pointer to **string** | Indicate the validity status of this AWS user. | [optional] [readonly] 

## Methods

### NewAwsuser

`func NewAwsuser() *Awsuser`

NewAwsuser instantiates a new Awsuser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsuserWithDefaults

`func NewAwsuserWithDefaults() *Awsuser`

NewAwsuserWithDefaults instantiates a new Awsuser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Awsuser) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Awsuser) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Awsuser) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Awsuser) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *Awsuser) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *Awsuser) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *Awsuser) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *Awsuser) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetAccountId

`func (o *Awsuser) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Awsuser) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Awsuser) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Awsuser) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetGovcloudEnabled

`func (o *Awsuser) GetGovcloudEnabled() bool`

GetGovcloudEnabled returns the GovcloudEnabled field if non-nil, zero value otherwise.

### GetGovcloudEnabledOk

`func (o *Awsuser) GetGovcloudEnabledOk() (*bool, bool)`

GetGovcloudEnabledOk returns a tuple with the GovcloudEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovcloudEnabled

`func (o *Awsuser) SetGovcloudEnabled(v bool)`

SetGovcloudEnabled sets GovcloudEnabled field to given value.

### HasGovcloudEnabled

`func (o *Awsuser) HasGovcloudEnabled() bool`

HasGovcloudEnabled returns a boolean if a field has been set.

### GetLastUsed

`func (o *Awsuser) GetLastUsed() int64`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *Awsuser) GetLastUsedOk() (*int64, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *Awsuser) SetLastUsed(v int64)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *Awsuser) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetName

`func (o *Awsuser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Awsuser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Awsuser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Awsuser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNiosUserName

`func (o *Awsuser) GetNiosUserName() string`

GetNiosUserName returns the NiosUserName field if non-nil, zero value otherwise.

### GetNiosUserNameOk

`func (o *Awsuser) GetNiosUserNameOk() (*string, bool)`

GetNiosUserNameOk returns a tuple with the NiosUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiosUserName

`func (o *Awsuser) SetNiosUserName(v string)`

SetNiosUserName sets NiosUserName field to given value.

### HasNiosUserName

`func (o *Awsuser) HasNiosUserName() bool`

HasNiosUserName returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *Awsuser) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *Awsuser) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *Awsuser) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *Awsuser) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetStatus

`func (o *Awsuser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Awsuser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Awsuser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Awsuser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


