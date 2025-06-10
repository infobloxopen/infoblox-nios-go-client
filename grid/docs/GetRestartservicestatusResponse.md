# GetRestartservicestatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**DhcpStatus** | Pointer to **string** | The status of the DHCP service. | [optional] [readonly] 
**DnsStatus** | Pointer to **string** | The status of the DNS service. | [optional] [readonly] 
**Member** | Pointer to **string** | The name of this Grid member in FQDN format. | [optional] [readonly] 
**ReportingStatus** | Pointer to **string** | The status of the reporting service. | [optional] [readonly] 
**Result** | Pointer to [**Restartservicestatus**](Restartservicestatus.md) |  | [optional] 

## Methods

### NewGetRestartservicestatusResponse

`func NewGetRestartservicestatusResponse() *GetRestartservicestatusResponse`

NewGetRestartservicestatusResponse instantiates a new GetRestartservicestatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRestartservicestatusResponseWithDefaults

`func NewGetRestartservicestatusResponseWithDefaults() *GetRestartservicestatusResponse`

NewGetRestartservicestatusResponseWithDefaults instantiates a new GetRestartservicestatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRestartservicestatusResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRestartservicestatusResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRestartservicestatusResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRestartservicestatusResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDhcpStatus

`func (o *GetRestartservicestatusResponse) GetDhcpStatus() string`

GetDhcpStatus returns the DhcpStatus field if non-nil, zero value otherwise.

### GetDhcpStatusOk

`func (o *GetRestartservicestatusResponse) GetDhcpStatusOk() (*string, bool)`

GetDhcpStatusOk returns a tuple with the DhcpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpStatus

`func (o *GetRestartservicestatusResponse) SetDhcpStatus(v string)`

SetDhcpStatus sets DhcpStatus field to given value.

### HasDhcpStatus

`func (o *GetRestartservicestatusResponse) HasDhcpStatus() bool`

HasDhcpStatus returns a boolean if a field has been set.

### GetDnsStatus

`func (o *GetRestartservicestatusResponse) GetDnsStatus() string`

GetDnsStatus returns the DnsStatus field if non-nil, zero value otherwise.

### GetDnsStatusOk

`func (o *GetRestartservicestatusResponse) GetDnsStatusOk() (*string, bool)`

GetDnsStatusOk returns a tuple with the DnsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsStatus

`func (o *GetRestartservicestatusResponse) SetDnsStatus(v string)`

SetDnsStatus sets DnsStatus field to given value.

### HasDnsStatus

`func (o *GetRestartservicestatusResponse) HasDnsStatus() bool`

HasDnsStatus returns a boolean if a field has been set.

### GetMember

`func (o *GetRestartservicestatusResponse) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetRestartservicestatusResponse) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetRestartservicestatusResponse) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetRestartservicestatusResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetReportingStatus

`func (o *GetRestartservicestatusResponse) GetReportingStatus() string`

GetReportingStatus returns the ReportingStatus field if non-nil, zero value otherwise.

### GetReportingStatusOk

`func (o *GetRestartservicestatusResponse) GetReportingStatusOk() (*string, bool)`

GetReportingStatusOk returns a tuple with the ReportingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingStatus

`func (o *GetRestartservicestatusResponse) SetReportingStatus(v string)`

SetReportingStatus sets ReportingStatus field to given value.

### HasReportingStatus

`func (o *GetRestartservicestatusResponse) HasReportingStatus() bool`

HasReportingStatus returns a boolean if a field has been set.

### GetResult

`func (o *GetRestartservicestatusResponse) GetResult() Restartservicestatus`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRestartservicestatusResponse) GetResultOk() (*Restartservicestatus, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRestartservicestatusResponse) SetResult(v Restartservicestatus)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRestartservicestatusResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


