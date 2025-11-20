# GetIpv6dhcpoptionspaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment of a DHCP IPv6 option space object. | [optional] 
**EnterpriseNumber** | Pointer to **int64** | The enterprise number of a DHCP IPv6 option space object. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP IPv6 option space object. | [optional] 
**OptionDefinitions** | Pointer to **[]string** | The list of DHCP IPv6 option definition objects. | [optional] 
**Result** | Pointer to [**Ipv6dhcpoptionspace**](Ipv6dhcpoptionspace.md) |  | [optional] 

## Methods

### NewGetIpv6dhcpoptionspaceResponse

`func NewGetIpv6dhcpoptionspaceResponse() *GetIpv6dhcpoptionspaceResponse`

NewGetIpv6dhcpoptionspaceResponse instantiates a new GetIpv6dhcpoptionspaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6dhcpoptionspaceResponseWithDefaults

`func NewGetIpv6dhcpoptionspaceResponseWithDefaults() *GetIpv6dhcpoptionspaceResponse`

NewGetIpv6dhcpoptionspaceResponseWithDefaults instantiates a new GetIpv6dhcpoptionspaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6dhcpoptionspaceResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6dhcpoptionspaceResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6dhcpoptionspaceResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6dhcpoptionspaceResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv6dhcpoptionspaceResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv6dhcpoptionspaceResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv6dhcpoptionspaceResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv6dhcpoptionspaceResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnterpriseNumber

`func (o *GetIpv6dhcpoptionspaceResponse) GetEnterpriseNumber() int64`

GetEnterpriseNumber returns the EnterpriseNumber field if non-nil, zero value otherwise.

### GetEnterpriseNumberOk

`func (o *GetIpv6dhcpoptionspaceResponse) GetEnterpriseNumberOk() (*int64, bool)`

GetEnterpriseNumberOk returns a tuple with the EnterpriseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseNumber

`func (o *GetIpv6dhcpoptionspaceResponse) SetEnterpriseNumber(v int64)`

SetEnterpriseNumber sets EnterpriseNumber field to given value.

### HasEnterpriseNumber

`func (o *GetIpv6dhcpoptionspaceResponse) HasEnterpriseNumber() bool`

HasEnterpriseNumber returns a boolean if a field has been set.

### GetName

`func (o *GetIpv6dhcpoptionspaceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpv6dhcpoptionspaceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpv6dhcpoptionspaceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIpv6dhcpoptionspaceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptionDefinitions

`func (o *GetIpv6dhcpoptionspaceResponse) GetOptionDefinitions() []string`

GetOptionDefinitions returns the OptionDefinitions field if non-nil, zero value otherwise.

### GetOptionDefinitionsOk

`func (o *GetIpv6dhcpoptionspaceResponse) GetOptionDefinitionsOk() (*[]string, bool)`

GetOptionDefinitionsOk returns a tuple with the OptionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDefinitions

`func (o *GetIpv6dhcpoptionspaceResponse) SetOptionDefinitions(v []string)`

SetOptionDefinitions sets OptionDefinitions field to given value.

### HasOptionDefinitions

`func (o *GetIpv6dhcpoptionspaceResponse) HasOptionDefinitions() bool`

HasOptionDefinitions returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6dhcpoptionspaceResponse) GetResult() Ipv6dhcpoptionspace`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6dhcpoptionspaceResponse) GetResultOk() (*Ipv6dhcpoptionspace, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6dhcpoptionspaceResponse) SetResult(v Ipv6dhcpoptionspace)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6dhcpoptionspaceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


