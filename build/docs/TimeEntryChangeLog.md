# TimeEntryChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Time Entry Change Log ID. | [optional] 
**PartnerId** | Pointer to **string** | Partner ID. | [optional] 
**ProductInstanceId** | Pointer to **string** | Product Instance ID. | [optional] 
**Action** | Pointer to **string** | Action. | [optional] 
**ActivitySubject** | Pointer to **string** | Activity Subject. | [optional] 
**ActualUtilizedHrs** | Pointer to **NullableFloat64** | Actual Utilized Hours. | [optional] 
**AgreementAdjustmentFirm** | Pointer to **NullableFloat64** | Agreement Adjustment Firm. | [optional] 
**AgreementAdjustmentTotal** | Pointer to **NullableFloat64** | Agreement Adjustment Total. | [optional] 
**AgreementAmountCovered** | Pointer to **NullableFloat64** | Agreement Amount Covered. | [optional] 
**AgreementHoursCovered** | Pointer to **NullableFloat64** | Agreement Hours Covered. | [optional] 
**BillableAmount** | Pointer to **NullableFloat64** | Billable Amount. | [optional] 
**BillableFlag** | Pointer to **NullableBool** | Billable Flag. | [optional] 
**BillableHours** | Pointer to **NullableFloat64** | Billable Hours. | [optional] 
**BillableUtilizedHours** | Pointer to **NullableFloat64** | Billable Utilized Hours. | [optional] 
**MemberDailyCapacity** | Pointer to **NullableFloat64** | Member Daily Capacity. | [optional] 
**BillableOption** | Pointer to **NullableString** | Billable Option. | [optional] 
**BusinessGroup** | Pointer to **string** | Business Group. | [optional] 
**LocationName** | Pointer to **string** | Location Name. | [optional] 
**ChargeCode** | Pointer to **string** | Charge Code. | [optional] 
**ChargeTo** | Pointer to **string** | Charge To. | [optional] 
**ChargeToType** | Pointer to **NullableString** | Charge To Type. | [optional] 
**ChargeToRecId** | Pointer to **NullableInt32** | Charge To Record ID. | [optional] 
**CompanyAndAgreement** | Pointer to **string** | Company and Agreement. | [optional] 
**CompanyName** | Pointer to **string** | Company Name. | [optional] 
**TimeStart** | Pointer to **string** | Time Start. | [optional] 
**TimeStartUtc** | Pointer to **string** | Time Start UTC. | [optional] 
**TimeEnd** | Pointer to **string** | Time End. | [optional] 
**TimeEndUtc** | Pointer to **string** | Time End UTC. | [optional] 
**DateStart** | Pointer to **string** | Date Start. | [optional] 
**DateInvoice** | Pointer to **string** | Date Invoice. | [optional] 
**FirstName** | Pointer to **string** | First Name. | [optional] 
**HourlyCost** | Pointer to **string** | Hourly Cost. | [optional] 
**HourlyCostDecimal** | Pointer to **NullableFloat64** | Hourly Cost in Decimal. | [optional] 
**HourlyRate** | Pointer to **NullableFloat64** | Hourly Rate. | [optional] 
**HoursActual** | Pointer to **NullableFloat64** | Actual Hours. | [optional] 
**InternalNote** | Pointer to **string** | Internal Note. | [optional] 
**InvoiceAdjustmentFirm** | Pointer to **NullableFloat64** | Invoice Adjustment Firm. | [optional] 
**InvoiceAdjustmentTotal** | Pointer to **NullableFloat64** | Invoice Adjustment Total. | [optional] 
**InvoiceFlag** | Pointer to **bool** | Invoice Flag. | [optional] 
**InvoiceNumber** | Pointer to **string** | Invoice Number. | [optional] 
**InvoiceReady** | Pointer to **bool** | Invoice Ready status. | [optional] 
**LastName** | Pointer to **string** | Last Name. | [optional] 
**MemberId** | Pointer to **string** | Member ID. | [optional] 
**NonBillableAmt** | Pointer to **NullableFloat64** | Non-Billable Amount. | [optional] 
**NonBillableHrs** | Pointer to **NullableFloat64** | Non-Billable Hours. | [optional] 
**Notes** | Pointer to **string** | Notes. | [optional] 
**OpportunityRecId** | Pointer to **NullableInt32** | Opportunity Record ID. | [optional] 
**OptionId** | Pointer to **string** | Option ID. | [optional] 
**ProjectActivity** | Pointer to **string** | Project Activity. | [optional] 
**ProjectName** | Pointer to **string** | Project Name. | [optional] 
**ProjectPhase** | Pointer to **string** | Project Phase. | [optional] 
**ServiceRequestStatus** | Pointer to **string** | Service Request Status. | [optional] 
**ServiceRequestSummary** | Pointer to **string** | Service Request Summary. | [optional] 
**Territory** | Pointer to **string** | Territory. | [optional] 
**TimeRecId** | Pointer to **int32** | Time Record ID. | [optional] 
**TimeStatus** | Pointer to **string** | Time Status. | [optional] 
**UtilizationFlag** | Pointer to **NullableBool** | Utilization Flag. | [optional] 
**CompanyType** | Pointer to **string** | Company Type. | [optional] 
**TicketCurrentBoard** | Pointer to **string** | Current Board of the Ticket. | [optional] 
**TicketType** | Pointer to **string** | Type of the Ticket. | [optional] 
**TicketSubtype** | Pointer to **string** | Subtype of the Ticket. | [optional] 
**AgreementType** | Pointer to **string** | Type of the Agreement. | [optional] 
**BillingStatus** | Pointer to **string** | Billing Status. | [optional] 
**ProcessingStatus** | Pointer to **string** | Processing Status. | [optional] 
**Invoicedhours** | Pointer to **NullableFloat64** | Invoiced Hours. | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**Invoice** | Pointer to [**InvoiceReference**](InvoiceReference.md) |  | [optional] 
**Ticket** | Pointer to [**TicketReference**](TicketReference.md) |  | [optional] 
**Project** | Pointer to [**ProjectReference**](ProjectReference.md) |  | [optional] 
**Phase** | Pointer to [**ProjectPhaseReference**](ProjectPhaseReference.md) |  | [optional] 
**Activity** | Pointer to [**ActivityReference**](ActivityReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimeEntryChangeLog

`func NewTimeEntryChangeLog() *TimeEntryChangeLog`

NewTimeEntryChangeLog instantiates a new TimeEntryChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeEntryChangeLogWithDefaults

`func NewTimeEntryChangeLogWithDefaults() *TimeEntryChangeLog`

NewTimeEntryChangeLogWithDefaults instantiates a new TimeEntryChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeEntryChangeLog) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeEntryChangeLog) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeEntryChangeLog) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeEntryChangeLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartnerId

`func (o *TimeEntryChangeLog) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *TimeEntryChangeLog) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *TimeEntryChangeLog) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *TimeEntryChangeLog) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetProductInstanceId

`func (o *TimeEntryChangeLog) GetProductInstanceId() string`

GetProductInstanceId returns the ProductInstanceId field if non-nil, zero value otherwise.

### GetProductInstanceIdOk

`func (o *TimeEntryChangeLog) GetProductInstanceIdOk() (*string, bool)`

GetProductInstanceIdOk returns a tuple with the ProductInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInstanceId

`func (o *TimeEntryChangeLog) SetProductInstanceId(v string)`

SetProductInstanceId sets ProductInstanceId field to given value.

### HasProductInstanceId

`func (o *TimeEntryChangeLog) HasProductInstanceId() bool`

HasProductInstanceId returns a boolean if a field has been set.

### GetAction

`func (o *TimeEntryChangeLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TimeEntryChangeLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TimeEntryChangeLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TimeEntryChangeLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActivitySubject

`func (o *TimeEntryChangeLog) GetActivitySubject() string`

GetActivitySubject returns the ActivitySubject field if non-nil, zero value otherwise.

### GetActivitySubjectOk

`func (o *TimeEntryChangeLog) GetActivitySubjectOk() (*string, bool)`

GetActivitySubjectOk returns a tuple with the ActivitySubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitySubject

`func (o *TimeEntryChangeLog) SetActivitySubject(v string)`

SetActivitySubject sets ActivitySubject field to given value.

### HasActivitySubject

`func (o *TimeEntryChangeLog) HasActivitySubject() bool`

HasActivitySubject returns a boolean if a field has been set.

### GetActualUtilizedHrs

`func (o *TimeEntryChangeLog) GetActualUtilizedHrs() float64`

GetActualUtilizedHrs returns the ActualUtilizedHrs field if non-nil, zero value otherwise.

### GetActualUtilizedHrsOk

`func (o *TimeEntryChangeLog) GetActualUtilizedHrsOk() (*float64, bool)`

GetActualUtilizedHrsOk returns a tuple with the ActualUtilizedHrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualUtilizedHrs

`func (o *TimeEntryChangeLog) SetActualUtilizedHrs(v float64)`

SetActualUtilizedHrs sets ActualUtilizedHrs field to given value.

### HasActualUtilizedHrs

`func (o *TimeEntryChangeLog) HasActualUtilizedHrs() bool`

HasActualUtilizedHrs returns a boolean if a field has been set.

### SetActualUtilizedHrsNil

`func (o *TimeEntryChangeLog) SetActualUtilizedHrsNil(b bool)`

 SetActualUtilizedHrsNil sets the value for ActualUtilizedHrs to be an explicit nil

### UnsetActualUtilizedHrs
`func (o *TimeEntryChangeLog) UnsetActualUtilizedHrs()`

UnsetActualUtilizedHrs ensures that no value is present for ActualUtilizedHrs, not even an explicit nil
### GetAgreementAdjustmentFirm

`func (o *TimeEntryChangeLog) GetAgreementAdjustmentFirm() float64`

GetAgreementAdjustmentFirm returns the AgreementAdjustmentFirm field if non-nil, zero value otherwise.

### GetAgreementAdjustmentFirmOk

`func (o *TimeEntryChangeLog) GetAgreementAdjustmentFirmOk() (*float64, bool)`

GetAgreementAdjustmentFirmOk returns a tuple with the AgreementAdjustmentFirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAdjustmentFirm

`func (o *TimeEntryChangeLog) SetAgreementAdjustmentFirm(v float64)`

SetAgreementAdjustmentFirm sets AgreementAdjustmentFirm field to given value.

### HasAgreementAdjustmentFirm

`func (o *TimeEntryChangeLog) HasAgreementAdjustmentFirm() bool`

HasAgreementAdjustmentFirm returns a boolean if a field has been set.

### SetAgreementAdjustmentFirmNil

`func (o *TimeEntryChangeLog) SetAgreementAdjustmentFirmNil(b bool)`

 SetAgreementAdjustmentFirmNil sets the value for AgreementAdjustmentFirm to be an explicit nil

### UnsetAgreementAdjustmentFirm
`func (o *TimeEntryChangeLog) UnsetAgreementAdjustmentFirm()`

UnsetAgreementAdjustmentFirm ensures that no value is present for AgreementAdjustmentFirm, not even an explicit nil
### GetAgreementAdjustmentTotal

`func (o *TimeEntryChangeLog) GetAgreementAdjustmentTotal() float64`

GetAgreementAdjustmentTotal returns the AgreementAdjustmentTotal field if non-nil, zero value otherwise.

### GetAgreementAdjustmentTotalOk

`func (o *TimeEntryChangeLog) GetAgreementAdjustmentTotalOk() (*float64, bool)`

GetAgreementAdjustmentTotalOk returns a tuple with the AgreementAdjustmentTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAdjustmentTotal

`func (o *TimeEntryChangeLog) SetAgreementAdjustmentTotal(v float64)`

SetAgreementAdjustmentTotal sets AgreementAdjustmentTotal field to given value.

### HasAgreementAdjustmentTotal

`func (o *TimeEntryChangeLog) HasAgreementAdjustmentTotal() bool`

HasAgreementAdjustmentTotal returns a boolean if a field has been set.

### SetAgreementAdjustmentTotalNil

`func (o *TimeEntryChangeLog) SetAgreementAdjustmentTotalNil(b bool)`

 SetAgreementAdjustmentTotalNil sets the value for AgreementAdjustmentTotal to be an explicit nil

### UnsetAgreementAdjustmentTotal
`func (o *TimeEntryChangeLog) UnsetAgreementAdjustmentTotal()`

UnsetAgreementAdjustmentTotal ensures that no value is present for AgreementAdjustmentTotal, not even an explicit nil
### GetAgreementAmountCovered

`func (o *TimeEntryChangeLog) GetAgreementAmountCovered() float64`

GetAgreementAmountCovered returns the AgreementAmountCovered field if non-nil, zero value otherwise.

### GetAgreementAmountCoveredOk

`func (o *TimeEntryChangeLog) GetAgreementAmountCoveredOk() (*float64, bool)`

GetAgreementAmountCoveredOk returns a tuple with the AgreementAmountCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementAmountCovered

`func (o *TimeEntryChangeLog) SetAgreementAmountCovered(v float64)`

SetAgreementAmountCovered sets AgreementAmountCovered field to given value.

### HasAgreementAmountCovered

`func (o *TimeEntryChangeLog) HasAgreementAmountCovered() bool`

HasAgreementAmountCovered returns a boolean if a field has been set.

### SetAgreementAmountCoveredNil

`func (o *TimeEntryChangeLog) SetAgreementAmountCoveredNil(b bool)`

 SetAgreementAmountCoveredNil sets the value for AgreementAmountCovered to be an explicit nil

### UnsetAgreementAmountCovered
`func (o *TimeEntryChangeLog) UnsetAgreementAmountCovered()`

UnsetAgreementAmountCovered ensures that no value is present for AgreementAmountCovered, not even an explicit nil
### GetAgreementHoursCovered

`func (o *TimeEntryChangeLog) GetAgreementHoursCovered() float64`

GetAgreementHoursCovered returns the AgreementHoursCovered field if non-nil, zero value otherwise.

### GetAgreementHoursCoveredOk

`func (o *TimeEntryChangeLog) GetAgreementHoursCoveredOk() (*float64, bool)`

GetAgreementHoursCoveredOk returns a tuple with the AgreementHoursCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementHoursCovered

`func (o *TimeEntryChangeLog) SetAgreementHoursCovered(v float64)`

SetAgreementHoursCovered sets AgreementHoursCovered field to given value.

### HasAgreementHoursCovered

`func (o *TimeEntryChangeLog) HasAgreementHoursCovered() bool`

HasAgreementHoursCovered returns a boolean if a field has been set.

### SetAgreementHoursCoveredNil

`func (o *TimeEntryChangeLog) SetAgreementHoursCoveredNil(b bool)`

 SetAgreementHoursCoveredNil sets the value for AgreementHoursCovered to be an explicit nil

### UnsetAgreementHoursCovered
`func (o *TimeEntryChangeLog) UnsetAgreementHoursCovered()`

UnsetAgreementHoursCovered ensures that no value is present for AgreementHoursCovered, not even an explicit nil
### GetBillableAmount

`func (o *TimeEntryChangeLog) GetBillableAmount() float64`

GetBillableAmount returns the BillableAmount field if non-nil, zero value otherwise.

### GetBillableAmountOk

`func (o *TimeEntryChangeLog) GetBillableAmountOk() (*float64, bool)`

GetBillableAmountOk returns a tuple with the BillableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableAmount

`func (o *TimeEntryChangeLog) SetBillableAmount(v float64)`

SetBillableAmount sets BillableAmount field to given value.

### HasBillableAmount

`func (o *TimeEntryChangeLog) HasBillableAmount() bool`

HasBillableAmount returns a boolean if a field has been set.

### SetBillableAmountNil

`func (o *TimeEntryChangeLog) SetBillableAmountNil(b bool)`

 SetBillableAmountNil sets the value for BillableAmount to be an explicit nil

### UnsetBillableAmount
`func (o *TimeEntryChangeLog) UnsetBillableAmount()`

UnsetBillableAmount ensures that no value is present for BillableAmount, not even an explicit nil
### GetBillableFlag

`func (o *TimeEntryChangeLog) GetBillableFlag() bool`

GetBillableFlag returns the BillableFlag field if non-nil, zero value otherwise.

### GetBillableFlagOk

`func (o *TimeEntryChangeLog) GetBillableFlagOk() (*bool, bool)`

GetBillableFlagOk returns a tuple with the BillableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableFlag

`func (o *TimeEntryChangeLog) SetBillableFlag(v bool)`

SetBillableFlag sets BillableFlag field to given value.

### HasBillableFlag

`func (o *TimeEntryChangeLog) HasBillableFlag() bool`

HasBillableFlag returns a boolean if a field has been set.

### SetBillableFlagNil

`func (o *TimeEntryChangeLog) SetBillableFlagNil(b bool)`

 SetBillableFlagNil sets the value for BillableFlag to be an explicit nil

### UnsetBillableFlag
`func (o *TimeEntryChangeLog) UnsetBillableFlag()`

UnsetBillableFlag ensures that no value is present for BillableFlag, not even an explicit nil
### GetBillableHours

`func (o *TimeEntryChangeLog) GetBillableHours() float64`

GetBillableHours returns the BillableHours field if non-nil, zero value otherwise.

### GetBillableHoursOk

`func (o *TimeEntryChangeLog) GetBillableHoursOk() (*float64, bool)`

GetBillableHoursOk returns a tuple with the BillableHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableHours

`func (o *TimeEntryChangeLog) SetBillableHours(v float64)`

SetBillableHours sets BillableHours field to given value.

### HasBillableHours

`func (o *TimeEntryChangeLog) HasBillableHours() bool`

HasBillableHours returns a boolean if a field has been set.

### SetBillableHoursNil

`func (o *TimeEntryChangeLog) SetBillableHoursNil(b bool)`

 SetBillableHoursNil sets the value for BillableHours to be an explicit nil

### UnsetBillableHours
`func (o *TimeEntryChangeLog) UnsetBillableHours()`

UnsetBillableHours ensures that no value is present for BillableHours, not even an explicit nil
### GetBillableUtilizedHours

`func (o *TimeEntryChangeLog) GetBillableUtilizedHours() float64`

GetBillableUtilizedHours returns the BillableUtilizedHours field if non-nil, zero value otherwise.

### GetBillableUtilizedHoursOk

`func (o *TimeEntryChangeLog) GetBillableUtilizedHoursOk() (*float64, bool)`

GetBillableUtilizedHoursOk returns a tuple with the BillableUtilizedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableUtilizedHours

`func (o *TimeEntryChangeLog) SetBillableUtilizedHours(v float64)`

SetBillableUtilizedHours sets BillableUtilizedHours field to given value.

### HasBillableUtilizedHours

`func (o *TimeEntryChangeLog) HasBillableUtilizedHours() bool`

HasBillableUtilizedHours returns a boolean if a field has been set.

### SetBillableUtilizedHoursNil

`func (o *TimeEntryChangeLog) SetBillableUtilizedHoursNil(b bool)`

 SetBillableUtilizedHoursNil sets the value for BillableUtilizedHours to be an explicit nil

### UnsetBillableUtilizedHours
`func (o *TimeEntryChangeLog) UnsetBillableUtilizedHours()`

UnsetBillableUtilizedHours ensures that no value is present for BillableUtilizedHours, not even an explicit nil
### GetMemberDailyCapacity

`func (o *TimeEntryChangeLog) GetMemberDailyCapacity() float64`

GetMemberDailyCapacity returns the MemberDailyCapacity field if non-nil, zero value otherwise.

### GetMemberDailyCapacityOk

`func (o *TimeEntryChangeLog) GetMemberDailyCapacityOk() (*float64, bool)`

GetMemberDailyCapacityOk returns a tuple with the MemberDailyCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberDailyCapacity

`func (o *TimeEntryChangeLog) SetMemberDailyCapacity(v float64)`

SetMemberDailyCapacity sets MemberDailyCapacity field to given value.

### HasMemberDailyCapacity

`func (o *TimeEntryChangeLog) HasMemberDailyCapacity() bool`

HasMemberDailyCapacity returns a boolean if a field has been set.

### SetMemberDailyCapacityNil

`func (o *TimeEntryChangeLog) SetMemberDailyCapacityNil(b bool)`

 SetMemberDailyCapacityNil sets the value for MemberDailyCapacity to be an explicit nil

### UnsetMemberDailyCapacity
`func (o *TimeEntryChangeLog) UnsetMemberDailyCapacity()`

UnsetMemberDailyCapacity ensures that no value is present for MemberDailyCapacity, not even an explicit nil
### GetBillableOption

`func (o *TimeEntryChangeLog) GetBillableOption() string`

GetBillableOption returns the BillableOption field if non-nil, zero value otherwise.

### GetBillableOptionOk

`func (o *TimeEntryChangeLog) GetBillableOptionOk() (*string, bool)`

GetBillableOptionOk returns a tuple with the BillableOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOption

`func (o *TimeEntryChangeLog) SetBillableOption(v string)`

SetBillableOption sets BillableOption field to given value.

### HasBillableOption

`func (o *TimeEntryChangeLog) HasBillableOption() bool`

HasBillableOption returns a boolean if a field has been set.

### SetBillableOptionNil

`func (o *TimeEntryChangeLog) SetBillableOptionNil(b bool)`

 SetBillableOptionNil sets the value for BillableOption to be an explicit nil

### UnsetBillableOption
`func (o *TimeEntryChangeLog) UnsetBillableOption()`

UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
### GetBusinessGroup

`func (o *TimeEntryChangeLog) GetBusinessGroup() string`

GetBusinessGroup returns the BusinessGroup field if non-nil, zero value otherwise.

### GetBusinessGroupOk

`func (o *TimeEntryChangeLog) GetBusinessGroupOk() (*string, bool)`

GetBusinessGroupOk returns a tuple with the BusinessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessGroup

`func (o *TimeEntryChangeLog) SetBusinessGroup(v string)`

SetBusinessGroup sets BusinessGroup field to given value.

### HasBusinessGroup

`func (o *TimeEntryChangeLog) HasBusinessGroup() bool`

HasBusinessGroup returns a boolean if a field has been set.

### GetLocationName

`func (o *TimeEntryChangeLog) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *TimeEntryChangeLog) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *TimeEntryChangeLog) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *TimeEntryChangeLog) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetChargeCode

`func (o *TimeEntryChangeLog) GetChargeCode() string`

GetChargeCode returns the ChargeCode field if non-nil, zero value otherwise.

### GetChargeCodeOk

`func (o *TimeEntryChangeLog) GetChargeCodeOk() (*string, bool)`

GetChargeCodeOk returns a tuple with the ChargeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCode

`func (o *TimeEntryChangeLog) SetChargeCode(v string)`

SetChargeCode sets ChargeCode field to given value.

### HasChargeCode

`func (o *TimeEntryChangeLog) HasChargeCode() bool`

HasChargeCode returns a boolean if a field has been set.

### GetChargeTo

`func (o *TimeEntryChangeLog) GetChargeTo() string`

GetChargeTo returns the ChargeTo field if non-nil, zero value otherwise.

### GetChargeToOk

`func (o *TimeEntryChangeLog) GetChargeToOk() (*string, bool)`

GetChargeToOk returns a tuple with the ChargeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeTo

`func (o *TimeEntryChangeLog) SetChargeTo(v string)`

SetChargeTo sets ChargeTo field to given value.

### HasChargeTo

`func (o *TimeEntryChangeLog) HasChargeTo() bool`

HasChargeTo returns a boolean if a field has been set.

### GetChargeToType

`func (o *TimeEntryChangeLog) GetChargeToType() string`

GetChargeToType returns the ChargeToType field if non-nil, zero value otherwise.

### GetChargeToTypeOk

`func (o *TimeEntryChangeLog) GetChargeToTypeOk() (*string, bool)`

GetChargeToTypeOk returns a tuple with the ChargeToType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeToType

`func (o *TimeEntryChangeLog) SetChargeToType(v string)`

SetChargeToType sets ChargeToType field to given value.

### HasChargeToType

`func (o *TimeEntryChangeLog) HasChargeToType() bool`

HasChargeToType returns a boolean if a field has been set.

### SetChargeToTypeNil

`func (o *TimeEntryChangeLog) SetChargeToTypeNil(b bool)`

 SetChargeToTypeNil sets the value for ChargeToType to be an explicit nil

### UnsetChargeToType
`func (o *TimeEntryChangeLog) UnsetChargeToType()`

UnsetChargeToType ensures that no value is present for ChargeToType, not even an explicit nil
### GetChargeToRecId

`func (o *TimeEntryChangeLog) GetChargeToRecId() int32`

GetChargeToRecId returns the ChargeToRecId field if non-nil, zero value otherwise.

### GetChargeToRecIdOk

`func (o *TimeEntryChangeLog) GetChargeToRecIdOk() (*int32, bool)`

GetChargeToRecIdOk returns a tuple with the ChargeToRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeToRecId

`func (o *TimeEntryChangeLog) SetChargeToRecId(v int32)`

SetChargeToRecId sets ChargeToRecId field to given value.

### HasChargeToRecId

`func (o *TimeEntryChangeLog) HasChargeToRecId() bool`

HasChargeToRecId returns a boolean if a field has been set.

### SetChargeToRecIdNil

`func (o *TimeEntryChangeLog) SetChargeToRecIdNil(b bool)`

 SetChargeToRecIdNil sets the value for ChargeToRecId to be an explicit nil

### UnsetChargeToRecId
`func (o *TimeEntryChangeLog) UnsetChargeToRecId()`

UnsetChargeToRecId ensures that no value is present for ChargeToRecId, not even an explicit nil
### GetCompanyAndAgreement

`func (o *TimeEntryChangeLog) GetCompanyAndAgreement() string`

GetCompanyAndAgreement returns the CompanyAndAgreement field if non-nil, zero value otherwise.

### GetCompanyAndAgreementOk

`func (o *TimeEntryChangeLog) GetCompanyAndAgreementOk() (*string, bool)`

GetCompanyAndAgreementOk returns a tuple with the CompanyAndAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAndAgreement

`func (o *TimeEntryChangeLog) SetCompanyAndAgreement(v string)`

SetCompanyAndAgreement sets CompanyAndAgreement field to given value.

### HasCompanyAndAgreement

`func (o *TimeEntryChangeLog) HasCompanyAndAgreement() bool`

HasCompanyAndAgreement returns a boolean if a field has been set.

### GetCompanyName

`func (o *TimeEntryChangeLog) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *TimeEntryChangeLog) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *TimeEntryChangeLog) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *TimeEntryChangeLog) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetTimeStart

