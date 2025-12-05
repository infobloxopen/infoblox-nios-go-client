# PxgridEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The pxgrid endpoint IPv4 Address or IPv6 Address or Fully-Qualified Domain Name (FQDN) | [optional] 
**ClientCertificateSubject** | Pointer to **string** | The Cisco ISE client certificate subject. | [optional] [readonly] 
**ClientCertificateToken** | Pointer to **string** | The token returned by the uploadinit function call in object fileop for Cisco ISE client certificate. | [optional] 
**ClientCertificateValidFrom** | Pointer to **int64** | The pxgrid endpoint client certificate valid from. | [optional] [readonly] 
**ClientCertificateValidTo** | Pointer to **int64** | The pxgrid endpoint client certificate valid to. | [optional] [readonly] 
**Comment** | Pointer to **string** | The Cisco ISE endpoint descriptive comment. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a Cisco ISE endpoint is disabled or not. When this is set to False, the Cisco ISE endpoint is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LogLevel** | Pointer to **string** | The log level for a notification pxgrid endpoint. | [optional] 
**Name** | Pointer to **string** | The name of the pxgrid endpoint. | [optional] 
**NetworkView** | Pointer to **string** | The pxgrid network view name. | [optional] 
**OutboundMemberType** | Pointer to **string** | The outbound member that will generate events. | [optional] 
**OutboundMembers** | Pointer to **[]string** | The list of members for outbound events. | [optional] 
**PublishSettings** | Pointer to [**PxgridEndpointPublishSettings**](PxgridEndpointPublishSettings.md) |  | [optional] 
**SubscribeSettings** | Pointer to [**PxgridEndpointSubscribeSettings**](PxgridEndpointSubscribeSettings.md) |  | [optional] 
**TemplateInstance** | Pointer to [**PxgridEndpointTemplateInstance**](PxgridEndpointTemplateInstance.md) |  | [optional] 
**Timeout** | Pointer to **int64** | The timeout of session management (in seconds). | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**VendorIdentifier** | Pointer to **string** | The vendor identifier. | [optional] 
**WapiUserName** | Pointer to **string** | The user name for WAPI integration. | [optional] 
**WapiUserPassword** | Pointer to **string** | The user password for WAPI integration. | [optional] 

## Methods

### NewPxgridEndpoint

`func NewPxgridEndpoint() *PxgridEndpoint`

NewPxgridEndpoint instantiates a new PxgridEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPxgridEndpointWithDefaults

`func NewPxgridEndpointWithDefaults() *PxgridEndpoint`

NewPxgridEndpointWithDefaults instantiates a new PxgridEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *PxgridEndpoint) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PxgridEndpoint) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PxgridEndpoint) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PxgridEndpoint) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *PxgridEndpoint) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PxgridEndpoint) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PxgridEndpoint) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PxgridEndpoint) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetClientCertificateSubject

`func (o *PxgridEndpoint) GetClientCertificateSubject() string`

GetClientCertificateSubject returns the ClientCertificateSubject field if non-nil, zero value otherwise.

### GetClientCertificateSubjectOk

`func (o *PxgridEndpoint) GetClientCertificateSubjectOk() (*string, bool)`

GetClientCertificateSubjectOk returns a tuple with the ClientCertificateSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateSubject

`func (o *PxgridEndpoint) SetClientCertificateSubject(v string)`

SetClientCertificateSubject sets ClientCertificateSubject field to given value.

### HasClientCertificateSubject

`func (o *PxgridEndpoint) HasClientCertificateSubject() bool`

HasClientCertificateSubject returns a boolean if a field has been set.

### GetClientCertificateToken

`func (o *PxgridEndpoint) GetClientCertificateToken() string`

GetClientCertificateToken returns the ClientCertificateToken field if non-nil, zero value otherwise.

### GetClientCertificateTokenOk

`func (o *PxgridEndpoint) GetClientCertificateTokenOk() (*string, bool)`

GetClientCertificateTokenOk returns a tuple with the ClientCertificateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateToken

`func (o *PxgridEndpoint) SetClientCertificateToken(v string)`

SetClientCertificateToken sets ClientCertificateToken field to given value.

### HasClientCertificateToken

`func (o *PxgridEndpoint) HasClientCertificateToken() bool`

HasClientCertificateToken returns a boolean if a field has been set.

### GetClientCertificateValidFrom

`func (o *PxgridEndpoint) GetClientCertificateValidFrom() int64`

GetClientCertificateValidFrom returns the ClientCertificateValidFrom field if non-nil, zero value otherwise.

### GetClientCertificateValidFromOk

`func (o *PxgridEndpoint) GetClientCertificateValidFromOk() (*int64, bool)`

GetClientCertificateValidFromOk returns a tuple with the ClientCertificateValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidFrom

`func (o *PxgridEndpoint) SetClientCertificateValidFrom(v int64)`

SetClientCertificateValidFrom sets ClientCertificateValidFrom field to given value.

### HasClientCertificateValidFrom

`func (o *PxgridEndpoint) HasClientCertificateValidFrom() bool`

HasClientCertificateValidFrom returns a boolean if a field has been set.

### GetClientCertificateValidTo

`func (o *PxgridEndpoint) GetClientCertificateValidTo() int64`

GetClientCertificateValidTo returns the ClientCertificateValidTo field if non-nil, zero value otherwise.

### GetClientCertificateValidToOk

`func (o *PxgridEndpoint) GetClientCertificateValidToOk() (*int64, bool)`

GetClientCertificateValidToOk returns a tuple with the ClientCertificateValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidTo

`func (o *PxgridEndpoint) SetClientCertificateValidTo(v int64)`

