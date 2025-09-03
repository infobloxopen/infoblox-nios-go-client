# GetBulkhostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CloudInfo** | Pointer to [**BulkhostCloudInfo**](BulkhostCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | The descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | The disable flag of a DNS BulkHost record. | [optional] 
**DnsPrefix** | Pointer to **string** | The prefix, in punycode format, for the bulk host. | [optional] [readonly] 
**EndAddr** | Pointer to **string** | The last IP address in the address range for the bulk host. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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
**Result** | Pointer to [**Bulkhost**](Bulkhost.md) |  | [optional] 

## Methods

### NewGetBulkhostResponse

`func NewGetBulkhostResponse() *GetBulkhostResponse`

NewGetBulkhostResponse instantiates a new GetBulkhostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBulkhostResponseWithDefaults

`func NewGetBulkhostResponseWithDefaults() *GetBulkhostResponse`

NewGetBulkhostResponseWithDefaults instantiates a new GetBulkhostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetBulkhostResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetBulkhostResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetBulkhostResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetBulkhostResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetBulkhostResponse) GetCloudInfo() BulkhostCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetBulkhostResponse) GetCloudInfoOk() (*BulkhostCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetBulkhostResponse) SetCloudInfo(v BulkhostCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetBulkhostResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetBulkhostResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetBulkhostResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetBulkhostResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetBulkhostResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetBulkhostResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetBulkhostResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetBulkhostResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetBulkhostResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsPrefix

`func (o *GetBulkhostResponse) GetDnsPrefix() string`

GetDnsPrefix returns the DnsPrefix field if non-nil, zero value otherwise.

### GetDnsPrefixOk

`func (o *GetBulkhostResponse) GetDnsPrefixOk() (*string, bool)`

GetDnsPrefixOk returns a tuple with the DnsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrefix

`func (o *GetBulkhostResponse) SetDnsPrefix(v string)`

SetDnsPrefix sets DnsPrefix field to given value.

### HasDnsPrefix

`func (o *GetBulkhostResponse) HasDnsPrefix() bool`

HasDnsPrefix returns a boolean if a field has been set.

### GetEndAddr

`func (o *GetBulkhostResponse) GetEndAddr() string`

GetEndAddr returns the EndAddr field if non-nil, zero value otherwise.

### GetEndAddrOk

`func (o *GetBulkhostResponse) GetEndAddrOk() (*string, bool)`

GetEndAddrOk returns a tuple with the EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddr

`func (o *GetBulkhostResponse) SetEndAddr(v string)`

SetEndAddr sets EndAddr field to given value.

### HasEndAddr

`func (o *GetBulkhostResponse) HasEndAddr() bool`

HasEndAddr returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetBulkhostResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetBulkhostResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetBulkhostResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetBulkhostResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetBulkhostResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetBulkhostResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetBulkhostResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetBulkhostResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetBulkhostResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetBulkhostResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetBulkhostResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetBulkhostResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetBulkhostResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetBulkhostResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetBulkhostResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetBulkhostResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetNameTemplate

`func (o *GetBulkhostResponse) GetNameTemplate() string`

GetNameTemplate returns the NameTemplate field if non-nil, zero value otherwise.

### GetNameTemplateOk

`func (o *GetBulkhostResponse) GetNameTemplateOk() (*string, bool)`

GetNameTemplateOk returns a tuple with the NameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTemplate

`func (o *GetBulkhostResponse) SetNameTemplate(v string)`

SetNameTemplate sets NameTemplate field to given value.

### HasNameTemplate

`func (o *GetBulkhostResponse) HasNameTemplate() bool`

HasNameTemplate returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetBulkhostResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetBulkhostResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetBulkhostResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetBulkhostResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetPolicy

`func (o *GetBulkhostResponse) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetBulkhostResponse) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetBulkhostResponse) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetBulkhostResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPrefix

`func (o *GetBulkhostResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetBulkhostResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetBulkhostResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetBulkhostResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetReverse

`func (o *GetBulkhostResponse) GetReverse() bool`

GetReverse returns the Reverse field if non-nil, zero value otherwise.

### GetReverseOk

`func (o *GetBulkhostResponse) GetReverseOk() (*bool, bool)`

GetReverseOk returns a tuple with the Reverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverse

`func (o *GetBulkhostResponse) SetReverse(v bool)`

SetReverse sets Reverse field to given value.

### HasReverse

`func (o *GetBulkhostResponse) HasReverse() bool`

HasReverse returns a boolean if a field has been set.

### GetStartAddr

`func (o *GetBulkhostResponse) GetStartAddr() string`

GetStartAddr returns the StartAddr field if non-nil, zero value otherwise.

### GetStartAddrOk

`func (o *GetBulkhostResponse) GetStartAddrOk() (*string, bool)`

GetStartAddrOk returns a tuple with the StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddr

`func (o *GetBulkhostResponse) SetStartAddr(v string)`

SetStartAddr sets StartAddr field to given value.

### HasStartAddr

`func (o *GetBulkhostResponse) HasStartAddr() bool`

HasStartAddr returns a boolean if a field has been set.

### GetTemplateFormat

`func (o *GetBulkhostResponse) GetTemplateFormat() string`

GetTemplateFormat returns the TemplateFormat field if non-nil, zero value otherwise.

### GetTemplateFormatOk

`func (o *GetBulkhostResponse) GetTemplateFormatOk() (*string, bool)`

GetTemplateFormatOk returns a tuple with the TemplateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateFormat

`func (o *GetBulkhostResponse) SetTemplateFormat(v string)`

SetTemplateFormat sets TemplateFormat field to given value.

### HasTemplateFormat

`func (o *GetBulkhostResponse) HasTemplateFormat() bool`

HasTemplateFormat returns a boolean if a field has been set.

### GetTtl

`func (o *GetBulkhostResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetBulkhostResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetBulkhostResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetBulkhostResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseNameTemplate

`func (o *GetBulkhostResponse) GetUseNameTemplate() bool`

GetUseNameTemplate returns the UseNameTemplate field if non-nil, zero value otherwise.

### GetUseNameTemplateOk

`func (o *GetBulkhostResponse) GetUseNameTemplateOk() (*bool, bool)`

GetUseNameTemplateOk returns a tuple with the UseNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNameTemplate

`func (o *GetBulkhostResponse) SetUseNameTemplate(v bool)`

SetUseNameTemplate sets UseNameTemplate field to given value.

### HasUseNameTemplate

`func (o *GetBulkhostResponse) HasUseNameTemplate() bool`

HasUseNameTemplate returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetBulkhostResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetBulkhostResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetBulkhostResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetBulkhostResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *GetBulkhostResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetBulkhostResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetBulkhostResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetBulkhostResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetBulkhostResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetBulkhostResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetBulkhostResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetBulkhostResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetBulkhostResponse) GetResult() Bulkhost`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetBulkhostResponse) GetResultOk() (*Bulkhost, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetBulkhostResponse) SetResult(v Bulkhost)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetBulkhostResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


