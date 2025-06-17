# PortalConfigurationPasswordEmailSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ValidPasswordEmailUseCustomEmailFlag** | Pointer to **NullableBool** |  | [optional] 
**ValidPasswordEmailFromFirstName** | Pointer to **string** |  | [optional] 
**ValidPasswordEmailFromLastName** | Pointer to **string** |  | [optional] 
**ValidPasswordEmailFromEmail** | Pointer to **string** | Gets or sets             required when validPasswordEmailUseCustomEmailFlag is true. | [optional] 
**ValidPasswordEmailSubject** | Pointer to **string** |  | [optional] 
**ValidPasswordEmailBody** | Pointer to **string** |  | [optional] 
**InvalidPasswordEmailUseCustomEmailFlag** | Pointer to **NullableBool** |  | [optional] 
**InvalidPasswordEmailFromFirstName** | Pointer to **string** |  | [optional] 
**InvalidPasswordEmailFromLastName** | Pointer to **string** |  | [optional] 
**InvalidPasswordEmailFromEmail** | Pointer to **string** | Gets or sets             required when invalidPasswordEmailUseCustomEmailFlag is true. | [optional] 
**InvalidPasswordEmailSubject** | Pointer to **string** |  | [optional] 
**InvalidPasswordEmailBody** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalConfigurationPasswordEmailSetup

`func NewPortalConfigurationPasswordEmailSetup() *PortalConfigurationPasswordEmailSetup`

NewPortalConfigurationPasswordEmailSetup instantiates a new PortalConfigurationPasswordEmailSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalConfigurationPasswordEmailSetupWithDefaults

`func NewPortalConfigurationPasswordEmailSetupWithDefaults() *PortalConfigurationPasswordEmailSetup`

NewPortalConfigurationPasswordEmailSetupWithDefaults instantiates a new PortalConfigurationPasswordEmailSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalConfigurationPasswordEmailSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalConfigurationPasswordEmailSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalConfigurationPasswordEmailSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalConfigurationPasswordEmailSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValidPasswordEmailUseCustomEmailFlag

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailUseCustomEmailFlag() bool`

GetValidPasswordEmailUseCustomEmailFlag returns the ValidPasswordEmailUseCustomEmailFlag field if non-nil, zero value otherwise.

### GetValidPasswordEmailUseCustomEmailFlagOk

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailUseCustomEmailFlagOk() (*bool, bool)`

GetValidPasswordEmailUseCustomEmailFlagOk returns a tuple with the ValidPasswordEmailUseCustomEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPasswordEmailUseCustomEmailFlag

`func (o *PortalConfigurationPasswordEmailSetup) SetValidPasswordEmailUseCustomEmailFlag(v bool)`

SetValidPasswordEmailUseCustomEmailFlag sets ValidPasswordEmailUseCustomEmailFlag field to given value.

### HasValidPasswordEmailUseCustomEmailFlag

`func (o *PortalConfigurationPasswordEmailSetup) HasValidPasswordEmailUseCustomEmailFlag() bool`

HasValidPasswordEmailUseCustomEmailFlag returns a boolean if a field has been set.

### SetValidPasswordEmailUseCustomEmailFlagNil

`func (o *PortalConfigurationPasswordEmailSetup) SetValidPasswordEmailUseCustomEmailFlagNil(b bool)`

 SetValidPasswordEmailUseCustomEmailFlagNil sets the value for ValidPasswordEmailUseCustomEmailFlag to be an explicit nil

### UnsetValidPasswordEmailUseCustomEmailFlag
`func (o *PortalConfigurationPasswordEmailSetup) UnsetValidPasswordEmailUseCustomEmailFlag()`

UnsetValidPasswordEmailUseCustomEmailFlag ensures that no value is present for ValidPasswordEmailUseCustomEmailFlag, not even an explicit nil
### GetValidPasswordEmailFromFirstName

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailFromFirstName() string`

GetValidPasswordEmailFromFirstName returns the ValidPasswordEmailFromFirstName field if non-nil, zero value otherwise.

### GetValidPasswordEmailFromFirstNameOk

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailFromFirstNameOk() (*string, bool)`

GetValidPasswordEmailFromFirstNameOk returns a tuple with the ValidPasswordEmailFromFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPasswordEmailFromFirstName

`func (o *PortalConfigurationPasswordEmailSetup) SetValidPasswordEmailFromFirstName(v string)`

SetValidPasswordEmailFromFirstName sets ValidPasswordEmailFromFirstName field to given value.

### HasValidPasswordEmailFromFirstName

`func (o *PortalConfigurationPasswordEmailSetup) HasValidPasswordEmailFromFirstName() bool`

HasValidPasswordEmailFromFirstName returns a boolean if a field has been set.

### GetValidPasswordEmailFromLastName

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailFromLastName() string`

GetValidPasswordEmailFromLastName returns the ValidPasswordEmailFromLastName field if non-nil, zero value otherwise.

### GetValidPasswordEmailFromLastNameOk

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailFromLastNameOk() (*string, bool)`

