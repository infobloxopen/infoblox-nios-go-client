# DxlEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Brokers** | Pointer to [**[]DxlEndpointBrokers**](DxlEndpointBrokers.md) | The list of DXL endpoint brokers. Note that you cannot specify brokers and brokers_import_token at the same time. | [optional] 
**BrokersImportToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop for a DXL broker configuration file. Note that you cannot specify brokers and brokers_import_token at the same time. | [optional] 
**ClientCertificateSubject** | Pointer to **string** | The client certificate subject of a DXL endpoint. | [optional] [readonly] 
**ClientCertificateToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop for a DXL endpoint client certificate. | [optional] 
**ClientCertificateValidFrom** | Pointer to **int64** | The timestamp when client certificate for a DXL endpoint was created. | [optional] [readonly] 
**ClientCertificateValidTo** | Pointer to **int64** | The timestamp when the client certificate for a DXL endpoint expires. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment of a DXL endpoint. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a DXL endpoint is disabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LogLevel** | Pointer to **string** | The log level for a DXL endpoint. | [optional] 
**Name** | Pointer to **string** | The name of a DXL endpoint. | [optional] 
**OutboundMemberType** | Pointer to **string** | The outbound member that will generate events. | [optional] 
**OutboundMembers** | Pointer to **[]string** | The list of members for outbound events. | [optional] 
**TemplateInstance** | Pointer to [**DxlEndpointTemplateInstance**](DxlEndpointTemplateInstance.md) |  | [optional] 
**Timeout** | Pointer to **int64** | The timeout of session management (in seconds). | [optional] 
**Topics** | Pointer to **[]string** | DXL topics | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**VendorIdentifier** | Pointer to **string** | The vendor identifier. | [optional] 
**WapiUserName** | Pointer to **string** | The user name for WAPI integration. | [optional] 
**WapiUserPassword** | Pointer to **string** | The user password for WAPI integration. | [optional] 

## Methods

### NewDxlEndpoint

`func NewDxlEndpoint() *DxlEndpoint`

NewDxlEndpoint instantiates a new DxlEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDxlEndpointWithDefaults

`func NewDxlEndpointWithDefaults() *DxlEndpoint`

NewDxlEndpointWithDefaults instantiates a new DxlEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DxlEndpoint) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DxlEndpoint) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DxlEndpoint) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DxlEndpoint) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBrokers

`func (o *DxlEndpoint) GetBrokers() []DxlEndpointBrokers`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *DxlEndpoint) GetBrokersOk() (*[]DxlEndpointBrokers, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *DxlEndpoint) SetBrokers(v []DxlEndpointBrokers)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *DxlEndpoint) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.

### GetBrokersImportToken

`func (o *DxlEndpoint) GetBrokersImportToken() string`

GetBrokersImportToken returns the BrokersImportToken field if non-nil, zero value otherwise.

### GetBrokersImportTokenOk

`func (o *DxlEndpoint) GetBrokersImportTokenOk() (*string, bool)`

GetBrokersImportTokenOk returns a tuple with the BrokersImportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokersImportToken

`func (o *DxlEndpoint) SetBrokersImportToken(v string)`

SetBrokersImportToken sets BrokersImportToken field to given value.

### HasBrokersImportToken

`func (o *DxlEndpoint) HasBrokersImportToken() bool`

HasBrokersImportToken returns a boolean if a field has been set.

### GetClientCertificateSubject

`func (o *DxlEndpoint) GetClientCertificateSubject() string`

GetClientCertificateSubject returns the ClientCertificateSubject field if non-nil, zero value otherwise.

### GetClientCertificateSubjectOk

`func (o *DxlEndpoint) GetClientCertificateSubjectOk() (*string, bool)`

GetClientCertificateSubjectOk returns a tuple with the ClientCertificateSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateSubject

`func (o *DxlEndpoint) SetClientCertificateSubject(v string)`

SetClientCertificateSubject sets ClientCertificateSubject field to given value.

### HasClientCertificateSubject

`func (o *DxlEndpoint) HasClientCertificateSubject() bool`

HasClientCertificateSubject returns a boolean if a field has been set.

### GetClientCertificateToken

`func (o *DxlEndpoint) GetClientCertificateToken() string`

GetClientCertificateToken returns the ClientCertificateToken field if non-nil, zero value otherwise.

### GetClientCertificateTokenOk

`func (o *DxlEndpoint) GetClientCertificateTokenOk() (*string, bool)`

GetClientCertificateTokenOk returns a tuple with the ClientCertificateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateToken

`func (o *DxlEndpoint) SetClientCertificateToken(v string)`

SetClientCertificateToken sets ClientCertificateToken field to given value.

### HasClientCertificateToken

`func (o *DxlEndpoint) HasClientCertificateToken() bool`

HasClientCertificateToken returns a boolean if a field has been set.

### GetClientCertificateValidFrom

`func (o *DxlEndpoint) GetClientCertificateValidFrom() int64`

GetClientCertificateValidFrom returns the ClientCertificateValidFrom field if non-nil, zero value otherwise.

### GetClientCertificateValidFromOk

`func (o *DxlEndpoint) GetClientCertificateValidFromOk() (*int64, bool)`

GetClientCertificateValidFromOk returns a tuple with the ClientCertificateValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidFrom

