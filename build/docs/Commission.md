# Commission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | [**MemberReference**](MemberReference.md) |  | 
**CommissionPercent** | Pointer to **NullableFloat64** |  | [optional] 
**DateStart** | Pointer to **time.Time** |  | [optional] 
**DateEnd** | Pointer to **time.Time** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**ServiceBoard** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Territory** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**BillingMethod** | Pointer to **NullableString** |  | [optional] 
**ServiceType** | Pointer to [**ServiceTypeReference**](ServiceTypeReference.md) |  | [optional] 
**ProjectBoard** | Pointer to [**ProjectBoardReference**](ProjectBoardReference.md) |  | [optional] 
**ProjectType** | Pointer to [**ProjectTypeReference**](ProjectTypeReference.md) |  | [optional] 
**AgreementType** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**NumberOfMonths** | Pointer to **NullableInt32** |  | [optional] 
**ProductCategory** | Pointer to [**ProductCategoryReference**](ProductCategoryReference.md) |  | [optional] 
**ProductSubCategory** | Pointer to [**ProductSubCategoryReference**](ProductSubCategoryReference.md) |  | [optional] 
**Item** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**CommissionBasis** | Pointer to **NullableString** |  | [optional] 
**InvoiceOption** | Pointer to **NullableString** |  | [optional] 
**ServicesFlag** | Pointer to **NullableBool** |  | [optional] 
**AgreementsFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**MyOpportunitiesFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCommission

`func NewCommission(member MemberReference, ) *Commission`

NewCommission instantiates a new Commission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionWithDefaults

`func NewCommissionWithDefaults() *Commission`

NewCommissionWithDefaults instantiates a new Commission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Commission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Commission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Commission) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Commission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *Commission) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Commission) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Commission) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetCommissionPercent

`func (o *Commission) GetCommissionPercent() float64`

GetCommissionPercent returns the CommissionPercent field if non-nil, zero value otherwise.

### GetCommissionPercentOk

`func (o *Commission) GetCommissionPercentOk() (*float64, bool)`

GetCommissionPercentOk returns a tuple with the CommissionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionPercent

`func (o *Commission) SetCommissionPercent(v float64)`

SetCommissionPercent sets CommissionPercent field to given value.

### HasCommissionPercent

`func (o *Commission) HasCommissionPercent() bool`

HasCommissionPercent returns a boolean if a field has been set.

### SetCommissionPercentNil

`func (o *Commission) SetCommissionPercentNil(b bool)`

 SetCommissionPercentNil sets the value for CommissionPercent to be an explicit nil

### UnsetCommissionPercent
`func (o *Commission) UnsetCommissionPercent()`

UnsetCommissionPercent ensures that no value is present for CommissionPercent, not even an explicit nil
### GetDateStart

`func (o *Commission) GetDateStart() time.Time`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *Commission) GetDateStartOk() (*time.Time, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *Commission) SetDateStart(v time.Time)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *Commission) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *Commission) GetDateEnd() time.Time`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *Commission) GetDateEndOk() (*time.Time, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *Commission) SetDateEnd(v time.Time)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *Commission) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### GetLocation

`func (o *Commission) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Commission) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Commission) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Commission) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *Commission) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Commission) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Commission) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Commission) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetCompany

`func (o *Commission) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Commission) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Commission) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Commission) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetSite

`func (o *Commission) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Commission) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Commission) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *Commission) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAgreement

`func (o *Commission) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *Commission) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *Commission) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *Commission) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetProject

`func (o *Commission) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Commission) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Commission) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *Commission) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetServiceBoard

`func (o *Commission) GetServiceBoard() BoardReference`

GetServiceBoard returns the ServiceBoard field if non-nil, zero value otherwise.

### GetServiceBoardOk

`func (o *Commission) GetServiceBoardOk() (*BoardReference, bool)`

GetServiceBoardOk returns a tuple with the ServiceBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoard

`func (o *Commission) SetServiceBoard(v BoardReference)`

SetServiceBoard sets ServiceBoard field to given value.

### HasServiceBoard

`func (o *Commission) HasServiceBoard() bool`