GetValidPasswordEmailFromLastNameOk returns a tuple with the ValidPasswordEmailFromLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPasswordEmailFromLastName

`func (o *PortalConfigurationPasswordEmailSetup) SetValidPasswordEmailFromLastName(v string)`

SetValidPasswordEmailFromLastName sets ValidPasswordEmailFromLastName field to given value.

### HasValidPasswordEmailFromLastName

`func (o *PortalConfigurationPasswordEmailSetup) HasValidPasswordEmailFromLastName() bool`

HasValidPasswordEmailFromLastName returns a boolean if a field has been set.

### GetValidPasswordEmailFromEmail

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailFromEmail() string`

GetValidPasswordEmailFromEmail returns the ValidPasswordEmailFromEmail field if non-nil, zero value otherwise.

### GetValidPasswordEmailFromEmailOk

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailFromEmailOk() (*string, bool)`

GetValidPasswordEmailFromEmailOk returns a tuple with the ValidPasswordEmailFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPasswordEmailFromEmail

`func (o *PortalConfigurationPasswordEmailSetup) SetValidPasswordEmailFromEmail(v string)`

SetValidPasswordEmailFromEmail sets ValidPasswordEmailFromEmail field to given value.

### HasValidPasswordEmailFromEmail

`func (o *PortalConfigurationPasswordEmailSetup) HasValidPasswordEmailFromEmail() bool`

HasValidPasswordEmailFromEmail returns a boolean if a field has been set.

### GetValidPasswordEmailSubject

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailSubject() string`

GetValidPasswordEmailSubject returns the ValidPasswordEmailSubject field if non-nil, zero value otherwise.

### GetValidPasswordEmailSubjectOk

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailSubjectOk() (*string, bool)`

GetValidPasswordEmailSubjectOk returns a tuple with the ValidPasswordEmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPasswordEmailSubject

`func (o *PortalConfigurationPasswordEmailSetup) SetValidPasswordEmailSubject(v string)`

SetValidPasswordEmailSubject sets ValidPasswordEmailSubject field to given value.

### HasValidPasswordEmailSubject

`func (o *PortalConfigurationPasswordEmailSetup) HasValidPasswordEmailSubject() bool`

HasValidPasswordEmailSubject returns a boolean if a field has been set.

### GetValidPasswordEmailBody

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailBody() string`

GetValidPasswordEmailBody returns the ValidPasswordEmailBody field if non-nil, zero value otherwise.

### GetValidPasswordEmailBodyOk

`func (o *PortalConfigurationPasswordEmailSetup) GetValidPasswordEmailBodyOk() (*string, bool)`

GetValidPasswordEmailBodyOk returns a tuple with the ValidPasswordEmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidPasswordEmailBody

`func (o *PortalConfigurationPasswordEmailSetup) SetValidPasswordEmailBody(v string)`

SetValidPasswordEmailBody sets ValidPasswordEmailBody field to given value.

### HasValidPasswordEmailBody

`func (o *PortalConfigurationPasswordEmailSetup) HasValidPasswordEmailBody() bool`

HasValidPasswordEmailBody returns a boolean if a field has been set.

### GetInvalidPasswordEmailUseCustomEmailFlag

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailUseCustomEmailFlag() bool`

GetInvalidPasswordEmailUseCustomEmailFlag returns the InvalidPasswordEmailUseCustomEmailFlag field if non-nil, zero value otherwise.

### GetInvalidPasswordEmailUseCustomEmailFlagOk

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailUseCustomEmailFlagOk() (*bool, bool)`

GetInvalidPasswordEmailUseCustomEmailFlagOk returns a tuple with the InvalidPasswordEmailUseCustomEmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPasswordEmailUseCustomEmailFlag

`func (o *PortalConfigurationPasswordEmailSetup) SetInvalidPasswordEmailUseCustomEmailFlag(v bool)`

SetInvalidPasswordEmailUseCustomEmailFlag sets InvalidPasswordEmailUseCustomEmailFlag field to given value.

### HasInvalidPasswordEmailUseCustomEmailFlag

`func (o *PortalConfigurationPasswordEmailSetup) HasInvalidPasswordEmailUseCustomEmailFlag() bool`

HasInvalidPasswordEmailUseCustomEmailFlag returns a boolean if a field has been set.

### SetInvalidPasswordEmailUseCustomEmailFlagNil

`func (o *PortalConfigurationPasswordEmailSetup) SetInvalidPasswordEmailUseCustomEmailFlagNil(b bool)`

 SetInvalidPasswordEmailUseCustomEmailFlagNil sets the value for InvalidPasswordEmailUseCustomEmailFlag to be an explicit nil

### UnsetInvalidPasswordEmailUseCustomEmailFlag
`func (o *PortalConfigurationPasswordEmailSetup) UnsetInvalidPasswordEmailUseCustomEmailFlag()`

UnsetInvalidPasswordEmailUseCustomEmailFlag ensures that no value is present for InvalidPasswordEmailUseCustomEmailFlag, not even an explicit nil
### GetInvalidPasswordEmailFromFirstName

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailFromFirstName() string`