`func (o *TimeEntryChangeLog) GetTimeStart() string`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *TimeEntryChangeLog) GetTimeStartOk() (*string, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *TimeEntryChangeLog) SetTimeStart(v string)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *TimeEntryChangeLog) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetTimeStartUtc

`func (o *TimeEntryChangeLog) GetTimeStartUtc() string`

GetTimeStartUtc returns the TimeStartUtc field if non-nil, zero value otherwise.

### GetTimeStartUtcOk

`func (o *TimeEntryChangeLog) GetTimeStartUtcOk() (*string, bool)`

GetTimeStartUtcOk returns a tuple with the TimeStartUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStartUtc

`func (o *TimeEntryChangeLog) SetTimeStartUtc(v string)`

SetTimeStartUtc sets TimeStartUtc field to given value.

### HasTimeStartUtc

`func (o *TimeEntryChangeLog) HasTimeStartUtc() bool`

HasTimeStartUtc returns a boolean if a field has been set.

### GetTimeEnd

`func (o *TimeEntryChangeLog) GetTimeEnd() string`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *TimeEntryChangeLog) GetTimeEndOk() (*string, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *TimeEntryChangeLog) SetTimeEnd(v string)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *TimeEntryChangeLog) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.

### GetTimeEndUtc

`func (o *TimeEntryChangeLog) GetTimeEndUtc() string`