`func (o *DxlEndpoint) SetClientCertificateValidFrom(v int64)`

SetClientCertificateValidFrom sets ClientCertificateValidFrom field to given value.

### HasClientCertificateValidFrom

`func (o *DxlEndpoint) HasClientCertificateValidFrom() bool`

HasClientCertificateValidFrom returns a boolean if a field has been set.

### GetClientCertificateValidTo

`func (o *DxlEndpoint) GetClientCertificateValidTo() int64`

GetClientCertificateValidTo returns the ClientCertificateValidTo field if non-nil, zero value otherwise.

### GetClientCertificateValidToOk

`func (o *DxlEndpoint) GetClientCertificateValidToOk() (*int64, bool)`

GetClientCertificateValidToOk returns a tuple with the ClientCertificateValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidTo

`func (o *DxlEndpoint) SetClientCertificateValidTo(v int64)`

SetClientCertificateValidTo sets ClientCertificateValidTo field to given value.

### HasClientCertificateValidTo

`func (o *DxlEndpoint) HasClientCertificateValidTo() bool`

HasClientCertificateValidTo returns a boolean if a field has been set.

### GetComment

`func (o *DxlEndpoint) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DxlEndpoint) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DxlEndpoint) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DxlEndpoint) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *DxlEndpoint) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DxlEndpoint) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DxlEndpoint) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DxlEndpoint) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *DxlEndpoint) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *DxlEndpoint) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *DxlEndpoint) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *DxlEndpoint) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *DxlEndpoint) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *DxlEndpoint) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *DxlEndpoint) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *DxlEndpoint) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *DxlEndpoint) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *DxlEndpoint) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *DxlEndpoint) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *DxlEndpoint) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLogLevel

`func (o *DxlEndpoint) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *DxlEndpoint) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *DxlEndpoint) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *DxlEndpoint) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetName

`func (o *DxlEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DxlEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DxlEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DxlEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundMemberType

`func (o *DxlEndpoint) GetOutboundMemberType() string`

GetOutboundMemberType returns the OutboundMemberType field if non-nil, zero value otherwise.

### GetOutboundMemberTypeOk

`func (o *DxlEndpoint) GetOutboundMemberTypeOk() (*string, bool)`

GetOutboundMemberTypeOk returns a tuple with the OutboundMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMemberType

`func (o *DxlEndpoint) SetOutboundMemberType(v string)`

SetOutboundMemberType sets OutboundMemberType field to given value.

### HasOutboundMemberType

`func (o *DxlEndpoint) HasOutboundMemberType() bool`

HasOutboundMemberType returns a boolean if a field has been set.

### GetOutboundMembers

`func (o *DxlEndpoint) GetOutboundMembers() []string`

GetOutboundMembers returns the OutboundMembers field if non-nil, zero value otherwise.

### GetOutboundMembersOk

`func (o *DxlEndpoint) GetOutboundMembersOk() (*[]string, bool)`

GetOutboundMembersOk returns a tuple with the OutboundMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMembers

`func (o *DxlEndpoint) SetOutboundMembers(v []string)`

SetOutboundMembers sets OutboundMembers field to given value.

### HasOutboundMembers

`func (o *DxlEndpoint) HasOutboundMembers() bool`

HasOutboundMembers returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *DxlEndpoint) GetTemplateInstance() DxlEndpointTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *DxlEndpoint) GetTemplateInstanceOk() (*DxlEndpointTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *DxlEndpoint) SetTemplateInstance(v DxlEndpointTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *DxlEndpoint) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetTimeout

`func (o *DxlEndpoint) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DxlEndpoint) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DxlEndpoint) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DxlEndpoint) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTopics

`func (o *DxlEndpoint) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *DxlEndpoint) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *DxlEndpoint) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *DxlEndpoint) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetUuid

`func (o *DxlEndpoint) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DxlEndpoint) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DxlEndpoint) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DxlEndpoint) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *DxlEndpoint) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *DxlEndpoint) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *DxlEndpoint) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *DxlEndpoint) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetWapiUserName

`func (o *DxlEndpoint) GetWapiUserName() string`

GetWapiUserName returns the WapiUserName field if non-nil, zero value otherwise.

### GetWapiUserNameOk

`func (o *DxlEndpoint) GetWapiUserNameOk() (*string, bool)`

GetWapiUserNameOk returns a tuple with the WapiUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserName

`func (o *DxlEndpoint) SetWapiUserName(v string)`

SetWapiUserName sets WapiUserName field to given value.

### HasWapiUserName

`func (o *DxlEndpoint) HasWapiUserName() bool`

HasWapiUserName returns a boolean if a field has been set.

### GetWapiUserPassword

`func (o *DxlEndpoint) GetWapiUserPassword() string`

GetWapiUserPassword returns the WapiUserPassword field if non-nil, zero value otherwise.

### GetWapiUserPasswordOk

`func (o *DxlEndpoint) GetWapiUserPasswordOk() (*string, bool)`

GetWapiUserPasswordOk returns a tuple with the WapiUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserPassword

`func (o *DxlEndpoint) SetWapiUserPassword(v string)`

SetWapiUserPassword sets WapiUserPassword field to given value.

### HasWapiUserPassword

`func (o *DxlEndpoint) HasWapiUserPassword() bool`

HasWapiUserPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


