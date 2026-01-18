# ParentalcontrolSubscribersite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Abss** | Pointer to [**[]ParentalcontrolSubscribersiteAbss**](ParentalcontrolSubscribersiteAbss.md) | The list of ABS for the site. | [optional] 
**ApiMembers** | Pointer to [**[]ParentalcontrolSubscribersiteApiMembers**](ParentalcontrolSubscribersiteApiMembers.md) | The list of API members for the site. | [optional] 
**ApiPort** | Pointer to **int64** | The port number for gRPC API server. | [optional] [readonly] 
**BlockSize** | Pointer to **int64** | The size of the Deterministic NAT block-size. | [optional] 
**BlockingIpv4Vip1** | Pointer to **string** | The IPv4 Address of the blocking server. | [optional] 
**BlockingIpv4Vip2** | Pointer to **string** | The IPv4 Address of the blocking server. | [optional] 
**BlockingIpv6Vip1** | Pointer to **string** | The IPv6 Address of the blocking server. | [optional] 
**BlockingIpv6Vip2** | Pointer to **string** | The IPv6 Address of the blocking server. | [optional] 
**Comment** | Pointer to **string** | The human readable comment for the site. | [optional] 
**DcaSubBwList** | Pointer to **bool** | Enable/disable the DCA subscriber B/W list support. | [optional] 
**DcaSubQueryCount** | Pointer to **bool** | Enable/disable the DCA subscriber query count. | [optional] 
**EnableGlobalAllowListRpz** | Pointer to **bool** | Enable/disable global allow list RPZ setting. | [optional] 
**EnableRpzFilteringBypass** | Pointer to **bool** | Enable/disable Subscriber Secure Policy Bypass for Allowed list. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FirstPort** | Pointer to **int64** | The start of the first Deterministic block. | [optional] 
**GlobalAllowListRpz** | Pointer to **int64** | Global allow list RPZ index. Valid values are between 0 and 63. | [optional] 
**MaximumSubscribers** | Pointer to **int64** | The max number of subscribers for the site. It is used to configure the cache size. | [optional] 
**Members** | Pointer to [**[]ParentalcontrolSubscribersiteMembers**](ParentalcontrolSubscribersiteMembers.md) | The list of members for the site. | [optional] 
**Msps** | Pointer to [**[]ParentalcontrolSubscribersiteMsps**](ParentalcontrolSubscribersiteMsps.md) | The list of MSP for the site. | [optional] 
**Name** | Pointer to **string** | The name of the site. | [optional] 
**NasGateways** | Pointer to [**[]ParentalcontrolSubscribersiteNasGateways**](ParentalcontrolSubscribersiteNasGateways.md) | The list of accounting log servers. | [optional] 
**NasPort** | Pointer to **int64** | The port number to reach the collector. | [optional] 
**ProxyRpzPassthru** | Pointer to **bool** | Enforce the global proxy list. | [optional] 
**Spms** | Pointer to [**[]ParentalcontrolSubscribersiteSpms**](ParentalcontrolSubscribersiteSpms.md) | The list of SPM for the site. | [optional] 
**StopAnycast** | Pointer to **bool** | Stop the anycast service when the subscriber service is in the interim state. | [optional] 
**StrictNat** | Pointer to **bool** | Restrict subscriber cache entries to NATed clients. | [optional] 
**SubscriberCollectionType** | Pointer to **string** | Subscriber collection type either RADIUS or API. | [optional] 

## Methods

### NewParentalcontrolSubscribersite

`func NewParentalcontrolSubscribersite() *ParentalcontrolSubscribersite`

NewParentalcontrolSubscribersite instantiates a new ParentalcontrolSubscribersite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentalcontrolSubscribersiteWithDefaults

`func NewParentalcontrolSubscribersiteWithDefaults() *ParentalcontrolSubscribersite`

NewParentalcontrolSubscribersiteWithDefaults instantiates a new ParentalcontrolSubscribersite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ParentalcontrolSubscribersite) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ParentalcontrolSubscribersite) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ParentalcontrolSubscribersite) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ParentalcontrolSubscribersite) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAbss

`func (o *ParentalcontrolSubscribersite) GetAbss() []ParentalcontrolSubscribersiteAbss`

GetAbss returns the Abss field if non-nil, zero value otherwise.

### GetAbssOk

`func (o *ParentalcontrolSubscribersite) GetAbssOk() (*[]ParentalcontrolSubscribersiteAbss, bool)`

GetAbssOk returns a tuple with the Abss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbss

`func (o *ParentalcontrolSubscribersite) SetAbss(v []ParentalcontrolSubscribersiteAbss)`