GetTimeEndUtc returns the TimeEndUtc field if non-nil, zero value otherwise.

### GetTimeEndUtcOk

`func (o *TimeEntryChangeLog) GetTimeEndUtcOk() (*string, bool)`

GetTimeEndUtcOk returns a tuple with the TimeEndUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEndUtc

`func (o *TimeEntryChangeLog) SetTimeEndUtc(v string)`

SetTimeEndUtc sets TimeEndUtc field to given value.

### HasTimeEndUtc

`func (o *TimeEntryChangeLog) HasTimeEndUtc() bool`

HasTimeEndUtc returns a boolean if a field has been set.

### GetDateStart

`func (o *TimeEntryChangeLog) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *TimeEntryChangeLog) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *TimeEntryChangeLog) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *TimeEntryChangeLog) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateInvoice

`func (o *TimeEntryChangeLog) GetDateInvoice() string`

GetDateInvoice returns the DateInvoice field if non-nil, zero value otherwise.

### GetDateInvoiceOk

`func (o *TimeEntryChangeLog) GetDateInvoiceOk() (*string, bool)`

GetDateInvoiceOk returns a tuple with the DateInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateInvoice

`func (o *TimeEntryChangeLog) SetDateInvoice(v string)`

SetDateInvoice sets DateInvoice field to given value.

### HasDateInvoice

`func (o *TimeEntryChangeLog) HasDateInvoice() bool`

HasDateInvoice returns a boolean if a field has been set.

### GetFirstName

`func (o *TimeEntryChangeLog) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TimeEntryChangeLog) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TimeEntryChangeLog) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *TimeEntryChangeLog) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHourlyCost

