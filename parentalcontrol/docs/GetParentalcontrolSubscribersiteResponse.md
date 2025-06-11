# GetParentalcontrolSubscribersiteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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
**Result** | Pointer to [**ParentalcontrolSubscribersite**](ParentalcontrolSubscribersite.md) |  | [optional] 

## Methods

### NewGetParentalcontrolSubscribersiteResponse

`func NewGetParentalcontrolSubscribersiteResponse() *GetParentalcontrolSubscribersiteResponse`

NewGetParentalcontrolSubscribersiteResponse instantiates a new GetParentalcontrolSubscribersiteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetParentalcontrolSubscribersiteResponseWithDefaults

`func NewGetParentalcontrolSubscribersiteResponseWithDefaults() *GetParentalcontrolSubscribersiteResponse`

NewGetParentalcontrolSubscribersiteResponseWithDefaults instantiates a new GetParentalcontrolSubscribersiteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetParentalcontrolSubscribersiteResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetParentalcontrolSubscribersiteResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetParentalcontrolSubscribersiteResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAbss

`func (o *GetParentalcontrolSubscribersiteResponse) GetAbss() []ParentalcontrolSubscribersiteAbss`

GetAbss returns the Abss field if non-nil, zero value otherwise.

### GetAbssOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetAbssOk() (*[]ParentalcontrolSubscribersiteAbss, bool)`

GetAbssOk returns a tuple with the Abss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbss

`func (o *GetParentalcontrolSubscribersiteResponse) SetAbss(v []ParentalcontrolSubscribersiteAbss)`

SetAbss sets Abss field to given value.

### HasAbss

`func (o *GetParentalcontrolSubscribersiteResponse) HasAbss() bool`

HasAbss returns a boolean if a field has been set.

### GetApiMembers

`func (o *GetParentalcontrolSubscribersiteResponse) GetApiMembers() []ParentalcontrolSubscribersiteApiMembers`

GetApiMembers returns the ApiMembers field if non-nil, zero value otherwise.

### GetApiMembersOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetApiMembersOk() (*[]ParentalcontrolSubscribersiteApiMembers, bool)`

GetApiMembersOk returns a tuple with the ApiMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMembers

`func (o *GetParentalcontrolSubscribersiteResponse) SetApiMembers(v []ParentalcontrolSubscribersiteApiMembers)`

SetApiMembers sets ApiMembers field to given value.

### HasApiMembers

`func (o *GetParentalcontrolSubscribersiteResponse) HasApiMembers() bool`

HasApiMembers returns a boolean if a field has been set.

### GetApiPort

`func (o *GetParentalcontrolSubscribersiteResponse) GetApiPort() int64`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetApiPortOk() (*int64, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *GetParentalcontrolSubscribersiteResponse) SetApiPort(v int64)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *GetParentalcontrolSubscribersiteResponse) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetBlockSize

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockSize() int64`

GetBlockSize returns the BlockSize field if non-nil, zero value otherwise.

### GetBlockSizeOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockSizeOk() (*int64, bool)`

GetBlockSizeOk returns a tuple with the BlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockSize

`func (o *GetParentalcontrolSubscribersiteResponse) SetBlockSize(v int64)`

SetBlockSize sets BlockSize field to given value.

### HasBlockSize

`func (o *GetParentalcontrolSubscribersiteResponse) HasBlockSize() bool`

HasBlockSize returns a boolean if a field has been set.

### GetBlockingIpv4Vip1

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockingIpv4Vip1() string`

GetBlockingIpv4Vip1 returns the BlockingIpv4Vip1 field if non-nil, zero value otherwise.

### GetBlockingIpv4Vip1Ok

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockingIpv4Vip1Ok() (*string, bool)`

GetBlockingIpv4Vip1Ok returns a tuple with the BlockingIpv4Vip1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv4Vip1

`func (o *GetParentalcontrolSubscribersiteResponse) SetBlockingIpv4Vip1(v string)`

