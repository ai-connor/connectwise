# SalesTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SalesTeamIdentifier** | **string** |  Max length: 20; | 
**SalesTeamDescription** | **string** |  Max length: 50; | 
**SalesTeamLocation** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSalesTeam

`func NewSalesTeam(salesTeamIdentifier string, salesTeamDescription string, salesTeamLocation SystemLocationReference, ) *SalesTeam`

NewSalesTeam instantiates a new SalesTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesTeamWithDefaults

`func NewSalesTeamWithDefaults() *SalesTeam`

NewSalesTeamWithDefaults instantiates a new SalesTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SalesTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSalesTeamIdentifier

`func (o *SalesTeam) GetSalesTeamIdentifier() string`

GetSalesTeamIdentifier returns the SalesTeamIdentifier field if non-nil, zero value otherwise.

### GetSalesTeamIdentifierOk

`func (o *SalesTeam) GetSalesTeamIdentifierOk() (*string, bool)`

GetSalesTeamIdentifierOk returns a tuple with the SalesTeamIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTeamIdentifier

`func (o *SalesTeam) SetSalesTeamIdentifier(v string)`

SetSalesTeamIdentifier sets SalesTeamIdentifier field to given value.


### GetSalesTeamDescription

`func (o *SalesTeam) GetSalesTeamDescription() string`

GetSalesTeamDescription returns the SalesTeamDescription field if non-nil, zero value otherwise.

### GetSalesTeamDescriptionOk

`func (o *SalesTeam) GetSalesTeamDescriptionOk() (*string, bool)`

GetSalesTeamDescriptionOk returns a tuple with the SalesTeamDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTeamDescription

`func (o *SalesTeam) SetSalesTeamDescription(v string)`

SetSalesTeamDescription sets SalesTeamDescription field to given value.


### GetSalesTeamLocation

`func (o *SalesTeam) GetSalesTeamLocation() SystemLocationReference`

GetSalesTeamLocation returns the SalesTeamLocation field if non-nil, zero value otherwise.

### GetSalesTeamLocationOk

`func (o *SalesTeam) GetSalesTeamLocationOk() (*SystemLocationReference, bool)`

GetSalesTeamLocationOk returns a tuple with the SalesTeamLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTeamLocation

`func (o *SalesTeam) SetSalesTeamLocation(v SystemLocationReference)`

SetSalesTeamLocation sets SalesTeamLocation field to given value.


### GetInactiveFlag

`func (o *SalesTeam) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *SalesTeam) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *SalesTeam) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *SalesTeam) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *SalesTeam) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *SalesTeam) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *SalesTeam) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SalesTeam) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SalesTeam) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SalesTeam) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


