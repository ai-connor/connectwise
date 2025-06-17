# LdapConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 100; | 
**Server** | **string** | FQDN of the Server. Max length: 200; | 
**Domain** | **string** | Domain Name of the server. Max length: 50; | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLdapConfiguration

`func NewLdapConfiguration(name string, server string, domain string, ) *LdapConfiguration`

NewLdapConfiguration instantiates a new LdapConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConfigurationWithDefaults

`func NewLdapConfigurationWithDefaults() *LdapConfiguration`

NewLdapConfigurationWithDefaults instantiates a new LdapConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LdapConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LdapConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *LdapConfiguration) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LdapConfiguration) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LdapConfiguration) SetServer(v string)`

SetServer sets Server field to given value.


### GetDomain

`func (o *LdapConfiguration) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *LdapConfiguration) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *LdapConfiguration) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetInfo

`func (o *LdapConfiguration) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *LdapConfiguration) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *LdapConfiguration) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *LdapConfiguration) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