`func (o *TimeEntryChangeLog) GetHourlyCost() string`

GetHourlyCost returns the HourlyCost field if non-nil, zero value otherwise.

### GetHourlyCostOk

`func (o *TimeEntryChangeLog) GetHourlyCostOk() (*string, bool)`

GetHourlyCostOk returns a tuple with the HourlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCost

`func (o *TimeEntryChangeLog) SetHourlyCost(v string)`

SetHourlyCost sets HourlyCost field to given value.

### HasHourlyCost

`func (o *TimeEntryChangeLog) HasHourlyCost() bool`

HasHourlyCost returns a boolean if a field has been set.

### GetHourlyCostDecimal

`func (o *TimeEntryChangeLog) GetHourlyCostDecimal() float64`

GetHourlyCostDecimal returns the HourlyCostDecimal field if non-nil, zero value otherwise.

### GetHourlyCostDecimalOk

`func (o *TimeEntryChangeLog) GetHourlyCostDecimalOk() (*float64, bool)`

GetHourlyCostDecimalOk returns a tuple with the HourlyCostDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyCostDecimal

`func (o *TimeEntryChangeLog) SetHourlyCostDecimal(v float64)`

SetHourlyCostDecimal sets HourlyCostDecimal field to given value.

### HasHourlyCostDecimal

`func (o *TimeEntryChangeLog) HasHourlyCostDecimal() bool`

HasHourlyCostDecimal returns a boolean if a field has been set.

### SetHourlyCostDecimalNil

`func (o *TimeEntryChangeLog) SetHourlyCostDecimalNil(b bool)`

 SetHourlyCostDecimalNil sets the value for HourlyCostDecimal to be an explicit nil

### UnsetHourlyCostDecimal
`func (o *TimeEntryChangeLog) UnsetHourlyCostDecimal()`

UnsetHourlyCostDecimal ensures that no value is present for HourlyCostDecimal, not even an explicit nil
### GetHourlyRate

`func (o *TimeEntryChangeLog) GetHourlyRate() float64`

GetHourlyRate returns the HourlyRate field if non-nil, zero value otherwise.

### GetHourlyRateOk

`func (o *TimeEntryChangeLog) GetHourlyRateOk() (*float64, bool)`

GetHourlyRateOk returns a tuple with the HourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyRate

`func (o *TimeEntryChangeLog) SetHourlyRate(v float64)`

SetHourlyRate sets HourlyRate field to given value.

### HasHourlyRate

`func (o *TimeEntryChangeLog) HasHourlyRate() bool`

HasHourlyRate returns a boolean if a field has been set.

### SetHourlyRateNil

`func (o *TimeEntryChangeLog) SetHourlyRateNil(b bool)`

 SetHourlyRateNil sets the value for HourlyRate to be an explicit nil

### UnsetHourlyRate
`func (o *TimeEntryChangeLog) UnsetHourlyRate()`

UnsetHourlyRate ensures that no value is present for HourlyRate, not even an explicit nil
### GetHoursActual

`func (o *TimeEntryChangeLog) GetHoursActual() float64`

GetHoursActual returns the HoursActual field if non-nil, zero value otherwise.

### GetHoursActualOk

`func (o *TimeEntryChangeLog) GetHoursActualOk() (*float64, bool)`

GetHoursActualOk returns a tuple with the HoursActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursActual

`func (o *TimeEntryChangeLog) SetHoursActual(v float64)`

SetHoursActual sets HoursActual field to given value.

### HasHoursActual

`func (o *TimeEntryChangeLog) HasHoursActual() bool`

HasHoursActual returns a boolean if a field has been set.

### SetHoursActualNil

`func (o *TimeEntryChangeLog) SetHoursActualNil(b bool)`

 SetHoursActualNil sets the value for HoursActual to be an explicit nil

### UnsetHoursActual
`func (o *TimeEntryChangeLog) UnsetHoursActual()`

UnsetHoursActual ensures that no value is present for HoursActual, not even an explicit nil
### GetInternalNote

`func (o *TimeEntryChangeLog) GetInternalNote() string`

GetInternalNote returns the InternalNote field if non-nil, zero value otherwise.

### GetInternalNoteOk

`func (o *TimeEntryChangeLog) GetInternalNoteOk() (*string, bool)`

GetInternalNoteOk returns a tuple with the InternalNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNote

`func (o *TimeEntryChangeLog) SetInternalNote(v string)`

SetInternalNote sets InternalNote field to given value.

### HasInternalNote

`func (o *TimeEntryChangeLog) HasInternalNote() bool`

HasInternalNote returns a boolean if a field has been set.

### GetInvoiceAdjustmentFirm

`func (o *TimeEntryChangeLog) GetInvoiceAdjustmentFirm() float64`

GetInvoiceAdjustmentFirm returns the InvoiceAdjustmentFirm field if non-nil, zero value otherwise.

### GetInvoiceAdjustmentFirmOk

`func (o *TimeEntryChangeLog) GetInvoiceAdjustmentFirmOk() (*float64, bool)`

GetInvoiceAdjustmentFirmOk returns a tuple with the InvoiceAdjustmentFirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAdjustmentFirm

`func (o *TimeEntryChangeLog) SetInvoiceAdjustmentFirm(v float64)`

SetInvoiceAdjustmentFirm sets InvoiceAdjustmentFirm field to given value.

### HasInvoiceAdjustmentFirm

`func (o *TimeEntryChangeLog) HasInvoiceAdjustmentFirm() bool`

HasInvoiceAdjustmentFirm returns a boolean if a field has been set.

