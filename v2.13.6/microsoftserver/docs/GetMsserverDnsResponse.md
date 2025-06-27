# GetMsserverDnsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Address** | Pointer to **string** | The address or FQDN of the DNS Microsoft Server. | [optional] [readonly] 
**EnableDnsReportsSync** | Pointer to **bool** | Determines if synchronization of DNS reporting data from the Microsoft server is enabled or not. | [optional] 
**LoginName** | Pointer to **string** | The login name of the DNS Microsoft Server. | [optional] 
**LoginPassword** | Pointer to **string** | The login password of the DNS Microsoft Server. | [optional] 
**SynchronizationInterval** | Pointer to **int64** | The minimum number of minutes between two synchronizations. | [optional] 
**UseEnableDnsReportsSync** | Pointer to **bool** | Use flag for: enable_dns_reports_sync | [optional] 
**UseLogin** | Pointer to **bool** | Use flag for: login_name , login_password | [optional] 
**UseSynchronizationInterval** | Pointer to **bool** | Use flag for: synchronization_interval | [optional] 
**Result** | Pointer to [**MsserverDns**](MsserverDns.md) |  | [optional] 

## Methods

### NewGetMsserverDnsResponse

`func NewGetMsserverDnsResponse() *GetMsserverDnsResponse`

NewGetMsserverDnsResponse instantiates a new GetMsserverDnsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMsserverDnsResponseWithDefaults

`func NewGetMsserverDnsResponseWithDefaults() *GetMsserverDnsResponse`

NewGetMsserverDnsResponseWithDefaults instantiates a new GetMsserverDnsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMsserverDnsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMsserverDnsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMsserverDnsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMsserverDnsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetMsserverDnsResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetMsserverDnsResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetMsserverDnsResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetMsserverDnsResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEnableDnsReportsSync

`func (o *GetMsserverDnsResponse) GetEnableDnsReportsSync() bool`

GetEnableDnsReportsSync returns the EnableDnsReportsSync field if non-nil, zero value otherwise.

### GetEnableDnsReportsSyncOk

`func (o *GetMsserverDnsResponse) GetEnableDnsReportsSyncOk() (*bool, bool)`

GetEnableDnsReportsSyncOk returns a tuple with the EnableDnsReportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsReportsSync

`func (o *GetMsserverDnsResponse) SetEnableDnsReportsSync(v bool)`

SetEnableDnsReportsSync sets EnableDnsReportsSync field to given value.

### HasEnableDnsReportsSync

`func (o *GetMsserverDnsResponse) HasEnableDnsReportsSync() bool`

HasEnableDnsReportsSync returns a boolean if a field has been set.

### GetLoginName

`func (o *GetMsserverDnsResponse) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *GetMsserverDnsResponse) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *GetMsserverDnsResponse) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *GetMsserverDnsResponse) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLoginPassword

`func (o *GetMsserverDnsResponse) GetLoginPassword() string`

GetLoginPassword returns the LoginPassword field if non-nil, zero value otherwise.

### GetLoginPasswordOk

`func (o *GetMsserverDnsResponse) GetLoginPasswordOk() (*string, bool)`

GetLoginPasswordOk returns a tuple with the LoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPassword

`func (o *GetMsserverDnsResponse) SetLoginPassword(v string)`

SetLoginPassword sets LoginPassword field to given value.

### HasLoginPassword

`func (o *GetMsserverDnsResponse) HasLoginPassword() bool`

HasLoginPassword returns a boolean if a field has been set.

### GetSynchronizationInterval

`func (o *GetMsserverDnsResponse) GetSynchronizationInterval() int64`

GetSynchronizationInterval returns the SynchronizationInterval field if non-nil, zero value otherwise.

### GetSynchronizationIntervalOk

`func (o *GetMsserverDnsResponse) GetSynchronizationIntervalOk() (*int64, bool)`

GetSynchronizationIntervalOk returns a tuple with the SynchronizationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationInterval

`func (o *GetMsserverDnsResponse) SetSynchronizationInterval(v int64)`

SetSynchronizationInterval sets SynchronizationInterval field to given value.

### HasSynchronizationInterval

`func (o *GetMsserverDnsResponse) HasSynchronizationInterval() bool`

HasSynchronizationInterval returns a boolean if a field has been set.

### GetUseEnableDnsReportsSync

`func (o *GetMsserverDnsResponse) GetUseEnableDnsReportsSync() bool`

GetUseEnableDnsReportsSync returns the UseEnableDnsReportsSync field if non-nil, zero value otherwise.

### GetUseEnableDnsReportsSyncOk

`func (o *GetMsserverDnsResponse) GetUseEnableDnsReportsSyncOk() (*bool, bool)`

GetUseEnableDnsReportsSyncOk returns a tuple with the UseEnableDnsReportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDnsReportsSync

`func (o *GetMsserverDnsResponse) SetUseEnableDnsReportsSync(v bool)`

SetUseEnableDnsReportsSync sets UseEnableDnsReportsSync field to given value.

### HasUseEnableDnsReportsSync

`func (o *GetMsserverDnsResponse) HasUseEnableDnsReportsSync() bool`

HasUseEnableDnsReportsSync returns a boolean if a field has been set.

### GetUseLogin

`func (o *GetMsserverDnsResponse) GetUseLogin() bool`

GetUseLogin returns the UseLogin field if non-nil, zero value otherwise.

### GetUseLoginOk

`func (o *GetMsserverDnsResponse) GetUseLoginOk() (*bool, bool)`

GetUseLoginOk returns a tuple with the UseLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogin

`func (o *GetMsserverDnsResponse) SetUseLogin(v bool)`

SetUseLogin sets UseLogin field to given value.

### HasUseLogin

`func (o *GetMsserverDnsResponse) HasUseLogin() bool`

HasUseLogin returns a boolean if a field has been set.

### GetUseSynchronizationInterval

`func (o *GetMsserverDnsResponse) GetUseSynchronizationInterval() bool`

GetUseSynchronizationInterval returns the UseSynchronizationInterval field if non-nil, zero value otherwise.

### GetUseSynchronizationIntervalOk

`func (o *GetMsserverDnsResponse) GetUseSynchronizationIntervalOk() (*bool, bool)`

GetUseSynchronizationIntervalOk returns a tuple with the UseSynchronizationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynchronizationInterval

`func (o *GetMsserverDnsResponse) SetUseSynchronizationInterval(v bool)`

SetUseSynchronizationInterval sets UseSynchronizationInterval field to given value.

### HasUseSynchronizationInterval

`func (o *GetMsserverDnsResponse) HasUseSynchronizationInterval() bool`

HasUseSynchronizationInterval returns a boolean if a field has been set.

### GetResult

`func (o *GetMsserverDnsResponse) GetResult() MsserverDns`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMsserverDnsResponse) GetResultOk() (*MsserverDns, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMsserverDnsResponse) SetResult(v MsserverDns)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMsserverDnsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


