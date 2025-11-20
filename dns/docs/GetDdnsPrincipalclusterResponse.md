# GetDdnsPrincipalclusterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment for the DDNS Principal Cluster. | [optional] 
**Group** | Pointer to **string** | The DDNS Principal cluster group name. | [optional] 
**Name** | Pointer to **string** | The name of this DDNS Principal Cluster. | [optional] 
**Principals** | Pointer to **[]string** | The list of equivalent principals. | [optional] 
**Result** | Pointer to [**DdnsPrincipalcluster**](DdnsPrincipalcluster.md) |  | [optional] 

## Methods

### NewGetDdnsPrincipalclusterResponse

`func NewGetDdnsPrincipalclusterResponse() *GetDdnsPrincipalclusterResponse`

NewGetDdnsPrincipalclusterResponse instantiates a new GetDdnsPrincipalclusterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDdnsPrincipalclusterResponseWithDefaults

`func NewGetDdnsPrincipalclusterResponseWithDefaults() *GetDdnsPrincipalclusterResponse`

NewGetDdnsPrincipalclusterResponseWithDefaults instantiates a new GetDdnsPrincipalclusterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDdnsPrincipalclusterResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDdnsPrincipalclusterResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDdnsPrincipalclusterResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDdnsPrincipalclusterResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetDdnsPrincipalclusterResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDdnsPrincipalclusterResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDdnsPrincipalclusterResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDdnsPrincipalclusterResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroup

`func (o *GetDdnsPrincipalclusterResponse) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetDdnsPrincipalclusterResponse) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetDdnsPrincipalclusterResponse) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetDdnsPrincipalclusterResponse) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetName

`func (o *GetDdnsPrincipalclusterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDdnsPrincipalclusterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDdnsPrincipalclusterResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDdnsPrincipalclusterResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrincipals

`func (o *GetDdnsPrincipalclusterResponse) GetPrincipals() []string`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *GetDdnsPrincipalclusterResponse) GetPrincipalsOk() (*[]string, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *GetDdnsPrincipalclusterResponse) SetPrincipals(v []string)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *GetDdnsPrincipalclusterResponse) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.

### GetResult

`func (o *GetDdnsPrincipalclusterResponse) GetResult() DdnsPrincipalcluster`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDdnsPrincipalclusterResponse) GetResultOk() (*DdnsPrincipalcluster, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDdnsPrincipalclusterResponse) SetResult(v DdnsPrincipalcluster)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDdnsPrincipalclusterResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


