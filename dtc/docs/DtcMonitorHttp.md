# DtcMonitorHttp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Ciphers** | Pointer to **string** | An optional cipher list for a secure HTTP/S connection. | [optional] 
**ClientCert** | Pointer to **string** | An optional client certificate, supplied in a secure HTTP/S mode if present. | [optional] 
**Comment** | Pointer to **string** | Comment for this DTC monitor; maximum 256 characters. | [optional] 
**ContentCheck** | Pointer to **string** | The content check type. | [optional] 
**ContentCheckInput** | Pointer to **string** | A portion of response to use as input for content check. | [optional] 
**ContentCheckOp** | Pointer to **string** | A content check success criteria operator. | [optional] 
**ContentCheckRegex** | Pointer to **string** | A content check regular expression. | [optional] 
**ContentExtractGroup** | Pointer to **int64** | A content extraction sub-expression to extract. | [optional] 
**ContentExtractType** | Pointer to **string** | A content extraction expected type for the extracted data. | [optional] 
**ContentExtractValue** | Pointer to **string** | A content extraction value to compare with extracted result. | [optional] 
**EnableSni** | Pointer to **bool** | Determines whether the Server Name Indication (SNI) for HTTPS monitor is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interval** | Pointer to **int64** | The interval for TCP health check. | [optional] 
**Name** | Pointer to **string** | The display name for this DTC monitor. | [optional] 
**Port** | Pointer to **int64** | Port for TCP requests. | [optional] 
**Request** | Pointer to **string** | An HTTP request to send. | [optional] 
**Result** | Pointer to **string** | The type of an expected result. | [optional] 
**ResultCode** | Pointer to **int64** | The expected return code. | [optional] 
**RetryDown** | Pointer to **int64** | The value of how many times the server should appear as down to be treated as dead after it was alive. | [optional] 
**RetryUp** | Pointer to **int64** | The value of how many times the server should appear as up to be treated as alive after it was dead. | [optional] 
**Secure** | Pointer to **bool** | The connection security status. | [optional] 
**Timeout** | Pointer to **int64** | The timeout for TCP health check in seconds. | [optional] 
**ValidateCert** | Pointer to **bool** | Determines whether the validation of the remote server&#39;s certificate is enabled. | [optional] 

## Methods

### NewDtcMonitorHttp

`func NewDtcMonitorHttp() *DtcMonitorHttp`

NewDtcMonitorHttp instantiates a new DtcMonitorHttp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcMonitorHttpWithDefaults

`func NewDtcMonitorHttpWithDefaults() *DtcMonitorHttp`

NewDtcMonitorHttpWithDefaults instantiates a new DtcMonitorHttp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcMonitorHttp) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcMonitorHttp) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcMonitorHttp) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcMonitorHttp) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCiphers

`func (o *DtcMonitorHttp) GetCiphers() string`

GetCiphers returns the Ciphers field if non-nil, zero value otherwise.

### GetCiphersOk

`func (o *DtcMonitorHttp) GetCiphersOk() (*string, bool)`

GetCiphersOk returns a tuple with the Ciphers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiphers

`func (o *DtcMonitorHttp) SetCiphers(v string)`

SetCiphers sets Ciphers field to given value.

### HasCiphers

`func (o *DtcMonitorHttp) HasCiphers() bool`

HasCiphers returns a boolean if a field has been set.

### GetClientCert

`func (o *DtcMonitorHttp) GetClientCert() string`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *DtcMonitorHttp) GetClientCertOk() (*string, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *DtcMonitorHttp) SetClientCert(v string)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *DtcMonitorHttp) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetComment

