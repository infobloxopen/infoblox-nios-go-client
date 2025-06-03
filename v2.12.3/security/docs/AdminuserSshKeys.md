# AdminuserSshKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | Pointer to **string** | Unique identifier for the key | [optional] 
**KeyType** | Pointer to **string** | ssh_key_types | [optional] 
**KeyValue** | Pointer to **string** | ssh key text | [optional] 

## Methods

### NewAdminuserSshKeys

`func NewAdminuserSshKeys() *AdminuserSshKeys`

NewAdminuserSshKeys instantiates a new AdminuserSshKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminuserSshKeysWithDefaults

`func NewAdminuserSshKeysWithDefaults() *AdminuserSshKeys`

NewAdminuserSshKeysWithDefaults instantiates a new AdminuserSshKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *AdminuserSshKeys) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *AdminuserSshKeys) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *AdminuserSshKeys) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *AdminuserSshKeys) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### GetKeyType

`func (o *AdminuserSshKeys) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *AdminuserSshKeys) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *AdminuserSshKeys) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *AdminuserSshKeys) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetKeyValue

`func (o *AdminuserSshKeys) GetKeyValue() string`

GetKeyValue returns the KeyValue field if non-nil, zero value otherwise.

### GetKeyValueOk

`func (o *AdminuserSshKeys) GetKeyValueOk() (*string, bool)`

GetKeyValueOk returns a tuple with the KeyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyValue

`func (o *AdminuserSshKeys) SetKeyValue(v string)`

SetKeyValue sets KeyValue field to given value.

### HasKeyValue

`func (o *AdminuserSshKeys) HasKeyValue() bool`

HasKeyValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


