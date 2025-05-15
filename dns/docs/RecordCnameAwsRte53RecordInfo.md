# RecordCnameAwsRte53RecordInfo

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

### NewRecordCnameAwsRte53RecordInfo

`func NewRecordCnameAwsRte53RecordInfo() *RecordCnameAwsRte53RecordInfo`

NewRecordCnameAwsRte53RecordInfo instantiates a new RecordCnameAwsRte53RecordInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordCnameAwsRte53RecordInfoWithDefaults

`func NewRecordCnameAwsRte53RecordInfoWithDefaults() *RecordCnameAwsRte53RecordInfo`

NewRecordCnameAwsRte53RecordInfoWithDefaults instantiates a new RecordCnameAwsRte53RecordInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasTargetDnsName

`func (o *RecordCnameAwsRte53RecordInfo) GetAliasTargetDnsName() string`

GetAliasTargetDnsName returns the AliasTargetDnsName field if non-nil, zero value otherwise.

### GetAliasTargetDnsNameOk

`func (o *RecordCnameAwsRte53RecordInfo) GetAliasTargetDnsNameOk() (*string, bool)`

GetAliasTargetDnsNameOk returns a tuple with the AliasTargetDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetDnsName

`func (o *RecordCnameAwsRte53RecordInfo) SetAliasTargetDnsName(v string)`

SetAliasTargetDnsName sets AliasTargetDnsName field to given value.

### HasAliasTargetDnsName

`func (o *RecordCnameAwsRte53RecordInfo) HasAliasTargetDnsName() bool`

HasAliasTargetDnsName returns a boolean if a field has been set.

### GetAliasTargetHostedZoneId

`func (o *RecordCnameAwsRte53RecordInfo) GetAliasTargetHostedZoneId() string`

GetAliasTargetHostedZoneId returns the AliasTargetHostedZoneId field if non-nil, zero value otherwise.

### GetAliasTargetHostedZoneIdOk

`func (o *RecordCnameAwsRte53RecordInfo) GetAliasTargetHostedZoneIdOk() (*string, bool)`

GetAliasTargetHostedZoneIdOk returns a tuple with the AliasTargetHostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetHostedZoneId

`func (o *RecordCnameAwsRte53RecordInfo) SetAliasTargetHostedZoneId(v string)`

SetAliasTargetHostedZoneId sets AliasTargetHostedZoneId field to given value.

### HasAliasTargetHostedZoneId

`func (o *RecordCnameAwsRte53RecordInfo) HasAliasTargetHostedZoneId() bool`

HasAliasTargetHostedZoneId returns a boolean if a field has been set.

### GetAliasTargetEvaluateTargetHealth

`func (o *RecordCnameAwsRte53RecordInfo) GetAliasTargetEvaluateTargetHealth() bool`

GetAliasTargetEvaluateTargetHealth returns the AliasTargetEvaluateTargetHealth field if non-nil, zero value otherwise.

### GetAliasTargetEvaluateTargetHealthOk

`func (o *RecordCnameAwsRte53RecordInfo) GetAliasTargetEvaluateTargetHealthOk() (*bool, bool)`

GetAliasTargetEvaluateTargetHealthOk returns a tuple with the AliasTargetEvaluateTargetHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasTargetEvaluateTargetHealth

`func (o *RecordCnameAwsRte53RecordInfo) SetAliasTargetEvaluateTargetHealth(v bool)`

SetAliasTargetEvaluateTargetHealth sets AliasTargetEvaluateTargetHealth field to given value.

### HasAliasTargetEvaluateTargetHealth

`func (o *RecordCnameAwsRte53RecordInfo) HasAliasTargetEvaluateTargetHealth() bool`

HasAliasTargetEvaluateTargetHealth returns a boolean if a field has been set.

### GetFailover

`func (o *RecordCnameAwsRte53RecordInfo) GetFailover() string`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *RecordCnameAwsRte53RecordInfo) GetFailoverOk() (*string, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *RecordCnameAwsRte53RecordInfo) SetFailover(v string)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *RecordCnameAwsRte53RecordInfo) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetGeolocationContinentCode

`func (o *RecordCnameAwsRte53RecordInfo) GetGeolocationContinentCode() string`

GetGeolocationContinentCode returns the GeolocationContinentCode field if non-nil, zero value otherwise.

### GetGeolocationContinentCodeOk

`func (o *RecordCnameAwsRte53RecordInfo) GetGeolocationContinentCodeOk() (*string, bool)`

