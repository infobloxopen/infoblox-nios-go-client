# GetRulesetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Descriptive comment about the Ruleset object. | [optional] 
**Disabled** | Pointer to **bool** | The flag that indicates if the Ruleset object is disabled. | [optional] 
**Name** | Pointer to **string** | The name of this Ruleset object. | [optional] 
**NxdomainRules** | Pointer to [**[]RulesetNxdomainRules**](RulesetNxdomainRules.md) | The list of Rules assigned to this Ruleset object. Rules can be set only when the Ruleset type is set to \&quot;NXDOMAIN\&quot;. | [optional] 
**Type** | Pointer to **string** | The type of this Ruleset object. | [optional] 
**Result** | Pointer to [**Ruleset**](Ruleset.md) |  | [optional] 

## Methods

### NewGetRulesetResponse

`func NewGetRulesetResponse() *GetRulesetResponse`

NewGetRulesetResponse instantiates a new GetRulesetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRulesetResponseWithDefaults

`func NewGetRulesetResponseWithDefaults() *GetRulesetResponse`

NewGetRulesetResponseWithDefaults instantiates a new GetRulesetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRulesetResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRulesetResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRulesetResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRulesetResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetRulesetResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRulesetResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRulesetResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRulesetResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *GetRulesetResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetRulesetResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetRulesetResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetRulesetResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetName

`func (o *GetRulesetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRulesetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRulesetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRulesetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNxdomainRules

`func (o *GetRulesetResponse) GetNxdomainRules() []RulesetNxdomainRules`

GetNxdomainRules returns the NxdomainRules field if non-nil, zero value otherwise.

### GetNxdomainRulesOk

`func (o *GetRulesetResponse) GetNxdomainRulesOk() (*[]RulesetNxdomainRules, bool)`

GetNxdomainRulesOk returns a tuple with the NxdomainRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNxdomainRules

`func (o *GetRulesetResponse) SetNxdomainRules(v []RulesetNxdomainRules)`

SetNxdomainRules sets NxdomainRules field to given value.

### HasNxdomainRules

`func (o *GetRulesetResponse) HasNxdomainRules() bool`

HasNxdomainRules returns a boolean if a field has been set.

### GetType

`func (o *GetRulesetResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetRulesetResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetRulesetResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetRulesetResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *GetRulesetResponse) GetResult() Ruleset`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRulesetResponse) GetResultOk() (*Ruleset, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRulesetResponse) SetResult(v Ruleset)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRulesetResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


