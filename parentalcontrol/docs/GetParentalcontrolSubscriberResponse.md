# GetParentalcontrolSubscriberResponse

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
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**ZveloUpdateFailureInDays** | Pointer to **int64** | Number of days since zvelo DB failed to update. | [optional] [readonly] 
**Result** | Pointer to [**ParentalcontrolSubscriber**](ParentalcontrolSubscriber.md) |  | [optional] 

## Methods

### NewGetParentalcontrolSubscriberResponse

`func NewGetParentalcontrolSubscriberResponse() *GetParentalcontrolSubscriberResponse`

NewGetParentalcontrolSubscriberResponse instantiates a new GetParentalcontrolSubscriberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetParentalcontrolSubscriberResponseWithDefaults

`func NewGetParentalcontrolSubscriberResponseWithDefaults() *GetParentalcontrolSubscriberResponse`

NewGetParentalcontrolSubscriberResponseWithDefaults instantiates a new GetParentalcontrolSubscriberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetParentalcontrolSubscriberResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetParentalcontrolSubscriberResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetParentalcontrolSubscriberResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetParentalcontrolSubscriberResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAltSubscriberId

`func (o *GetParentalcontrolSubscriberResponse) GetAltSubscriberId() string`

GetAltSubscriberId returns the AltSubscriberId field if non-nil, zero value otherwise.

### GetAltSubscriberIdOk

`func (o *GetParentalcontrolSubscriberResponse) GetAltSubscriberIdOk() (*string, bool)`

GetAltSubscriberIdOk returns a tuple with the AltSubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSubscriberId

`func (o *GetParentalcontrolSubscriberResponse) SetAltSubscriberId(v string)`

SetAltSubscriberId sets AltSubscriberId field to given value.

### HasAltSubscriberId

`func (o *GetParentalcontrolSubscriberResponse) HasAltSubscriberId() bool`

HasAltSubscriberId returns a boolean if a field has been set.

### GetAltSubscriberIdRegexp

`func (o *GetParentalcontrolSubscriberResponse) GetAltSubscriberIdRegexp() string`

GetAltSubscriberIdRegexp returns the AltSubscriberIdRegexp field if non-nil, zero value otherwise.

### GetAltSubscriberIdRegexpOk

`func (o *GetParentalcontrolSubscriberResponse) GetAltSubscriberIdRegexpOk() (*string, bool)`

GetAltSubscriberIdRegexpOk returns a tuple with the AltSubscriberIdRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSubscriberIdRegexp

`func (o *GetParentalcontrolSubscriberResponse) SetAltSubscriberIdRegexp(v string)`

SetAltSubscriberIdRegexp sets AltSubscriberIdRegexp field to given value.

### HasAltSubscriberIdRegexp

`func (o *GetParentalcontrolSubscriberResponse) HasAltSubscriberIdRegexp() bool`

HasAltSubscriberIdRegexp returns a boolean if a field has been set.

### GetAltSubscriberIdSubexpression

`func (o *GetParentalcontrolSubscriberResponse) GetAltSubscriberIdSubexpression() int64`

GetAltSubscriberIdSubexpression returns the AltSubscriberIdSubexpression field if non-nil, zero value otherwise.

### GetAltSubscriberIdSubexpressionOk

`func (o *GetParentalcontrolSubscriberResponse) GetAltSubscriberIdSubexpressionOk() (*int64, bool)`

GetAltSubscriberIdSubexpressionOk returns a tuple with the AltSubscriberIdSubexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltSubscriberIdSubexpression

`func (o *GetParentalcontrolSubscriberResponse) SetAltSubscriberIdSubexpression(v int64)`

SetAltSubscriberIdSubexpression sets AltSubscriberIdSubexpression field to given value.

### HasAltSubscriberIdSubexpression

`func (o *GetParentalcontrolSubscriberResponse) HasAltSubscriberIdSubexpression() bool`

HasAltSubscriberIdSubexpression returns a boolean if a field has been set.

### GetAncillaries

`func (o *GetParentalcontrolSubscriberResponse) GetAncillaries() []string`

GetAncillaries returns the Ancillaries field if non-nil, zero value otherwise.

