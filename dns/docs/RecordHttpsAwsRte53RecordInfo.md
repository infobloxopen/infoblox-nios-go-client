# RecordHttpsAwsRte53RecordInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasTargetDnsName** | Pointer to **string** | DNS name of the alias target. | [optional] [readonly] 
**AliasTargetHostedZoneId** | Pointer to **string** | Hosted zone ID of the alias target. | [optional] [readonly] 
**AliasTargetEvaluateTargetHealth** | Pointer to **bool** | Indicates if Amazon Route 53 evaluates the health of the alias target. | [optional] [readonly] 
**Failover** | Pointer to **string** | Indicates whether this is the primary or secondary resource record for Amazon Route 53 failover routing. | [optional] [readonly] 
**GeolocationContinentCode** | Pointer to **string** | Continent code for Amazon Route 53 geolocation routing. | [optional] [readonly] 
**GeolocationCountryCode** | Pointer to **string** | Country code for Amazon Route 53 geolocation routing. | [optional] [readonly] 
**GeolocationSubdivisionCode** | Pointer to **string** | Subdivision code for Amazon Route 53 geolocation routing. | [optional] [readonly] 
**HealthCheckId** | Pointer to **string** | ID of the health check that Amazon Route 53 performs for this resource record. | [optional] [readonly] 
**Region** | Pointer to **string** | Amazon EC2 region where this resource record resides for latency routing. | [optional] [readonly] 
**SetIdentifier** | Pointer to **string** | An identifier that differentiates records with the same DNS name and type for weighted, latency, geolocation, and failover routing. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of Amazon Route 53 resource record. | [optional] [readonly] 
**Weight** | Pointer to **int64** | Value that determines the portion of traffic for this record in weighted routing. The range is from 0 to 255. | [optional] [readonly] 

## Methods

### NewRecordHttpsAwsRte53RecordInfo

`func NewRecordHttpsAwsRte53RecordInfo() *RecordHttpsAwsRte53RecordInfo`

NewRecordHttpsAwsRte53RecordInfo instantiates a new RecordHttpsAwsRte53RecordInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordHttpsAwsRte53RecordInfoWithDefaults

`func NewRecordHttpsAwsRte53RecordInfoWithDefaults() *RecordHttpsAwsRte53RecordInfo`

NewRecordHttpsAwsRte53RecordInfoWithDefaults instantiates a new RecordHttpsAwsRte53RecordInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasTargetDnsName

`func (o *RecordHttpsAwsRte53RecordInfo) GetAliasTargetDnsName() string`

GetAliasTargetDnsName returns the AliasTargetDnsName field if non-nil, zero value otherwise.

### GetAliasTargetDnsNameOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetAliasTargetDnsNameOk() (*string, bool)`

GetAliasTargetDnsNameOk returns a tuple with the AliasTargetDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetDnsName

`func (o *RecordHttpsAwsRte53RecordInfo) SetAliasTargetDnsName(v string)`

SetAliasTargetDnsName sets AliasTargetDnsName field to given value.

### HasAliasTargetDnsName

`func (o *RecordHttpsAwsRte53RecordInfo) HasAliasTargetDnsName() bool`

HasAliasTargetDnsName returns a boolean if a field has been set.

### GetAliasTargetHostedZoneId

`func (o *RecordHttpsAwsRte53RecordInfo) GetAliasTargetHostedZoneId() string`

GetAliasTargetHostedZoneId returns the AliasTargetHostedZoneId field if non-nil, zero value otherwise.

### GetAliasTargetHostedZoneIdOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetAliasTargetHostedZoneIdOk() (*string, bool)`

GetAliasTargetHostedZoneIdOk returns a tuple with the AliasTargetHostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetHostedZoneId

`func (o *RecordHttpsAwsRte53RecordInfo) SetAliasTargetHostedZoneId(v string)`

SetAliasTargetHostedZoneId sets AliasTargetHostedZoneId field to given value.

### HasAliasTargetHostedZoneId

`func (o *RecordHttpsAwsRte53RecordInfo) HasAliasTargetHostedZoneId() bool`

HasAliasTargetHostedZoneId returns a boolean if a field has been set.

### GetAliasTargetEvaluateTargetHealth

`func (o *RecordHttpsAwsRte53RecordInfo) GetAliasTargetEvaluateTargetHealth() bool`

GetAliasTargetEvaluateTargetHealth returns the AliasTargetEvaluateTargetHealth field if non-nil, zero value otherwise.

### GetAliasTargetEvaluateTargetHealthOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetAliasTargetEvaluateTargetHealthOk() (*bool, bool)`

GetAliasTargetEvaluateTargetHealthOk returns a tuple with the AliasTargetEvaluateTargetHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetEvaluateTargetHealth

`func (o *RecordHttpsAwsRte53RecordInfo) SetAliasTargetEvaluateTargetHealth(v bool)`

SetAliasTargetEvaluateTargetHealth sets AliasTargetEvaluateTargetHealth field to given value.

### HasAliasTargetEvaluateTargetHealth

