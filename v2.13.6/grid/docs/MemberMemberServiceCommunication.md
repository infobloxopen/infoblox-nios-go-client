# MemberMemberServiceCommunication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to **string** | The service for a Grid member. | [optional] 
**Type** | Pointer to **string** | Communication type. | [optional] 
**Option** | Pointer to **string** | The option for communication type. | [optional] [readonly] 

## Methods

### NewMemberMemberServiceCommunication

`func NewMemberMemberServiceCommunication() *MemberMemberServiceCommunication`

NewMemberMemberServiceCommunication instantiates a new MemberMemberServiceCommunication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberMemberServiceCommunicationWithDefaults

`func NewMemberMemberServiceCommunicationWithDefaults() *MemberMemberServiceCommunication`

NewMemberMemberServiceCommunicationWithDefaults instantiates a new MemberMemberServiceCommunication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *MemberMemberServiceCommunication) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MemberMemberServiceCommunication) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MemberMemberServiceCommunication) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MemberMemberServiceCommunication) HasService() bool`

HasService returns a boolean if a field has been set.

### GetType

`func (o *MemberMemberServiceCommunication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemberMemberServiceCommunication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemberMemberServiceCommunication) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MemberMemberServiceCommunication) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOption

`func (o *MemberMemberServiceCommunication) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *MemberMemberServiceCommunication) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *MemberMemberServiceCommunication) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *MemberMemberServiceCommunication) HasOption() bool`

HasOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