### GetAncillariesOk

`func (o *GetParentalcontrolSubscriberResponse) GetAncillariesOk() (*[]string, bool)`

GetAncillariesOk returns a tuple with the Ancillaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncillaries

`func (o *GetParentalcontrolSubscriberResponse) SetAncillaries(v []string)`

SetAncillaries sets Ancillaries field to given value.

### HasAncillaries

`func (o *GetParentalcontrolSubscriberResponse) HasAncillaries() bool`

HasAncillaries returns a boolean if a field has been set.

### GetCatAcctname

`func (o *GetParentalcontrolSubscriberResponse) GetCatAcctname() string`

GetCatAcctname returns the CatAcctname field if non-nil, zero value otherwise.

### GetCatAcctnameOk

`func (o *GetParentalcontrolSubscriberResponse) GetCatAcctnameOk() (*string, bool)`

GetCatAcctnameOk returns a tuple with the CatAcctname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatAcctname

`func (o *GetParentalcontrolSubscriberResponse) SetCatAcctname(v string)`

SetCatAcctname sets CatAcctname field to given value.

### HasCatAcctname

`func (o *GetParentalcontrolSubscriberResponse) HasCatAcctname() bool`

HasCatAcctname returns a boolean if a field has been set.

### GetCatPassword

`func (o *GetParentalcontrolSubscriberResponse) GetCatPassword() string`

GetCatPassword returns the CatPassword field if non-nil, zero value otherwise.

### GetCatPasswordOk

`func (o *GetParentalcontrolSubscriberResponse) GetCatPasswordOk() (*string, bool)`

GetCatPasswordOk returns a tuple with the CatPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatPassword

`func (o *GetParentalcontrolSubscriberResponse) SetCatPassword(v string)`

SetCatPassword sets CatPassword field to given value.

### HasCatPassword

`func (o *GetParentalcontrolSubscriberResponse) HasCatPassword() bool`

HasCatPassword returns a boolean if a field has been set.

### GetCatUpdateFrequency

`func (o *GetParentalcontrolSubscriberResponse) GetCatUpdateFrequency() int64`

GetCatUpdateFrequency returns the CatUpdateFrequency field if non-nil, zero value otherwise.

### GetCatUpdateFrequencyOk

`func (o *GetParentalcontrolSubscriberResponse) GetCatUpdateFrequencyOk() (*int64, bool)`

GetCatUpdateFrequencyOk returns a tuple with the CatUpdateFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatUpdateFrequency

`func (o *GetParentalcontrolSubscriberResponse) SetCatUpdateFrequency(v int64)`

SetCatUpdateFrequency sets CatUpdateFrequency field to given value.

### HasCatUpdateFrequency

`func (o *GetParentalcontrolSubscriberResponse) HasCatUpdateFrequency() bool`

HasCatUpdateFrequency returns a boolean if a field has been set.

### GetCategoryUrl

`func (o *GetParentalcontrolSubscriberResponse) GetCategoryUrl() string`

GetCategoryUrl returns the CategoryUrl field if non-nil, zero value otherwise.

### GetCategoryUrlOk

`func (o *GetParentalcontrolSubscriberResponse) GetCategoryUrlOk() (*string, bool)`

GetCategoryUrlOk returns a tuple with the CategoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryUrl

`func (o *GetParentalcontrolSubscriberResponse) SetCategoryUrl(v string)`

SetCategoryUrl sets CategoryUrl field to given value.

### HasCategoryUrl

`func (o *GetParentalcontrolSubscriberResponse) HasCategoryUrl() bool`

HasCategoryUrl returns a boolean if a field has been set.

### GetEnableMgmtOnlyNas

`func (o *GetParentalcontrolSubscriberResponse) GetEnableMgmtOnlyNas() bool`

GetEnableMgmtOnlyNas returns the EnableMgmtOnlyNas field if non-nil, zero value otherwise.

### GetEnableMgmtOnlyNasOk

`func (o *GetParentalcontrolSubscriberResponse) GetEnableMgmtOnlyNasOk() (*bool, bool)`

