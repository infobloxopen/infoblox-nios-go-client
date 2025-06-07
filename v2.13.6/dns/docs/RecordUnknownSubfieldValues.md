# RecordUnknownSubfieldValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldValue** | Pointer to **string** | String representation of subfield value. | [optional] 
**FieldType** | Pointer to **string** | Type of field. \&quot;B\&quot;: unsigned 8-bit integer, \&quot;S\&quot;: unsigned 16-bit integer, \&quot;I\&quot;: unsigned 32-bit integer. \&quot;H\&quot;: BASE64, \&quot;6\&quot;: an IPv6 address, \&quot;4\&quot;: an IPv4 address, \&quot;N\&quot;: a domain name, \&quot;T\&quot;: text string, \&quot;X\&quot;: opaque binary data | [optional] 
**IncludeLength** | Pointer to **string** | The &#39;size of &#39;length&#39; sub-sub field to be included in RDATA. | [optional] 

## Methods

### NewRecordUnknownSubfieldValues

`func NewRecordUnknownSubfieldValues() *RecordUnknownSubfieldValues`

NewRecordUnknownSubfieldValues instantiates a new RecordUnknownSubfieldValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordUnknownSubfieldValuesWithDefaults

`func NewRecordUnknownSubfieldValuesWithDefaults() *RecordUnknownSubfieldValues`

NewRecordUnknownSubfieldValuesWithDefaults instantiates a new RecordUnknownSubfieldValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldValue

`func (o *RecordUnknownSubfieldValues) GetFieldValue() string`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *RecordUnknownSubfieldValues) GetFieldValueOk() (*string, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *RecordUnknownSubfieldValues) SetFieldValue(v string)`

SetFieldValue sets FieldValue field to given value.

### HasFieldValue

`func (o *RecordUnknownSubfieldValues) HasFieldValue() bool`

HasFieldValue returns a boolean if a field has been set.

### GetFieldType

`func (o *RecordUnknownSubfieldValues) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *RecordUnknownSubfieldValues) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *RecordUnknownSubfieldValues) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *RecordUnknownSubfieldValues) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetIncludeLength

`func (o *RecordUnknownSubfieldValues) GetIncludeLength() string`

GetIncludeLength returns the IncludeLength field if non-nil, zero value otherwise.

### GetIncludeLengthOk

`func (o *RecordUnknownSubfieldValues) GetIncludeLengthOk() (*string, bool)`

GetIncludeLengthOk returns a tuple with the IncludeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLength

`func (o *RecordUnknownSubfieldValues) SetIncludeLength(v string)`

SetIncludeLength sets IncludeLength field to given value.

### HasIncludeLength

`func (o *RecordUnknownSubfieldValues) HasIncludeLength() bool`

HasIncludeLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


