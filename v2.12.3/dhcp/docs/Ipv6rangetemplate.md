# Ipv6rangetemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CloudApiCompatible** | Pointer to **bool** | Determines whether the IPv6 DHCP range template can be used to create network objects in a cloud-computing deployment. | [optional] 
**Comment** | Pointer to **string** | The IPv6 DHCP range template descriptive comment. | [optional] 
**DelegatedMember** | Pointer to [**Ipv6rangetemplateDelegatedMember**](Ipv6rangetemplateDelegatedMember.md) |  | [optional] 
**Exclude** | Pointer to [**[]Ipv6rangetemplateExclude**](Ipv6rangetemplateExclude.md) | These are ranges of IPv6 addresses that the appliance does not use to assign to clients. You can use these excluded addresses as static IPv6 addresses. They contain the start and end addresses of the excluded range, and optionally, information about this excluded range. | [optional] 
**LogicFilterRules** | Pointer to [**[]Ipv6rangetemplateLogicFilterRules**](Ipv6rangetemplateLogicFilterRules.md) | This field contains the logic filters to be applied on this IPv6 range. This list corresponds to the match rules that are written to the DHCPv6 configuration file. | [optional] 
**Member** | Pointer to [**Ipv6rangetemplateMember**](Ipv6rangetemplateMember.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the IPv6 DHCP range template. | [optional] 
**NumberOfAddresses** | Pointer to **int64** | The number of addresses for the IPv6 DHCP range. | [optional] 
**Offset** | Pointer to **int64** | The start address offset for the IPv6 DHCP range. | [optional] 
**OptionFilterRules** | Pointer to [**[]Ipv6rangetemplateOptionFilterRules**](Ipv6rangetemplateOptionFilterRules.md) | This field contains the Option filters to be applied to this IPv6 range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease. | [optional] 
**RecycleLeases** | Pointer to **bool** | Determines whether the leases are kept in Recycle Bin until one week after expiry. If this is set to False, the leases are permanently deleted. | [optional] 
**ServerAssociationType** | Pointer to **string** | The type of server that is going to serve the IPv6 DHCP range. | [optional] 
**UseLogicFilterRules** | Pointer to **bool** | Use flag for: logic_filter_rules | [optional] 
**UseRecycleLeases** | Pointer to **bool** | Use flag for: recycle_leases | [optional] 

## Methods

### NewIpv6rangetemplate

`func NewIpv6rangetemplate() *Ipv6rangetemplate`

NewIpv6rangetemplate instantiates a new Ipv6rangetemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6rangetemplateWithDefaults

`func NewIpv6rangetemplateWithDefaults() *Ipv6rangetemplate`

NewIpv6rangetemplateWithDefaults instantiates a new Ipv6rangetemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Ipv6rangetemplate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Ipv6rangetemplate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Ipv6rangetemplate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Ipv6rangetemplate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudApiCompatible

`func (o *Ipv6rangetemplate) GetCloudApiCompatible() bool`

GetCloudApiCompatible returns the CloudApiCompatible field if non-nil, zero value otherwise.

### GetCloudApiCompatibleOk

`func (o *Ipv6rangetemplate) GetCloudApiCompatibleOk() (*bool, bool)`

GetCloudApiCompatibleOk returns a tuple with the CloudApiCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudApiCompatible

`func (o *Ipv6rangetemplate) SetCloudApiCompatible(v bool)`

SetCloudApiCompatible sets CloudApiCompatible field to given value.

### HasCloudApiCompatible

`func (o *Ipv6rangetemplate) HasCloudApiCompatible() bool`

HasCloudApiCompatible returns a boolean if a field has been set.

### GetComment

`func (o *Ipv6rangetemplate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Ipv6rangetemplate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Ipv6rangetemplate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Ipv6rangetemplate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelegatedMember

`func (o *Ipv6rangetemplate) GetDelegatedMember() Ipv6rangetemplateDelegatedMember`

GetDelegatedMember returns the DelegatedMember field if non-nil, zero value otherwise.

### GetDelegatedMemberOk

`func (o *Ipv6rangetemplate) GetDelegatedMemberOk() (*Ipv6rangetemplateDelegatedMember, bool)`

GetDelegatedMemberOk returns a tuple with the DelegatedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedMember

`func (o *Ipv6rangetemplate) SetDelegatedMember(v Ipv6rangetemplateDelegatedMember)`

SetDelegatedMember sets DelegatedMember field to given value.

### HasDelegatedMember

`func (o *Ipv6rangetemplate) HasDelegatedMember() bool`

HasDelegatedMember returns a boolean if a field has been set.

### GetExclude

`func (o *Ipv6rangetemplate) GetExclude() []Ipv6rangetemplateExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *Ipv6rangetemplate) GetExcludeOk() (*[]Ipv6rangetemplateExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *Ipv6rangetemplate) SetExclude(v []Ipv6rangetemplateExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *Ipv6rangetemplate) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *Ipv6rangetemplate) GetLogicFilterRules() []Ipv6rangetemplateLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *Ipv6rangetemplate) GetLogicFilterRulesOk() (*[]Ipv6rangetemplateLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *Ipv6rangetemplate) SetLogicFilterRules(v []Ipv6rangetemplateLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *Ipv6rangetemplate) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMember

