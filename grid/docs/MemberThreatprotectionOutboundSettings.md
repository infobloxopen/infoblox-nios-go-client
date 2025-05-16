# MemberThreatprotectionOutboundSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableQueryFqdn** | Pointer to **bool** | Flag to enable using DNS query FQDN for Outbound. | [optional] 
**QueryFqdnLimit** | Pointer to **int64** | Max domain level for DNS Query FQDN | [optional] 

## Methods

### NewMemberThreatprotectionOutboundSettings

`func NewMemberThreatprotectionOutboundSettings() *MemberThreatprotectionOutboundSettings`

NewMemberThreatprotectionOutboundSettings instantiates a new MemberThreatprotectionOutboundSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberThreatprotectionOutboundSettingsWithDefaults

`func NewMemberThreatprotectionOutboundSettingsWithDefaults() *MemberThreatprotectionOutboundSettings`

NewMemberThreatprotectionOutboundSettingsWithDefaults instantiates a new MemberThreatprotectionOutboundSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableQueryFqdn

`func (o *MemberThreatprotectionOutboundSettings) GetEnableQueryFqdn() bool`

GetEnableQueryFqdn returns the EnableQueryFqdn field if non-nil, zero value otherwise.

### GetEnableQueryFqdnOk

`func (o *MemberThreatprotectionOutboundSettings) GetEnableQueryFqdnOk() (*bool, bool)`

GetEnableQueryFqdnOk returns a tuple with the EnableQueryFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQueryFqdn

`func (o *MemberThreatprotectionOutboundSettings) SetEnableQueryFqdn(v bool)`

SetEnableQueryFqdn sets EnableQueryFqdn field to given value.

### HasEnableQueryFqdn

`func (o *MemberThreatprotectionOutboundSettings) HasEnableQueryFqdn() bool`

HasEnableQueryFqdn returns a boolean if a field has been set.

### GetQueryFqdnLimit

`func (o *MemberThreatprotectionOutboundSettings) GetQueryFqdnLimit() int64`

GetQueryFqdnLimit returns the QueryFqdnLimit field if non-nil, zero value otherwise.

### GetQueryFqdnLimitOk

`func (o *MemberThreatprotectionOutboundSettings) GetQueryFqdnLimitOk() (*int64, bool)`

GetQueryFqdnLimitOk returns a tuple with the QueryFqdnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryFqdnLimit

`func (o *MemberThreatprotectionOutboundSettings) SetQueryFqdnLimit(v int64)`

SetQueryFqdnLimit sets QueryFqdnLimit field to given value.

### HasQueryFqdnLimit

`func (o *MemberThreatprotectionOutboundSettings) HasQueryFqdnLimit() bool`

HasQueryFqdnLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


