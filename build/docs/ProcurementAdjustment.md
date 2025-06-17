# ProcurementAdjustment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 50; | 
**Type** | [**AdjustmentTypeReference**](AdjustmentTypeReference.md) |  | 
**Reason** | Pointer to **string** |  Max length: 100; | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedBy** | Pointer to **string** |  | [optional] 
**ClosedDate** | Pointer to **time.Time** |  | [optional] 
**AdjustmentDetails** | Pointer to [**[]AdjustmentDetail**](AdjustmentDetail.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProcurementAdjustment

`func NewProcurementAdjustment(identifier string, type_ AdjustmentTypeReference, ) *ProcurementAdjustment`

NewProcurementAdjustment instantiates a new ProcurementAdjustment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcurementAdjustmentWithDefaults

`func NewProcurementAdjustmentWithDefaults() *ProcurementAdjustment`

NewProcurementAdjustmentWithDefaults instantiates a new ProcurementAdjustment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcurementAdjustment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcurementAdjustment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcurementAdjustment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProcurementAdjustment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *ProcurementAdjustment) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ProcurementAdjustment) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ProcurementAdjustment) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetType

`func (o *ProcurementAdjustment) GetType() AdjustmentTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcurementAdjustment) GetTypeOk() (*AdjustmentTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcurementAdjustment) SetType(v AdjustmentTypeReference)`

SetType sets Type field to given value.


### GetReason

`func (o *ProcurementAdjustment) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ProcurementAdjustment) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ProcurementAdjustment) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ProcurementAdjustment) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetNotes

`func (o *ProcurementAdjustment) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ProcurementAdjustment) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ProcurementAdjustment) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ProcurementAdjustment) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetClosedFlag

`func (o *ProcurementAdjustment) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *ProcurementAdjustment) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *ProcurementAdjustment) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *ProcurementAdjustment) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *ProcurementAdjustment) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *ProcurementAdjustment) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetClosedBy

`func (o *ProcurementAdjustment) GetClosedBy() string`

GetClosedBy returns the ClosedBy field if non-nil, zero value otherwise.

### GetClosedByOk

`func (o *ProcurementAdjustment) GetClosedByOk() (*string, bool)`

GetClosedByOk returns a tuple with the ClosedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedBy

`func (o *ProcurementAdjustment) SetClosedBy(v string)`

SetClosedBy sets ClosedBy field to given value.

### HasClosedBy

`func (o *ProcurementAdjustment) HasClosedBy() bool`

HasClosedBy returns a boolean if a field has been set.

### GetClosedDate

`func (o *ProcurementAdjustment) GetClosedDate() time.Time`

GetClosedDate returns the ClosedDate field if non-nil, zero value otherwise.

### GetClosedDateOk

`func (o *ProcurementAdjustment) GetClosedDateOk() (*time.Time, bool)`

GetClosedDateOk returns a tuple with the ClosedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDate

`func (o *ProcurementAdjustment) SetClosedDate(v time.Time)`

SetClosedDate sets ClosedDate field to given value.

### HasClosedDate

`func (o *ProcurementAdjustment) HasClosedDate() bool`

HasClosedDate returns a boolean if a field has been set.

### GetAdjustmentDetails

`func (o *ProcurementAdjustment) GetAdjustmentDetails() []AdjustmentDetail`

GetAdjustmentDetails returns the AdjustmentDetails field if non-nil, zero value otherwise.

### GetAdjustmentDetailsOk

`func (o *ProcurementAdjustment) GetAdjustmentDetailsOk() (*[]AdjustmentDetail, bool)`

GetAdjustmentDetailsOk returns a tuple with the AdjustmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentDetails

`func (o *ProcurementAdjustment) SetAdjustmentDetails(v []AdjustmentDetail)`

SetAdjustmentDetails sets AdjustmentDetails field to given value.

### HasAdjustmentDetails

`func (o *ProcurementAdjustment) HasAdjustmentDetails() bool`

HasAdjustmentDetails returns a boolean if a field has been set.

### GetInfo

`func (o *ProcurementAdjustment) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProcurementAdjustment) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProcurementAdjustment) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProcurementAdjustment) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


