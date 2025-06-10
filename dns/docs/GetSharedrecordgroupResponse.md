# GetSharedrecordgroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment of this shared record group. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of this shared record group. | [optional] 
**RecordNamePolicy** | Pointer to **string** | The record name policy of this shared record group. | [optional] 
**UseRecordNamePolicy** | Pointer to **bool** | Use flag for: record_name_policy | [optional] 
**ZoneAssociations** | Pointer to [**[]SharedrecordgroupZoneAssociations**](SharedrecordgroupZoneAssociations.md) | The list of zones associated with this shared record group. Starting from NIOS-9.0.6, this field has been updated to a structure that includes FQDN and DNS view details. | [optional] 
**Result** | Pointer to [**Sharedrecordgroup**](Sharedrecordgroup.md) |  | [optional] 

## Methods

### NewGetSharedrecordgroupResponse

`func NewGetSharedrecordgroupResponse() *GetSharedrecordgroupResponse`

NewGetSharedrecordgroupResponse instantiates a new GetSharedrecordgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSharedrecordgroupResponseWithDefaults

`func NewGetSharedrecordgroupResponseWithDefaults() *GetSharedrecordgroupResponse`

NewGetSharedrecordgroupResponseWithDefaults instantiates a new GetSharedrecordgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSharedrecordgroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSharedrecordgroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSharedrecordgroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSharedrecordgroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetSharedrecordgroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSharedrecordgroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSharedrecordgroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSharedrecordgroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetSharedrecordgroupResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetSharedrecordgroupResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetSharedrecordgroupResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetSharedrecordgroupResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetName

`func (o *GetSharedrecordgroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSharedrecordgroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSharedrecordgroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSharedrecordgroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecordNamePolicy

`func (o *GetSharedrecordgroupResponse) GetRecordNamePolicy() string`

GetRecordNamePolicy returns the RecordNamePolicy field if non-nil, zero value otherwise.

### GetRecordNamePolicyOk

`func (o *GetSharedrecordgroupResponse) GetRecordNamePolicyOk() (*string, bool)`

GetRecordNamePolicyOk returns a tuple with the RecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNamePolicy

`func (o *GetSharedrecordgroupResponse) SetRecordNamePolicy(v string)`

SetRecordNamePolicy sets RecordNamePolicy field to given value.

### HasRecordNamePolicy

`func (o *GetSharedrecordgroupResponse) HasRecordNamePolicy() bool`

HasRecordNamePolicy returns a boolean if a field has been set.

### GetUseRecordNamePolicy

`func (o *GetSharedrecordgroupResponse) GetUseRecordNamePolicy() bool`

GetUseRecordNamePolicy returns the UseRecordNamePolicy field if non-nil, zero value otherwise.

### GetUseRecordNamePolicyOk

`func (o *GetSharedrecordgroupResponse) GetUseRecordNamePolicyOk() (*bool, bool)`

GetUseRecordNamePolicyOk returns a tuple with the UseRecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecordNamePolicy

`func (o *GetSharedrecordgroupResponse) SetUseRecordNamePolicy(v bool)`

SetUseRecordNamePolicy sets UseRecordNamePolicy field to given value.

### HasUseRecordNamePolicy

`func (o *GetSharedrecordgroupResponse) HasUseRecordNamePolicy() bool`

HasUseRecordNamePolicy returns a boolean if a field has been set.

### GetZoneAssociations

`func (o *GetSharedrecordgroupResponse) GetZoneAssociations() []SharedrecordgroupZoneAssociations`

GetZoneAssociations returns the ZoneAssociations field if non-nil, zero value otherwise.

### GetZoneAssociationsOk

`func (o *GetSharedrecordgroupResponse) GetZoneAssociationsOk() (*[]SharedrecordgroupZoneAssociations, bool)`

GetZoneAssociationsOk returns a tuple with the ZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAssociations

`func (o *GetSharedrecordgroupResponse) SetZoneAssociations(v []SharedrecordgroupZoneAssociations)`

SetZoneAssociations sets ZoneAssociations field to given value.

### HasZoneAssociations

`func (o *GetSharedrecordgroupResponse) HasZoneAssociations() bool`

HasZoneAssociations returns a boolean if a field has been set.

### GetResult

`func (o *GetSharedrecordgroupResponse) GetResult() Sharedrecordgroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSharedrecordgroupResponse) GetResultOk() (*Sharedrecordgroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSharedrecordgroupResponse) SetResult(v Sharedrecordgroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSharedrecordgroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


