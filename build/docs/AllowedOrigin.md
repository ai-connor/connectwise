# AllowedOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Origin** | **string** |  Max length: 2000; | 
**Description** | **string** |  Max length: 2000; | 
**LastUpdateUtc** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAllowedOrigin

`func NewAllowedOrigin(origin string, description string, ) *AllowedOrigin`

NewAllowedOrigin instantiates a new AllowedOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedOriginWithDefaults

`func NewAllowedOriginWithDefaults() *AllowedOrigin`

NewAllowedOriginWithDefaults instantiates a new AllowedOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowedOrigin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedOrigin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedOrigin) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AllowedOrigin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrigin

`func (o *AllowedOrigin) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AllowedOrigin) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AllowedOrigin) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetDescription

`func (o *AllowedOrigin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllowedOrigin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllowedOrigin) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLastUpdateUtc

`func (o *AllowedOrigin) GetLastUpdateUtc() time.Time`

GetLastUpdateUtc returns the LastUpdateUtc field if non-nil, zero value otherwise.

### GetLastUpdateUtcOk

`func (o *AllowedOrigin) GetLastUpdateUtcOk() (*time.Time, bool)`

GetLastUpdateUtcOk returns a tuple with the LastUpdateUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateUtc

`func (o *AllowedOrigin) SetLastUpdateUtc(v time.Time)`

SetLastUpdateUtc sets LastUpdateUtc field to given value.

### HasLastUpdateUtc

`func (o *AllowedOrigin) HasLastUpdateUtc() bool`

HasLastUpdateUtc returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AllowedOrigin) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AllowedOrigin) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AllowedOrigin) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AllowedOrigin) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetInfo

`func (o *AllowedOrigin) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AllowedOrigin) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AllowedOrigin) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AllowedOrigin) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


