# Opportunity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 100; | 
**ExpectedCloseDate** | Pointer to **time.Time** |  Required On Updates; | [optional] 
**Type** | Pointer to [**OpportunityTypeReference**](OpportunityTypeReference.md) |  | [optional] 
**Stage** | Pointer to [**OpportunityStageReference**](OpportunityStageReference.md) |  | [optional] 
**Status** | Pointer to [**OpportunityStatusReference**](OpportunityStatusReference.md) |  | [optional] 
**Priority** | Pointer to [**OpportunityPriorityReference**](OpportunityPriorityReference.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Probability** | Pointer to [**OpportunityProbabilityReference**](OpportunityProbabilityReference.md) |  | [optional] 
**Source** | Pointer to **string** |  Max length: 50; | [optional] 
**Rating** | Pointer to [**OpportunityRatingReference**](OpportunityRatingReference.md) |  | [optional] 
**Campaign** | Pointer to [**CampaignReference**](CampaignReference.md) |  | [optional] 
**PrimarySalesRep** | [**MemberReference**](MemberReference.md) |  | 
**SecondarySalesRep** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  Required On Updates; | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  Required On Updates; | [optional] 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Contact** | [**ContactReference**](ContactReference.md) |  | 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**CustomerPO** | Pointer to **string** |  Max length: 25; | [optional] 
**PipelineChangeDate** | Pointer to **time.Time** |  | [optional] 
**DateBecameLead** | Pointer to **time.Time** |  | [optional] 
**ClosedDate** | Pointer to **time.Time** |  | [optional] 
**ClosedBy** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**TotalSalesTax** | Pointer to **NullableFloat64** |  | [optional] 
**ShipToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ShipToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**ShipToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillToCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**BillToContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**BillToSite** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**CompanyLocationId** | Pointer to **NullableInt32** |  | [optional] 
**TechnicalContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewOpportunity

`func NewOpportunity(name string, primarySalesRep MemberReference, company CompanyReference, contact ContactReference, ) *Opportunity`

NewOpportunity instantiates a new Opportunity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityWithDefaults

`func NewOpportunityWithDefaults() *Opportunity`

NewOpportunityWithDefaults instantiates a new Opportunity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Opportunity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Opportunity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Opportunity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Opportunity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Opportunity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Opportunity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Opportunity) SetName(v string)`

SetName sets Name field to given value.


### GetExpectedCloseDate

`func (o *Opportunity) GetExpectedCloseDate() time.Time`

GetExpectedCloseDate returns the ExpectedCloseDate field if non-nil, zero value otherwise.

### GetExpectedCloseDateOk

`func (o *Opportunity) GetExpectedCloseDateOk() (*time.Time, bool)`

GetExpectedCloseDateOk returns a tuple with the ExpectedCloseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedCloseDate

`func (o *Opportunity) SetExpectedCloseDate(v time.Time)`

SetExpectedCloseDate sets ExpectedCloseDate field to given value.

### HasExpectedCloseDate

`func (o *Opportunity) HasExpectedCloseDate() bool`

HasExpectedCloseDate returns a boolean if a field has been set.

### GetType

`func (o *Opportunity) GetType() OpportunityTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Opportunity) GetTypeOk() (*OpportunityTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Opportunity) SetType(v OpportunityTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *Opportunity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStage

`func (o *Opportunity) GetStage() OpportunityStageReference`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Opportunity) GetStageOk() (*OpportunityStageReference, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Opportunity) SetStage(v OpportunityStageReference)`

SetStage sets Stage field to given value.

### HasStage

`func (o *Opportunity) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStatus

`func (o *Opportunity) GetStatus() OpportunityStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Opportunity) GetStatusOk() (*OpportunityStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Opportunity) SetStatus(v OpportunityStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Opportunity) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *Opportunity) GetPriority() OpportunityPriorityReference`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Opportunity) GetPriorityOk() (*OpportunityPriorityReference, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Opportunity) SetPriority(v OpportunityPriorityReference)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Opportunity) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetNotes

`func (o *Opportunity) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Opportunity) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Opportunity) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Opportunity) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetProbability

`func (o *Opportunity) GetProbability() OpportunityProbabilityReference`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *Opportunity) GetProbabilityOk() (*OpportunityProbabilityReference, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *Opportunity) SetProbability(v OpportunityProbabilityReference)`

