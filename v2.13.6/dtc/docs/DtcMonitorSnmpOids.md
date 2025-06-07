# DtcMonitorSnmpOids

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oid** | Pointer to **string** | The SNMP OID value for DTC SNMP Monitor health checks. | [optional] 
**Comment** | Pointer to **string** | The comment for a DTC SNMP Health Monitor OID object. | [optional] 
**Type** | Pointer to **string** | The value of the condition type for DTC SNMP Monitor health check results. | [optional] 
**Condition** | Pointer to **string** | The condition of the validation result for an SNMP health check. The following conditions can be applied to the health check results: &#39;ANY&#39; accepts any response; &#39;EXACT&#39; accepts result equal to &#39;first&#39;; &#39;LEQ&#39; accepts result which is less than &#39;first&#39;; &#39;GEQ&#39; accepts result which is greater than &#39;first&#39;; &#39;RANGE&#39; accepts result value of which is between &#39;first&#39; and &#39;last&#39;. | [optional] 
**First** | Pointer to **string** | The condition&#39;s first term to match against the SNMP health check result. | [optional] 
**Last** | Pointer to **string** | The condition&#39;s second term to match against the SNMP health check result with &#39;RANGE&#39; condition. | [optional] 

## Methods

### NewDtcMonitorSnmpOids

`func NewDtcMonitorSnmpOids() *DtcMonitorSnmpOids`

NewDtcMonitorSnmpOids instantiates a new DtcMonitorSnmpOids object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcMonitorSnmpOidsWithDefaults

`func NewDtcMonitorSnmpOidsWithDefaults() *DtcMonitorSnmpOids`

NewDtcMonitorSnmpOidsWithDefaults instantiates a new DtcMonitorSnmpOids object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOid

`func (o *DtcMonitorSnmpOids) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *DtcMonitorSnmpOids) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *DtcMonitorSnmpOids) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *DtcMonitorSnmpOids) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetComment

`func (o *DtcMonitorSnmpOids) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcMonitorSnmpOids) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcMonitorSnmpOids) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcMonitorSnmpOids) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetType

`func (o *DtcMonitorSnmpOids) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DtcMonitorSnmpOids) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DtcMonitorSnmpOids) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DtcMonitorSnmpOids) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCondition

`func (o *DtcMonitorSnmpOids) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *DtcMonitorSnmpOids) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *DtcMonitorSnmpOids) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *DtcMonitorSnmpOids) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetFirst

`func (o *DtcMonitorSnmpOids) GetFirst() string`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *DtcMonitorSnmpOids) GetFirstOk() (*string, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *DtcMonitorSnmpOids) SetFirst(v string)`

SetFirst sets First field to given value.

### HasFirst

`func (o *DtcMonitorSnmpOids) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetLast

`func (o *DtcMonitorSnmpOids) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *DtcMonitorSnmpOids) GetLastOk() (*string, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *DtcMonitorSnmpOids) SetLast(v string)`

SetLast sets Last field to given value.

### HasLast

`func (o *DtcMonitorSnmpOids) HasLast() bool`

HasLast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


