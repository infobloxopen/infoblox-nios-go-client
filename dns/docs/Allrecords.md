# Allrecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The record address. | [optional] [readonly] 
**Comment** | Pointer to **string** | The record comment. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. | [optional] [readonly] 
**DdnsPrincipal** | Pointer to **string** | The GSS-TSIG principal that owns this record. | [optional] [readonly] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] [readonly] 
**Disable** | Pointer to **bool** | The disable value determines if the record is disabled or not. \&quot;False\&quot; means the record is enabled. | [optional] [readonly] 
**DtcObscured** | Pointer to **string** | The specific LBDN record. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the record. | [optional] [readonly] 
**Reclaimable** | Pointer to **bool** | Determines if the record is reclaimable or not. | [optional] [readonly] 
**Record** | Pointer to **string** | The record object, if supported by the WAPI. Otherwise, the value is \&quot;None\&quot;. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for which the record is valid or being cached. The 32-bit unsigned integer represents the duration in seconds. Zero indicates that the record should not be cached. | [optional] [readonly] 
**Type** | Pointer to **string** | The record type. When searching for an unspecified record type, the search is performed for all records. On retrieval, the appliance returns \&quot;UNSUPPORTED\&quot; for unsupported records. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**View** | Pointer to **string** | Name of the DNS View in which the record resides. | [optional] [readonly] 
**Zone** | Pointer to **string** | Name of the zone in which the record resides. | [optional] [readonly] 

## Methods

### NewAllrecords

`func NewAllrecords() *Allrecords`

NewAllrecords instantiates a new Allrecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllrecordsWithDefaults

`func NewAllrecordsWithDefaults() *Allrecords`

NewAllrecordsWithDefaults instantiates a new Allrecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Allrecords) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Allrecords) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Allrecords) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Allrecords) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *Allrecords) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Allrecords) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Allrecords) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Allrecords) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *Allrecords) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Allrecords) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Allrecords) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Allrecords) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreator

`func (o *Allrecords) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Allrecords) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Allrecords) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Allrecords) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *Allrecords) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *Allrecords) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *Allrecords) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *Allrecords) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *Allrecords) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *Allrecords) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *Allrecords) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *Allrecords) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *Allrecords) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Allrecords) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Allrecords) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Allrecords) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDtcObscured

`func (o *Allrecords) GetDtcObscured() string`

GetDtcObscured returns the DtcObscured field if non-nil, zero value otherwise.

### GetDtcObscuredOk

`func (o *Allrecords) GetDtcObscuredOk() (*string, bool)`

GetDtcObscuredOk returns a tuple with the DtcObscured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcObscured

`func (o *Allrecords) SetDtcObscured(v string)`

SetDtcObscured sets DtcObscured field to given value.

### HasDtcObscured

`func (o *Allrecords) HasDtcObscured() bool`

HasDtcObscured returns a boolean if a field has been set.

### GetName

`func (o *Allrecords) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Allrecords) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Allrecords) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Allrecords) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReclaimable

`func (o *Allrecords) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *Allrecords) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *Allrecords) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *Allrecords) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetRecord

`func (o *Allrecords) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *Allrecords) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *Allrecords) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *Allrecords) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetTtl

`func (o *Allrecords) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Allrecords) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Allrecords) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Allrecords) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *Allrecords) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Allrecords) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Allrecords) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Allrecords) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *Allrecords) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Allrecords) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Allrecords) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Allrecords) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetView

`func (o *Allrecords) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Allrecords) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Allrecords) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *Allrecords) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *Allrecords) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Allrecords) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Allrecords) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Allrecords) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


