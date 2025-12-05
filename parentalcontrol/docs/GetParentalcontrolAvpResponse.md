# GetParentalcontrolAvpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for the AVP. | [optional] 
**DomainTypes** | Pointer to **[]string** | The list of domains applicable to AVP. | [optional] 
**IsRestricted** | Pointer to **bool** | Determines if AVP is restricted to domains. | [optional] 
**Name** | Pointer to **string** | The name of AVP. | [optional] 
**Type** | Pointer to **int64** | The type of AVP as per RFC 2865/2866. | [optional] 
**UserDefined** | Pointer to **bool** | Determines if AVP was defined by user. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**ValueType** | Pointer to **string** | The type of value. | [optional] 
**VendorId** | Pointer to **int64** | The vendor ID as per RFC 2865/2866. | [optional] 
**VendorType** | Pointer to **int64** | The vendor type as per RFC 2865/2866. | [optional] 
**Result** | Pointer to [**ParentalcontrolAvp**](ParentalcontrolAvp.md) |  | [optional] 

## Methods

### NewGetParentalcontrolAvpResponse

`func NewGetParentalcontrolAvpResponse() *GetParentalcontrolAvpResponse`

NewGetParentalcontrolAvpResponse instantiates a new GetParentalcontrolAvpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetParentalcontrolAvpResponseWithDefaults

`func NewGetParentalcontrolAvpResponseWithDefaults() *GetParentalcontrolAvpResponse`

NewGetParentalcontrolAvpResponseWithDefaults instantiates a new GetParentalcontrolAvpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetParentalcontrolAvpResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetParentalcontrolAvpResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetParentalcontrolAvpResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetParentalcontrolAvpResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetParentalcontrolAvpResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetParentalcontrolAvpResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetParentalcontrolAvpResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetParentalcontrolAvpResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDomainTypes

`func (o *GetParentalcontrolAvpResponse) GetDomainTypes() []string`

GetDomainTypes returns the DomainTypes field if non-nil, zero value otherwise.

### GetDomainTypesOk

`func (o *GetParentalcontrolAvpResponse) GetDomainTypesOk() (*[]string, bool)`

GetDomainTypesOk returns a tuple with the DomainTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainTypes

`func (o *GetParentalcontrolAvpResponse) SetDomainTypes(v []string)`

SetDomainTypes sets DomainTypes field to given value.

### HasDomainTypes

`func (o *GetParentalcontrolAvpResponse) HasDomainTypes() bool`

HasDomainTypes returns a boolean if a field has been set.

### GetIsRestricted

`func (o *GetParentalcontrolAvpResponse) GetIsRestricted() bool`

GetIsRestricted returns the IsRestricted field if non-nil, zero value otherwise.

### GetIsRestrictedOk

`func (o *GetParentalcontrolAvpResponse) GetIsRestrictedOk() (*bool, bool)`

GetIsRestrictedOk returns a tuple with the IsRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRestricted

`func (o *GetParentalcontrolAvpResponse) SetIsRestricted(v bool)`

SetIsRestricted sets IsRestricted field to given value.

### HasIsRestricted

`func (o *GetParentalcontrolAvpResponse) HasIsRestricted() bool`

HasIsRestricted returns a boolean if a field has been set.

### GetName

`func (o *GetParentalcontrolAvpResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetParentalcontrolAvpResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetParentalcontrolAvpResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetParentalcontrolAvpResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetParentalcontrolAvpResponse) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetParentalcontrolAvpResponse) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetParentalcontrolAvpResponse) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *GetParentalcontrolAvpResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserDefined

`func (o *GetParentalcontrolAvpResponse) GetUserDefined() bool`

GetUserDefined returns the UserDefined field if non-nil, zero value otherwise.

### GetUserDefinedOk

`func (o *GetParentalcontrolAvpResponse) GetUserDefinedOk() (*bool, bool)`

GetUserDefinedOk returns a tuple with the UserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefined

`func (o *GetParentalcontrolAvpResponse) SetUserDefined(v bool)`

SetUserDefined sets UserDefined field to given value.

### HasUserDefined

`func (o *GetParentalcontrolAvpResponse) HasUserDefined() bool`

HasUserDefined returns a boolean if a field has been set.

### GetUuid

`func (o *GetParentalcontrolAvpResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetParentalcontrolAvpResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetParentalcontrolAvpResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetParentalcontrolAvpResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetValueType

`func (o *GetParentalcontrolAvpResponse) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *GetParentalcontrolAvpResponse) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *GetParentalcontrolAvpResponse) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *GetParentalcontrolAvpResponse) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetVendorId

`func (o *GetParentalcontrolAvpResponse) GetVendorId() int64`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *GetParentalcontrolAvpResponse) GetVendorIdOk() (*int64, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *GetParentalcontrolAvpResponse) SetVendorId(v int64)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *GetParentalcontrolAvpResponse) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetVendorType

`func (o *GetParentalcontrolAvpResponse) GetVendorType() int64`

GetVendorType returns the VendorType field if non-nil, zero value otherwise.

### GetVendorTypeOk

`func (o *GetParentalcontrolAvpResponse) GetVendorTypeOk() (*int64, bool)`

GetVendorTypeOk returns a tuple with the VendorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorType

`func (o *GetParentalcontrolAvpResponse) SetVendorType(v int64)`

SetVendorType sets VendorType field to given value.

### HasVendorType

`func (o *GetParentalcontrolAvpResponse) HasVendorType() bool`

HasVendorType returns a boolean if a field has been set.

### GetResult

`func (o *GetParentalcontrolAvpResponse) GetResult() ParentalcontrolAvp`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetParentalcontrolAvpResponse) GetResultOk() (*ParentalcontrolAvp, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetParentalcontrolAvpResponse) SetResult(v ParentalcontrolAvp)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetParentalcontrolAvpResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


