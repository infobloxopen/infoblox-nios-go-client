# SyslogEndpointSyslogServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Syslog Server IP address | [optional] 
**ConnectionType** | Pointer to **string** | Connection type values | [optional] 
**Port** | Pointer to **int64** | The port this server listens on. | [optional] 
**Hostname** | Pointer to **string** | List of hostnames | [optional] 
**Format** | Pointer to **string** | Format vlues for syslog endpoint server | [optional] 
**Facility** | Pointer to **string** | Facility values for syslog endpoint server | [optional] 
**Severity** | Pointer to **string** | Severity values for syslog endpoint server. | [optional] 
**Certificate** | Pointer to **string** | Reference for creating sysog endpoint server. | [optional] [readonly] 
**CertificateToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop. | [optional] 

## Methods

### NewSyslogEndpointSyslogServers

`func NewSyslogEndpointSyslogServers() *SyslogEndpointSyslogServers`

NewSyslogEndpointSyslogServers instantiates a new SyslogEndpointSyslogServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogEndpointSyslogServersWithDefaults

`func NewSyslogEndpointSyslogServersWithDefaults() *SyslogEndpointSyslogServers`

NewSyslogEndpointSyslogServersWithDefaults instantiates a new SyslogEndpointSyslogServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SyslogEndpointSyslogServers) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SyslogEndpointSyslogServers) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SyslogEndpointSyslogServers) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SyslogEndpointSyslogServers) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetConnectionType

`func (o *SyslogEndpointSyslogServers) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *SyslogEndpointSyslogServers) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *SyslogEndpointSyslogServers) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *SyslogEndpointSyslogServers) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetPort

`func (o *SyslogEndpointSyslogServers) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SyslogEndpointSyslogServers) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SyslogEndpointSyslogServers) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SyslogEndpointSyslogServers) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetHostname

`func (o *SyslogEndpointSyslogServers) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SyslogEndpointSyslogServers) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SyslogEndpointSyslogServers) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SyslogEndpointSyslogServers) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetFormat

`func (o *SyslogEndpointSyslogServers) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SyslogEndpointSyslogServers) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SyslogEndpointSyslogServers) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SyslogEndpointSyslogServers) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetFacility

`func (o *SyslogEndpointSyslogServers) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *SyslogEndpointSyslogServers) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *SyslogEndpointSyslogServers) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *SyslogEndpointSyslogServers) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetSeverity

`func (o *SyslogEndpointSyslogServers) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SyslogEndpointSyslogServers) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SyslogEndpointSyslogServers) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SyslogEndpointSyslogServers) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCertificate

`func (o *SyslogEndpointSyslogServers) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SyslogEndpointSyslogServers) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SyslogEndpointSyslogServers) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SyslogEndpointSyslogServers) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateToken

`func (o *SyslogEndpointSyslogServers) GetCertificateToken() string`

GetCertificateToken returns the CertificateToken field if non-nil, zero value otherwise.

### GetCertificateTokenOk

`func (o *SyslogEndpointSyslogServers) GetCertificateTokenOk() (*string, bool)`

GetCertificateTokenOk returns a tuple with the CertificateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateToken

`func (o *SyslogEndpointSyslogServers) SetCertificateToken(v string)`

SetCertificateToken sets CertificateToken field to given value.

### HasCertificateToken

`func (o *SyslogEndpointSyslogServers) HasCertificateToken() bool`

HasCertificateToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


