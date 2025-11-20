# ParentalcontrolSubscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AltSubscriberId** | Pointer to **string** | The name of AVP to be used as an alternate subscriber ID for fixed lines. | [optional] 
**AltSubscriberIdRegexp** | Pointer to **string** | A character string to control aspects of rewriting of the fields. | [optional] 
**AltSubscriberIdSubexpression** | Pointer to **int64** | The subexpression indicates which subexpression to extract. If zero, then the text matching the entire regular expression is extracted. If non-zero, then the regex must contain at least that many sub-expression groups. It takes values from 0 to 8. | [optional] 
**Ancillaries** | Pointer to **[]string** | The list of ordered AVP Ancillary Fields. | [optional] 
**CatAcctname** | Pointer to **string** | Category content account name using the categorization service. | [optional] 
**CatPassword** | Pointer to **string** | Category content account password to access the categorization service. | [optional] 
**CatUpdateFrequency** | Pointer to **int64** | Category content updates every number of hours. | [optional] 
**CategoryUrl** | Pointer to **string** | Category content vendor url to download category data from and upload feedback to, configure for parental control. | [optional] 
**EnableMgmtOnlyNas** | Pointer to **bool** | Determines if NAS RADIUS traffic is accepted over MGMT only. | [optional] 
**EnableParentalControl** | Pointer to **bool** | Determines if parental control is enabled. | [optional] 
**InterimAccountingInterval** | Pointer to **int64** | The time for collector to be fully populated. Valid values are from 1 to 65535. | [optional] 
**IpAnchors** | Pointer to **[]string** | The ordered list of IP Anchors AVPs. The list content cannot be changed, but the order of elements. | [optional] 
**IpSpaceDiscRegexp** | Pointer to **string** | A character string to control aspects of rewriting of the fields. | [optional] 
**IpSpaceDiscSubexpression** | Pointer to **int64** | The subexpression indicates which subexpression to extract. If zero, then the text matching the entire regular expression is extracted. If non-zero, then the regex must contain at least that many sub-expression groups. It takes values from 0 to 8. | [optional] 
**IpSpaceDiscriminator** | Pointer to **string** | The name of AVP to be used as IP address discriminator. | [optional] 
**LocalId** | Pointer to **string** | The name of AVP to be used as local ID. | [optional] 
**LocalIdRegexp** | Pointer to **string** | A character string to control aspects of rewriting of the fields. | [optional] 
**LocalIdSubexpression** | Pointer to **int64** | The subexpression indicates which subexpression to extract. If zero, then the text matching the entire regular expression is extracted. If non-zero, then the regex must contain at least that many sub-expression groups. It takes values from 0 to 8. | [optional] 
**LogGuestLookups** | Pointer to **bool** | CEF log all guest lookups, will produce two logs in case of a violation. | [optional] 
**NasContextInfo** | Pointer to **string** | NAS contextual information AVP. | [optional] 
**PcZoneName** | Pointer to **string** | The SOA to store parental control records. | [optional] 
**ProxyPassword** | Pointer to **string** | Proxy server password used for authentication. | [optional] 
**ProxyUrl** | Pointer to **string** | Proxy url to download category data from and upload feedback to, configure for parental control. The default value &#39;None&#39; is no longer valid as it match url regex pattern \&quot;^http|https://\&quot;. The new default value does not get saved in database, but rather used for comparision with object created in unit test cases. | [optional] 
**ProxyUsername** | Pointer to **string** | Proxy server username used for authentication. | [optional] 
**SubscriberId** | Pointer to **string** | The name of AVP to be used as a subscriber. | [optional] 
**SubscriberIdRegexp** | Pointer to **string** | A character string to control aspects of rewriting of the fields. | [optional] 
**SubscriberIdSubexpression** | Pointer to **int64** | The subexpression indicates which subexpression to extract. If zero, then the text matching the entire regular expression is extracted. If non-zero, then the regex must contain at least that many sub-expression groups. It takes values from 0 to 8. | [optional] 
**ZveloUpdateFailureInDays** | Pointer to **int64** | Number of days since zvelo DB failed to update. | [optional] [readonly] 

## Methods

### NewParentalcontrolSubscriber