GetEnableMgmtOnlyNasOk returns a tuple with the EnableMgmtOnlyNas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMgmtOnlyNas

`func (o *GetParentalcontrolSubscriberResponse) SetEnableMgmtOnlyNas(v bool)`

SetEnableMgmtOnlyNas sets EnableMgmtOnlyNas field to given value.

### HasEnableMgmtOnlyNas

`func (o *GetParentalcontrolSubscriberResponse) HasEnableMgmtOnlyNas() bool`

HasEnableMgmtOnlyNas returns a boolean if a field has been set.

### GetEnableParentalControl

`func (o *GetParentalcontrolSubscriberResponse) GetEnableParentalControl() bool`

GetEnableParentalControl returns the EnableParentalControl field if non-nil, zero value otherwise.

### GetEnableParentalControlOk

`func (o *GetParentalcontrolSubscriberResponse) GetEnableParentalControlOk() (*bool, bool)`

GetEnableParentalControlOk returns a tuple with the EnableParentalControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableParentalControl

`func (o *GetParentalcontrolSubscriberResponse) SetEnableParentalControl(v bool)`

SetEnableParentalControl sets EnableParentalControl field to given value.

### HasEnableParentalControl

`func (o *GetParentalcontrolSubscriberResponse) HasEnableParentalControl() bool`

HasEnableParentalControl returns a boolean if a field has been set.

### GetInterimAccountingInterval

`func (o *GetParentalcontrolSubscriberResponse) GetInterimAccountingInterval() int64`

GetInterimAccountingInterval returns the InterimAccountingInterval field if non-nil, zero value otherwise.

### GetInterimAccountingIntervalOk

`func (o *GetParentalcontrolSubscriberResponse) GetInterimAccountingIntervalOk() (*int64, bool)`

GetInterimAccountingIntervalOk returns a tuple with the InterimAccountingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimAccountingInterval

`func (o *GetParentalcontrolSubscriberResponse) SetInterimAccountingInterval(v int64)`

SetInterimAccountingInterval sets InterimAccountingInterval field to given value.

### HasInterimAccountingInterval

`func (o *GetParentalcontrolSubscriberResponse) HasInterimAccountingInterval() bool`

HasInterimAccountingInterval returns a boolean if a field has been set.

### GetIpAnchors

`func (o *GetParentalcontrolSubscriberResponse) GetIpAnchors() []string`

GetIpAnchors returns the IpAnchors field if non-nil, zero value otherwise.

### GetIpAnchorsOk

`func (o *GetParentalcontrolSubscriberResponse) GetIpAnchorsOk() (*[]string, bool)`

GetIpAnchorsOk returns a tuple with the IpAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAnchors

`func (o *GetParentalcontrolSubscriberResponse) SetIpAnchors(v []string)`

SetIpAnchors sets IpAnchors field to given value.

### HasIpAnchors

`func (o *GetParentalcontrolSubscriberResponse) HasIpAnchors() bool`

HasIpAnchors returns a boolean if a field has been set.

### GetIpSpaceDiscRegexp

`func (o *GetParentalcontrolSubscriberResponse) GetIpSpaceDiscRegexp() string`

GetIpSpaceDiscRegexp returns the IpSpaceDiscRegexp field if non-nil, zero value otherwise.

### GetIpSpaceDiscRegexpOk

`func (o *GetParentalcontrolSubscriberResponse) GetIpSpaceDiscRegexpOk() (*string, bool)`

GetIpSpaceDiscRegexpOk returns a tuple with the IpSpaceDiscRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceDiscRegexp

`func (o *GetParentalcontrolSubscriberResponse) SetIpSpaceDiscRegexp(v string)`

SetIpSpaceDiscRegexp sets IpSpaceDiscRegexp field to given value.

### HasIpSpaceDiscRegexp

`func (o *GetParentalcontrolSubscriberResponse) HasIpSpaceDiscRegexp() bool`

HasIpSpaceDiscRegexp returns a boolean if a field has been set.

### GetIpSpaceDiscSubexpression

`func (o *GetParentalcontrolSubscriberResponse) GetIpSpaceDiscSubexpression() int64`

