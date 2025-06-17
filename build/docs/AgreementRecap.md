# AgreementRecap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AdjustmentAmount** | Pointer to **float64** |  | [optional] 
**AgreementStatus** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AvailableAmount** | Pointer to **float64** |  | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**IsUnlimited** | Pointer to **string** |  | [optional] 
**LastInvoiceAmount** | Pointer to **string** |  | [optional] 
**LastInvoiceDate** | Pointer to **string** |  | [optional] 
**LastInvoiceNumber** | Pointer to **string** |  | [optional] 
**NextInvoiceAmount** | Pointer to **float64** |  | [optional] 
**NextInvoiceDate** | Pointer to **string** |  | [optional] 
**OverrunAmount** | Pointer to **float64** |  | [optional] 
**RemainingAmount** | Pointer to **float64** |  | [optional] 
**StartingAmount** | Pointer to **float64** |  | [optional] 
**UnbilledOverageAmount** | Pointer to **float64** |  | [optional] 
**UnbilledPeriods** | Pointer to **int32** |  | [optional] 
**UsedAmount** | Pointer to **float64** |  | [optional] 

## Methods

### NewAgreementRecap

`func NewAgreementRecap() *AgreementRecap`

NewAgreementRecap instantiates a new AgreementRecap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementRecapWithDefaults

`func NewAgreementRecapWithDefaults() *AgreementRecap`

NewAgreementRecapWithDefaults instantiates a new AgreementRecap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementRecap) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementRecap) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementRecap) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementRecap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAdjustmentAmount

`func (o *AgreementRecap) GetAdjustmentAmount() float64`

GetAdjustmentAmount returns the AdjustmentAmount field if non-nil, zero value otherwise.

### GetAdjustmentAmountOk

`func (o *AgreementRecap) GetAdjustmentAmountOk() (*float64, bool)`

GetAdjustmentAmountOk returns a tuple with the AdjustmentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentAmount

`func (o *AgreementRecap) SetAdjustmentAmount(v float64)`

SetAdjustmentAmount sets AdjustmentAmount field to given value.

### HasAdjustmentAmount

`func (o *AgreementRecap) HasAdjustmentAmount() bool`

HasAdjustmentAmount returns a boolean if a field has been set.

### GetAgreementStatus

`func (o *AgreementRecap) GetAgreementStatus() string`

GetAgreementStatus returns the AgreementStatus field if non-nil, zero value otherwise.

### GetAgreementStatusOk

`func (o *AgreementRecap) GetAgreementStatusOk() (*string, bool)`

GetAgreementStatusOk returns a tuple with the AgreementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementStatus

`func (o *AgreementRecap) SetAgreementStatus(v string)`

SetAgreementStatus sets AgreementStatus field to given value.

### HasAgreementStatus

`func (o *AgreementRecap) HasAgreementStatus() bool`

HasAgreementStatus returns a boolean if a field has been set.

### GetName

`func (o *AgreementRecap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgreementRecap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgreementRecap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgreementRecap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvailableAmount

`func (o *AgreementRecap) GetAvailableAmount() float64`

GetAvailableAmount returns the AvailableAmount field if non-nil, zero value otherwise.

### GetAvailableAmountOk

`func (o *AgreementRecap) GetAvailableAmountOk() (*float64, bool)`

GetAvailableAmountOk returns a tuple with the AvailableAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAmount

`func (o *AgreementRecap) SetAvailableAmount(v float64)`

SetAvailableAmount sets AvailableAmount field to given value.

### HasAvailableAmount

`func (o *AgreementRecap) HasAvailableAmount() bool`

HasAvailableAmount returns a boolean if a field has been set.

### GetCompanyName

`func (o *AgreementRecap) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *AgreementRecap) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *AgreementRecap) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *AgreementRecap) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetIsUnlimited

`func (o *AgreementRecap) GetIsUnlimited() string`

GetIsUnlimited returns the IsUnlimited field if non-nil, zero value otherwise.

### GetIsUnlimitedOk

`func (o *AgreementRecap) GetIsUnlimitedOk() (*string, bool)`

