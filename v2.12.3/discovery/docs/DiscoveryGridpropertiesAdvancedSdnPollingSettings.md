# DiscoveryGridpropertiesAdvancedSdnPollingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworksMappingPolicy** | Pointer to **string** | Policy of mapping SDN networks to network views. To obtain mapping table use the object discovery:sdnnetwork | [optional] 
**DisableSdnDiscoveryOutsideIpam** | Pointer to **bool** | Disable discovery of SDN subnets that are not in IPAM. | [optional] 

## Methods

### NewDiscoveryGridpropertiesAdvancedSdnPollingSettings

`func NewDiscoveryGridpropertiesAdvancedSdnPollingSettings() *DiscoveryGridpropertiesAdvancedSdnPollingSettings`

NewDiscoveryGridpropertiesAdvancedSdnPollingSettings instantiates a new DiscoveryGridpropertiesAdvancedSdnPollingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryGridpropertiesAdvancedSdnPollingSettingsWithDefaults

`func NewDiscoveryGridpropertiesAdvancedSdnPollingSettingsWithDefaults() *DiscoveryGridpropertiesAdvancedSdnPollingSettings`

NewDiscoveryGridpropertiesAdvancedSdnPollingSettingsWithDefaults instantiates a new DiscoveryGridpropertiesAdvancedSdnPollingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworksMappingPolicy

`func (o *DiscoveryGridpropertiesAdvancedSdnPollingSettings) GetNetworksMappingPolicy() string`

GetNetworksMappingPolicy returns the NetworksMappingPolicy field if non-nil, zero value otherwise.

### GetNetworksMappingPolicyOk

`func (o *DiscoveryGridpropertiesAdvancedSdnPollingSettings) GetNetworksMappingPolicyOk() (*string, bool)`

GetNetworksMappingPolicyOk returns a tuple with the NetworksMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworksMappingPolicy

`func (o *DiscoveryGridpropertiesAdvancedSdnPollingSettings) SetNetworksMappingPolicy(v string)`

SetNetworksMappingPolicy sets NetworksMappingPolicy field to given value.

### HasNetworksMappingPolicy

`func (o *DiscoveryGridpropertiesAdvancedSdnPollingSettings) HasNetworksMappingPolicy() bool`

HasNetworksMappingPolicy returns a boolean if a field has been set.

### GetDisableSdnDiscoveryOutsideIpam

`func (o *DiscoveryGridpropertiesAdvancedSdnPollingSettings) GetDisableSdnDiscoveryOutsideIpam() bool`

GetDisableSdnDiscoveryOutsideIpam returns the DisableSdnDiscoveryOutsideIpam field if non-nil, zero value otherwise.

### GetDisableSdnDiscoveryOutsideIpamOk

`func (o *DiscoveryGridpropertiesAdvancedSdnPollingSettings) GetDisableSdnDiscoveryOutsideIpamOk() (*bool, bool)`

GetDisableSdnDiscoveryOutsideIpamOk returns a tuple with the DisableSdnDiscoveryOutsideIpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSdnDiscoveryOutsideIpam

`func (o *DiscoveryGridpropertiesAdvancedSdnPollingSettings) SetDisableSdnDiscoveryOutsideIpam(v bool)`

SetDisableSdnDiscoveryOutsideIpam sets DisableSdnDiscoveryOutsideIpam field to given value.

### HasDisableSdnDiscoveryOutsideIpam

`func (o *DiscoveryGridpropertiesAdvancedSdnPollingSettings) HasDisableSdnDiscoveryOutsideIpam() bool`

HasDisableSdnDiscoveryOutsideIpam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


