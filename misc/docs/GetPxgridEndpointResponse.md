# GetPxgridEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**VendorIdentifier** | Pointer to **string** | The vendor identifier. | [optional] 
**WapiUserName** | Pointer to **string** | The user name for WAPI integration. | [optional] 
**WapiUserPassword** | Pointer to **string** | The user password for WAPI integration. | [optional] 
**Result** | Pointer to [**PxgridEndpoint**](PxgridEndpoint.md) |  | [optional] 

## Methods

### NewGetPxgridEndpointResponse

`func NewGetPxgridEndpointResponse() *GetPxgridEndpointResponse`

NewGetPxgridEndpointResponse instantiates a new GetPxgridEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPxgridEndpointResponseWithDefaults

`func NewGetPxgridEndpointResponseWithDefaults() *GetPxgridEndpointResponse`

NewGetPxgridEndpointResponseWithDefaults instantiates a new GetPxgridEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetPxgridEndpointResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetPxgridEndpointResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetPxgridEndpointResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetPxgridEndpointResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetPxgridEndpointResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetPxgridEndpointResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetPxgridEndpointResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetPxgridEndpointResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetClientCertificateSubject

`func (o *GetPxgridEndpointResponse) GetClientCertificateSubject() string`

GetClientCertificateSubject returns the ClientCertificateSubject field if non-nil, zero value otherwise.

### GetClientCertificateSubjectOk

`func (o *GetPxgridEndpointResponse) GetClientCertificateSubjectOk() (*string, bool)`

GetClientCertificateSubjectOk returns a tuple with the ClientCertificateSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateSubject

`func (o *GetPxgridEndpointResponse) SetClientCertificateSubject(v string)`

SetClientCertificateSubject sets ClientCertificateSubject field to given value.

### HasClientCertificateSubject

`func (o *GetPxgridEndpointResponse) HasClientCertificateSubject() bool`

HasClientCertificateSubject returns a boolean if a field has been set.

### GetClientCertificateToken

`func (o *GetPxgridEndpointResponse) GetClientCertificateToken() string`

GetClientCertificateToken returns the ClientCertificateToken field if non-nil, zero value otherwise.

### GetClientCertificateTokenOk

`func (o *GetPxgridEndpointResponse) GetClientCertificateTokenOk() (*string, bool)`

GetClientCertificateTokenOk returns a tuple with the ClientCertificateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateToken

`func (o *GetPxgridEndpointResponse) SetClientCertificateToken(v string)`

SetClientCertificateToken sets ClientCertificateToken field to given value.

### HasClientCertificateToken

`func (o *GetPxgridEndpointResponse) HasClientCertificateToken() bool`

HasClientCertificateToken returns a boolean if a field has been set.

### GetClientCertificateValidFrom

`func (o *GetPxgridEndpointResponse) GetClientCertificateValidFrom() int64`

GetClientCertificateValidFrom returns the ClientCertificateValidFrom field if non-nil, zero value otherwise.

### GetClientCertificateValidFromOk

`func (o *GetPxgridEndpointResponse) GetClientCertificateValidFromOk() (*int64, bool)`

GetClientCertificateValidFromOk returns a tuple with the ClientCertificateValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidFrom

`func (o *GetPxgridEndpointResponse) SetClientCertificateValidFrom(v int64)`

SetClientCertificateValidFrom sets ClientCertificateValidFrom field to given value.

### HasClientCertificateValidFrom

`func (o *GetPxgridEndpointResponse) HasClientCertificateValidFrom() bool`

HasClientCertificateValidFrom returns a boolean if a field has been set.

### GetClientCertificateValidTo

`func (o *GetPxgridEndpointResponse) GetClientCertificateValidTo() int64`

GetClientCertificateValidTo returns the ClientCertificateValidTo field if non-nil, zero value otherwise.

### GetClientCertificateValidToOk

`func (o *GetPxgridEndpointResponse) GetClientCertificateValidToOk() (*int64, bool)`

GetClientCertificateValidToOk returns a tuple with the ClientCertificateValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateValidTo

