# Sharedrecordgroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The descriptive comment of this shared record group. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of this shared record group. | [optional] 
**RecordNamePolicy** | Pointer to **string** | The record name policy of this shared record group. | [optional] 
**UseRecordNamePolicy** | Pointer to **bool** | Use flag for: record_name_policy | [optional] 
**ZoneAssociations** | Pointer to [**[]SharedrecordgroupZoneAssociations**](SharedrecordgroupZoneAssociations.md) | The list of zones associated with this shared record group. Starting from NIOS-9.0.6, this field has been updated to a structure that includes FQDN and DNS view details. | [optional] 

## Methods

### NewSharedrecordgroup

`func NewSharedrecordgroup() *Sharedrecordgroup`

NewSharedrecordgroup instantiates a new Sharedrecordgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedrecordgroupWithDefaults

`func NewSharedrecordgroupWithDefaults() *Sharedrecordgroup`

NewSharedrecordgroupWithDefaults instantiates a new Sharedrecordgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Sharedrecordgroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Sharedrecordgroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Sharedrecordgroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Sharedrecordgroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Sharedrecordgroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Sharedrecordgroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Sharedrecordgroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Sharedrecordgroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Sharedrecordgroup) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Sharedrecordgroup) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Sharedrecordgroup) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Sharedrecordgroup) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Sharedrecordgroup) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Sharedrecordgroup) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Sharedrecordgroup) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Sharedrecordgroup) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Sharedrecordgroup) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Sharedrecordgroup) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Sharedrecordgroup) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Sharedrecordgroup) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *Sharedrecordgroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sharedrecordgroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sharedrecordgroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Sharedrecordgroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecordNamePolicy

`func (o *Sharedrecordgroup) GetRecordNamePolicy() string`

GetRecordNamePolicy returns the RecordNamePolicy field if non-nil, zero value otherwise.

### GetRecordNamePolicyOk

`func (o *Sharedrecordgroup) GetRecordNamePolicyOk() (*string, bool)`

GetRecordNamePolicyOk returns a tuple with the RecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNamePolicy

`func (o *Sharedrecordgroup) SetRecordNamePolicy(v string)`

SetRecordNamePolicy sets RecordNamePolicy field to given value.

### HasRecordNamePolicy

`func (o *Sharedrecordgroup) HasRecordNamePolicy() bool`

HasRecordNamePolicy returns a boolean if a field has been set.

### GetUseRecordNamePolicy

`func (o *Sharedrecordgroup) GetUseRecordNamePolicy() bool`

GetUseRecordNamePolicy returns the UseRecordNamePolicy field if non-nil, zero value otherwise.

### GetUseRecordNamePolicyOk

`func (o *Sharedrecordgroup) GetUseRecordNamePolicyOk() (*bool, bool)`

GetUseRecordNamePolicyOk returns a tuple with the UseRecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecordNamePolicy

`func (o *Sharedrecordgroup) SetUseRecordNamePolicy(v bool)`

SetUseRecordNamePolicy sets UseRecordNamePolicy field to given value.

### HasUseRecordNamePolicy

`func (o *Sharedrecordgroup) HasUseRecordNamePolicy() bool`

HasUseRecordNamePolicy returns a boolean if a field has been set.

### GetZoneAssociations

`func (o *Sharedrecordgroup) GetZoneAssociations() []SharedrecordgroupZoneAssociations`

GetZoneAssociations returns the ZoneAssociations field if non-nil, zero value otherwise.

### GetZoneAssociationsOk

`func (o *Sharedrecordgroup) GetZoneAssociationsOk() (*[]SharedrecordgroupZoneAssociations, bool)`

GetZoneAssociationsOk returns a tuple with the ZoneAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAssociations

`func (o *Sharedrecordgroup) SetZoneAssociations(v []SharedrecordgroupZoneAssociations)`

SetZoneAssociations sets ZoneAssociations field to given value.

### HasZoneAssociations

`func (o *Sharedrecordgroup) HasZoneAssociations() bool`

HasZoneAssociations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


