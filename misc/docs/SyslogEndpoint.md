# SyslogEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LogLevel** | Pointer to **string** | The log level for a notification REST endpoint. | [optional] 
**Name** | Pointer to **string** | The name of a Syslog endpoint. | [optional] 
**OutboundMemberType** | Pointer to **string** | The outbound member that will generate events. | [optional] 
**OutboundMembers** | Pointer to **[]string** | The list of members for outbound events. | [optional] 
**SyslogServers** | Pointer to [**[]SyslogEndpointSyslogServers**](SyslogEndpointSyslogServers.md) | List of syslog servers | [optional] 
**TemplateInstance** | Pointer to [**SyslogEndpointTemplateInstance**](SyslogEndpointTemplateInstance.md) |  | [optional] 
**Timeout** | Pointer to **int64** | The timeout of session management (in seconds). | [optional] 
**VendorIdentifier** | Pointer to **string** | The vendor identifier. | [optional] 
**WapiUserName** | Pointer to **string** | The user name for WAPI integration. | [optional] 
**WapiUserPassword** | Pointer to **string** | The user password for WAPI integration. | [optional] 

## Methods

### NewSyslogEndpoint

`func NewSyslogEndpoint() *SyslogEndpoint`

NewSyslogEndpoint instantiates a new SyslogEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyslogEndpointWithDefaults

`func NewSyslogEndpointWithDefaults() *SyslogEndpoint`

NewSyslogEndpointWithDefaults instantiates a new SyslogEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *SyslogEndpoint) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *SyslogEndpoint) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *SyslogEndpoint) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *SyslogEndpoint) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetExtattrs

`func (o *SyslogEndpoint) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *SyslogEndpoint) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *SyslogEndpoint) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *SyslogEndpoint) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetLogLevel

`func (o *SyslogEndpoint) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *SyslogEndpoint) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *SyslogEndpoint) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *SyslogEndpoint) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetName

`func (o *SyslogEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyslogEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyslogEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SyslogEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundMemberType

`func (o *SyslogEndpoint) GetOutboundMemberType() string`

GetOutboundMemberType returns the OutboundMemberType field if non-nil, zero value otherwise.

### GetOutboundMemberTypeOk

`func (o *SyslogEndpoint) GetOutboundMemberTypeOk() (*string, bool)`

GetOutboundMemberTypeOk returns a tuple with the OutboundMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMemberType

`func (o *SyslogEndpoint) SetOutboundMemberType(v string)`

SetOutboundMemberType sets OutboundMemberType field to given value.

### HasOutboundMemberType

`func (o *SyslogEndpoint) HasOutboundMemberType() bool`

HasOutboundMemberType returns a boolean if a field has been set.

### GetOutboundMembers

`func (o *SyslogEndpoint) GetOutboundMembers() []string`

GetOutboundMembers returns the OutboundMembers field if non-nil, zero value otherwise.

### GetOutboundMembersOk

`func (o *SyslogEndpoint) GetOutboundMembersOk() (*[]string, bool)`

GetOutboundMembersOk returns a tuple with the OutboundMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMembers

`func (o *SyslogEndpoint) SetOutboundMembers(v []string)`

SetOutboundMembers sets OutboundMembers field to given value.

### HasOutboundMembers

`func (o *SyslogEndpoint) HasOutboundMembers() bool`

HasOutboundMembers returns a boolean if a field has been set.

### GetSyslogServers

`func (o *SyslogEndpoint) GetSyslogServers() []SyslogEndpointSyslogServers`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *SyslogEndpoint) GetSyslogServersOk() (*[]SyslogEndpointSyslogServers, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *SyslogEndpoint) SetSyslogServers(v []SyslogEndpointSyslogServers)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *SyslogEndpoint) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *SyslogEndpoint) GetTemplateInstance() SyslogEndpointTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *SyslogEndpoint) GetTemplateInstanceOk() (*SyslogEndpointTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *SyslogEndpoint) SetTemplateInstance(v SyslogEndpointTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *SyslogEndpoint) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetTimeout

`func (o *SyslogEndpoint) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SyslogEndpoint) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SyslogEndpoint) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SyslogEndpoint) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *SyslogEndpoint) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *SyslogEndpoint) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *SyslogEndpoint) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *SyslogEndpoint) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetWapiUserName

`func (o *SyslogEndpoint) GetWapiUserName() string`

GetWapiUserName returns the WapiUserName field if non-nil, zero value otherwise.

### GetWapiUserNameOk

`func (o *SyslogEndpoint) GetWapiUserNameOk() (*string, bool)`

GetWapiUserNameOk returns a tuple with the WapiUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserName

`func (o *SyslogEndpoint) SetWapiUserName(v string)`

SetWapiUserName sets WapiUserName field to given value.

### HasWapiUserName

`func (o *SyslogEndpoint) HasWapiUserName() bool`

HasWapiUserName returns a boolean if a field has been set.

### GetWapiUserPassword

`func (o *SyslogEndpoint) GetWapiUserPassword() string`

GetWapiUserPassword returns the WapiUserPassword field if non-nil, zero value otherwise.

### GetWapiUserPasswordOk

`func (o *SyslogEndpoint) GetWapiUserPasswordOk() (*string, bool)`

GetWapiUserPasswordOk returns a tuple with the WapiUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserPassword

`func (o *SyslogEndpoint) SetWapiUserPassword(v string)`

SetWapiUserPassword sets WapiUserPassword field to given value.

### HasWapiUserPassword

`func (o *SyslogEndpoint) HasWapiUserPassword() bool`

HasWapiUserPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


