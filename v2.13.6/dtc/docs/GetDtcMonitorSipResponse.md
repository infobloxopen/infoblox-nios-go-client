# GetDtcMonitorSipResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Ciphers** | Pointer to **string** | An optional cipher list for secure TLS/SIPS connection. | [optional] 
**ClientCert** | Pointer to **string** | An optional client certificate, supplied in TLS and SIPS mode if present. | [optional] 
**Comment** | Pointer to **string** | Comment for this DTC monitor; maximum 256 characters. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interval** | Pointer to **int64** | The interval for TCP health check. | [optional] 
**Name** | Pointer to **string** | The display name for this DTC monitor. | [optional] 
**Port** | Pointer to **int64** | The port value for SIP requests. | [optional] 
**Request** | Pointer to **string** | A SIP request to send | [optional] 
**Result** | Pointer to [**DtcMonitorSip**](DtcMonitorSip.md) |  | [optional] 
**ResultCode** | Pointer to **int64** | The expected return code value. | [optional] 
**RetryDown** | Pointer to **int64** | The value of how many times the server should appear as down to be treated as dead after it was alive. | [optional] 
**RetryUp** | Pointer to **int64** | The value of how many times the server should appear as up to be treated as alive after it was dead. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for TCP health check in seconds. | [optional] 
**Transport** | Pointer to **string** | The transport layer protocol to use for SIP check. | [optional] 
**ValidateCert** | Pointer to **bool** | Determines whether the validation of the remote server&#39;s certificate is enabled. | [optional] 

## Methods

### NewGetDtcMonitorSipResponse

`func NewGetDtcMonitorSipResponse() *GetDtcMonitorSipResponse`

NewGetDtcMonitorSipResponse instantiates a new GetDtcMonitorSipResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcMonitorSipResponseWithDefaults

`func NewGetDtcMonitorSipResponseWithDefaults() *GetDtcMonitorSipResponse`

NewGetDtcMonitorSipResponseWithDefaults instantiates a new GetDtcMonitorSipResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcMonitorSipResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcMonitorSipResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcMonitorSipResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcMonitorSipResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCiphers

`func (o *GetDtcMonitorSipResponse) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *GetDtcMonitorSipResponse) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *GetDtcMonitorSipResponse) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *GetDtcMonitorSipResponse) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetClientCert

`func (o *GetDtcMonitorSipResponse) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *GetDtcMonitorSipResponse) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *GetDtcMonitorSipResponse) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *GetDtcMonitorSipResponse) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcMonitorSipResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcMonitorSipResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcMonitorSipResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcMonitorSipResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetDtcMonitorSipResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetDtcMonitorSipResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetDtcMonitorSipResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetDtcMonitorSipResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetInterval

`func (o *GetDtcMonitorSipResponse) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetDtcMonitorSipResponse) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetDtcMonitorSipResponse) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GetDtcMonitorSipResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *GetDtcMonitorSipResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDtcMonitorSipResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDtcMonitorSipResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDtcMonitorSipResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *GetDtcMonitorSipResponse) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetDtcMonitorSipResponse) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetDtcMonitorSipResponse) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetDtcMonitorSipResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRequest

`func (o *GetDtcMonitorSipResponse) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *GetDtcMonitorSipResponse) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *GetDtcMonitorSipResponse) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *GetDtcMonitorSipResponse) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcMonitorSipResponse) GetResult() DtcMonitorSip`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcMonitorSipResponse) GetResultOk() (*DtcMonitorSip, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcMonitorSipResponse) SetResult(v DtcMonitorSip)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcMonitorSipResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResultCode

`func (o *GetDtcMonitorSipResponse) GetResultCode() int64`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *GetDtcMonitorSipResponse) GetResultCodeOk() (*int64, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *GetDtcMonitorSipResponse) SetResultCode(v int64)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *GetDtcMonitorSipResponse) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetRetryDown

`func (o *GetDtcMonitorSipResponse) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *GetDtcMonitorSipResponse) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *GetDtcMonitorSipResponse) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *GetDtcMonitorSipResponse) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *GetDtcMonitorSipResponse) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *GetDtcMonitorSipResponse) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *GetDtcMonitorSipResponse) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *GetDtcMonitorSipResponse) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetTimeout

`func (o *GetDtcMonitorSipResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetDtcMonitorSipResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetDtcMonitorSipResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetDtcMonitorSipResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTransport

`func (o *GetDtcMonitorSipResponse) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *GetDtcMonitorSipResponse) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *GetDtcMonitorSipResponse) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *GetDtcMonitorSipResponse) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetValidateCert

`func (o *GetDtcMonitorSipResponse) GetValidateCert() bool`

GetValidateCert returns the ValidateCert field if non-nil, zero value otherwise.

### GetValidateCertOk

`func (o *GetDtcMonitorSipResponse) GetValidateCertOk() (*bool, bool)`

GetValidateCertOk returns a tuple with the ValidateCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCert

`func (o *GetDtcMonitorSipResponse) SetValidateCert(v bool)`

SetValidateCert sets ValidateCert field to given value.

### HasValidateCert

`func (o *GetDtcMonitorSipResponse) HasValidateCert() bool`

HasValidateCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


