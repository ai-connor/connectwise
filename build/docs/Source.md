# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**EnteredBy** | Pointer to **string** |  | [optional] 
**DateEntered** | Pointer to **time.Time** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSource

`func NewSource(name string, ) *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Source) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Source) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Source) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Source) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Source) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Source) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Source) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *Source) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *Source) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *Source) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *Source) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *Source) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *Source) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetEnteredBy

`func (o *Source) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *Source) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *Source) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.

### HasEnteredBy

`func (o *Source) HasEnteredBy() bool`

HasEnteredBy returns a boolean if a field has been set.

### GetDateEntered

`func (o *Source) GetDateEntered() time.Time`

GetDateEntered returns the DateEntered field if non-nil, zero value otherwise.

### GetDateEnteredOk

`func (o *Source) GetDateEnteredOk() (*time.Time, bool)`

GetDateEnteredOk returns a tuple with the DateEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEntered

`func (o *Source) SetDateEntered(v time.Time)`

SetDateEntered sets DateEntered field to given value.

### HasDateEntered

`func (o *Source) HasDateEntered() bool`

HasDateEntered returns a boolean if a field has been set.

### GetInfo

`func (o *Source) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Source) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Source) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Source) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


