# RulesetNxdomainRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action to perform when a domain name matches the pattern defined in this Ruleset. | [optional] 
**Pattern** | Pointer to **string** | The pattern that is used to match the domain name. | [optional] 

## Methods

### NewRulesetNxdomainRules

`func NewRulesetNxdomainRules() *RulesetNxdomainRules`

NewRulesetNxdomainRules instantiates a new RulesetNxdomainRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulesetNxdomainRulesWithDefaults

`func NewRulesetNxdomainRulesWithDefaults() *RulesetNxdomainRules`

NewRulesetNxdomainRulesWithDefaults instantiates a new RulesetNxdomainRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RulesetNxdomainRules) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RulesetNxdomainRules) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RulesetNxdomainRules) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RulesetNxdomainRules) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPattern

`func (o *RulesetNxdomainRules) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *RulesetNxdomainRules) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *RulesetNxdomainRules) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *RulesetNxdomainRules) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


