# State

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 50; | 
**Name** | **string** |  Max length: 50; | 
**Country** | [**CountryReference**](CountryReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewState

`func NewState(identifier string, name string, country CountryReference, ) *State`

NewState instantiates a new State object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateWithDefaults

`func NewStateWithDefaults() *State`

NewStateWithDefaults instantiates a new State object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *State) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *State) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *State) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *State) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *State) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *State) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *State) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *State) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *State) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *State) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *State) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *State) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *State) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.


### GetInfo

`func (o *State) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *State) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *State) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *State) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


