# DtcObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AbstractType** | Pointer to **string** | The abstract object type. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for the DTC object; maximum 256 characters. | [optional] [readonly] 
**DisplayType** | Pointer to **string** | The display object type. | [optional] [readonly] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv4AddressList** | Pointer to **[]string** | The list of IPv4 addresses. | [optional] [readonly] 
**Ipv6AddressList** | Pointer to **[]string** | The list of IPv6 addresses. | [optional] [readonly] 
**Name** | Pointer to **string** | The display name of the DTC object. | [optional] [readonly] 
**Object** | Pointer to **string** | The specific DTC object. | [optional] [readonly] 
**Status** | Pointer to **string** | The availability color status. | [optional] [readonly] 
**StatusTime** | Pointer to **int64** | The timestamp when status or health was last determined. | [optional] [readonly] 

## Methods

### NewDtcObject

`func NewDtcObject() *DtcObject`

NewDtcObject instantiates a new DtcObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcObjectWithDefaults

`func NewDtcObjectWithDefaults() *DtcObject`

NewDtcObjectWithDefaults instantiates a new DtcObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcObject) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcObject) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcObject) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcObject) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAbstractType

`func (o *DtcObject) GetAbstractType() string`

GetAbstractType returns the AbstractType field if non-nil, zero value otherwise.

### GetAbstractTypeOk

`func (o *DtcObject) GetAbstractTypeOk() (*string, bool)`

GetAbstractTypeOk returns a tuple with the AbstractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstractType

`func (o *DtcObject) SetAbstractType(v string)`

SetAbstractType sets AbstractType field to given value.

### HasAbstractType

`func (o *DtcObject) HasAbstractType() bool`

HasAbstractType returns a boolean if a field has been set.

### GetComment

`func (o *DtcObject) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcObject) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcObject) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcObject) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisplayType

`func (o *DtcObject) GetDisplayType() string`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *DtcObject) GetDisplayTypeOk() (*string, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *DtcObject) SetDisplayType(v string)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *DtcObject) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetExtAttrs

`func (o *DtcObject) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *DtcObject) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *DtcObject) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *DtcObject) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIpv4AddressList

`func (o *DtcObject) GetIpv4AddressList() []string`

GetIpv4AddressList returns the Ipv4AddressList field if non-nil, zero value otherwise.

### GetIpv4AddressListOk

`func (o *DtcObject) GetIpv4AddressListOk() (*[]string, bool)`

GetIpv4AddressListOk returns a tuple with the Ipv4AddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4AddressList

`func (o *DtcObject) SetIpv4AddressList(v []string)`

SetIpv4AddressList sets Ipv4AddressList field to given value.

### HasIpv4AddressList

`func (o *DtcObject) HasIpv4AddressList() bool`

HasIpv4AddressList returns a boolean if a field has been set.

### GetIpv6AddressList

`func (o *DtcObject) GetIpv6AddressList() []string`

GetIpv6AddressList returns the Ipv6AddressList field if non-nil, zero value otherwise.

### GetIpv6AddressListOk

`func (o *DtcObject) GetIpv6AddressListOk() (*[]string, bool)`

GetIpv6AddressListOk returns a tuple with the Ipv6AddressList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6AddressList

`func (o *DtcObject) SetIpv6AddressList(v []string)`

SetIpv6AddressList sets Ipv6AddressList field to given value.

### HasIpv6AddressList

`func (o *DtcObject) HasIpv6AddressList() bool`

HasIpv6AddressList returns a boolean if a field has been set.

### GetName

`func (o *DtcObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcObject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcObject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *DtcObject) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DtcObject) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DtcObject) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *DtcObject) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatus

`func (o *DtcObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DtcObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DtcObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DtcObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTime

`func (o *DtcObject) GetStatusTime() int64`

GetStatusTime returns the StatusTime field if non-nil, zero value otherwise.

### GetStatusTimeOk

`func (o *DtcObject) GetStatusTimeOk() (*int64, bool)`

GetStatusTimeOk returns a tuple with the StatusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTime

`func (o *DtcObject) SetStatusTime(v int64)`

SetStatusTime sets StatusTime field to given value.

### HasStatusTime

`func (o *DtcObject) HasStatusTime() bool`

HasStatusTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


