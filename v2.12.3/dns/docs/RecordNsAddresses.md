# RecordNsAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address of the Zone Name Server. | [optional] 
**AutoCreatePtr** | Pointer to **bool** | Flag to indicate if ptr records need to be auto created. | [optional] 

## Methods

### NewRecordNsAddresses

`func NewRecordNsAddresses() *RecordNsAddresses`

NewRecordNsAddresses instantiates a new RecordNsAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordNsAddressesWithDefaults

`func NewRecordNsAddressesWithDefaults() *RecordNsAddresses`

NewRecordNsAddressesWithDefaults instantiates a new RecordNsAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RecordNsAddresses) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RecordNsAddresses) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RecordNsAddresses) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RecordNsAddresses) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAutoCreatePtr

`func (o *RecordNsAddresses) GetAutoCreatePtr() bool`

GetAutoCreatePtr returns the AutoCreatePtr field if non-nil, zero value otherwise.

### GetAutoCreatePtrOk

`func (o *RecordNsAddresses) GetAutoCreatePtrOk() (*bool, bool)`

GetAutoCreatePtrOk returns a tuple with the AutoCreatePtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreatePtr

`func (o *RecordNsAddresses) SetAutoCreatePtr(v bool)`

SetAutoCreatePtr sets AutoCreatePtr field to given value.

### HasAutoCreatePtr

`func (o *RecordNsAddresses) HasAutoCreatePtr() bool`

HasAutoCreatePtr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


