# ZoneAuthMsDcNsRecordCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IPv4 address of the domain controller that is allowed to create NS records. | [optional] 
**Comment** | Pointer to **string** | Optional user comment. | [optional] 

## Methods

### NewZoneAuthMsDcNsRecordCreation

`func NewZoneAuthMsDcNsRecordCreation() *ZoneAuthMsDcNsRecordCreation`

NewZoneAuthMsDcNsRecordCreation instantiates a new ZoneAuthMsDcNsRecordCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthMsDcNsRecordCreationWithDefaults

`func NewZoneAuthMsDcNsRecordCreationWithDefaults() *ZoneAuthMsDcNsRecordCreation`

NewZoneAuthMsDcNsRecordCreationWithDefaults instantiates a new ZoneAuthMsDcNsRecordCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ZoneAuthMsDcNsRecordCreation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ZoneAuthMsDcNsRecordCreation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ZoneAuthMsDcNsRecordCreation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ZoneAuthMsDcNsRecordCreation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetComment

`func (o *ZoneAuthMsDcNsRecordCreation) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ZoneAuthMsDcNsRecordCreation) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ZoneAuthMsDcNsRecordCreation) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ZoneAuthMsDcNsRecordCreation) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


