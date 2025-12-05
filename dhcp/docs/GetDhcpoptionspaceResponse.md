# GetDhcpoptionspaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment of a DHCP option space object. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP option space object. | [optional] 
**OptionDefinitions** | Pointer to **[]string** | The list of DHCP option definition objects. | [optional] 
**SpaceType** | Pointer to **string** | The type of a DHCP option space object. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Dhcpoptionspace**](Dhcpoptionspace.md) |  | [optional] 

## Methods

### NewGetDhcpoptionspaceResponse

`func NewGetDhcpoptionspaceResponse() *GetDhcpoptionspaceResponse`

NewGetDhcpoptionspaceResponse instantiates a new GetDhcpoptionspaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDhcpoptionspaceResponseWithDefaults

`func NewGetDhcpoptionspaceResponseWithDefaults() *GetDhcpoptionspaceResponse`

NewGetDhcpoptionspaceResponseWithDefaults instantiates a new GetDhcpoptionspaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDhcpoptionspaceResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDhcpoptionspaceResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDhcpoptionspaceResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDhcpoptionspaceResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetDhcpoptionspaceResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDhcpoptionspaceResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDhcpoptionspaceResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDhcpoptionspaceResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *GetDhcpoptionspaceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDhcpoptionspaceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDhcpoptionspaceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDhcpoptionspaceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptionDefinitions

`func (o *GetDhcpoptionspaceResponse) GetOptionDefinitions() []string`

GetOptionDefinitions returns the OptionDefinitions field if non-nil, zero value otherwise.

### GetOptionDefinitionsOk

`func (o *GetDhcpoptionspaceResponse) GetOptionDefinitionsOk() (*[]string, bool)`

GetOptionDefinitionsOk returns a tuple with the OptionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDefinitions

`func (o *GetDhcpoptionspaceResponse) SetOptionDefinitions(v []string)`

SetOptionDefinitions sets OptionDefinitions field to given value.

### HasOptionDefinitions

`func (o *GetDhcpoptionspaceResponse) HasOptionDefinitions() bool`

HasOptionDefinitions returns a boolean if a field has been set.

### GetSpaceType

`func (o *GetDhcpoptionspaceResponse) GetSpaceType() string`

GetSpaceType returns the SpaceType field if non-nil, zero value otherwise.

### GetSpaceTypeOk

`func (o *GetDhcpoptionspaceResponse) GetSpaceTypeOk() (*string, bool)`

GetSpaceTypeOk returns a tuple with the SpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceType

`func (o *GetDhcpoptionspaceResponse) SetSpaceType(v string)`

SetSpaceType sets SpaceType field to given value.

### HasSpaceType

`func (o *GetDhcpoptionspaceResponse) HasSpaceType() bool`

HasSpaceType returns a boolean if a field has been set.

### GetUuid

`func (o *GetDhcpoptionspaceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDhcpoptionspaceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDhcpoptionspaceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDhcpoptionspaceResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDhcpoptionspaceResponse) GetResult() Dhcpoptionspace`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDhcpoptionspaceResponse) GetResultOk() (*Dhcpoptionspace, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDhcpoptionspaceResponse) SetResult(v Dhcpoptionspace)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDhcpoptionspaceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


