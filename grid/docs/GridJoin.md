# GridJoin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GridName** | Pointer to **string** | The name of the Grid. | [optional] 
**Master** | Pointer to **string** | The virtual IP address of the grid master. | [optional] 
**SharedSecret** | Pointer to **string** | The shared secret string of the grid. | [optional] 

## Methods

### NewGridJoin

`func NewGridJoin() *GridJoin`

NewGridJoin instantiates a new GridJoin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridJoinWithDefaults

`func NewGridJoinWithDefaults() *GridJoin`

NewGridJoinWithDefaults instantiates a new GridJoin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGridName

`func (o *GridJoin) GetGridName() string`

GetGridName returns the GridName field if non-nil, zero value otherwise.

### GetGridNameOk

`func (o *GridJoin) GetGridNameOk() (*string, bool)`

GetGridNameOk returns a tuple with the GridName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridName

`func (o *GridJoin) SetGridName(v string)`

SetGridName sets GridName field to given value.

### HasGridName

`func (o *GridJoin) HasGridName() bool`

HasGridName returns a boolean if a field has been set.

### GetMaster

`func (o *GridJoin) GetMaster() string`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *GridJoin) GetMasterOk() (*string, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *GridJoin) SetMaster(v string)`

SetMaster sets Master field to given value.

### HasMaster

`func (o *GridJoin) HasMaster() bool`

HasMaster returns a boolean if a field has been set.

### GetSharedSecret

`func (o *GridJoin) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *GridJoin) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *GridJoin) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *GridJoin) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


