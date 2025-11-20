# GetSuperhostchildResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AssociatedObject** | Pointer to **string** | The record object, if supported by the WAPI. Otherwise, the value is \&quot;None\&quot;. | [optional] [readonly] 
**Comment** | Pointer to **string** | The record comment. | [optional] [readonly] 
**CreationTimestamp** | Pointer to **int64** | Time at which DNS RR was created. | [optional] [readonly] 
**Data** | Pointer to **string** | Specific data of DNS/DHCP records. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | True if the child DNS/DHCP object is disabled. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the associated DNS/DHCP object. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this network record resides. | [optional] [readonly] 
**Parent** | Pointer to **string** | Name of the Super Host object in which record resides. | [optional] [readonly] 
**RecordParent** | Pointer to **string** | Name of a parent zone/network. | [optional] [readonly] 
**Type** | Pointer to **string** | The record type. When searching for an unspecified record type, the search is performed for all records. | [optional] [readonly] 
**View** | Pointer to **string** | Name of the DNS View in which the record resides. | [optional] [readonly] 
**Result** | Pointer to [**Superhostchild**](Superhostchild.md) |  | [optional] 

## Methods

### NewGetSuperhostchildResponse

`func NewGetSuperhostchildResponse() *GetSuperhostchildResponse`

NewGetSuperhostchildResponse instantiates a new GetSuperhostchildResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSuperhostchildResponseWithDefaults

`func NewGetSuperhostchildResponseWithDefaults() *GetSuperhostchildResponse`

NewGetSuperhostchildResponseWithDefaults instantiates a new GetSuperhostchildResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSuperhostchildResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSuperhostchildResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSuperhostchildResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSuperhostchildResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAssociatedObject

`func (o *GetSuperhostchildResponse) GetAssociatedObject() string`

GetAssociatedObject returns the AssociatedObject field if non-nil, zero value otherwise.

### GetAssociatedObjectOk

`func (o *GetSuperhostchildResponse) GetAssociatedObjectOk() (*string, bool)`

GetAssociatedObjectOk returns a tuple with the AssociatedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObject

`func (o *GetSuperhostchildResponse) SetAssociatedObject(v string)`

SetAssociatedObject sets AssociatedObject field to given value.

### HasAssociatedObject

`func (o *GetSuperhostchildResponse) HasAssociatedObject() bool`

HasAssociatedObject returns a boolean if a field has been set.

### GetComment

`func (o *GetSuperhostchildResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetSuperhostchildResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetSuperhostchildResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetSuperhostchildResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTimestamp

`func (o *GetSuperhostchildResponse) GetCreationTimestamp() int64`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *GetSuperhostchildResponse) GetCreationTimestampOk() (*int64, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *GetSuperhostchildResponse) SetCreationTimestamp(v int64)`

SetCreationTimestamp sets CreationTimestamp field to given value.

### HasCreationTimestamp

`func (o *GetSuperhostchildResponse) HasCreationTimestamp() bool`

HasCreationTimestamp returns a boolean if a field has been set.

### GetData

`func (o *GetSuperhostchildResponse) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSuperhostchildResponse) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSuperhostchildResponse) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *GetSuperhostchildResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDisabled

`func (o *GetSuperhostchildResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetSuperhostchildResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetSuperhostchildResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetSuperhostchildResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetName

`func (o *GetSuperhostchildResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSuperhostchildResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSuperhostchildResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSuperhostchildResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetSuperhostchildResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetSuperhostchildResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetSuperhostchildResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetSuperhostchildResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetParent

`func (o *GetSuperhostchildResponse) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetSuperhostchildResponse) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetSuperhostchildResponse) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetSuperhostchildResponse) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetRecordParent

`func (o *GetSuperhostchildResponse) GetRecordParent() string`

GetRecordParent returns the RecordParent field if non-nil, zero value otherwise.

### GetRecordParentOk

`func (o *GetSuperhostchildResponse) GetRecordParentOk() (*string, bool)`

GetRecordParentOk returns a tuple with the RecordParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordParent

`func (o *GetSuperhostchildResponse) SetRecordParent(v string)`

SetRecordParent sets RecordParent field to given value.

### HasRecordParent

`func (o *GetSuperhostchildResponse) HasRecordParent() bool`

HasRecordParent returns a boolean if a field has been set.

### GetType

`func (o *GetSuperhostchildResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSuperhostchildResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSuperhostchildResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetSuperhostchildResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetView

`func (o *GetSuperhostchildResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetSuperhostchildResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetSuperhostchildResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetSuperhostchildResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetResult

`func (o *GetSuperhostchildResponse) GetResult() Superhostchild`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSuperhostchildResponse) GetResultOk() (*Superhostchild, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSuperhostchildResponse) SetResult(v Superhostchild)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSuperhostchildResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