HasServiceBoard returns a boolean if a field has been set.

### GetTicket

`func (o *Commission) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *Commission) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *Commission) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *Commission) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetTerritory

`func (o *Commission) GetTerritory() SystemLocationReference`

GetTerritory returns the Territory field if non-nil, zero value otherwise.

### GetTerritoryOk

`func (o *Commission) GetTerritoryOk() (*SystemLocationReference, bool)`

GetTerritoryOk returns a tuple with the Territory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritory

`func (o *Commission) SetTerritory(v SystemLocationReference)`

SetTerritory sets Territory field to given value.

### HasTerritory

`func (o *Commission) HasTerritory() bool`

HasTerritory returns a boolean if a field has been set.

### GetBillingMethod

`func (o *Commission) GetBillingMethod() string`

GetBillingMethod returns the BillingMethod field if non-nil, zero value otherwise.

### GetBillingMethodOk

`func (o *Commission) GetBillingMethodOk() (*string, bool)`

GetBillingMethodOk returns a tuple with the BillingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMethod

`func (o *Commission) SetBillingMethod(v string)`

SetBillingMethod sets BillingMethod field to given value.

### HasBillingMethod

`func (o *Commission) HasBillingMethod() bool`

HasBillingMethod returns a boolean if a field has been set.

### SetBillingMethodNil

`func (o *Commission) SetBillingMethodNil(b bool)`

 SetBillingMethodNil sets the value for BillingMethod to be an explicit nil

### UnsetBillingMethod
`func (o *Commission) UnsetBillingMethod()`

UnsetBillingMethod ensures that no value is present for BillingMethod, not even an explicit nil
### GetServiceType

`func (o *Commission) GetServiceType() ServiceTypeReference`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Commission) GetServiceTypeOk() (*ServiceTypeReference, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Commission) SetServiceType(v ServiceTypeReference)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Commission) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetProjectBoard

`func (o *Commission) GetProjectBoard() ProjectBoardReference`

GetProjectBoard returns the ProjectBoard field if non-nil, zero value otherwise.

### GetProjectBoardOk

`func (o *Commission) GetProjectBoardOk() (*ProjectBoardReference, bool)`

GetProjectBoardOk returns a tuple with the ProjectBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectBoard

`func (o *Commission) SetProjectBoard(v ProjectBoardReference)`

SetProjectBoard sets ProjectBoard field to given value.

### HasProjectBoard

`func (o *Commission) HasProjectBoard() bool`

HasProjectBoard returns a boolean if a field has been set.

### GetProjectType

`func (o *Commission) GetProjectType() ProjectTypeReference`

GetProjectType returns the ProjectType field if non-nil, zero value otherwise.

### GetProjectTypeOk

`func (o *Commission) GetProjectTypeOk() (*ProjectTypeReference, bool)`

GetProjectTypeOk returns a tuple with the ProjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectType

`func (o *Commission) SetProjectType(v ProjectTypeReference)`

SetProjectType sets ProjectType field to given value.

### HasProjectType

`func (o *Commission) HasProjectType() bool`

HasProjectType returns a boolean if a field has been set.

### GetAgreementType

`func (o *Commission) GetAgreementType() AgreementTypeReference`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *Commission) GetAgreementTypeOk() (*AgreementTypeReference, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *Commission) SetAgreementType(v AgreementTypeReference)`

SetAgreementType sets AgreementType field to given value.

### HasAgreementType

`func (o *Commission) HasAgreementType() bool`

HasAgreementType returns a boolean if a field has been set.

### GetNumberOfMonths

`func (o *Commission) GetNumberOfMonths() int32`

GetNumberOfMonths returns the NumberOfMonths field if non-nil, zero value otherwise.

### GetNumberOfMonthsOk

`func (o *Commission) GetNumberOfMonthsOk() (*int32, bool)`

GetNumberOfMonthsOk returns a tuple with the NumberOfMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfMonths

`func (o *Commission) SetNumberOfMonths(v int32)`

SetNumberOfMonths sets NumberOfMonths field to given value.

### HasNumberOfMonths

`func (o *Commission) HasNumberOfMonths() bool`

