# StandardNoteInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewStandardNoteInfo

`func NewStandardNoteInfo() *StandardNoteInfo`

NewStandardNoteInfo instantiates a new StandardNoteInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardNoteInfoWithDefaults

`func NewStandardNoteInfoWithDefaults() *StandardNoteInfo`

NewStandardNoteInfoWithDefaults instantiates a new StandardNoteInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StandardNoteInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StandardNoteInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StandardNoteInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StandardNoteInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StandardNoteInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StandardNoteInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StandardNoteInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StandardNoteInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetContents

`func (o *StandardNoteInfo) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *StandardNoteInfo) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *StandardNoteInfo) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *StandardNoteInfo) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetLocation

`func (o *StandardNoteInfo) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StandardNoteInfo) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StandardNoteInfo) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *StandardNoteInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *StandardNoteInfo) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *StandardNoteInfo) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *StandardNoteInfo) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *StandardNoteInfo) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetBoard

`func (o *StandardNoteInfo) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *StandardNoteInfo) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *StandardNoteInfo) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *StandardNoteInfo) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetInfo

`func (o *StandardNoteInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *StandardNoteInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *StandardNoteInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *StandardNoteInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


