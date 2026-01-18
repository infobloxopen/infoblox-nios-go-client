# GetIpv6dhcpoptiondefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Code** | Pointer to **int64** | The code of a DHCP IPv6 option definition object. An option code number is used to identify the DHCP option. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP IPv6 option definition object. | [optional] 
**Space** | Pointer to **string** | The space of a DHCP option definition object. | [optional] 
**Type** | Pointer to **string** | The data type of the Grid DHCP IPv6 option. | [optional] 
**Result** | Pointer to [**Ipv6dhcpoptiondefinition**](Ipv6dhcpoptiondefinition.md) |  | [optional] 

## Methods

### NewGetIpv6dhcpoptiondefinitionResponse

`func NewGetIpv6dhcpoptiondefinitionResponse() *GetIpv6dhcpoptiondefinitionResponse`

NewGetIpv6dhcpoptiondefinitionResponse instantiates a new GetIpv6dhcpoptiondefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6dhcpoptiondefinitionResponseWithDefaults

`func NewGetIpv6dhcpoptiondefinitionResponseWithDefaults() *GetIpv6dhcpoptiondefinitionResponse`

NewGetIpv6dhcpoptiondefinitionResponseWithDefaults instantiates a new GetIpv6dhcpoptiondefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6dhcpoptiondefinitionResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6dhcpoptiondefinitionResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCode

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetIpv6dhcpoptiondefinitionResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetIpv6dhcpoptiondefinitionResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpv6dhcpoptiondefinitionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIpv6dhcpoptiondefinitionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpace

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *GetIpv6dhcpoptiondefinitionResponse) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *GetIpv6dhcpoptiondefinitionResponse) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetType

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIpv6dhcpoptiondefinitionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIpv6dhcpoptiondefinitionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetResult() Ipv6dhcpoptiondefinition`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6dhcpoptiondefinitionResponse) GetResultOk() (*Ipv6dhcpoptiondefinition, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6dhcpoptiondefinitionResponse) SetResult(v Ipv6dhcpoptiondefinition)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6dhcpoptiondefinitionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


