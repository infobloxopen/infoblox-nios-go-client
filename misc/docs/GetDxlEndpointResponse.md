# GetDxlEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Brokers** | Pointer to [**[]DxlEndpointBrokers**](DxlEndpointBrokers.md) | The list of DXL endpoint brokers. Note that you cannot specify brokers and brokers_import_token at the same time. | [optional] 
**BrokersImportToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop for a DXL broker configuration file. Note that you cannot specify brokers and brokers_import_token at the same time. | [optional] 
**ClientCertificateSubject** | Pointer to **string** | The client certificate subject of a DXL endpoint. | [optional] [readonly] 
**ClientCertificateToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop for a DXL endpoint client certificate. | [optional] 
**ClientCertificateValidFrom** | Pointer to **int64** | The timestamp when client certificate for a DXL endpoint was created. | [optional] [readonly] 
**ClientCertificateValidTo** | Pointer to **int64** | The timestamp when the client certificate for a DXL endpoint expires. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment of a DXL endpoint. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a DXL endpoint is disabled. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LogLevel** | Pointer to **string** | The log level for a DXL endpoint. | [optional] 
**Name** | Pointer to **string** | The name of a DXL endpoint. | [optional] 
**OutboundMemberType** | Pointer to **string** | The outbound member that will generate events. | [optional] 
**OutboundMembers** | Pointer to **[]string** | The list of members for outbound events. | [optional] 
**TemplateInstance** | Pointer to [**DxlEndpointTemplateInstance**](DxlEndpointTemplateInstance.md) |  | [optional] 
**Timeout** | Pointer to **int64** | The timeout of session management (in seconds). | [optional] 
**Topics** | Pointer to **[]string** | DXL topics | [optional] 
**VendorIdentifier** | Pointer to **string** | The vendor identifier. | [optional] 
**WapiUserName** | Pointer to **string** | The user name for WAPI integration. | [optional] 
**WapiUserPassword** | Pointer to **string** | The user password for WAPI integration. | [optional] 
**Result** | Pointer to [**DxlEndpoint**](DxlEndpoint.md) |  | [optional] 

## Methods

### NewGetDxlEndpointResponse

`func NewGetDxlEndpointResponse() *GetDxlEndpointResponse`

NewGetDxlEndpointResponse instantiates a new GetDxlEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDxlEndpointResponseWithDefaults

`func NewGetDxlEndpointResponseWithDefaults() *GetDxlEndpointResponse`

NewGetDxlEndpointResponseWithDefaults instantiates a new GetDxlEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDxlEndpointResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDxlEndpointResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDxlEndpointResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDxlEndpointResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBrokers

`func (o *GetDxlEndpointResponse) GetBrokers() []DxlEndpointBrokers`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *GetDxlEndpointResponse) GetBrokersOk() (*[]DxlEndpointBrokers, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *GetDxlEndpointResponse) SetBrokers(v []DxlEndpointBrokers)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *GetDxlEndpointResponse) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.

### GetBrokersImportToken

`func (o *GetDxlEndpointResponse) GetBrokersImportToken() string`

GetBrokersImportToken returns the BrokersImportToken field if non-nil, zero value otherwise.

### GetBrokersImportTokenOk

`func (o *GetDxlEndpointResponse) GetBrokersImportTokenOk() (*string, bool)`

GetBrokersImportTokenOk returns a tuple with the BrokersImportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokersImportToken

`func (o *GetDxlEndpointResponse) SetBrokersImportToken(v string)`

SetBrokersImportToken sets BrokersImportToken field to given value.

### HasBrokersImportToken

`func (o *GetDxlEndpointResponse) HasBrokersImportToken() bool`

HasBrokersImportToken returns a boolean if a field has been set.

### GetClientCertificateSubject

`func (o *GetDxlEndpointResponse) GetClientCertificateSubject() string`

GetClientCertificateSubject returns the ClientCertificateSubject field if non-nil, zero value otherwise.

### GetClientCertificateSubjectOk

`func (o *GetDxlEndpointResponse) GetClientCertificateSubjectOk() (*string, bool)`

GetClientCertificateSubjectOk returns a tuple with the ClientCertificateSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateSubject

`func (o *GetDxlEndpointResponse) SetClientCertificateSubject(v string)`

SetClientCertificateSubject sets ClientCertificateSubject field to given value.

### HasClientCertificateSubject

`func (o *GetDxlEndpointResponse) HasClientCertificateSubject() bool`

HasClientCertificateSubject returns a boolean if a field has been set.

### GetClientCertificateToken

`func (o *GetDxlEndpointResponse) GetClientCertificateToken() string`

GetClientCertificateToken returns the ClientCertificateToken field if non-nil, zero value otherwise.

### GetClientCertificateTokenOk

`func (o *GetDxlEndpointResponse) GetClientCertificateTokenOk() (*string, bool)`

GetClientCertificateTokenOk returns a tuple with the ClientCertificateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateToken

`func (o *GetDxlEndpointResponse) SetClientCertificateToken(v string)`

SetClientCertificateToken sets ClientCertificateToken field to given value.

### HasClientCertificateToken

`func (o *GetDxlEndpointResponse) HasClientCertificateToken() bool`

HasClientCertificateToken returns a boolean if a field has been set.

### GetClientCertificateValidFrom

`func (o *GetDxlEndpointResponse) GetClientCertificateValidFrom() int64`

GetClientCertificateValidFrom returns the ClientCertificateValidFrom field if non-nil, zero value otherwise.

### GetClientCertificateValidFromOk

`func (o *GetDxlEndpointResponse) GetClientCertificateValidFromOk() (*int64, bool)`

GetClientCertificateValidFromOk returns a tuple with the ClientCertificateValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidFrom

`func (o *GetDxlEndpointResponse) SetClientCertificateValidFrom(v int64)`

SetClientCertificateValidFrom sets ClientCertificateValidFrom field to given value.

### HasClientCertificateValidFrom

`func (o *GetDxlEndpointResponse) HasClientCertificateValidFrom() bool`

HasClientCertificateValidFrom returns a boolean if a field has been set.

### GetClientCertificateValidTo

`func (o *GetDxlEndpointResponse) GetClientCertificateValidTo() int64`

GetClientCertificateValidTo returns the ClientCertificateValidTo field if non-nil, zero value otherwise.

### GetClientCertificateValidToOk

`func (o *GetDxlEndpointResponse) GetClientCertificateValidToOk() (*int64, bool)`

GetClientCertificateValidToOk returns a tuple with the ClientCertificateValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidTo

`func (o *GetDxlEndpointResponse) SetClientCertificateValidTo(v int64)`

SetClientCertificateValidTo sets ClientCertificateValidTo field to given value.

### HasClientCertificateValidTo

`func (o *GetDxlEndpointResponse) HasClientCertificateValidTo() bool`

HasClientCertificateValidTo returns a boolean if a field has been set.

### GetComment

`func (o *GetDxlEndpointResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDxlEndpointResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDxlEndpointResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDxlEndpointResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetDxlEndpointResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetDxlEndpointResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetDxlEndpointResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetDxlEndpointResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetDxlEndpointResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetDxlEndpointResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetDxlEndpointResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetDxlEndpointResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetLogLevel

`func (o *GetDxlEndpointResponse) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *GetDxlEndpointResponse) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *GetDxlEndpointResponse) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *GetDxlEndpointResponse) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetName

`func (o *GetDxlEndpointResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDxlEndpointResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDxlEndpointResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDxlEndpointResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundMemberType

`func (o *GetDxlEndpointResponse) GetOutboundMemberType() string`

GetOutboundMemberType returns the OutboundMemberType field if non-nil, zero value otherwise.

### GetOutboundMemberTypeOk

`func (o *GetDxlEndpointResponse) GetOutboundMemberTypeOk() (*string, bool)`

GetOutboundMemberTypeOk returns a tuple with the OutboundMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMemberType

`func (o *GetDxlEndpointResponse) SetOutboundMemberType(v string)`

SetOutboundMemberType sets OutboundMemberType field to given value.

### HasOutboundMemberType

`func (o *GetDxlEndpointResponse) HasOutboundMemberType() bool`

HasOutboundMemberType returns a boolean if a field has been set.

### GetOutboundMembers

`func (o *GetDxlEndpointResponse) GetOutboundMembers() []string`

GetOutboundMembers returns the OutboundMembers field if non-nil, zero value otherwise.

### GetOutboundMembersOk

`func (o *GetDxlEndpointResponse) GetOutboundMembersOk() (*[]string, bool)`

GetOutboundMembersOk returns a tuple with the OutboundMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMembers

`func (o *GetDxlEndpointResponse) SetOutboundMembers(v []string)`

SetOutboundMembers sets OutboundMembers field to given value.

### HasOutboundMembers

`func (o *GetDxlEndpointResponse) HasOutboundMembers() bool`

HasOutboundMembers returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *GetDxlEndpointResponse) GetTemplateInstance() DxlEndpointTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *GetDxlEndpointResponse) GetTemplateInstanceOk() (*DxlEndpointTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *GetDxlEndpointResponse) SetTemplateInstance(v DxlEndpointTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *GetDxlEndpointResponse) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetTimeout

`func (o *GetDxlEndpointResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetDxlEndpointResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetDxlEndpointResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetDxlEndpointResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTopics

`func (o *GetDxlEndpointResponse) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *GetDxlEndpointResponse) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *GetDxlEndpointResponse) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *GetDxlEndpointResponse) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *GetDxlEndpointResponse) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *GetDxlEndpointResponse) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *GetDxlEndpointResponse) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *GetDxlEndpointResponse) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetWapiUserName

`func (o *GetDxlEndpointResponse) GetWapiUserName() string`

GetWapiUserName returns the WapiUserName field if non-nil, zero value otherwise.

### GetWapiUserNameOk

`func (o *GetDxlEndpointResponse) GetWapiUserNameOk() (*string, bool)`

GetWapiUserNameOk returns a tuple with the WapiUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserName

`func (o *GetDxlEndpointResponse) SetWapiUserName(v string)`

SetWapiUserName sets WapiUserName field to given value.

### HasWapiUserName

`func (o *GetDxlEndpointResponse) HasWapiUserName() bool`

HasWapiUserName returns a boolean if a field has been set.

### GetWapiUserPassword

`func (o *GetDxlEndpointResponse) GetWapiUserPassword() string`

GetWapiUserPassword returns the WapiUserPassword field if non-nil, zero value otherwise.

### GetWapiUserPasswordOk

`func (o *GetDxlEndpointResponse) GetWapiUserPasswordOk() (*string, bool)`

GetWapiUserPasswordOk returns a tuple with the WapiUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserPassword

`func (o *GetDxlEndpointResponse) SetWapiUserPassword(v string)`

SetWapiUserPassword sets WapiUserPassword field to given value.

### HasWapiUserPassword

`func (o *GetDxlEndpointResponse) HasWapiUserPassword() bool`

HasWapiUserPassword returns a boolean if a field has been set.

### GetResult

`func (o *GetDxlEndpointResponse) GetResult() DxlEndpoint`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDxlEndpointResponse) GetResultOk() (*DxlEndpoint, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDxlEndpointResponse) SetResult(v DxlEndpoint)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDxlEndpointResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


