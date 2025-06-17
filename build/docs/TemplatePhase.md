# TemplatePhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentPhase** | Pointer to [**ProjectTemplatePhaseReference**](ProjectTemplatePhaseReference.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**TemplateId** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**MarkAsMilestoneFlag** | Pointer to **NullableBool** |  | [optional] 
**BillPhaseSeparately** | Pointer to **NullableBool** |  | [optional] 
**WbsCode** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTemplatePhase

`func NewTemplatePhase() *TemplatePhase`

NewTemplatePhase instantiates a new TemplatePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePhaseWithDefaults

`func NewTemplatePhaseWithDefaults() *TemplatePhase`

NewTemplatePhaseWithDefaults instantiates a new TemplatePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentPhase

`func (o *TemplatePhase) GetParentPhase() ProjectTemplatePhaseReference`

GetParentPhase returns the ParentPhase field if non-nil, zero value otherwise.

### GetParentPhaseOk

`func (o *TemplatePhase) GetParentPhaseOk() (*ProjectTemplatePhaseReference, bool)`

GetParentPhaseOk returns a tuple with the ParentPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPhase

`func (o *TemplatePhase) SetParentPhase(v ProjectTemplatePhaseReference)`

SetParentPhase sets ParentPhase field to given value.

### HasParentPhase

`func (o *TemplatePhase) HasParentPhase() bool`

HasParentPhase returns a boolean if a field has been set.

### GetId

`func (o *TemplatePhase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplatePhase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplatePhase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TemplatePhase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplateId

`func (o *TemplatePhase) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplatePhase) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplatePhase) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *TemplatePhase) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *TemplatePhase) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *TemplatePhase) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetDescription

`func (o *TemplatePhase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplatePhase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplatePhase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplatePhase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotes

`func (o *TemplatePhase) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TemplatePhase) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TemplatePhase) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TemplatePhase) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetMarkAsMilestoneFlag

`func (o *TemplatePhase) GetMarkAsMilestoneFlag() bool`

GetMarkAsMilestoneFlag returns the MarkAsMilestoneFlag field if non-nil, zero value otherwise.

### GetMarkAsMilestoneFlagOk

`func (o *TemplatePhase) GetMarkAsMilestoneFlagOk() (*bool, bool)`

GetMarkAsMilestoneFlagOk returns a tuple with the MarkAsMilestoneFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkAsMilestoneFlag

`func (o *TemplatePhase) SetMarkAsMilestoneFlag(v bool)`

SetMarkAsMilestoneFlag sets MarkAsMilestoneFlag field to given value.

### HasMarkAsMilestoneFlag

`func (o *TemplatePhase) HasMarkAsMilestoneFlag() bool`

HasMarkAsMilestoneFlag returns a boolean if a field has been set.

### SetMarkAsMilestoneFlagNil

`func (o *TemplatePhase) SetMarkAsMilestoneFlagNil(b bool)`

 SetMarkAsMilestoneFlagNil sets the value for MarkAsMilestoneFlag to be an explicit nil

### UnsetMarkAsMilestoneFlag
`func (o *TemplatePhase) UnsetMarkAsMilestoneFlag()`

UnsetMarkAsMilestoneFlag ensures that no value is present for MarkAsMilestoneFlag, not even an explicit nil
### GetBillPhaseSeparately

`func (o *TemplatePhase) GetBillPhaseSeparately() bool`

GetBillPhaseSeparately returns the BillPhaseSeparately field if non-nil, zero value otherwise.

### GetBillPhaseSeparatelyOk

`func (o *TemplatePhase) GetBillPhaseSeparatelyOk() (*bool, bool)`

GetBillPhaseSeparatelyOk returns a tuple with the BillPhaseSeparately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillPhaseSeparately

`func (o *TemplatePhase) SetBillPhaseSeparately(v bool)`

SetBillPhaseSeparately sets BillPhaseSeparately field to given value.

### HasBillPhaseSeparately

`func (o *TemplatePhase) HasBillPhaseSeparately() bool`

HasBillPhaseSeparately returns a boolean if a field has been set.

### SetBillPhaseSeparatelyNil

`func (o *TemplatePhase) SetBillPhaseSeparatelyNil(b bool)`

 SetBillPhaseSeparatelyNil sets the value for BillPhaseSeparately to be an explicit nil

### UnsetBillPhaseSeparately
`func (o *TemplatePhase) UnsetBillPhaseSeparately()`

UnsetBillPhaseSeparately ensures that no value is present for BillPhaseSeparately, not even an explicit nil
### GetWbsCode

`func (o *TemplatePhase) GetWbsCode() string`

GetWbsCode returns the WbsCode field if non-nil, zero value otherwise.

### GetWbsCodeOk

`func (o *TemplatePhase) GetWbsCodeOk() (*string, bool)`

GetWbsCodeOk returns a tuple with the WbsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWbsCode

`func (o *TemplatePhase) SetWbsCode(v string)`

SetWbsCode sets WbsCode field to given value.

### HasWbsCode

`func (o *TemplatePhase) HasWbsCode() bool`

HasWbsCode returns a boolean if a field has been set.

### GetInfo

`func (o *TemplatePhase) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TemplatePhase) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TemplatePhase) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TemplatePhase) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


