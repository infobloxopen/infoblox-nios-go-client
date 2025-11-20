# Restartservicestatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**DhcpStatus** | Pointer to **string** | The status of the DHCP service. | [optional] [readonly] 
**DnsStatus** | Pointer to **string** | The status of the DNS service. | [optional] [readonly] 
**Member** | Pointer to **string** | The name of this Grid member in FQDN format. | [optional] [readonly] 
**ReportingStatus** | Pointer to **string** | The status of the reporting service. | [optional] [readonly] 

## Methods

### NewRestartservicestatus

`func NewRestartservicestatus() *Restartservicestatus`

NewRestartservicestatus instantiates a new Restartservicestatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartservicestatusWithDefaults

`func NewRestartservicestatusWithDefaults() *Restartservicestatus`

NewRestartservicestatusWithDefaults instantiates a new Restartservicestatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Restartservicestatus) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Restartservicestatus) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Restartservicestatus) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Restartservicestatus) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDhcpStatus

`func (o *Restartservicestatus) GetDhcpStatus() string`

GetDhcpStatus returns the DhcpStatus field if non-nil, zero value otherwise.

### GetDhcpStatusOk

`func (o *Restartservicestatus) GetDhcpStatusOk() (*string, bool)`

GetDhcpStatusOk returns a tuple with the DhcpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpStatus

`func (o *Restartservicestatus) SetDhcpStatus(v string)`

SetDhcpStatus sets DhcpStatus field to given value.

### HasDhcpStatus

`func (o *Restartservicestatus) HasDhcpStatus() bool`

HasDhcpStatus returns a boolean if a field has been set.

### GetDnsStatus

`func (o *Restartservicestatus) GetDnsStatus() string`

GetDnsStatus returns the DnsStatus field if non-nil, zero value otherwise.

### GetDnsStatusOk

`func (o *Restartservicestatus) GetDnsStatusOk() (*string, bool)`

GetDnsStatusOk returns a tuple with the DnsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsStatus

`func (o *Restartservicestatus) SetDnsStatus(v string)`

SetDnsStatus sets DnsStatus field to given value.

### HasDnsStatus

`func (o *Restartservicestatus) HasDnsStatus() bool`

HasDnsStatus returns a boolean if a field has been set.

### GetMember

`func (o *Restartservicestatus) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Restartservicestatus) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Restartservicestatus) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *Restartservicestatus) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetReportingStatus

`func (o *Restartservicestatus) GetReportingStatus() string`

GetReportingStatus returns the ReportingStatus field if non-nil, zero value otherwise.

### GetReportingStatusOk

`func (o *Restartservicestatus) GetReportingStatusOk() (*string, bool)`

GetReportingStatusOk returns a tuple with the ReportingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingStatus

`func (o *Restartservicestatus) SetReportingStatus(v string)`

SetReportingStatus sets ReportingStatus field to given value.

### HasReportingStatus

`func (o *Restartservicestatus) HasReportingStatus() bool`

HasReportingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


