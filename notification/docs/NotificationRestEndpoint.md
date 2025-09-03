# NotificationRestEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**VendorIdentifier** | Pointer to **string** | The vendor identifier. | [optional] 
**WapiUserName** | Pointer to **string** | The user name for WAPI integration. | [optional] 
**WapiUserPassword** | Pointer to **string** | The user password for WAPI integration. | [optional] 

## Methods

### NewNotificationRestEndpoint

`func NewNotificationRestEndpoint() *NotificationRestEndpoint`

NewNotificationRestEndpoint instantiates a new NotificationRestEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRestEndpointWithDefaults

`func NewNotificationRestEndpointWithDefaults() *NotificationRestEndpoint`

NewNotificationRestEndpointWithDefaults instantiates a new NotificationRestEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *NotificationRestEndpoint) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *NotificationRestEndpoint) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *NotificationRestEndpoint) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *NotificationRestEndpoint) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetClientCertificateSubject

`func (o *NotificationRestEndpoint) GetClientCertificateSubject() string`

GetClientCertificateSubject returns the ClientCertificateSubject field if non-nil, zero value otherwise.

### GetClientCertificateSubjectOk

`func (o *NotificationRestEndpoint) GetClientCertificateSubjectOk() (*string, bool)`

GetClientCertificateSubjectOk returns a tuple with the ClientCertificateSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateSubject

`func (o *NotificationRestEndpoint) SetClientCertificateSubject(v string)`

SetClientCertificateSubject sets ClientCertificateSubject field to given value.

### HasClientCertificateSubject

`func (o *NotificationRestEndpoint) HasClientCertificateSubject() bool`

HasClientCertificateSubject returns a boolean if a field has been set.

### GetClientCertificateToken

`func (o *NotificationRestEndpoint) GetClientCertificateToken() string`

GetClientCertificateToken returns the ClientCertificateToken field if non-nil, zero value otherwise.

### GetClientCertificateTokenOk

`func (o *NotificationRestEndpoint) GetClientCertificateTokenOk() (*string, bool)`

GetClientCertificateTokenOk returns a tuple with the ClientCertificateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateToken

`func (o *NotificationRestEndpoint) SetClientCertificateToken(v string)`

SetClientCertificateToken sets ClientCertificateToken field to given value.

### HasClientCertificateToken

`func (o *NotificationRestEndpoint) HasClientCertificateToken() bool`

HasClientCertificateToken returns a boolean if a field has been set.

### GetClientCertificateValidFrom

`func (o *NotificationRestEndpoint) GetClientCertificateValidFrom() int64`

GetClientCertificateValidFrom returns the ClientCertificateValidFrom field if non-nil, zero value otherwise.

### GetClientCertificateValidFromOk

`func (o *NotificationRestEndpoint) GetClientCertificateValidFromOk() (*int64, bool)`

GetClientCertificateValidFromOk returns a tuple with the ClientCertificateValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidFrom

`func (o *NotificationRestEndpoint) SetClientCertificateValidFrom(v int64)`

SetClientCertificateValidFrom sets ClientCertificateValidFrom field to given value.

### HasClientCertificateValidFrom

`func (o *NotificationRestEndpoint) HasClientCertificateValidFrom() bool`

HasClientCertificateValidFrom returns a boolean if a field has been set.

### GetClientCertificateValidTo

`func (o *NotificationRestEndpoint) GetClientCertificateValidTo() int64`

GetClientCertificateValidTo returns the ClientCertificateValidTo field if non-nil, zero value otherwise.

### GetClientCertificateValidToOk

`func (o *NotificationRestEndpoint) GetClientCertificateValidToOk() (*int64, bool)`

GetClientCertificateValidToOk returns a tuple with the ClientCertificateValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidTo

`func (o *NotificationRestEndpoint) SetClientCertificateValidTo(v int64)`

SetClientCertificateValidTo sets ClientCertificateValidTo field to given value.

### HasClientCertificateValidTo

`func (o *NotificationRestEndpoint) HasClientCertificateValidTo() bool`

HasClientCertificateValidTo returns a boolean if a field has been set.

### GetComment

`func (o *NotificationRestEndpoint) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NotificationRestEndpoint) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NotificationRestEndpoint) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NotificationRestEndpoint) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *NotificationRestEndpoint) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *NotificationRestEndpoint) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *NotificationRestEndpoint) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *NotificationRestEndpoint) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *NotificationRestEndpoint) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *NotificationRestEndpoint) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *NotificationRestEndpoint) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *NotificationRestEndpoint) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *NotificationRestEndpoint) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *NotificationRestEndpoint) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *NotificationRestEndpoint) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *NotificationRestEndpoint) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLogLevel

`func (o *NotificationRestEndpoint) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *NotificationRestEndpoint) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *NotificationRestEndpoint) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *NotificationRestEndpoint) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetName

