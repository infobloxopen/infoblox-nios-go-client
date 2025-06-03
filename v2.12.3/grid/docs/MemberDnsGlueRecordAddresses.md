# MemberDnsGlueRecordAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachEmptyRecursiveView** | Pointer to **bool** | Determines if empty view with recursion enabled will be written into the conf file. | [optional] 
**GlueRecordAddress** | Pointer to **string** | The address the appliance uses to generate the glue record. | [optional] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**GlueAddressChoice** | Pointer to **string** | The address choice for auto-created glue records for this view. | [optional] 

## Methods

### NewMemberDnsGlueRecordAddresses

`func NewMemberDnsGlueRecordAddresses() *MemberDnsGlueRecordAddresses`

NewMemberDnsGlueRecordAddresses instantiates a new MemberDnsGlueRecordAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsGlueRecordAddressesWithDefaults

`func NewMemberDnsGlueRecordAddressesWithDefaults() *MemberDnsGlueRecordAddresses`

NewMemberDnsGlueRecordAddressesWithDefaults instantiates a new MemberDnsGlueRecordAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachEmptyRecursiveView

`func (o *MemberDnsGlueRecordAddresses) GetAttachEmptyRecursiveView() bool`

GetAttachEmptyRecursiveView returns the AttachEmptyRecursiveView field if non-nil, zero value otherwise.

### GetAttachEmptyRecursiveViewOk

`func (o *MemberDnsGlueRecordAddresses) GetAttachEmptyRecursiveViewOk() (*bool, bool)`

GetAttachEmptyRecursiveViewOk returns a tuple with the AttachEmptyRecursiveView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachEmptyRecursiveView

`func (o *MemberDnsGlueRecordAddresses) SetAttachEmptyRecursiveView(v bool)`

SetAttachEmptyRecursiveView sets AttachEmptyRecursiveView field to given value.

### HasAttachEmptyRecursiveView

`func (o *MemberDnsGlueRecordAddresses) HasAttachEmptyRecursiveView() bool`

HasAttachEmptyRecursiveView returns a boolean if a field has been set.

### GetGlueRecordAddress

`func (o *MemberDnsGlueRecordAddresses) GetGlueRecordAddress() string`

GetGlueRecordAddress returns the GlueRecordAddress field if non-nil, zero value otherwise.

### GetGlueRecordAddressOk

`func (o *MemberDnsGlueRecordAddresses) GetGlueRecordAddressOk() (*string, bool)`

GetGlueRecordAddressOk returns a tuple with the GlueRecordAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlueRecordAddress

`func (o *MemberDnsGlueRecordAddresses) SetGlueRecordAddress(v string)`

SetGlueRecordAddress sets GlueRecordAddress field to given value.

### HasGlueRecordAddress

`func (o *MemberDnsGlueRecordAddresses) HasGlueRecordAddress() bool`

HasGlueRecordAddress returns a boolean if a field has been set.

### GetView

`func (o *MemberDnsGlueRecordAddresses) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *MemberDnsGlueRecordAddresses) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *MemberDnsGlueRecordAddresses) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *MemberDnsGlueRecordAddresses) HasView() bool`

HasView returns a boolean if a field has been set.

### GetGlueAddressChoice

`func (o *MemberDnsGlueRecordAddresses) GetGlueAddressChoice() string`

GetGlueAddressChoice returns the GlueAddressChoice field if non-nil, zero value otherwise.

### GetGlueAddressChoiceOk

`func (o *MemberDnsGlueRecordAddresses) GetGlueAddressChoiceOk() (*string, bool)`

GetGlueAddressChoiceOk returns a tuple with the GlueAddressChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlueAddressChoice

`func (o *MemberDnsGlueRecordAddresses) SetGlueAddressChoice(v string)`

SetGlueAddressChoice sets GlueAddressChoice field to given value.

### HasGlueAddressChoice

`func (o *MemberDnsGlueRecordAddresses) HasGlueAddressChoice() bool`

HasGlueAddressChoice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


