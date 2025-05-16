# MemberThreatprotectionNatRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleType** | Pointer to **string** | The rule type for the threat protection NAT mapping rule. | [optional] 
**Address** | Pointer to **string** | The IP address for the threat protection NAT mapping rule. | [optional] 
**Network** | Pointer to **string** | The network address for the threat protection NAT mapping rule. | [optional] 
**Cidr** | Pointer to **int64** | The network CIDR for the threat protection NAT mapping rule. | [optional] 
**StartAddress** | Pointer to **string** | The start address for the range of the threat protection NAT mapping rule. | [optional] 
**EndAddress** | Pointer to **string** | The end address for the range of the threat protection NAT mapping rule. | [optional] 
**NatPorts** | Pointer to [**MemberthreatprotectionnatrulesNatPorts**](MemberthreatprotectionnatrulesNatPorts.md) |  | [optional] 

## Methods

### NewMemberThreatprotectionNatRules

`func NewMemberThreatprotectionNatRules() *MemberThreatprotectionNatRules`

NewMemberThreatprotectionNatRules instantiates a new MemberThreatprotectionNatRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberThreatprotectionNatRulesWithDefaults

`func NewMemberThreatprotectionNatRulesWithDefaults() *MemberThreatprotectionNatRules`

NewMemberThreatprotectionNatRulesWithDefaults instantiates a new MemberThreatprotectionNatRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleType

`func (o *MemberThreatprotectionNatRules) GetRuleType() string`

GetRuleType returns the RuleType field if non-nil, zero value otherwise.

### GetRuleTypeOk

`func (o *MemberThreatprotectionNatRules) GetRuleTypeOk() (*string, bool)`

GetRuleTypeOk returns a tuple with the RuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleType

`func (o *MemberThreatprotectionNatRules) SetRuleType(v string)`

SetRuleType sets RuleType field to given value.

### HasRuleType

`func (o *MemberThreatprotectionNatRules) HasRuleType() bool`

HasRuleType returns a boolean if a field has been set.

### GetAddress

`func (o *MemberThreatprotectionNatRules) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberThreatprotectionNatRules) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberThreatprotectionNatRules) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberThreatprotectionNatRules) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNetwork

`func (o *MemberThreatprotectionNatRules) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MemberThreatprotectionNatRules) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MemberThreatprotectionNatRules) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MemberThreatprotectionNatRules) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCidr

`func (o *MemberThreatprotectionNatRules) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *MemberThreatprotectionNatRules) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *MemberThreatprotectionNatRules) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *MemberThreatprotectionNatRules) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetStartAddress

`func (o *MemberThreatprotectionNatRules) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *MemberThreatprotectionNatRules) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *MemberThreatprotectionNatRules) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *MemberThreatprotectionNatRules) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *MemberThreatprotectionNatRules) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *MemberThreatprotectionNatRules) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *MemberThreatprotectionNatRules) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *MemberThreatprotectionNatRules) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetNatPorts

`func (o *MemberThreatprotectionNatRules) GetNatPorts() MemberthreatprotectionnatrulesNatPorts`

GetNatPorts returns the NatPorts field if non-nil, zero value otherwise.

### GetNatPortsOk

`func (o *MemberThreatprotectionNatRules) GetNatPortsOk() (*MemberthreatprotectionnatrulesNatPorts, bool)`

GetNatPortsOk returns a tuple with the NatPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatPorts

`func (o *MemberThreatprotectionNatRules) SetNatPorts(v MemberthreatprotectionnatrulesNatPorts)`

SetNatPorts sets NatPorts field to given value.

### HasNatPorts

`func (o *MemberThreatprotectionNatRules) HasNatPorts() bool`

HasNatPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