GetIsUnlimitedOk returns a tuple with the IsUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnlimited

`func (o *AgreementRecap) SetIsUnlimited(v string)`

SetIsUnlimited sets IsUnlimited field to given value.

### HasIsUnlimited

`func (o *AgreementRecap) HasIsUnlimited() bool`

HasIsUnlimited returns a boolean if a field has been set.

### GetLastInvoiceAmount

`func (o *AgreementRecap) GetLastInvoiceAmount() string`

GetLastInvoiceAmount returns the LastInvoiceAmount field if non-nil, zero value otherwise.

### GetLastInvoiceAmountOk

`func (o *AgreementRecap) GetLastInvoiceAmountOk() (*string, bool)`

GetLastInvoiceAmountOk returns a tuple with the LastInvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInvoiceAmount

`func (o *AgreementRecap) SetLastInvoiceAmount(v string)`

SetLastInvoiceAmount sets LastInvoiceAmount field to given value.

### HasLastInvoiceAmount

`func (o *AgreementRecap) HasLastInvoiceAmount() bool`

HasLastInvoiceAmount returns a boolean if a field has been set.

### GetLastInvoiceDate

`func (o *AgreementRecap) GetLastInvoiceDate() string`

GetLastInvoiceDate returns the LastInvoiceDate field if non-nil, zero value otherwise.

### GetLastInvoiceDateOk

`func (o *AgreementRecap) GetLastInvoiceDateOk() (*string, bool)`

GetLastInvoiceDateOk returns a tuple with the LastInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInvoiceDate

`func (o *AgreementRecap) SetLastInvoiceDate(v string)`

SetLastInvoiceDate sets LastInvoiceDate field to given value.

### HasLastInvoiceDate

`func (o *AgreementRecap) HasLastInvoiceDate() bool`

HasLastInvoiceDate returns a boolean if a field has been set.

### GetLastInvoiceNumber

`func (o *AgreementRecap) GetLastInvoiceNumber() string`

GetLastInvoiceNumber returns the LastInvoiceNumber field if non-nil, zero value otherwise.

### GetLastInvoiceNumberOk

`func (o *AgreementRecap) GetLastInvoiceNumberOk() (*string, bool)`

GetLastInvoiceNumberOk returns a tuple with the LastInvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastInvoiceNumber

`func (o *AgreementRecap) SetLastInvoiceNumber(v string)`

SetLastInvoiceNumber sets LastInvoiceNumber field to given value.

### HasLastInvoiceNumber

`func (o *AgreementRecap) HasLastInvoiceNumber() bool`

HasLastInvoiceNumber returns a boolean if a field has been set.

### GetNextInvoiceAmount

`func (o *AgreementRecap) GetNextInvoiceAmount() float64`

GetNextInvoiceAmount returns the NextInvoiceAmount field if non-nil, zero value otherwise.

### GetNextInvoiceAmountOk

`func (o *AgreementRecap) GetNextInvoiceAmountOk() (*float64, bool)`

GetNextInvoiceAmountOk returns a tuple with the NextInvoiceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceAmount

`func (o *AgreementRecap) SetNextInvoiceAmount(v float64)`

SetNextInvoiceAmount sets NextInvoiceAmount field to given value.

### HasNextInvoiceAmount

`func (o *AgreementRecap) HasNextInvoiceAmount() bool`

HasNextInvoiceAmount returns a boolean if a field has been set.

### GetNextInvoiceDate

`func (o *AgreementRecap) GetNextInvoiceDate() string`

GetNextInvoiceDate returns the NextInvoiceDate field if non-nil, zero value otherwise.

### GetNextInvoiceDateOk

`func (o *AgreementRecap) GetNextInvoiceDateOk() (*string, bool)`

GetNextInvoiceDateOk returns a tuple with the NextInvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextInvoiceDate

`func (o *AgreementRecap) SetNextInvoiceDate(v string)`

SetNextInvoiceDate sets NextInvoiceDate field to given value.

### HasNextInvoiceDate

`func (o *AgreementRecap) HasNextInvoiceDate() bool`

HasNextInvoiceDate returns a boolean if a field has been set.

