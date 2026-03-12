# RecordSvcbSvcParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SvcKey** | Pointer to **string** | Svc param key. | [optional] 
**SvcValue** | Pointer to **[]string** | Svc param value. | [optional] 
**Mandatory** | Pointer to **bool** | Specifies if this is mandatory key for this RR. | [optional] 

## Methods

### NewRecordSvcbSvcParameters

`func NewRecordSvcbSvcParameters() *RecordSvcbSvcParameters`

NewRecordSvcbSvcParameters instantiates a new RecordSvcbSvcParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordSvcbSvcParametersWithDefaults

`func NewRecordSvcbSvcParametersWithDefaults() *RecordSvcbSvcParameters`

NewRecordSvcbSvcParametersWithDefaults instantiates a new RecordSvcbSvcParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSvcKey

`func (o *RecordSvcbSvcParameters) GetSvcKey() string`

GetSvcKey returns the SvcKey field if non-nil, zero value otherwise.

### GetSvcKeyOk

`func (o *RecordSvcbSvcParameters) GetSvcKeyOk() (*string, bool)`

GetSvcKeyOk returns a tuple with the SvcKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcKey

`func (o *RecordSvcbSvcParameters) SetSvcKey(v string)`

SetSvcKey sets SvcKey field to given value.

### HasSvcKey

`func (o *RecordSvcbSvcParameters) HasSvcKey() bool`

HasSvcKey returns a boolean if a field has been set.

### GetSvcValue

`func (o *RecordSvcbSvcParameters) GetSvcValue() []string`

GetSvcValue returns the SvcValue field if non-nil, zero value otherwise.

### GetSvcValueOk

`func (o *RecordSvcbSvcParameters) GetSvcValueOk() (*[]string, bool)`

GetSvcValueOk returns a tuple with the SvcValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcValue

`func (o *RecordSvcbSvcParameters) SetSvcValue(v []string)`

SetSvcValue sets SvcValue field to given value.

### HasSvcValue

`func (o *RecordSvcbSvcParameters) HasSvcValue() bool`

HasSvcValue returns a boolean if a field has been set.

### GetMandatory

`func (o *RecordSvcbSvcParameters) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *RecordSvcbSvcParameters) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *RecordSvcbSvcParameters) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *RecordSvcbSvcParameters) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


