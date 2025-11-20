# Recordnamepolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Determines whether the record name policy is Grid default. | [optional] 
**Name** | Pointer to **string** | The name of the record name policy object. | [optional] 
**PreDefined** | Pointer to **bool** | Determines whether the record name policy is a predefined one. | [optional] [readonly] 
**Regex** | Pointer to **string** | The POSIX regular expression the record names should match in order to comply with the record name policy. | [optional] 

## Methods

### NewRecordnamepolicy

`func NewRecordnamepolicy() *Recordnamepolicy`

NewRecordnamepolicy instantiates a new Recordnamepolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordnamepolicyWithDefaults

`func NewRecordnamepolicyWithDefaults() *Recordnamepolicy`

NewRecordnamepolicyWithDefaults instantiates a new Recordnamepolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Recordnamepolicy) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Recordnamepolicy) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Recordnamepolicy) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Recordnamepolicy) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetIsDefault

`func (o *Recordnamepolicy) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Recordnamepolicy) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Recordnamepolicy) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Recordnamepolicy) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *Recordnamepolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Recordnamepolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Recordnamepolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Recordnamepolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreDefined

`func (o *Recordnamepolicy) GetPreDefined() bool`

GetPreDefined returns the PreDefined field if non-nil, zero value otherwise.

### GetPreDefinedOk

`func (o *Recordnamepolicy) GetPreDefinedOk() (*bool, bool)`

GetPreDefinedOk returns a tuple with the PreDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreDefined

`func (o *Recordnamepolicy) SetPreDefined(v bool)`

SetPreDefined sets PreDefined field to given value.

### HasPreDefined

`func (o *Recordnamepolicy) HasPreDefined() bool`

HasPreDefined returns a boolean if a field has been set.

### GetRegex

`func (o *Recordnamepolicy) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *Recordnamepolicy) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *Recordnamepolicy) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *Recordnamepolicy) HasRegex() bool`

HasRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


