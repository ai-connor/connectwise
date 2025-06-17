# EmailConnectorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DefaultCompany** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ImapSetup** | Pointer to [**ImapSetupReference**](ImapSetupReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEmailConnectorInfo

`func NewEmailConnectorInfo() *EmailConnectorInfo`

NewEmailConnectorInfo instantiates a new EmailConnectorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConnectorInfoWithDefaults

`func NewEmailConnectorInfoWithDefaults() *EmailConnectorInfo`

NewEmailConnectorInfoWithDefaults instantiates a new EmailConnectorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailConnectorInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailConnectorInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailConnectorInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EmailConnectorInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultCompany

`func (o *EmailConnectorInfo) GetDefaultCompany() CompanyReference`

GetDefaultCompany returns the DefaultCompany field if non-nil, zero value otherwise.

### GetDefaultCompanyOk

`func (o *EmailConnectorInfo) GetDefaultCompanyOk() (*CompanyReference, bool)`

GetDefaultCompanyOk returns a tuple with the DefaultCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCompany

`func (o *EmailConnectorInfo) SetDefaultCompany(v CompanyReference)`

SetDefaultCompany sets DefaultCompany field to given value.

### HasDefaultCompany

`func (o *EmailConnectorInfo) HasDefaultCompany() bool`

HasDefaultCompany returns a boolean if a field has been set.

### GetImapSetup

`func (o *EmailConnectorInfo) GetImapSetup() ImapSetupReference`

GetImapSetup returns the ImapSetup field if non-nil, zero value otherwise.

### GetImapSetupOk

`func (o *EmailConnectorInfo) GetImapSetupOk() (*ImapSetupReference, bool)`

GetImapSetupOk returns a tuple with the ImapSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapSetup

`func (o *EmailConnectorInfo) SetImapSetup(v ImapSetupReference)`

SetImapSetup sets ImapSetup field to given value.

### HasImapSetup

`func (o *EmailConnectorInfo) HasImapSetup() bool`

HasImapSetup returns a boolean if a field has been set.

### GetInfo

`func (o *EmailConnectorInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *EmailConnectorInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *EmailConnectorInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *EmailConnectorInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


