# Dhcpoptionspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment of a DHCP option space object. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP option space object. | [optional] 
**OptionDefinitions** | Pointer to **[]string** | The list of DHCP option definition objects. | [optional] 
**SpaceType** | Pointer to **string** | The type of a DHCP option space object. | [optional] [readonly] 

## Methods

### NewDhcpoptionspace

`func NewDhcpoptionspace() *Dhcpoptionspace`

NewDhcpoptionspace instantiates a new Dhcpoptionspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpoptionspaceWithDefaults

`func NewDhcpoptionspaceWithDefaults() *Dhcpoptionspace`

NewDhcpoptionspaceWithDefaults instantiates a new Dhcpoptionspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Dhcpoptionspace) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Dhcpoptionspace) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Dhcpoptionspace) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Dhcpoptionspace) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Dhcpoptionspace) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Dhcpoptionspace) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Dhcpoptionspace) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Dhcpoptionspace) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *Dhcpoptionspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dhcpoptionspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dhcpoptionspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dhcpoptionspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptionDefinitions

`func (o *Dhcpoptionspace) GetOptionDefinitions() []string`

GetOptionDefinitions returns the OptionDefinitions field if non-nil, zero value otherwise.

### GetOptionDefinitionsOk

`func (o *Dhcpoptionspace) GetOptionDefinitionsOk() (*[]string, bool)`

GetOptionDefinitionsOk returns a tuple with the OptionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDefinitions

`func (o *Dhcpoptionspace) SetOptionDefinitions(v []string)`

SetOptionDefinitions sets OptionDefinitions field to given value.

### HasOptionDefinitions

`func (o *Dhcpoptionspace) HasOptionDefinitions() bool`

HasOptionDefinitions returns a boolean if a field has been set.

### GetSpaceType

`func (o *Dhcpoptionspace) GetSpaceType() string`

GetSpaceType returns the SpaceType field if non-nil, zero value otherwise.

### GetSpaceTypeOk

`func (o *Dhcpoptionspace) GetSpaceTypeOk() (*string, bool)`

GetSpaceTypeOk returns a tuple with the SpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceType

`func (o *Dhcpoptionspace) SetSpaceType(v string)`

SetSpaceType sets SpaceType field to given value.

### HasSpaceType

`func (o *Dhcpoptionspace) HasSpaceType() bool`

HasSpaceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


