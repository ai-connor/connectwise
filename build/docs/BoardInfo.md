# BoardInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ProjectFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedLoopDiscussionsFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedLoopInternalAnalysisFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedLoopResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedLoopAllFlag** | Pointer to **NullableBool** |  | [optional] 
**OverrideBillingSetupFlag** | Pointer to **NullableBool** |  | [optional] 
**BillTicketsAfterClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**BillUnapprovedTimeExpenseFlag** | Pointer to **NullableBool** |  | [optional] 
**BillTime** | Pointer to **NullableString** |  | [optional] 
**BillExpense** | Pointer to **NullableString** |  | [optional] 
**BillProduct** | Pointer to **NullableString** |  | [optional] 
**ProblemSort** | Pointer to **NullableString** |  | [optional] 
**InternalAnalysisSort** | Pointer to **NullableString** |  | [optional] 
**ResolutionSort** | Pointer to **NullableString** |  | [optional] 
**AllSort** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardInfo

`func NewBoardInfo() *BoardInfo`

NewBoardInfo instantiates a new BoardInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardInfoWithDefaults

`func NewBoardInfoWithDefaults() *BoardInfo`

NewBoardInfoWithDefaults instantiates a new BoardInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoardInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoardInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *BoardInfo) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BoardInfo) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BoardInfo) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BoardInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *BoardInfo) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *BoardInfo) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *BoardInfo) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *BoardInfo) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetProjectFlag

`func (o *BoardInfo) GetProjectFlag() bool`

GetProjectFlag returns the ProjectFlag field if non-nil, zero value otherwise.

### GetProjectFlagOk

`func (o *BoardInfo) GetProjectFlagOk() (*bool, bool)`

GetProjectFlagOk returns a tuple with the ProjectFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectFlag

`func (o *BoardInfo) SetProjectFlag(v bool)`

SetProjectFlag sets ProjectFlag field to given value.

### HasProjectFlag

`func (o *BoardInfo) HasProjectFlag() bool`

HasProjectFlag returns a boolean if a field has been set.

### SetProjectFlagNil

`func (o *BoardInfo) SetProjectFlagNil(b bool)`

 SetProjectFlagNil sets the value for ProjectFlag to be an explicit nil

### UnsetProjectFlag
`func (o *BoardInfo) UnsetProjectFlag()`

UnsetProjectFlag ensures that no value is present for ProjectFlag, not even an explicit nil
### GetInactiveFlag

