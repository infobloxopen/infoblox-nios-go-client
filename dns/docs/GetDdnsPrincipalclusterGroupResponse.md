# GetDdnsPrincipalclusterGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Clusters** | Pointer to **[]string** | The list of equivalent DDNS principal clusters. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the DDNS Principal Cluster Group. | [optional] 
**Name** | Pointer to **string** | The name of this DDNS Principal Cluster Group. | [optional] 
**Result** | Pointer to [**DdnsPrincipalclusterGroup**](DdnsPrincipalclusterGroup.md) |  | [optional] 

## Methods

### NewGetDdnsPrincipalclusterGroupResponse

`func NewGetDdnsPrincipalclusterGroupResponse() *GetDdnsPrincipalclusterGroupResponse`

NewGetDdnsPrincipalclusterGroupResponse instantiates a new GetDdnsPrincipalclusterGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDdnsPrincipalclusterGroupResponseWithDefaults

`func NewGetDdnsPrincipalclusterGroupResponseWithDefaults() *GetDdnsPrincipalclusterGroupResponse`

NewGetDdnsPrincipalclusterGroupResponseWithDefaults instantiates a new GetDdnsPrincipalclusterGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDdnsPrincipalclusterGroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDdnsPrincipalclusterGroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDdnsPrincipalclusterGroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDdnsPrincipalclusterGroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetClusters

`func (o *GetDdnsPrincipalclusterGroupResponse) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *GetDdnsPrincipalclusterGroupResponse) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *GetDdnsPrincipalclusterGroupResponse) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *GetDdnsPrincipalclusterGroupResponse) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetComment

`func (o *GetDdnsPrincipalclusterGroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDdnsPrincipalclusterGroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDdnsPrincipalclusterGroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDdnsPrincipalclusterGroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *GetDdnsPrincipalclusterGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDdnsPrincipalclusterGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDdnsPrincipalclusterGroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDdnsPrincipalclusterGroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetDdnsPrincipalclusterGroupResponse) GetResult() DdnsPrincipalclusterGroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDdnsPrincipalclusterGroupResponse) GetResultOk() (*DdnsPrincipalclusterGroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDdnsPrincipalclusterGroupResponse) SetResult(v DdnsPrincipalclusterGroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDdnsPrincipalclusterGroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


