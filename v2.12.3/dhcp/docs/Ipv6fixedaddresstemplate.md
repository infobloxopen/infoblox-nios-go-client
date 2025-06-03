# Ipv6fixedaddresstemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | A descriptive comment of an IPv6 fixed address template object. | [optional] 
**DomainName** | Pointer to **string** | Domain name of the IPv6 fixed address template object. | [optional] 
**DomainNameServers** | Pointer to **[]string** | The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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

## Methods

### NewIpv6fixedaddresstemplate

`func NewIpv6fixedaddresstemplate() *Ipv6fixedaddresstemplate`

NewIpv6fixedaddresstemplate instantiates a new Ipv6fixedaddresstemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6fixedaddresstemplateWithDefaults

`func NewIpv6fixedaddresstemplateWithDefaults() *Ipv6fixedaddresstemplate`

NewIpv6fixedaddresstemplateWithDefaults instantiates a new Ipv6fixedaddresstemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6fixedaddresstemplate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6fixedaddresstemplate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6fixedaddresstemplate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6fixedaddresstemplate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6fixedaddresstemplate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6fixedaddresstemplate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6fixedaddresstemplate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6fixedaddresstemplate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDomainName

`func (o *Ipv6fixedaddresstemplate) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Ipv6fixedaddresstemplate) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Ipv6fixedaddresstemplate) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Ipv6fixedaddresstemplate) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainNameServers

`func (o *Ipv6fixedaddresstemplate) GetDomainNameServers() []string`

GetDomainNameServers returns the DomainNameServers field if non-nil, zero value otherwise.

### GetDomainNameServersOk

`func (o *Ipv6fixedaddresstemplate) GetDomainNameServersOk() (*[]string, bool)`

GetDomainNameServersOk returns a tuple with the DomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNameServers

`func (o *Ipv6fixedaddresstemplate) SetDomainNameServers(v []string)`

SetDomainNameServers sets DomainNameServers field to given value.

### HasDomainNameServers

`func (o *Ipv6fixedaddresstemplate) HasDomainNameServers() bool`

HasDomainNameServers returns a boolean if a field has been set.

### GetExtattrs

`func (o *Ipv6fixedaddresstemplate) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Ipv6fixedaddresstemplate) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Ipv6fixedaddresstemplate) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Ipv6fixedaddresstemplate) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Ipv6fixedaddresstemplate) GetLogicFilterRules() []Ipv6fixedaddresstemplateLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Ipv6fixedaddresstemplate) GetLogicFilterRulesOk() (*[]Ipv6fixedaddresstemplateLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Ipv6fixedaddresstemplate) SetLogicFilterRules(v []Ipv6fixedaddresstemplateLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Ipv6fixedaddresstemplate) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetName

`func (o *Ipv6fixedaddresstemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ipv6fixedaddresstemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ipv6fixedaddresstemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ipv6fixedaddresstemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfAddresses

`func (o *Ipv6fixedaddresstemplate) GetNumberOfAddresses() int64`

GetNumberOfAddresses returns the NumberOfAddresses field if non-nil, zero value otherwise.

### GetNumberOfAddressesOk

`func (o *Ipv6fixedaddresstemplate) GetNumberOfAddressesOk() (*int64, bool)`

GetNumberOfAddressesOk returns a tuple with the NumberOfAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAddresses

`func (o *Ipv6fixedaddresstemplate) SetNumberOfAddresses(v int64)`

SetNumberOfAddresses sets NumberOfAddresses field to given value.

### HasNumberOfAddresses

`func (o *Ipv6fixedaddresstemplate) HasNumberOfAddresses() bool`

HasNumberOfAddresses returns a boolean if a field has been set.

### GetOffset

`func (o *Ipv6fixedaddresstemplate) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Ipv6fixedaddresstemplate) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Ipv6fixedaddresstemplate) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Ipv6fixedaddresstemplate) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOptions

`func (o *Ipv6fixedaddresstemplate) GetOptions() []Ipv6fixedaddresstemplateOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Ipv6fixedaddresstemplate) GetOptionsOk() (*[]Ipv6fixedaddresstemplateOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Ipv6fixedaddresstemplate) SetOptions(v []Ipv6fixedaddresstemplateOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Ipv6fixedaddresstemplate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *Ipv6fixedaddresstemplate) GetPreferredLifetime() int64`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *Ipv6fixedaddresstemplate) GetPreferredLifetimeOk() (*int64, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *Ipv6fixedaddresstemplate) SetPreferredLifetime(v int64)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *Ipv6fixedaddresstemplate) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetUseDomainName

`func (o *Ipv6fixedaddresstemplate) GetUseDomainName() bool`

