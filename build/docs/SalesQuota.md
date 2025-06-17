# SalesQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | [**MemberReference**](MemberReference.md) |  | 
**ForecastYear** | Pointer to **NullableInt32** |  | [optional] 
**Location** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Category** | Pointer to [**ProductCategoryReference**](ProductCategoryReference.md) |  | [optional] 
**SubCategory** | Pointer to [**ProductSubCategoryReference**](ProductSubCategoryReference.md) |  | [optional] 
**JanuaryRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**JanuaryMargin** | Pointer to **NullableFloat64** |  | [optional] 
**FebruaryRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**FebruaryMargin** | Pointer to **NullableFloat64** |  | [optional] 
**MarchRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**MarchMargin** | Pointer to **NullableFloat64** |  | [optional] 
**AprilRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**AprilMargin** | Pointer to **NullableFloat64** |  | [optional] 
**MayRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**MayMargin** | Pointer to **NullableFloat64** |  | [optional] 
**JuneRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**JuneMargin** | Pointer to **NullableFloat64** |  | [optional] 
**JulyRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**JulyMargin** | Pointer to **NullableFloat64** |  | [optional] 
**AugustRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**AugustMargin** | Pointer to **NullableFloat64** |  | [optional] 
**SeptemberRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**SeptemberMargin** | Pointer to **NullableFloat64** |  | [optional] 
**OctoberRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**OctoberMargin** | Pointer to **NullableFloat64** |  | [optional] 
**NovemberRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**NovemberMargin** | Pointer to **NullableFloat64** |  | [optional] 
**DecemberRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**DecemberMargin** | Pointer to **NullableFloat64** |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSalesQuota

`func NewSalesQuota(member MemberReference, location SystemLocationReference, ) *SalesQuota`

NewSalesQuota instantiates a new SalesQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesQuotaWithDefaults

`func NewSalesQuotaWithDefaults() *SalesQuota`

NewSalesQuotaWithDefaults instantiates a new SalesQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesQuota) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesQuota) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesQuota) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SalesQuota) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *SalesQuota) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *SalesQuota) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *SalesQuota) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetForecastYear

`func (o *SalesQuota) GetForecastYear() int32`

GetForecastYear returns the ForecastYear field if non-nil, zero value otherwise.

### GetForecastYearOk

`func (o *SalesQuota) GetForecastYearOk() (*int32, bool)`

GetForecastYearOk returns a tuple with the ForecastYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastYear

`func (o *SalesQuota) SetForecastYear(v int32)`

SetForecastYear sets ForecastYear field to given value.

### HasForecastYear

`func (o *SalesQuota) HasForecastYear() bool`

HasForecastYear returns a boolean if a field has been set.

### SetForecastYearNil

`func (o *SalesQuota) SetForecastYearNil(b bool)`

 SetForecastYearNil sets the value for ForecastYear to be an explicit nil

### UnsetForecastYear
`func (o *SalesQuota) UnsetForecastYear()`

UnsetForecastYear ensures that no value is present for ForecastYear, not even an explicit nil
### GetLocation

`func (o *SalesQuota) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SalesQuota) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SalesQuota) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.


### GetDepartment

`func (o *SalesQuota) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *SalesQuota) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *SalesQuota) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *SalesQuota) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetCategory

`func (o *SalesQuota) GetCategory() ProductCategoryReference`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SalesQuota) GetCategoryOk() (*ProductCategoryReference, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SalesQuota) SetCategory(v ProductCategoryReference)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SalesQuota) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSubCategory

`func (o *SalesQuota) GetSubCategory() ProductSubCategoryReference`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *SalesQuota) GetSubCategoryOk() (*ProductSubCategoryReference, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *SalesQuota) SetSubCategory(v ProductSubCategoryReference)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *SalesQuota) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetJanuaryRevenue

`func (o *SalesQuota) GetJanuaryRevenue() float64`

GetJanuaryRevenue returns the JanuaryRevenue field if non-nil, zero value otherwise.

### GetJanuaryRevenueOk

`func (o *SalesQuota) GetJanuaryRevenueOk() (*float64, bool)`

GetJanuaryRevenueOk returns a tuple with the JanuaryRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJanuaryRevenue

`func (o *SalesQuota) SetJanuaryRevenue(v float64)`

SetJanuaryRevenue sets JanuaryRevenue field to given value.

### HasJanuaryRevenue

`func (o *SalesQuota) HasJanuaryRevenue() bool`

HasJanuaryRevenue returns a boolean if a field has been set.

### SetJanuaryRevenueNil

`func (o *SalesQuota) SetJanuaryRevenueNil(b bool)`

 SetJanuaryRevenueNil sets the value for JanuaryRevenue to be an explicit nil

### UnsetJanuaryRevenue
`func (o *SalesQuota) UnsetJanuaryRevenue()`

UnsetJanuaryRevenue ensures that no value is present for JanuaryRevenue, not even an explicit nil
### GetJanuaryMargin

`func (o *SalesQuota) GetJanuaryMargin() float64`

GetJanuaryMargin returns the JanuaryMargin field if non-nil, zero value otherwise.

### GetJanuaryMarginOk

`func (o *SalesQuota) GetJanuaryMarginOk() (*float64, bool)`

GetJanuaryMarginOk returns a tuple with the JanuaryMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJanuaryMargin

`func (o *SalesQuota) SetJanuaryMargin(v float64)`

SetJanuaryMargin sets JanuaryMargin field to given value.

### HasJanuaryMargin

`func (o *SalesQuota) HasJanuaryMargin() bool`

HasJanuaryMargin returns a boolean if a field has been set.

### SetJanuaryMarginNil

`func (o *SalesQuota) SetJanuaryMarginNil(b bool)`

 SetJanuaryMarginNil sets the value for JanuaryMargin to be an explicit nil

### UnsetJanuaryMargin
`func (o *SalesQuota) UnsetJanuaryMargin()`

UnsetJanuaryMargin ensures that no value is present for JanuaryMargin, not even an explicit nil
### GetFebruaryRevenue

`func (o *SalesQuota) GetFebruaryRevenue() float64`

GetFebruaryRevenue returns the FebruaryRevenue field if non-nil, zero value otherwise.

### GetFebruaryRevenueOk

`func (o *SalesQuota) GetFebruaryRevenueOk() (*float64, bool)`

GetFebruaryRevenueOk returns a tuple with the FebruaryRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFebruaryRevenue

`func (o *SalesQuota) SetFebruaryRevenue(v float64)`

SetFebruaryRevenue sets FebruaryRevenue field to given value.

### HasFebruaryRevenue

`func (o *SalesQuota) HasFebruaryRevenue() bool`

HasFebruaryRevenue returns a boolean if a field has been set.

### SetFebruaryRevenueNil

`func (o *SalesQuota) SetFebruaryRevenueNil(b bool)`

 SetFebruaryRevenueNil sets the value for FebruaryRevenue to be an explicit nil

### UnsetFebruaryRevenue
`func (o *SalesQuota) UnsetFebruaryRevenue()`

UnsetFebruaryRevenue ensures that no value is present for FebruaryRevenue, not even an explicit nil
### GetFebruaryMargin

`func (o *SalesQuota) GetFebruaryMargin() float64`

GetFebruaryMargin returns the FebruaryMargin field if non-nil, zero value otherwise.

### GetFebruaryMarginOk

`func (o *SalesQuota) GetFebruaryMarginOk() (*float64, bool)`

GetFebruaryMarginOk returns a tuple with the FebruaryMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFebruaryMargin

`func (o *SalesQuota) SetFebruaryMargin(v float64)`

SetFebruaryMargin sets FebruaryMargin field to given value.

### HasFebruaryMargin

`func (o *SalesQuota) HasFebruaryMargin() bool`

HasFebruaryMargin returns a boolean if a field has been set.

### SetFebruaryMarginNil

`func (o *SalesQuota) SetFebruaryMarginNil(b bool)`

 SetFebruaryMarginNil sets the value for FebruaryMargin to be an explicit nil

### UnsetFebruaryMargin
`func (o *SalesQuota) UnsetFebruaryMargin()`

UnsetFebruaryMargin ensures that no value is present for FebruaryMargin, not even an explicit nil
### GetMarchRevenue

`func (o *SalesQuota) GetMarchRevenue() float64`

