# OpportunityStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Probability** | Pointer to [**OpportunityProbabilityReference**](OpportunityProbabilityReference.md) |  | [optional] 
**Color** | Pointer to **string** |  Max length: 25; | [optional] 
**SequenceNumber** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOpportunityStage

`func NewOpportunityStage(name string, ) *OpportunityStage`

NewOpportunityStage instantiates a new OpportunityStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityStageWithDefaults

`func NewOpportunityStageWithDefaults() *OpportunityStage`

NewOpportunityStageWithDefaults instantiates a new OpportunityStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpportunityStage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpportunityStage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpportunityStage) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpportunityStage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OpportunityStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpportunityStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpportunityStage) SetName(v string)`

SetName sets Name field to given value.


### GetProbability

`func (o *OpportunityStage) GetProbability() OpportunityProbabilityReference`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *OpportunityStage) GetProbabilityOk() (*OpportunityProbabilityReference, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *OpportunityStage) SetProbability(v OpportunityProbabilityReference)`

SetProbability sets Probability field to given value.

### HasProbability

`func (o *OpportunityStage) HasProbability() bool`

HasProbability returns a boolean if a field has been set.

### GetColor

`func (o *OpportunityStage) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *OpportunityStage) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *OpportunityStage) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *OpportunityStage) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *OpportunityStage) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *OpportunityStage) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *OpportunityStage) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *OpportunityStage) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *OpportunityStage) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *OpportunityStage) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetInfo

`func (o *OpportunityStage) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OpportunityStage) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OpportunityStage) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OpportunityStage) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