### SetInvoiceAdjustmentFirmNil

`func (o *TimeEntryChangeLog) SetInvoiceAdjustmentFirmNil(b bool)`

 SetInvoiceAdjustmentFirmNil sets the value for InvoiceAdjustmentFirm to be an explicit nil

### UnsetInvoiceAdjustmentFirm
`func (o *TimeEntryChangeLog) UnsetInvoiceAdjustmentFirm()`

UnsetInvoiceAdjustmentFirm ensures that no value is present for InvoiceAdjustmentFirm, not even an explicit nil
### GetInvoiceAdjustmentTotal

`func (o *TimeEntryChangeLog) GetInvoiceAdjustmentTotal() float64`

GetInvoiceAdjustmentTotal returns the InvoiceAdjustmentTotal field if non-nil, zero value otherwise.

### GetInvoiceAdjustmentTotalOk

`func (o *TimeEntryChangeLog) GetInvoiceAdjustmentTotalOk() (*float64, bool)`

GetInvoiceAdjustmentTotalOk returns a tuple with the InvoiceAdjustmentTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceAdjustmentTotal

`func (o *TimeEntryChangeLog) SetInvoiceAdjustmentTotal(v float64)`

SetInvoiceAdjustmentTotal sets InvoiceAdjustmentTotal field to given value.

### HasInvoiceAdjustmentTotal

`func (o *TimeEntryChangeLog) HasInvoiceAdjustmentTotal() bool`

HasInvoiceAdjustmentTotal returns a boolean if a field has been set.

### SetInvoiceAdjustmentTotalNil

`func (o *TimeEntryChangeLog) SetInvoiceAdjustmentTotalNil(b bool)`

 SetInvoiceAdjustmentTotalNil sets the value for InvoiceAdjustmentTotal to be an explicit nil

### UnsetInvoiceAdjustmentTotal
`func (o *TimeEntryChangeLog) UnsetInvoiceAdjustmentTotal()`

UnsetInvoiceAdjustmentTotal ensures that no value is present for InvoiceAdjustmentTotal, not even an explicit nil
### GetInvoiceFlag

`func (o *TimeEntryChangeLog) GetInvoiceFlag() bool`

GetInvoiceFlag returns the InvoiceFlag field if non-nil, zero value otherwise.

### GetInvoiceFlagOk

`func (o *TimeEntryChangeLog) GetInvoiceFlagOk() (*bool, bool)`

GetInvoiceFlagOk returns a tuple with the InvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFlag

`func (o *TimeEntryChangeLog) SetInvoiceFlag(v bool)`

SetInvoiceFlag sets InvoiceFlag field to given value.

### HasInvoiceFlag

`func (o *TimeEntryChangeLog) HasInvoiceFlag() bool`

HasInvoiceFlag returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *TimeEntryChangeLog) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *TimeEntryChangeLog) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *TimeEntryChangeLog) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *TimeEntryChangeLog) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetInvoiceReady

`func (o *TimeEntryChangeLog) GetInvoiceReady() bool`

GetInvoiceReady returns the InvoiceReady field if non-nil, zero value otherwise.

### GetInvoiceReadyOk

`func (o *TimeEntryChangeLog) GetInvoiceReadyOk() (*bool, bool)`

GetInvoiceReadyOk returns a tuple with the InvoiceReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceReady

`func (o *TimeEntryChangeLog) SetInvoiceReady(v bool)`

SetInvoiceReady sets InvoiceReady field to given value.

### HasInvoiceReady

`func (o *TimeEntryChangeLog) HasInvoiceReady() bool`

HasInvoiceReady returns a boolean if a field has been set.

### GetLastName

`func (o *TimeEntryChangeLog) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TimeEntryChangeLog) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TimeEntryChangeLog) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *TimeEntryChangeLog) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMemberId

`func (o *TimeEntryChangeLog) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *TimeEntryChangeLog) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *TimeEntryChangeLog) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *TimeEntryChangeLog) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetNonBillableAmt

`func (o *TimeEntryChangeLog) GetNonBillableAmt() float64`

GetNonBillableAmt returns the NonBillableAmt field if non-nil, zero value otherwise.

### GetNonBillableAmtOk

`func (o *TimeEntryChangeLog) GetNonBillableAmtOk() (*float64, bool)`

GetNonBillableAmtOk returns a tuple with the NonBillableAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonBillableAmt

`func (o *TimeEntryChangeLog) SetNonBillableAmt(v float64)`

SetNonBillableAmt sets NonBillableAmt field to given value.

### HasNonBillableAmt

`func (o *TimeEntryChangeLog) HasNonBillableAmt() bool`

HasNonBillableAmt returns a boolean if a field has been set.

### SetNonBillableAmtNil

`func (o *TimeEntryChangeLog) SetNonBillableAmtNil(b bool)`

 SetNonBillableAmtNil sets the value for NonBillableAmt to be an explicit nil

### UnsetNonBillableAmt
`func (o *TimeEntryChangeLog) UnsetNonBillableAmt()`

UnsetNonBillableAmt ensures that no value is present for NonBillableAmt, not even an explicit nil
### GetNonBillableHrs

`func (o *TimeEntryChangeLog) GetNonBillableHrs() float64`

GetNonBillableHrs returns the NonBillableHrs field if non-nil, zero value otherwise.

### GetNonBillableHrsOk

`func (o *TimeEntryChangeLog) GetNonBillableHrsOk() (*float64, bool)`

GetNonBillableHrsOk returns a tuple with the NonBillableHrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonBillableHrs

`func (o *TimeEntryChangeLog) SetNonBillableHrs(v float64)`

SetNonBillableHrs sets NonBillableHrs field to given value.

### HasNonBillableHrs

`func (o *TimeEntryChangeLog) HasNonBillableHrs() bool`

HasNonBillableHrs returns a boolean if a field has been set.

### SetNonBillableHrsNil

`func (o *TimeEntryChangeLog) SetNonBillableHrsNil(b bool)`

 SetNonBillableHrsNil sets the value for NonBillableHrs to be an explicit nil

### UnsetNonBillableHrs
`func (o *TimeEntryChangeLog) UnsetNonBillableHrs()`

UnsetNonBillableHrs ensures that no value is present for NonBillableHrs, not even an explicit nil
### GetNotes

`func (o *TimeEntryChangeLog) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TimeEntryChangeLog) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TimeEntryChangeLog) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TimeEntryChangeLog) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetOpportunityRecId

`func (o *TimeEntryChangeLog) GetOpportunityRecId() int32`

GetOpportunityRecId returns the OpportunityRecId field if non-nil, zero value otherwise.

### GetOpportunityRecIdOk

`func (o *TimeEntryChangeLog) GetOpportunityRecIdOk() (*int32, bool)`

GetOpportunityRecIdOk returns a tuple with the OpportunityRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityRecId

`func (o *TimeEntryChangeLog) SetOpportunityRecId(v int32)`

SetOpportunityRecId sets OpportunityRecId field to given value.

### HasOpportunityRecId

`func (o *TimeEntryChangeLog) HasOpportunityRecId() bool`

HasOpportunityRecId returns a boolean if a field has been set.

### SetOpportunityRecIdNil

`func (o *TimeEntryChangeLog) SetOpportunityRecIdNil(b bool)`

 SetOpportunityRecIdNil sets the value for OpportunityRecId to be an explicit nil

### UnsetOpportunityRecId
`func (o *TimeEntryChangeLog) UnsetOpportunityRecId()`

UnsetOpportunityRecId ensures that no value is present for OpportunityRecId, not even an explicit nil
### GetOptionId

`func (o *TimeEntryChangeLog) GetOptionId() string`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *TimeEntryChangeLog) GetOptionIdOk() (*string, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *TimeEntryChangeLog) SetOptionId(v string)`