GetIpSpaceDiscSubexpression returns the IpSpaceDiscSubexpression field if non-nil, zero value otherwise.

### GetIpSpaceDiscSubexpressionOk

`func (o *GetParentalcontrolSubscriberResponse) GetIpSpaceDiscSubexpressionOk() (*int64, bool)`

GetIpSpaceDiscSubexpressionOk returns a tuple with the IpSpaceDiscSubexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceDiscSubexpression

`func (o *GetParentalcontrolSubscriberResponse) SetIpSpaceDiscSubexpression(v int64)`

SetIpSpaceDiscSubexpression sets IpSpaceDiscSubexpression field to given value.

### HasIpSpaceDiscSubexpression

`func (o *GetParentalcontrolSubscriberResponse) HasIpSpaceDiscSubexpression() bool`

HasIpSpaceDiscSubexpression returns a boolean if a field has been set.

### GetIpSpaceDiscriminator

`func (o *GetParentalcontrolSubscriberResponse) GetIpSpaceDiscriminator() string`

GetIpSpaceDiscriminator returns the IpSpaceDiscriminator field if non-nil, zero value otherwise.

### GetIpSpaceDiscriminatorOk

`func (o *GetParentalcontrolSubscriberResponse) GetIpSpaceDiscriminatorOk() (*string, bool)`

GetIpSpaceDiscriminatorOk returns a tuple with the IpSpaceDiscriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceDiscriminator

`func (o *GetParentalcontrolSubscriberResponse) SetIpSpaceDiscriminator(v string)`

SetIpSpaceDiscriminator sets IpSpaceDiscriminator field to given value.

### HasIpSpaceDiscriminator

`func (o *GetParentalcontrolSubscriberResponse) HasIpSpaceDiscriminator() bool`

HasIpSpaceDiscriminator returns a boolean if a field has been set.

### GetLocalId

`func (o *GetParentalcontrolSubscriberResponse) GetLocalId() string`

GetLocalId returns the LocalId field if non-nil, zero value otherwise.

### GetLocalIdOk

`func (o *GetParentalcontrolSubscriberResponse) GetLocalIdOk() (*string, bool)`

GetLocalIdOk returns a tuple with the LocalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalId

`func (o *GetParentalcontrolSubscriberResponse) SetLocalId(v string)`

SetLocalId sets LocalId field to given value.

### HasLocalId

`func (o *GetParentalcontrolSubscriberResponse) HasLocalId() bool`

HasLocalId returns a boolean if a field has been set.

### GetLocalIdRegexp

`func (o *GetParentalcontrolSubscriberResponse) GetLocalIdRegexp() string`

GetLocalIdRegexp returns the LocalIdRegexp field if non-nil, zero value otherwise.

### GetLocalIdRegexpOk

`func (o *GetParentalcontrolSubscriberResponse) GetLocalIdRegexpOk() (*string, bool)`

GetLocalIdRegexpOk returns a tuple with the LocalIdRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIdRegexp

`func (o *GetParentalcontrolSubscriberResponse) SetLocalIdRegexp(v string)`

SetLocalIdRegexp sets LocalIdRegexp field to given value.

### HasLocalIdRegexp

`func (o *GetParentalcontrolSubscriberResponse) HasLocalIdRegexp() bool`

HasLocalIdRegexp returns a boolean if a field has been set.

### GetLocalIdSubexpression

`func (o *GetParentalcontrolSubscriberResponse) GetLocalIdSubexpression() int64`

GetLocalIdSubexpression returns the LocalIdSubexpression field if non-nil, zero value otherwise.

### GetLocalIdSubexpressionOk

`func (o *GetParentalcontrolSubscriberResponse) GetLocalIdSubexpressionOk() (*int64, bool)`

GetLocalIdSubexpressionOk returns a tuple with the LocalIdSubexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIdSubexpression

`func (o *GetParentalcontrolSubscriberResponse) SetLocalIdSubexpression(v int64)`

SetLocalIdSubexpression sets LocalIdSubexpression field to given value.

### HasLocalIdSubexpression

`func (o *GetParentalcontrolSubscriberResponse) HasLocalIdSubexpression() bool`