SetProbability sets Probability field to given value.

### HasProbability

`func (o *Opportunity) HasProbability() bool`

HasProbability returns a boolean if a field has been set.

### GetSource

`func (o *Opportunity) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Opportunity) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Opportunity) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Opportunity) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRating

`func (o *Opportunity) GetRating() OpportunityRatingReference`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Opportunity) GetRatingOk() (*OpportunityRatingReference, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Opportunity) SetRating(v OpportunityRatingReference)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Opportunity) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetCampaign

`func (o *Opportunity) GetCampaign() CampaignReference`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *Opportunity) GetCampaignOk() (*CampaignReference, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *Opportunity) SetCampaign(v CampaignReference)`

SetCampaign sets Campaign field to given value.

### HasCampaign

`func (o *Opportunity) HasCampaign() bool`

HasCampaign returns a boolean if a field has been set.

### GetPrimarySalesRep

`func (o *Opportunity) GetPrimarySalesRep() MemberReference`

GetPrimarySalesRep returns the PrimarySalesRep field if non-nil, zero value otherwise.

### GetPrimarySalesRepOk

`func (o *Opportunity) GetPrimarySalesRepOk() (*MemberReference, bool)`

GetPrimarySalesRepOk returns a tuple with the PrimarySalesRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySalesRep

`func (o *Opportunity) SetPrimarySalesRep(v MemberReference)`

SetPrimarySalesRep sets PrimarySalesRep field to given value.


### GetSecondarySalesRep

`func (o *Opportunity) GetSecondarySalesRep() MemberReference`

GetSecondarySalesRep returns the SecondarySalesRep field if non-nil, zero value otherwise.

### GetSecondarySalesRepOk

`func (o *Opportunity) GetSecondarySalesRepOk() (*MemberReference, bool)`

GetSecondarySalesRepOk returns a tuple with the SecondarySalesRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondarySalesRep

`func (o *Opportunity) SetSecondarySalesRep(v MemberReference)`

SetSecondarySalesRep sets SecondarySalesRep field to given value.

### HasSecondarySalesRep

`func (o *Opportunity) HasSecondarySalesRep() bool`

HasSecondarySalesRep returns a boolean if a field has been set.

### GetLocationId

`func (o *Opportunity) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Opportunity) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Opportunity) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Opportunity) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *Opportunity) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *Opportunity) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetBusinessUnitId

`func (o *Opportunity) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *Opportunity) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *Opportunity) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *Opportunity) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *Opportunity) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *Opportunity) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetCompany

`func (o *Opportunity) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Opportunity) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Opportunity) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetContact

`func (o *Opportunity) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Opportunity) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Opportunity) SetContact(v ContactReference)`

SetContact sets Contact field to given value.


### GetSite

`func (o *Opportunity) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Opportunity) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Opportunity) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *Opportunity) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetCustomerPO

`func (o *Opportunity) GetCustomerPO() string`

GetCustomerPO returns the CustomerPO field if non-nil, zero value otherwise.

### GetCustomerPOOk

`func (o *Opportunity) GetCustomerPOOk() (*string, bool)`

GetCustomerPOOk returns a tuple with the CustomerPO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPO

`func (o *Opportunity) SetCustomerPO(v string)`

SetCustomerPO sets CustomerPO field to given value.

### HasCustomerPO

`func (o *Opportunity) HasCustomerPO() bool`

HasCustomerPO returns a boolean if a field has been set.

### GetPipelineChangeDate

`func (o *Opportunity) GetPipelineChangeDate() time.Time`

GetPipelineChangeDate returns the PipelineChangeDate field if non-nil, zero value otherwise.

### GetPipelineChangeDateOk

`func (o *Opportunity) GetPipelineChangeDateOk() (*time.Time, bool)`

GetPipelineChangeDateOk returns a tuple with the PipelineChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineChangeDate

`func (o *Opportunity) SetPipelineChangeDate(v time.Time)`

SetPipelineChangeDate sets PipelineChangeDate field to given value.

### HasPipelineChangeDate

`func (o *Opportunity) HasPipelineChangeDate() bool`

HasPipelineChangeDate returns a boolean if a field has been set.

### GetDateBecameLead

