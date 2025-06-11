# GetMultiregionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CloudPlatform** | Pointer to **string** | Type of cloud_platform to get the supported regions. | [optional] 
**GovcloudRegions** | Pointer to **string** | Comma separated sting containing only GovCloud supported regions. | [optional] 
**Regions** | Pointer to **string** | Comma separated string which contains all supported regions. | [optional] 
**Result** | Pointer to [**Multiregions**](Multiregions.md) |  | [optional] 

## Methods

### NewGetMultiregionsResponse

`func NewGetMultiregionsResponse() *GetMultiregionsResponse`

NewGetMultiregionsResponse instantiates a new GetMultiregionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMultiregionsResponseWithDefaults

`func NewGetMultiregionsResponseWithDefaults() *GetMultiregionsResponse`

NewGetMultiregionsResponseWithDefaults instantiates a new GetMultiregionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMultiregionsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMultiregionsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMultiregionsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMultiregionsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudPlatform

`func (o *GetMultiregionsResponse) GetCloudPlatform() string`

GetCloudPlatform returns the CloudPlatform field if non-nil, zero value otherwise.

### GetCloudPlatformOk

`func (o *GetMultiregionsResponse) GetCloudPlatformOk() (*string, bool)`

GetCloudPlatformOk returns a tuple with the CloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudPlatform

`func (o *GetMultiregionsResponse) SetCloudPlatform(v string)`

SetCloudPlatform sets CloudPlatform field to given value.

### HasCloudPlatform

`func (o *GetMultiregionsResponse) HasCloudPlatform() bool`

HasCloudPlatform returns a boolean if a field has been set.

### GetGovcloudRegions

`func (o *GetMultiregionsResponse) GetGovcloudRegions() string`

GetGovcloudRegions returns the GovcloudRegions field if non-nil, zero value otherwise.

### GetGovcloudRegionsOk

`func (o *GetMultiregionsResponse) GetGovcloudRegionsOk() (*string, bool)`

GetGovcloudRegionsOk returns a tuple with the GovcloudRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovcloudRegions

`func (o *GetMultiregionsResponse) SetGovcloudRegions(v string)`

SetGovcloudRegions sets GovcloudRegions field to given value.

### HasGovcloudRegions

`func (o *GetMultiregionsResponse) HasGovcloudRegions() bool`

HasGovcloudRegions returns a boolean if a field has been set.

### GetRegions

`func (o *GetMultiregionsResponse) GetRegions() string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GetMultiregionsResponse) GetRegionsOk() (*string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GetMultiregionsResponse) SetRegions(v string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *GetMultiregionsResponse) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetResult

`func (o *GetMultiregionsResponse) GetResult() Multiregions`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMultiregionsResponse) GetResultOk() (*Multiregions, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMultiregionsResponse) SetResult(v Multiregions)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMultiregionsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