`func (o *NotificationRestEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRestEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRestEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationRestEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundMemberType

`func (o *NotificationRestEndpoint) GetOutboundMemberType() string`

GetOutboundMemberType returns the OutboundMemberType field if non-nil, zero value otherwise.

### GetOutboundMemberTypeOk

`func (o *NotificationRestEndpoint) GetOutboundMemberTypeOk() (*string, bool)`

GetOutboundMemberTypeOk returns a tuple with the OutboundMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMemberType

`func (o *NotificationRestEndpoint) SetOutboundMemberType(v string)`

SetOutboundMemberType sets OutboundMemberType field to given value.

### HasOutboundMemberType

`func (o *NotificationRestEndpoint) HasOutboundMemberType() bool`

HasOutboundMemberType returns a boolean if a field has been set.

### GetOutboundMembers

`func (o *NotificationRestEndpoint) GetOutboundMembers() []string`

GetOutboundMembers returns the OutboundMembers field if non-nil, zero value otherwise.

### GetOutboundMembersOk

`func (o *NotificationRestEndpoint) GetOutboundMembersOk() (*[]string, bool)`

GetOutboundMembersOk returns a tuple with the OutboundMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMembers

`func (o *NotificationRestEndpoint) SetOutboundMembers(v []string)`

SetOutboundMembers sets OutboundMembers field to given value.

### HasOutboundMembers

`func (o *NotificationRestEndpoint) HasOutboundMembers() bool`

HasOutboundMembers returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationRestEndpoint) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationRestEndpoint) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationRestEndpoint) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationRestEndpoint) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetServerCertValidation

`func (o *NotificationRestEndpoint) GetServerCertValidation() string`

GetServerCertValidation returns the ServerCertValidation field if non-nil, zero value otherwise.

### GetServerCertValidationOk

`func (o *NotificationRestEndpoint) GetServerCertValidationOk() (*string, bool)`

GetServerCertValidationOk returns a tuple with the ServerCertValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCertValidation

`func (o *NotificationRestEndpoint) SetServerCertValidation(v string)`

SetServerCertValidation sets ServerCertValidation field to given value.

### HasServerCertValidation

`func (o *NotificationRestEndpoint) HasServerCertValidation() bool`

HasServerCertValidation returns a boolean if a field has been set.

### GetSyncDisabled

`func (o *NotificationRestEndpoint) GetSyncDisabled() bool`

GetSyncDisabled returns the SyncDisabled field if non-nil, zero value otherwise.

### GetSyncDisabledOk

`func (o *NotificationRestEndpoint) GetSyncDisabledOk() (*bool, bool)`

GetSyncDisabledOk returns a tuple with the SyncDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncDisabled

`func (o *NotificationRestEndpoint) SetSyncDisabled(v bool)`

SetSyncDisabled sets SyncDisabled field to given value.

### HasSyncDisabled

`func (o *NotificationRestEndpoint) HasSyncDisabled() bool`

HasSyncDisabled returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *NotificationRestEndpoint) GetTemplateInstance() NotificationRestEndpointTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *NotificationRestEndpoint) GetTemplateInstanceOk() (*NotificationRestEndpointTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *NotificationRestEndpoint) SetTemplateInstance(v NotificationRestEndpointTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *NotificationRestEndpoint) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetTimeout

`func (o *NotificationRestEndpoint) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *NotificationRestEndpoint) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *NotificationRestEndpoint) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *NotificationRestEndpoint) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUri

`func (o *NotificationRestEndpoint) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *NotificationRestEndpoint) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *NotificationRestEndpoint) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *NotificationRestEndpoint) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationRestEndpoint) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationRestEndpoint) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationRestEndpoint) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationRestEndpoint) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *NotificationRestEndpoint) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *NotificationRestEndpoint) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *NotificationRestEndpoint) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *NotificationRestEndpoint) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetWapiUserName

`func (o *NotificationRestEndpoint) GetWapiUserName() string`

GetWapiUserName returns the WapiUserName field if non-nil, zero value otherwise.

### GetWapiUserNameOk

`func (o *NotificationRestEndpoint) GetWapiUserNameOk() (*string, bool)`

GetWapiUserNameOk returns a tuple with the WapiUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserName

`func (o *NotificationRestEndpoint) SetWapiUserName(v string)`

SetWapiUserName sets WapiUserName field to given value.

### HasWapiUserName

`func (o *NotificationRestEndpoint) HasWapiUserName() bool`

HasWapiUserName returns a boolean if a field has been set.

### GetWapiUserPassword

`func (o *NotificationRestEndpoint) GetWapiUserPassword() string`

GetWapiUserPassword returns the WapiUserPassword field if non-nil, zero value otherwise.

### GetWapiUserPasswordOk

`func (o *NotificationRestEndpoint) GetWapiUserPasswordOk() (*string, bool)`

GetWapiUserPasswordOk returns a tuple with the WapiUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserPassword

`func (o *NotificationRestEndpoint) SetWapiUserPassword(v string)`

SetWapiUserPassword sets WapiUserPassword field to given value.

### HasWapiUserPassword

`func (o *NotificationRestEndpoint) HasWapiUserPassword() bool`

HasWapiUserPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


