# MsserverAdsitesDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**EaDefinition** | Pointer to **string** | The name of the Extensible Attribute Definition object that is linked to the Active Directory Sites Domain. | [optional] [readonly] 
**MsSyncMasterName** | Pointer to **string** | The IP address or FQDN of the managing master for the MS server, if applicable. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Active Directory Domain properties object. | [optional] [readonly] 
**Netbios** | Pointer to **string** | The NetBIOS name of the Active Directory Domain properties object. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which the Active Directory Domain resides. | [optional] [readonly] 
**ReadOnly** | Pointer to **bool** | Determines whether the Active Directory Domain properties object is a read-only object. | [optional] [readonly] 

## Methods

### NewMsserverAdsitesDomain

`func NewMsserverAdsitesDomain() *MsserverAdsitesDomain`

NewMsserverAdsitesDomain instantiates a new MsserverAdsitesDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsserverAdsitesDomainWithDefaults

`func NewMsserverAdsitesDomainWithDefaults() *MsserverAdsitesDomain`

NewMsserverAdsitesDomainWithDefaults instantiates a new MsserverAdsitesDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MsserverAdsitesDomain) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MsserverAdsitesDomain) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MsserverAdsitesDomain) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MsserverAdsitesDomain) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetEaDefinition

`func (o *MsserverAdsitesDomain) GetEaDefinition() string`

GetEaDefinition returns the EaDefinition field if non-nil, zero value otherwise.

### GetEaDefinitionOk

`func (o *MsserverAdsitesDomain) GetEaDefinitionOk() (*string, bool)`

GetEaDefinitionOk returns a tuple with the EaDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaDefinition

`func (o *MsserverAdsitesDomain) SetEaDefinition(v string)`

SetEaDefinition sets EaDefinition field to given value.

### HasEaDefinition

`func (o *MsserverAdsitesDomain) HasEaDefinition() bool`

HasEaDefinition returns a boolean if a field has been set.

### GetMsSyncMasterName

`func (o *MsserverAdsitesDomain) GetMsSyncMasterName() string`

GetMsSyncMasterName returns the MsSyncMasterName field if non-nil, zero value otherwise.

### GetMsSyncMasterNameOk

`func (o *MsserverAdsitesDomain) GetMsSyncMasterNameOk() (*string, bool)`

GetMsSyncMasterNameOk returns a tuple with the MsSyncMasterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSyncMasterName

`func (o *MsserverAdsitesDomain) SetMsSyncMasterName(v string)`

SetMsSyncMasterName sets MsSyncMasterName field to given value.

### HasMsSyncMasterName

`func (o *MsserverAdsitesDomain) HasMsSyncMasterName() bool`

HasMsSyncMasterName returns a boolean if a field has been set.

### GetName

`func (o *MsserverAdsitesDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MsserverAdsitesDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MsserverAdsitesDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MsserverAdsitesDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetbios

`func (o *MsserverAdsitesDomain) GetNetbios() string`

GetNetbios returns the Netbios field if non-nil, zero value otherwise.

### GetNetbiosOk

`func (o *MsserverAdsitesDomain) GetNetbiosOk() (*string, bool)`

GetNetbiosOk returns a tuple with the Netbios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbios

`func (o *MsserverAdsitesDomain) SetNetbios(v string)`

SetNetbios sets Netbios field to given value.

### HasNetbios

`func (o *MsserverAdsitesDomain) HasNetbios() bool`

HasNetbios returns a boolean if a field has been set.

### GetNetworkView

`func (o *MsserverAdsitesDomain) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *MsserverAdsitesDomain) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *MsserverAdsitesDomain) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *MsserverAdsitesDomain) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetReadOnly

`func (o *MsserverAdsitesDomain) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MsserverAdsitesDomain) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *MsserverAdsitesDomain) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *MsserverAdsitesDomain) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


