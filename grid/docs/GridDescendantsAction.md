# GridDescendantsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionWithEa** | Pointer to **string** | This option describes which action must be taken if the extensible attribute exists for both the parent and descendant objects: * INHERIT: inherit the extensible attribute from the parent object. * RETAIN: retain the value of an extensible attribute that was set for the child object. * CONVERT: the value of the extensible attribute must be copied from the parent object. | [optional] 
**OptionWithoutEa** | Pointer to **string** | This option describes which action must be taken if the extensible attribute exists for the parent, but is absent from the descendant object: * INHERIT: inherit the extensible attribute from the parent object. * NOT_INHERIT: do nothing. | [optional] 
**OptionDeleteEa** | Pointer to **string** | This option describes which action must be taken if the extensible attribute exists for the descendant, but is absent for the parent object: * RETAIN: retain the extensible attribute value for the descendant object. * REMOVE: remove this extensible attribute from the descendant object. | [optional] 

## Methods

### NewGridDescendantsAction

`func NewGridDescendantsAction() *GridDescendantsAction`

NewGridDescendantsAction instantiates a new GridDescendantsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDescendantsActionWithDefaults

`func NewGridDescendantsActionWithDefaults() *GridDescendantsAction`

NewGridDescendantsActionWithDefaults instantiates a new GridDescendantsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionWithEa

`func (o *GridDescendantsAction) GetOptionWithEa() string`

GetOptionWithEa returns the OptionWithEa field if non-nil, zero value otherwise.

### GetOptionWithEaOk

`func (o *GridDescendantsAction) GetOptionWithEaOk() (*string, bool)`

GetOptionWithEaOk returns a tuple with the OptionWithEa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionWithEa

`func (o *GridDescendantsAction) SetOptionWithEa(v string)`

SetOptionWithEa sets OptionWithEa field to given value.

### HasOptionWithEa

`func (o *GridDescendantsAction) HasOptionWithEa() bool`

HasOptionWithEa returns a boolean if a field has been set.

### GetOptionWithoutEa

`func (o *GridDescendantsAction) GetOptionWithoutEa() string`

GetOptionWithoutEa returns the OptionWithoutEa field if non-nil, zero value otherwise.

### GetOptionWithoutEaOk

`func (o *GridDescendantsAction) GetOptionWithoutEaOk() (*string, bool)`

GetOptionWithoutEaOk returns a tuple with the OptionWithoutEa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionWithoutEa

`func (o *GridDescendantsAction) SetOptionWithoutEa(v string)`

SetOptionWithoutEa sets OptionWithoutEa field to given value.

### HasOptionWithoutEa

`func (o *GridDescendantsAction) HasOptionWithoutEa() bool`

HasOptionWithoutEa returns a boolean if a field has been set.

### GetOptionDeleteEa

`func (o *GridDescendantsAction) GetOptionDeleteEa() string`

GetOptionDeleteEa returns the OptionDeleteEa field if non-nil, zero value otherwise.

### GetOptionDeleteEaOk

`func (o *GridDescendantsAction) GetOptionDeleteEaOk() (*string, bool)`

GetOptionDeleteEaOk returns a tuple with the OptionDeleteEa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDeleteEa

`func (o *GridDescendantsAction) SetOptionDeleteEa(v string)`

SetOptionDeleteEa sets OptionDeleteEa field to given value.

### HasOptionDeleteEa

`func (o *GridDescendantsAction) HasOptionDeleteEa() bool`

HasOptionDeleteEa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