`func NewParentalcontrolSubscriber() *ParentalcontrolSubscriber`

NewParentalcontrolSubscriber instantiates a new ParentalcontrolSubscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentalcontrolSubscriberWithDefaults

`func NewParentalcontrolSubscriberWithDefaults() *ParentalcontrolSubscriber`

NewParentalcontrolSubscriberWithDefaults instantiates a new ParentalcontrolSubscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ParentalcontrolSubscriber) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ParentalcontrolSubscriber) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ParentalcontrolSubscriber) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ParentalcontrolSubscriber) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAltSubscriberId

`func (o *ParentalcontrolSubscriber) GetAltSubscriberId() string`

GetAltSubscriberId returns the AltSubscriberId field if non-nil, zero value otherwise.

### GetAltSubscriberIdOk

`func (o *ParentalcontrolSubscriber) GetAltSubscriberIdOk() (*string, bool)`

GetAltSubscriberIdOk returns a tuple with the AltSubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSubscriberId

`func (o *ParentalcontrolSubscriber) SetAltSubscriberId(v string)`

SetAltSubscriberId sets AltSubscriberId field to given value.

### HasAltSubscriberId

`func (o *ParentalcontrolSubscriber) HasAltSubscriberId() bool`

HasAltSubscriberId returns a boolean if a field has been set.

### GetAltSubscriberIdRegexp

`func (o *ParentalcontrolSubscriber) GetAltSubscriberIdRegexp() string`

GetAltSubscriberIdRegexp returns the AltSubscriberIdRegexp field if non-nil, zero value otherwise.

### GetAltSubscriberIdRegexpOk

`func (o *ParentalcontrolSubscriber) GetAltSubscriberIdRegexpOk() (*string, bool)`

GetAltSubscriberIdRegexpOk returns a tuple with the AltSubscriberIdRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSubscriberIdRegexp

`func (o *ParentalcontrolSubscriber) SetAltSubscriberIdRegexp(v string)`

SetAltSubscriberIdRegexp sets AltSubscriberIdRegexp field to given value.

### HasAltSubscriberIdRegexp

`func (o *ParentalcontrolSubscriber) HasAltSubscriberIdRegexp() bool`

HasAltSubscriberIdRegexp returns a boolean if a field has been set.

### GetAltSubscriberIdSubexpression

`func (o *ParentalcontrolSubscriber) GetAltSubscriberIdSubexpression() int64`

GetAltSubscriberIdSubexpression returns the AltSubscriberIdSubexpression field if non-nil, zero value otherwise.

### GetAltSubscriberIdSubexpressionOk

`func (o *ParentalcontrolSubscriber) GetAltSubscriberIdSubexpressionOk() (*int64, bool)`

GetAltSubscriberIdSubexpressionOk returns a tuple with the AltSubscriberIdSubexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSubscriberIdSubexpression

`func (o *ParentalcontrolSubscriber) SetAltSubscriberIdSubexpression(v int64)`

SetAltSubscriberIdSubexpression sets AltSubscriberIdSubexpression field to given value.

### HasAltSubscriberIdSubexpression

`func (o *ParentalcontrolSubscriber) HasAltSubscriberIdSubexpression() bool`

HasAltSubscriberIdSubexpression returns a boolean if a field has been set.

### GetAncillaries

`func (o *ParentalcontrolSubscriber) GetAncillaries() []string`

GetAncillaries returns the Ancillaries field if non-nil, zero value otherwise.

### GetAncillariesOk

`func (o *ParentalcontrolSubscriber) GetAncillariesOk() (*[]string, bool)`

GetAncillariesOk returns a tuple with the Ancillaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncillaries

`func (o *ParentalcontrolSubscriber) SetAncillaries(v []string)`

SetAncillaries sets Ancillaries field to given value.

### HasAncillaries

`func (o *ParentalcontrolSubscriber) HasAncillaries() bool`

HasAncillaries returns a boolean if a field has been set.

### GetCatAcctname

`func (o *ParentalcontrolSubscriber) GetCatAcctname() string`

GetCatAcctname returns the CatAcctname field if non-nil, zero value otherwise.

### GetCatAcctnameOk

`func (o *ParentalcontrolSubscriber) GetCatAcctnameOk() (*string, bool)`

