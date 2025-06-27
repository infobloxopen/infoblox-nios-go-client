# GetIpv6fixedaddresstemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | A descriptive comment of an IPv6 fixed address template object. | [optional] 
**DomainName** | Pointer to **string** | Domain name of the IPv6 fixed address template object. | [optional] 
**DomainNameServers** | Pointer to **[]string** | The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LogicFilterRules** | Pointer to [**[]Ipv6fixedaddresstemplateLogicFilterRules**](Ipv6fixedaddresstemplateLogicFilterRules.md) | This field contains the logic filters to be applied to this IPv6 fixed address. This list corresponds to the match rules that are written to the DHCPv6 configuration file. | [optional] 
**Name** | Pointer to **string** | Name of an IPv6 fixed address template object. | [optional] 
**NumberOfAddresses** | Pointer to **int64** | The number of IPv6 addresses for this fixed address. | [optional] 
**Offset** | Pointer to **int64** | The start address offset for this IPv6 fixed address. | [optional] 
**Options** | Pointer to [**[]Ipv6fixedaddresstemplateOptions**](Ipv6fixedaddresstemplateOptions.md) | An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object. | [optional] 
**PreferredLifetime** | Pointer to **int64** | The preferred lifetime value for this DHCP IPv6 fixed address template object. | [optional] 
**UseDomainName** | Pointer to **bool** | Use flag for: domain_name | [optional] 
**UseDomainNameServers** | Pointer to **bool** | Use flag for: domain_name_servers | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseOptions** | Pointer to **bool** | Use flag for: options | [optional] 
**UsePreferredLifetime** | Pointer to **bool** | Use flag for: preferred_lifetime | [optional] 
**UseValidLifetime** | Pointer to **bool** | Use flag for: valid_lifetime | [optional] 
**ValidLifetime** | Pointer to **int64** | The valid lifetime value for this DHCP IPv6 fixed address template object. | [optional] 
**Result** | Pointer to [**Ipv6fixedaddresstemplate**](Ipv6fixedaddresstemplate.md) |  | [optional] 

## Methods

### NewGetIpv6fixedaddresstemplateResponse

`func NewGetIpv6fixedaddresstemplateResponse() *GetIpv6fixedaddresstemplateResponse`

NewGetIpv6fixedaddresstemplateResponse instantiates a new GetIpv6fixedaddresstemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6fixedaddresstemplateResponseWithDefaults

`func NewGetIpv6fixedaddresstemplateResponseWithDefaults() *GetIpv6fixedaddresstemplateResponse`

NewGetIpv6fixedaddresstemplateResponseWithDefaults instantiates a new GetIpv6fixedaddresstemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6fixedaddresstemplateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6fixedaddresstemplateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6fixedaddresstemplateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv6fixedaddresstemplateResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv6fixedaddresstemplateResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv6fixedaddresstemplateResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDomainName

`func (o *GetIpv6fixedaddresstemplateResponse) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetIpv6fixedaddresstemplateResponse) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetIpv6fixedaddresstemplateResponse) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *GetIpv6fixedaddresstemplateResponse) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *GetIpv6fixedaddresstemplateResponse) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *GetIpv6fixedaddresstemplateResponse) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetIpv6fixedaddresstemplateResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetIpv6fixedaddresstemplateResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetIpv6fixedaddresstemplateResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetIpv6fixedaddresstemplateResponse) GetLogicFilterRules() []Ipv6fixedaddresstemplateLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetLogicFilterRulesOk() (*[]Ipv6fixedaddresstemplateLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetIpv6fixedaddresstemplateResponse) SetLogicFilterRules(v []Ipv6fixedaddresstemplateLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetIpv6fixedaddresstemplateResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetName

`func (o *GetIpv6fixedaddresstemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpv6fixedaddresstemplateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIpv6fixedaddresstemplateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfAddresses

`func (o *GetIpv6fixedaddresstemplateResponse) GetNumberOfAddresses() int64`

GetNumberOfAddresses returns the NumberOfAddresses field if non-nil, zero value otherwise.

### GetNumberOfAddressesOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetNumberOfAddressesOk() (*int64, bool)`

GetNumberOfAddressesOk returns a tuple with the NumberOfAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAddresses

`func (o *GetIpv6fixedaddresstemplateResponse) SetNumberOfAddresses(v int64)`

SetNumberOfAddresses sets NumberOfAddresses field to given value.

### HasNumberOfAddresses

`func (o *GetIpv6fixedaddresstemplateResponse) HasNumberOfAddresses() bool`

HasNumberOfAddresses returns a boolean if a field has been set.

### GetOffset

`func (o *GetIpv6fixedaddresstemplateResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetIpv6fixedaddresstemplateResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetIpv6fixedaddresstemplateResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOptions

`func (o *GetIpv6fixedaddresstemplateResponse) GetOptions() []Ipv6fixedaddresstemplateOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetOptionsOk() (*[]Ipv6fixedaddresstemplateOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetIpv6fixedaddresstemplateResponse) SetOptions(v []Ipv6fixedaddresstemplateOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetIpv6fixedaddresstemplateResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetUseDomainName

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseDomainName() bool`

GetUseDomainName returns the UseDomainName field if non-nil, zero value otherwise.

### GetUseDomainNameOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseDomainNameOk() (*bool, bool)`

GetUseDomainNameOk returns a tuple with the UseDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainName

`func (o *GetIpv6fixedaddresstemplateResponse) SetUseDomainName(v bool)`

SetUseDomainName sets UseDomainName field to given value.

### HasUseDomainName

`func (o *GetIpv6fixedaddresstemplateResponse) HasUseDomainName() bool`

HasUseDomainName returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *GetIpv6fixedaddresstemplateResponse) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *GetIpv6fixedaddresstemplateResponse) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetIpv6fixedaddresstemplateResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetIpv6fixedaddresstemplateResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseOptions

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *GetIpv6fixedaddresstemplateResponse) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *GetIpv6fixedaddresstemplateResponse) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *GetIpv6fixedaddresstemplateResponse) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6fixedaddresstemplateResponse) GetResult() Ipv6fixedaddresstemplate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6fixedaddresstemplateResponse) GetResultOk() (*Ipv6fixedaddresstemplate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6fixedaddresstemplateResponse) SetResult(v Ipv6fixedaddresstemplate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6fixedaddresstemplateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


