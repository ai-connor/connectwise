# M365ContactSyncProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IncludeExcludeType** | Pointer to **NullableString** |  | [optional] 
**PropertyType** | Pointer to **NullableString** |  | [optional] 
**ExcludeIncludeFlag** | Pointer to **bool** |  | [optional] 
**WildCard** | Pointer to **string** |  | [optional] 
**CompanyRecID** | Pointer to **int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewM365ContactSyncProperty

`func NewM365ContactSyncProperty() *M365ContactSyncProperty`

NewM365ContactSyncProperty instantiates a new M365ContactSyncProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewM365ContactSyncPropertyWithDefaults

`func NewM365ContactSyncPropertyWithDefaults() *M365ContactSyncProperty`

NewM365ContactSyncPropertyWithDefaults instantiates a new M365ContactSyncProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *M365ContactSyncProperty) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *M365ContactSyncProperty) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *M365ContactSyncProperty) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *M365ContactSyncProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludeExcludeType

`func (o *M365ContactSyncProperty) GetIncludeExcludeType() string`

GetIncludeExcludeType returns the IncludeExcludeType field if non-nil, zero value otherwise.

### GetIncludeExcludeTypeOk

`func (o *M365ContactSyncProperty) GetIncludeExcludeTypeOk() (*string, bool)`

GetIncludeExcludeTypeOk returns a tuple with the IncludeExcludeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExcludeType

`func (o *M365ContactSyncProperty) SetIncludeExcludeType(v string)`

SetIncludeExcludeType sets IncludeExcludeType field to given value.

### HasIncludeExcludeType

`func (o *M365ContactSyncProperty) HasIncludeExcludeType() bool`

HasIncludeExcludeType returns a boolean if a field has been set.

### SetIncludeExcludeTypeNil

`func (o *M365ContactSyncProperty) SetIncludeExcludeTypeNil(b bool)`

 SetIncludeExcludeTypeNil sets the value for IncludeExcludeType to be an explicit nil

### UnsetIncludeExcludeType
`func (o *M365ContactSyncProperty) UnsetIncludeExcludeType()`

UnsetIncludeExcludeType ensures that no value is present for IncludeExcludeType, not even an explicit nil
### GetPropertyType

`func (o *M365ContactSyncProperty) GetPropertyType() string`

GetPropertyType returns the PropertyType field if non-nil, zero value otherwise.

### GetPropertyTypeOk

`func (o *M365ContactSyncProperty) GetPropertyTypeOk() (*string, bool)`

GetPropertyTypeOk returns a tuple with the PropertyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyType

`func (o *M365ContactSyncProperty) SetPropertyType(v string)`

SetPropertyType sets PropertyType field to given value.

### HasPropertyType

`func (o *M365ContactSyncProperty) HasPropertyType() bool`

HasPropertyType returns a boolean if a field has been set.

### SetPropertyTypeNil

`func (o *M365ContactSyncProperty) SetPropertyTypeNil(b bool)`

 SetPropertyTypeNil sets the value for PropertyType to be an explicit nil

### UnsetPropertyType
`func (o *M365ContactSyncProperty) UnsetPropertyType()`

UnsetPropertyType ensures that no value is present for PropertyType, not even an explicit nil
### GetExcludeIncludeFlag

`func (o *M365ContactSyncProperty) GetExcludeIncludeFlag() bool`

GetExcludeIncludeFlag returns the ExcludeIncludeFlag field if non-nil, zero value otherwise.

### GetExcludeIncludeFlagOk

`func (o *M365ContactSyncProperty) GetExcludeIncludeFlagOk() (*bool, bool)`

GetExcludeIncludeFlagOk returns a tuple with the ExcludeIncludeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeIncludeFlag

`func (o *M365ContactSyncProperty) SetExcludeIncludeFlag(v bool)`

SetExcludeIncludeFlag sets ExcludeIncludeFlag field to given value.

### HasExcludeIncludeFlag

`func (o *M365ContactSyncProperty) HasExcludeIncludeFlag() bool`

HasExcludeIncludeFlag returns a boolean if a field has been set.

### GetWildCard

`func (o *M365ContactSyncProperty) GetWildCard() string`

GetWildCard returns the WildCard field if non-nil, zero value otherwise.

### GetWildCardOk

`func (o *M365ContactSyncProperty) GetWildCardOk() (*string, bool)`

GetWildCardOk returns a tuple with the WildCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildCard

`func (o *M365ContactSyncProperty) SetWildCard(v string)`

SetWildCard sets WildCard field to given value.

### HasWildCard

`func (o *M365ContactSyncProperty) HasWildCard() bool`

HasWildCard returns a boolean if a field has been set.

### GetCompanyRecID

`func (o *M365ContactSyncProperty) GetCompanyRecID() int32`

GetCompanyRecID returns the CompanyRecID field if non-nil, zero value otherwise.

### GetCompanyRecIDOk

`func (o *M365ContactSyncProperty) GetCompanyRecIDOk() (*int32, bool)`

GetCompanyRecIDOk returns a tuple with the CompanyRecID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyRecID

`func (o *M365ContactSyncProperty) SetCompanyRecID(v int32)`

SetCompanyRecID sets CompanyRecID field to given value.

### HasCompanyRecID

`func (o *M365ContactSyncProperty) HasCompanyRecID() bool`

HasCompanyRecID returns a boolean if a field has been set.

### GetInfo

`func (o *M365ContactSyncProperty) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *M365ContactSyncProperty) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *M365ContactSyncProperty) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *M365ContactSyncProperty) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


