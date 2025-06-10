# Extensibleattributedef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllowedObjectTypes** | Pointer to **[]string** | The object types this extensible attribute is allowed to associate with. | [optional] 
**Comment** | Pointer to **string** | Comment for the Extensible Attribute Definition; maximum 256 characters. | [optional] 
**DefaultValue** | Pointer to **string** | Default value used to pre-populate the attribute value in the GUI. For email, URL, and string types, the value is a string with a maximum of 256 characters. For an integer, the value is an integer from -2147483648 through 2147483647. For a date, the value is the number of seconds that have elapsed since January 1st, 1970 UTC. | [optional] 
**DescendantsAction** | Pointer to [**ExtensibleattributedefDescendantsAction**](ExtensibleattributedefDescendantsAction.md) |  | [optional] 
**Flags** | Pointer to **string** | This field contains extensible attribute flags. Possible values: (A)udited, (C)loud API, Cloud (G)master, (I)nheritable, (L)isted, (M)andatory value, MGM (P)rivate, (R)ead Only, (S)ort enum values, Multiple (V)alues If there are two or more flags in the field, you must list them according to the order they are listed above. For example, &#39;CR&#39; is a valid value for the &#39;flags&#39; field because C &#x3D; Cloud API is listed before R &#x3D; Read only. However, the value &#39;RC&#39; is invalid because the order for the &#39;flags&#39; field is broken. | [optional] 
**ListValues** | Pointer to [**[]ExtensibleattributedefListValues**](ExtensibleattributedefListValues.md) | List of Values. Applicable if the extensible attribute type is ENUM. | [optional] 
**Max** | Pointer to **int64** | Maximum allowed value of extensible attribute. Applicable if the extensible attribute type is INTEGER. | [optional] 
**Min** | Pointer to **int64** | Minimum allowed value of extensible attribute. Applicable if the extensible attribute type is INTEGER. | [optional] 
**Name** | Pointer to **string** | The name of the Extensible Attribute Definition. | [optional] 
**Namespace** | Pointer to **string** | Namespace for the Extensible Attribute Definition. | [optional] [readonly] 
**Type** | Pointer to **string** | Type for the Extensible Attribute Definition. | [optional] 

## Methods

### NewExtensibleattributedef

`func NewExtensibleattributedef() *Extensibleattributedef`

NewExtensibleattributedef instantiates a new Extensibleattributedef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensibleattributedefWithDefaults

`func NewExtensibleattributedefWithDefaults() *Extensibleattributedef`

NewExtensibleattributedefWithDefaults instantiates a new Extensibleattributedef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Extensibleattributedef) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Extensibleattributedef) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Extensibleattributedef) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Extensibleattributedef) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowedObjectTypes

`func (o *Extensibleattributedef) GetAllowedObjectTypes() []string`

GetAllowedObjectTypes returns the AllowedObjectTypes field if non-nil, zero value otherwise.

### GetAllowedObjectTypesOk

`func (o *Extensibleattributedef) GetAllowedObjectTypesOk() (*[]string, bool)`

GetAllowedObjectTypesOk returns a tuple with the AllowedObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedObjectTypes

`func (o *Extensibleattributedef) SetAllowedObjectTypes(v []string)`

SetAllowedObjectTypes sets AllowedObjectTypes field to given value.

### HasAllowedObjectTypes

`func (o *Extensibleattributedef) HasAllowedObjectTypes() bool`

HasAllowedObjectTypes returns a boolean if a field has been set.

### GetComment

`func (o *Extensibleattributedef) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Extensibleattributedef) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Extensibleattributedef) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Extensibleattributedef) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultValue

`func (o *Extensibleattributedef) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *Extensibleattributedef) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *Extensibleattributedef) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *Extensibleattributedef) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetDescendantsAction

`func (o *Extensibleattributedef) GetDescendantsAction() ExtensibleattributedefDescendantsAction`

GetDescendantsAction returns the DescendantsAction field if non-nil, zero value otherwise.

### GetDescendantsActionOk

`func (o *Extensibleattributedef) GetDescendantsActionOk() (*ExtensibleattributedefDescendantsAction, bool)`

GetDescendantsActionOk returns a tuple with the DescendantsAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescendantsAction

`func (o *Extensibleattributedef) SetDescendantsAction(v ExtensibleattributedefDescendantsAction)`

SetDescendantsAction sets DescendantsAction field to given value.

### HasDescendantsAction

`func (o *Extensibleattributedef) HasDescendantsAction() bool`

HasDescendantsAction returns a boolean if a field has been set.

### GetFlags

`func (o *Extensibleattributedef) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Extensibleattributedef) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Extensibleattributedef) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Extensibleattributedef) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetListValues

`func (o *Extensibleattributedef) GetListValues() []ExtensibleattributedefListValues`

GetListValues returns the ListValues field if non-nil, zero value otherwise.

### GetListValuesOk

`func (o *Extensibleattributedef) GetListValuesOk() (*[]ExtensibleattributedefListValues, bool)`

GetListValuesOk returns a tuple with the ListValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListValues

`func (o *Extensibleattributedef) SetListValues(v []ExtensibleattributedefListValues)`

SetListValues sets ListValues field to given value.

### HasListValues

`func (o *Extensibleattributedef) HasListValues() bool`

HasListValues returns a boolean if a field has been set.

### GetMax

`func (o *Extensibleattributedef) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *Extensibleattributedef) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *Extensibleattributedef) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *Extensibleattributedef) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *Extensibleattributedef) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *Extensibleattributedef) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *Extensibleattributedef) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *Extensibleattributedef) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetName

`func (o *Extensibleattributedef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Extensibleattributedef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Extensibleattributedef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Extensibleattributedef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *Extensibleattributedef) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Extensibleattributedef) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Extensibleattributedef) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Extensibleattributedef) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetType

`func (o *Extensibleattributedef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Extensibleattributedef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Extensibleattributedef) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Extensibleattributedef) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