SetBlockingIpv4Vip1 sets BlockingIpv4Vip1 field to given value.

### HasBlockingIpv4Vip1

`func (o *GetParentalcontrolSubscribersiteResponse) HasBlockingIpv4Vip1() bool`

HasBlockingIpv4Vip1 returns a boolean if a field has been set.

### GetBlockingIpv4Vip2

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockingIpv4Vip2() string`

GetBlockingIpv4Vip2 returns the BlockingIpv4Vip2 field if non-nil, zero value otherwise.

### GetBlockingIpv4Vip2Ok

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockingIpv4Vip2Ok() (*string, bool)`

GetBlockingIpv4Vip2Ok returns a tuple with the BlockingIpv4Vip2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv4Vip2

`func (o *GetParentalcontrolSubscribersiteResponse) SetBlockingIpv4Vip2(v string)`

SetBlockingIpv4Vip2 sets BlockingIpv4Vip2 field to given value.

### HasBlockingIpv4Vip2

`func (o *GetParentalcontrolSubscribersiteResponse) HasBlockingIpv4Vip2() bool`

HasBlockingIpv4Vip2 returns a boolean if a field has been set.

### GetBlockingIpv6Vip1

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockingIpv6Vip1() string`

GetBlockingIpv6Vip1 returns the BlockingIpv6Vip1 field if non-nil, zero value otherwise.

### GetBlockingIpv6Vip1Ok

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockingIpv6Vip1Ok() (*string, bool)`

GetBlockingIpv6Vip1Ok returns a tuple with the BlockingIpv6Vip1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv6Vip1

`func (o *GetParentalcontrolSubscribersiteResponse) SetBlockingIpv6Vip1(v string)`

SetBlockingIpv6Vip1 sets BlockingIpv6Vip1 field to given value.

### HasBlockingIpv6Vip1

`func (o *GetParentalcontrolSubscribersiteResponse) HasBlockingIpv6Vip1() bool`

HasBlockingIpv6Vip1 returns a boolean if a field has been set.

