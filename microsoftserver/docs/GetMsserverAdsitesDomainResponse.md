# GetMsserverAdsitesDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**EaDefinition** | Pointer to **string** | The name of the Extensible Attribute Definition object that is linked to the Active Directory Sites Domain. | [optional] [readonly] 
**MsSyncMasterName** | Pointer to **string** | The IP address or FQDN of the managing master for the MS server, if applicable. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Active Directory Domain properties object. | [optional] [readonly] 
**Netbios** | Pointer to **string** | The NetBIOS name of the Active Directory Domain properties object. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which the Active Directory Domain resides. | [optional] [readonly] 
**ReadOnly** | Pointer to **bool** | Determines whether the Active Directory Domain properties object is a read-only object. | [optional] [readonly] 
**Result** | Pointer to [**MsserverAdsitesDomain**](MsserverAdsitesDomain.md) |  | [optional] 

## Methods

### NewGetMsserverAdsitesDomainResponse

`func NewGetMsserverAdsitesDomainResponse() *GetMsserverAdsitesDomainResponse`

NewGetMsserverAdsitesDomainResponse instantiates a new GetMsserverAdsitesDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMsserverAdsitesDomainResponseWithDefaults

`func NewGetMsserverAdsitesDomainResponseWithDefaults() *GetMsserverAdsitesDomainResponse`

NewGetMsserverAdsitesDomainResponseWithDefaults instantiates a new GetMsserverAdsitesDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMsserverAdsitesDomainResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMsserverAdsitesDomainResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMsserverAdsitesDomainResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMsserverAdsitesDomainResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetEaDefinition

`func (o *GetMsserverAdsitesDomainResponse) GetEaDefinition() string`

GetEaDefinition returns the EaDefinition field if non-nil, zero value otherwise.

### GetEaDefinitionOk

`func (o *GetMsserverAdsitesDomainResponse) GetEaDefinitionOk() (*string, bool)`

GetEaDefinitionOk returns a tuple with the EaDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaDefinition

`func (o *GetMsserverAdsitesDomainResponse) SetEaDefinition(v string)`

SetEaDefinition sets EaDefinition field to given value.

### HasEaDefinition

`func (o *GetMsserverAdsitesDomainResponse) HasEaDefinition() bool`

HasEaDefinition returns a boolean if a field has been set.

### GetMsSyncMasterName

`func (o *GetMsserverAdsitesDomainResponse) GetMsSyncMasterName() string`

GetMsSyncMasterName returns the MsSyncMasterName field if non-nil, zero value otherwise.

### GetMsSyncMasterNameOk

`func (o *GetMsserverAdsitesDomainResponse) GetMsSyncMasterNameOk() (*string, bool)`

GetMsSyncMasterNameOk returns a tuple with the MsSyncMasterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSyncMasterName

`func (o *GetMsserverAdsitesDomainResponse) SetMsSyncMasterName(v string)`

SetMsSyncMasterName sets MsSyncMasterName field to given value.

### HasMsSyncMasterName

`func (o *GetMsserverAdsitesDomainResponse) HasMsSyncMasterName() bool`

HasMsSyncMasterName returns a boolean if a field has been set.

### GetName

`func (o *GetMsserverAdsitesDomainResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMsserverAdsitesDomainResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMsserverAdsitesDomainResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetMsserverAdsitesDomainResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetbios

`func (o *GetMsserverAdsitesDomainResponse) GetNetbios() string`

GetNetbios returns the Netbios field if non-nil, zero value otherwise.

### GetNetbiosOk

`func (o *GetMsserverAdsitesDomainResponse) GetNetbiosOk() (*string, bool)`

GetNetbiosOk returns a tuple with the Netbios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbios

`func (o *GetMsserverAdsitesDomainResponse) SetNetbios(v string)`

SetNetbios sets Netbios field to given value.

### HasNetbios

`func (o *GetMsserverAdsitesDomainResponse) HasNetbios() bool`

HasNetbios returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetMsserverAdsitesDomainResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetMsserverAdsitesDomainResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetMsserverAdsitesDomainResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetMsserverAdsitesDomainResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetReadOnly

`func (o *GetMsserverAdsitesDomainResponse) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *GetMsserverAdsitesDomainResponse) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *GetMsserverAdsitesDomainResponse) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *GetMsserverAdsitesDomainResponse) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetResult

`func (o *GetMsserverAdsitesDomainResponse) GetResult() MsserverAdsitesDomain`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMsserverAdsitesDomainResponse) GetResultOk() (*MsserverAdsitesDomain, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMsserverAdsitesDomainResponse) SetResult(v MsserverAdsitesDomain)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMsserverAdsitesDomainResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


