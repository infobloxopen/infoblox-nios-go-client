# MemberDnsSortlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The source address of a sortlist object. | [optional] 
**MatchList** | Pointer to **[]string** | The match list of a sortlist. | [optional] 

## Methods

### NewMemberDnsSortlist

`func NewMemberDnsSortlist() *MemberDnsSortlist`

NewMemberDnsSortlist instantiates a new MemberDnsSortlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsSortlistWithDefaults

`func NewMemberDnsSortlistWithDefaults() *MemberDnsSortlist`

NewMemberDnsSortlistWithDefaults instantiates a new MemberDnsSortlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MemberDnsSortlist) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MemberDnsSortlist) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MemberDnsSortlist) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MemberDnsSortlist) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetMatchList

`func (o *MemberDnsSortlist) GetMatchList() []string`

GetMatchList returns the MatchList field if non-nil, zero value otherwise.

### GetMatchListOk

`func (o *MemberDnsSortlist) GetMatchListOk() (*[]string, bool)`

GetMatchListOk returns a tuple with the MatchList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchList

`func (o *MemberDnsSortlist) SetMatchList(v []string)`

SetMatchList sets MatchList field to given value.

### HasMatchList

`func (o *MemberDnsSortlist) HasMatchList() bool`

HasMatchList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


