# GridSyslogServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressOrFqdn** | Pointer to **string** | The server address or FQDN. | [optional] 
**Certificate** | Pointer to **string** | Reference to the underlying X509Certificate object grid:x509certificate. | [optional] [readonly] 
**CertificateToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop. | [optional] 
**ConnectionType** | Pointer to **string** | The connection type for communicating with this server. | [optional] 
**Port** | Pointer to **int64** | The port this server listens on. | [optional] 
**LocalInterface** | Pointer to **string** | The local interface through which the appliance sends syslog messages to the syslog server. | [optional] 
**MessageSource** | Pointer to **string** | The source of syslog messages to be sent to the external syslog server. If set to &#39;INTERNAL&#39;, only messages the appliance generates will be sent to the syslog server. If set to &#39;EXTERNAL&#39;, the appliance sends syslog messages that it receives from other devices, such as syslog servers and routers. If set to &#39;ANY&#39;, the appliance sends both internal and external syslog messages. | [optional] 
**MessageNodeId** | Pointer to **string** | Identify the node in the syslog message. | [optional] 
**Severity** | Pointer to **string** | The severity filter. The appliance sends log messages of the specified severity and above to the external syslog server. | [optional] 
**CategoryList** | Pointer to **[]string** | The list of all syslog logging categories. | [optional] 
**OnlyCategoryList** | Pointer to **bool** | The list of selected syslog logging categories. The appliance forwards syslog messages that belong to the selected categories. | [optional] 

## Methods

### NewGridSyslogServers

`func NewGridSyslogServers() *GridSyslogServers`

NewGridSyslogServers instantiates a new GridSyslogServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridSyslogServersWithDefaults

`func NewGridSyslogServersWithDefaults() *GridSyslogServers`

NewGridSyslogServersWithDefaults instantiates a new GridSyslogServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressOrFqdn

`func (o *GridSyslogServers) GetAddressOrFqdn() string`

GetAddressOrFqdn returns the AddressOrFqdn field if non-nil, zero value otherwise.

### GetAddressOrFqdnOk

`func (o *GridSyslogServers) GetAddressOrFqdnOk() (*string, bool)`

GetAddressOrFqdnOk returns a tuple with the AddressOrFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressOrFqdn

`func (o *GridSyslogServers) SetAddressOrFqdn(v string)`

SetAddressOrFqdn sets AddressOrFqdn field to given value.

### HasAddressOrFqdn

`func (o *GridSyslogServers) HasAddressOrFqdn() bool`

HasAddressOrFqdn returns a boolean if a field has been set.

### GetCertificate

`func (o *GridSyslogServers) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *GridSyslogServers) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *GridSyslogServers) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *GridSyslogServers) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateToken

`func (o *GridSyslogServers) GetCertificateToken() string`

GetCertificateToken returns the CertificateToken field if non-nil, zero value otherwise.

### GetCertificateTokenOk

`func (o *GridSyslogServers) GetCertificateTokenOk() (*string, bool)`

GetCertificateTokenOk returns a tuple with the CertificateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateToken

`func (o *GridSyslogServers) SetCertificateToken(v string)`

SetCertificateToken sets CertificateToken field to given value.

### HasCertificateToken

`func (o *GridSyslogServers) HasCertificateToken() bool`

HasCertificateToken returns a boolean if a field has been set.

### GetConnectionType

`func (o *GridSyslogServers) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *GridSyslogServers) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *GridSyslogServers) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *GridSyslogServers) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetPort

`func (o *GridSyslogServers) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GridSyslogServers) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GridSyslogServers) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GridSyslogServers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetLocalInterface

`func (o *GridSyslogServers) GetLocalInterface() string`

GetLocalInterface returns the LocalInterface field if non-nil, zero value otherwise.

### GetLocalInterfaceOk

`func (o *GridSyslogServers) GetLocalInterfaceOk() (*string, bool)`

GetLocalInterfaceOk returns a tuple with the LocalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInterface

`func (o *GridSyslogServers) SetLocalInterface(v string)`

SetLocalInterface sets LocalInterface field to given value.

### HasLocalInterface

`func (o *GridSyslogServers) HasLocalInterface() bool`

HasLocalInterface returns a boolean if a field has been set.

### GetMessageSource

`func (o *GridSyslogServers) GetMessageSource() string`

GetMessageSource returns the MessageSource field if non-nil, zero value otherwise.

### GetMessageSourceOk

`func (o *GridSyslogServers) GetMessageSourceOk() (*string, bool)`

GetMessageSourceOk returns a tuple with the MessageSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSource

`func (o *GridSyslogServers) SetMessageSource(v string)`

SetMessageSource sets MessageSource field to given value.

### HasMessageSource

`func (o *GridSyslogServers) HasMessageSource() bool`

HasMessageSource returns a boolean if a field has been set.

### GetMessageNodeId

`func (o *GridSyslogServers) GetMessageNodeId() string`

GetMessageNodeId returns the MessageNodeId field if non-nil, zero value otherwise.

### GetMessageNodeIdOk

`func (o *GridSyslogServers) GetMessageNodeIdOk() (*string, bool)`

GetMessageNodeIdOk returns a tuple with the MessageNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageNodeId

`func (o *GridSyslogServers) SetMessageNodeId(v string)`

SetMessageNodeId sets MessageNodeId field to given value.

### HasMessageNodeId

`func (o *GridSyslogServers) HasMessageNodeId() bool`

HasMessageNodeId returns a boolean if a field has been set.

### GetSeverity

`func (o *GridSyslogServers) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GridSyslogServers) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GridSyslogServers) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GridSyslogServers) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCategoryList

`func (o *GridSyslogServers) GetCategoryList() []string`

GetCategoryList returns the CategoryList field if non-nil, zero value otherwise.

### GetCategoryListOk

`func (o *GridSyslogServers) GetCategoryListOk() (*[]string, bool)`

GetCategoryListOk returns a tuple with the CategoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryList

`func (o *GridSyslogServers) SetCategoryList(v []string)`

SetCategoryList sets CategoryList field to given value.

### HasCategoryList

`func (o *GridSyslogServers) HasCategoryList() bool`

HasCategoryList returns a boolean if a field has been set.

### GetOnlyCategoryList

`func (o *GridSyslogServers) GetOnlyCategoryList() bool`

GetOnlyCategoryList returns the OnlyCategoryList field if non-nil, zero value otherwise.

### GetOnlyCategoryListOk

`func (o *GridSyslogServers) GetOnlyCategoryListOk() (*bool, bool)`

GetOnlyCategoryListOk returns a tuple with the OnlyCategoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyCategoryList

`func (o *GridSyslogServers) SetOnlyCategoryList(v bool)`

SetOnlyCategoryList sets OnlyCategoryList field to given value.

### HasOnlyCategoryList

`func (o *GridSyslogServers) HasOnlyCategoryList() bool`

HasOnlyCategoryList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


