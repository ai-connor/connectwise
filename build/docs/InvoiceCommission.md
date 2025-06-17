# InvoiceCommission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Percent** | Pointer to **NullableFloat64** |  | [optional] 
**SplitPercent** | Pointer to **NullableFloat64** |  | [optional] 
**Adjustment** | Pointer to **NullableFloat64** |  | [optional] 
**NetAmount** | Pointer to **NullableFloat64** |  | [optional] 
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**Activity** | Pointer to [**ActivityReference**](ActivityReference.md) |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**SalesOrder** | Pointer to [**SalesOrderReference**](SalesOrderReference.md) |  | [optional] 
**AdjustedBy** | Pointer to **string** |  | [optional] 
**AdjustedDate** | Pointer to **string** |  | [optional] 
**AdjustmentReason** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInvoiceCommission

`func NewInvoiceCommission() *InvoiceCommission`

NewInvoiceCommission instantiates a new InvoiceCommission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCommissionWithDefaults

`func NewInvoiceCommissionWithDefaults() *InvoiceCommission`

NewInvoiceCommissionWithDefaults instantiates a new InvoiceCommission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceCommission) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceCommission) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceCommission) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceCommission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *InvoiceCommission) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *InvoiceCommission) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *InvoiceCommission) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *InvoiceCommission) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetPercent

`func (o *InvoiceCommission) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *InvoiceCommission) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *InvoiceCommission) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *InvoiceCommission) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### SetPercentNil

`func (o *InvoiceCommission) SetPercentNil(b bool)`

 SetPercentNil sets the value for Percent to be an explicit nil

### UnsetPercent
`func (o *InvoiceCommission) UnsetPercent()`

UnsetPercent ensures that no value is present for Percent, not even an explicit nil
### GetSplitPercent

`func (o *InvoiceCommission) GetSplitPercent() float64`

GetSplitPercent returns the SplitPercent field if non-nil, zero value otherwise.

### GetSplitPercentOk

`func (o *InvoiceCommission) GetSplitPercentOk() (*float64, bool)`

GetSplitPercentOk returns a tuple with the SplitPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitPercent

`func (o *InvoiceCommission) SetSplitPercent(v float64)`

SetSplitPercent sets SplitPercent field to given value.

### HasSplitPercent

`func (o *InvoiceCommission) HasSplitPercent() bool`

HasSplitPercent returns a boolean if a field has been set.

### SetSplitPercentNil

`func (o *InvoiceCommission) SetSplitPercentNil(b bool)`

 SetSplitPercentNil sets the value for SplitPercent to be an explicit nil

### UnsetSplitPercent
`func (o *InvoiceCommission) UnsetSplitPercent()`

UnsetSplitPercent ensures that no value is present for SplitPercent, not even an explicit nil
### GetAdjustment

`func (o *InvoiceCommission) GetAdjustment() float64`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *InvoiceCommission) GetAdjustmentOk() (*float64, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *InvoiceCommission) SetAdjustment(v float64)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *InvoiceCommission) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### SetAdjustmentNil

`func (o *InvoiceCommission) SetAdjustmentNil(b bool)`

 SetAdjustmentNil sets the value for Adjustment to be an explicit nil

### UnsetAdjustment
`func (o *InvoiceCommission) UnsetAdjustment()`

UnsetAdjustment ensures that no value is present for Adjustment, not even an explicit nil
### GetNetAmount

`func (o *InvoiceCommission) GetNetAmount() float64`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *InvoiceCommission) GetNetAmountOk() (*float64, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *InvoiceCommission) SetNetAmount(v float64)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *InvoiceCommission) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### SetNetAmountNil

`func (o *InvoiceCommission) SetNetAmountNil(b bool)`

 SetNetAmountNil sets the value for NetAmount to be an explicit nil

### UnsetNetAmount
`func (o *InvoiceCommission) UnsetNetAmount()`

UnsetNetAmount ensures that no value is present for NetAmount, not even an explicit nil
### GetAmount

`func (o *InvoiceCommission) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceCommission) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceCommission) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoiceCommission) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *InvoiceCommission) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *InvoiceCommission) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetInvoice

