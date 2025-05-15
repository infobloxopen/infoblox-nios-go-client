# ZoneRpFireeyeRuleMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AptOverride** | Pointer to **string** | The override setting for APT alerts. | [optional] 
**FireeyeAlertMapping** | Pointer to [**ZonerpfireeyerulemappingFireeyeAlertMapping**](ZonerpfireeyerulemappingFireeyeAlertMapping.md) |  | [optional] 
**SubstitutedDomainName** | Pointer to **string** | The domain name to be substituted, this is applicable only when apt_override is set to \&quot;SUBSTITUTE\&quot;. | [optional] 

## Methods

### NewZoneRpFireeyeRuleMapping

`func NewZoneRpFireeyeRuleMapping() *ZoneRpFireeyeRuleMapping`

NewZoneRpFireeyeRuleMapping instantiates a new ZoneRpFireeyeRuleMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneRpFireeyeRuleMappingWithDefaults

`func NewZoneRpFireeyeRuleMappingWithDefaults() *ZoneRpFireeyeRuleMapping`

NewZoneRpFireeyeRuleMappingWithDefaults instantiates a new ZoneRpFireeyeRuleMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAptOverride

`func (o *ZoneRpFireeyeRuleMapping) GetAptOverride() string`

GetAptOverride returns the AptOverride field if non-nil, zero value otherwise.

### GetAptOverrideOk

`func (o *ZoneRpFireeyeRuleMapping) GetAptOverrideOk() (*string, bool)`

GetAptOverrideOk returns a tuple with the AptOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAptOverride

`func (o *ZoneRpFireeyeRuleMapping) SetAptOverride(v string)`

SetAptOverride sets AptOverride field to given value.

### HasAptOverride

`func (o *ZoneRpFireeyeRuleMapping) HasAptOverride() bool`

HasAptOverride returns a boolean if a field has been set.

### GetFireeyeAlertMapping

`func (o *ZoneRpFireeyeRuleMapping) GetFireeyeAlertMapping() ZonerpfireeyerulemappingFireeyeAlertMapping`

GetFireeyeAlertMapping returns the FireeyeAlertMapping field if non-nil, zero value otherwise.

### GetFireeyeAlertMappingOk

`func (o *ZoneRpFireeyeRuleMapping) GetFireeyeAlertMappingOk() (*ZonerpfireeyerulemappingFireeyeAlertMapping, bool)`

GetFireeyeAlertMappingOk returns a tuple with the FireeyeAlertMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireeyeAlertMapping

`func (o *ZoneRpFireeyeRuleMapping) SetFireeyeAlertMapping(v ZonerpfireeyerulemappingFireeyeAlertMapping)`

SetFireeyeAlertMapping sets FireeyeAlertMapping field to given value.

### HasFireeyeAlertMapping

`func (o *ZoneRpFireeyeRuleMapping) HasFireeyeAlertMapping() bool`

HasFireeyeAlertMapping returns a boolean if a field has been set.

### GetSubstitutedDomainName

`func (o *ZoneRpFireeyeRuleMapping) GetSubstitutedDomainName() string`

GetSubstitutedDomainName returns the SubstitutedDomainName field if non-nil, zero value otherwise.

### GetSubstitutedDomainNameOk

`func (o *ZoneRpFireeyeRuleMapping) GetSubstitutedDomainNameOk() (*string, bool)`

GetSubstitutedDomainNameOk returns a tuple with the SubstitutedDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutedDomainName

`func (o *ZoneRpFireeyeRuleMapping) SetSubstitutedDomainName(v string)`

SetSubstitutedDomainName sets SubstitutedDomainName field to given value.

### HasSubstitutedDomainName

`func (o *ZoneRpFireeyeRuleMapping) HasSubstitutedDomainName() bool`

HasSubstitutedDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


