# ViewResponseRateLimiting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableRrl** | Pointer to **bool** | Determines if the response rate limiting is enabled or not. | [optional] 
**LogOnly** | Pointer to **bool** | Determines if logging for response rate limiting without dropping any requests is enabled or not. | [optional] 
**ResponsesPerSecond** | Pointer to **int64** | The number of responses per client per second. | [optional] 
**Window** | Pointer to **int64** | The time interval in seconds over which responses are tracked. | [optional] 
**Slip** | Pointer to **int64** | The response rate limiting slip. Note that if slip is not equal to 0 every n-th rate-limited UDP request is sent a truncated response instead of being dropped. | [optional] 

## Methods

### NewViewResponseRateLimiting

`func NewViewResponseRateLimiting() *ViewResponseRateLimiting`

NewViewResponseRateLimiting instantiates a new ViewResponseRateLimiting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewResponseRateLimitingWithDefaults

`func NewViewResponseRateLimitingWithDefaults() *ViewResponseRateLimiting`

NewViewResponseRateLimitingWithDefaults instantiates a new ViewResponseRateLimiting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableRrl

`func (o *ViewResponseRateLimiting) GetEnableRrl() bool`

GetEnableRrl returns the EnableRrl field if non-nil, zero value otherwise.

### GetEnableRrlOk

`func (o *ViewResponseRateLimiting) GetEnableRrlOk() (*bool, bool)`

GetEnableRrlOk returns a tuple with the EnableRrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRrl

`func (o *ViewResponseRateLimiting) SetEnableRrl(v bool)`

SetEnableRrl sets EnableRrl field to given value.

### HasEnableRrl

`func (o *ViewResponseRateLimiting) HasEnableRrl() bool`

HasEnableRrl returns a boolean if a field has been set.

### GetLogOnly

`func (o *ViewResponseRateLimiting) GetLogOnly() bool`

GetLogOnly returns the LogOnly field if non-nil, zero value otherwise.

### GetLogOnlyOk

`func (o *ViewResponseRateLimiting) GetLogOnlyOk() (*bool, bool)`

GetLogOnlyOk returns a tuple with the LogOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOnly

`func (o *ViewResponseRateLimiting) SetLogOnly(v bool)`

SetLogOnly sets LogOnly field to given value.

### HasLogOnly

`func (o *ViewResponseRateLimiting) HasLogOnly() bool`

HasLogOnly returns a boolean if a field has been set.

### GetResponsesPerSecond

`func (o *ViewResponseRateLimiting) GetResponsesPerSecond() int64`

GetResponsesPerSecond returns the ResponsesPerSecond field if non-nil, zero value otherwise.

### GetResponsesPerSecondOk

`func (o *ViewResponseRateLimiting) GetResponsesPerSecondOk() (*int64, bool)`

GetResponsesPerSecondOk returns a tuple with the ResponsesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsesPerSecond

`func (o *ViewResponseRateLimiting) SetResponsesPerSecond(v int64)`

SetResponsesPerSecond sets ResponsesPerSecond field to given value.

### HasResponsesPerSecond

`func (o *ViewResponseRateLimiting) HasResponsesPerSecond() bool`

HasResponsesPerSecond returns a boolean if a field has been set.

### GetWindow

`func (o *ViewResponseRateLimiting) GetWindow() int64`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *ViewResponseRateLimiting) GetWindowOk() (*int64, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *ViewResponseRateLimiting) SetWindow(v int64)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *ViewResponseRateLimiting) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetSlip

`func (o *ViewResponseRateLimiting) GetSlip() int64`

GetSlip returns the Slip field if non-nil, zero value otherwise.

### GetSlipOk

`func (o *ViewResponseRateLimiting) GetSlipOk() (*int64, bool)`

GetSlipOk returns a tuple with the Slip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlip

`func (o *ViewResponseRateLimiting) SetSlip(v int64)`

SetSlip sets Slip field to given value.

### HasSlip

`func (o *ViewResponseRateLimiting) HasSlip() bool`

HasSlip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


