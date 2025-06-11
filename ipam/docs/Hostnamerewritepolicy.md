# Hostnamerewritepolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**IsDefault** | Pointer to **bool** | True if the policy is the Grid default. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of a hostname rewrite policy object. | [optional] 
**PreDefined** | Pointer to **bool** | Determines whether the policy is a predefined one. | [optional] [readonly] 
**ReplacementCharacter** | Pointer to **string** | The replacement character for symbols in hostnames that do not conform to the hostname policy. | [optional] 
**ValidCharacters** | Pointer to **string** | The set of valid characters represented in string format. | [optional] 

## Methods

### NewHostnamerewritepolicy

`func NewHostnamerewritepolicy() *Hostnamerewritepolicy`

NewHostnamerewritepolicy instantiates a new Hostnamerewritepolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostnamerewritepolicyWithDefaults

`func NewHostnamerewritepolicyWithDefaults() *Hostnamerewritepolicy`

NewHostnamerewritepolicyWithDefaults instantiates a new Hostnamerewritepolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Hostnamerewritepolicy) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Hostnamerewritepolicy) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Hostnamerewritepolicy) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Hostnamerewritepolicy) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIsDefault

`func (o *Hostnamerewritepolicy) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Hostnamerewritepolicy) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Hostnamerewritepolicy) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Hostnamerewritepolicy) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *Hostnamerewritepolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Hostnamerewritepolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Hostnamerewritepolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Hostnamerewritepolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreDefined

`func (o *Hostnamerewritepolicy) GetPreDefined() bool`

GetPreDefined returns the PreDefined field if non-nil, zero value otherwise.

### GetPreDefinedOk

`func (o *Hostnamerewritepolicy) GetPreDefinedOk() (*bool, bool)`

GetPreDefinedOk returns a tuple with the PreDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDefined

`func (o *Hostnamerewritepolicy) SetPreDefined(v bool)`

SetPreDefined sets PreDefined field to given value.

### HasPreDefined

`func (o *Hostnamerewritepolicy) HasPreDefined() bool`

HasPreDefined returns a boolean if a field has been set.

### GetReplacementCharacter

`func (o *Hostnamerewritepolicy) GetReplacementCharacter() string`

GetReplacementCharacter returns the ReplacementCharacter field if non-nil, zero value otherwise.

### GetReplacementCharacterOk

`func (o *Hostnamerewritepolicy) GetReplacementCharacterOk() (*string, bool)`

GetReplacementCharacterOk returns a tuple with the ReplacementCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementCharacter

`func (o *Hostnamerewritepolicy) SetReplacementCharacter(v string)`

SetReplacementCharacter sets ReplacementCharacter field to given value.

### HasReplacementCharacter

`func (o *Hostnamerewritepolicy) HasReplacementCharacter() bool`

HasReplacementCharacter returns a boolean if a field has been set.

### GetValidCharacters

`func (o *Hostnamerewritepolicy) GetValidCharacters() string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *Hostnamerewritepolicy) GetValidCharactersOk() (*string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *Hostnamerewritepolicy) SetValidCharacters(v string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *Hostnamerewritepolicy) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


