# RecordNs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Addresses** | Pointer to [**[]RecordNsAddresses**](RecordNsAddresses.md) | The list of zone name servers. | [optional] 
**CloudInfo** | Pointer to [**RecordNsCloudInfo**](RecordNsCloudInfo.md) |  | [optional] 
**Creator** | Pointer to **string** | The record creator. | [optional] [readonly] 
**DnsName** | Pointer to **string** | The name of the NS record in punycode format. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**MsDelegationName** | Pointer to **string** | The MS delegation point name. | [optional] 
**Name** | Pointer to **string** | The name of the NS record in FQDN format. This value can be in unicode format. | [optional] 
**Nameserver** | Pointer to **string** | The domain name of an authoritative server for the redirected zone. | [optional] 
**Policy** | Pointer to **string** | The host name policy for the record. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordNs

`func NewRecordNs() *RecordNs`

NewRecordNs instantiates a new RecordNs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordNsWithDefaults

`func NewRecordNsWithDefaults() *RecordNs`

NewRecordNsWithDefaults instantiates a new RecordNs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordNs) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordNs) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordNs) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordNs) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddresses

`func (o *RecordNs) GetAddresses() []RecordNsAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *RecordNs) GetAddressesOk() (*[]RecordNsAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *RecordNs) SetAddresses(v []RecordNsAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *RecordNs) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordNs) GetCloudInfo() RecordNsCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordNs) GetCloudInfoOk() (*RecordNsCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordNs) SetCloudInfo(v RecordNsCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordNs) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetCreator

`func (o *RecordNs) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordNs) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordNs) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordNs) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordNs) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordNs) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordNs) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordNs) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordNs) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordNs) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordNs) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordNs) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordNs) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordNs) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordNs) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordNs) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordNs) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordNs) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordNs) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordNs) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordNs) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordNs) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordNs) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordNs) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMsDelegationName

`func (o *RecordNs) GetMsDelegationName() string`

GetMsDelegationName returns the MsDelegationName field if non-nil, zero value otherwise.

### GetMsDelegationNameOk

`func (o *RecordNs) GetMsDelegationNameOk() (*string, bool)`

GetMsDelegationNameOk returns a tuple with the MsDelegationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDelegationName

`func (o *RecordNs) SetMsDelegationName(v string)`

SetMsDelegationName sets MsDelegationName field to given value.

### HasMsDelegationName

`func (o *RecordNs) HasMsDelegationName() bool`

HasMsDelegationName returns a boolean if a field has been set.

### GetName

`func (o *RecordNs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordNs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordNs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordNs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameserver

`func (o *RecordNs) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *RecordNs) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *RecordNs) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *RecordNs) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetPolicy

`func (o *RecordNs) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *RecordNs) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *RecordNs) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *RecordNs) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetUuid

`func (o *RecordNs) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RecordNs) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RecordNs) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RecordNs) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetView

`func (o *RecordNs) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordNs) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordNs) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordNs) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordNs) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordNs) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordNs) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordNs) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


