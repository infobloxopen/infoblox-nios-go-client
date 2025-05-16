# DiscoveryGridpropertiesAdvisorSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableProxy** | Pointer to **bool** | Use proxy server if true. | [optional] 
**ProxyAddress** | Pointer to **string** | Host name or ip address of the proxy server. | [optional] 
**ProxyPort** | Pointer to **int64** | Port number the proxy listens on. | [optional] 
**UseProxyUsernamePasswd** | Pointer to **bool** | Is proxy authorization required? | [optional] 
**ProxyUsername** | Pointer to **string** | Proxy username. | [optional] 
**ProxyPassword** | Pointer to **string** | Proxy password. | [optional] 
**ExecutionInterval** | Pointer to **int64** | Application execution interval in seconds. Must be greater than or equal to 86400 seconds (1 day). | [optional] 
**ExecutionHour** | Pointer to **int64** | Application execution hour. | [optional] 
**NetworkInterfaceType** | Pointer to **string** | The type of the network interface on consolidator member. | [optional] 
**NetworkInterfaceVirtualIp** | Pointer to **string** | The interface for Advisor application on consolidator member. | [optional] 
**Address** | Pointer to **string** | Host name or ip address of the Advisor server | [optional] 
**Port** | Pointer to **int64** | Port number the Advisor server listens on | [optional] 
**AuthType** | Pointer to **string** | Authentication type used for Advisor server | [optional] 
**AuthToken** | Pointer to **string** | Advisor AUTH token | [optional] 
**Username** | Pointer to **string** | Username for Advisor server | [optional] 
**Password** | Pointer to **string** | Password for Advisor server | [optional] 
**MinSeverity** | Pointer to **string** | Advisor minimal severity | [optional] 
**LastExecTime** | Pointer to **int64** | Timestamp of the Advisor last execution attempt&#39; | [optional] [readonly] 
**LastExecStatus** | Pointer to **string** | Result of the last execution attempt of Advisor | [optional] [readonly] 
**LastExecDetails** | Pointer to **string** | Details of the last execution attempt of Advisor. Describes the error or warning with a string | [optional] [readonly] 
**LastRunNowTime** | Pointer to **int64** | Timestamp of the Advisor last Run Now attempt&#39; | [optional] [readonly] 
**LastRunNowStatus** | Pointer to **string** | Result of the last Run Now attempt of Advisor | [optional] [readonly] 
**LastRunNowDetails** | Pointer to **string** | Details of the last Run Now attempt of Advisor. Describes the error or warning with a string | [optional] [readonly] 

## Methods

### NewDiscoveryGridpropertiesAdvisorSettings

`func NewDiscoveryGridpropertiesAdvisorSettings() *DiscoveryGridpropertiesAdvisorSettings`

NewDiscoveryGridpropertiesAdvisorSettings instantiates a new DiscoveryGridpropertiesAdvisorSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryGridpropertiesAdvisorSettingsWithDefaults

`func NewDiscoveryGridpropertiesAdvisorSettingsWithDefaults() *DiscoveryGridpropertiesAdvisorSettings`

NewDiscoveryGridpropertiesAdvisorSettingsWithDefaults instantiates a new DiscoveryGridpropertiesAdvisorSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableProxy

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetEnableProxy() bool`

GetEnableProxy returns the EnableProxy field if non-nil, zero value otherwise.

### GetEnableProxyOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetEnableProxyOk() (*bool, bool)`

GetEnableProxyOk returns a tuple with the EnableProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxy

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetEnableProxy(v bool)`

SetEnableProxy sets EnableProxy field to given value.

### HasEnableProxy

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasEnableProxy() bool`

HasEnableProxy returns a boolean if a field has been set.

### GetProxyAddress

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetProxyAddress() string`

GetProxyAddress returns the ProxyAddress field if non-nil, zero value otherwise.

### GetProxyAddressOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetProxyAddressOk() (*string, bool)`

