# RecordHttpsSvcParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SvcKey** | Pointer to **string** | Svc param key. | [optional] 
**SvcValue** | Pointer to **[]string** | Svc param value. | [optional] 
**Mandatory** | Pointer to **bool** | Specifies if this is mandatory key for this RR. | [optional] 

## Methods

### NewRecordHttpsSvcParameters

`func NewRecordHttpsSvcParameters() *RecordHttpsSvcParameters`

NewRecordHttpsSvcParameters instantiates a new RecordHttpsSvcParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordHttpsSvcParametersWithDefaults

`func NewRecordHttpsSvcParametersWithDefaults() *RecordHttpsSvcParameters`

NewRecordHttpsSvcParametersWithDefaults instantiates a new RecordHttpsSvcParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSvcKey

`func (o *RecordHttpsSvcParameters) GetSvcKey() string`

GetSvcKey returns the SvcKey field if non-nil, zero value otherwise.

### GetSvcKeyOk

`func (o *RecordHttpsSvcParameters) GetSvcKeyOk() (*string, bool)`

GetSvcKeyOk returns a tuple with the SvcKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcKey

`func (o *RecordHttpsSvcParameters) SetSvcKey(v string)`

SetSvcKey sets SvcKey field to given value.

### HasSvcKey

`func (o *RecordHttpsSvcParameters) HasSvcKey() bool`

HasSvcKey returns a boolean if a field has been set.

### GetSvcValue

`func (o *RecordHttpsSvcParameters) GetSvcValue() []string`

GetSvcValue returns the SvcValue field if non-nil, zero value otherwise.

### GetSvcValueOk

`func (o *RecordHttpsSvcParameters) GetSvcValueOk() (*[]string, bool)`

GetSvcValueOk returns a tuple with the SvcValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcValue

`func (o *RecordHttpsSvcParameters) SetSvcValue(v []string)`

SetSvcValue sets SvcValue field to given value.

### HasSvcValue

`func (o *RecordHttpsSvcParameters) HasSvcValue() bool`

HasSvcValue returns a boolean if a field has been set.

### GetMandatory

`func (o *RecordHttpsSvcParameters) GetMandatory() bool`

GetMandatory returns the Mandatory field if non-nil, zero value otherwise.

### GetMandatoryOk

`func (o *RecordHttpsSvcParameters) GetMandatoryOk() (*bool, bool)`

GetMandatoryOk returns a tuple with the Mandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatory

`func (o *RecordHttpsSvcParameters) SetMandatory(v bool)`

SetMandatory sets Mandatory field to given value.

### HasMandatory

`func (o *RecordHttpsSvcParameters) HasMandatory() bool`

HasMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


