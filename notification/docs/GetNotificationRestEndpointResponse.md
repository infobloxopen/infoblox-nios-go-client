# GetNotificationRestEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**ClientCertificateSubject** | Pointer to **string** | The client certificate subject of a notification REST endpoint. | [optional] [readonly] 
**ClientCertificateToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop for a notification REST endpoit client certificate. | [optional] 
**ClientCertificateValidFrom** | Pointer to **int64** | The timestamp when client certificate for a notification REST endpoint was created. | [optional] [readonly] 
**ClientCertificateValidTo** | Pointer to **int64** | The timestamp when client certificate for a notification REST endpoint expires. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment of a notification REST endpoint. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LogLevel** | Pointer to **string** | The log level for a notification REST endpoint. | [optional] 
**Name** | Pointer to **string** | The name of a notification REST endpoint. | [optional] 
**OutboundMemberType** | Pointer to **string** | The outbound member which will generate an event. | [optional] 
**OutboundMembers** | Pointer to **[]string** | The list of members for outbound events. | [optional] 
**Password** | Pointer to **string** | The password of the user that can log into a notification REST endpoint. | [optional] 
**ServerCertValidation** | Pointer to **string** | The server certificate validation type. | [optional] 
**SyncDisabled** | Pointer to **bool** | Determines if the sync process is disabled for a notification REST endpoint. | [optional] 
**TemplateInstance** | Pointer to [**NotificationRestEndpointTemplateInstance**](NotificationRestEndpointTemplateInstance.md) |  | [optional] 
**Timeout** | Pointer to **int64** | The timeout of session management (in seconds). | [optional] 
**Uri** | Pointer to **string** | The URI of a notification REST endpoint. | [optional] 
**Username** | Pointer to **string** | The username of the user that can log into a notification REST endpoint. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**VendorIdentifier** | Pointer to **string** | The vendor identifier. | [optional] 
**WapiUserName** | Pointer to **string** | The user name for WAPI integration. | [optional] 
**WapiUserPassword** | Pointer to **string** | The user password for WAPI integration. | [optional] 
**Result** | Pointer to [**NotificationRestEndpoint**](NotificationRestEndpoint.md) |  | [optional] 

## Methods

### NewGetNotificationRestEndpointResponse

`func NewGetNotificationRestEndpointResponse() *GetNotificationRestEndpointResponse`

NewGetNotificationRestEndpointResponse instantiates a new GetNotificationRestEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationRestEndpointResponseWithDefaults

`func NewGetNotificationRestEndpointResponseWithDefaults() *GetNotificationRestEndpointResponse`

NewGetNotificationRestEndpointResponseWithDefaults instantiates a new GetNotificationRestEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNotificationRestEndpointResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNotificationRestEndpointResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNotificationRestEndpointResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNotificationRestEndpointResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetClientCertificateSubject

`func (o *GetNotificationRestEndpointResponse) GetClientCertificateSubject() string`

GetClientCertificateSubject returns the ClientCertificateSubject field if non-nil, zero value otherwise.

### GetClientCertificateSubjectOk

`func (o *GetNotificationRestEndpointResponse) GetClientCertificateSubjectOk() (*string, bool)`

GetClientCertificateSubjectOk returns a tuple with the ClientCertificateSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateSubject

`func (o *GetNotificationRestEndpointResponse) SetClientCertificateSubject(v string)`

SetClientCertificateSubject sets ClientCertificateSubject field to given value.

### HasClientCertificateSubject

`func (o *GetNotificationRestEndpointResponse) HasClientCertificateSubject() bool`

HasClientCertificateSubject returns a boolean if a field has been set.

### GetClientCertificateToken

`func (o *GetNotificationRestEndpointResponse) GetClientCertificateToken() string`

GetClientCertificateToken returns the ClientCertificateToken field if non-nil, zero value otherwise.

### GetClientCertificateTokenOk

`func (o *GetNotificationRestEndpointResponse) GetClientCertificateTokenOk() (*string, bool)`

GetClientCertificateTokenOk returns a tuple with the ClientCertificateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateToken

`func (o *GetNotificationRestEndpointResponse) SetClientCertificateToken(v string)`

SetClientCertificateToken sets ClientCertificateToken field to given value.

### HasClientCertificateToken

`func (o *GetNotificationRestEndpointResponse) HasClientCertificateToken() bool`

HasClientCertificateToken returns a boolean if a field has been set.

