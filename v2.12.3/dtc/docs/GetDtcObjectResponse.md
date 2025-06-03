# GetDtcObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AbstractType** | Pointer to **string** | The abstract object type. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for the DTC object; maximum 256 characters. | [optional] [readonly] 
**DisplayType** | Pointer to **string** | The display object type. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv4AddressList** | Pointer to **[]string** | The list of IPv4 addresses. | [optional] [readonly] 
**Ipv6AddressList** | Pointer to **[]string** | The list of IPv6 addresses. | [optional] [readonly] 
**Name** | Pointer to **string** | The display name of the DTC object. | [optional] [readonly] 
**Object** | Pointer to **string** | The specific DTC object. | [optional] [readonly] 
**Status** | Pointer to **string** | The availability color status. | [optional] [readonly] 
**StatusTime** | Pointer to **int64** | The timestamp when status or health was last determined. | [optional] [readonly] 
**Result** | Pointer to [**DtcObject**](DtcObject.md) |  | [optional] 

## Methods

### NewGetDtcObjectResponse

`func NewGetDtcObjectResponse() *GetDtcObjectResponse`

NewGetDtcObjectResponse instantiates a new GetDtcObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcObjectResponseWithDefaults

`func NewGetDtcObjectResponseWithDefaults() *GetDtcObjectResponse`

NewGetDtcObjectResponseWithDefaults instantiates a new GetDtcObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcObjectResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcObjectResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcObjectResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcObjectResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAbstractType

`func (o *GetDtcObjectResponse) GetAbstractType() string`

GetAbstractType returns the AbstractType field if non-nil, zero value otherwise.

### GetAbstractTypeOk

`func (o *GetDtcObjectResponse) GetAbstractTypeOk() (*string, bool)`

GetAbstractTypeOk returns a tuple with the AbstractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstractType

`func (o *GetDtcObjectResponse) SetAbstractType(v string)`

SetAbstractType sets AbstractType field to given value.

### HasAbstractType

`func (o *GetDtcObjectResponse) HasAbstractType() bool`

HasAbstractType returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcObjectResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcObjectResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcObjectResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcObjectResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisplayType

`func (o *GetDtcObjectResponse) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *GetDtcObjectResponse) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *GetDtcObjectResponse) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *GetDtcObjectResponse) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetDtcObjectResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetDtcObjectResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetDtcObjectResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetDtcObjectResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIpv4AddressList

`func (o *GetDtcObjectResponse) GetIpv4AddressList() []string`

GetIpv4AddressList returns the Ipv4AddressList field if non-nil, zero value otherwise.

### GetIpv4AddressListOk

`func (o *GetDtcObjectResponse) GetIpv4AddressListOk() (*[]string, bool)`

GetIpv4AddressListOk returns a tuple with the Ipv4AddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddressList

`func (o *GetDtcObjectResponse) SetIpv4AddressList(v []string)`

SetIpv4AddressList sets Ipv4AddressList field to given value.

### HasIpv4AddressList

`func (o *GetDtcObjectResponse) HasIpv4AddressList() bool`

HasIpv4AddressList returns a boolean if a field has been set.

### GetIpv6AddressList

`func (o *GetDtcObjectResponse) GetIpv6AddressList() []string`

GetIpv6AddressList returns the Ipv6AddressList field if non-nil, zero value otherwise.

### GetIpv6AddressListOk

`func (o *GetDtcObjectResponse) GetIpv6AddressListOk() (*[]string, bool)`

GetIpv6AddressListOk returns a tuple with the Ipv6AddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6AddressList

`func (o *GetDtcObjectResponse) SetIpv6AddressList(v []string)`

SetIpv6AddressList sets Ipv6AddressList field to given value.

### HasIpv6AddressList

`func (o *GetDtcObjectResponse) HasIpv6AddressList() bool`

HasIpv6AddressList returns a boolean if a field has been set.

### GetName

`func (o *GetDtcObjectResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDtcObjectResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDtcObjectResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDtcObjectResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *GetDtcObjectResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GetDtcObjectResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GetDtcObjectResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GetDtcObjectResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatus

`func (o *GetDtcObjectResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDtcObjectResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDtcObjectResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDtcObjectResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTime

`func (o *GetDtcObjectResponse) GetStatusTime() int64`

GetStatusTime returns the StatusTime field if non-nil, zero value otherwise.

### GetStatusTimeOk

`func (o *GetDtcObjectResponse) GetStatusTimeOk() (*int64, bool)`

GetStatusTimeOk returns a tuple with the StatusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTime

`func (o *GetDtcObjectResponse) SetStatusTime(v int64)`

SetStatusTime sets StatusTime field to given value.

### HasStatusTime

`func (o *GetDtcObjectResponse) HasStatusTime() bool`

HasStatusTime returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcObjectResponse) GetResult() DtcObject`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcObjectResponse) GetResultOk() (*DtcObject, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcObjectResponse) SetResult(v DtcObject)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcObjectResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


