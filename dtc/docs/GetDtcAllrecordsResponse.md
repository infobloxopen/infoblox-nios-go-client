# GetDtcAllrecordsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The record comment. | [optional] [readonly] 
**Disable** | Pointer to **bool** | The disable value determines if the record is disabled or not. \&quot;False\&quot; means the record is enabled. | [optional] [readonly] 
**DtcServer** | Pointer to **string** | The name of the DTC Server object with which the record is associated. | [optional] [readonly] 
**Record** | Pointer to **string** | The record object, if supported by the WAPI. Otherwise, the value is \&quot;None\&quot;. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The TTL value of the record associated with the DTC AllRecords object. | [optional] 
**Type** | Pointer to **string** | The record type. When searching for an unspecified record type, the search is performed for all records. On retrieval, the appliance returns \&quot;UNSUPPORTED\&quot; for unsupported records. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**DtcAllrecords**](DtcAllrecords.md) |  | [optional] 

## Methods

### NewGetDtcAllrecordsResponse

`func NewGetDtcAllrecordsResponse() *GetDtcAllrecordsResponse`

NewGetDtcAllrecordsResponse instantiates a new GetDtcAllrecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcAllrecordsResponseWithDefaults

`func NewGetDtcAllrecordsResponseWithDefaults() *GetDtcAllrecordsResponse`

NewGetDtcAllrecordsResponseWithDefaults instantiates a new GetDtcAllrecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcAllrecordsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcAllrecordsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcAllrecordsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcAllrecordsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcAllrecordsResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcAllrecordsResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcAllrecordsResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcAllrecordsResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetDtcAllrecordsResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetDtcAllrecordsResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetDtcAllrecordsResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetDtcAllrecordsResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDtcServer

`func (o *GetDtcAllrecordsResponse) GetDtcServer() string`

GetDtcServer returns the DtcServer field if non-nil, zero value otherwise.

### GetDtcServerOk

`func (o *GetDtcAllrecordsResponse) GetDtcServerOk() (*string, bool)`

GetDtcServerOk returns a tuple with the DtcServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtcServer

`func (o *GetDtcAllrecordsResponse) SetDtcServer(v string)`

SetDtcServer sets DtcServer field to given value.

### HasDtcServer

`func (o *GetDtcAllrecordsResponse) HasDtcServer() bool`

HasDtcServer returns a boolean if a field has been set.

### GetRecord

`func (o *GetDtcAllrecordsResponse) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *GetDtcAllrecordsResponse) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *GetDtcAllrecordsResponse) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *GetDtcAllrecordsResponse) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetTtl

`func (o *GetDtcAllrecordsResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetDtcAllrecordsResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetDtcAllrecordsResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetDtcAllrecordsResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *GetDtcAllrecordsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDtcAllrecordsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDtcAllrecordsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDtcAllrecordsResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *GetDtcAllrecordsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDtcAllrecordsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDtcAllrecordsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDtcAllrecordsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcAllrecordsResponse) GetResult() DtcAllrecords`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcAllrecordsResponse) GetResultOk() (*DtcAllrecords, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcAllrecordsResponse) SetResult(v DtcAllrecords)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcAllrecordsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


