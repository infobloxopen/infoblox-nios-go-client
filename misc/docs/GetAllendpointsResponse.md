# GetAllendpointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The Grid endpoint IPv4 Address or IPv6 Address or Fully-Qualified Domain Name (FQDN). | [optional] [readonly] 
**Comment** | Pointer to **string** | The Grid endpoint descriptive comment. | [optional] [readonly] 
**Disable** | Pointer to **bool** | Determines whether a Grid endpoint is disabled or not. When this is set to False, the Grid endpoint is enabled. | [optional] [readonly] 
**SubscribingMember** | Pointer to **string** | The name of the Grid Member object that is serving Grid endpoint. | [optional] [readonly] 
**Type** | Pointer to **string** | The Grid endpoint type. | [optional] [readonly] 
**Version** | Pointer to **string** | The Grid endpoint version. | [optional] [readonly] 
**Result** | Pointer to [**Allendpoints**](Allendpoints.md) |  | [optional] 

## Methods

### NewGetAllendpointsResponse

`func NewGetAllendpointsResponse() *GetAllendpointsResponse`

NewGetAllendpointsResponse instantiates a new GetAllendpointsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllendpointsResponseWithDefaults

`func NewGetAllendpointsResponseWithDefaults() *GetAllendpointsResponse`

NewGetAllendpointsResponseWithDefaults instantiates a new GetAllendpointsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAllendpointsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAllendpointsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAllendpointsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAllendpointsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetAllendpointsResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetAllendpointsResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetAllendpointsResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetAllendpointsResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *GetAllendpointsResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAllendpointsResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAllendpointsResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAllendpointsResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetAllendpointsResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetAllendpointsResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetAllendpointsResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetAllendpointsResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetSubscribingMember

`func (o *GetAllendpointsResponse) GetSubscribingMember() string`

GetSubscribingMember returns the SubscribingMember field if non-nil, zero value otherwise.

### GetSubscribingMemberOk

`func (o *GetAllendpointsResponse) GetSubscribingMemberOk() (*string, bool)`

GetSubscribingMemberOk returns a tuple with the SubscribingMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribingMember

`func (o *GetAllendpointsResponse) SetSubscribingMember(v string)`

SetSubscribingMember sets SubscribingMember field to given value.

### HasSubscribingMember

`func (o *GetAllendpointsResponse) HasSubscribingMember() bool`

HasSubscribingMember returns a boolean if a field has been set.

### GetType

`func (o *GetAllendpointsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAllendpointsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAllendpointsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAllendpointsResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *GetAllendpointsResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetAllendpointsResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetAllendpointsResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetAllendpointsResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetResult

`func (o *GetAllendpointsResponse) GetResult() Allendpoints`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAllendpointsResponse) GetResultOk() (*Allendpoints, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAllendpointsResponse) SetResult(v Allendpoints)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAllendpointsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


