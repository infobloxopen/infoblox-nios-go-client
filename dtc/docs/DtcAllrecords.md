# DtcAllrecords

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The record comment. | [optional] [readonly] 
**Disable** | Pointer to **bool** | The disable value determines if the record is disabled or not. \&quot;False\&quot; means the record is enabled. | [optional] [readonly] 
**DtcServer** | Pointer to **string** | The name of the DTC Server object with which the record is associated. | [optional] [readonly] 
**Record** | Pointer to **string** | The record object, if supported by the WAPI. Otherwise, the value is \&quot;None\&quot;. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The TTL value of the record associated with the DTC AllRecords object. | [optional] 
**Type** | Pointer to **string** | The record type. When searching for an unspecified record type, the search is performed for all records. On retrieval, the appliance returns \&quot;UNSUPPORTED\&quot; for unsupported records. | [optional] [readonly] 

## Methods

### NewDtcAllrecords

`func NewDtcAllrecords() *DtcAllrecords`

NewDtcAllrecords instantiates a new DtcAllrecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcAllrecordsWithDefaults

`func NewDtcAllrecordsWithDefaults() *DtcAllrecords`

NewDtcAllrecordsWithDefaults instantiates a new DtcAllrecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcAllrecords) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcAllrecords) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcAllrecords) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcAllrecords) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *DtcAllrecords) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcAllrecords) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcAllrecords) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcAllrecords) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *DtcAllrecords) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DtcAllrecords) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DtcAllrecords) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DtcAllrecords) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDtcServer

`func (o *DtcAllrecords) GetDtcServer() string`

GetDtcServer returns the DtcServer field if non-nil, zero value otherwise.

### GetDtcServerOk

`func (o *DtcAllrecords) GetDtcServerOk() (*string, bool)`

GetDtcServerOk returns a tuple with the DtcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcServer

`func (o *DtcAllrecords) SetDtcServer(v string)`

SetDtcServer sets DtcServer field to given value.

### HasDtcServer

`func (o *DtcAllrecords) HasDtcServer() bool`

HasDtcServer returns a boolean if a field has been set.

### GetRecord

`func (o *DtcAllrecords) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *DtcAllrecords) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *DtcAllrecords) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *DtcAllrecords) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetTtl

`func (o *DtcAllrecords) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DtcAllrecords) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DtcAllrecords) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DtcAllrecords) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *DtcAllrecords) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtcAllrecords) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtcAllrecords) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DtcAllrecords) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


