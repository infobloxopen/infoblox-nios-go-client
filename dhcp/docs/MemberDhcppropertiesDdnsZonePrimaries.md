# MemberDhcppropertiesDdnsZonePrimaries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneMatch** | Pointer to **string** | Indicate matching type. | [optional] 
**DnsGridZone** | Pointer to **string** | The ref of a DNS zone. | [optional] 
**DnsGridPrimary** | Pointer to **string** | The name of a Grid member. | [optional] 
**DnsExtZone** | Pointer to **string** | The name of external zone in FQDN format. | [optional] 
**DnsExtPrimary** | Pointer to **string** | The IP address of the External server. Valid when zone_match is \&quot;EXTERNAL\&quot; or \&quot;ANY_EXTERNAL\&quot;. | [optional] 

## Methods

### NewMemberDhcppropertiesDdnsZonePrimaries

`func NewMemberDhcppropertiesDdnsZonePrimaries() *MemberDhcppropertiesDdnsZonePrimaries`

NewMemberDhcppropertiesDdnsZonePrimaries instantiates a new MemberDhcppropertiesDdnsZonePrimaries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDhcppropertiesDdnsZonePrimariesWithDefaults

`func NewMemberDhcppropertiesDdnsZonePrimariesWithDefaults() *MemberDhcppropertiesDdnsZonePrimaries`

NewMemberDhcppropertiesDdnsZonePrimariesWithDefaults instantiates a new MemberDhcppropertiesDdnsZonePrimaries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneMatch

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetZoneMatch() string`

GetZoneMatch returns the ZoneMatch field if non-nil, zero value otherwise.

### GetZoneMatchOk

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetZoneMatchOk() (*string, bool)`

GetZoneMatchOk returns a tuple with the ZoneMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMatch

`func (o *MemberDhcppropertiesDdnsZonePrimaries) SetZoneMatch(v string)`

SetZoneMatch sets ZoneMatch field to given value.

### HasZoneMatch

`func (o *MemberDhcppropertiesDdnsZonePrimaries) HasZoneMatch() bool`

HasZoneMatch returns a boolean if a field has been set.

### GetDnsGridZone

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetDnsGridZone() string`

GetDnsGridZone returns the DnsGridZone field if non-nil, zero value otherwise.

### GetDnsGridZoneOk

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetDnsGridZoneOk() (*string, bool)`

GetDnsGridZoneOk returns a tuple with the DnsGridZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsGridZone

`func (o *MemberDhcppropertiesDdnsZonePrimaries) SetDnsGridZone(v string)`

SetDnsGridZone sets DnsGridZone field to given value.

### HasDnsGridZone

`func (o *MemberDhcppropertiesDdnsZonePrimaries) HasDnsGridZone() bool`

HasDnsGridZone returns a boolean if a field has been set.

### GetDnsGridPrimary

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetDnsGridPrimary() string`

GetDnsGridPrimary returns the DnsGridPrimary field if non-nil, zero value otherwise.

### GetDnsGridPrimaryOk

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetDnsGridPrimaryOk() (*string, bool)`

GetDnsGridPrimaryOk returns a tuple with the DnsGridPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsGridPrimary

`func (o *MemberDhcppropertiesDdnsZonePrimaries) SetDnsGridPrimary(v string)`

SetDnsGridPrimary sets DnsGridPrimary field to given value.

### HasDnsGridPrimary

`func (o *MemberDhcppropertiesDdnsZonePrimaries) HasDnsGridPrimary() bool`

HasDnsGridPrimary returns a boolean if a field has been set.

### GetDnsExtZone

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetDnsExtZone() string`

GetDnsExtZone returns the DnsExtZone field if non-nil, zero value otherwise.

### GetDnsExtZoneOk

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetDnsExtZoneOk() (*string, bool)`

GetDnsExtZoneOk returns a tuple with the DnsExtZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsExtZone

`func (o *MemberDhcppropertiesDdnsZonePrimaries) SetDnsExtZone(v string)`

SetDnsExtZone sets DnsExtZone field to given value.

### HasDnsExtZone

`func (o *MemberDhcppropertiesDdnsZonePrimaries) HasDnsExtZone() bool`

HasDnsExtZone returns a boolean if a field has been set.

### GetDnsExtPrimary

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetDnsExtPrimary() string`

GetDnsExtPrimary returns the DnsExtPrimary field if non-nil, zero value otherwise.

### GetDnsExtPrimaryOk

`func (o *MemberDhcppropertiesDdnsZonePrimaries) GetDnsExtPrimaryOk() (*string, bool)`

GetDnsExtPrimaryOk returns a tuple with the DnsExtPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsExtPrimary

`func (o *MemberDhcppropertiesDdnsZonePrimaries) SetDnsExtPrimary(v string)`

SetDnsExtPrimary sets DnsExtPrimary field to given value.

### HasDnsExtPrimary

`func (o *MemberDhcppropertiesDdnsZonePrimaries) HasDnsExtPrimary() bool`

HasDnsExtPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