SetAbss sets Abss field to given value.

### HasAbss

`func (o *ParentalcontrolSubscribersite) HasAbss() bool`

HasAbss returns a boolean if a field has been set.

### GetApiMembers

`func (o *ParentalcontrolSubscribersite) GetApiMembers() []ParentalcontrolSubscribersiteApiMembers`

GetApiMembers returns the ApiMembers field if non-nil, zero value otherwise.

### GetApiMembersOk

`func (o *ParentalcontrolSubscribersite) GetApiMembersOk() (*[]ParentalcontrolSubscribersiteApiMembers, bool)`

GetApiMembersOk returns a tuple with the ApiMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMembers

`func (o *ParentalcontrolSubscribersite) SetApiMembers(v []ParentalcontrolSubscribersiteApiMembers)`

SetApiMembers sets ApiMembers field to given value.

### HasApiMembers

`func (o *ParentalcontrolSubscribersite) HasApiMembers() bool`

HasApiMembers returns a boolean if a field has been set.

### GetApiPort

`func (o *ParentalcontrolSubscribersite) GetApiPort() int64`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *ParentalcontrolSubscribersite) GetApiPortOk() (*int64, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *ParentalcontrolSubscribersite) SetApiPort(v int64)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *ParentalcontrolSubscribersite) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetBlockSize

`func (o *ParentalcontrolSubscribersite) GetBlockSize() int64`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *ParentalcontrolSubscribersite) GetBlockSizeOk() (*int64, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *ParentalcontrolSubscribersite) SetBlockSize(v int64)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *ParentalcontrolSubscribersite) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBlockingIpv4Vip1

`func (o *ParentalcontrolSubscribersite) GetBlockingIpv4Vip1() string`

GetBlockingIpv4Vip1 returns the BlockingIpv4Vip1 field if non-nil, zero value otherwise.

### GetBlockingIpv4Vip1Ok

`func (o *ParentalcontrolSubscribersite) GetBlockingIpv4Vip1Ok() (*string, bool)`

GetBlockingIpv4Vip1Ok returns a tuple with the BlockingIpv4Vip1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv4Vip1

`func (o *ParentalcontrolSubscribersite) SetBlockingIpv4Vip1(v string)`

SetBlockingIpv4Vip1 sets BlockingIpv4Vip1 field to given value.

### HasBlockingIpv4Vip1

`func (o *ParentalcontrolSubscribersite) HasBlockingIpv4Vip1() bool`

HasBlockingIpv4Vip1 returns a boolean if a field has been set.

### GetBlockingIpv4Vip2

`func (o *ParentalcontrolSubscribersite) GetBlockingIpv4Vip2() string`

GetBlockingIpv4Vip2 returns the BlockingIpv4Vip2 field if non-nil, zero value otherwise.

### GetBlockingIpv4Vip2Ok

`func (o *ParentalcontrolSubscribersite) GetBlockingIpv4Vip2Ok() (*string, bool)`

GetBlockingIpv4Vip2Ok returns a tuple with the BlockingIpv4Vip2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv4Vip2

`func (o *ParentalcontrolSubscribersite) SetBlockingIpv4Vip2(v string)`

SetBlockingIpv4Vip2 sets BlockingIpv4Vip2 field to given value.

### HasBlockingIpv4Vip2

`func (o *ParentalcontrolSubscribersite) HasBlockingIpv4Vip2() bool`

HasBlockingIpv4Vip2 returns a boolean if a field has been set.

### GetBlockingIpv6Vip1

`func (o *ParentalcontrolSubscribersite) GetBlockingIpv6Vip1() string`

GetBlockingIpv6Vip1 returns the BlockingIpv6Vip1 field if non-nil, zero value otherwise.

### GetBlockingIpv6Vip1Ok

`func (o *ParentalcontrolSubscribersite) GetBlockingIpv6Vip1Ok() (*string, bool)`

GetBlockingIpv6Vip1Ok returns a tuple with the BlockingIpv6Vip1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv6Vip1

`func (o *ParentalcontrolSubscribersite) SetBlockingIpv6Vip1(v string)`

SetBlockingIpv6Vip1 sets BlockingIpv6Vip1 field to given value.

### HasBlockingIpv6Vip1

`func (o *ParentalcontrolSubscribersite) HasBlockingIpv6Vip1() bool`

HasBlockingIpv6Vip1 returns a boolean if a field has been set.

### GetBlockingIpv6Vip2

`func (o *ParentalcontrolSubscribersite) GetBlockingIpv6Vip2() string`

GetBlockingIpv6Vip2 returns the BlockingIpv6Vip2 field if non-nil, zero value otherwise.

