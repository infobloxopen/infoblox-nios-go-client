# Multiregions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CloudPlatform** | Pointer to **string** | Type of cloud_platform to get the supported regions. | [optional] 
**GovcloudRegions** | Pointer to **string** | Comma separated sting containing only GovCloud supported regions. | [optional] 
**Regions** | Pointer to **string** | Comma separated string which contains all supported regions. | [optional] 

## Methods

### NewMultiregions

`func NewMultiregions() *Multiregions`

NewMultiregions instantiates a new Multiregions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiregionsWithDefaults

`func NewMultiregionsWithDefaults() *Multiregions`

NewMultiregionsWithDefaults instantiates a new Multiregions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Multiregions) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Multiregions) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Multiregions) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Multiregions) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudPlatform

`func (o *Multiregions) GetCloudPlatform() string`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *Multiregions) GetCloudPlatformOk() (*string, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *Multiregions) SetCloudPlatform(v string)`

SetCloudPlatform sets CloudPlatform field to given value.

### HasCloudPlatform

`func (o *Multiregions) HasCloudPlatform() bool`

HasCloudPlatform returns a boolean if a field has been set.

### GetGovcloudRegions

`func (o *Multiregions) GetGovcloudRegions() string`

GetGovcloudRegions returns the GovcloudRegions field if non-nil, zero value otherwise.

### GetGovcloudRegionsOk

`func (o *Multiregions) GetGovcloudRegionsOk() (*string, bool)`

GetGovcloudRegionsOk returns a tuple with the GovcloudRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovcloudRegions

`func (o *Multiregions) SetGovcloudRegions(v string)`

SetGovcloudRegions sets GovcloudRegions field to given value.

### HasGovcloudRegions

`func (o *Multiregions) HasGovcloudRegions() bool`

HasGovcloudRegions returns a boolean if a field has been set.

### GetRegions

`func (o *Multiregions) GetRegions() string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Multiregions) GetRegionsOk() (*string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Multiregions) SetRegions(v string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Multiregions) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