HasNumberOfMonths returns a boolean if a field has been set.

### SetNumberOfMonthsNil

`func (o *Commission) SetNumberOfMonthsNil(b bool)`

 SetNumberOfMonthsNil sets the value for NumberOfMonths to be an explicit nil

### UnsetNumberOfMonths
`func (o *Commission) UnsetNumberOfMonths()`

UnsetNumberOfMonths ensures that no value is present for NumberOfMonths, not even an explicit nil
### GetProductCategory

`func (o *Commission) GetProductCategory() ProductCategoryReference`

GetProductCategory returns the ProductCategory field if non-nil, zero value otherwise.

### GetProductCategoryOk

`func (o *Commission) GetProductCategoryOk() (*ProductCategoryReference, bool)`

GetProductCategoryOk returns a tuple with the ProductCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCategory

`func (o *Commission) SetProductCategory(v ProductCategoryReference)`

SetProductCategory sets ProductCategory field to given value.

### HasProductCategory

`func (o *Commission) HasProductCategory() bool`

HasProductCategory returns a boolean if a field has been set.

### GetProductSubCategory

`func (o *Commission) GetProductSubCategory() ProductSubCategoryReference`

GetProductSubCategory returns the ProductSubCategory field if non-nil, zero value otherwise.

### GetProductSubCategoryOk

`func (o *Commission) GetProductSubCategoryOk() (*ProductSubCategoryReference, bool)`

GetProductSubCategoryOk returns a tuple with the ProductSubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubCategory

`func (o *Commission) SetProductSubCategory(v ProductSubCategoryReference)`

SetProductSubCategory sets ProductSubCategory field to given value.

### HasProductSubCategory

`func (o *Commission) HasProductSubCategory() bool`

HasProductSubCategory returns a boolean if a field has been set.

### GetItem

