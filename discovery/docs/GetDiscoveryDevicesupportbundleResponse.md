# GetDiscoveryDevicesupportbundleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Author** | Pointer to **string** | The developer of the device support bundle. | [optional] [readonly] 
**IntegratedInd** | Pointer to **bool** | Determines whether the device support bundle is integrated or imported. Note that integrated support bundles cannot be removed. | [optional] [readonly] 
**Name** | Pointer to **string** | The descriptive device name for the device support bundle. | [optional] [readonly] 
**Version** | Pointer to **string** | The version of the currently active device support bundle. | [optional] [readonly] 
**Result** | Pointer to [**DiscoveryDevicesupportbundle**](DiscoveryDevicesupportbundle.md) |  | [optional] 

## Methods

### NewGetDiscoveryDevicesupportbundleResponse

`func NewGetDiscoveryDevicesupportbundleResponse() *GetDiscoveryDevicesupportbundleResponse`

NewGetDiscoveryDevicesupportbundleResponse instantiates a new GetDiscoveryDevicesupportbundleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveryDevicesupportbundleResponseWithDefaults

`func NewGetDiscoveryDevicesupportbundleResponseWithDefaults() *GetDiscoveryDevicesupportbundleResponse`

NewGetDiscoveryDevicesupportbundleResponseWithDefaults instantiates a new GetDiscoveryDevicesupportbundleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoveryDevicesupportbundleResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoveryDevicesupportbundleResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoveryDevicesupportbundleResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoveryDevicesupportbundleResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthor

`func (o *GetDiscoveryDevicesupportbundleResponse) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GetDiscoveryDevicesupportbundleResponse) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GetDiscoveryDevicesupportbundleResponse) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GetDiscoveryDevicesupportbundleResponse) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetIntegratedInd

`func (o *GetDiscoveryDevicesupportbundleResponse) GetIntegratedInd() bool`

GetIntegratedInd returns the IntegratedInd field if non-nil, zero value otherwise.

### GetIntegratedIndOk

`func (o *GetDiscoveryDevicesupportbundleResponse) GetIntegratedIndOk() (*bool, bool)`

GetIntegratedIndOk returns a tuple with the IntegratedInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratedInd

`func (o *GetDiscoveryDevicesupportbundleResponse) SetIntegratedInd(v bool)`

SetIntegratedInd sets IntegratedInd field to given value.

### HasIntegratedInd

`func (o *GetDiscoveryDevicesupportbundleResponse) HasIntegratedInd() bool`

HasIntegratedInd returns a boolean if a field has been set.

### GetName

`func (o *GetDiscoveryDevicesupportbundleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDiscoveryDevicesupportbundleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDiscoveryDevicesupportbundleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDiscoveryDevicesupportbundleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *GetDiscoveryDevicesupportbundleResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetDiscoveryDevicesupportbundleResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetDiscoveryDevicesupportbundleResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetDiscoveryDevicesupportbundleResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoveryDevicesupportbundleResponse) GetResult() DiscoveryDevicesupportbundle`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoveryDevicesupportbundleResponse) GetResultOk() (*DiscoveryDevicesupportbundle, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoveryDevicesupportbundleResponse) SetResult(v DiscoveryDevicesupportbundle)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoveryDevicesupportbundleResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


