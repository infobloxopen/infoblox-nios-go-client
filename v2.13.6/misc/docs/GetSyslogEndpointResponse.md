# GetSyslogEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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
**Result** | Pointer to [**SyslogEndpoint**](SyslogEndpoint.md) |  | [optional] 

## Methods

### NewGetSyslogEndpointResponse

`func NewGetSyslogEndpointResponse() *GetSyslogEndpointResponse`

NewGetSyslogEndpointResponse instantiates a new GetSyslogEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSyslogEndpointResponseWithDefaults

`func NewGetSyslogEndpointResponseWithDefaults() *GetSyslogEndpointResponse`

NewGetSyslogEndpointResponseWithDefaults instantiates a new GetSyslogEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSyslogEndpointResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSyslogEndpointResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSyslogEndpointResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSyslogEndpointResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetSyslogEndpointResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetSyslogEndpointResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetSyslogEndpointResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetSyslogEndpointResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLogLevel

`func (o *GetSyslogEndpointResponse) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *GetSyslogEndpointResponse) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *GetSyslogEndpointResponse) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *GetSyslogEndpointResponse) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetName

`func (o *GetSyslogEndpointResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSyslogEndpointResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSyslogEndpointResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSyslogEndpointResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundMemberType

`func (o *GetSyslogEndpointResponse) GetOutboundMemberType() string`

GetOutboundMemberType returns the OutboundMemberType field if non-nil, zero value otherwise.

### GetOutboundMemberTypeOk

`func (o *GetSyslogEndpointResponse) GetOutboundMemberTypeOk() (*string, bool)`

GetOutboundMemberTypeOk returns a tuple with the OutboundMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMemberType

`func (o *GetSyslogEndpointResponse) SetOutboundMemberType(v string)`

SetOutboundMemberType sets OutboundMemberType field to given value.

### HasOutboundMemberType

`func (o *GetSyslogEndpointResponse) HasOutboundMemberType() bool`

HasOutboundMemberType returns a boolean if a field has been set.

### GetOutboundMembers

`func (o *GetSyslogEndpointResponse) GetOutboundMembers() []string`

GetOutboundMembers returns the OutboundMembers field if non-nil, zero value otherwise.

### GetOutboundMembersOk

`func (o *GetSyslogEndpointResponse) GetOutboundMembersOk() (*[]string, bool)`

GetOutboundMembersOk returns a tuple with the OutboundMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMembers

`func (o *GetSyslogEndpointResponse) SetOutboundMembers(v []string)`

SetOutboundMembers sets OutboundMembers field to given value.

### HasOutboundMembers

`func (o *GetSyslogEndpointResponse) HasOutboundMembers() bool`

HasOutboundMembers returns a boolean if a field has been set.

### GetSyslogServers

`func (o *GetSyslogEndpointResponse) GetSyslogServers() []SyslogEndpointSyslogServers`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *GetSyslogEndpointResponse) GetSyslogServersOk() (*[]SyslogEndpointSyslogServers, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *GetSyslogEndpointResponse) SetSyslogServers(v []SyslogEndpointSyslogServers)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *GetSyslogEndpointResponse) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *GetSyslogEndpointResponse) GetTemplateInstance() SyslogEndpointTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *GetSyslogEndpointResponse) GetTemplateInstanceOk() (*SyslogEndpointTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *GetSyslogEndpointResponse) SetTemplateInstance(v SyslogEndpointTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *GetSyslogEndpointResponse) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetTimeout

`func (o *GetSyslogEndpointResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetSyslogEndpointResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetSyslogEndpointResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetSyslogEndpointResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *GetSyslogEndpointResponse) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *GetSyslogEndpointResponse) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *GetSyslogEndpointResponse) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *GetSyslogEndpointResponse) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetWapiUserName

`func (o *GetSyslogEndpointResponse) GetWapiUserName() string`

GetWapiUserName returns the WapiUserName field if non-nil, zero value otherwise.

### GetWapiUserNameOk

`func (o *GetSyslogEndpointResponse) GetWapiUserNameOk() (*string, bool)`

GetWapiUserNameOk returns a tuple with the WapiUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserName

`func (o *GetSyslogEndpointResponse) SetWapiUserName(v string)`

SetWapiUserName sets WapiUserName field to given value.

### HasWapiUserName

`func (o *GetSyslogEndpointResponse) HasWapiUserName() bool`

HasWapiUserName returns a boolean if a field has been set.

### GetWapiUserPassword

`func (o *GetSyslogEndpointResponse) GetWapiUserPassword() string`

GetWapiUserPassword returns the WapiUserPassword field if non-nil, zero value otherwise.

### GetWapiUserPasswordOk

`func (o *GetSyslogEndpointResponse) GetWapiUserPasswordOk() (*string, bool)`

GetWapiUserPasswordOk returns a tuple with the WapiUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserPassword

`func (o *GetSyslogEndpointResponse) SetWapiUserPassword(v string)`

SetWapiUserPassword sets WapiUserPassword field to given value.

### HasWapiUserPassword

`func (o *GetSyslogEndpointResponse) HasWapiUserPassword() bool`

HasWapiUserPassword returns a boolean if a field has been set.

### GetResult

`func (o *GetSyslogEndpointResponse) GetResult() SyslogEndpoint`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSyslogEndpointResponse) GetResultOk() (*SyslogEndpoint, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSyslogEndpointResponse) SetResult(v SyslogEndpoint)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSyslogEndpointResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