### GetBlockingIpv6Vip2Ok

`func (o *ParentalcontrolSubscribersite) GetBlockingIpv6Vip2Ok() (*string, bool)`

GetBlockingIpv6Vip2Ok returns a tuple with the BlockingIpv6Vip2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv6Vip2

`func (o *ParentalcontrolSubscribersite) SetBlockingIpv6Vip2(v string)`

SetBlockingIpv6Vip2 sets BlockingIpv6Vip2 field to given value.

### HasBlockingIpv6Vip2

`func (o *ParentalcontrolSubscribersite) HasBlockingIpv6Vip2() bool`

HasBlockingIpv6Vip2 returns a boolean if a field has been set.

### GetComment

`func (o *ParentalcontrolSubscribersite) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ParentalcontrolSubscribersite) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ParentalcontrolSubscribersite) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ParentalcontrolSubscribersite) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDcaSubBwList

`func (o *ParentalcontrolSubscribersite) GetDcaSubBwList() bool`

GetDcaSubBwList returns the DcaSubBwList field if non-nil, zero value otherwise.

### GetDcaSubBwListOk

`func (o *ParentalcontrolSubscribersite) GetDcaSubBwListOk() (*bool, bool)`

GetDcaSubBwListOk returns a tuple with the DcaSubBwList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcaSubBwList

`func (o *ParentalcontrolSubscribersite) SetDcaSubBwList(v bool)`

SetDcaSubBwList sets DcaSubBwList field to given value.

### HasDcaSubBwList

`func (o *ParentalcontrolSubscribersite) HasDcaSubBwList() bool`

HasDcaSubBwList returns a boolean if a field has been set.

### GetDcaSubQueryCount

`func (o *ParentalcontrolSubscribersite) GetDcaSubQueryCount() bool`

GetDcaSubQueryCount returns the DcaSubQueryCount field if non-nil, zero value otherwise.

### GetDcaSubQueryCountOk

`func (o *ParentalcontrolSubscribersite) GetDcaSubQueryCountOk() (*bool, bool)`

GetDcaSubQueryCountOk returns a tuple with the DcaSubQueryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcaSubQueryCount

`func (o *ParentalcontrolSubscribersite) SetDcaSubQueryCount(v bool)`

SetDcaSubQueryCount sets DcaSubQueryCount field to given value.

### HasDcaSubQueryCount

`func (o *ParentalcontrolSubscribersite) HasDcaSubQueryCount() bool`

HasDcaSubQueryCount returns a boolean if a field has been set.

### GetEnableGlobalAllowListRpz

`func (o *ParentalcontrolSubscribersite) GetEnableGlobalAllowListRpz() bool`

GetEnableGlobalAllowListRpz returns the EnableGlobalAllowListRpz field if non-nil, zero value otherwise.

### GetEnableGlobalAllowListRpzOk

`func (o *ParentalcontrolSubscribersite) GetEnableGlobalAllowListRpzOk() (*bool, bool)`

GetEnableGlobalAllowListRpzOk returns a tuple with the EnableGlobalAllowListRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGlobalAllowListRpz

`func (o *ParentalcontrolSubscribersite) SetEnableGlobalAllowListRpz(v bool)`

SetEnableGlobalAllowListRpz sets EnableGlobalAllowListRpz field to given value.

### HasEnableGlobalAllowListRpz

`func (o *ParentalcontrolSubscribersite) HasEnableGlobalAllowListRpz() bool`

HasEnableGlobalAllowListRpz returns a boolean if a field has been set.

### GetEnableRpzFilteringBypass

`func (o *ParentalcontrolSubscribersite) GetEnableRpzFilteringBypass() bool`

GetEnableRpzFilteringBypass returns the EnableRpzFilteringBypass field if non-nil, zero value otherwise.

### GetEnableRpzFilteringBypassOk

`func (o *ParentalcontrolSubscribersite) GetEnableRpzFilteringBypassOk() (*bool, bool)`

GetEnableRpzFilteringBypassOk returns a tuple with the EnableRpzFilteringBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRpzFilteringBypass

`func (o *ParentalcontrolSubscribersite) SetEnableRpzFilteringBypass(v bool)`

SetEnableRpzFilteringBypass sets EnableRpzFilteringBypass field to given value.

### HasEnableRpzFilteringBypass

`func (o *ParentalcontrolSubscribersite) HasEnableRpzFilteringBypass() bool`

HasEnableRpzFilteringBypass returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *ParentalcontrolSubscribersite) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *ParentalcontrolSubscribersite) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *ParentalcontrolSubscribersite) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *ParentalcontrolSubscribersite) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *ParentalcontrolSubscribersite) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *ParentalcontrolSubscribersite) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *ParentalcontrolSubscribersite) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *ParentalcontrolSubscribersite) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *ParentalcontrolSubscribersite) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *ParentalcontrolSubscribersite) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *ParentalcontrolSubscribersite) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *ParentalcontrolSubscribersite) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFirstPort

`func (o *ParentalcontrolSubscribersite) GetFirstPort() int64`

GetFirstPort returns the FirstPort field if non-nil, zero value otherwise.

### GetFirstPortOk

`func (o *ParentalcontrolSubscribersite) GetFirstPortOk() (*int64, bool)`

GetFirstPortOk returns a tuple with the FirstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPort

`func (o *ParentalcontrolSubscribersite) SetFirstPort(v int64)`

SetFirstPort sets FirstPort field to given value.

### HasFirstPort

`func (o *ParentalcontrolSubscribersite) HasFirstPort() bool`

HasFirstPort returns a boolean if a field has been set.

### GetGlobalAllowListRpz

`func (o *ParentalcontrolSubscribersite) GetGlobalAllowListRpz() int64`

GetGlobalAllowListRpz returns the GlobalAllowListRpz field if non-nil, zero value otherwise.

### GetGlobalAllowListRpzOk

`func (o *ParentalcontrolSubscribersite) GetGlobalAllowListRpzOk() (*int64, bool)`

GetGlobalAllowListRpzOk returns a tuple with the GlobalAllowListRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAllowListRpz

`func (o *ParentalcontrolSubscribersite) SetGlobalAllowListRpz(v int64)`

SetGlobalAllowListRpz sets GlobalAllowListRpz field to given value.

### HasGlobalAllowListRpz

`func (o *ParentalcontrolSubscribersite) HasGlobalAllowListRpz() bool`

HasGlobalAllowListRpz returns a boolean if a field has been set.

### GetMaximumSubscribers

`func (o *ParentalcontrolSubscribersite) GetMaximumSubscribers() int64`

GetMaximumSubscribers returns the MaximumSubscribers field if non-nil, zero value otherwise.

### GetMaximumSubscribersOk

`func (o *ParentalcontrolSubscribersite) GetMaximumSubscribersOk() (*int64, bool)`

GetMaximumSubscribersOk returns a tuple with the MaximumSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSubscribers

`func (o *ParentalcontrolSubscribersite) SetMaximumSubscribers(v int64)`

SetMaximumSubscribers sets MaximumSubscribers field to given value.

### HasMaximumSubscribers

`func (o *ParentalcontrolSubscribersite) HasMaximumSubscribers() bool`

HasMaximumSubscribers returns a boolean if a field has been set.

### GetMembers

`func (o *ParentalcontrolSubscribersite) GetMembers() []ParentalcontrolSubscribersiteMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ParentalcontrolSubscribersite) GetMembersOk() (*[]ParentalcontrolSubscribersiteMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ParentalcontrolSubscribersite) SetMembers(v []ParentalcontrolSubscribersiteMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ParentalcontrolSubscribersite) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMsps

