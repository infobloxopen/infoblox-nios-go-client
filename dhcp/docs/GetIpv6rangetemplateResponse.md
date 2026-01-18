# GetIpv6rangetemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
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
**Result** | Pointer to [**Ipv6rangetemplate**](Ipv6rangetemplate.md) |  | [optional] 

## Methods

### NewGetIpv6rangetemplateResponse

`func NewGetIpv6rangetemplateResponse() *GetIpv6rangetemplateResponse`

NewGetIpv6rangetemplateResponse instantiates a new GetIpv6rangetemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpv6rangetemplateResponseWithDefaults

`func NewGetIpv6rangetemplateResponseWithDefaults() *GetIpv6rangetemplateResponse`

NewGetIpv6rangetemplateResponseWithDefaults instantiates a new GetIpv6rangetemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetIpv6rangetemplateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetIpv6rangetemplateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetIpv6rangetemplateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetIpv6rangetemplateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudApiCompatible

`func (o *GetIpv6rangetemplateResponse) GetCloudApiCompatible() bool`

GetCloudApiCompatible returns the CloudApiCompatible field if non-nil, zero value otherwise.

### GetCloudApiCompatibleOk

`func (o *GetIpv6rangetemplateResponse) GetCloudApiCompatibleOk() (*bool, bool)`

GetCloudApiCompatibleOk returns a tuple with the CloudApiCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudApiCompatible

`func (o *GetIpv6rangetemplateResponse) SetCloudApiCompatible(v bool)`

SetCloudApiCompatible sets CloudApiCompatible field to given value.

### HasCloudApiCompatible

`func (o *GetIpv6rangetemplateResponse) HasCloudApiCompatible() bool`

HasCloudApiCompatible returns a boolean if a field has been set.

### GetComment

`func (o *GetIpv6rangetemplateResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetIpv6rangetemplateResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetIpv6rangetemplateResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetIpv6rangetemplateResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelegatedMember

`func (o *GetIpv6rangetemplateResponse) GetDelegatedMember() Ipv6rangetemplateDelegatedMember`

GetDelegatedMember returns the DelegatedMember field if non-nil, zero value otherwise.

### GetDelegatedMemberOk

`func (o *GetIpv6rangetemplateResponse) GetDelegatedMemberOk() (*Ipv6rangetemplateDelegatedMember, bool)`

GetDelegatedMemberOk returns a tuple with the DelegatedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedMember

`func (o *GetIpv6rangetemplateResponse) SetDelegatedMember(v Ipv6rangetemplateDelegatedMember)`

SetDelegatedMember sets DelegatedMember field to given value.

### HasDelegatedMember

`func (o *GetIpv6rangetemplateResponse) HasDelegatedMember() bool`

HasDelegatedMember returns a boolean if a field has been set.

### GetExclude

`func (o *GetIpv6rangetemplateResponse) GetExclude() []Ipv6rangetemplateExclude`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *GetIpv6rangetemplateResponse) GetExcludeOk() (*[]Ipv6rangetemplateExclude, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *GetIpv6rangetemplateResponse) SetExclude(v []Ipv6rangetemplateExclude)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *GetIpv6rangetemplateResponse) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetLogicFilterRules

`func (o *GetIpv6rangetemplateResponse) GetLogicFilterRules() []Ipv6rangetemplateLogicFilterRules`

GetLogicFilterRules returns the LogicFilterRules field if non-nil, zero value otherwise.

### GetLogicFilterRulesOk

`func (o *GetIpv6rangetemplateResponse) GetLogicFilterRulesOk() (*[]Ipv6rangetemplateLogicFilterRules, bool)`

GetLogicFilterRulesOk returns a tuple with the LogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicFilterRules

`func (o *GetIpv6rangetemplateResponse) SetLogicFilterRules(v []Ipv6rangetemplateLogicFilterRules)`

SetLogicFilterRules sets LogicFilterRules field to given value.

### HasLogicFilterRules

`func (o *GetIpv6rangetemplateResponse) HasLogicFilterRules() bool`

HasLogicFilterRules returns a boolean if a field has been set.

### GetMember

`func (o *GetIpv6rangetemplateResponse) GetMember() Ipv6rangetemplateMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetIpv6rangetemplateResponse) GetMemberOk() (*Ipv6rangetemplateMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetIpv6rangetemplateResponse) SetMember(v Ipv6rangetemplateMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetIpv6rangetemplateResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetName

`func (o *GetIpv6rangetemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpv6rangetemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpv6rangetemplateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIpv6rangetemplateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfAddresses

`func (o *GetIpv6rangetemplateResponse) GetNumberOfAddresses() int64`

GetNumberOfAddresses returns the NumberOfAddresses field if non-nil, zero value otherwise.

### GetNumberOfAddressesOk

`func (o *GetIpv6rangetemplateResponse) GetNumberOfAddressesOk() (*int64, bool)`

GetNumberOfAddressesOk returns a tuple with the NumberOfAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAddresses

`func (o *GetIpv6rangetemplateResponse) SetNumberOfAddresses(v int64)`

SetNumberOfAddresses sets NumberOfAddresses field to given value.

### HasNumberOfAddresses

`func (o *GetIpv6rangetemplateResponse) HasNumberOfAddresses() bool`

HasNumberOfAddresses returns a boolean if a field has been set.

### GetOffset

`func (o *GetIpv6rangetemplateResponse) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetIpv6rangetemplateResponse) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetIpv6rangetemplateResponse) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetIpv6rangetemplateResponse) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOptionFilterRules

`func (o *GetIpv6rangetemplateResponse) GetOptionFilterRules() []Ipv6rangetemplateOptionFilterRules`

GetOptionFilterRules returns the OptionFilterRules field if non-nil, zero value otherwise.

### GetOptionFilterRulesOk

`func (o *GetIpv6rangetemplateResponse) GetOptionFilterRulesOk() (*[]Ipv6rangetemplateOptionFilterRules, bool)`

GetOptionFilterRulesOk returns a tuple with the OptionFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionFilterRules

`func (o *GetIpv6rangetemplateResponse) SetOptionFilterRules(v []Ipv6rangetemplateOptionFilterRules)`

SetOptionFilterRules sets OptionFilterRules field to given value.

### HasOptionFilterRules

`func (o *GetIpv6rangetemplateResponse) HasOptionFilterRules() bool`

HasOptionFilterRules returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *GetIpv6rangetemplateResponse) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *GetIpv6rangetemplateResponse) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *GetIpv6rangetemplateResponse) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *GetIpv6rangetemplateResponse) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetServerAssociationType

`func (o *GetIpv6rangetemplateResponse) GetServerAssociationType() string`

GetServerAssociationType returns the ServerAssociationType field if non-nil, zero value otherwise.

### GetServerAssociationTypeOk

`func (o *GetIpv6rangetemplateResponse) GetServerAssociationTypeOk() (*string, bool)`

GetServerAssociationTypeOk returns a tuple with the ServerAssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAssociationType

`func (o *GetIpv6rangetemplateResponse) SetServerAssociationType(v string)`

SetServerAssociationType sets ServerAssociationType field to given value.

### HasServerAssociationType

`func (o *GetIpv6rangetemplateResponse) HasServerAssociationType() bool`

HasServerAssociationType returns a boolean if a field has been set.

### GetUseLogicFilterRules

`func (o *GetIpv6rangetemplateResponse) GetUseLogicFilterRules() bool`

GetUseLogicFilterRules returns the UseLogicFilterRules field if non-nil, zero value otherwise.

### GetUseLogicFilterRulesOk

`func (o *GetIpv6rangetemplateResponse) GetUseLogicFilterRulesOk() (*bool, bool)`

GetUseLogicFilterRulesOk returns a tuple with the UseLogicFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogicFilterRules

`func (o *GetIpv6rangetemplateResponse) SetUseLogicFilterRules(v bool)`

SetUseLogicFilterRules sets UseLogicFilterRules field to given value.

### HasUseLogicFilterRules

`func (o *GetIpv6rangetemplateResponse) HasUseLogicFilterRules() bool`

HasUseLogicFilterRules returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *GetIpv6rangetemplateResponse) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *GetIpv6rangetemplateResponse) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *GetIpv6rangetemplateResponse) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *GetIpv6rangetemplateResponse) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetResult

`func (o *GetIpv6rangetemplateResponse) GetResult() Ipv6rangetemplate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetIpv6rangetemplateResponse) GetResultOk() (*Ipv6rangetemplate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetIpv6rangetemplateResponse) SetResult(v Ipv6rangetemplate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetIpv6rangetemplateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