`func (o *InvoiceCommission) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *InvoiceCommission) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *InvoiceCommission) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *InvoiceCommission) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetOpportunity

`func (o *InvoiceCommission) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *InvoiceCommission) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *InvoiceCommission) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *InvoiceCommission) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetAgreement

`func (o *InvoiceCommission) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *InvoiceCommission) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *InvoiceCommission) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *InvoiceCommission) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetActivity

`func (o *InvoiceCommission) GetActivity() ActivityReference`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *InvoiceCommission) GetActivityOk() (*ActivityReference, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *InvoiceCommission) SetActivity(v ActivityReference)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *InvoiceCommission) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetTicket

`func (o *InvoiceCommission) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *InvoiceCommission) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *InvoiceCommission) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *InvoiceCommission) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetProject

`func (o *InvoiceCommission) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *InvoiceCommission) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *InvoiceCommission) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *InvoiceCommission) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetSalesOrder

`func (o *InvoiceCommission) GetSalesOrder() SalesOrderReference`

GetSalesOrder returns the SalesOrder field if non-nil, zero value otherwise.

### GetSalesOrderOk

`func (o *InvoiceCommission) GetSalesOrderOk() (*SalesOrderReference, bool)`

GetSalesOrderOk returns a tuple with the SalesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrder

`func (o *InvoiceCommission) SetSalesOrder(v SalesOrderReference)`

SetSalesOrder sets SalesOrder field to given value.

### HasSalesOrder

`func (o *InvoiceCommission) HasSalesOrder() bool`

HasSalesOrder returns a boolean if a field has been set.

### GetAdjustedBy

`func (o *InvoiceCommission) GetAdjustedBy() string`

GetAdjustedBy returns the AdjustedBy field if non-nil, zero value otherwise.

### GetAdjustedByOk

`func (o *InvoiceCommission) GetAdjustedByOk() (*string, bool)`

GetAdjustedByOk returns a tuple with the AdjustedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedBy

`func (o *InvoiceCommission) SetAdjustedBy(v string)`

SetAdjustedBy sets AdjustedBy field to given value.

### HasAdjustedBy

`func (o *InvoiceCommission) HasAdjustedBy() bool`

HasAdjustedBy returns a boolean if a field has been set.

### GetAdjustedDate

`func (o *InvoiceCommission) GetAdjustedDate() string`

GetAdjustedDate returns the AdjustedDate field if non-nil, zero value otherwise.

### GetAdjustedDateOk

`func (o *InvoiceCommission) GetAdjustedDateOk() (*string, bool)`

GetAdjustedDateOk returns a tuple with the AdjustedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedDate

`func (o *InvoiceCommission) SetAdjustedDate(v string)`

SetAdjustedDate sets AdjustedDate field to given value.

### HasAdjustedDate

`func (o *InvoiceCommission) HasAdjustedDate() bool`

HasAdjustedDate returns a boolean if a field has been set.

### GetAdjustmentReason

`func (o *InvoiceCommission) GetAdjustmentReason() string`

GetAdjustmentReason returns the AdjustmentReason field if non-nil, zero value otherwise.

### GetAdjustmentReasonOk

`func (o *InvoiceCommission) GetAdjustmentReasonOk() (*string, bool)`

GetAdjustmentReasonOk returns a tuple with the AdjustmentReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentReason

`func (o *InvoiceCommission) SetAdjustmentReason(v string)`

SetAdjustmentReason sets AdjustmentReason field to given value.

### HasAdjustmentReason

`func (o *InvoiceCommission) HasAdjustmentReason() bool`

HasAdjustmentReason returns a boolean if a field has been set.

### GetInfo

`func (o *InvoiceCommission) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InvoiceCommission) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InvoiceCommission) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InvoiceCommission) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