GetProxyAddressOk returns a tuple with the ProxyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAddress

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetProxyAddress(v string)`

SetProxyAddress sets ProxyAddress field to given value.

### HasProxyAddress

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasProxyAddress() bool`

HasProxyAddress returns a boolean if a field has been set.

### GetProxyPort

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetProxyPort() int64`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetProxyPortOk() (*int64, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetProxyPort(v int64)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetUseProxyUsernamePasswd

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetUseProxyUsernamePasswd() bool`

GetUseProxyUsernamePasswd returns the UseProxyUsernamePasswd field if non-nil, zero value otherwise.

### GetUseProxyUsernamePasswdOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetUseProxyUsernamePasswdOk() (*bool, bool)`

GetUseProxyUsernamePasswdOk returns a tuple with the UseProxyUsernamePasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseProxyUsernamePasswd

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetUseProxyUsernamePasswd(v bool)`

SetUseProxyUsernamePasswd sets UseProxyUsernamePasswd field to given value.

### HasUseProxyUsernamePasswd

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasUseProxyUsernamePasswd() bool`

HasUseProxyUsernamePasswd returns a boolean if a field has been set.

### GetProxyUsername

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetProxyUsername() string`

GetProxyUsername returns the ProxyUsername field if non-nil, zero value otherwise.

### GetProxyUsernameOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetProxyUsernameOk() (*string, bool)`

GetProxyUsernameOk returns a tuple with the ProxyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUsername

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetProxyUsername(v string)`

SetProxyUsername sets ProxyUsername field to given value.

### HasProxyUsername

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasProxyUsername() bool`

HasProxyUsername returns a boolean if a field has been set.

### GetProxyPassword

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetExecutionInterval

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetExecutionInterval() int64`

GetExecutionInterval returns the ExecutionInterval field if non-nil, zero value otherwise.

### GetExecutionIntervalOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetExecutionIntervalOk() (*int64, bool)`

GetExecutionIntervalOk returns a tuple with the ExecutionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionInterval

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetExecutionInterval(v int64)`

SetExecutionInterval sets ExecutionInterval field to given value.

### HasExecutionInterval

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasExecutionInterval() bool`

HasExecutionInterval returns a boolean if a field has been set.

### GetExecutionHour

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetExecutionHour() int64`

GetExecutionHour returns the ExecutionHour field if non-nil, zero value otherwise.

### GetExecutionHourOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetExecutionHourOk() (*int64, bool)`

GetExecutionHourOk returns a tuple with the ExecutionHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionHour

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetExecutionHour(v int64)`

SetExecutionHour sets ExecutionHour field to given value.

### HasExecutionHour

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasExecutionHour() bool`

HasExecutionHour returns a boolean if a field has been set.

### GetNetworkInterfaceType

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetNetworkInterfaceType() string`

GetNetworkInterfaceType returns the NetworkInterfaceType field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetNetworkInterfaceTypeOk() (*string, bool)`

GetNetworkInterfaceTypeOk returns a tuple with the NetworkInterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceType

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetNetworkInterfaceType(v string)`

SetNetworkInterfaceType sets NetworkInterfaceType field to given value.

### HasNetworkInterfaceType

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasNetworkInterfaceType() bool`

HasNetworkInterfaceType returns a boolean if a field has been set.

### GetNetworkInterfaceVirtualIp

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetNetworkInterfaceVirtualIp() string`

GetNetworkInterfaceVirtualIp returns the NetworkInterfaceVirtualIp field if non-nil, zero value otherwise.

### GetNetworkInterfaceVirtualIpOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetNetworkInterfaceVirtualIpOk() (*string, bool)`

GetNetworkInterfaceVirtualIpOk returns a tuple with the NetworkInterfaceVirtualIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceVirtualIp

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetNetworkInterfaceVirtualIp(v string)`

SetNetworkInterfaceVirtualIp sets NetworkInterfaceVirtualIp field to given value.

### HasNetworkInterfaceVirtualIp

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasNetworkInterfaceVirtualIp() bool`

HasNetworkInterfaceVirtualIp returns a boolean if a field has been set.

