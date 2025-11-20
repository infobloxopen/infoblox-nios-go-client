# RecordDhcid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CreationTime** | Pointer to **int64** | The creation time of the record. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. | [optional] [readonly] 
**Dhcid** | Pointer to **string** | The Base64 encoded DHCP client information. | [optional] [readonly] 
**DnsName** | Pointer to **string** | The name for the DHCID record in punycode format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the DHCID record in FQDN format. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] [readonly] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] [readonly] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordDhcid

`func NewRecordDhcid() *RecordDhcid`

NewRecordDhcid instantiates a new RecordDhcid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordDhcidWithDefaults

`func NewRecordDhcidWithDefaults() *RecordDhcid`

NewRecordDhcidWithDefaults instantiates a new RecordDhcid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordDhcid) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordDhcid) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordDhcid) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordDhcid) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordDhcid) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordDhcid) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordDhcid) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordDhcid) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordDhcid) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordDhcid) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordDhcid) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordDhcid) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDhcid

`func (o *RecordDhcid) GetDhcid() string`

GetDhcid returns the Dhcid field if non-nil, zero value otherwise.

### GetDhcidOk

`func (o *RecordDhcid) GetDhcidOk() (*string, bool)`

GetDhcidOk returns a tuple with the Dhcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcid

`func (o *RecordDhcid) SetDhcid(v string)`

SetDhcid sets Dhcid field to given value.

### HasDhcid

`func (o *RecordDhcid) HasDhcid() bool`

HasDhcid returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordDhcid) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordDhcid) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordDhcid) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordDhcid) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetName

`func (o *RecordDhcid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordDhcid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordDhcid) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordDhcid) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTtl

`func (o *RecordDhcid) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordDhcid) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordDhcid) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordDhcid) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordDhcid) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordDhcid) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordDhcid) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordDhcid) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordDhcid) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordDhcid) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordDhcid) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordDhcid) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordDhcid) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordDhcid) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordDhcid) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordDhcid) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


