# GetRecordNsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Addresses** | Pointer to [**[]RecordNsAddresses**](RecordNsAddresses.md) | The list of zone name servers. | [optional] 
**CloudInfo** | Pointer to [**RecordNsCloudInfo**](RecordNsCloudInfo.md) |  | [optional] 
**Creator** | Pointer to **string** | The record creator. | [optional] [readonly] 
**DnsName** | Pointer to **string** | The name of the NS record in punycode format. | [optional] [readonly] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**MsDelegationName** | Pointer to **string** | The MS delegation point name. | [optional] 
**Name** | Pointer to **string** | The name of the NS record in FQDN format. This value can be in unicode format. | [optional] 
**Nameserver** | Pointer to **string** | The domain name of an authoritative server for the redirected zone. | [optional] 
**Policy** | Pointer to **string** | The host name policy for the record. | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 
**Result** | Pointer to [**RecordNs**](RecordNs.md) |  | [optional] 

## Methods

### NewGetRecordNsResponse

`func NewGetRecordNsResponse() *GetRecordNsResponse`

NewGetRecordNsResponse instantiates a new GetRecordNsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordNsResponseWithDefaults

`func NewGetRecordNsResponseWithDefaults() *GetRecordNsResponse`

NewGetRecordNsResponseWithDefaults instantiates a new GetRecordNsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordNsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordNsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordNsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordNsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddresses

`func (o *GetRecordNsResponse) GetAddresses() []RecordNsAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetRecordNsResponse) GetAddressesOk() (*[]RecordNsAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetRecordNsResponse) SetAddresses(v []RecordNsAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *GetRecordNsResponse) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetRecordNsResponse) GetCloudInfo() RecordNsCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetRecordNsResponse) GetCloudInfoOk() (*RecordNsCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetRecordNsResponse) SetCloudInfo(v RecordNsCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetRecordNsResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetCreator

`func (o *GetRecordNsResponse) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetRecordNsResponse) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetRecordNsResponse) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetRecordNsResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDnsName

`func (o *GetRecordNsResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *GetRecordNsResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *GetRecordNsResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *GetRecordNsResponse) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetRecordNsResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetRecordNsResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetRecordNsResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetRecordNsResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMsDelegationName

`func (o *GetRecordNsResponse) GetMsDelegationName() string`

GetMsDelegationName returns the MsDelegationName field if non-nil, zero value otherwise.

### GetMsDelegationNameOk

`func (o *GetRecordNsResponse) GetMsDelegationNameOk() (*string, bool)`

GetMsDelegationNameOk returns a tuple with the MsDelegationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDelegationName

`func (o *GetRecordNsResponse) SetMsDelegationName(v string)`

SetMsDelegationName sets MsDelegationName field to given value.

### HasMsDelegationName

`func (o *GetRecordNsResponse) HasMsDelegationName() bool`

HasMsDelegationName returns a boolean if a field has been set.

### GetName

`func (o *GetRecordNsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRecordNsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRecordNsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRecordNsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameserver

`func (o *GetRecordNsResponse) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *GetRecordNsResponse) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *GetRecordNsResponse) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *GetRecordNsResponse) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetPolicy

`func (o *GetRecordNsResponse) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *GetRecordNsResponse) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *GetRecordNsResponse) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *GetRecordNsResponse) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetView

`func (o *GetRecordNsResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetRecordNsResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetRecordNsResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetRecordNsResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetRecordNsResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetRecordNsResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetRecordNsResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetRecordNsResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordNsResponse) GetResult() RecordNs`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordNsResponse) GetResultOk() (*RecordNs, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordNsResponse) SetResult(v RecordNs)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordNsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


