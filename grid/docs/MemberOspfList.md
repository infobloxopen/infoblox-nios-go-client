# MemberOspfList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaId** | Pointer to **string** | The area ID value of the OSPF settings. | [optional] 
**AreaType** | Pointer to **string** | The OSPF area type. | [optional] 
**AuthenticationKey** | Pointer to **string** | The authentication password to use for OSPF. The authentication key is valid only when authentication type is \&quot;SIMPLE\&quot; or \&quot;MESSAGE_DIGEST\&quot;. | [optional] 
**AuthenticationType** | Pointer to **string** | The authentication type used for the OSPF advertisement. | [optional] 
**AutoCalcCostEnabled** | Pointer to **bool** | Determines if auto calculate cost is enabled or not. | [optional] 
**Comment** | Pointer to **string** | A descriptive comment of the OSPF configuration. | [optional] 
**Cost** | Pointer to **int64** | The cost metric associated with the OSPF advertisement. | [optional] 
**DeadInterval** | Pointer to **int64** | The dead interval value of OSPF (in seconds). The dead interval describes the time to wait before declaring the device is unavailable and down. | [optional] 
**HelloInterval** | Pointer to **int64** | The hello interval value of OSPF. The hello interval specifies how often to send OSPF hello advertisement, in seconds. | [optional] 
**Interface** | Pointer to **string** | The interface that sends out OSPF advertisement information. | [optional] 
**IsIpv4** | Pointer to **bool** | The OSPF protocol version. Specify \&quot;true\&quot; if the IPv4 version of OSPF is used, or \&quot;false\&quot; if the IPv6 version of OSPF is used. | [optional] 
**KeyId** | Pointer to **int64** | The hash key identifier to use for \&quot;MESSAGE_DIGEST\&quot; authentication. The hash key identifier is valid only when authentication type is \&quot;MESSAGE_DIGEST\&quot;. | [optional] 
**RetransmitInterval** | Pointer to **int64** | The retransmit interval time of OSPF (in seconds). The retransmit interval describes the time to wait before retransmitting OSPF advertisement. | [optional] 
**TransmitDelay** | Pointer to **int64** | The transmit delay value of OSPF (in seconds). The transmit delay describes the time to wait before sending an advertisement. | [optional] 
**AdvertiseInterfaceVlan** | Pointer to **string** | The VLAN used as the advertising interface for sending OSPF announcements. | [optional] 
**BfdTemplate** | Pointer to **string** | Determines BFD template name. | [optional] 
**EnableBfd** | Pointer to **bool** | Determines if the BFD is enabled or not. | [optional] 
**EnableBfdDnscheck** | Pointer to **bool** | Determines if BFD internal DNS monitor is enabled or not. | [optional] 

## Methods

### NewMemberOspfList

`func NewMemberOspfList() *MemberOspfList`

NewMemberOspfList instantiates a new MemberOspfList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberOspfListWithDefaults

`func NewMemberOspfListWithDefaults() *MemberOspfList`

NewMemberOspfListWithDefaults instantiates a new MemberOspfList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaId

`func (o *MemberOspfList) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *MemberOspfList) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *MemberOspfList) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.

### HasAreaId

`func (o *MemberOspfList) HasAreaId() bool`

HasAreaId returns a boolean if a field has been set.

### GetAreaType

`func (o *MemberOspfList) GetAreaType() string`

GetAreaType returns the AreaType field if non-nil, zero value otherwise.

### GetAreaTypeOk

`func (o *MemberOspfList) GetAreaTypeOk() (*string, bool)`

GetAreaTypeOk returns a tuple with the AreaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaType

`func (o *MemberOspfList) SetAreaType(v string)`

SetAreaType sets AreaType field to given value.

### HasAreaType

`func (o *MemberOspfList) HasAreaType() bool`

HasAreaType returns a boolean if a field has been set.

### GetAuthenticationKey

`func (o *MemberOspfList) GetAuthenticationKey() string`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *MemberOspfList) GetAuthenticationKeyOk() (*string, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *MemberOspfList) SetAuthenticationKey(v string)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *MemberOspfList) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetAuthenticationType

`func (o *MemberOspfList) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *MemberOspfList) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *MemberOspfList) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.

### HasAuthenticationType

`func (o *MemberOspfList) HasAuthenticationType() bool`

HasAuthenticationType returns a boolean if a field has been set.

### GetAutoCalcCostEnabled

`func (o *MemberOspfList) GetAutoCalcCostEnabled() bool`

GetAutoCalcCostEnabled returns the AutoCalcCostEnabled field if non-nil, zero value otherwise.

### GetAutoCalcCostEnabledOk

`func (o *MemberOspfList) GetAutoCalcCostEnabledOk() (*bool, bool)`

GetAutoCalcCostEnabledOk returns a tuple with the AutoCalcCostEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCalcCostEnabled

`func (o *MemberOspfList) SetAutoCalcCostEnabled(v bool)`

SetAutoCalcCostEnabled sets AutoCalcCostEnabled field to given value.

### HasAutoCalcCostEnabled

`func (o *MemberOspfList) HasAutoCalcCostEnabled() bool`

HasAutoCalcCostEnabled returns a boolean if a field has been set.

### GetComment

`func (o *MemberOspfList) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MemberOspfList) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MemberOspfList) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MemberOspfList) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCost

`func (o *MemberOspfList) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *MemberOspfList) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *MemberOspfList) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *MemberOspfList) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadInterval

`func (o *MemberOspfList) GetDeadInterval() int64`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *MemberOspfList) GetDeadIntervalOk() (*int64, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *MemberOspfList) SetDeadInterval(v int64)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *MemberOspfList) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetHelloInterval

`func (o *MemberOspfList) GetHelloInterval() int64`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *MemberOspfList) GetHelloIntervalOk() (*int64, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *MemberOspfList) SetHelloInterval(v int64)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *MemberOspfList) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetInterface

`func (o *MemberOspfList) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *MemberOspfList) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *MemberOspfList) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *MemberOspfList) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIsIpv4

`func (o *MemberOspfList) GetIsIpv4() bool`

GetIsIpv4 returns the IsIpv4 field if non-nil, zero value otherwise.

### GetIsIpv4Ok

`func (o *MemberOspfList) GetIsIpv4Ok() (*bool, bool)`

GetIsIpv4Ok returns a tuple with the IsIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIpv4

`func (o *MemberOspfList) SetIsIpv4(v bool)`

SetIsIpv4 sets IsIpv4 field to given value.

### HasIsIpv4

`func (o *MemberOspfList) HasIsIpv4() bool`

HasIsIpv4 returns a boolean if a field has been set.

### GetKeyId

`func (o *MemberOspfList) GetKeyId() int64`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *MemberOspfList) GetKeyIdOk() (*int64, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *MemberOspfList) SetKeyId(v int64)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *MemberOspfList) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetRetransmitInterval

`func (o *MemberOspfList) GetRetransmitInterval() int64`

GetRetransmitInterval returns the RetransmitInterval field if non-nil, zero value otherwise.

### GetRetransmitIntervalOk

`func (o *MemberOspfList) GetRetransmitIntervalOk() (*int64, bool)`

GetRetransmitIntervalOk returns a tuple with the RetransmitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitInterval

`func (o *MemberOspfList) SetRetransmitInterval(v int64)`

SetRetransmitInterval sets RetransmitInterval field to given value.

### HasRetransmitInterval

`func (o *MemberOspfList) HasRetransmitInterval() bool`

HasRetransmitInterval returns a boolean if a field has been set.

### GetTransmitDelay

`func (o *MemberOspfList) GetTransmitDelay() int64`

GetTransmitDelay returns the TransmitDelay field if non-nil, zero value otherwise.

### GetTransmitDelayOk

`func (o *MemberOspfList) GetTransmitDelayOk() (*int64, bool)`

GetTransmitDelayOk returns a tuple with the TransmitDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitDelay

`func (o *MemberOspfList) SetTransmitDelay(v int64)`

SetTransmitDelay sets TransmitDelay field to given value.

### HasTransmitDelay

`func (o *MemberOspfList) HasTransmitDelay() bool`

HasTransmitDelay returns a boolean if a field has been set.

### GetAdvertiseInterfaceVlan

`func (o *MemberOspfList) GetAdvertiseInterfaceVlan() string`

GetAdvertiseInterfaceVlan returns the AdvertiseInterfaceVlan field if non-nil, zero value otherwise.

### GetAdvertiseInterfaceVlanOk

`func (o *MemberOspfList) GetAdvertiseInterfaceVlanOk() (*string, bool)`

GetAdvertiseInterfaceVlanOk returns a tuple with the AdvertiseInterfaceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseInterfaceVlan

`func (o *MemberOspfList) SetAdvertiseInterfaceVlan(v string)`

SetAdvertiseInterfaceVlan sets AdvertiseInterfaceVlan field to given value.

### HasAdvertiseInterfaceVlan

`func (o *MemberOspfList) HasAdvertiseInterfaceVlan() bool`

HasAdvertiseInterfaceVlan returns a boolean if a field has been set.

### GetBfdTemplate

`func (o *MemberOspfList) GetBfdTemplate() string`

GetBfdTemplate returns the BfdTemplate field if non-nil, zero value otherwise.

### GetBfdTemplateOk

`func (o *MemberOspfList) GetBfdTemplateOk() (*string, bool)`

GetBfdTemplateOk returns a tuple with the BfdTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfdTemplate

`func (o *MemberOspfList) SetBfdTemplate(v string)`

SetBfdTemplate sets BfdTemplate field to given value.

### HasBfdTemplate

`func (o *MemberOspfList) HasBfdTemplate() bool`

HasBfdTemplate returns a boolean if a field has been set.

### GetEnableBfd

`func (o *MemberOspfList) GetEnableBfd() bool`

GetEnableBfd returns the EnableBfd field if non-nil, zero value otherwise.

### GetEnableBfdOk

`func (o *MemberOspfList) GetEnableBfdOk() (*bool, bool)`

GetEnableBfdOk returns a tuple with the EnableBfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBfd

`func (o *MemberOspfList) SetEnableBfd(v bool)`

SetEnableBfd sets EnableBfd field to given value.

### HasEnableBfd

`func (o *MemberOspfList) HasEnableBfd() bool`

HasEnableBfd returns a boolean if a field has been set.

### GetEnableBfdDnscheck

`func (o *MemberOspfList) GetEnableBfdDnscheck() bool`

GetEnableBfdDnscheck returns the EnableBfdDnscheck field if non-nil, zero value otherwise.

### GetEnableBfdDnscheckOk

`func (o *MemberOspfList) GetEnableBfdDnscheckOk() (*bool, bool)`

GetEnableBfdDnscheckOk returns a tuple with the EnableBfdDnscheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBfdDnscheck

`func (o *MemberOspfList) SetEnableBfdDnscheck(v bool)`

SetEnableBfdDnscheck sets EnableBfdDnscheck field to given value.

### HasEnableBfdDnscheck

`func (o *MemberOspfList) HasEnableBfdDnscheck() bool`

HasEnableBfdDnscheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