`func (o *BoardInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *BoardInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *BoardInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *BoardInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *BoardInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *BoardInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetClosedLoopDiscussionsFlag

`func (o *BoardInfo) GetClosedLoopDiscussionsFlag() bool`

GetClosedLoopDiscussionsFlag returns the ClosedLoopDiscussionsFlag field if non-nil, zero value otherwise.

### GetClosedLoopDiscussionsFlagOk

`func (o *BoardInfo) GetClosedLoopDiscussionsFlagOk() (*bool, bool)`

GetClosedLoopDiscussionsFlagOk returns a tuple with the ClosedLoopDiscussionsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedLoopDiscussionsFlag

`func (o *BoardInfo) SetClosedLoopDiscussionsFlag(v bool)`

SetClosedLoopDiscussionsFlag sets ClosedLoopDiscussionsFlag field to given value.

### HasClosedLoopDiscussionsFlag

`func (o *BoardInfo) HasClosedLoopDiscussionsFlag() bool`

HasClosedLoopDiscussionsFlag returns a boolean if a field has been set.

### SetClosedLoopDiscussionsFlagNil

`func (o *BoardInfo) SetClosedLoopDiscussionsFlagNil(b bool)`

 SetClosedLoopDiscussionsFlagNil sets the value for ClosedLoopDiscussionsFlag to be an explicit nil

### UnsetClosedLoopDiscussionsFlag
`func (o *BoardInfo) UnsetClosedLoopDiscussionsFlag()`

UnsetClosedLoopDiscussionsFlag ensures that no value is present for ClosedLoopDiscussionsFlag, not even an explicit nil
### GetClosedLoopInternalAnalysisFlag

`func (o *BoardInfo) GetClosedLoopInternalAnalysisFlag() bool`

GetClosedLoopInternalAnalysisFlag returns the ClosedLoopInternalAnalysisFlag field if non-nil, zero value otherwise.

### GetClosedLoopInternalAnalysisFlagOk

`func (o *BoardInfo) GetClosedLoopInternalAnalysisFlagOk() (*bool, bool)`

GetClosedLoopInternalAnalysisFlagOk returns a tuple with the ClosedLoopInternalAnalysisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedLoopInternalAnalysisFlag

`func (o *BoardInfo) SetClosedLoopInternalAnalysisFlag(v bool)`

SetClosedLoopInternalAnalysisFlag sets ClosedLoopInternalAnalysisFlag field to given value.

### HasClosedLoopInternalAnalysisFlag

`func (o *BoardInfo) HasClosedLoopInternalAnalysisFlag() bool`

HasClosedLoopInternalAnalysisFlag returns a boolean if a field has been set.

### SetClosedLoopInternalAnalysisFlagNil

`func (o *BoardInfo) SetClosedLoopInternalAnalysisFlagNil(b bool)`

 SetClosedLoopInternalAnalysisFlagNil sets the value for ClosedLoopInternalAnalysisFlag to be an explicit nil

### UnsetClosedLoopInternalAnalysisFlag
`func (o *BoardInfo) UnsetClosedLoopInternalAnalysisFlag()`

UnsetClosedLoopInternalAnalysisFlag ensures that no value is present for ClosedLoopInternalAnalysisFlag, not even an explicit nil
### GetClosedLoopResolutionFlag

`func (o *BoardInfo) GetClosedLoopResolutionFlag() bool`

GetClosedLoopResolutionFlag returns the ClosedLoopResolutionFlag field if non-nil, zero value otherwise.

### GetClosedLoopResolutionFlagOk

`func (o *BoardInfo) GetClosedLoopResolutionFlagOk() (*bool, bool)`

GetClosedLoopResolutionFlagOk returns a tuple with the ClosedLoopResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedLoopResolutionFlag

`func (o *BoardInfo) SetClosedLoopResolutionFlag(v bool)`

SetClosedLoopResolutionFlag sets ClosedLoopResolutionFlag field to given value.

### HasClosedLoopResolutionFlag

`func (o *BoardInfo) HasClosedLoopResolutionFlag() bool`

HasClosedLoopResolutionFlag returns a boolean if a field has been set.

### SetClosedLoopResolutionFlagNil

`func (o *BoardInfo) SetClosedLoopResolutionFlagNil(b bool)`

 SetClosedLoopResolutionFlagNil sets the value for ClosedLoopResolutionFlag to be an explicit nil

### UnsetClosedLoopResolutionFlag
`func (o *BoardInfo) UnsetClosedLoopResolutionFlag()`

UnsetClosedLoopResolutionFlag ensures that no value is present for ClosedLoopResolutionFlag, not even an explicit nil
### GetClosedLoopAllFlag

`func (o *BoardInfo) GetClosedLoopAllFlag() bool`

GetClosedLoopAllFlag returns the ClosedLoopAllFlag field if non-nil, zero value otherwise.

### GetClosedLoopAllFlagOk

`func (o *BoardInfo) GetClosedLoopAllFlagOk() (*bool, bool)`

GetClosedLoopAllFlagOk returns a tuple with the ClosedLoopAllFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedLoopAllFlag

`func (o *BoardInfo) SetClosedLoopAllFlag(v bool)`

SetClosedLoopAllFlag sets ClosedLoopAllFlag field to given value.

### HasClosedLoopAllFlag

`func (o *BoardInfo) HasClosedLoopAllFlag() bool`

HasClosedLoopAllFlag returns a boolean if a field has been set.

### SetClosedLoopAllFlagNil

`func (o *BoardInfo) SetClosedLoopAllFlagNil(b bool)`

 SetClosedLoopAllFlagNil sets the value for ClosedLoopAllFlag to be an explicit nil

### UnsetClosedLoopAllFlag
`func (o *BoardInfo) UnsetClosedLoopAllFlag()`

UnsetClosedLoopAllFlag ensures that no value is present for ClosedLoopAllFlag, not even an explicit nil
### GetOverrideBillingSetupFlag

`func (o *BoardInfo) GetOverrideBillingSetupFlag() bool`

GetOverrideBillingSetupFlag returns the OverrideBillingSetupFlag field if non-nil, zero value otherwise.

### GetOverrideBillingSetupFlagOk

`func (o *BoardInfo) GetOverrideBillingSetupFlagOk() (*bool, bool)`

GetOverrideBillingSetupFlagOk returns a tuple with the OverrideBillingSetupFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideBillingSetupFlag

`func (o *BoardInfo) SetOverrideBillingSetupFlag(v bool)`

SetOverrideBillingSetupFlag sets OverrideBillingSetupFlag field to given value.

### HasOverrideBillingSetupFlag

`func (o *BoardInfo) HasOverrideBillingSetupFlag() bool`

HasOverrideBillingSetupFlag returns a boolean if a field has been set.

### SetOverrideBillingSetupFlagNil

`func (o *BoardInfo) SetOverrideBillingSetupFlagNil(b bool)`

 SetOverrideBillingSetupFlagNil sets the value for OverrideBillingSetupFlag to be an explicit nil

### UnsetOverrideBillingSetupFlag
`func (o *BoardInfo) UnsetOverrideBillingSetupFlag()`

UnsetOverrideBillingSetupFlag ensures that no value is present for OverrideBillingSetupFlag, not even an explicit nil
### GetBillTicketsAfterClosedFlag

`func (o *BoardInfo) GetBillTicketsAfterClosedFlag() bool`

GetBillTicketsAfterClosedFlag returns the BillTicketsAfterClosedFlag field if non-nil, zero value otherwise.

### GetBillTicketsAfterClosedFlagOk

`func (o *BoardInfo) GetBillTicketsAfterClosedFlagOk() (*bool, bool)`

GetBillTicketsAfterClosedFlagOk returns a tuple with the BillTicketsAfterClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTicketsAfterClosedFlag

`func (o *BoardInfo) SetBillTicketsAfterClosedFlag(v bool)`

SetBillTicketsAfterClosedFlag sets BillTicketsAfterClosedFlag field to given value.

### HasBillTicketsAfterClosedFlag

`func (o *BoardInfo) HasBillTicketsAfterClosedFlag() bool`

HasBillTicketsAfterClosedFlag returns a boolean if a field has been set.

### SetBillTicketsAfterClosedFlagNil

`func (o *BoardInfo) SetBillTicketsAfterClosedFlagNil(b bool)`

 SetBillTicketsAfterClosedFlagNil sets the value for BillTicketsAfterClosedFlag to be an explicit nil

### UnsetBillTicketsAfterClosedFlag
`func (o *BoardInfo) UnsetBillTicketsAfterClosedFlag()`

UnsetBillTicketsAfterClosedFlag ensures that no value is present for BillTicketsAfterClosedFlag, not even an explicit nil
### GetBillUnapprovedTimeExpenseFlag

`func (o *BoardInfo) GetBillUnapprovedTimeExpenseFlag() bool`

GetBillUnapprovedTimeExpenseFlag returns the BillUnapprovedTimeExpenseFlag field if non-nil, zero value otherwise.

### GetBillUnapprovedTimeExpenseFlagOk

`func (o *BoardInfo) GetBillUnapprovedTimeExpenseFlagOk() (*bool, bool)`

GetBillUnapprovedTimeExpenseFlagOk returns a tuple with the BillUnapprovedTimeExpenseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillUnapprovedTimeExpenseFlag

`func (o *BoardInfo) SetBillUnapprovedTimeExpenseFlag(v bool)`

SetBillUnapprovedTimeExpenseFlag sets BillUnapprovedTimeExpenseFlag field to given value.

### HasBillUnapprovedTimeExpenseFlag

`func (o *BoardInfo) HasBillUnapprovedTimeExpenseFlag() bool`

HasBillUnapprovedTimeExpenseFlag returns a boolean if a field has been set.

### SetBillUnapprovedTimeExpenseFlagNil

`func (o *BoardInfo) SetBillUnapprovedTimeExpenseFlagNil(b bool)`

 SetBillUnapprovedTimeExpenseFlagNil sets the value for BillUnapprovedTimeExpenseFlag to be an explicit nil

### UnsetBillUnapprovedTimeExpenseFlag
`func (o *BoardInfo) UnsetBillUnapprovedTimeExpenseFlag()`

UnsetBillUnapprovedTimeExpenseFlag ensures that no value is present for BillUnapprovedTimeExpenseFlag, not even an explicit nil
### GetBillTime

`func (o *BoardInfo) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *BoardInfo) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *BoardInfo) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.

