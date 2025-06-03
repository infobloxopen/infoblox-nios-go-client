# RangeMacFilterRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to **string** | The name of the DHCP filter. | [optional] 
**Permission** | Pointer to **string** | The permission to be applied. | [optional] 

## Methods

### NewRangeMacFilterRules

`func NewRangeMacFilterRules() *RangeMacFilterRules`

NewRangeMacFilterRules instantiates a new RangeMacFilterRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeMacFilterRulesWithDefaults

`func NewRangeMacFilterRulesWithDefaults() *RangeMacFilterRules`

NewRangeMacFilterRulesWithDefaults instantiates a new RangeMacFilterRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RangeMacFilterRules) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RangeMacFilterRules) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RangeMacFilterRules) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RangeMacFilterRules) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPermission

`func (o *RangeMacFilterRules) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *RangeMacFilterRules) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *RangeMacFilterRules) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *RangeMacFilterRules) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