### GetOverrunAmount

`func (o *AgreementRecap) GetOverrunAmount() float64`

GetOverrunAmount returns the OverrunAmount field if non-nil, zero value otherwise.

### GetOverrunAmountOk

`func (o *AgreementRecap) GetOverrunAmountOk() (*float64, bool)`

GetOverrunAmountOk returns a tuple with the OverrunAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrunAmount

`func (o *AgreementRecap) SetOverrunAmount(v float64)`

SetOverrunAmount sets OverrunAmount field to given value.

### HasOverrunAmount

`func (o *AgreementRecap) HasOverrunAmount() bool`

HasOverrunAmount returns a boolean if a field has been set.

### GetRemainingAmount

`func (o *AgreementRecap) GetRemainingAmount() float64`

GetRemainingAmount returns the RemainingAmount field if non-nil, zero value otherwise.

### GetRemainingAmountOk

`func (o *AgreementRecap) GetRemainingAmountOk() (*float64, bool)`

GetRemainingAmountOk returns a tuple with the RemainingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingAmount

`func (o *AgreementRecap) SetRemainingAmount(v float64)`

SetRemainingAmount sets RemainingAmount field to given value.

### HasRemainingAmount

`func (o *AgreementRecap) HasRemainingAmount() bool`

HasRemainingAmount returns a boolean if a field has been set.

### GetStartingAmount

`func (o *AgreementRecap) GetStartingAmount() float64`

GetStartingAmount returns the StartingAmount field if non-nil, zero value otherwise.

### GetStartingAmountOk

`func (o *AgreementRecap) GetStartingAmountOk() (*float64, bool)`

GetStartingAmountOk returns a tuple with the StartingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAmount

`func (o *AgreementRecap) SetStartingAmount(v float64)`

SetStartingAmount sets StartingAmount field to given value.

### HasStartingAmount

`func (o *AgreementRecap) HasStartingAmount() bool`

HasStartingAmount returns a boolean if a field has been set.

### GetUnbilledOverageAmount

`func (o *AgreementRecap) GetUnbilledOverageAmount() float64`

GetUnbilledOverageAmount returns the UnbilledOverageAmount field if non-nil, zero value otherwise.

### GetUnbilledOverageAmountOk

`func (o *AgreementRecap) GetUnbilledOverageAmountOk() (*float64, bool)`

GetUnbilledOverageAmountOk returns a tuple with the UnbilledOverageAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbilledOverageAmount

`func (o *AgreementRecap) SetUnbilledOverageAmount(v float64)`

SetUnbilledOverageAmount sets UnbilledOverageAmount field to given value.

### HasUnbilledOverageAmount

`func (o *AgreementRecap) HasUnbilledOverageAmount() bool`

HasUnbilledOverageAmount returns a boolean if a field has been set.

### GetUnbilledPeriods

`func (o *AgreementRecap) GetUnbilledPeriods() int32`

GetUnbilledPeriods returns the UnbilledPeriods field if non-nil, zero value otherwise.

### GetUnbilledPeriodsOk

`func (o *AgreementRecap) GetUnbilledPeriodsOk() (*int32, bool)`

GetUnbilledPeriodsOk returns a tuple with the UnbilledPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbilledPeriods

`func (o *AgreementRecap) SetUnbilledPeriods(v int32)`

SetUnbilledPeriods sets UnbilledPeriods field to given value.

### HasUnbilledPeriods

`func (o *AgreementRecap) HasUnbilledPeriods() bool`

HasUnbilledPeriods returns a boolean if a field has been set.

### GetUsedAmount

`func (o *AgreementRecap) GetUsedAmount() float64`

GetUsedAmount returns the UsedAmount field if non-nil, zero value otherwise.

### GetUsedAmountOk

`func (o *AgreementRecap) GetUsedAmountOk() (*float64, bool)`

GetUsedAmountOk returns a tuple with the UsedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAmount

`func (o *AgreementRecap) SetUsedAmount(v float64)`

SetUsedAmount sets UsedAmount field to given value.

### HasUsedAmount

`func (o *AgreementRecap) HasUsedAmount() bool`

HasUsedAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