`func (o *Ipv6rangetemplate) GetMember() Ipv6rangetemplateMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Ipv6rangetemplate) GetMemberOk() (*Ipv6rangetemplateMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Ipv6rangetemplate) SetMember(v Ipv6rangetemplateMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *Ipv6rangetemplate) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *Ipv6rangetemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ipv6rangetemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ipv6rangetemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ipv6rangetemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfAddresses

`func (o *Ipv6rangetemplate) GetNumberOfAddresses() int64`

GetNumberOfAddresses returns the NumberOfAddresses field if non-nil, zero value otherwise.

### GetNumberOfAddressesOk

`func (o *Ipv6rangetemplate) GetNumberOfAddressesOk() (*int64, bool)`

GetNumberOfAddressesOk returns a tuple with the NumberOfAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAddresses

`func (o *Ipv6rangetemplate) SetNumberOfAddresses(v int64)`

SetNumberOfAddresses sets NumberOfAddresses field to given value.

### HasNumberOfAddresses

`func (o *Ipv6rangetemplate) HasNumberOfAddresses() bool`

HasNumberOfAddresses returns a boolean if a field has been set.

### GetOffset

`func (o *Ipv6rangetemplate) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Ipv6rangetemplate) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Ipv6rangetemplate) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *Ipv6rangetemplate) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOptionFilterRules

`func (o *Ipv6rangetemplate) GetOptionFilterRules() []Ipv6rangetemplateOptionFilterRules`

GetOptionFilterRules returns the OptionFilterRules field if non-nil, zero value otherwise.

### GetOptionFilterRulesOk

`func (o *Ipv6rangetemplate) GetOptionFilterRulesOk() (*[]Ipv6rangetemplateOptionFilterRules, bool)`

GetOptionFilterRulesOk returns a tuple with the OptionFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFilterRules

`func (o *Ipv6rangetemplate) SetOptionFilterRules(v []Ipv6rangetemplateOptionFilterRules)`

SetOptionFilterRules sets OptionFilterRules field to given value.

### HasOptionFilterRules

`func (o *Ipv6rangetemplate) HasOptionFilterRules() bool`

HasOptionFilterRules returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *Ipv6rangetemplate) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *Ipv6rangetemplate) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *Ipv6rangetemplate) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *Ipv6rangetemplate) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetServerAssociationType

`func (o *Ipv6rangetemplate) GetServerAssociationType() string`

GetServerAssociationType returns the ServerAssociationType field if non-nil, zero value otherwise.

### GetServerAssociationTypeOk

`func (o *Ipv6rangetemplate) GetServerAssociationTypeOk() (*string, bool)`

GetServerAssociationTypeOk returns a tuple with the ServerAssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAssociationType

`func (o *Ipv6rangetemplate) SetServerAssociationType(v string)`

SetServerAssociationType sets ServerAssociationType field to given value.

### HasServerAssociationType

`func (o *Ipv6rangetemplate) HasServerAssociationType() bool`

HasServerAssociationType returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *Ipv6rangetemplate) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *Ipv6rangetemplate) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *Ipv6rangetemplate) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *Ipv6rangetemplate) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *Ipv6rangetemplate) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *Ipv6rangetemplate) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *Ipv6rangetemplate) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *Ipv6rangetemplate) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


