# Dbsnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The descriptive comment. | [optional] [readonly] 
**Timestamp** | Pointer to **int64** | The time when the latest OneDB snapshot was taken in Epoch seconds format. | [optional] [readonly] 

## Methods

### NewDbsnapshot

`func NewDbsnapshot() *Dbsnapshot`

NewDbsnapshot instantiates a new Dbsnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbsnapshotWithDefaults

`func NewDbsnapshotWithDefaults() *Dbsnapshot`

NewDbsnapshotWithDefaults instantiates a new Dbsnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Dbsnapshot) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Dbsnapshot) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Dbsnapshot) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Dbsnapshot) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Dbsnapshot) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Dbsnapshot) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Dbsnapshot) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Dbsnapshot) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTimestamp

`func (o *Dbsnapshot) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Dbsnapshot) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Dbsnapshot) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Dbsnapshot) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