GetInvalidPasswordEmailFromFirstName returns the InvalidPasswordEmailFromFirstName field if non-nil, zero value otherwise.

### GetInvalidPasswordEmailFromFirstNameOk

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailFromFirstNameOk() (*string, bool)`

GetInvalidPasswordEmailFromFirstNameOk returns a tuple with the InvalidPasswordEmailFromFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPasswordEmailFromFirstName

`func (o *PortalConfigurationPasswordEmailSetup) SetInvalidPasswordEmailFromFirstName(v string)`

SetInvalidPasswordEmailFromFirstName sets InvalidPasswordEmailFromFirstName field to given value.

### HasInvalidPasswordEmailFromFirstName

`func (o *PortalConfigurationPasswordEmailSetup) HasInvalidPasswordEmailFromFirstName() bool`

HasInvalidPasswordEmailFromFirstName returns a boolean if a field has been set.

### GetInvalidPasswordEmailFromLastName

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailFromLastName() string`

GetInvalidPasswordEmailFromLastName returns the InvalidPasswordEmailFromLastName field if non-nil, zero value otherwise.

### GetInvalidPasswordEmailFromLastNameOk

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailFromLastNameOk() (*string, bool)`

GetInvalidPasswordEmailFromLastNameOk returns a tuple with the InvalidPasswordEmailFromLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPasswordEmailFromLastName

`func (o *PortalConfigurationPasswordEmailSetup) SetInvalidPasswordEmailFromLastName(v string)`

SetInvalidPasswordEmailFromLastName sets InvalidPasswordEmailFromLastName field to given value.

### HasInvalidPasswordEmailFromLastName

`func (o *PortalConfigurationPasswordEmailSetup) HasInvalidPasswordEmailFromLastName() bool`

HasInvalidPasswordEmailFromLastName returns a boolean if a field has been set.

### GetInvalidPasswordEmailFromEmail

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailFromEmail() string`

GetInvalidPasswordEmailFromEmail returns the InvalidPasswordEmailFromEmail field if non-nil, zero value otherwise.

### GetInvalidPasswordEmailFromEmailOk

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailFromEmailOk() (*string, bool)`

GetInvalidPasswordEmailFromEmailOk returns a tuple with the InvalidPasswordEmailFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPasswordEmailFromEmail

`func (o *PortalConfigurationPasswordEmailSetup) SetInvalidPasswordEmailFromEmail(v string)`

SetInvalidPasswordEmailFromEmail sets InvalidPasswordEmailFromEmail field to given value.

### HasInvalidPasswordEmailFromEmail

`func (o *PortalConfigurationPasswordEmailSetup) HasInvalidPasswordEmailFromEmail() bool`

HasInvalidPasswordEmailFromEmail returns a boolean if a field has been set.

### GetInvalidPasswordEmailSubject

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailSubject() string`

GetInvalidPasswordEmailSubject returns the InvalidPasswordEmailSubject field if non-nil, zero value otherwise.

### GetInvalidPasswordEmailSubjectOk

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailSubjectOk() (*string, bool)`

GetInvalidPasswordEmailSubjectOk returns a tuple with the InvalidPasswordEmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPasswordEmailSubject

`func (o *PortalConfigurationPasswordEmailSetup) SetInvalidPasswordEmailSubject(v string)`

SetInvalidPasswordEmailSubject sets InvalidPasswordEmailSubject field to given value.

### HasInvalidPasswordEmailSubject

`func (o *PortalConfigurationPasswordEmailSetup) HasInvalidPasswordEmailSubject() bool`

HasInvalidPasswordEmailSubject returns a boolean if a field has been set.

### GetInvalidPasswordEmailBody

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailBody() string`

GetInvalidPasswordEmailBody returns the InvalidPasswordEmailBody field if non-nil, zero value otherwise.

### GetInvalidPasswordEmailBodyOk

`func (o *PortalConfigurationPasswordEmailSetup) GetInvalidPasswordEmailBodyOk() (*string, bool)`

GetInvalidPasswordEmailBodyOk returns a tuple with the InvalidPasswordEmailBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidPasswordEmailBody

`func (o *PortalConfigurationPasswordEmailSetup) SetInvalidPasswordEmailBody(v string)`

SetInvalidPasswordEmailBody sets InvalidPasswordEmailBody field to given value.

### HasInvalidPasswordEmailBody

`func (o *PortalConfigurationPasswordEmailSetup) HasInvalidPasswordEmailBody() bool`

HasInvalidPasswordEmailBody returns a boolean if a field has been set.

### GetInfo

`func (o *PortalConfigurationPasswordEmailSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalConfigurationPasswordEmailSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalConfigurationPasswordEmailSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalConfigurationPasswordEmailSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


