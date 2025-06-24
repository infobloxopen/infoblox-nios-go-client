# GetDns64groupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Clients** | Pointer to [**[]Dns64groupClients**](Dns64groupClients.md) | Access Control settings that contain IPv4 and IPv6 DNS clients and networks to which the DNS server is allowed to send synthesized AAAA records with the specified IPv6 prefix. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment for the DNS64 synthesis group object. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the DNS64 synthesis group is disabled. | [optional] 
**EnableDnssecDns64** | Pointer to **bool** | Determines whether the DNS64 synthesis of AAAA records is enabled for DNS64 synthesis groups that request DNSSEC data. | [optional] 
**Exclude** | Pointer to [**[]Dns64groupExclude**](Dns64groupExclude.md) | Access Control settings that contain IPv6 addresses or prefix ranges that cannot be used by IPv6-only hosts, such as IP addresses in the ::ffff:0:0/96 network. When DNS server retrieves an AAAA record that contains an IPv6 address that matches an excluded address, it does not return the AAAA record. Instead it synthesizes an AAAA record from the A record. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Mapped** | Pointer to [**[]Dns64groupMapped**](Dns64groupMapped.md) | Access Control settings that contain IPv4 addresses and networks for which the DNS server can synthesize AAAA records with the specified prefix. | [optional] 
**Name** | Pointer to **string** | The name of the DNS64 synthesis group object. | [optional] 
**Prefix** | Pointer to **string** | The IPv6 prefix used for the synthesized AAAA records. The prefix length must be /32, /40, /48, /56, /64 or /96, and all bits beyond the specified length must be zero. | [optional] 
**Result** | Pointer to [**Dns64group**](Dns64group.md) |  | [optional] 

## Methods

### NewGetDns64groupResponse

`func NewGetDns64groupResponse() *GetDns64groupResponse`

NewGetDns64groupResponse instantiates a new GetDns64groupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDns64groupResponseWithDefaults

`func NewGetDns64groupResponseWithDefaults() *GetDns64groupResponse`

NewGetDns64groupResponseWithDefaults instantiates a new GetDns64groupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDns64groupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDns64groupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDns64groupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDns64groupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetClients

`func (o *GetDns64groupResponse) GetClients() []Dns64groupClients`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *GetDns64groupResponse) GetClientsOk() (*[]Dns64groupClients, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *GetDns64groupResponse) SetClients(v []Dns64groupClients)`

SetClients sets Clients field to given value.

### HasClients

`func (o *GetDns64groupResponse) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetComment

`func (o *GetDns64groupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDns64groupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDns64groupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDns64groupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetDns64groupResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetDns64groupResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetDns64groupResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetDns64groupResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEnableDnssecDns64

`func (o *GetDns64groupResponse) GetEnableDnssecDns64() bool`

GetEnableDnssecDns64 returns the EnableDnssecDns64 field if non-nil, zero value otherwise.

### GetEnableDnssecDns64Ok

`func (o *GetDns64groupResponse) GetEnableDnssecDns64Ok() (*bool, bool)`

GetEnableDnssecDns64Ok returns a tuple with the EnableDnssecDns64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnssecDns64

`func (o *GetDns64groupResponse) SetEnableDnssecDns64(v bool)`

SetEnableDnssecDns64 sets EnableDnssecDns64 field to given value.

### HasEnableDnssecDns64

`func (o *GetDns64groupResponse) HasEnableDnssecDns64() bool`

HasEnableDnssecDns64 returns a boolean if a field has been set.

### GetExclude

`func (o *GetDns64groupResponse) GetExclude() []Dns64groupExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *GetDns64groupResponse) GetExcludeOk() (*[]Dns64groupExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *GetDns64groupResponse) SetExclude(v []Dns64groupExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *GetDns64groupResponse) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetDns64groupResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetDns64groupResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetDns64groupResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetDns64groupResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetMapped

`func (o *GetDns64groupResponse) GetMapped() []Dns64groupMapped`

GetMapped returns the Mapped field if non-nil, zero value otherwise.

### GetMappedOk

`func (o *GetDns64groupResponse) GetMappedOk() (*[]Dns64groupMapped, bool)`

GetMappedOk returns a tuple with the Mapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapped

`func (o *GetDns64groupResponse) SetMapped(v []Dns64groupMapped)`

SetMapped sets Mapped field to given value.

### HasMapped

`func (o *GetDns64groupResponse) HasMapped() bool`

HasMapped returns a boolean if a field has been set.

### GetName

`func (o *GetDns64groupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDns64groupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDns64groupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDns64groupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrefix

`func (o *GetDns64groupResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetDns64groupResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetDns64groupResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetDns64groupResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetResult

`func (o *GetDns64groupResponse) GetResult() Dns64group`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDns64groupResponse) GetResultOk() (*Dns64group, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDns64groupResponse) SetResult(v Dns64group)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDns64groupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