GetMarchRevenue returns the MarchRevenue field if non-nil, zero value otherwise.

### GetMarchRevenueOk

`func (o *SalesQuota) GetMarchRevenueOk() (*float64, bool)`

GetMarchRevenueOk returns a tuple with the MarchRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarchRevenue

`func (o *SalesQuota) SetMarchRevenue(v float64)`

SetMarchRevenue sets MarchRevenue field to given value.

### HasMarchRevenue

`func (o *SalesQuota) HasMarchRevenue() bool`

HasMarchRevenue returns a boolean if a field has been set.

### SetMarchRevenueNil

`func (o *SalesQuota) SetMarchRevenueNil(b bool)`

 SetMarchRevenueNil sets the value for MarchRevenue to be an explicit nil

### UnsetMarchRevenue
`func (o *SalesQuota) UnsetMarchRevenue()`

UnsetMarchRevenue ensures that no value is present for MarchRevenue, not even an explicit nil
### GetMarchMargin

`func (o *SalesQuota) GetMarchMargin() float64`

GetMarchMargin returns the MarchMargin field if non-nil, zero value otherwise.

### GetMarchMarginOk

`func (o *SalesQuota) GetMarchMarginOk() (*float64, bool)`

GetMarchMarginOk returns a tuple with the MarchMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarchMargin

`func (o *SalesQuota) SetMarchMargin(v float64)`

SetMarchMargin sets MarchMargin field to given value.

### HasMarchMargin

`func (o *SalesQuota) HasMarchMargin() bool`

HasMarchMargin returns a boolean if a field has been set.

### SetMarchMarginNil

`func (o *SalesQuota) SetMarchMarginNil(b bool)`

 SetMarchMarginNil sets the value for MarchMargin to be an explicit nil

### UnsetMarchMargin
`func (o *SalesQuota) UnsetMarchMargin()`

UnsetMarchMargin ensures that no value is present for MarchMargin, not even an explicit nil
### GetAprilRevenue

`func (o *SalesQuota) GetAprilRevenue() float64`

GetAprilRevenue returns the AprilRevenue field if non-nil, zero value otherwise.

### GetAprilRevenueOk

`func (o *SalesQuota) GetAprilRevenueOk() (*float64, bool)`

GetAprilRevenueOk returns a tuple with the AprilRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprilRevenue

`func (o *SalesQuota) SetAprilRevenue(v float64)`

SetAprilRevenue sets AprilRevenue field to given value.

### HasAprilRevenue

`func (o *SalesQuota) HasAprilRevenue() bool`

HasAprilRevenue returns a boolean if a field has been set.

### SetAprilRevenueNil

`func (o *SalesQuota) SetAprilRevenueNil(b bool)`

 SetAprilRevenueNil sets the value for AprilRevenue to be an explicit nil

### UnsetAprilRevenue
`func (o *SalesQuota) UnsetAprilRevenue()`

UnsetAprilRevenue ensures that no value is present for AprilRevenue, not even an explicit nil
### GetAprilMargin

`func (o *SalesQuota) GetAprilMargin() float64`

GetAprilMargin returns the AprilMargin field if non-nil, zero value otherwise.

### GetAprilMarginOk

`func (o *SalesQuota) GetAprilMarginOk() (*float64, bool)`

GetAprilMarginOk returns a tuple with the AprilMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprilMargin

`func (o *SalesQuota) SetAprilMargin(v float64)`

SetAprilMargin sets AprilMargin field to given value.

### HasAprilMargin

`func (o *SalesQuota) HasAprilMargin() bool`

HasAprilMargin returns a boolean if a field has been set.

### SetAprilMarginNil

`func (o *SalesQuota) SetAprilMarginNil(b bool)`

 SetAprilMarginNil sets the value for AprilMargin to be an explicit nil

### UnsetAprilMargin
`func (o *SalesQuota) UnsetAprilMargin()`

UnsetAprilMargin ensures that no value is present for AprilMargin, not even an explicit nil
### GetMayRevenue

`func (o *SalesQuota) GetMayRevenue() float64`

GetMayRevenue returns the MayRevenue field if non-nil, zero value otherwise.

### GetMayRevenueOk

