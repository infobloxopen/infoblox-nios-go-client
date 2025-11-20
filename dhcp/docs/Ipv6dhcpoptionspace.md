# Ipv6dhcpoptionspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment of a DHCP IPv6 option space object. | [optional] 
**EnterpriseNumber** | Pointer to **int64** | The enterprise number of a DHCP IPv6 option space object. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP IPv6 option space object. | [optional] 
**OptionDefinitions** | Pointer to **[]string** | The list of DHCP IPv6 option definition objects. | [optional] 

## Methods

### NewIpv6dhcpoptionspace

`func NewIpv6dhcpoptionspace() *Ipv6dhcpoptionspace`

NewIpv6dhcpoptionspace instantiates a new Ipv6dhcpoptionspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6dhcpoptionspaceWithDefaults

`func NewIpv6dhcpoptionspaceWithDefaults() *Ipv6dhcpoptionspace`

NewIpv6dhcpoptionspaceWithDefaults instantiates a new Ipv6dhcpoptionspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6dhcpoptionspace) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6dhcpoptionspace) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6dhcpoptionspace) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6dhcpoptionspace) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6dhcpoptionspace) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6dhcpoptionspace) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6dhcpoptionspace) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6dhcpoptionspace) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnterpriseNumber

`func (o *Ipv6dhcpoptionspace) GetEnterpriseNumber() int64`

GetEnterpriseNumber returns the EnterpriseNumber field if non-nil, zero value otherwise.

### GetEnterpriseNumberOk

`func (o *Ipv6dhcpoptionspace) GetEnterpriseNumberOk() (*int64, bool)`

GetEnterpriseNumberOk returns a tuple with the EnterpriseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseNumber

`func (o *Ipv6dhcpoptionspace) SetEnterpriseNumber(v int64)`

SetEnterpriseNumber sets EnterpriseNumber field to given value.

### HasEnterpriseNumber

`func (o *Ipv6dhcpoptionspace) HasEnterpriseNumber() bool`

HasEnterpriseNumber returns a boolean if a field has been set.

### GetName

`func (o *Ipv6dhcpoptionspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ipv6dhcpoptionspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ipv6dhcpoptionspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ipv6dhcpoptionspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptionDefinitions

`func (o *Ipv6dhcpoptionspace) GetOptionDefinitions() []string`

GetOptionDefinitions returns the OptionDefinitions field if non-nil, zero value otherwise.

### GetOptionDefinitionsOk

`func (o *Ipv6dhcpoptionspace) GetOptionDefinitionsOk() (*[]string, bool)`

GetOptionDefinitionsOk returns a tuple with the OptionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDefinitions

`func (o *Ipv6dhcpoptionspace) SetOptionDefinitions(v []string)`

SetOptionDefinitions sets OptionDefinitions field to given value.

### HasOptionDefinitions

`func (o *Ipv6dhcpoptionspace) HasOptionDefinitions() bool`

HasOptionDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