### GetBlockingIpv6Vip2

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockingIpv6Vip2() string`

GetBlockingIpv6Vip2 returns the BlockingIpv6Vip2 field if non-nil, zero value otherwise.

### GetBlockingIpv6Vip2Ok

`func (o *GetParentalcontrolSubscribersiteResponse) GetBlockingIpv6Vip2Ok() (*string, bool)`

GetBlockingIpv6Vip2Ok returns a tuple with the BlockingIpv6Vip2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockingIpv6Vip2

`func (o *GetParentalcontrolSubscribersiteResponse) SetBlockingIpv6Vip2(v string)`

SetBlockingIpv6Vip2 sets BlockingIpv6Vip2 field to given value.

### HasBlockingIpv6Vip2

`func (o *GetParentalcontrolSubscribersiteResponse) HasBlockingIpv6Vip2() bool`

HasBlockingIpv6Vip2 returns a boolean if a field has been set.

### GetComment

`func (o *GetParentalcontrolSubscribersiteResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetParentalcontrolSubscribersiteResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetParentalcontrolSubscribersiteResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDcaSubBwList

`func (o *GetParentalcontrolSubscribersiteResponse) GetDcaSubBwList() bool`

GetDcaSubBwList returns the DcaSubBwList field if non-nil, zero value otherwise.

### GetDcaSubBwListOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetDcaSubBwListOk() (*bool, bool)`

GetDcaSubBwListOk returns a tuple with the DcaSubBwList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcaSubBwList

`func (o *GetParentalcontrolSubscribersiteResponse) SetDcaSubBwList(v bool)`

SetDcaSubBwList sets DcaSubBwList field to given value.

### HasDcaSubBwList

`func (o *GetParentalcontrolSubscribersiteResponse) HasDcaSubBwList() bool`

HasDcaSubBwList returns a boolean if a field has been set.

### GetDcaSubQueryCount

`func (o *GetParentalcontrolSubscribersiteResponse) GetDcaSubQueryCount() bool`

GetDcaSubQueryCount returns the DcaSubQueryCount field if non-nil, zero value otherwise.

### GetDcaSubQueryCountOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetDcaSubQueryCountOk() (*bool, bool)`

GetDcaSubQueryCountOk returns a tuple with the DcaSubQueryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDcaSubQueryCount

`func (o *GetParentalcontrolSubscribersiteResponse) SetDcaSubQueryCount(v bool)`

SetDcaSubQueryCount sets DcaSubQueryCount field to given value.

### HasDcaSubQueryCount

`func (o *GetParentalcontrolSubscribersiteResponse) HasDcaSubQueryCount() bool`

HasDcaSubQueryCount returns a boolean if a field has been set.

### GetEnableGlobalAllowListRpz

`func (o *GetParentalcontrolSubscribersiteResponse) GetEnableGlobalAllowListRpz() bool`

GetEnableGlobalAllowListRpz returns the EnableGlobalAllowListRpz field if non-nil, zero value otherwise.

### GetEnableGlobalAllowListRpzOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetEnableGlobalAllowListRpzOk() (*bool, bool)`

GetEnableGlobalAllowListRpzOk returns a tuple with the EnableGlobalAllowListRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGlobalAllowListRpz

`func (o *GetParentalcontrolSubscribersiteResponse) SetEnableGlobalAllowListRpz(v bool)`

SetEnableGlobalAllowListRpz sets EnableGlobalAllowListRpz field to given value.

### HasEnableGlobalAllowListRpz

`func (o *GetParentalcontrolSubscribersiteResponse) HasEnableGlobalAllowListRpz() bool`

HasEnableGlobalAllowListRpz returns a boolean if a field has been set.

### GetEnableRpzFilteringBypass

`func (o *GetParentalcontrolSubscribersiteResponse) GetEnableRpzFilteringBypass() bool`

GetEnableRpzFilteringBypass returns the EnableRpzFilteringBypass field if non-nil, zero value otherwise.

### GetEnableRpzFilteringBypassOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetEnableRpzFilteringBypassOk() (*bool, bool)`

GetEnableRpzFilteringBypassOk returns a tuple with the EnableRpzFilteringBypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRpzFilteringBypass

`func (o *GetParentalcontrolSubscribersiteResponse) SetEnableRpzFilteringBypass(v bool)`

SetEnableRpzFilteringBypass sets EnableRpzFilteringBypass field to given value.

### HasEnableRpzFilteringBypass

`func (o *GetParentalcontrolSubscribersiteResponse) HasEnableRpzFilteringBypass() bool`

HasEnableRpzFilteringBypass returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetParentalcontrolSubscribersiteResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetParentalcontrolSubscribersiteResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetParentalcontrolSubscribersiteResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFirstPort

`func (o *GetParentalcontrolSubscribersiteResponse) GetFirstPort() int64`

GetFirstPort returns the FirstPort field if non-nil, zero value otherwise.

### GetFirstPortOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetFirstPortOk() (*int64, bool)`

GetFirstPortOk returns a tuple with the FirstPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPort

`func (o *GetParentalcontrolSubscribersiteResponse) SetFirstPort(v int64)`

SetFirstPort sets FirstPort field to given value.

### HasFirstPort

`func (o *GetParentalcontrolSubscribersiteResponse) HasFirstPort() bool`

HasFirstPort returns a boolean if a field has been set.

### GetGlobalAllowListRpz

`func (o *GetParentalcontrolSubscribersiteResponse) GetGlobalAllowListRpz() int64`

GetGlobalAllowListRpz returns the GlobalAllowListRpz field if non-nil, zero value otherwise.

### GetGlobalAllowListRpzOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetGlobalAllowListRpzOk() (*int64, bool)`

GetGlobalAllowListRpzOk returns a tuple with the GlobalAllowListRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAllowListRpz

`func (o *GetParentalcontrolSubscribersiteResponse) SetGlobalAllowListRpz(v int64)`

SetGlobalAllowListRpz sets GlobalAllowListRpz field to given value.

### HasGlobalAllowListRpz

`func (o *GetParentalcontrolSubscribersiteResponse) HasGlobalAllowListRpz() bool`

HasGlobalAllowListRpz returns a boolean if a field has been set.

### GetMaximumSubscribers

`func (o *GetParentalcontrolSubscribersiteResponse) GetMaximumSubscribers() int64`

GetMaximumSubscribers returns the MaximumSubscribers field if non-nil, zero value otherwise.

### GetMaximumSubscribersOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetMaximumSubscribersOk() (*int64, bool)`

GetMaximumSubscribersOk returns a tuple with the MaximumSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumSubscribers

`func (o *GetParentalcontrolSubscribersiteResponse) SetMaximumSubscribers(v int64)`

SetMaximumSubscribers sets MaximumSubscribers field to given value.

### HasMaximumSubscribers

`func (o *GetParentalcontrolSubscribersiteResponse) HasMaximumSubscribers() bool`

HasMaximumSubscribers returns a boolean if a field has been set.

### GetMembers

`func (o *GetParentalcontrolSubscribersiteResponse) GetMembers() []ParentalcontrolSubscribersiteMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetMembersOk() (*[]ParentalcontrolSubscribersiteMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetParentalcontrolSubscribersiteResponse) SetMembers(v []ParentalcontrolSubscribersiteMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetParentalcontrolSubscribersiteResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMsps

`func (o *GetParentalcontrolSubscribersiteResponse) GetMsps() []ParentalcontrolSubscribersiteMsps`

GetMsps returns the Msps field if non-nil, zero value otherwise.

### GetMspsOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetMspsOk() (*[]ParentalcontrolSubscribersiteMsps, bool)`

GetMspsOk returns a tuple with the Msps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsps

`func (o *GetParentalcontrolSubscribersiteResponse) SetMsps(v []ParentalcontrolSubscribersiteMsps)`

SetMsps sets Msps field to given value.

### HasMsps

`func (o *GetParentalcontrolSubscribersiteResponse) HasMsps() bool`

HasMsps returns a boolean if a field has been set.

### GetName

`func (o *GetParentalcontrolSubscribersiteResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetParentalcontrolSubscribersiteResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetParentalcontrolSubscribersiteResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNasGateways

`func (o *GetParentalcontrolSubscribersiteResponse) GetNasGateways() []ParentalcontrolSubscribersiteNasGateways`

GetNasGateways returns the NasGateways field if non-nil, zero value otherwise.

### GetNasGatewaysOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetNasGatewaysOk() (*[]ParentalcontrolSubscribersiteNasGateways, bool)`

GetNasGatewaysOk returns a tuple with the NasGateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasGateways

`func (o *GetParentalcontrolSubscribersiteResponse) SetNasGateways(v []ParentalcontrolSubscribersiteNasGateways)`

SetNasGateways sets NasGateways field to given value.

### HasNasGateways

`func (o *GetParentalcontrolSubscribersiteResponse) HasNasGateways() bool`

HasNasGateways returns a boolean if a field has been set.

### GetNasPort

`func (o *GetParentalcontrolSubscribersiteResponse) GetNasPort() int64`

GetNasPort returns the NasPort field if non-nil, zero value otherwise.

### GetNasPortOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetNasPortOk() (*int64, bool)`

GetNasPortOk returns a tuple with the NasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasPort

`func (o *GetParentalcontrolSubscribersiteResponse) SetNasPort(v int64)`

SetNasPort sets NasPort field to given value.

### HasNasPort

`func (o *GetParentalcontrolSubscribersiteResponse) HasNasPort() bool`

HasNasPort returns a boolean if a field has been set.

### GetProxyRpzPassthru

`func (o *GetParentalcontrolSubscribersiteResponse) GetProxyRpzPassthru() bool`

GetProxyRpzPassthru returns the ProxyRpzPassthru field if non-nil, zero value otherwise.

### GetProxyRpzPassthruOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetProxyRpzPassthruOk() (*bool, bool)`

GetProxyRpzPassthruOk returns a tuple with the ProxyRpzPassthru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyRpzPassthru

`func (o *GetParentalcontrolSubscribersiteResponse) SetProxyRpzPassthru(v bool)`

SetProxyRpzPassthru sets ProxyRpzPassthru field to given value.

### HasProxyRpzPassthru

`func (o *GetParentalcontrolSubscribersiteResponse) HasProxyRpzPassthru() bool`

HasProxyRpzPassthru returns a boolean if a field has been set.

### GetSpms

`func (o *GetParentalcontrolSubscribersiteResponse) GetSpms() []ParentalcontrolSubscribersiteSpms`

GetSpms returns the Spms field if non-nil, zero value otherwise.

### GetSpmsOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetSpmsOk() (*[]ParentalcontrolSubscribersiteSpms, bool)`

GetSpmsOk returns a tuple with the Spms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpms

`func (o *GetParentalcontrolSubscribersiteResponse) SetSpms(v []ParentalcontrolSubscribersiteSpms)`

SetSpms sets Spms field to given value.

### HasSpms

`func (o *GetParentalcontrolSubscribersiteResponse) HasSpms() bool`

HasSpms returns a boolean if a field has been set.

### GetStopAnycast

`func (o *GetParentalcontrolSubscribersiteResponse) GetStopAnycast() bool`

GetStopAnycast returns the StopAnycast field if non-nil, zero value otherwise.

### GetStopAnycastOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetStopAnycastOk() (*bool, bool)`

GetStopAnycastOk returns a tuple with the StopAnycast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopAnycast

`func (o *GetParentalcontrolSubscribersiteResponse) SetStopAnycast(v bool)`

SetStopAnycast sets StopAnycast field to given value.

### HasStopAnycast

`func (o *GetParentalcontrolSubscribersiteResponse) HasStopAnycast() bool`

HasStopAnycast returns a boolean if a field has been set.

### GetStrictNat

`func (o *GetParentalcontrolSubscribersiteResponse) GetStrictNat() bool`

GetStrictNat returns the StrictNat field if non-nil, zero value otherwise.

### GetStrictNatOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetStrictNatOk() (*bool, bool)`

GetStrictNatOk returns a tuple with the StrictNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictNat

`func (o *GetParentalcontrolSubscribersiteResponse) SetStrictNat(v bool)`

SetStrictNat sets StrictNat field to given value.

### HasStrictNat

`func (o *GetParentalcontrolSubscribersiteResponse) HasStrictNat() bool`

HasStrictNat returns a boolean if a field has been set.

### GetSubscriberCollectionType

`func (o *GetParentalcontrolSubscribersiteResponse) GetSubscriberCollectionType() string`

GetSubscriberCollectionType returns the SubscriberCollectionType field if non-nil, zero value otherwise.

### GetSubscriberCollectionTypeOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetSubscriberCollectionTypeOk() (*string, bool)`

GetSubscriberCollectionTypeOk returns a tuple with the SubscriberCollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCollectionType

`func (o *GetParentalcontrolSubscribersiteResponse) SetSubscriberCollectionType(v string)`

SetSubscriberCollectionType sets SubscriberCollectionType field to given value.

### HasSubscriberCollectionType

`func (o *GetParentalcontrolSubscribersiteResponse) HasSubscriberCollectionType() bool`

HasSubscriberCollectionType returns a boolean if a field has been set.

### GetResult

`func (o *GetParentalcontrolSubscribersiteResponse) GetResult() ParentalcontrolSubscribersite`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetParentalcontrolSubscribersiteResponse) GetResultOk() (*ParentalcontrolSubscribersite, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetParentalcontrolSubscribersiteResponse) SetResult(v ParentalcontrolSubscribersite)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetParentalcontrolSubscribersiteResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