`func (o *GetPxgridEndpointResponse) SetClientCertificateValidTo(v int64)`

SetClientCertificateValidTo sets ClientCertificateValidTo field to given value.

### HasClientCertificateValidTo

`func (o *GetPxgridEndpointResponse) HasClientCertificateValidTo() bool`

HasClientCertificateValidTo returns a boolean if a field has been set.

### GetComment

`func (o *GetPxgridEndpointResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetPxgridEndpointResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetPxgridEndpointResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetPxgridEndpointResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetPxgridEndpointResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetPxgridEndpointResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetPxgridEndpointResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetPxgridEndpointResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetPxgridEndpointResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetPxgridEndpointResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetPxgridEndpointResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetPxgridEndpointResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetPxgridEndpointResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetPxgridEndpointResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetPxgridEndpointResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetPxgridEndpointResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetPxgridEndpointResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetPxgridEndpointResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetPxgridEndpointResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetPxgridEndpointResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLogLevel

`func (o *GetPxgridEndpointResponse) GetLogLevel() string`

GetLogLevel returns the LogLevel field if non-nil, zero value otherwise.

### GetLogLevelOk

`func (o *GetPxgridEndpointResponse) GetLogLevelOk() (*string, bool)`

GetLogLevelOk returns a tuple with the LogLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLevel

`func (o *GetPxgridEndpointResponse) SetLogLevel(v string)`

SetLogLevel sets LogLevel field to given value.

### HasLogLevel

`func (o *GetPxgridEndpointResponse) HasLogLevel() bool`

HasLogLevel returns a boolean if a field has been set.

### GetName

`func (o *GetPxgridEndpointResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPxgridEndpointResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPxgridEndpointResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPxgridEndpointResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetPxgridEndpointResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetPxgridEndpointResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetPxgridEndpointResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetPxgridEndpointResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetOutboundMemberType

`func (o *GetPxgridEndpointResponse) GetOutboundMemberType() string`

GetOutboundMemberType returns the OutboundMemberType field if non-nil, zero value otherwise.

### GetOutboundMemberTypeOk

`func (o *GetPxgridEndpointResponse) GetOutboundMemberTypeOk() (*string, bool)`

GetOutboundMemberTypeOk returns a tuple with the OutboundMemberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMemberType

`func (o *GetPxgridEndpointResponse) SetOutboundMemberType(v string)`

SetOutboundMemberType sets OutboundMemberType field to given value.

### HasOutboundMemberType

`func (o *GetPxgridEndpointResponse) HasOutboundMemberType() bool`

HasOutboundMemberType returns a boolean if a field has been set.

### GetOutboundMembers

`func (o *GetPxgridEndpointResponse) GetOutboundMembers() []string`

GetOutboundMembers returns the OutboundMembers field if non-nil, zero value otherwise.

### GetOutboundMembersOk

`func (o *GetPxgridEndpointResponse) GetOutboundMembersOk() (*[]string, bool)`

GetOutboundMembersOk returns a tuple with the OutboundMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundMembers

`func (o *GetPxgridEndpointResponse) SetOutboundMembers(v []string)`

SetOutboundMembers sets OutboundMembers field to given value.

### HasOutboundMembers

`func (o *GetPxgridEndpointResponse) HasOutboundMembers() bool`

HasOutboundMembers returns a boolean if a field has been set.

### GetPublishSettings

`func (o *GetPxgridEndpointResponse) GetPublishSettings() PxgridEndpointPublishSettings`

GetPublishSettings returns the PublishSettings field if non-nil, zero value otherwise.

### GetPublishSettingsOk

`func (o *GetPxgridEndpointResponse) GetPublishSettingsOk() (*PxgridEndpointPublishSettings, bool)`

GetPublishSettingsOk returns a tuple with the PublishSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishSettings

`func (o *GetPxgridEndpointResponse) SetPublishSettings(v PxgridEndpointPublishSettings)`

SetPublishSettings sets PublishSettings field to given value.

### HasPublishSettings

`func (o *GetPxgridEndpointResponse) HasPublishSettings() bool`

HasPublishSettings returns a boolean if a field has been set.

### GetSubscribeSettings

`func (o *GetPxgridEndpointResponse) GetSubscribeSettings() PxgridEndpointSubscribeSettings`

GetSubscribeSettings returns the SubscribeSettings field if non-nil, zero value otherwise.

### GetSubscribeSettingsOk

`func (o *GetPxgridEndpointResponse) GetSubscribeSettingsOk() (*PxgridEndpointSubscribeSettings, bool)`

GetSubscribeSettingsOk returns a tuple with the SubscribeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeSettings

`func (o *GetPxgridEndpointResponse) SetSubscribeSettings(v PxgridEndpointSubscribeSettings)`

SetSubscribeSettings sets SubscribeSettings field to given value.

### HasSubscribeSettings

`func (o *GetPxgridEndpointResponse) HasSubscribeSettings() bool`

HasSubscribeSettings returns a boolean if a field has been set.

### GetTemplateInstance

`func (o *GetPxgridEndpointResponse) GetTemplateInstance() PxgridEndpointTemplateInstance`

GetTemplateInstance returns the TemplateInstance field if non-nil, zero value otherwise.

### GetTemplateInstanceOk

`func (o *GetPxgridEndpointResponse) GetTemplateInstanceOk() (*PxgridEndpointTemplateInstance, bool)`

GetTemplateInstanceOk returns a tuple with the TemplateInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInstance

`func (o *GetPxgridEndpointResponse) SetTemplateInstance(v PxgridEndpointTemplateInstance)`

SetTemplateInstance sets TemplateInstance field to given value.

### HasTemplateInstance

`func (o *GetPxgridEndpointResponse) HasTemplateInstance() bool`

HasTemplateInstance returns a boolean if a field has been set.

### GetTimeout

`func (o *GetPxgridEndpointResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetPxgridEndpointResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetPxgridEndpointResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetPxgridEndpointResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetVendorIdentifier

`func (o *GetPxgridEndpointResponse) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *GetPxgridEndpointResponse) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *GetPxgridEndpointResponse) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.

