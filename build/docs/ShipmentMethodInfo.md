# ShipmentMethodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**TrackingUrl** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewShipmentMethodInfo

`func NewShipmentMethodInfo() *ShipmentMethodInfo`

NewShipmentMethodInfo instantiates a new ShipmentMethodInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentMethodInfoWithDefaults

`func NewShipmentMethodInfoWithDefaults() *ShipmentMethodInfo`

NewShipmentMethodInfoWithDefaults instantiates a new ShipmentMethodInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipmentMethodInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentMethodInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentMethodInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ShipmentMethodInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ShipmentMethodInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipmentMethodInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipmentMethodInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ShipmentMethodInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ShipmentMethodInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ShipmentMethodInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ShipmentMethodInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ShipmentMethodInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ShipmentMethodInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ShipmentMethodInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetTrackingUrl

`func (o *ShipmentMethodInfo) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *ShipmentMethodInfo) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *ShipmentMethodInfo) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *ShipmentMethodInfo) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.

### GetInfo

`func (o *ShipmentMethodInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ShipmentMethodInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ShipmentMethodInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ShipmentMethodInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


