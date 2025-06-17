# SurveyOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Caption** | **string** |  Max length: 100; | 
**Points** | **NullableInt32** |  | 
**Visibleflag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSurveyOption

`func NewSurveyOption(caption string, points NullableInt32, ) *SurveyOption`

NewSurveyOption instantiates a new SurveyOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSurveyOptionWithDefaults

`func NewSurveyOptionWithDefaults() *SurveyOption`

NewSurveyOptionWithDefaults instantiates a new SurveyOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SurveyOption) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SurveyOption) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SurveyOption) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SurveyOption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCaption

`func (o *SurveyOption) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *SurveyOption) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *SurveyOption) SetCaption(v string)`

SetCaption sets Caption field to given value.


### GetPoints

`func (o *SurveyOption) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *SurveyOption) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *SurveyOption) SetPoints(v int32)`

SetPoints sets Points field to given value.


### SetPointsNil

`func (o *SurveyOption) SetPointsNil(b bool)`

 SetPointsNil sets the value for Points to be an explicit nil

### UnsetPoints
`func (o *SurveyOption) UnsetPoints()`

UnsetPoints ensures that no value is present for Points, not even an explicit nil
### GetVisibleflag

`func (o *SurveyOption) GetVisibleflag() bool`

GetVisibleflag returns the Visibleflag field if non-nil, zero value otherwise.

### GetVisibleflagOk

`func (o *SurveyOption) GetVisibleflagOk() (*bool, bool)`

GetVisibleflagOk returns a tuple with the Visibleflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleflag

`func (o *SurveyOption) SetVisibleflag(v bool)`

SetVisibleflag sets Visibleflag field to given value.

### HasVisibleflag

`func (o *SurveyOption) HasVisibleflag() bool`

HasVisibleflag returns a boolean if a field has been set.

### SetVisibleflagNil

`func (o *SurveyOption) SetVisibleflagNil(b bool)`

 SetVisibleflagNil sets the value for Visibleflag to be an explicit nil

### UnsetVisibleflag
`func (o *SurveyOption) UnsetVisibleflag()`

UnsetVisibleflag ensures that no value is present for Visibleflag, not even an explicit nil
### GetInfo

`func (o *SurveyOption) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SurveyOption) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SurveyOption) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SurveyOption) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