### GetAddress

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAuthType

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthToken

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetUsername

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetMinSeverity

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetMinSeverity() string`

GetMinSeverity returns the MinSeverity field if non-nil, zero value otherwise.

### GetMinSeverityOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetMinSeverityOk() (*string, bool)`

GetMinSeverityOk returns a tuple with the MinSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeverity

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetMinSeverity(v string)`

SetMinSeverity sets MinSeverity field to given value.

### HasMinSeverity

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasMinSeverity() bool`

HasMinSeverity returns a boolean if a field has been set.

### GetLastExecTime

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastExecTime() int64`

GetLastExecTime returns the LastExecTime field if non-nil, zero value otherwise.

### GetLastExecTimeOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastExecTimeOk() (*int64, bool)`

GetLastExecTimeOk returns a tuple with the LastExecTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecTime

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetLastExecTime(v int64)`

SetLastExecTime sets LastExecTime field to given value.

### HasLastExecTime

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasLastExecTime() bool`

HasLastExecTime returns a boolean if a field has been set.

### GetLastExecStatus

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastExecStatus() string`

GetLastExecStatus returns the LastExecStatus field if non-nil, zero value otherwise.

### GetLastExecStatusOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastExecStatusOk() (*string, bool)`

GetLastExecStatusOk returns a tuple with the LastExecStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecStatus

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetLastExecStatus(v string)`

SetLastExecStatus sets LastExecStatus field to given value.

### HasLastExecStatus

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasLastExecStatus() bool`

HasLastExecStatus returns a boolean if a field has been set.

### GetLastExecDetails

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastExecDetails() string`

GetLastExecDetails returns the LastExecDetails field if non-nil, zero value otherwise.

### GetLastExecDetailsOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastExecDetailsOk() (*string, bool)`

GetLastExecDetailsOk returns a tuple with the LastExecDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecDetails

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetLastExecDetails(v string)`

SetLastExecDetails sets LastExecDetails field to given value.

### HasLastExecDetails

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasLastExecDetails() bool`

HasLastExecDetails returns a boolean if a field has been set.

### GetLastRunNowTime

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastRunNowTime() int64`

GetLastRunNowTime returns the LastRunNowTime field if non-nil, zero value otherwise.

### GetLastRunNowTimeOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastRunNowTimeOk() (*int64, bool)`

GetLastRunNowTimeOk returns a tuple with the LastRunNowTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunNowTime

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetLastRunNowTime(v int64)`

SetLastRunNowTime sets LastRunNowTime field to given value.

### HasLastRunNowTime

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasLastRunNowTime() bool`

HasLastRunNowTime returns a boolean if a field has been set.

### GetLastRunNowStatus

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastRunNowStatus() string`

GetLastRunNowStatus returns the LastRunNowStatus field if non-nil, zero value otherwise.

### GetLastRunNowStatusOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastRunNowStatusOk() (*string, bool)`

GetLastRunNowStatusOk returns a tuple with the LastRunNowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunNowStatus

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetLastRunNowStatus(v string)`

SetLastRunNowStatus sets LastRunNowStatus field to given value.

### HasLastRunNowStatus

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasLastRunNowStatus() bool`

HasLastRunNowStatus returns a boolean if a field has been set.

### GetLastRunNowDetails

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastRunNowDetails() string`

GetLastRunNowDetails returns the LastRunNowDetails field if non-nil, zero value otherwise.

### GetLastRunNowDetailsOk

`func (o *DiscoveryGridpropertiesAdvisorSettings) GetLastRunNowDetailsOk() (*string, bool)`

GetLastRunNowDetailsOk returns a tuple with the LastRunNowDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunNowDetails

`func (o *DiscoveryGridpropertiesAdvisorSettings) SetLastRunNowDetails(v string)`

SetLastRunNowDetails sets LastRunNowDetails field to given value.

### HasLastRunNowDetails

`func (o *DiscoveryGridpropertiesAdvisorSettings) HasLastRunNowDetails() bool`

HasLastRunNowDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


