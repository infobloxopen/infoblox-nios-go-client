# GetRirResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**CommunicationMode** | Pointer to **string** | The communication mode for RIR. | [optional] 
**Email** | Pointer to **string** | The e-mail address for RIR. | [optional] 
**Name** | Pointer to **string** | The name of RIR. | [optional] 
**Url** | Pointer to **string** | The WebAPI URL for RIR. | [optional] 
**UseEmail** | Pointer to **bool** | Use flag for: email | [optional] 
**UseUrl** | Pointer to **bool** | Use flag for: url | [optional] 
**Result** | Pointer to [**Rir**](Rir.md) |  | [optional] 

## Methods

### NewGetRirResponse

`func NewGetRirResponse() *GetRirResponse`

NewGetRirResponse instantiates a new GetRirResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRirResponseWithDefaults

`func NewGetRirResponseWithDefaults() *GetRirResponse`

NewGetRirResponseWithDefaults instantiates a new GetRirResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRirResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRirResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRirResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRirResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCommunicationMode

`func (o *GetRirResponse) GetCommunicationMode() string`

GetCommunicationMode returns the CommunicationMode field if non-nil, zero value otherwise.

### GetCommunicationModeOk

`func (o *GetRirResponse) GetCommunicationModeOk() (*string, bool)`

GetCommunicationModeOk returns a tuple with the CommunicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationMode

`func (o *GetRirResponse) SetCommunicationMode(v string)`

SetCommunicationMode sets CommunicationMode field to given value.

### HasCommunicationMode

`func (o *GetRirResponse) HasCommunicationMode() bool`

HasCommunicationMode returns a boolean if a field has been set.

### GetEmail

`func (o *GetRirResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetRirResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetRirResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetRirResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *GetRirResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRirResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRirResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRirResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *GetRirResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetRirResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetRirResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetRirResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUseEmail

`func (o *GetRirResponse) GetUseEmail() bool`

GetUseEmail returns the UseEmail field if non-nil, zero value otherwise.

### GetUseEmailOk

`func (o *GetRirResponse) GetUseEmailOk() (*bool, bool)`

GetUseEmailOk returns a tuple with the UseEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmail

`func (o *GetRirResponse) SetUseEmail(v bool)`

SetUseEmail sets UseEmail field to given value.

### HasUseEmail

`func (o *GetRirResponse) HasUseEmail() bool`

HasUseEmail returns a boolean if a field has been set.

### GetUseUrl

`func (o *GetRirResponse) GetUseUrl() bool`

GetUseUrl returns the UseUrl field if non-nil, zero value otherwise.

### GetUseUrlOk

`func (o *GetRirResponse) GetUseUrlOk() (*bool, bool)`

GetUseUrlOk returns a tuple with the UseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUrl

`func (o *GetRirResponse) SetUseUrl(v bool)`

SetUseUrl sets UseUrl field to given value.

### HasUseUrl

`func (o *GetRirResponse) HasUseUrl() bool`

HasUseUrl returns a boolean if a field has been set.

### GetResult

`func (o *GetRirResponse) GetResult() Rir`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRirResponse) GetResultOk() (*Rir, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRirResponse) SetResult(v Rir)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRirResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