`func (o *DtcMonitorHttp) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcMonitorHttp) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcMonitorHttp) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcMonitorHttp) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContentCheck

`func (o *DtcMonitorHttp) GetContentCheck() string`

GetContentCheck returns the ContentCheck field if non-nil, zero value otherwise.

### GetContentCheckOk

`func (o *DtcMonitorHttp) GetContentCheckOk() (*string, bool)`

GetContentCheckOk returns a tuple with the ContentCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCheck

`func (o *DtcMonitorHttp) SetContentCheck(v string)`

SetContentCheck sets ContentCheck field to given value.

### HasContentCheck

`func (o *DtcMonitorHttp) HasContentCheck() bool`

HasContentCheck returns a boolean if a field has been set.

### GetContentCheckInput

`func (o *DtcMonitorHttp) GetContentCheckInput() string`

GetContentCheckInput returns the ContentCheckInput field if non-nil, zero value otherwise.

### GetContentCheckInputOk

`func (o *DtcMonitorHttp) GetContentCheckInputOk() (*string, bool)`

GetContentCheckInputOk returns a tuple with the ContentCheckInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCheckInput

`func (o *DtcMonitorHttp) SetContentCheckInput(v string)`

SetContentCheckInput sets ContentCheckInput field to given value.

### HasContentCheckInput

`func (o *DtcMonitorHttp) HasContentCheckInput() bool`

HasContentCheckInput returns a boolean if a field has been set.

### GetContentCheckOp

`func (o *DtcMonitorHttp) GetContentCheckOp() string`

GetContentCheckOp returns the ContentCheckOp field if non-nil, zero value otherwise.

### GetContentCheckOpOk

`func (o *DtcMonitorHttp) GetContentCheckOpOk() (*string, bool)`

GetContentCheckOpOk returns a tuple with the ContentCheckOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCheckOp

`func (o *DtcMonitorHttp) SetContentCheckOp(v string)`

SetContentCheckOp sets ContentCheckOp field to given value.

### HasContentCheckOp

`func (o *DtcMonitorHttp) HasContentCheckOp() bool`

HasContentCheckOp returns a boolean if a field has been set.

### GetContentCheckRegex

`func (o *DtcMonitorHttp) GetContentCheckRegex() string`

GetContentCheckRegex returns the ContentCheckRegex field if non-nil, zero value otherwise.

### GetContentCheckRegexOk

`func (o *DtcMonitorHttp) GetContentCheckRegexOk() (*string, bool)`

GetContentCheckRegexOk returns a tuple with the ContentCheckRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentCheckRegex

`func (o *DtcMonitorHttp) SetContentCheckRegex(v string)`

SetContentCheckRegex sets ContentCheckRegex field to given value.

### HasContentCheckRegex

`func (o *DtcMonitorHttp) HasContentCheckRegex() bool`

HasContentCheckRegex returns a boolean if a field has been set.

### GetContentExtractGroup

`func (o *DtcMonitorHttp) GetContentExtractGroup() int64`

GetContentExtractGroup returns the ContentExtractGroup field if non-nil, zero value otherwise.

### GetContentExtractGroupOk

`func (o *DtcMonitorHttp) GetContentExtractGroupOk() (*int64, bool)`

GetContentExtractGroupOk returns a tuple with the ContentExtractGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentExtractGroup

`func (o *DtcMonitorHttp) SetContentExtractGroup(v int64)`

SetContentExtractGroup sets ContentExtractGroup field to given value.

### HasContentExtractGroup

`func (o *DtcMonitorHttp) HasContentExtractGroup() bool`

HasContentExtractGroup returns a boolean if a field has been set.

### GetContentExtractType

`func (o *DtcMonitorHttp) GetContentExtractType() string`

GetContentExtractType returns the ContentExtractType field if non-nil, zero value otherwise.

### GetContentExtractTypeOk

`func (o *DtcMonitorHttp) GetContentExtractTypeOk() (*string, bool)`

GetContentExtractTypeOk returns a tuple with the ContentExtractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentExtractType

`func (o *DtcMonitorHttp) SetContentExtractType(v string)`

SetContentExtractType sets ContentExtractType field to given value.

### HasContentExtractType

`func (o *DtcMonitorHttp) HasContentExtractType() bool`

HasContentExtractType returns a boolean if a field has been set.

### GetContentExtractValue

`func (o *DtcMonitorHttp) GetContentExtractValue() string`

GetContentExtractValue returns the ContentExtractValue field if non-nil, zero value otherwise.

### GetContentExtractValueOk

`func (o *DtcMonitorHttp) GetContentExtractValueOk() (*string, bool)`

GetContentExtractValueOk returns a tuple with the ContentExtractValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentExtractValue

`func (o *DtcMonitorHttp) SetContentExtractValue(v string)`

SetContentExtractValue sets ContentExtractValue field to given value.

### HasContentExtractValue

`func (o *DtcMonitorHttp) HasContentExtractValue() bool`

HasContentExtractValue returns a boolean if a field has been set.

### GetEnableSni

`func (o *DtcMonitorHttp) GetEnableSni() bool`

GetEnableSni returns the EnableSni field if non-nil, zero value otherwise.

### GetEnableSniOk

`func (o *DtcMonitorHttp) GetEnableSniOk() (*bool, bool)`

GetEnableSniOk returns a tuple with the EnableSni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSni

`func (o *DtcMonitorHttp) SetEnableSni(v bool)`

SetEnableSni sets EnableSni field to given value.

### HasEnableSni

`func (o *DtcMonitorHttp) HasEnableSni() bool`

HasEnableSni returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *DtcMonitorHttp) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *DtcMonitorHttp) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *DtcMonitorHttp) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *DtcMonitorHttp) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *DtcMonitorHttp) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *DtcMonitorHttp) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *DtcMonitorHttp) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *DtcMonitorHttp) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *DtcMonitorHttp) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *DtcMonitorHttp) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *DtcMonitorHttp) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *DtcMonitorHttp) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetInterval

`func (o *DtcMonitorHttp) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *DtcMonitorHttp) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *DtcMonitorHttp) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *DtcMonitorHttp) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *DtcMonitorHttp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcMonitorHttp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcMonitorHttp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcMonitorHttp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *DtcMonitorHttp) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DtcMonitorHttp) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DtcMonitorHttp) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DtcMonitorHttp) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRequest

`func (o *DtcMonitorHttp) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *DtcMonitorHttp) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *DtcMonitorHttp) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *DtcMonitorHttp) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResult

`func (o *DtcMonitorHttp) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DtcMonitorHttp) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DtcMonitorHttp) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *DtcMonitorHttp) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetResultCode

`func (o *DtcMonitorHttp) GetResultCode() int64`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *DtcMonitorHttp) GetResultCodeOk() (*int64, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *DtcMonitorHttp) SetResultCode(v int64)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *DtcMonitorHttp) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.

### GetRetryDown

`func (o *DtcMonitorHttp) GetRetryDown() int64`

GetRetryDown returns the RetryDown field if non-nil, zero value otherwise.

### GetRetryDownOk

`func (o *DtcMonitorHttp) GetRetryDownOk() (*int64, bool)`

GetRetryDownOk returns a tuple with the RetryDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDown

`func (o *DtcMonitorHttp) SetRetryDown(v int64)`

SetRetryDown sets RetryDown field to given value.

### HasRetryDown

`func (o *DtcMonitorHttp) HasRetryDown() bool`

HasRetryDown returns a boolean if a field has been set.

### GetRetryUp

`func (o *DtcMonitorHttp) GetRetryUp() int64`

GetRetryUp returns the RetryUp field if non-nil, zero value otherwise.

### GetRetryUpOk

`func (o *DtcMonitorHttp) GetRetryUpOk() (*int64, bool)`

GetRetryUpOk returns a tuple with the RetryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryUp

`func (o *DtcMonitorHttp) SetRetryUp(v int64)`

SetRetryUp sets RetryUp field to given value.

### HasRetryUp

`func (o *DtcMonitorHttp) HasRetryUp() bool`

HasRetryUp returns a boolean if a field has been set.

### GetSecure

`func (o *DtcMonitorHttp) GetSecure() bool`

GetSecure returns the Secure field if non-nil, zero value otherwise.

### GetSecureOk

`func (o *DtcMonitorHttp) GetSecureOk() (*bool, bool)`

GetSecureOk returns a tuple with the Secure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecure

`func (o *DtcMonitorHttp) SetSecure(v bool)`

SetSecure sets Secure field to given value.

### HasSecure

`func (o *DtcMonitorHttp) HasSecure() bool`

HasSecure returns a boolean if a field has been set.

### GetTimeout

`func (o *DtcMonitorHttp) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DtcMonitorHttp) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DtcMonitorHttp) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DtcMonitorHttp) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetValidateCert

`func (o *DtcMonitorHttp) GetValidateCert() bool`

GetValidateCert returns the ValidateCert field if non-nil, zero value otherwise.

### GetValidateCertOk

`func (o *DtcMonitorHttp) GetValidateCertOk() (*bool, bool)`

GetValidateCertOk returns a tuple with the ValidateCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCert

`func (o *DtcMonitorHttp) SetValidateCert(v bool)`

SetValidateCert sets ValidateCert field to given value.

### HasValidateCert

`func (o *DtcMonitorHttp) HasValidateCert() bool`

HasValidateCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