SetClientCertificateValidTo sets ClientCertificateValidTo field to given value.

### HasClientCertificateValidTo

`func (o *PxgridEndpoint) HasClientCertificateValidTo() bool`

HasClientCertificateValidTo returns a boolean if a field has been set.

### GetComment

`func (o *PxgridEndpoint) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PxgridEndpoint) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PxgridEndpoint) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PxgridEndpoint) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *PxgridEndpoint) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *PxgridEndpoint) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *PxgridEndpoint) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *PxgridEndpoint) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *PxgridEndpoint) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *PxgridEndpoint) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *PxgridEndpoint) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *PxgridEndpoint) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *PxgridEndpoint) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *PxgridEndpoint) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *PxgridEndpoint) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *PxgridEndpoint) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *PxgridEndpoint) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *PxgridEndpoint) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *PxgridEndpoint) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *PxgridEndpoint) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLogLevel

`func (o *PxgridEndpoint) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *PxgridEndpoint) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *PxgridEndpoint) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *PxgridEndpoint) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetName

`func (o *PxgridEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PxgridEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PxgridEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PxgridEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *PxgridEndpoint) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *PxgridEndpoint) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *PxgridEndpoint) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *PxgridEndpoint) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetOutboundMemberType

`func (o *PxgridEndpoint) GetOutboundMemberType() string`

GetOutboundMemberType returns the OutboundMemberType field if non-nil, zero value otherwise.

### GetOutboundMemberTypeOk

`func (o *PxgridEndpoint) GetOutboundMemberTypeOk() (*string, bool)`

GetOutboundMemberTypeOk returns a tuple with the OutboundMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMemberType

`func (o *PxgridEndpoint) SetOutboundMemberType(v string)`

SetOutboundMemberType sets OutboundMemberType field to given value.

### HasOutboundMemberType

`func (o *PxgridEndpoint) HasOutboundMemberType() bool`

HasOutboundMemberType returns a boolean if a field has been set.

### GetOutboundMembers

`func (o *PxgridEndpoint) GetOutboundMembers() []string`

GetOutboundMembers returns the OutboundMembers field if non-nil, zero value otherwise.

### GetOutboundMembersOk

`func (o *PxgridEndpoint) GetOutboundMembersOk() (*[]string, bool)`

GetOutboundMembersOk returns a tuple with the OutboundMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMembers

`func (o *PxgridEndpoint) SetOutboundMembers(v []string)`

SetOutboundMembers sets OutboundMembers field to given value.

### HasOutboundMembers

`func (o *PxgridEndpoint) HasOutboundMembers() bool`

HasOutboundMembers returns a boolean if a field has been set.

### GetPublishSettings

`func (o *PxgridEndpoint) GetPublishSettings() PxgridEndpointPublishSettings`

GetPublishSettings returns the PublishSettings field if non-nil, zero value otherwise.

### GetPublishSettingsOk

`func (o *PxgridEndpoint) GetPublishSettingsOk() (*PxgridEndpointPublishSettings, bool)`

GetPublishSettingsOk returns a tuple with the PublishSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishSettings

`func (o *PxgridEndpoint) SetPublishSettings(v PxgridEndpointPublishSettings)`

SetPublishSettings sets PublishSettings field to given value.

### HasPublishSettings

`func (o *PxgridEndpoint) HasPublishSettings() bool`

HasPublishSettings returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *PxgridEndpoint) GetSubscribeSettings() PxgridEndpointSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *PxgridEndpoint) GetSubscribeSettingsOk() (*PxgridEndpointSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *PxgridEndpoint) SetSubscribeSettings(v PxgridEndpointSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *PxgridEndpoint) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *PxgridEndpoint) GetTemplateInstance() PxgridEndpointTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *PxgridEndpoint) GetTemplateInstanceOk() (*PxgridEndpointTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *PxgridEndpoint) SetTemplateInstance(v PxgridEndpointTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *PxgridEndpoint) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetTimeout

`func (o *PxgridEndpoint) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PxgridEndpoint) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PxgridEndpoint) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PxgridEndpoint) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUuid

`func (o *PxgridEndpoint) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PxgridEndpoint) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PxgridEndpoint) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PxgridEndpoint) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *PxgridEndpoint) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *PxgridEndpoint) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *PxgridEndpoint) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *PxgridEndpoint) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetWapiUserName

`func (o *PxgridEndpoint) GetWapiUserName() string`

GetWapiUserName returns the WapiUserName field if non-nil, zero value otherwise.

### GetWapiUserNameOk

`func (o *PxgridEndpoint) GetWapiUserNameOk() (*string, bool)`

GetWapiUserNameOk returns a tuple with the WapiUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserName

`func (o *PxgridEndpoint) SetWapiUserName(v string)`

SetWapiUserName sets WapiUserName field to given value.

### HasWapiUserName

`func (o *PxgridEndpoint) HasWapiUserName() bool`

HasWapiUserName returns a boolean if a field has been set.

### GetWapiUserPassword

`func (o *PxgridEndpoint) GetWapiUserPassword() string`

GetWapiUserPassword returns the WapiUserPassword field if non-nil, zero value otherwise.

### GetWapiUserPasswordOk

`func (o *PxgridEndpoint) GetWapiUserPasswordOk() (*string, bool)`

GetWapiUserPasswordOk returns a tuple with the WapiUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserPassword

`func (o *PxgridEndpoint) SetWapiUserPassword(v string)`

SetWapiUserPassword sets WapiUserPassword field to given value.

### HasWapiUserPassword

`func (o *PxgridEndpoint) HasWapiUserPassword() bool`

HasWapiUserPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


