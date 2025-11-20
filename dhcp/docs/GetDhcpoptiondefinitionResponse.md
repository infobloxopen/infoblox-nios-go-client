# GetDhcpoptiondefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Code** | Pointer to **int64** | The code of a DHCP option definition object. An option code number is used to identify the DHCP option. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP option definition object. | [optional] 
**Space** | Pointer to **string** | The space of a DHCP option definition object. | [optional] 
**Type** | Pointer to **string** | The data type of the Grid DHCP option. | [optional] 
**Result** | Pointer to [**Dhcpoptiondefinition**](Dhcpoptiondefinition.md) |  | [optional] 

## Methods

### NewGetDhcpoptiondefinitionResponse

`func NewGetDhcpoptiondefinitionResponse() *GetDhcpoptiondefinitionResponse`

NewGetDhcpoptiondefinitionResponse instantiates a new GetDhcpoptiondefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDhcpoptiondefinitionResponseWithDefaults

`func NewGetDhcpoptiondefinitionResponseWithDefaults() *GetDhcpoptiondefinitionResponse`

NewGetDhcpoptiondefinitionResponseWithDefaults instantiates a new GetDhcpoptiondefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDhcpoptiondefinitionResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDhcpoptiondefinitionResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDhcpoptiondefinitionResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDhcpoptiondefinitionResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCode

`func (o *GetDhcpoptiondefinitionResponse) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetDhcpoptiondefinitionResponse) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetDhcpoptiondefinitionResponse) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetDhcpoptiondefinitionResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetDhcpoptiondefinitionResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDhcpoptiondefinitionResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDhcpoptiondefinitionResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDhcpoptiondefinitionResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpace

`func (o *GetDhcpoptiondefinitionResponse) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *GetDhcpoptiondefinitionResponse) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *GetDhcpoptiondefinitionResponse) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *GetDhcpoptiondefinitionResponse) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetType

`func (o *GetDhcpoptiondefinitionResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDhcpoptiondefinitionResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDhcpoptiondefinitionResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDhcpoptiondefinitionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *GetDhcpoptiondefinitionResponse) GetResult() Dhcpoptiondefinition`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDhcpoptiondefinitionResponse) GetResultOk() (*Dhcpoptiondefinition, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDhcpoptiondefinitionResponse) SetResult(v Dhcpoptiondefinition)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDhcpoptiondefinitionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


