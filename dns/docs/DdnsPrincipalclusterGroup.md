# DdnsPrincipalclusterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Clusters** | Pointer to **[]string** | The list of equivalent DDNS principal clusters. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the DDNS Principal Cluster Group. | [optional] 
**Name** | Pointer to **string** | The name of this DDNS Principal Cluster Group. | [optional] 

## Methods

### NewDdnsPrincipalclusterGroup

`func NewDdnsPrincipalclusterGroup() *DdnsPrincipalclusterGroup`

NewDdnsPrincipalclusterGroup instantiates a new DdnsPrincipalclusterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdnsPrincipalclusterGroupWithDefaults

`func NewDdnsPrincipalclusterGroupWithDefaults() *DdnsPrincipalclusterGroup`

NewDdnsPrincipalclusterGroupWithDefaults instantiates a new DdnsPrincipalclusterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DdnsPrincipalclusterGroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DdnsPrincipalclusterGroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DdnsPrincipalclusterGroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DdnsPrincipalclusterGroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetClusters

`func (o *DdnsPrincipalclusterGroup) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *DdnsPrincipalclusterGroup) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *DdnsPrincipalclusterGroup) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *DdnsPrincipalclusterGroup) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetComment

`func (o *DdnsPrincipalclusterGroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DdnsPrincipalclusterGroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DdnsPrincipalclusterGroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DdnsPrincipalclusterGroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *DdnsPrincipalclusterGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DdnsPrincipalclusterGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DdnsPrincipalclusterGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DdnsPrincipalclusterGroup) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


