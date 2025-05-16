# Bulkhost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CloudInfo** | Pointer to [**BulkhostCloudInfo**](BulkhostCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | The descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | The disable flag of a DNS BulkHost record. | [optional] 
**DnsPrefix** | Pointer to **string** | The prefix, in punycode format, for the bulk host. | [optional] [readonly] 
**EndAddr** | Pointer to **string** | The last IP address in the address range for the bulk host. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**NameTemplate** | Pointer to **string** | The bulk host name template. | [optional] 
**NetworkView** | Pointer to **string** | The network view associated with the bulk host view. | [optional] [readonly] 
**Policy** | Pointer to **string** | The hostname policy for records under the bulk host parent zone. | [optional] [readonly] 
**Prefix** | Pointer to **string** | The prefix for the bulk host. The prefix is the name (or a series of characters) inserted at the beginning of each host name. | [optional] 
**Reverse** | Pointer to **bool** | The reverse flag of the BulkHost record. | [optional] 
**StartAddr** | Pointer to **string** | The first IP address in the address range for the bulk host. | [optional] 
**TemplateFormat** | Pointer to **string** | The bulk host name template format. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time to Live (TTL) value. | [optional] 
**UseNameTemplate** | Pointer to **bool** | Use flag for: name_template | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The view for the bulk host. | [optional] 
**Zone** | Pointer to **string** | The zone name. | [optional] 

## Methods

### NewBulkhost

`func NewBulkhost() *Bulkhost`

NewBulkhost instantiates a new Bulkhost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkhostWithDefaults

`func NewBulkhostWithDefaults() *Bulkhost`

NewBulkhostWithDefaults instantiates a new Bulkhost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Bulkhost) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Bulkhost) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Bulkhost) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Bulkhost) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudInfo

`func (o *Bulkhost) GetCloudInfo() BulkhostCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *Bulkhost) GetCloudInfoOk() (*BulkhostCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *Bulkhost) SetCloudInfo(v BulkhostCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *Bulkhost) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *Bulkhost) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Bulkhost) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Bulkhost) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Bulkhost) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *Bulkhost) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Bulkhost) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Bulkhost) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Bulkhost) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsPrefix

`func (o *Bulkhost) GetDnsPrefix() string`

GetDnsPrefix returns the DnsPrefix field if non-nil, zero value otherwise.

### GetDnsPrefixOk

`func (o *Bulkhost) GetDnsPrefixOk() (*string, bool)`

GetDnsPrefixOk returns a tuple with the DnsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrefix

`func (o *Bulkhost) SetDnsPrefix(v string)`

SetDnsPrefix sets DnsPrefix field to given value.

### HasDnsPrefix

`func (o *Bulkhost) HasDnsPrefix() bool`

HasDnsPrefix returns a boolean if a field has been set.

### GetEndAddr

`func (o *Bulkhost) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *Bulkhost) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *Bulkhost) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *Bulkhost) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetExtattrs

`func (o *Bulkhost) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Bulkhost) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Bulkhost) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Bulkhost) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetLastQueried

`func (o *Bulkhost) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *Bulkhost) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *Bulkhost) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *Bulkhost) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetNameTemplate

`func (o *Bulkhost) GetNameTemplate() string`

GetNameTemplate returns the NameTemplate field if non-nil, zero value otherwise.

### GetNameTemplateOk

`func (o *Bulkhost) GetNameTemplateOk() (*string, bool)`

GetNameTemplateOk returns a tuple with the NameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTemplate

`func (o *Bulkhost) SetNameTemplate(v string)`

SetNameTemplate sets NameTemplate field to given value.

### HasNameTemplate

`func (o *Bulkhost) HasNameTemplate() bool`

HasNameTemplate returns a boolean if a field has been set.

### GetNetworkView

`func (o *Bulkhost) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Bulkhost) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Bulkhost) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Bulkhost) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetPolicy

`func (o *Bulkhost) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Bulkhost) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Bulkhost) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *Bulkhost) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPrefix

`func (o *Bulkhost) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *Bulkhost) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *Bulkhost) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *Bulkhost) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetReverse

`func (o *Bulkhost) GetReverse() bool`

GetReverse returns the Reverse field if non-nil, zero value otherwise.

### GetReverseOk

`func (o *Bulkhost) GetReverseOk() (*bool, bool)`

GetReverseOk returns a tuple with the Reverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverse

`func (o *Bulkhost) SetReverse(v bool)`

SetReverse sets Reverse field to given value.

### HasReverse

`func (o *Bulkhost) HasReverse() bool`

HasReverse returns a boolean if a field has been set.

### GetStartAddr

`func (o *Bulkhost) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *Bulkhost) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *Bulkhost) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *Bulkhost) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.

### GetTemplateFormat

`func (o *Bulkhost) GetTemplateFormat() string`

GetTemplateFormat returns the TemplateFormat field if non-nil, zero value otherwise.

### GetTemplateFormatOk

`func (o *Bulkhost) GetTemplateFormatOk() (*string, bool)`

GetTemplateFormatOk returns a tuple with the TemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFormat

`func (o *Bulkhost) SetTemplateFormat(v string)`

SetTemplateFormat sets TemplateFormat field to given value.

### HasTemplateFormat

`func (o *Bulkhost) HasTemplateFormat() bool`

HasTemplateFormat returns a boolean if a field has been set.

### GetTtl

`func (o *Bulkhost) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Bulkhost) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Bulkhost) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Bulkhost) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseNameTemplate

`func (o *Bulkhost) GetUseNameTemplate() bool`

GetUseNameTemplate returns the UseNameTemplate field if non-nil, zero value otherwise.

### GetUseNameTemplateOk

`func (o *Bulkhost) GetUseNameTemplateOk() (*bool, bool)`

GetUseNameTemplateOk returns a tuple with the UseNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNameTemplate

`func (o *Bulkhost) SetUseNameTemplate(v bool)`

SetUseNameTemplate sets UseNameTemplate field to given value.

### HasUseNameTemplate

`func (o *Bulkhost) HasUseNameTemplate() bool`

HasUseNameTemplate returns a boolean if a field has been set.

### GetUseTtl

`func (o *Bulkhost) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *Bulkhost) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *Bulkhost) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *Bulkhost) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *Bulkhost) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Bulkhost) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Bulkhost) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *Bulkhost) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *Bulkhost) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Bulkhost) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Bulkhost) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Bulkhost) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