GetGeolocationContinentCodeOk returns a tuple with the GeolocationContinentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocationContinentCode

`func (o *RecordCnameAwsRte53RecordInfo) SetGeolocationContinentCode(v string)`

SetGeolocationContinentCode sets GeolocationContinentCode field to given value.

### HasGeolocationContinentCode

`func (o *RecordCnameAwsRte53RecordInfo) HasGeolocationContinentCode() bool`

HasGeolocationContinentCode returns a boolean if a field has been set.

### GetGeolocationCountryCode

`func (o *RecordCnameAwsRte53RecordInfo) GetGeolocationCountryCode() string`

GetGeolocationCountryCode returns the GeolocationCountryCode field if non-nil, zero value otherwise.

### GetGeolocationCountryCodeOk

`func (o *RecordCnameAwsRte53RecordInfo) GetGeolocationCountryCodeOk() (*string, bool)`

GetGeolocationCountryCodeOk returns a tuple with the GeolocationCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocationCountryCode

`func (o *RecordCnameAwsRte53RecordInfo) SetGeolocationCountryCode(v string)`

SetGeolocationCountryCode sets GeolocationCountryCode field to given value.

### HasGeolocationCountryCode

`func (o *RecordCnameAwsRte53RecordInfo) HasGeolocationCountryCode() bool`

HasGeolocationCountryCode returns a boolean if a field has been set.

### GetGeolocationSubdivisionCode

`func (o *RecordCnameAwsRte53RecordInfo) GetGeolocationSubdivisionCode() string`

GetGeolocationSubdivisionCode returns the GeolocationSubdivisionCode field if non-nil, zero value otherwise.

### GetGeolocationSubdivisionCodeOk

`func (o *RecordCnameAwsRte53RecordInfo) GetGeolocationSubdivisionCodeOk() (*string, bool)`

GetGeolocationSubdivisionCodeOk returns a tuple with the GeolocationSubdivisionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeolocationSubdivisionCode

`func (o *RecordCnameAwsRte53RecordInfo) SetGeolocationSubdivisionCode(v string)`

SetGeolocationSubdivisionCode sets GeolocationSubdivisionCode field to given value.

### HasGeolocationSubdivisionCode

`func (o *RecordCnameAwsRte53RecordInfo) HasGeolocationSubdivisionCode() bool`

HasGeolocationSubdivisionCode returns a boolean if a field has been set.

### GetHealthCheckId

`func (o *RecordCnameAwsRte53RecordInfo) GetHealthCheckId() string`

GetHealthCheckId returns the HealthCheckId field if non-nil, zero value otherwise.

### GetHealthCheckIdOk

`func (o *RecordCnameAwsRte53RecordInfo) GetHealthCheckIdOk() (*string, bool)`

GetHealthCheckIdOk returns a tuple with the HealthCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckId

`func (o *RecordCnameAwsRte53RecordInfo) SetHealthCheckId(v string)`

SetHealthCheckId sets HealthCheckId field to given value.

### HasHealthCheckId

`func (o *RecordCnameAwsRte53RecordInfo) HasHealthCheckId() bool`

HasHealthCheckId returns a boolean if a field has been set.

### GetRegion

`func (o *RecordCnameAwsRte53RecordInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RecordCnameAwsRte53RecordInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RecordCnameAwsRte53RecordInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RecordCnameAwsRte53RecordInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSetIdentifier

`func (o *RecordCnameAwsRte53RecordInfo) GetSetIdentifier() string`

GetSetIdentifier returns the SetIdentifier field if non-nil, zero value otherwise.

### GetSetIdentifierOk

`func (o *RecordCnameAwsRte53RecordInfo) GetSetIdentifierOk() (*string, bool)`

GetSetIdentifierOk returns a tuple with the SetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetIdentifier

`func (o *RecordCnameAwsRte53RecordInfo) SetSetIdentifier(v string)`

SetSetIdentifier sets SetIdentifier field to given value.

### HasSetIdentifier

`func (o *RecordCnameAwsRte53RecordInfo) HasSetIdentifier() bool`

HasSetIdentifier returns a boolean if a field has been set.

### GetType

`func (o *RecordCnameAwsRte53RecordInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordCnameAwsRte53RecordInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordCnameAwsRte53RecordInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RecordCnameAwsRte53RecordInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWeight

`func (o *RecordCnameAwsRte53RecordInfo) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RecordCnameAwsRte53RecordInfo) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RecordCnameAwsRte53RecordInfo) SetWeight(v int64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *RecordCnameAwsRte53RecordInfo) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