HasLocalIdSubexpression returns a boolean if a field has been set.

### GetLogGuestLookups

`func (o *GetParentalcontrolSubscriberResponse) GetLogGuestLookups() bool`

GetLogGuestLookups returns the LogGuestLookups field if non-nil, zero value otherwise.

### GetLogGuestLookupsOk

`func (o *GetParentalcontrolSubscriberResponse) GetLogGuestLookupsOk() (*bool, bool)`

GetLogGuestLookupsOk returns a tuple with the LogGuestLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogGuestLookups

`func (o *GetParentalcontrolSubscriberResponse) SetLogGuestLookups(v bool)`

SetLogGuestLookups sets LogGuestLookups field to given value.

### HasLogGuestLookups

`func (o *GetParentalcontrolSubscriberResponse) HasLogGuestLookups() bool`

HasLogGuestLookups returns a boolean if a field has been set.

### GetNasContextInfo

`func (o *GetParentalcontrolSubscriberResponse) GetNasContextInfo() string`

GetNasContextInfo returns the NasContextInfo field if non-nil, zero value otherwise.

### GetNasContextInfoOk

`func (o *GetParentalcontrolSubscriberResponse) GetNasContextInfoOk() (*string, bool)`

GetNasContextInfoOk returns a tuple with the NasContextInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasContextInfo

`func (o *GetParentalcontrolSubscriberResponse) SetNasContextInfo(v string)`

SetNasContextInfo sets NasContextInfo field to given value.

### HasNasContextInfo

`func (o *GetParentalcontrolSubscriberResponse) HasNasContextInfo() bool`

HasNasContextInfo returns a boolean if a field has been set.

### GetPcZoneName

`func (o *GetParentalcontrolSubscriberResponse) GetPcZoneName() string`

GetPcZoneName returns the PcZoneName field if non-nil, zero value otherwise.

### GetPcZoneNameOk

`func (o *GetParentalcontrolSubscriberResponse) GetPcZoneNameOk() (*string, bool)`

GetPcZoneNameOk returns a tuple with the PcZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcZoneName

`func (o *GetParentalcontrolSubscriberResponse) SetPcZoneName(v string)`

SetPcZoneName sets PcZoneName field to given value.

### HasPcZoneName

`func (o *GetParentalcontrolSubscriberResponse) HasPcZoneName() bool`

HasPcZoneName returns a boolean if a field has been set.

### GetProxyPassword

`func (o *GetParentalcontrolSubscriberResponse) GetProxyPassword() string`

GetProxyPassword returns the ProxyPassword field if non-nil, zero value otherwise.

### GetProxyPasswordOk

`func (o *GetParentalcontrolSubscriberResponse) GetProxyPasswordOk() (*string, bool)`

GetProxyPasswordOk returns a tuple with the ProxyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPassword

`func (o *GetParentalcontrolSubscriberResponse) SetProxyPassword(v string)`

SetProxyPassword sets ProxyPassword field to given value.

### HasProxyPassword

`func (o *GetParentalcontrolSubscriberResponse) HasProxyPassword() bool`

HasProxyPassword returns a boolean if a field has been set.

### GetProxyUrl

`func (o *GetParentalcontrolSubscriberResponse) GetProxyUrl() string`

GetProxyUrl returns the ProxyUrl field if non-nil, zero value otherwise.

### GetProxyUrlOk

`func (o *GetParentalcontrolSubscriberResponse) GetProxyUrlOk() (*string, bool)`

GetProxyUrlOk returns a tuple with the ProxyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUrl

`func (o *GetParentalcontrolSubscriberResponse) SetProxyUrl(v string)`

SetProxyUrl sets ProxyUrl field to given value.

### HasProxyUrl

`func (o *GetParentalcontrolSubscriberResponse) HasProxyUrl() bool`

HasProxyUrl returns a boolean if a field has been set.

### GetProxyUsername

`func (o *GetParentalcontrolSubscriberResponse) GetProxyUsername() string`

GetProxyUsername returns the ProxyUsername field if non-nil, zero value otherwise.

### GetProxyUsernameOk

`func (o *GetParentalcontrolSubscriberResponse) GetProxyUsernameOk() (*string, bool)`