GetCatAcctnameOk returns a tuple with the CatAcctname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatAcctname

`func (o *ParentalcontrolSubscriber) SetCatAcctname(v string)`

SetCatAcctname sets CatAcctname field to given value.

### HasCatAcctname

`func (o *ParentalcontrolSubscriber) HasCatAcctname() bool`

HasCatAcctname returns a boolean if a field has been set.

### GetCatPassword

`func (o *ParentalcontrolSubscriber) GetCatPassword() string`

GetCatPassword returns the CatPassword field if non-nil, zero value otherwise.

### GetCatPasswordOk

`func (o *ParentalcontrolSubscriber) GetCatPasswordOk() (*string, bool)`

GetCatPasswordOk returns a tuple with the CatPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatPassword

`func (o *ParentalcontrolSubscriber) SetCatPassword(v string)`

SetCatPassword sets CatPassword field to given value.

### HasCatPassword

`func (o *ParentalcontrolSubscriber) HasCatPassword() bool`

HasCatPassword returns a boolean if a field has been set.

### GetCatUpdateFrequency

`func (o *ParentalcontrolSubscriber) GetCatUpdateFrequency() int64`

GetCatUpdateFrequency returns the CatUpdateFrequency field if non-nil, zero value otherwise.

### GetCatUpdateFrequencyOk

`func (o *ParentalcontrolSubscriber) GetCatUpdateFrequencyOk() (*int64, bool)`

GetCatUpdateFrequencyOk returns a tuple with the CatUpdateFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatUpdateFrequency

`func (o *ParentalcontrolSubscriber) SetCatUpdateFrequency(v int64)`

SetCatUpdateFrequency sets CatUpdateFrequency field to given value.

### HasCatUpdateFrequency

`func (o *ParentalcontrolSubscriber) HasCatUpdateFrequency() bool`

HasCatUpdateFrequency returns a boolean if a field has been set.

### GetCategoryUrl

`func (o *ParentalcontrolSubscriber) GetCategoryUrl() string`

GetCategoryUrl returns the CategoryUrl field if non-nil, zero value otherwise.

### GetCategoryUrlOk

`func (o *ParentalcontrolSubscriber) GetCategoryUrlOk() (*string, bool)`

GetCategoryUrlOk returns a tuple with the CategoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryUrl

`func (o *ParentalcontrolSubscriber) SetCategoryUrl(v string)`

SetCategoryUrl sets CategoryUrl field to given value.

### HasCategoryUrl

`func (o *ParentalcontrolSubscriber) HasCategoryUrl() bool`

HasCategoryUrl returns a boolean if a field has been set.

### GetEnableMgmtOnlyNas

`func (o *ParentalcontrolSubscriber) GetEnableMgmtOnlyNas() bool`

GetEnableMgmtOnlyNas returns the EnableMgmtOnlyNas field if non-nil, zero value otherwise.

### GetEnableMgmtOnlyNasOk

`func (o *ParentalcontrolSubscriber) GetEnableMgmtOnlyNasOk() (*bool, bool)`

GetEnableMgmtOnlyNasOk returns a tuple with the EnableMgmtOnlyNas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMgmtOnlyNas

`func (o *ParentalcontrolSubscriber) SetEnableMgmtOnlyNas(v bool)`

SetEnableMgmtOnlyNas sets EnableMgmtOnlyNas field to given value.

### HasEnableMgmtOnlyNas

`func (o *ParentalcontrolSubscriber) HasEnableMgmtOnlyNas() bool`

HasEnableMgmtOnlyNas returns a boolean if a field has been set.

### GetEnableParentalControl

`func (o *ParentalcontrolSubscriber) GetEnableParentalControl() bool`

GetEnableParentalControl returns the EnableParentalControl field if non-nil, zero value otherwise.

### GetEnableParentalControlOk

`func (o *ParentalcontrolSubscriber) GetEnableParentalControlOk() (*bool, bool)`

GetEnableParentalControlOk returns a tuple with the EnableParentalControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableParentalControl

`func (o *ParentalcontrolSubscriber) SetEnableParentalControl(v bool)`

SetEnableParentalControl sets EnableParentalControl field to given value.