`func (o *Opportunity) GetDateBecameLead() time.Time`

GetDateBecameLead returns the DateBecameLead field if non-nil, zero value otherwise.

### GetDateBecameLeadOk

`func (o *Opportunity) GetDateBecameLeadOk() (*time.Time, bool)`

GetDateBecameLeadOk returns a tuple with the DateBecameLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateBecameLead

`func (o *Opportunity) SetDateBecameLead(v time.Time)`

SetDateBecameLead sets DateBecameLead field to given value.

### HasDateBecameLead

`func (o *Opportunity) HasDateBecameLead() bool`

HasDateBecameLead returns a boolean if a field has been set.

### GetClosedDate

`func (o *Opportunity) GetClosedDate() time.Time`

GetClosedDate returns the ClosedDate field if non-nil, zero value otherwise.

### GetClosedDateOk

`func (o *Opportunity) GetClosedDateOk() (*time.Time, bool)`

GetClosedDateOk returns a tuple with the ClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDate

`func (o *Opportunity) SetClosedDate(v time.Time)`

SetClosedDate sets ClosedDate field to given value.

### HasClosedDate

`func (o *Opportunity) HasClosedDate() bool`

HasClosedDate returns a boolean if a field has been set.

### GetClosedBy

`func (o *Opportunity) GetClosedBy() MemberReference`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *Opportunity) GetClosedByOk() (*MemberReference, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *Opportunity) SetClosedBy(v MemberReference)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *Opportunity) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetTotalSalesTax

`func (o *Opportunity) GetTotalSalesTax() float64`

GetTotalSalesTax returns the TotalSalesTax field if non-nil, zero value otherwise.

### GetTotalSalesTaxOk

`func (o *Opportunity) GetTotalSalesTaxOk() (*float64, bool)`

GetTotalSalesTaxOk returns a tuple with the TotalSalesTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSalesTax

`func (o *Opportunity) SetTotalSalesTax(v float64)`

SetTotalSalesTax sets TotalSalesTax field to given value.

### HasTotalSalesTax

`func (o *Opportunity) HasTotalSalesTax() bool`

HasTotalSalesTax returns a boolean if a field has been set.

### SetTotalSalesTaxNil

`func (o *Opportunity) SetTotalSalesTaxNil(b bool)`

 SetTotalSalesTaxNil sets the value for TotalSalesTax to be an explicit nil

### UnsetTotalSalesTax
`func (o *Opportunity) UnsetTotalSalesTax()`

UnsetTotalSalesTax ensures that no value is present for TotalSalesTax, not even an explicit nil
### GetShipToCompany

`func (o *Opportunity) GetShipToCompany() CompanyReference`

GetShipToCompany returns the ShipToCompany field if non-nil, zero value otherwise.

### GetShipToCompanyOk

`func (o *Opportunity) GetShipToCompanyOk() (*CompanyReference, bool)`

GetShipToCompanyOk returns a tuple with the ShipToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCompany

`func (o *Opportunity) SetShipToCompany(v CompanyReference)`

SetShipToCompany sets ShipToCompany field to given value.

### HasShipToCompany

`func (o *Opportunity) HasShipToCompany() bool`

HasShipToCompany returns a boolean if a field has been set.

### GetShipToContact

`func (o *Opportunity) GetShipToContact() ContactReference`

GetShipToContact returns the ShipToContact field if non-nil, zero value otherwise.

### GetShipToContactOk

`func (o *Opportunity) GetShipToContactOk() (*ContactReference, bool)`

GetShipToContactOk returns a tuple with the ShipToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToContact

`func (o *Opportunity) SetShipToContact(v ContactReference)`

SetShipToContact sets ShipToContact field to given value.

### HasShipToContact

`func (o *Opportunity) HasShipToContact() bool`

HasShipToContact returns a boolean if a field has been set.

### GetShipToSite

`func (o *Opportunity) GetShipToSite() SiteReference`

GetShipToSite returns the ShipToSite field if non-nil, zero value otherwise.

### GetShipToSiteOk

`func (o *Opportunity) GetShipToSiteOk() (*SiteReference, bool)`

GetShipToSiteOk returns a tuple with the ShipToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToSite

`func (o *Opportunity) SetShipToSite(v SiteReference)`