### GetClientCertificateValidFrom

`func (o *GetNotificationRestEndpointResponse) GetClientCertificateValidFrom() int64`

GetClientCertificateValidFrom returns the ClientCertificateValidFrom field if non-nil, zero value otherwise.

### GetClientCertificateValidFromOk

`func (o *GetNotificationRestEndpointResponse) GetClientCertificateValidFromOk() (*int64, bool)`

GetClientCertificateValidFromOk returns a tuple with the ClientCertificateValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidFrom

`func (o *GetNotificationRestEndpointResponse) SetClientCertificateValidFrom(v int64)`

SetClientCertificateValidFrom sets ClientCertificateValidFrom field to given value.

### HasClientCertificateValidFrom

`func (o *GetNotificationRestEndpointResponse) HasClientCertificateValidFrom() bool`

HasClientCertificateValidFrom returns a boolean if a field has been set.

### GetClientCertificateValidTo

`func (o *GetNotificationRestEndpointResponse) GetClientCertificateValidTo() int64`

GetClientCertificateValidTo returns the ClientCertificateValidTo field if non-nil, zero value otherwise.

### GetClientCertificateValidToOk

`func (o *GetNotificationRestEndpointResponse) GetClientCertificateValidToOk() (*int64, bool)`

GetClientCertificateValidToOk returns a tuple with the ClientCertificateValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidTo

`func (o *GetNotificationRestEndpointResponse) SetClientCertificateValidTo(v int64)`

SetClientCertificateValidTo sets ClientCertificateValidTo field to given value.

### HasClientCertificateValidTo

`func (o *GetNotificationRestEndpointResponse) HasClientCertificateValidTo() bool`

HasClientCertificateValidTo returns a boolean if a field has been set.

### GetComment

`func (o *GetNotificationRestEndpointResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNotificationRestEndpointResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNotificationRestEndpointResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNotificationRestEndpointResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetNotificationRestEndpointResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetNotificationRestEndpointResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetNotificationRestEndpointResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetNotificationRestEndpointResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetNotificationRestEndpointResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetNotificationRestEndpointResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetNotificationRestEndpointResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetNotificationRestEndpointResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetNotificationRestEndpointResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetNotificationRestEndpointResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetNotificationRestEndpointResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetNotificationRestEndpointResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLogLevel

`func (o *GetNotificationRestEndpointResponse) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *GetNotificationRestEndpointResponse) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *GetNotificationRestEndpointResponse) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *GetNotificationRestEndpointResponse) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetName

`func (o *GetNotificationRestEndpointResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNotificationRestEndpointResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNotificationRestEndpointResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNotificationRestEndpointResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundMemberType

`func (o *GetNotificationRestEndpointResponse) GetOutboundMemberType() string`

GetOutboundMemberType returns the OutboundMemberType field if non-nil, zero value otherwise.

### GetOutboundMemberTypeOk

`func (o *GetNotificationRestEndpointResponse) GetOutboundMemberTypeOk() (*string, bool)`

GetOutboundMemberTypeOk returns a tuple with the OutboundMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMemberType

`func (o *GetNotificationRestEndpointResponse) SetOutboundMemberType(v string)`

SetOutboundMemberType sets OutboundMemberType field to given value.

### HasOutboundMemberType

`func (o *GetNotificationRestEndpointResponse) HasOutboundMemberType() bool`

HasOutboundMemberType returns a boolean if a field has been set.

### GetOutboundMembers

`func (o *GetNotificationRestEndpointResponse) GetOutboundMembers() []string`

GetOutboundMembers returns the OutboundMembers field if non-nil, zero value otherwise.

### GetOutboundMembersOk

`func (o *GetNotificationRestEndpointResponse) GetOutboundMembersOk() (*[]string, bool)`

GetOutboundMembersOk returns a tuple with the OutboundMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMembers

`func (o *GetNotificationRestEndpointResponse) SetOutboundMembers(v []string)`

SetOutboundMembers sets OutboundMembers field to given value.

### HasOutboundMembers

`func (o *GetNotificationRestEndpointResponse) HasOutboundMembers() bool`

HasOutboundMembers returns a boolean if a field has been set.

### GetPassword

`func (o *GetNotificationRestEndpointResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetNotificationRestEndpointResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetNotificationRestEndpointResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetNotificationRestEndpointResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServerCertValidation

`func (o *GetNotificationRestEndpointResponse) GetServerCertValidation() string`

GetServerCertValidation returns the ServerCertValidation field if non-nil, zero value otherwise.

### GetServerCertValidationOk

`func (o *GetNotificationRestEndpointResponse) GetServerCertValidationOk() (*string, bool)`

GetServerCertValidationOk returns a tuple with the ServerCertValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertValidation

`func (o *GetNotificationRestEndpointResponse) SetServerCertValidation(v string)`

SetServerCertValidation sets ServerCertValidation field to given value.

### HasServerCertValidation

`func (o *GetNotificationRestEndpointResponse) HasServerCertValidation() bool`

HasServerCertValidation returns a boolean if a field has been set.

### GetSyncDisabled

`func (o *GetNotificationRestEndpointResponse) GetSyncDisabled() bool`

GetSyncDisabled returns the SyncDisabled field if non-nil, zero value otherwise.

### GetSyncDisabledOk

`func (o *GetNotificationRestEndpointResponse) GetSyncDisabledOk() (*bool, bool)`

GetSyncDisabledOk returns a tuple with the SyncDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDisabled

`func (o *GetNotificationRestEndpointResponse) SetSyncDisabled(v bool)`

SetSyncDisabled sets SyncDisabled field to given value.

### HasSyncDisabled

`func (o *GetNotificationRestEndpointResponse) HasSyncDisabled() bool`

HasSyncDisabled returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *GetNotificationRestEndpointResponse) GetTemplateInstance() NotificationRestEndpointTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *GetNotificationRestEndpointResponse) GetTemplateInstanceOk() (*NotificationRestEndpointTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *GetNotificationRestEndpointResponse) SetTemplateInstance(v NotificationRestEndpointTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *GetNotificationRestEndpointResponse) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetTimeout

`func (o *GetNotificationRestEndpointResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetNotificationRestEndpointResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetNotificationRestEndpointResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetNotificationRestEndpointResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUri

`func (o *GetNotificationRestEndpointResponse) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *GetNotificationRestEndpointResponse) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *GetNotificationRestEndpointResponse) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *GetNotificationRestEndpointResponse) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUsername

`func (o *GetNotificationRestEndpointResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetNotificationRestEndpointResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetNotificationRestEndpointResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetNotificationRestEndpointResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *GetNotificationRestEndpointResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetNotificationRestEndpointResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetNotificationRestEndpointResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetNotificationRestEndpointResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *GetNotificationRestEndpointResponse) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *GetNotificationRestEndpointResponse) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *GetNotificationRestEndpointResponse) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *GetNotificationRestEndpointResponse) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetWapiUserName

`func (o *GetNotificationRestEndpointResponse) GetWapiUserName() string`

GetWapiUserName returns the WapiUserName field if non-nil, zero value otherwise.

### GetWapiUserNameOk

`func (o *GetNotificationRestEndpointResponse) GetWapiUserNameOk() (*string, bool)`

GetWapiUserNameOk returns a tuple with the WapiUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserName

`func (o *GetNotificationRestEndpointResponse) SetWapiUserName(v string)`

SetWapiUserName sets WapiUserName field to given value.

### HasWapiUserName

`func (o *GetNotificationRestEndpointResponse) HasWapiUserName() bool`

HasWapiUserName returns a boolean if a field has been set.

### GetWapiUserPassword

`func (o *GetNotificationRestEndpointResponse) GetWapiUserPassword() string`

GetWapiUserPassword returns the WapiUserPassword field if non-nil, zero value otherwise.

### GetWapiUserPasswordOk

`func (o *GetNotificationRestEndpointResponse) GetWapiUserPasswordOk() (*string, bool)`

GetWapiUserPasswordOk returns a tuple with the WapiUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserPassword

`func (o *GetNotificationRestEndpointResponse) SetWapiUserPassword(v string)`

SetWapiUserPassword sets WapiUserPassword field to given value.

### HasWapiUserPassword

`func (o *GetNotificationRestEndpointResponse) HasWapiUserPassword() bool`

HasWapiUserPassword returns a boolean if a field has been set.

### GetResult

`func (o *GetNotificationRestEndpointResponse) GetResult() NotificationRestEndpoint`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNotificationRestEndpointResponse) GetResultOk() (*NotificationRestEndpoint, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNotificationRestEndpointResponse) SetResult(v NotificationRestEndpoint)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNotificationRestEndpointResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