GetProxyUsernameOk returns a tuple with the ProxyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyUsername

`func (o *GetParentalcontrolSubscriberResponse) SetProxyUsername(v string)`

SetProxyUsername sets ProxyUsername field to given value.

### HasProxyUsername

`func (o *GetParentalcontrolSubscriberResponse) HasProxyUsername() bool`

HasProxyUsername returns a boolean if a field has been set.

### GetSubscriberId

`func (o *GetParentalcontrolSubscriberResponse) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *GetParentalcontrolSubscriberResponse) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *GetParentalcontrolSubscriberResponse) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *GetParentalcontrolSubscriberResponse) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetSubscriberIdRegexp

`func (o *GetParentalcontrolSubscriberResponse) GetSubscriberIdRegexp() string`

GetSubscriberIdRegexp returns the SubscriberIdRegexp field if non-nil, zero value otherwise.

### GetSubscriberIdRegexpOk

`func (o *GetParentalcontrolSubscriberResponse) GetSubscriberIdRegexpOk() (*string, bool)`

GetSubscriberIdRegexpOk returns a tuple with the SubscriberIdRegexp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdRegexp

`func (o *GetParentalcontrolSubscriberResponse) SetSubscriberIdRegexp(v string)`

SetSubscriberIdRegexp sets SubscriberIdRegexp field to given value.

### HasSubscriberIdRegexp

`func (o *GetParentalcontrolSubscriberResponse) HasSubscriberIdRegexp() bool`

HasSubscriberIdRegexp returns a boolean if a field has been set.

### GetSubscriberIdSubexpression

`func (o *GetParentalcontrolSubscriberResponse) GetSubscriberIdSubexpression() int64`

GetSubscriberIdSubexpression returns the SubscriberIdSubexpression field if non-nil, zero value otherwise.

### GetSubscriberIdSubexpressionOk

`func (o *GetParentalcontrolSubscriberResponse) GetSubscriberIdSubexpressionOk() (*int64, bool)`

GetSubscriberIdSubexpressionOk returns a tuple with the SubscriberIdSubexpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberIdSubexpression

`func (o *GetParentalcontrolSubscriberResponse) SetSubscriberIdSubexpression(v int64)`

SetSubscriberIdSubexpression sets SubscriberIdSubexpression field to given value.

### HasSubscriberIdSubexpression

`func (o *GetParentalcontrolSubscriberResponse) HasSubscriberIdSubexpression() bool`

HasSubscriberIdSubexpression returns a boolean if a field has been set.

### GetUuid

`func (o *GetParentalcontrolSubscriberResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetParentalcontrolSubscriberResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetParentalcontrolSubscriberResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetParentalcontrolSubscriberResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetZveloUpdateFailureInDays

`func (o *GetParentalcontrolSubscriberResponse) GetZveloUpdateFailureInDays() int64`

GetZveloUpdateFailureInDays returns the ZveloUpdateFailureInDays field if non-nil, zero value otherwise.

### GetZveloUpdateFailureInDaysOk

`func (o *GetParentalcontrolSubscriberResponse) GetZveloUpdateFailureInDaysOk() (*int64, bool)`

GetZveloUpdateFailureInDaysOk returns a tuple with the ZveloUpdateFailureInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZveloUpdateFailureInDays

`func (o *GetParentalcontrolSubscriberResponse) SetZveloUpdateFailureInDays(v int64)`

SetZveloUpdateFailureInDays sets ZveloUpdateFailureInDays field to given value.

### HasZveloUpdateFailureInDays

`func (o *GetParentalcontrolSubscriberResponse) HasZveloUpdateFailureInDays() bool`

HasZveloUpdateFailureInDays returns a boolean if a field has been set.

### GetResult

`func (o *GetParentalcontrolSubscriberResponse) GetResult() ParentalcontrolSubscriber`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetParentalcontrolSubscriberResponse) GetResultOk() (*ParentalcontrolSubscriber, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetParentalcontrolSubscriberResponse) SetResult(v ParentalcontrolSubscriber)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetParentalcontrolSubscriberResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