SetShipToSite sets ShipToSite field to given value.

### HasShipToSite

`func (o *Opportunity) HasShipToSite() bool`

HasShipToSite returns a boolean if a field has been set.

### GetBillToCompany

`func (o *Opportunity) GetBillToCompany() CompanyReference`

GetBillToCompany returns the BillToCompany field if non-nil, zero value otherwise.

### GetBillToCompanyOk

`func (o *Opportunity) GetBillToCompanyOk() (*CompanyReference, bool)`

GetBillToCompanyOk returns a tuple with the BillToCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToCompany

`func (o *Opportunity) SetBillToCompany(v CompanyReference)`

SetBillToCompany sets BillToCompany field to given value.

### HasBillToCompany

`func (o *Opportunity) HasBillToCompany() bool`

HasBillToCompany returns a boolean if a field has been set.

### GetBillToContact

`func (o *Opportunity) GetBillToContact() ContactReference`

GetBillToContact returns the BillToContact field if non-nil, zero value otherwise.

### GetBillToContactOk

`func (o *Opportunity) GetBillToContactOk() (*ContactReference, bool)`

GetBillToContactOk returns a tuple with the BillToContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToContact

`func (o *Opportunity) SetBillToContact(v ContactReference)`

SetBillToContact sets BillToContact field to given value.

### HasBillToContact

`func (o *Opportunity) HasBillToContact() bool`

HasBillToContact returns a boolean if a field has been set.

### GetBillToSite

`func (o *Opportunity) GetBillToSite() SiteReference`

GetBillToSite returns the BillToSite field if non-nil, zero value otherwise.

### GetBillToSiteOk

`func (o *Opportunity) GetBillToSiteOk() (*SiteReference, bool)`

GetBillToSiteOk returns a tuple with the BillToSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillToSite

`func (o *Opportunity) SetBillToSite(v SiteReference)`

SetBillToSite sets BillToSite field to given value.

### HasBillToSite

`func (o *Opportunity) HasBillToSite() bool`

HasBillToSite returns a boolean if a field has been set.

### GetBillingTerms

`func (o *Opportunity) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *Opportunity) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *Opportunity) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *Opportunity) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetTaxCode

`func (o *Opportunity) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Opportunity) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Opportunity) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Opportunity) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetCurrency

`func (o *Opportunity) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Opportunity) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Opportunity) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Opportunity) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCompanyLocationId

`func (o *Opportunity) GetCompanyLocationId() int32`

GetCompanyLocationId returns the CompanyLocationId field if non-nil, zero value otherwise.

### GetCompanyLocationIdOk

`func (o *Opportunity) GetCompanyLocationIdOk() (*int32, bool)`

GetCompanyLocationIdOk returns a tuple with the CompanyLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLocationId

`func (o *Opportunity) SetCompanyLocationId(v int32)`

SetCompanyLocationId sets CompanyLocationId field to given value.

### HasCompanyLocationId

`func (o *Opportunity) HasCompanyLocationId() bool`

HasCompanyLocationId returns a boolean if a field has been set.

### SetCompanyLocationIdNil

`func (o *Opportunity) SetCompanyLocationIdNil(b bool)`

 SetCompanyLocationIdNil sets the value for CompanyLocationId to be an explicit nil

### UnsetCompanyLocationId
`func (o *Opportunity) UnsetCompanyLocationId()`

UnsetCompanyLocationId ensures that no value is present for CompanyLocationId, not even an explicit nil
### GetTechnicalContact

`func (o *Opportunity) GetTechnicalContact() ContactReference`

GetTechnicalContact returns the TechnicalContact field if non-nil, zero value otherwise.

### GetTechnicalContactOk

`func (o *Opportunity) GetTechnicalContactOk() (*ContactReference, bool)`

GetTechnicalContactOk returns a tuple with the TechnicalContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContact

`func (o *Opportunity) SetTechnicalContact(v ContactReference)`

SetTechnicalContact sets TechnicalContact field to given value.

### HasTechnicalContact

`func (o *Opportunity) HasTechnicalContact() bool`

HasTechnicalContact returns a boolean if a field has been set.

### GetInfo

`func (o *Opportunity) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Opportunity) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Opportunity) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Opportunity) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *Opportunity) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Opportunity) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Opportunity) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Opportunity) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