`func (o *RecordHttpsAwsRte53RecordInfo) HasAliasTargetEvaluateTargetHealth() bool`

HasAliasTargetEvaluateTargetHealth returns a boolean if a field has been set.

### GetFailover

`func (o *RecordHttpsAwsRte53RecordInfo) GetFailover() string`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetFailoverOk() (*string, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *RecordHttpsAwsRte53RecordInfo) SetFailover(v string)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *RecordHttpsAwsRte53RecordInfo) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetGeolocationContinentCode

`func (o *RecordHttpsAwsRte53RecordInfo) GetGeolocationContinentCode() string`

GetGeolocationContinentCode returns the GeolocationContinentCode field if non-nil, zero value otherwise.

### GetGeolocationContinentCodeOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetGeolocationContinentCodeOk() (*string, bool)`

GetGeolocationContinentCodeOk returns a tuple with the GeolocationContinentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocationContinentCode

`func (o *RecordHttpsAwsRte53RecordInfo) SetGeolocationContinentCode(v string)`

SetGeolocationContinentCode sets GeolocationContinentCode field to given value.

### HasGeolocationContinentCode

`func (o *RecordHttpsAwsRte53RecordInfo) HasGeolocationContinentCode() bool`

HasGeolocationContinentCode returns a boolean if a field has been set.

### GetGeolocationCountryCode

`func (o *RecordHttpsAwsRte53RecordInfo) GetGeolocationCountryCode() string`

GetGeolocationCountryCode returns the GeolocationCountryCode field if non-nil, zero value otherwise.

### GetGeolocationCountryCodeOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetGeolocationCountryCodeOk() (*string, bool)`

GetGeolocationCountryCodeOk returns a tuple with the GeolocationCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocationCountryCode

`func (o *RecordHttpsAwsRte53RecordInfo) SetGeolocationCountryCode(v string)`

SetGeolocationCountryCode sets GeolocationCountryCode field to given value.

### HasGeolocationCountryCode

`func (o *RecordHttpsAwsRte53RecordInfo) HasGeolocationCountryCode() bool`

HasGeolocationCountryCode returns a boolean if a field has been set.

### GetGeolocationSubdivisionCode

`func (o *RecordHttpsAwsRte53RecordInfo) GetGeolocationSubdivisionCode() string`

GetGeolocationSubdivisionCode returns the GeolocationSubdivisionCode field if non-nil, zero value otherwise.

### GetGeolocationSubdivisionCodeOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetGeolocationSubdivisionCodeOk() (*string, bool)`

GetGeolocationSubdivisionCodeOk returns a tuple with the GeolocationSubdivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocationSubdivisionCode

`func (o *RecordHttpsAwsRte53RecordInfo) SetGeolocationSubdivisionCode(v string)`

SetGeolocationSubdivisionCode sets GeolocationSubdivisionCode field to given value.

### HasGeolocationSubdivisionCode

`func (o *RecordHttpsAwsRte53RecordInfo) HasGeolocationSubdivisionCode() bool`

HasGeolocationSubdivisionCode returns a boolean if a field has been set.

### GetHealthCheckId

`func (o *RecordHttpsAwsRte53RecordInfo) GetHealthCheckId() string`

GetHealthCheckId returns the HealthCheckId field if non-nil, zero value otherwise.

### GetHealthCheckIdOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetHealthCheckIdOk() (*string, bool)`

GetHealthCheckIdOk returns a tuple with the HealthCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckId

`func (o *RecordHttpsAwsRte53RecordInfo) SetHealthCheckId(v string)`

SetHealthCheckId sets HealthCheckId field to given value.

### HasHealthCheckId

`func (o *RecordHttpsAwsRte53RecordInfo) HasHealthCheckId() bool`

HasHealthCheckId returns a boolean if a field has been set.

### GetRegion

`func (o *RecordHttpsAwsRte53RecordInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RecordHttpsAwsRte53RecordInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RecordHttpsAwsRte53RecordInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSetIdentifier

`func (o *RecordHttpsAwsRte53RecordInfo) GetSetIdentifier() string`

GetSetIdentifier returns the SetIdentifier field if non-nil, zero value otherwise.

### GetSetIdentifierOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetSetIdentifierOk() (*string, bool)`

GetSetIdentifierOk returns a tuple with the SetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetIdentifier

`func (o *RecordHttpsAwsRte53RecordInfo) SetSetIdentifier(v string)`

SetSetIdentifier sets SetIdentifier field to given value.

### HasSetIdentifier

`func (o *RecordHttpsAwsRte53RecordInfo) HasSetIdentifier() bool`

HasSetIdentifier returns a boolean if a field has been set.

### GetType

`func (o *RecordHttpsAwsRte53RecordInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordHttpsAwsRte53RecordInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RecordHttpsAwsRte53RecordInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWeight

`func (o *RecordHttpsAwsRte53RecordInfo) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RecordHttpsAwsRte53RecordInfo) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RecordHttpsAwsRte53RecordInfo) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RecordHttpsAwsRte53RecordInfo) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


