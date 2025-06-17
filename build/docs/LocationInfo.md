# LocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LocationFlag** | Pointer to **bool** |  | [optional] 
**StructureLevel** | Pointer to [**CorporateStructureLevelReference**](CorporateStructureLevelReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLocationInfo

`func NewLocationInfo() *LocationInfo`

NewLocationInfo instantiates a new LocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationInfoWithDefaults

`func NewLocationInfoWithDefaults() *LocationInfo`

NewLocationInfoWithDefaults instantiates a new LocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LocationInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LocationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocationInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocationFlag

`func (o *LocationInfo) GetLocationFlag() bool`

GetLocationFlag returns the LocationFlag field if non-nil, zero value otherwise.

### GetLocationFlagOk

`func (o *LocationInfo) GetLocationFlagOk() (*bool, bool)`

GetLocationFlagOk returns a tuple with the LocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationFlag

`func (o *LocationInfo) SetLocationFlag(v bool)`

SetLocationFlag sets LocationFlag field to given value.

### HasLocationFlag

`func (o *LocationInfo) HasLocationFlag() bool`

HasLocationFlag returns a boolean if a field has been set.

### GetStructureLevel

`func (o *LocationInfo) GetStructureLevel() CorporateStructureLevelReference`

GetStructureLevel returns the StructureLevel field if non-nil, zero value otherwise.

### GetStructureLevelOk

`func (o *LocationInfo) GetStructureLevelOk() (*CorporateStructureLevelReference, bool)`

GetStructureLevelOk returns a tuple with the StructureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureLevel

`func (o *LocationInfo) SetStructureLevel(v CorporateStructureLevelReference)`

SetStructureLevel sets StructureLevel field to given value.

### HasStructureLevel

`func (o *LocationInfo) HasStructureLevel() bool`

HasStructureLevel returns a boolean if a field has been set.

### GetInfo

`func (o *LocationInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *LocationInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *LocationInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *LocationInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


