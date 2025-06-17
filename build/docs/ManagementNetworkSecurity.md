# ManagementNetworkSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**Username** | Pointer to **string** |  Max length: 50; | [optional] 
**Password** | Pointer to **string** |  Max length: 50; | [optional] 
**Site** | **string** |  Max length: 100; | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagementNetworkSecurity

`func NewManagementNetworkSecurity(name string, site string, ) *ManagementNetworkSecurity`

NewManagementNetworkSecurity instantiates a new ManagementNetworkSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementNetworkSecurityWithDefaults

`func NewManagementNetworkSecurityWithDefaults() *ManagementNetworkSecurity`

NewManagementNetworkSecurityWithDefaults instantiates a new ManagementNetworkSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagementNetworkSecurity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementNetworkSecurity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementNetworkSecurity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagementNetworkSecurity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManagementNetworkSecurity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementNetworkSecurity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementNetworkSecurity) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *ManagementNetworkSecurity) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ManagementNetworkSecurity) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ManagementNetworkSecurity) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ManagementNetworkSecurity) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ManagementNetworkSecurity) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ManagementNetworkSecurity) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ManagementNetworkSecurity) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ManagementNetworkSecurity) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSite

`func (o *ManagementNetworkSecurity) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ManagementNetworkSecurity) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ManagementNetworkSecurity) SetSite(v string)`

SetSite sets Site field to given value.


### GetInfo

`func (o *ManagementNetworkSecurity) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagementNetworkSecurity) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagementNetworkSecurity) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagementNetworkSecurity) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


