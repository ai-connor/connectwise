# PortalReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PortalConfiguration** | Pointer to [**PortalConfigurationReference**](PortalConfigurationReference.md) |  | [optional] 
**Name** | **string** |  Max length: 255; | 
**Url** | **string** |  Max length: 255; | 
**OpenSameWindowFlag** | Pointer to **NullableBool** |  | [optional] 
**CustomFlag** | Pointer to **NullableBool** |  | [optional] 
**DisplayFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalReport

`func NewPortalReport(name string, url string, ) *PortalReport`

NewPortalReport instantiates a new PortalReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalReportWithDefaults

`func NewPortalReportWithDefaults() *PortalReport`

NewPortalReportWithDefaults instantiates a new PortalReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalReport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalReport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalReport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortalConfiguration

`func (o *PortalReport) GetPortalConfiguration() PortalConfigurationReference`

GetPortalConfiguration returns the PortalConfiguration field if non-nil, zero value otherwise.

### GetPortalConfigurationOk

`func (o *PortalReport) GetPortalConfigurationOk() (*PortalConfigurationReference, bool)`

GetPortalConfigurationOk returns a tuple with the PortalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalConfiguration

`func (o *PortalReport) SetPortalConfiguration(v PortalConfigurationReference)`

SetPortalConfiguration sets PortalConfiguration field to given value.

### HasPortalConfiguration

`func (o *PortalReport) HasPortalConfiguration() bool`

HasPortalConfiguration returns a boolean if a field has been set.

### GetName

`func (o *PortalReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortalReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortalReport) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *PortalReport) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PortalReport) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PortalReport) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetOpenSameWindowFlag

`func (o *PortalReport) GetOpenSameWindowFlag() bool`

GetOpenSameWindowFlag returns the OpenSameWindowFlag field if non-nil, zero value otherwise.

### GetOpenSameWindowFlagOk

`func (o *PortalReport) GetOpenSameWindowFlagOk() (*bool, bool)`

GetOpenSameWindowFlagOk returns a tuple with the OpenSameWindowFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSameWindowFlag

`func (o *PortalReport) SetOpenSameWindowFlag(v bool)`

SetOpenSameWindowFlag sets OpenSameWindowFlag field to given value.

### HasOpenSameWindowFlag

`func (o *PortalReport) HasOpenSameWindowFlag() bool`

HasOpenSameWindowFlag returns a boolean if a field has been set.

### SetOpenSameWindowFlagNil

`func (o *PortalReport) SetOpenSameWindowFlagNil(b bool)`

 SetOpenSameWindowFlagNil sets the value for OpenSameWindowFlag to be an explicit nil

### UnsetOpenSameWindowFlag
`func (o *PortalReport) UnsetOpenSameWindowFlag()`

UnsetOpenSameWindowFlag ensures that no value is present for OpenSameWindowFlag, not even an explicit nil
### GetCustomFlag

`func (o *PortalReport) GetCustomFlag() bool`

GetCustomFlag returns the CustomFlag field if non-nil, zero value otherwise.

### GetCustomFlagOk

`func (o *PortalReport) GetCustomFlagOk() (*bool, bool)`

GetCustomFlagOk returns a tuple with the CustomFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFlag

`func (o *PortalReport) SetCustomFlag(v bool)`

SetCustomFlag sets CustomFlag field to given value.

### HasCustomFlag

`func (o *PortalReport) HasCustomFlag() bool`

HasCustomFlag returns a boolean if a field has been set.

### SetCustomFlagNil

`func (o *PortalReport) SetCustomFlagNil(b bool)`

 SetCustomFlagNil sets the value for CustomFlag to be an explicit nil

### UnsetCustomFlag
`func (o *PortalReport) UnsetCustomFlag()`

UnsetCustomFlag ensures that no value is present for CustomFlag, not even an explicit nil
### GetDisplayFlag

`func (o *PortalReport) GetDisplayFlag() bool`

GetDisplayFlag returns the DisplayFlag field if non-nil, zero value otherwise.

### GetDisplayFlagOk

`func (o *PortalReport) GetDisplayFlagOk() (*bool, bool)`

GetDisplayFlagOk returns a tuple with the DisplayFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayFlag

`func (o *PortalReport) SetDisplayFlag(v bool)`

SetDisplayFlag sets DisplayFlag field to given value.

### HasDisplayFlag

`func (o *PortalReport) HasDisplayFlag() bool`

HasDisplayFlag returns a boolean if a field has been set.

### SetDisplayFlagNil

`func (o *PortalReport) SetDisplayFlagNil(b bool)`

 SetDisplayFlagNil sets the value for DisplayFlag to be an explicit nil

### UnsetDisplayFlag
`func (o *PortalReport) UnsetDisplayFlag()`

UnsetDisplayFlag ensures that no value is present for DisplayFlag, not even an explicit nil
### GetInfo

`func (o *PortalReport) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalReport) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalReport) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalReport) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