`func (o *SalesQuota) GetMayRevenueOk() (*float64, bool)`

GetMayRevenueOk returns a tuple with the MayRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayRevenue

`func (o *SalesQuota) SetMayRevenue(v float64)`

SetMayRevenue sets MayRevenue field to given value.

### HasMayRevenue

`func (o *SalesQuota) HasMayRevenue() bool`

HasMayRevenue returns a boolean if a field has been set.

### SetMayRevenueNil

`func (o *SalesQuota) SetMayRevenueNil(b bool)`

 SetMayRevenueNil sets the value for MayRevenue to be an explicit nil

### UnsetMayRevenue
`func (o *SalesQuota) UnsetMayRevenue()`

UnsetMayRevenue ensures that no value is present for MayRevenue, not even an explicit nil
### GetMayMargin

`func (o *SalesQuota) GetMayMargin() float64`

GetMayMargin returns the MayMargin field if non-nil, zero value otherwise.

### GetMayMarginOk

`func (o *SalesQuota) GetMayMarginOk() (*float64, bool)`

GetMayMarginOk returns a tuple with the MayMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayMargin

`func (o *SalesQuota) SetMayMargin(v float64)`

SetMayMargin sets MayMargin field to given value.

### HasMayMargin

`func (o *SalesQuota) HasMayMargin() bool`

HasMayMargin returns a boolean if a field has been set.

### SetMayMarginNil

`func (o *SalesQuota) SetMayMarginNil(b bool)`

 SetMayMarginNil sets the value for MayMargin to be an explicit nil

### UnsetMayMargin
`func (o *SalesQuota) UnsetMayMargin()`

UnsetMayMargin ensures that no value is present for MayMargin, not even an explicit nil
### GetJuneRevenue

`func (o *SalesQuota) GetJuneRevenue() float64`

GetJuneRevenue returns the JuneRevenue field if non-nil, zero value otherwise.

### GetJuneRevenueOk

`func (o *SalesQuota) GetJuneRevenueOk() (*float64, bool)`

GetJuneRevenueOk returns a tuple with the JuneRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJuneRevenue

`func (o *SalesQuota) SetJuneRevenue(v float64)`

SetJuneRevenue sets JuneRevenue field to given value.

### HasJuneRevenue

`func (o *SalesQuota) HasJuneRevenue() bool`

HasJuneRevenue returns a boolean if a field has been set.

### SetJuneRevenueNil

`func (o *SalesQuota) SetJuneRevenueNil(b bool)`

 SetJuneRevenueNil sets the value for JuneRevenue to be an explicit nil

### UnsetJuneRevenue
`func (o *SalesQuota) UnsetJuneRevenue()`

UnsetJuneRevenue ensures that no value is present for JuneRevenue, not even an explicit nil
### GetJuneMargin

`func (o *SalesQuota) GetJuneMargin() float64`

GetJuneMargin returns the JuneMargin field if non-nil, zero value otherwise.

### GetJuneMarginOk

`func (o *SalesQuota) GetJuneMarginOk() (*float64, bool)`

GetJuneMarginOk returns a tuple with the JuneMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJuneMargin

`func (o *SalesQuota) SetJuneMargin(v float64)`

SetJuneMargin sets JuneMargin field to given value.

### HasJuneMargin

`func (o *SalesQuota) HasJuneMargin() bool`

HasJuneMargin returns a boolean if a field has been set.

### SetJuneMarginNil

`func (o *SalesQuota) SetJuneMarginNil(b bool)`

 SetJuneMarginNil sets the value for JuneMargin to be an explicit nil

### UnsetJuneMargin
`func (o *SalesQuota) UnsetJuneMargin()`

UnsetJuneMargin ensures that no value is present for JuneMargin, not even an explicit nil
### GetJulyRevenue

`func (o *SalesQuota) GetJulyRevenue() float64`

GetJulyRevenue returns the JulyRevenue field if non-nil, zero value otherwise.

### GetJulyRevenueOk

`func (o *SalesQuota) GetJulyRevenueOk() (*float64, bool)`

GetJulyRevenueOk returns a tuple with the JulyRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJulyRevenue

`func (o *SalesQuota) SetJulyRevenue(v float64)`