### HasEnableParentalControl

`func (o *ParentalcontrolSubscriber) HasEnableParentalControl() bool`

HasEnableParentalControl returns a boolean if a field has been set.

### GetInterimAccountingInterval

`func (o *ParentalcontrolSubscriber) GetInterimAccountingInterval() int64`

GetInterimAccountingInterval returns the InterimAccountingInterval field if non-nil, zero value otherwise.

### GetInterimAccountingIntervalOk

`func (o *ParentalcontrolSubscriber) GetInterimAccountingIntervalOk() (*int64, bool)`

GetInterimAccountingIntervalOk returns a tuple with the InterimAccountingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimAccountingInterval

`func (o *ParentalcontrolSubscriber) SetInterimAccountingInterval(v int64)`

SetInterimAccountingInterval sets InterimAccountingInterval field to given value.

### HasInterimAccountingInterval

`func (o *ParentalcontrolSubscriber) HasInterimAccountingInterval() bool`

HasInterimAccountingInterval returns a boolean if a field has been set.

### GetIpAnchors

`func (o *ParentalcontrolSubscriber) GetIpAnchors() []string`

GetIpAnchors returns the IpAnchors field if non-nil, zero value otherwise.

### GetIpAnchorsOk

`func (o *ParentalcontrolSubscriber) GetIpAnchorsOk() (*[]string, bool)`

GetIpAnchorsOk returns a tuple with the IpAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAnchors

`func (o *ParentalcontrolSubscriber) SetIpAnchors(v []string)`

SetIpAnchors sets IpAnchors field to given value.

### HasIpAnchors

`func (o *ParentalcontrolSubscriber) HasIpAnchors() bool`

HasIpAnchors returns a boolean if a field has been set.

### GetIpSpaceDiscRegexp

`func (o *ParentalcontrolSubscriber) GetIpSpaceDiscRegexp() string`

GetIpSpaceDiscRegexp returns the IpSpaceDiscRegexp field if non-nil, zero value otherwise.

### GetIpSpaceDiscRegexpOk

`func (o *ParentalcontrolSubscriber) GetIpSpaceDiscRegexpOk() (*string, bool)`

GetIpSpaceDiscRegexpOk returns a tuple with the IpSpaceDiscRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceDiscRegexp

`func (o *ParentalcontrolSubscriber) SetIpSpaceDiscRegexp(v string)`

SetIpSpaceDiscRegexp sets IpSpaceDiscRegexp field to given value.

### HasIpSpaceDiscRegexp

`func (o *ParentalcontrolSubscriber) HasIpSpaceDiscRegexp() bool`

HasIpSpaceDiscRegexp returns a boolean if a field has been set.

### GetIpSpaceDiscSubexpression

`func (o *ParentalcontrolSubscriber) GetIpSpaceDiscSubexpression() int64`

GetIpSpaceDiscSubexpression returns the IpSpaceDiscSubexpression field if non-nil, zero value otherwise.

### GetIpSpaceDiscSubexpressionOk

`func (o *ParentalcontrolSubscriber) GetIpSpaceDiscSubexpressionOk() (*int64, bool)`

GetIpSpaceDiscSubexpressionOk returns a tuple with the IpSpaceDiscSubexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceDiscSubexpression

`func (o *ParentalcontrolSubscriber) SetIpSpaceDiscSubexpression(v int64)`

SetIpSpaceDiscSubexpression sets IpSpaceDiscSubexpression field to given value.

### HasIpSpaceDiscSubexpression

`func (o *ParentalcontrolSubscriber) HasIpSpaceDiscSubexpression() bool`

HasIpSpaceDiscSubexpression returns a boolean if a field has been set.

### GetIpSpaceDiscriminator

`func (o *ParentalcontrolSubscriber) GetIpSpaceDiscriminator() string`

GetIpSpaceDiscriminator returns the IpSpaceDiscriminator field if non-nil, zero value otherwise.

### GetIpSpaceDiscriminatorOk

`func (o *ParentalcontrolSubscriber) GetIpSpaceDiscriminatorOk() (*string, bool)`

GetIpSpaceDiscriminatorOk returns a tuple with the IpSpaceDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceDiscriminator

`func (o *ParentalcontrolSubscriber) SetIpSpaceDiscriminator(v string)`

SetIpSpaceDiscriminator sets IpSpaceDiscriminator field to given value.

### HasIpSpaceDiscriminator

`func (o *ParentalcontrolSubscriber) HasIpSpaceDiscriminator() bool`

HasIpSpaceDiscriminator returns a boolean if a field has been set.

### GetLocalId

`func (o *ParentalcontrolSubscriber) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *ParentalcontrolSubscriber) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *ParentalcontrolSubscriber) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *ParentalcontrolSubscriber) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetLocalIdRegexp

`func (o *ParentalcontrolSubscriber) GetLocalIdRegexp() string`

GetLocalIdRegexp returns the LocalIdRegexp field if non-nil, zero value otherwise.

### GetLocalIdRegexpOk

`func (o *ParentalcontrolSubscriber) GetLocalIdRegexpOk() (*string, bool)`

GetLocalIdRegexpOk returns a tuple with the LocalIdRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIdRegexp

`func (o *ParentalcontrolSubscriber) SetLocalIdRegexp(v string)`

SetLocalIdRegexp sets LocalIdRegexp field to given value.

### HasLocalIdRegexp

`func (o *ParentalcontrolSubscriber) HasLocalIdRegexp() bool`

HasLocalIdRegexp returns a boolean if a field has been set.

### GetLocalIdSubexpression

`func (o *ParentalcontrolSubscriber) GetLocalIdSubexpression() int64`

GetLocalIdSubexpression returns the LocalIdSubexpression field if non-nil, zero value otherwise.

### GetLocalIdSubexpressionOk

`func (o *ParentalcontrolSubscriber) GetLocalIdSubexpressionOk() (*int64, bool)`

GetLocalIdSubexpressionOk returns a tuple with the LocalIdSubexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIdSubexpression

`func (o *ParentalcontrolSubscriber) SetLocalIdSubexpression(v int64)`

SetLocalIdSubexpression sets LocalIdSubexpression field to given value.

### HasLocalIdSubexpression

`func (o *ParentalcontrolSubscriber) HasLocalIdSubexpression() bool`

HasLocalIdSubexpression returns a boolean if a field has been set.

### GetLogGuestLookups

`func (o *ParentalcontrolSubscriber) GetLogGuestLookups() bool`

GetLogGuestLookups returns the LogGuestLookups field if non-nil, zero value otherwise.

### GetLogGuestLookupsOk

`func (o *ParentalcontrolSubscriber) GetLogGuestLookupsOk() (*bool, bool)`

GetLogGuestLookupsOk returns a tuple with the LogGuestLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogGuestLookups

`func (o *ParentalcontrolSubscriber) SetLogGuestLookups(v bool)`

SetLogGuestLookups sets LogGuestLookups field to given value.

### HasLogGuestLookups

`func (o *ParentalcontrolSubscriber) HasLogGuestLookups() bool`

HasLogGuestLookups returns a boolean if a field has been set.

### GetNasContextInfo

`func (o *ParentalcontrolSubscriber) GetNasContextInfo() string`

GetNasContextInfo returns the NasContextInfo field if non-nil, zero value otherwise.

### GetNasContextInfoOk

`func (o *ParentalcontrolSubscriber) GetNasContextInfoOk() (*string, bool)`

GetNasContextInfoOk returns a tuple with the NasContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasContextInfo

`func (o *ParentalcontrolSubscriber) SetNasContextInfo(v string)`

SetNasContextInfo sets NasContextInfo field to given value.

### HasNasContextInfo

`func (o *ParentalcontrolSubscriber) HasNasContextInfo() bool`

HasNasContextInfo returns a boolean if a field has been set.

### GetPcZoneName

`func (o *ParentalcontrolSubscriber) GetPcZoneName() string`

GetPcZoneName returns the PcZoneName field if non-nil, zero value otherwise.

### GetPcZoneNameOk

`func (o *ParentalcontrolSubscriber) GetPcZoneNameOk() (*string, bool)`

GetPcZoneNameOk returns a tuple with the PcZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcZoneName

`func (o *ParentalcontrolSubscriber) SetPcZoneName(v string)`

SetPcZoneName sets PcZoneName field to given value.

### HasPcZoneName

`func (o *ParentalcontrolSubscriber) HasPcZoneName() bool`

HasPcZoneName returns a boolean if a field has been set.

### GetProxyPassword

`func (o *ParentalcontrolSubscriber) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *ParentalcontrolSubscriber) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *ParentalcontrolSubscriber) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *ParentalcontrolSubscriber) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetProxyUrl

`func (o *ParentalcontrolSubscriber) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *ParentalcontrolSubscriber) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *ParentalcontrolSubscriber) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *ParentalcontrolSubscriber) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### GetProxyUsername

`func (o *ParentalcontrolSubscriber) GetProxyUsername() string`

GetProxyUsername returns the ProxyUsername field if non-nil, zero value otherwise.

### GetProxyUsernameOk

`func (o *ParentalcontrolSubscriber) GetProxyUsernameOk() (*string, bool)`

GetProxyUsernameOk returns a tuple with the ProxyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUsername

`func (o *ParentalcontrolSubscriber) SetProxyUsername(v string)`

SetProxyUsername sets ProxyUsername field to given value.

### HasProxyUsername

`func (o *ParentalcontrolSubscriber) HasProxyUsername() bool`

HasProxyUsername returns a boolean if a field has been set.

### GetSubscriberId

`func (o *ParentalcontrolSubscriber) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *ParentalcontrolSubscriber) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *ParentalcontrolSubscriber) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *ParentalcontrolSubscriber) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetSubscriberIdRegexp

`func (o *ParentalcontrolSubscriber) GetSubscriberIdRegexp() string`

GetSubscriberIdRegexp returns the SubscriberIdRegexp field if non-nil, zero value otherwise.

### GetSubscriberIdRegexpOk

`func (o *ParentalcontrolSubscriber) GetSubscriberIdRegexpOk() (*string, bool)`

GetSubscriberIdRegexpOk returns a tuple with the SubscriberIdRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdRegexp

`func (o *ParentalcontrolSubscriber) SetSubscriberIdRegexp(v string)`

SetSubscriberIdRegexp sets SubscriberIdRegexp field to given value.

### HasSubscriberIdRegexp

`func (o *ParentalcontrolSubscriber) HasSubscriberIdRegexp() bool`

HasSubscriberIdRegexp returns a boolean if a field has been set.

### GetSubscriberIdSubexpression

`func (o *ParentalcontrolSubscriber) GetSubscriberIdSubexpression() int64`

GetSubscriberIdSubexpression returns the SubscriberIdSubexpression field if non-nil, zero value otherwise.

### GetSubscriberIdSubexpressionOk

`func (o *ParentalcontrolSubscriber) GetSubscriberIdSubexpressionOk() (*int64, bool)`

GetSubscriberIdSubexpressionOk returns a tuple with the SubscriberIdSubexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdSubexpression

`func (o *ParentalcontrolSubscriber) SetSubscriberIdSubexpression(v int64)`

SetSubscriberIdSubexpression sets SubscriberIdSubexpression field to given value.

### HasSubscriberIdSubexpression

`func (o *ParentalcontrolSubscriber) HasSubscriberIdSubexpression() bool`

HasSubscriberIdSubexpression returns a boolean if a field has been set.

### GetZveloUpdateFailureInDays

`func (o *ParentalcontrolSubscriber) GetZveloUpdateFailureInDays() int64`

GetZveloUpdateFailureInDays returns the ZveloUpdateFailureInDays field if non-nil, zero value otherwise.

### GetZveloUpdateFailureInDaysOk

`func (o *ParentalcontrolSubscriber) GetZveloUpdateFailureInDaysOk() (*int64, bool)`

GetZveloUpdateFailureInDaysOk returns a tuple with the ZveloUpdateFailureInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZveloUpdateFailureInDays

`func (o *ParentalcontrolSubscriber) SetZveloUpdateFailureInDays(v int64)`

SetZveloUpdateFailureInDays sets ZveloUpdateFailureInDays field to given value.

### HasZveloUpdateFailureInDays

`func (o *ParentalcontrolSubscriber) HasZveloUpdateFailureInDays() bool`

HasZveloUpdateFailureInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


