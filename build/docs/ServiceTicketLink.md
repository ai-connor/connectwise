# ServiceTicketLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**EnabledFlag** | Pointer to **NullableBool** |  | [optional] 
**LinkText** | **string** |  Max length: 50; | 
**Url** | **string** |  Max length: 1000; | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceTicketLink

`func NewServiceTicketLink(name string, linkText string, url string, ) *ServiceTicketLink`

NewServiceTicketLink instantiates a new ServiceTicketLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTicketLinkWithDefaults

`func NewServiceTicketLinkWithDefaults() *ServiceTicketLink`

NewServiceTicketLinkWithDefaults instantiates a new ServiceTicketLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTicketLink) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTicketLink) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTicketLink) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceTicketLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceTicketLink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceTicketLink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceTicketLink) SetName(v string)`

SetName sets Name field to given value.


### GetEnabledFlag

`func (o *ServiceTicketLink) GetEnabledFlag() bool`

GetEnabledFlag returns the EnabledFlag field if non-nil, zero value otherwise.

### GetEnabledFlagOk

`func (o *ServiceTicketLink) GetEnabledFlagOk() (*bool, bool)`

GetEnabledFlagOk returns a tuple with the EnabledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFlag

`func (o *ServiceTicketLink) SetEnabledFlag(v bool)`

SetEnabledFlag sets EnabledFlag field to given value.

### HasEnabledFlag

`func (o *ServiceTicketLink) HasEnabledFlag() bool`

HasEnabledFlag returns a boolean if a field has been set.

### SetEnabledFlagNil

`func (o *ServiceTicketLink) SetEnabledFlagNil(b bool)`

 SetEnabledFlagNil sets the value for EnabledFlag to be an explicit nil

### UnsetEnabledFlag
`func (o *ServiceTicketLink) UnsetEnabledFlag()`

UnsetEnabledFlag ensures that no value is present for EnabledFlag, not even an explicit nil
### GetLinkText

`func (o *ServiceTicketLink) GetLinkText() string`

GetLinkText returns the LinkText field if non-nil, zero value otherwise.

### GetLinkTextOk

`func (o *ServiceTicketLink) GetLinkTextOk() (*string, bool)`

GetLinkTextOk returns a tuple with the LinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkText

`func (o *ServiceTicketLink) SetLinkText(v string)`

SetLinkText sets LinkText field to given value.


### GetUrl

`func (o *ServiceTicketLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceTicketLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceTicketLink) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetInfo

`func (o *ServiceTicketLink) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceTicketLink) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceTicketLink) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceTicketLink) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