### HasVendorIdentifier

`func (o *GetPxgridEndpointResponse) HasVendorIdentifier() bool`

HasVendorIdentifier returns a boolean if a field has been set.

### GetWapiUserName

`func (o *GetPxgridEndpointResponse) GetWapiUserName() string`

GetWapiUserName returns the WapiUserName field if non-nil, zero value otherwise.

### GetWapiUserNameOk

`func (o *GetPxgridEndpointResponse) GetWapiUserNameOk() (*string, bool)`

GetWapiUserNameOk returns a tuple with the WapiUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserName

`func (o *GetPxgridEndpointResponse) SetWapiUserName(v string)`

SetWapiUserName sets WapiUserName field to given value.

### HasWapiUserName

`func (o *GetPxgridEndpointResponse) HasWapiUserName() bool`

HasWapiUserName returns a boolean if a field has been set.

### GetWapiUserPassword

`func (o *GetPxgridEndpointResponse) GetWapiUserPassword() string`

GetWapiUserPassword returns the WapiUserPassword field if non-nil, zero value otherwise.

### GetWapiUserPasswordOk

`func (o *GetPxgridEndpointResponse) GetWapiUserPasswordOk() (*string, bool)`

GetWapiUserPasswordOk returns a tuple with the WapiUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWapiUserPassword

`func (o *GetPxgridEndpointResponse) SetWapiUserPassword(v string)`

SetWapiUserPassword sets WapiUserPassword field to given value.

### HasWapiUserPassword

`func (o *GetPxgridEndpointResponse) HasWapiUserPassword() bool`

HasWapiUserPassword returns a boolean if a field has been set.

### GetResult

`func (o *GetPxgridEndpointResponse) GetResult() PxgridEndpoint`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetPxgridEndpointResponse) GetResultOk() (*PxgridEndpoint, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetPxgridEndpointResponse) SetResult(v PxgridEndpoint)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetPxgridEndpointResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


