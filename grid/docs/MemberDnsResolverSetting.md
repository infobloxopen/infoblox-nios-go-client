# MemberDnsResolverSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resolvers** | Pointer to **[]string** | The resolvers of a Grid member. The Grid member sends queries to the first name server address in the list. The second name server address is used if first one does not response. | [optional] 
**SearchDomains** | Pointer to **[]string** | The Search Domain Group, which is a group of domain names that the Infoblox device can add to partial queries that do not specify a domain name. Note that you can set this parameter only when prefer_resolver or alternate_resolver is set. | [optional] 

## Methods

### NewMemberDnsResolverSetting

`func NewMemberDnsResolverSetting() *MemberDnsResolverSetting`

NewMemberDnsResolverSetting instantiates a new MemberDnsResolverSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsResolverSettingWithDefaults

`func NewMemberDnsResolverSettingWithDefaults() *MemberDnsResolverSetting`

NewMemberDnsResolverSettingWithDefaults instantiates a new MemberDnsResolverSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolvers

`func (o *MemberDnsResolverSetting) GetResolvers() []string`

GetResolvers returns the Resolvers field if non-nil, zero value otherwise.

### GetResolversOk

`func (o *MemberDnsResolverSetting) GetResolversOk() (*[]string, bool)`

GetResolversOk returns a tuple with the Resolvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvers

`func (o *MemberDnsResolverSetting) SetResolvers(v []string)`

SetResolvers sets Resolvers field to given value.

### HasResolvers

`func (o *MemberDnsResolverSetting) HasResolvers() bool`

HasResolvers returns a boolean if a field has been set.

### GetSearchDomains

`func (o *MemberDnsResolverSetting) GetSearchDomains() []string`

GetSearchDomains returns the SearchDomains field if non-nil, zero value otherwise.

### GetSearchDomainsOk

`func (o *MemberDnsResolverSetting) GetSearchDomainsOk() (*[]string, bool)`

GetSearchDomainsOk returns a tuple with the SearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomains

`func (o *MemberDnsResolverSetting) SetSearchDomains(v []string)`

SetSearchDomains sets SearchDomains field to given value.

### HasSearchDomains

`func (o *MemberDnsResolverSetting) HasSearchDomains() bool`

HasSearchDomains returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


