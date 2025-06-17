# ShipmentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**TrackingUrl** | Pointer to **string** |  Max length: 200; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewShipmentMethod

`func NewShipmentMethod(name string, ) *ShipmentMethod`

NewShipmentMethod instantiates a new ShipmentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentMethodWithDefaults

`func NewShipmentMethodWithDefaults() *ShipmentMethod`

NewShipmentMethodWithDefaults instantiates a new ShipmentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipmentMethod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentMethod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentMethod) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ShipmentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentMethod) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *ShipmentMethod) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ShipmentMethod) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ShipmentMethod) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ShipmentMethod) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ShipmentMethod) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ShipmentMethod) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetTrackingUrl

`func (o *ShipmentMethod) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *ShipmentMethod) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *ShipmentMethod) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *ShipmentMethod) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.

### GetInfo

`func (o *ShipmentMethod) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ShipmentMethod) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ShipmentMethod) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ShipmentMethod) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