SetJulyRevenue sets JulyRevenue field to given value.

### HasJulyRevenue

`func (o *SalesQuota) HasJulyRevenue() bool`

HasJulyRevenue returns a boolean if a field has been set.

### SetJulyRevenueNil

`func (o *SalesQuota) SetJulyRevenueNil(b bool)`

 SetJulyRevenueNil sets the value for JulyRevenue to be an explicit nil

### UnsetJulyRevenue
`func (o *SalesQuota) UnsetJulyRevenue()`

UnsetJulyRevenue ensures that no value is present for JulyRevenue, not even an explicit nil
### GetJulyMargin

`func (o *SalesQuota) GetJulyMargin() float64`

GetJulyMargin returns the JulyMargin field if non-nil, zero value otherwise.

### GetJulyMarginOk

`func (o *SalesQuota) GetJulyMarginOk() (*float64, bool)`

GetJulyMarginOk returns a tuple with the JulyMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJulyMargin

`func (o *SalesQuota) SetJulyMargin(v float64)`

SetJulyMargin sets JulyMargin field to given value.

### HasJulyMargin

`func (o *SalesQuota) HasJulyMargin() bool`

HasJulyMargin returns a boolean if a field has been set.

### SetJulyMarginNil

`func (o *SalesQuota) SetJulyMarginNil(b bool)`

 SetJulyMarginNil sets the value for JulyMargin to be an explicit nil

### UnsetJulyMargin
`func (o *SalesQuota) UnsetJulyMargin()`

UnsetJulyMargin ensures that no value is present for JulyMargin, not even an explicit nil
### GetAugustRevenue

`func (o *SalesQuota) GetAugustRevenue() float64`

GetAugustRevenue returns the AugustRevenue field if non-nil, zero value otherwise.

### GetAugustRevenueOk

`func (o *SalesQuota) GetAugustRevenueOk() (*float64, bool)`

GetAugustRevenueOk returns a tuple with the AugustRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAugustRevenue

`func (o *SalesQuota) SetAugustRevenue(v float64)`

SetAugustRevenue sets AugustRevenue field to given value.

### HasAugustRevenue

`func (o *SalesQuota) HasAugustRevenue() bool`

HasAugustRevenue returns a boolean if a field has been set.

### SetAugustRevenueNil

`func (o *SalesQuota) SetAugustRevenueNil(b bool)`

 SetAugustRevenueNil sets the value for AugustRevenue to be an explicit nil

### UnsetAugustRevenue
`func (o *SalesQuota) UnsetAugustRevenue()`

UnsetAugustRevenue ensures that no value is present for AugustRevenue, not even an explicit nil
### GetAugustMargin

`func (o *SalesQuota) GetAugustMargin() float64`

GetAugustMargin returns the AugustMargin field if non-nil, zero value otherwise.

### GetAugustMarginOk

`func (o *SalesQuota) GetAugustMarginOk() (*float64, bool)`

GetAugustMarginOk returns a tuple with the AugustMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAugustMargin

`func (o *SalesQuota) SetAugustMargin(v float64)`

SetAugustMargin sets AugustMargin field to given value.

### HasAugustMargin

`func (o *SalesQuota) HasAugustMargin() bool`

HasAugustMargin returns a boolean if a field has been set.

### SetAugustMarginNil

`func (o *SalesQuota) SetAugustMarginNil(b bool)`

 SetAugustMarginNil sets the value for AugustMargin to be an explicit nil

### UnsetAugustMargin
`func (o *SalesQuota) UnsetAugustMargin()`

UnsetAugustMargin ensures that no value is present for AugustMargin, not even an explicit nil
### GetSeptemberRevenue

`func (o *SalesQuota) GetSeptemberRevenue() float64`

GetSeptemberRevenue returns the SeptemberRevenue field if non-nil, zero value otherwise.

### GetSeptemberRevenueOk

`func (o *SalesQuota) GetSeptemberRevenueOk() (*float64, bool)`

GetSeptemberRevenueOk returns a tuple with the SeptemberRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeptemberRevenue

`func (o *SalesQuota) SetSeptemberRevenue(v float64)`

SetSeptemberRevenue sets SeptemberRevenue field to given value.

