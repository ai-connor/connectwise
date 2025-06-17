# QuoteLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Link** | **string** |  Max length: 2000; | 
**AllLocationsFlag** | Pointer to **NullableBool** |  | [optional] 
**NewWindowFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewQuoteLink

`func NewQuoteLink(link string, ) *QuoteLink`

NewQuoteLink instantiates a new QuoteLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteLinkWithDefaults

`func NewQuoteLinkWithDefaults() *QuoteLink`

NewQuoteLinkWithDefaults instantiates a new QuoteLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuoteLink) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuoteLink) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuoteLink) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QuoteLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *QuoteLink) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *QuoteLink) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *QuoteLink) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *QuoteLink) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLink

`func (o *QuoteLink) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *QuoteLink) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *QuoteLink) SetLink(v string)`

SetLink sets Link field to given value.


### GetAllLocationsFlag

`func (o *QuoteLink) GetAllLocationsFlag() bool`

GetAllLocationsFlag returns the AllLocationsFlag field if non-nil, zero value otherwise.

### GetAllLocationsFlagOk

`func (o *QuoteLink) GetAllLocationsFlagOk() (*bool, bool)`

GetAllLocationsFlagOk returns a tuple with the AllLocationsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLocationsFlag

`func (o *QuoteLink) SetAllLocationsFlag(v bool)`

SetAllLocationsFlag sets AllLocationsFlag field to given value.

### HasAllLocationsFlag

`func (o *QuoteLink) HasAllLocationsFlag() bool`

HasAllLocationsFlag returns a boolean if a field has been set.

### SetAllLocationsFlagNil

`func (o *QuoteLink) SetAllLocationsFlagNil(b bool)`

 SetAllLocationsFlagNil sets the value for AllLocationsFlag to be an explicit nil

### UnsetAllLocationsFlag
`func (o *QuoteLink) UnsetAllLocationsFlag()`

UnsetAllLocationsFlag ensures that no value is present for AllLocationsFlag, not even an explicit nil
### GetNewWindowFlag

`func (o *QuoteLink) GetNewWindowFlag() bool`

GetNewWindowFlag returns the NewWindowFlag field if non-nil, zero value otherwise.

### GetNewWindowFlagOk

`func (o *QuoteLink) GetNewWindowFlagOk() (*bool, bool)`

GetNewWindowFlagOk returns a tuple with the NewWindowFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWindowFlag

`func (o *QuoteLink) SetNewWindowFlag(v bool)`

SetNewWindowFlag sets NewWindowFlag field to given value.

### HasNewWindowFlag

`func (o *QuoteLink) HasNewWindowFlag() bool`

HasNewWindowFlag returns a boolean if a field has been set.

### SetNewWindowFlagNil

`func (o *QuoteLink) SetNewWindowFlagNil(b bool)`

 SetNewWindowFlagNil sets the value for NewWindowFlag to be an explicit nil

### UnsetNewWindowFlag
`func (o *QuoteLink) UnsetNewWindowFlag()`

UnsetNewWindowFlag ensures that no value is present for NewWindowFlag, not even an explicit nil
### GetInfo

`func (o *QuoteLink) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *QuoteLink) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *QuoteLink) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *QuoteLink) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