SetOptionId sets OptionId field to given value.

### HasOptionId

`func (o *TimeEntryChangeLog) HasOptionId() bool`

HasOptionId returns a boolean if a field has been set.

### GetProjectActivity

`func (o *TimeEntryChangeLog) GetProjectActivity() string`

GetProjectActivity returns the ProjectActivity field if non-nil, zero value otherwise.

### GetProjectActivityOk

`func (o *TimeEntryChangeLog) GetProjectActivityOk() (*string, bool)`

GetProjectActivityOk returns a tuple with the ProjectActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectActivity

`func (o *TimeEntryChangeLog) SetProjectActivity(v string)`

SetProjectActivity sets ProjectActivity field to given value.

### HasProjectActivity

`func (o *TimeEntryChangeLog) HasProjectActivity() bool`

HasProjectActivity returns a boolean if a field has been set.

### GetProjectName

`func (o *TimeEntryChangeLog) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *TimeEntryChangeLog) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *TimeEntryChangeLog) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *TimeEntryChangeLog) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetProjectPhase

`func (o *TimeEntryChangeLog) GetProjectPhase() string`

GetProjectPhase returns the ProjectPhase field if non-nil, zero value otherwise.

### GetProjectPhaseOk

`func (o *TimeEntryChangeLog) GetProjectPhaseOk() (*string, bool)`

GetProjectPhaseOk returns a tuple with the ProjectPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPhase

`func (o *TimeEntryChangeLog) SetProjectPhase(v string)`

SetProjectPhase sets ProjectPhase field to given value.

### HasProjectPhase

`func (o *TimeEntryChangeLog) HasProjectPhase() bool`

HasProjectPhase returns a boolean if a field has been set.

### GetServiceRequestStatus

`func (o *TimeEntryChangeLog) GetServiceRequestStatus() string`

GetServiceRequestStatus returns the ServiceRequestStatus field if non-nil, zero value otherwise.

### GetServiceRequestStatusOk

`func (o *TimeEntryChangeLog) GetServiceRequestStatusOk() (*string, bool)`

GetServiceRequestStatusOk returns a tuple with the ServiceRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestStatus

`func (o *TimeEntryChangeLog) SetServiceRequestStatus(v string)`

SetServiceRequestStatus sets ServiceRequestStatus field to given value.

### HasServiceRequestStatus

`func (o *TimeEntryChangeLog) HasServiceRequestStatus() bool`

HasServiceRequestStatus returns a boolean if a field has been set.

### GetServiceRequestSummary

`func (o *TimeEntryChangeLog) GetServiceRequestSummary() string`

GetServiceRequestSummary returns the ServiceRequestSummary field if non-nil, zero value otherwise.

### GetServiceRequestSummaryOk

`func (o *TimeEntryChangeLog) GetServiceRequestSummaryOk() (*string, bool)`

GetServiceRequestSummaryOk returns a tuple with the ServiceRequestSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestSummary

`func (o *TimeEntryChangeLog) SetServiceRequestSummary(v string)`

SetServiceRequestSummary sets ServiceRequestSummary field to given value.

### HasServiceRequestSummary

`func (o *TimeEntryChangeLog) HasServiceRequestSummary() bool`

HasServiceRequestSummary returns a boolean if a field has been set.

### GetTerritory

`func (o *TimeEntryChangeLog) GetTerritory() string`

GetTerritory returns the Territory field if non-nil, zero value otherwise.

### GetTerritoryOk

`func (o *TimeEntryChangeLog) GetTerritoryOk() (*string, bool)`

GetTerritoryOk returns a tuple with the Territory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritory

`func (o *TimeEntryChangeLog) SetTerritory(v string)`

SetTerritory sets Territory field to given value.

### HasTerritory

`func (o *TimeEntryChangeLog) HasTerritory() bool`

HasTerritory returns a boolean if a field has been set.

### GetTimeRecId

`func (o *TimeEntryChangeLog) GetTimeRecId() int32`

GetTimeRecId returns the TimeRecId field if non-nil, zero value otherwise.

### GetTimeRecIdOk

`func (o *TimeEntryChangeLog) GetTimeRecIdOk() (*int32, bool)`

GetTimeRecIdOk returns a tuple with the TimeRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRecId

`func (o *TimeEntryChangeLog) SetTimeRecId(v int32)`

SetTimeRecId sets TimeRecId field to given value.

### HasTimeRecId

`func (o *TimeEntryChangeLog) HasTimeRecId() bool`

HasTimeRecId returns a boolean if a field has been set.

### GetTimeStatus

`func (o *TimeEntryChangeLog) GetTimeStatus() string`

GetTimeStatus returns the TimeStatus field if non-nil, zero value otherwise.

### GetTimeStatusOk

`func (o *TimeEntryChangeLog) GetTimeStatusOk() (*string, bool)`

GetTimeStatusOk returns a tuple with the TimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStatus

`func (o *TimeEntryChangeLog) SetTimeStatus(v string)`

SetTimeStatus sets TimeStatus field to given value.

### HasTimeStatus

`func (o *TimeEntryChangeLog) HasTimeStatus() bool`

HasTimeStatus returns a boolean if a field has been set.

### GetUtilizationFlag

`func (o *TimeEntryChangeLog) GetUtilizationFlag() bool`

GetUtilizationFlag returns the UtilizationFlag field if non-nil, zero value otherwise.

### GetUtilizationFlagOk

`func (o *TimeEntryChangeLog) GetUtilizationFlagOk() (*bool, bool)`

GetUtilizationFlagOk returns a tuple with the UtilizationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationFlag

`func (o *TimeEntryChangeLog) SetUtilizationFlag(v bool)`

SetUtilizationFlag sets UtilizationFlag field to given value.

### HasUtilizationFlag

`func (o *TimeEntryChangeLog) HasUtilizationFlag() bool`

HasUtilizationFlag returns a boolean if a field has been set.

### SetUtilizationFlagNil

`func (o *TimeEntryChangeLog) SetUtilizationFlagNil(b bool)`

 SetUtilizationFlagNil sets the value for UtilizationFlag to be an explicit nil

### UnsetUtilizationFlag
`func (o *TimeEntryChangeLog) UnsetUtilizationFlag()`

UnsetUtilizationFlag ensures that no value is present for UtilizationFlag, not even an explicit nil
### GetCompanyType

`func (o *TimeEntryChangeLog) GetCompanyType() string`

GetCompanyType returns the CompanyType field if non-nil, zero value otherwise.

### GetCompanyTypeOk

`func (o *TimeEntryChangeLog) GetCompanyTypeOk() (*string, bool)`

GetCompanyTypeOk returns a tuple with the CompanyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyType

`func (o *TimeEntryChangeLog) SetCompanyType(v string)`

SetCompanyType sets CompanyType field to given value.

### HasCompanyType

`func (o *TimeEntryChangeLog) HasCompanyType() bool`

HasCompanyType returns a boolean if a field has been set.

### GetTicketCurrentBoard

`func (o *TimeEntryChangeLog) GetTicketCurrentBoard() string`

GetTicketCurrentBoard returns the TicketCurrentBoard field if non-nil, zero value otherwise.

### GetTicketCurrentBoardOk

`func (o *TimeEntryChangeLog) GetTicketCurrentBoardOk() (*string, bool)`

GetTicketCurrentBoardOk returns a tuple with the TicketCurrentBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketCurrentBoard

`func (o *TimeEntryChangeLog) SetTicketCurrentBoard(v string)`

SetTicketCurrentBoard sets TicketCurrentBoard field to given value.

### HasTicketCurrentBoard

`func (o *TimeEntryChangeLog) HasTicketCurrentBoard() bool`

HasTicketCurrentBoard returns a boolean if a field has been set.

### GetTicketType

`func (o *TimeEntryChangeLog) GetTicketType() string`

GetTicketType returns the TicketType field if non-nil, zero value otherwise.

### GetTicketTypeOk

`func (o *TimeEntryChangeLog) GetTicketTypeOk() (*string, bool)`

GetTicketTypeOk returns a tuple with the TicketType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketType

`func (o *TimeEntryChangeLog) SetTicketType(v string)`

SetTicketType sets TicketType field to given value.

### HasTicketType

`func (o *TimeEntryChangeLog) HasTicketType() bool`

HasTicketType returns a boolean if a field has been set.

### GetTicketSubtype

`func (o *TimeEntryChangeLog) GetTicketSubtype() string`

GetTicketSubtype returns the TicketSubtype field if non-nil, zero value otherwise.

### GetTicketSubtypeOk

`func (o *TimeEntryChangeLog) GetTicketSubtypeOk() (*string, bool)`

GetTicketSubtypeOk returns a tuple with the TicketSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketSubtype

`func (o *TimeEntryChangeLog) SetTicketSubtype(v string)`

SetTicketSubtype sets TicketSubtype field to given value.

### HasTicketSubtype

`func (o *TimeEntryChangeLog) HasTicketSubtype() bool`

HasTicketSubtype returns a boolean if a field has been set.

### GetAgreementType

`func (o *TimeEntryChangeLog) GetAgreementType() string`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *TimeEntryChangeLog) GetAgreementTypeOk() (*string, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *TimeEntryChangeLog) SetAgreementType(v string)`

