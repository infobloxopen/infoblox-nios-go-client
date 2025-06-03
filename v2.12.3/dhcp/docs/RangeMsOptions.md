# RangeMsOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Num** | Pointer to **int64** | The code of the DHCP option. | [optional] 
**Value** | Pointer to **string** | Value of the DHCP option. | [optional] 
**Name** | Pointer to **string** | The name of the DHCP option. | [optional] 
**VendorClass** | Pointer to **string** | The name of the vendor class with which this DHCP option is associated. | [optional] 
**UserClass** | Pointer to **string** | The name of the user class with which this DHCP option is associated. | [optional] 
**Type** | Pointer to **string** | The DHCP option type. Valid values are: * \&quot;16-bit signed integer\&quot; * \&quot;16-bit unsigned integer\&quot; * \&quot;32-bit signed integer\&quot; * \&quot;32-bit unsigned integer\&quot; * \&quot;64-bit unsigned integer\&quot; * \&quot;8-bit signed integer\&quot; * \&quot;8-bit unsigned integer (1,2,4,8)\&quot; * \&quot;8-bit unsigned integer\&quot; * \&quot;array of 16-bit integer\&quot; * \&quot;array of 16-bit unsigned integer\&quot; * \&quot;array of 32-bit integer\&quot; * \&quot;array of 32-bit unsigned integer\&quot; * \&quot;array of 64-bit unsigned integer\&quot; * \&quot;array of 8-bit integer\&quot; * \&quot;array of 8-bit unsigned integer\&quot; * \&quot;array of ip-address pair\&quot; * \&quot;array of ip-address\&quot; * \&quot;array of string\&quot; * \&quot;binary\&quot; * \&quot;boolean array of ip-address\&quot; * \&quot;boolean\&quot; * \&quot;boolean-text\&quot; * \&quot;domain-list\&quot; * \&quot;domain-name\&quot; * \&quot;encapsulated\&quot; * \&quot;ip-address\&quot; * \&quot;string\&quot; * \&quot;text\&quot; | [optional] [readonly] 

## Methods

### NewRangeMsOptions

`func NewRangeMsOptions() *RangeMsOptions`

NewRangeMsOptions instantiates a new RangeMsOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeMsOptionsWithDefaults

`func NewRangeMsOptionsWithDefaults() *RangeMsOptions`

NewRangeMsOptionsWithDefaults instantiates a new RangeMsOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNum

`func (o *RangeMsOptions) GetNum() int64`

GetNum returns the Num field if non-nil, zero value otherwise.

### GetNumOk

`func (o *RangeMsOptions) GetNumOk() (*int64, bool)`

GetNumOk returns a tuple with the Num field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNum

`func (o *RangeMsOptions) SetNum(v int64)`

SetNum sets Num field to given value.

### HasNum

`func (o *RangeMsOptions) HasNum() bool`

HasNum returns a boolean if a field has been set.

### GetValue

`func (o *RangeMsOptions) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RangeMsOptions) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RangeMsOptions) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RangeMsOptions) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetName

`func (o *RangeMsOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RangeMsOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RangeMsOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RangeMsOptions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVendorClass

`func (o *RangeMsOptions) GetVendorClass() string`

GetVendorClass returns the VendorClass field if non-nil, zero value otherwise.

### GetVendorClassOk

`func (o *RangeMsOptions) GetVendorClassOk() (*string, bool)`

GetVendorClassOk returns a tuple with the VendorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorClass

`func (o *RangeMsOptions) SetVendorClass(v string)`

SetVendorClass sets VendorClass field to given value.

### HasVendorClass

`func (o *RangeMsOptions) HasVendorClass() bool`

HasVendorClass returns a boolean if a field has been set.

### GetUserClass

`func (o *RangeMsOptions) GetUserClass() string`

GetUserClass returns the UserClass field if non-nil, zero value otherwise.

### GetUserClassOk

`func (o *RangeMsOptions) GetUserClassOk() (*string, bool)`

GetUserClassOk returns a tuple with the UserClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserClass

`func (o *RangeMsOptions) SetUserClass(v string)`

SetUserClass sets UserClass field to given value.

### HasUserClass

`func (o *RangeMsOptions) HasUserClass() bool`

HasUserClass returns a boolean if a field has been set.

### GetType

`func (o *RangeMsOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RangeMsOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RangeMsOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RangeMsOptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


