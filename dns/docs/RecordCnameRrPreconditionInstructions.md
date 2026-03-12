# RecordCnameRrPreconditionInstructions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to **string** | Condition type: exist or notexist. | [optional] 
**Name** | Pointer to **string** | Record name | [optional] 
**Type** | Pointer to **string** | Record type | [optional] 
**Rdata** | Pointer to **string** | Record data (optional) | [optional] 
**Action** | Pointer to **string** | Action to perform if condition is met: none or delete. | [optional] 

## Methods

### NewRecordCnameRrPreconditionInstructions

`func NewRecordCnameRrPreconditionInstructions() *RecordCnameRrPreconditionInstructions`

NewRecordCnameRrPreconditionInstructions instantiates a new RecordCnameRrPreconditionInstructions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordCnameRrPreconditionInstructionsWithDefaults

`func NewRecordCnameRrPreconditionInstructionsWithDefaults() *RecordCnameRrPreconditionInstructions`

NewRecordCnameRrPreconditionInstructionsWithDefaults instantiates a new RecordCnameRrPreconditionInstructions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *RecordCnameRrPreconditionInstructions) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RecordCnameRrPreconditionInstructions) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RecordCnameRrPreconditionInstructions) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RecordCnameRrPreconditionInstructions) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetName

`func (o *RecordCnameRrPreconditionInstructions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordCnameRrPreconditionInstructions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordCnameRrPreconditionInstructions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordCnameRrPreconditionInstructions) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RecordCnameRrPreconditionInstructions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordCnameRrPreconditionInstructions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordCnameRrPreconditionInstructions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RecordCnameRrPreconditionInstructions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRdata

`func (o *RecordCnameRrPreconditionInstructions) GetRdata() string`

GetRdata returns the Rdata field if non-nil, zero value otherwise.

### GetRdataOk

`func (o *RecordCnameRrPreconditionInstructions) GetRdataOk() (*string, bool)`

GetRdataOk returns a tuple with the Rdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdata

`func (o *RecordCnameRrPreconditionInstructions) SetRdata(v string)`

SetRdata sets Rdata field to given value.

### HasRdata

`func (o *RecordCnameRrPreconditionInstructions) HasRdata() bool`

HasRdata returns a boolean if a field has been set.

### GetAction

`func (o *RecordCnameRrPreconditionInstructions) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RecordCnameRrPreconditionInstructions) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RecordCnameRrPreconditionInstructions) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RecordCnameRrPreconditionInstructions) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