`func (o *ParentalcontrolSubscribersite) GetMsps() []ParentalcontrolSubscribersiteMsps`

GetMsps returns the Msps field if non-nil, zero value otherwise.

### GetMspsOk

`func (o *ParentalcontrolSubscribersite) GetMspsOk() (*[]ParentalcontrolSubscribersiteMsps, bool)`

GetMspsOk returns a tuple with the Msps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsps

`func (o *ParentalcontrolSubscribersite) SetMsps(v []ParentalcontrolSubscribersiteMsps)`

SetMsps sets Msps field to given value.

### HasMsps

`func (o *ParentalcontrolSubscribersite) HasMsps() bool`

HasMsps returns a boolean if a field has been set.

### GetName

`func (o *ParentalcontrolSubscribersite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParentalcontrolSubscribersite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParentalcontrolSubscribersite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParentalcontrolSubscribersite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNasGateways

`func (o *ParentalcontrolSubscribersite) GetNasGateways() []ParentalcontrolSubscribersiteNasGateways`

GetNasGateways returns the NasGateways field if non-nil, zero value otherwise.

### GetNasGatewaysOk

`func (o *ParentalcontrolSubscribersite) GetNasGatewaysOk() (*[]ParentalcontrolSubscribersiteNasGateways, bool)`

GetNasGatewaysOk returns a tuple with the NasGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasGateways

`func (o *ParentalcontrolSubscribersite) SetNasGateways(v []ParentalcontrolSubscribersiteNasGateways)`

SetNasGateways sets NasGateways field to given value.

### HasNasGateways

`func (o *ParentalcontrolSubscribersite) HasNasGateways() bool`

HasNasGateways returns a boolean if a field has been set.

### GetNasPort

`func (o *ParentalcontrolSubscribersite) GetNasPort() int64`

GetNasPort returns the NasPort field if non-nil, zero value otherwise.

### GetNasPortOk

`func (o *ParentalcontrolSubscribersite) GetNasPortOk() (*int64, bool)`

GetNasPortOk returns a tuple with the NasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasPort

`func (o *ParentalcontrolSubscribersite) SetNasPort(v int64)`

SetNasPort sets NasPort field to given value.

### HasNasPort

`func (o *ParentalcontrolSubscribersite) HasNasPort() bool`

HasNasPort returns a boolean if a field has been set.

### GetProxyRpzPassthru

`func (o *ParentalcontrolSubscribersite) GetProxyRpzPassthru() bool`

GetProxyRpzPassthru returns the ProxyRpzPassthru field if non-nil, zero value otherwise.

### GetProxyRpzPassthruOk

`func (o *ParentalcontrolSubscribersite) GetProxyRpzPassthruOk() (*bool, bool)`

GetProxyRpzPassthruOk returns a tuple with the ProxyRpzPassthru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyRpzPassthru

`func (o *ParentalcontrolSubscribersite) SetProxyRpzPassthru(v bool)`

SetProxyRpzPassthru sets ProxyRpzPassthru field to given value.

### HasProxyRpzPassthru

`func (o *ParentalcontrolSubscribersite) HasProxyRpzPassthru() bool`

HasProxyRpzPassthru returns a boolean if a field has been set.

### GetSpms

`func (o *ParentalcontrolSubscribersite) GetSpms() []ParentalcontrolSubscribersiteSpms`

GetSpms returns the Spms field if non-nil, zero value otherwise.

### GetSpmsOk

`func (o *ParentalcontrolSubscribersite) GetSpmsOk() (*[]ParentalcontrolSubscribersiteSpms, bool)`

GetSpmsOk returns a tuple with the Spms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpms

`func (o *ParentalcontrolSubscribersite) SetSpms(v []ParentalcontrolSubscribersiteSpms)`

SetSpms sets Spms field to given value.

### HasSpms

`func (o *ParentalcontrolSubscribersite) HasSpms() bool`

HasSpms returns a boolean if a field has been set.

### GetStopAnycast

`func (o *ParentalcontrolSubscribersite) GetStopAnycast() bool`

GetStopAnycast returns the StopAnycast field if non-nil, zero value otherwise.

### GetStopAnycastOk

`func (o *ParentalcontrolSubscribersite) GetStopAnycastOk() (*bool, bool)`

GetStopAnycastOk returns a tuple with the StopAnycast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopAnycast

`func (o *ParentalcontrolSubscribersite) SetStopAnycast(v bool)`

SetStopAnycast sets StopAnycast field to given value.

### HasStopAnycast

`func (o *ParentalcontrolSubscribersite) HasStopAnycast() bool`

HasStopAnycast returns a boolean if a field has been set.

### GetStrictNat

`func (o *ParentalcontrolSubscribersite) GetStrictNat() bool`

GetStrictNat returns the StrictNat field if non-nil, zero value otherwise.

### GetStrictNatOk

`func (o *ParentalcontrolSubscribersite) GetStrictNatOk() (*bool, bool)`

GetStrictNatOk returns a tuple with the StrictNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictNat

`func (o *ParentalcontrolSubscribersite) SetStrictNat(v bool)`

SetStrictNat sets StrictNat field to given value.

### HasStrictNat

`func (o *ParentalcontrolSubscribersite) HasStrictNat() bool`

HasStrictNat returns a boolean if a field has been set.

### GetSubscriberCollectionType

`func (o *ParentalcontrolSubscribersite) GetSubscriberCollectionType() string`

GetSubscriberCollectionType returns the SubscriberCollectionType field if non-nil, zero value otherwise.

### GetSubscriberCollectionTypeOk

`func (o *ParentalcontrolSubscribersite) GetSubscriberCollectionTypeOk() (*string, bool)`

GetSubscriberCollectionTypeOk returns a tuple with the SubscriberCollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCollectionType

`func (o *ParentalcontrolSubscribersite) SetSubscriberCollectionType(v string)`

SetSubscriberCollectionType sets SubscriberCollectionType field to given value.

### HasSubscriberCollectionType

`func (o *ParentalcontrolSubscribersite) HasSubscriberCollectionType() bool`

HasSubscriberCollectionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