### HasSeptemberRevenue

`func (o *SalesQuota) HasSeptemberRevenue() bool`

HasSeptemberRevenue returns a boolean if a field has been set.

### SetSeptemberRevenueNil

`func (o *SalesQuota) SetSeptemberRevenueNil(b bool)`

 SetSeptemberRevenueNil sets the value for SeptemberRevenue to be an explicit nil

### UnsetSeptemberRevenue
`func (o *SalesQuota) UnsetSeptemberRevenue()`

UnsetSeptemberRevenue ensures that no value is present for SeptemberRevenue, not even an explicit nil
### GetSeptemberMargin

`func (o *SalesQuota) GetSeptemberMargin() float64`

GetSeptemberMargin returns the SeptemberMargin field if non-nil, zero value otherwise.

### GetSeptemberMarginOk

`func (o *SalesQuota) GetSeptemberMarginOk() (*float64, bool)`

GetSeptemberMarginOk returns a tuple with the SeptemberMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeptemberMargin

`func (o *SalesQuota) SetSeptemberMargin(v float64)`

SetSeptemberMargin sets SeptemberMargin field to given value.

### HasSeptemberMargin

`func (o *SalesQuota) HasSeptemberMargin() bool`

HasSeptemberMargin returns a boolean if a field has been set.

### SetSeptemberMarginNil

`func (o *SalesQuota) SetSeptemberMarginNil(b bool)`

 SetSeptemberMarginNil sets the value for SeptemberMargin to be an explicit nil

### UnsetSeptemberMargin
`func (o *SalesQuota) UnsetSeptemberMargin()`

UnsetSeptemberMargin ensures that no value is present for SeptemberMargin, not even an explicit nil
### GetOctoberRevenue

`func (o *SalesQuota) GetOctoberRevenue() float64`

GetOctoberRevenue returns the OctoberRevenue field if non-nil, zero value otherwise.

### GetOctoberRevenueOk

`func (o *SalesQuota) GetOctoberRevenueOk() (*float64, bool)`

GetOctoberRevenueOk returns a tuple with the OctoberRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoberRevenue

`func (o *SalesQuota) SetOctoberRevenue(v float64)`

SetOctoberRevenue sets OctoberRevenue field to given value.

### HasOctoberRevenue

`func (o *SalesQuota) HasOctoberRevenue() bool`

HasOctoberRevenue returns a boolean if a field has been set.

### SetOctoberRevenueNil

`func (o *SalesQuota) SetOctoberRevenueNil(b bool)`

 SetOctoberRevenueNil sets the value for OctoberRevenue to be an explicit nil

### UnsetOctoberRevenue
`func (o *SalesQuota) UnsetOctoberRevenue()`

UnsetOctoberRevenue ensures that no value is present for OctoberRevenue, not even an explicit nil
### GetOctoberMargin

`func (o *SalesQuota) GetOctoberMargin() float64`

GetOctoberMargin returns the OctoberMargin field if non-nil, zero value otherwise.

### GetOctoberMarginOk

`func (o *SalesQuota) GetOctoberMarginOk() (*float64, bool)`

GetOctoberMarginOk returns a tuple with the OctoberMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctoberMargin

`func (o *SalesQuota) SetOctoberMargin(v float64)`

SetOctoberMargin sets OctoberMargin field to given value.

### HasOctoberMargin

`func (o *SalesQuota) HasOctoberMargin() bool`

HasOctoberMargin returns a boolean if a field has been set.

### SetOctoberMarginNil

`func (o *SalesQuota) SetOctoberMarginNil(b bool)`

 SetOctoberMarginNil sets the value for OctoberMargin to be an explicit nil

### UnsetOctoberMargin
`func (o *SalesQuota) UnsetOctoberMargin()`

UnsetOctoberMargin ensures that no value is present for OctoberMargin, not even an explicit nil
### GetNovemberRevenue

`func (o *SalesQuota) GetNovemberRevenue() float64`

GetNovemberRevenue returns the NovemberRevenue field if non-nil, zero value otherwise.

### GetNovemberRevenueOk

`func (o *SalesQuota) GetNovemberRevenueOk() (*float64, bool)`