`func (o *Commission) GetItem() IvItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *Commission) GetItemOk() (*IvItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *Commission) SetItem(v IvItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *Commission) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetCommissionBasis

`func (o *Commission) GetCommissionBasis() string`

GetCommissionBasis returns the CommissionBasis field if non-nil, zero value otherwise.

### GetCommissionBasisOk

`func (o *Commission) GetCommissionBasisOk() (*string, bool)`

GetCommissionBasisOk returns a tuple with the CommissionBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionBasis

`func (o *Commission) SetCommissionBasis(v string)`

SetCommissionBasis sets CommissionBasis field to given value.

### HasCommissionBasis

`func (o *Commission) HasCommissionBasis() bool`

HasCommissionBasis returns a boolean if a field has been set.

### SetCommissionBasisNil

`func (o *Commission) SetCommissionBasisNil(b bool)`

 SetCommissionBasisNil sets the value for CommissionBasis to be an explicit nil

### UnsetCommissionBasis
`func (o *Commission) UnsetCommissionBasis()`

UnsetCommissionBasis ensures that no value is present for CommissionBasis, not even an explicit nil
### GetInvoiceOption

`func (o *Commission) GetInvoiceOption() string`

GetInvoiceOption returns the InvoiceOption field if non-nil, zero value otherwise.

### GetInvoiceOptionOk

`func (o *Commission) GetInvoiceOptionOk() (*string, bool)`

GetInvoiceOptionOk returns a tuple with the InvoiceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceOption

`func (o *Commission) SetInvoiceOption(v string)`

SetInvoiceOption sets InvoiceOption field to given value.

### HasInvoiceOption

`func (o *Commission) HasInvoiceOption() bool`

HasInvoiceOption returns a boolean if a field has been set.

### SetInvoiceOptionNil

`func (o *Commission) SetInvoiceOptionNil(b bool)`

 SetInvoiceOptionNil sets the value for InvoiceOption to be an explicit nil

### UnsetInvoiceOption
`func (o *Commission) UnsetInvoiceOption()`

UnsetInvoiceOption ensures that no value is present for InvoiceOption, not even an explicit nil
### GetServicesFlag

`func (o *Commission) GetServicesFlag() bool`

GetServicesFlag returns the ServicesFlag field if non-nil, zero value otherwise.

### GetServicesFlagOk

`func (o *Commission) GetServicesFlagOk() (*bool, bool)`

GetServicesFlagOk returns a tuple with the ServicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesFlag

`func (o *Commission) SetServicesFlag(v bool)`

SetServicesFlag sets ServicesFlag field to given value.

### HasServicesFlag

`func (o *Commission) HasServicesFlag() bool`

HasServicesFlag returns a boolean if a field has been set.

### SetServicesFlagNil

`func (o *Commission) SetServicesFlagNil(b bool)`

 SetServicesFlagNil sets the value for ServicesFlag to be an explicit nil

### UnsetServicesFlag
`func (o *Commission) UnsetServicesFlag()`

UnsetServicesFlag ensures that no value is present for ServicesFlag, not even an explicit nil
### GetAgreementsFlag

`func (o *Commission) GetAgreementsFlag() bool`

GetAgreementsFlag returns the AgreementsFlag field if non-nil, zero value otherwise.

### GetAgreementsFlagOk

`func (o *Commission) GetAgreementsFlagOk() (*bool, bool)`

GetAgreementsFlagOk returns a tuple with the AgreementsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementsFlag

`func (o *Commission) SetAgreementsFlag(v bool)`

SetAgreementsFlag sets AgreementsFlag field to given value.

### HasAgreementsFlag

`func (o *Commission) HasAgreementsFlag() bool`

HasAgreementsFlag returns a boolean if a field has been set.

### SetAgreementsFlagNil

`func (o *Commission) SetAgreementsFlagNil(b bool)`

 SetAgreementsFlagNil sets the value for AgreementsFlag to be an explicit nil

### UnsetAgreementsFlag
`func (o *Commission) UnsetAgreementsFlag()`

UnsetAgreementsFlag ensures that no value is present for AgreementsFlag, not even an explicit nil
### GetProductsFlag

`func (o *Commission) GetProductsFlag() bool`

GetProductsFlag returns the ProductsFlag field if non-nil, zero value otherwise.

### GetProductsFlagOk

`func (o *Commission) GetProductsFlagOk() (*bool, bool)`

GetProductsFlagOk returns a tuple with the ProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductsFlag

`func (o *Commission) SetProductsFlag(v bool)`

SetProductsFlag sets ProductsFlag field to given value.

### HasProductsFlag

`func (o *Commission) HasProductsFlag() bool`

HasProductsFlag returns a boolean if a field has been set.

### SetProductsFlagNil

`func (o *Commission) SetProductsFlagNil(b bool)`

 SetProductsFlagNil sets the value for ProductsFlag to be an explicit nil

### UnsetProductsFlag
`func (o *Commission) UnsetProductsFlag()`

UnsetProductsFlag ensures that no value is present for ProductsFlag, not even an explicit nil
### GetMyOpportunitiesFlag

`func (o *Commission) GetMyOpportunitiesFlag() bool`

GetMyOpportunitiesFlag returns the MyOpportunitiesFlag field if non-nil, zero value otherwise.

### GetMyOpportunitiesFlagOk

`func (o *Commission) GetMyOpportunitiesFlagOk() (*bool, bool)`

GetMyOpportunitiesFlagOk returns a tuple with the MyOpportunitiesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyOpportunitiesFlag

`func (o *Commission) SetMyOpportunitiesFlag(v bool)`

SetMyOpportunitiesFlag sets MyOpportunitiesFlag field to given value.

### HasMyOpportunitiesFlag

`func (o *Commission) HasMyOpportunitiesFlag() bool`

HasMyOpportunitiesFlag returns a boolean if a field has been set.

### SetMyOpportunitiesFlagNil

`func (o *Commission) SetMyOpportunitiesFlagNil(b bool)`

 SetMyOpportunitiesFlagNil sets the value for MyOpportunitiesFlag to be an explicit nil

### UnsetMyOpportunitiesFlag
`func (o *Commission) UnsetMyOpportunitiesFlag()`

UnsetMyOpportunitiesFlag ensures that no value is present for MyOpportunitiesFlag, not even an explicit nil
### GetInfo

`func (o *Commission) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Commission) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Commission) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Commission) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


