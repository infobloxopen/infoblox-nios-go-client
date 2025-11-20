# GetHostnamerewritepolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | True if the policy is the Grid default. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of a hostname rewrite policy object. | [optional] 
**PreDefined** | Pointer to **bool** | Determines whether the policy is a predefined one. | [optional] [readonly] 
**ReplacementCharacter** | Pointer to **string** | The replacement character for symbols in hostnames that do not conform to the hostname policy. | [optional] 
**ValidCharacters** | Pointer to **string** | The set of valid characters represented in string format. | [optional] 
**Result** | Pointer to [**Hostnamerewritepolicy**](Hostnamerewritepolicy.md) |  | [optional] 

## Methods

### NewGetHostnamerewritepolicyResponse

`func NewGetHostnamerewritepolicyResponse() *GetHostnamerewritepolicyResponse`

NewGetHostnamerewritepolicyResponse instantiates a new GetHostnamerewritepolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHostnamerewritepolicyResponseWithDefaults

`func NewGetHostnamerewritepolicyResponseWithDefaults() *GetHostnamerewritepolicyResponse`

NewGetHostnamerewritepolicyResponseWithDefaults instantiates a new GetHostnamerewritepolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetHostnamerewritepolicyResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetHostnamerewritepolicyResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetHostnamerewritepolicyResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetHostnamerewritepolicyResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetHostnamerewritepolicyResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetHostnamerewritepolicyResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetHostnamerewritepolicyResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetHostnamerewritepolicyResponse) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *GetHostnamerewritepolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetHostnamerewritepolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetHostnamerewritepolicyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetHostnamerewritepolicyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreDefined

`func (o *GetHostnamerewritepolicyResponse) GetPreDefined() bool`

GetPreDefined returns the PreDefined field if non-nil, zero value otherwise.

### GetPreDefinedOk

`func (o *GetHostnamerewritepolicyResponse) GetPreDefinedOk() (*bool, bool)`

GetPreDefinedOk returns a tuple with the PreDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDefined

`func (o *GetHostnamerewritepolicyResponse) SetPreDefined(v bool)`

SetPreDefined sets PreDefined field to given value.

### HasPreDefined

`func (o *GetHostnamerewritepolicyResponse) HasPreDefined() bool`

HasPreDefined returns a boolean if a field has been set.

### GetReplacementCharacter

`func (o *GetHostnamerewritepolicyResponse) GetReplacementCharacter() string`

GetReplacementCharacter returns the ReplacementCharacter field if non-nil, zero value otherwise.

### GetReplacementCharacterOk

`func (o *GetHostnamerewritepolicyResponse) GetReplacementCharacterOk() (*string, bool)`

GetReplacementCharacterOk returns a tuple with the ReplacementCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementCharacter

`func (o *GetHostnamerewritepolicyResponse) SetReplacementCharacter(v string)`

SetReplacementCharacter sets ReplacementCharacter field to given value.

### HasReplacementCharacter

`func (o *GetHostnamerewritepolicyResponse) HasReplacementCharacter() bool`

HasReplacementCharacter returns a boolean if a field has been set.

### GetValidCharacters

`func (o *GetHostnamerewritepolicyResponse) GetValidCharacters() string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *GetHostnamerewritepolicyResponse) GetValidCharactersOk() (*string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *GetHostnamerewritepolicyResponse) SetValidCharacters(v string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *GetHostnamerewritepolicyResponse) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### GetResult

`func (o *GetHostnamerewritepolicyResponse) GetResult() Hostnamerewritepolicy`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetHostnamerewritepolicyResponse) GetResultOk() (*Hostnamerewritepolicy, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetHostnamerewritepolicyResponse) SetResult(v Hostnamerewritepolicy)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetHostnamerewritepolicyResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


