# ParentalcontrolAvp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The comment for the AVP. | [optional] 
**DomainTypes** | Pointer to **[]string** | The list of domains applicable to AVP. | [optional] 
**IsRestricted** | Pointer to **bool** | Determines if AVP is restricted to domains. | [optional] 
**Name** | Pointer to **string** | The name of AVP. | [optional] 
**Type** | Pointer to **int64** | The type of AVP as per RFC 2865/2866. | [optional] 
**UserDefined** | Pointer to **bool** | Determines if AVP was defined by user. | [optional] [readonly] 
**ValueType** | Pointer to **string** | The type of value. | [optional] 
**VendorId** | Pointer to **int64** | The vendor ID as per RFC 2865/2866. | [optional] 
**VendorType** | Pointer to **int64** | The vendor type as per RFC 2865/2866. | [optional] 

## Methods

### NewParentalcontrolAvp

`func NewParentalcontrolAvp() *ParentalcontrolAvp`

NewParentalcontrolAvp instantiates a new ParentalcontrolAvp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentalcontrolAvpWithDefaults

`func NewParentalcontrolAvpWithDefaults() *ParentalcontrolAvp`

NewParentalcontrolAvpWithDefaults instantiates a new ParentalcontrolAvp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ParentalcontrolAvp) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ParentalcontrolAvp) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ParentalcontrolAvp) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ParentalcontrolAvp) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *ParentalcontrolAvp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ParentalcontrolAvp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ParentalcontrolAvp) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ParentalcontrolAvp) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDomainTypes

`func (o *ParentalcontrolAvp) GetDomainTypes() []string`

GetDomainTypes returns the DomainTypes field if non-nil, zero value otherwise.

### GetDomainTypesOk

`func (o *ParentalcontrolAvp) GetDomainTypesOk() (*[]string, bool)`

GetDomainTypesOk returns a tuple with the DomainTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainTypes

`func (o *ParentalcontrolAvp) SetDomainTypes(v []string)`

SetDomainTypes sets DomainTypes field to given value.

### HasDomainTypes

`func (o *ParentalcontrolAvp) HasDomainTypes() bool`

HasDomainTypes returns a boolean if a field has been set.

### GetIsRestricted

`func (o *ParentalcontrolAvp) GetIsRestricted() bool`

GetIsRestricted returns the IsRestricted field if non-nil, zero value otherwise.

### GetIsRestrictedOk

`func (o *ParentalcontrolAvp) GetIsRestrictedOk() (*bool, bool)`

GetIsRestrictedOk returns a tuple with the IsRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestricted

`func (o *ParentalcontrolAvp) SetIsRestricted(v bool)`

SetIsRestricted sets IsRestricted field to given value.

### HasIsRestricted

`func (o *ParentalcontrolAvp) HasIsRestricted() bool`

HasIsRestricted returns a boolean if a field has been set.

### GetName

`func (o *ParentalcontrolAvp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParentalcontrolAvp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParentalcontrolAvp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParentalcontrolAvp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ParentalcontrolAvp) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ParentalcontrolAvp) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ParentalcontrolAvp) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *ParentalcontrolAvp) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserDefined

`func (o *ParentalcontrolAvp) GetUserDefined() bool`

GetUserDefined returns the UserDefined field if non-nil, zero value otherwise.

### GetUserDefinedOk

`func (o *ParentalcontrolAvp) GetUserDefinedOk() (*bool, bool)`

GetUserDefinedOk returns a tuple with the UserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefined

`func (o *ParentalcontrolAvp) SetUserDefined(v bool)`

SetUserDefined sets UserDefined field to given value.

### HasUserDefined

`func (o *ParentalcontrolAvp) HasUserDefined() bool`

HasUserDefined returns a boolean if a field has been set.

### GetValueType

`func (o *ParentalcontrolAvp) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ParentalcontrolAvp) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ParentalcontrolAvp) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *ParentalcontrolAvp) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetVendorId

`func (o *ParentalcontrolAvp) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *ParentalcontrolAvp) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *ParentalcontrolAvp) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *ParentalcontrolAvp) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetVendorType

`func (o *ParentalcontrolAvp) GetVendorType() int64`

GetVendorType returns the VendorType field if non-nil, zero value otherwise.

### GetVendorTypeOk

`func (o *ParentalcontrolAvp) GetVendorTypeOk() (*int64, bool)`

GetVendorTypeOk returns a tuple with the VendorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorType

`func (o *ParentalcontrolAvp) SetVendorType(v int64)`

SetVendorType sets VendorType field to given value.

### HasVendorType

`func (o *ParentalcontrolAvp) HasVendorType() bool`

HasVendorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