GetNovemberRevenueOk returns a tuple with the NovemberRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNovemberRevenue

`func (o *SalesQuota) SetNovemberRevenue(v float64)`

SetNovemberRevenue sets NovemberRevenue field to given value.

### HasNovemberRevenue

`func (o *SalesQuota) HasNovemberRevenue() bool`

HasNovemberRevenue returns a boolean if a field has been set.

### SetNovemberRevenueNil

`func (o *SalesQuota) SetNovemberRevenueNil(b bool)`

 SetNovemberRevenueNil sets the value for NovemberRevenue to be an explicit nil

### UnsetNovemberRevenue
`func (o *SalesQuota) UnsetNovemberRevenue()`

UnsetNovemberRevenue ensures that no value is present for NovemberRevenue, not even an explicit nil
### GetNovemberMargin

`func (o *SalesQuota) GetNovemberMargin() float64`

GetNovemberMargin returns the NovemberMargin field if non-nil, zero value otherwise.

### GetNovemberMarginOk

`func (o *SalesQuota) GetNovemberMarginOk() (*float64, bool)`

GetNovemberMarginOk returns a tuple with the NovemberMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNovemberMargin

`func (o *SalesQuota) SetNovemberMargin(v float64)`

SetNovemberMargin sets NovemberMargin field to given value.

### HasNovemberMargin

`func (o *SalesQuota) HasNovemberMargin() bool`

HasNovemberMargin returns a boolean if a field has been set.

### SetNovemberMarginNil

`func (o *SalesQuota) SetNovemberMarginNil(b bool)`

 SetNovemberMarginNil sets the value for NovemberMargin to be an explicit nil

### UnsetNovemberMargin
`func (o *SalesQuota) UnsetNovemberMargin()`

UnsetNovemberMargin ensures that no value is present for NovemberMargin, not even an explicit nil
### GetDecemberRevenue

`func (o *SalesQuota) GetDecemberRevenue() float64`

GetDecemberRevenue returns the DecemberRevenue field if non-nil, zero value otherwise.

### GetDecemberRevenueOk

`func (o *SalesQuota) GetDecemberRevenueOk() (*float64, bool)`

GetDecemberRevenueOk returns a tuple with the DecemberRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecemberRevenue

`func (o *SalesQuota) SetDecemberRevenue(v float64)`

SetDecemberRevenue sets DecemberRevenue field to given value.

### HasDecemberRevenue

`func (o *SalesQuota) HasDecemberRevenue() bool`

HasDecemberRevenue returns a boolean if a field has been set.

### SetDecemberRevenueNil

`func (o *SalesQuota) SetDecemberRevenueNil(b bool)`

 SetDecemberRevenueNil sets the value for DecemberRevenue to be an explicit nil

### UnsetDecemberRevenue
`func (o *SalesQuota) UnsetDecemberRevenue()`

UnsetDecemberRevenue ensures that no value is present for DecemberRevenue, not even an explicit nil
### GetDecemberMargin

`func (o *SalesQuota) GetDecemberMargin() float64`

GetDecemberMargin returns the DecemberMargin field if non-nil, zero value otherwise.

### GetDecemberMarginOk

`func (o *SalesQuota) GetDecemberMarginOk() (*float64, bool)`

GetDecemberMarginOk returns a tuple with the DecemberMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecemberMargin

`func (o *SalesQuota) SetDecemberMargin(v float64)`

SetDecemberMargin sets DecemberMargin field to given value.

### HasDecemberMargin

`func (o *SalesQuota) HasDecemberMargin() bool`

HasDecemberMargin returns a boolean if a field has been set.

### SetDecemberMarginNil

`func (o *SalesQuota) SetDecemberMarginNil(b bool)`

 SetDecemberMarginNil sets the value for DecemberMargin to be an explicit nil

### UnsetDecemberMargin
`func (o *SalesQuota) UnsetDecemberMargin()`

UnsetDecemberMargin ensures that no value is present for DecemberMargin, not even an explicit nil
### GetCurrency

`func (o *SalesQuota) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SalesQuota) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SalesQuota) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SalesQuota) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInfo

`func (o *SalesQuota) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SalesQuota) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SalesQuota) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SalesQuota) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