GetUseDomainName returns the UseDomainName field if non-nil, zero value otherwise.

### GetUseDomainNameOk

`func (o *Ipv6fixedaddresstemplate) GetUseDomainNameOk() (*bool, bool)`

GetUseDomainNameOk returns a tuple with the UseDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainName

`func (o *Ipv6fixedaddresstemplate) SetUseDomainName(v bool)`

SetUseDomainName sets UseDomainName field to given value.

### HasUseDomainName

`func (o *Ipv6fixedaddresstemplate) HasUseDomainName() bool`

HasUseDomainName returns a boolean if a field has been set.

### GetUseDomainNameServers

`func (o *Ipv6fixedaddresstemplate) GetUseDomainNameServers() bool`

GetUseDomainNameServers returns the UseDomainNameServers field if non-nil, zero value otherwise.

### GetUseDomainNameServersOk

`func (o *Ipv6fixedaddresstemplate) GetUseDomainNameServersOk() (*bool, bool)`

GetUseDomainNameServersOk returns a tuple with the UseDomainNameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDomainNameServers

`func (o *Ipv6fixedaddresstemplate) SetUseDomainNameServers(v bool)`

SetUseDomainNameServers sets UseDomainNameServers field to given value.

### HasUseDomainNameServers

`func (o *Ipv6fixedaddresstemplate) HasUseDomainNameServers() bool`

HasUseDomainNameServers returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Ipv6fixedaddresstemplate) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Ipv6fixedaddresstemplate) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Ipv6fixedaddresstemplate) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Ipv6fixedaddresstemplate) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseOptions

`func (o *Ipv6fixedaddresstemplate) GetUseOptions() bool`

GetUseOptions returns the UseOptions field if non-nil, zero value otherwise.

### GetUseOptionsOk

`func (o *Ipv6fixedaddresstemplate) GetUseOptionsOk() (*bool, bool)`

GetUseOptionsOk returns a tuple with the UseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOptions

`func (o *Ipv6fixedaddresstemplate) SetUseOptions(v bool)`

SetUseOptions sets UseOptions field to given value.

### HasUseOptions

`func (o *Ipv6fixedaddresstemplate) HasUseOptions() bool`

HasUseOptions returns a boolean if a field has been set.

### GetUsePreferredLifetime

`func (o *Ipv6fixedaddresstemplate) GetUsePreferredLifetime() bool`

GetUsePreferredLifetime returns the UsePreferredLifetime field if non-nil, zero value otherwise.

### GetUsePreferredLifetimeOk

`func (o *Ipv6fixedaddresstemplate) GetUsePreferredLifetimeOk() (*bool, bool)`

GetUsePreferredLifetimeOk returns a tuple with the UsePreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePreferredLifetime

`func (o *Ipv6fixedaddresstemplate) SetUsePreferredLifetime(v bool)`

SetUsePreferredLifetime sets UsePreferredLifetime field to given value.

### HasUsePreferredLifetime

`func (o *Ipv6fixedaddresstemplate) HasUsePreferredLifetime() bool`

HasUsePreferredLifetime returns a boolean if a field has been set.

### GetUseValidLifetime

`func (o *Ipv6fixedaddresstemplate) GetUseValidLifetime() bool`

GetUseValidLifetime returns the UseValidLifetime field if non-nil, zero value otherwise.

### GetUseValidLifetimeOk

`func (o *Ipv6fixedaddresstemplate) GetUseValidLifetimeOk() (*bool, bool)`

GetUseValidLifetimeOk returns a tuple with the UseValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseValidLifetime

`func (o *Ipv6fixedaddresstemplate) SetUseValidLifetime(v bool)`

SetUseValidLifetime sets UseValidLifetime field to given value.

### HasUseValidLifetime

`func (o *Ipv6fixedaddresstemplate) HasUseValidLifetime() bool`

HasUseValidLifetime returns a boolean if a field has been set.

### GetValidLifetime

`func (o *Ipv6fixedaddresstemplate) GetValidLifetime() int64`

GetValidLifetime returns the ValidLifetime field if non-nil, zero value otherwise.

### GetValidLifetimeOk

`func (o *Ipv6fixedaddresstemplate) GetValidLifetimeOk() (*int64, bool)`

GetValidLifetimeOk returns a tuple with the ValidLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidLifetime

`func (o *Ipv6fixedaddresstemplate) SetValidLifetime(v int64)`

SetValidLifetime sets ValidLifetime field to given value.

### HasValidLifetime

`func (o *Ipv6fixedaddresstemplate) HasValidLifetime() bool`

HasValidLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