SetAgreementType sets AgreementType field to given value.

### HasAgreementType

`func (o *TimeEntryChangeLog) HasAgreementType() bool`

HasAgreementType returns a boolean if a field has been set.

### GetBillingStatus

`func (o *TimeEntryChangeLog) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *TimeEntryChangeLog) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *TimeEntryChangeLog) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.

### HasBillingStatus

`func (o *TimeEntryChangeLog) HasBillingStatus() bool`

HasBillingStatus returns a boolean if a field has been set.

### GetProcessingStatus

`func (o *TimeEntryChangeLog) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *TimeEntryChangeLog) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *TimeEntryChangeLog) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.

### HasProcessingStatus

`func (o *TimeEntryChangeLog) HasProcessingStatus() bool`

HasProcessingStatus returns a boolean if a field has been set.

### GetInvoicedhours

`func (o *TimeEntryChangeLog) GetInvoicedhours() float64`

GetInvoicedhours returns the Invoicedhours field if non-nil, zero value otherwise.

### GetInvoicedhoursOk

`func (o *TimeEntryChangeLog) GetInvoicedhoursOk() (*float64, bool)`

GetInvoicedhoursOk returns a tuple with the Invoicedhours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicedhours

`func (o *TimeEntryChangeLog) SetInvoicedhours(v float64)`

SetInvoicedhours sets Invoicedhours field to given value.

### HasInvoicedhours

`func (o *TimeEntryChangeLog) HasInvoicedhours() bool`

HasInvoicedhours returns a boolean if a field has been set.

### SetInvoicedhoursNil

`func (o *TimeEntryChangeLog) SetInvoicedhoursNil(b bool)`

 SetInvoicedhoursNil sets the value for Invoicedhours to be an explicit nil

### UnsetInvoicedhours
`func (o *TimeEntryChangeLog) UnsetInvoicedhours()`

UnsetInvoicedhours ensures that no value is present for Invoicedhours, not even an explicit nil
### GetCompany

`func (o *TimeEntryChangeLog) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TimeEntryChangeLog) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TimeEntryChangeLog) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *TimeEntryChangeLog) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetMember

`func (o *TimeEntryChangeLog) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *TimeEntryChangeLog) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *TimeEntryChangeLog) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *TimeEntryChangeLog) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetWorkType

`func (o *TimeEntryChangeLog) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *TimeEntryChangeLog) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *TimeEntryChangeLog) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *TimeEntryChangeLog) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetWorkRole

`func (o *TimeEntryChangeLog) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *TimeEntryChangeLog) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *TimeEntryChangeLog) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *TimeEntryChangeLog) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetAgreement

`func (o *TimeEntryChangeLog) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *TimeEntryChangeLog) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *TimeEntryChangeLog) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *TimeEntryChangeLog) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetInvoice

`func (o *TimeEntryChangeLog) GetInvoice() InvoiceReference`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *TimeEntryChangeLog) GetInvoiceOk() (*InvoiceReference, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *TimeEntryChangeLog) SetInvoice(v InvoiceReference)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *TimeEntryChangeLog) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetTicket

`func (o *TimeEntryChangeLog) GetTicket() TicketReference`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *TimeEntryChangeLog) GetTicketOk() (*TicketReference, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *TimeEntryChangeLog) SetTicket(v TicketReference)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *TimeEntryChangeLog) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetProject

`func (o *TimeEntryChangeLog) GetProject() ProjectReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *TimeEntryChangeLog) GetProjectOk() (*ProjectReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *TimeEntryChangeLog) SetProject(v ProjectReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *TimeEntryChangeLog) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetPhase

`func (o *TimeEntryChangeLog) GetPhase() ProjectPhaseReference`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *TimeEntryChangeLog) GetPhaseOk() (*ProjectPhaseReference, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *TimeEntryChangeLog) SetPhase(v ProjectPhaseReference)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *TimeEntryChangeLog) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetActivity

`func (o *TimeEntryChangeLog) GetActivity() ActivityReference`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *TimeEntryChangeLog) GetActivityOk() (*ActivityReference, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *TimeEntryChangeLog) SetActivity(v ActivityReference)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *TimeEntryChangeLog) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetInfo

`func (o *TimeEntryChangeLog) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeEntryChangeLog) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeEntryChangeLog) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeEntryChangeLog) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


