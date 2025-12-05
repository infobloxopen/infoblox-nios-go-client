# MsserverDns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The address or FQDN of the DNS Microsoft Server. | [optional] [readonly] 
**EnableDnsReportsSync** | Pointer to **bool** | Determines if synchronization of DNS reporting data from the Microsoft server is enabled or not. | [optional] 
**LoginName** | Pointer to **string** | The login name of the DNS Microsoft Server. | [optional] 
**LoginPassword** | Pointer to **string** | The login password of the DNS Microsoft Server. | [optional] 
**SynchronizationInterval** | Pointer to **int64** | The minimum number of minutes between two synchronizations. | [optional] 
**UseEnableDnsReportsSync** | Pointer to **bool** | Use flag for: enable_dns_reports_sync | [optional] 
**UseLogin** | Pointer to **bool** | Use flag for: login_name , login_password | [optional] 
**UseSynchronizationInterval** | Pointer to **bool** | Use flag for: synchronization_interval | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewMsserverDns

`func NewMsserverDns() *MsserverDns`

NewMsserverDns instantiates a new MsserverDns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsserverDnsWithDefaults

`func NewMsserverDnsWithDefaults() *MsserverDns`

NewMsserverDnsWithDefaults instantiates a new MsserverDns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MsserverDns) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MsserverDns) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MsserverDns) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MsserverDns) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *MsserverDns) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MsserverDns) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MsserverDns) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MsserverDns) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEnableDnsReportsSync

`func (o *MsserverDns) GetEnableDnsReportsSync() bool`

GetEnableDnsReportsSync returns the EnableDnsReportsSync field if non-nil, zero value otherwise.

### GetEnableDnsReportsSyncOk

`func (o *MsserverDns) GetEnableDnsReportsSyncOk() (*bool, bool)`

GetEnableDnsReportsSyncOk returns a tuple with the EnableDnsReportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDnsReportsSync

`func (o *MsserverDns) SetEnableDnsReportsSync(v bool)`

SetEnableDnsReportsSync sets EnableDnsReportsSync field to given value.

### HasEnableDnsReportsSync

`func (o *MsserverDns) HasEnableDnsReportsSync() bool`

HasEnableDnsReportsSync returns a boolean if a field has been set.

### GetLoginName

`func (o *MsserverDns) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *MsserverDns) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *MsserverDns) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *MsserverDns) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLoginPassword

`func (o *MsserverDns) GetLoginPassword() string`

GetLoginPassword returns the LoginPassword field if non-nil, zero value otherwise.

### GetLoginPasswordOk

`func (o *MsserverDns) GetLoginPasswordOk() (*string, bool)`

GetLoginPasswordOk returns a tuple with the LoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPassword

`func (o *MsserverDns) SetLoginPassword(v string)`

SetLoginPassword sets LoginPassword field to given value.

### HasLoginPassword

`func (o *MsserverDns) HasLoginPassword() bool`

HasLoginPassword returns a boolean if a field has been set.

### GetSynchronizationInterval

`func (o *MsserverDns) GetSynchronizationInterval() int64`

GetSynchronizationInterval returns the SynchronizationInterval field if non-nil, zero value otherwise.

### GetSynchronizationIntervalOk

`func (o *MsserverDns) GetSynchronizationIntervalOk() (*int64, bool)`

GetSynchronizationIntervalOk returns a tuple with the SynchronizationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationInterval

`func (o *MsserverDns) SetSynchronizationInterval(v int64)`

SetSynchronizationInterval sets SynchronizationInterval field to given value.

### HasSynchronizationInterval

`func (o *MsserverDns) HasSynchronizationInterval() bool`

HasSynchronizationInterval returns a boolean if a field has been set.

### GetUseEnableDnsReportsSync

`func (o *MsserverDns) GetUseEnableDnsReportsSync() bool`

GetUseEnableDnsReportsSync returns the UseEnableDnsReportsSync field if non-nil, zero value otherwise.

### GetUseEnableDnsReportsSyncOk

`func (o *MsserverDns) GetUseEnableDnsReportsSyncOk() (*bool, bool)`

GetUseEnableDnsReportsSyncOk returns a tuple with the UseEnableDnsReportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableDnsReportsSync

`func (o *MsserverDns) SetUseEnableDnsReportsSync(v bool)`

SetUseEnableDnsReportsSync sets UseEnableDnsReportsSync field to given value.

### HasUseEnableDnsReportsSync

`func (o *MsserverDns) HasUseEnableDnsReportsSync() bool`

HasUseEnableDnsReportsSync returns a boolean if a field has been set.

### GetUseLogin

`func (o *MsserverDns) GetUseLogin() bool`

GetUseLogin returns the UseLogin field if non-nil, zero value otherwise.

### GetUseLoginOk

`func (o *MsserverDns) GetUseLoginOk() (*bool, bool)`

GetUseLoginOk returns a tuple with the UseLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogin

`func (o *MsserverDns) SetUseLogin(v bool)`

SetUseLogin sets UseLogin field to given value.

### HasUseLogin

`func (o *MsserverDns) HasUseLogin() bool`

HasUseLogin returns a boolean if a field has been set.

### GetUseSynchronizationInterval

`func (o *MsserverDns) GetUseSynchronizationInterval() bool`

GetUseSynchronizationInterval returns the UseSynchronizationInterval field if non-nil, zero value otherwise.

### GetUseSynchronizationIntervalOk

`func (o *MsserverDns) GetUseSynchronizationIntervalOk() (*bool, bool)`

GetUseSynchronizationIntervalOk returns a tuple with the UseSynchronizationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynchronizationInterval

`func (o *MsserverDns) SetUseSynchronizationInterval(v bool)`

SetUseSynchronizationInterval sets UseSynchronizationInterval field to given value.

### HasUseSynchronizationInterval

`func (o *MsserverDns) HasUseSynchronizationInterval() bool`

HasUseSynchronizationInterval returns a boolean if a field has been set.

### GetUuid

`func (o *MsserverDns) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MsserverDns) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MsserverDns) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MsserverDns) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


