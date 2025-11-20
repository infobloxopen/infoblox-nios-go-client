# RecordNsec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CloudInfo** | Pointer to [**RecordNsecCloudInfo**](RecordNsecCloudInfo.md) |  | [optional] 
**CreationTime** | Pointer to **int64** | Time that the record was created. | [optional] [readonly] 
**Creator** | Pointer to **string** | Creator of the record. | [optional] [readonly] 
**DnsName** | Pointer to **string** | Name for an NSEC record in punycode format. | [optional] [readonly] 
**DnsNextOwnerName** | Pointer to **string** | Name of the next owner in punycode format. | [optional] [readonly] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the NSEC record in FQDN format. | [optional] [readonly] 
**NextOwnerName** | Pointer to **string** | Name of the next owner that has authoritative data or that contains a delegation point NS record. | [optional] [readonly] 
**RrsetTypes** | Pointer to **[]string** | The RRSet types that exist at the original owner name of the NSEC RR. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] [readonly] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] [readonly] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordNsec

`func NewRecordNsec() *RecordNsec`

NewRecordNsec instantiates a new RecordNsec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordNsecWithDefaults

`func NewRecordNsecWithDefaults() *RecordNsec`

NewRecordNsecWithDefaults instantiates a new RecordNsec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordNsec) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordNsec) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordNsec) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordNsec) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordNsec) GetCloudInfo() RecordNsecCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordNsec) GetCloudInfoOk() (*RecordNsecCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordNsec) SetCloudInfo(v RecordNsecCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordNsec) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordNsec) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordNsec) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordNsec) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordNsec) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordNsec) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordNsec) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordNsec) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordNsec) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordNsec) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordNsec) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordNsec) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordNsec) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDnsNextOwnerName

`func (o *RecordNsec) GetDnsNextOwnerName() string`

GetDnsNextOwnerName returns the DnsNextOwnerName field if non-nil, zero value otherwise.

### GetDnsNextOwnerNameOk

`func (o *RecordNsec) GetDnsNextOwnerNameOk() (*string, bool)`

GetDnsNextOwnerNameOk returns a tuple with the DnsNextOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNextOwnerName

`func (o *RecordNsec) SetDnsNextOwnerName(v string)`

SetDnsNextOwnerName sets DnsNextOwnerName field to given value.

### HasDnsNextOwnerName

`func (o *RecordNsec) HasDnsNextOwnerName() bool`

HasDnsNextOwnerName returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordNsec) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordNsec) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordNsec) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordNsec) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordNsec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordNsec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordNsec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordNsec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextOwnerName

`func (o *RecordNsec) GetNextOwnerName() string`

GetNextOwnerName returns the NextOwnerName field if non-nil, zero value otherwise.

### GetNextOwnerNameOk

`func (o *RecordNsec) GetNextOwnerNameOk() (*string, bool)`

GetNextOwnerNameOk returns a tuple with the NextOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextOwnerName

`func (o *RecordNsec) SetNextOwnerName(v string)`

SetNextOwnerName sets NextOwnerName field to given value.

### HasNextOwnerName

`func (o *RecordNsec) HasNextOwnerName() bool`

HasNextOwnerName returns a boolean if a field has been set.

### GetRrsetTypes

`func (o *RecordNsec) GetRrsetTypes() []string`

GetRrsetTypes returns the RrsetTypes field if non-nil, zero value otherwise.

### GetRrsetTypesOk

`func (o *RecordNsec) GetRrsetTypesOk() (*[]string, bool)`

GetRrsetTypesOk returns a tuple with the RrsetTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrsetTypes

`func (o *RecordNsec) SetRrsetTypes(v []string)`

SetRrsetTypes sets RrsetTypes field to given value.

### HasRrsetTypes

`func (o *RecordNsec) HasRrsetTypes() bool`

HasRrsetTypes returns a boolean if a field has been set.

### GetTtl

`func (o *RecordNsec) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordNsec) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordNsec) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordNsec) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordNsec) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordNsec) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordNsec) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordNsec) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordNsec) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordNsec) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordNsec) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordNsec) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordNsec) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordNsec) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordNsec) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordNsec) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