### HasBillTime

`func (o *BoardInfo) HasBillTime() bool`

HasBillTime returns a boolean if a field has been set.

### SetBillTimeNil

`func (o *BoardInfo) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *BoardInfo) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetBillExpense

`func (o *BoardInfo) GetBillExpense() string`

GetBillExpense returns the BillExpense field if non-nil, zero value otherwise.

### GetBillExpenseOk

`func (o *BoardInfo) GetBillExpenseOk() (*string, bool)`

GetBillExpenseOk returns a tuple with the BillExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillExpense

`func (o *BoardInfo) SetBillExpense(v string)`

SetBillExpense sets BillExpense field to given value.

### HasBillExpense

`func (o *BoardInfo) HasBillExpense() bool`

HasBillExpense returns a boolean if a field has been set.

### SetBillExpenseNil

`func (o *BoardInfo) SetBillExpenseNil(b bool)`

 SetBillExpenseNil sets the value for BillExpense to be an explicit nil

### UnsetBillExpense
`func (o *BoardInfo) UnsetBillExpense()`

UnsetBillExpense ensures that no value is present for BillExpense, not even an explicit nil
### GetBillProduct

`func (o *BoardInfo) GetBillProduct() string`

GetBillProduct returns the BillProduct field if non-nil, zero value otherwise.

### GetBillProductOk

`func (o *BoardInfo) GetBillProductOk() (*string, bool)`

GetBillProductOk returns a tuple with the BillProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillProduct

`func (o *BoardInfo) SetBillProduct(v string)`

SetBillProduct sets BillProduct field to given value.

### HasBillProduct

`func (o *BoardInfo) HasBillProduct() bool`

HasBillProduct returns a boolean if a field has been set.

### SetBillProductNil

`func (o *BoardInfo) SetBillProductNil(b bool)`

 SetBillProductNil sets the value for BillProduct to be an explicit nil

### UnsetBillProduct
`func (o *BoardInfo) UnsetBillProduct()`

UnsetBillProduct ensures that no value is present for BillProduct, not even an explicit nil
### GetProblemSort

`func (o *BoardInfo) GetProblemSort() string`

GetProblemSort returns the ProblemSort field if non-nil, zero value otherwise.

### GetProblemSortOk

`func (o *BoardInfo) GetProblemSortOk() (*string, bool)`

GetProblemSortOk returns a tuple with the ProblemSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemSort

`func (o *BoardInfo) SetProblemSort(v string)`

SetProblemSort sets ProblemSort field to given value.

### HasProblemSort

`func (o *BoardInfo) HasProblemSort() bool`

HasProblemSort returns a boolean if a field has been set.

### SetProblemSortNil

`func (o *BoardInfo) SetProblemSortNil(b bool)`

 SetProblemSortNil sets the value for ProblemSort to be an explicit nil

### UnsetProblemSort
`func (o *BoardInfo) UnsetProblemSort()`

UnsetProblemSort ensures that no value is present for ProblemSort, not even an explicit nil
### GetInternalAnalysisSort

`func (o *BoardInfo) GetInternalAnalysisSort() string`

GetInternalAnalysisSort returns the InternalAnalysisSort field if non-nil, zero value otherwise.

### GetInternalAnalysisSortOk

`func (o *BoardInfo) GetInternalAnalysisSortOk() (*string, bool)`

GetInternalAnalysisSortOk returns a tuple with the InternalAnalysisSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysisSort

`func (o *BoardInfo) SetInternalAnalysisSort(v string)`

SetInternalAnalysisSort sets InternalAnalysisSort field to given value.

### HasInternalAnalysisSort

`func (o *BoardInfo) HasInternalAnalysisSort() bool`

HasInternalAnalysisSort returns a boolean if a field has been set.

### SetInternalAnalysisSortNil

`func (o *BoardInfo) SetInternalAnalysisSortNil(b bool)`

 SetInternalAnalysisSortNil sets the value for InternalAnalysisSort to be an explicit nil

### UnsetInternalAnalysisSort
`func (o *BoardInfo) UnsetInternalAnalysisSort()`

UnsetInternalAnalysisSort ensures that no value is present for InternalAnalysisSort, not even an explicit nil
### GetResolutionSort

`func (o *BoardInfo) GetResolutionSort() string`

GetResolutionSort returns the ResolutionSort field if non-nil, zero value otherwise.

### GetResolutionSortOk

`func (o *BoardInfo) GetResolutionSortOk() (*string, bool)`

GetResolutionSortOk returns a tuple with the ResolutionSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionSort

`func (o *BoardInfo) SetResolutionSort(v string)`

SetResolutionSort sets ResolutionSort field to given value.

### HasResolutionSort

`func (o *BoardInfo) HasResolutionSort() bool`

HasResolutionSort returns a boolean if a field has been set.

### SetResolutionSortNil

`func (o *BoardInfo) SetResolutionSortNil(b bool)`

 SetResolutionSortNil sets the value for ResolutionSort to be an explicit nil

### UnsetResolutionSort
`func (o *BoardInfo) UnsetResolutionSort()`

UnsetResolutionSort ensures that no value is present for ResolutionSort, not even an explicit nil
### GetAllSort

`func (o *BoardInfo) GetAllSort() string`

GetAllSort returns the AllSort field if non-nil, zero value otherwise.

### GetAllSortOk

`func (o *BoardInfo) GetAllSortOk() (*string, bool)`

GetAllSortOk returns a tuple with the AllSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSort

`func (o *BoardInfo) SetAllSort(v string)`

SetAllSort sets AllSort field to given value.

### HasAllSort

`func (o *BoardInfo) HasAllSort() bool`

HasAllSort returns a boolean if a field has been set.

### SetAllSortNil

`func (o *BoardInfo) SetAllSortNil(b bool)`

 SetAllSortNil sets the value for AllSort to be an explicit nil

### UnsetAllSort
`func (o *BoardInfo) UnsetAllSort()`

UnsetAllSort ensures that no value is present for AllSort, not even an explicit nil
### GetInfo

`func (o *BoardInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


