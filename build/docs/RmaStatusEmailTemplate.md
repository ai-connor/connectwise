# RmaStatusEmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**RmaStatusReference**](RmaStatusReference.md) |  | [optional] 
**UseSenderFlag** | Pointer to **NullableBool** |  | [optional] 
**FirstName** | Pointer to **string** |  Max length: 100; | [optional] 
**LastName** | Pointer to **string** |  Max length: 100; | [optional] 
**EmailAddress** | Pointer to **string** |  Max length: 100; | [optional] 
**Subject** | **string** |  Max length: 200; | 
**Body** | **string** |  | 
**CopySenderFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewRmaStatusEmailTemplate

`func NewRmaStatusEmailTemplate(subject string, body string, ) *RmaStatusEmailTemplate`

NewRmaStatusEmailTemplate instantiates a new RmaStatusEmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRmaStatusEmailTemplateWithDefaults

`func NewRmaStatusEmailTemplateWithDefaults() *RmaStatusEmailTemplate`

NewRmaStatusEmailTemplateWithDefaults instantiates a new RmaStatusEmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RmaStatusEmailTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RmaStatusEmailTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RmaStatusEmailTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RmaStatusEmailTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *RmaStatusEmailTemplate) GetStatus() RmaStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RmaStatusEmailTemplate) GetStatusOk() (*RmaStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RmaStatusEmailTemplate) SetStatus(v RmaStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RmaStatusEmailTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUseSenderFlag

`func (o *RmaStatusEmailTemplate) GetUseSenderFlag() bool`

GetUseSenderFlag returns the UseSenderFlag field if non-nil, zero value otherwise.

### GetUseSenderFlagOk

`func (o *RmaStatusEmailTemplate) GetUseSenderFlagOk() (*bool, bool)`

GetUseSenderFlagOk returns a tuple with the UseSenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSenderFlag

`func (o *RmaStatusEmailTemplate) SetUseSenderFlag(v bool)`

SetUseSenderFlag sets UseSenderFlag field to given value.

### HasUseSenderFlag

`func (o *RmaStatusEmailTemplate) HasUseSenderFlag() bool`

HasUseSenderFlag returns a boolean if a field has been set.

### SetUseSenderFlagNil

`func (o *RmaStatusEmailTemplate) SetUseSenderFlagNil(b bool)`

 SetUseSenderFlagNil sets the value for UseSenderFlag to be an explicit nil

### UnsetUseSenderFlag
`func (o *RmaStatusEmailTemplate) UnsetUseSenderFlag()`

UnsetUseSenderFlag ensures that no value is present for UseSenderFlag, not even an explicit nil
### GetFirstName

`func (o *RmaStatusEmailTemplate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *RmaStatusEmailTemplate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *RmaStatusEmailTemplate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *RmaStatusEmailTemplate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *RmaStatusEmailTemplate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *RmaStatusEmailTemplate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *RmaStatusEmailTemplate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *RmaStatusEmailTemplate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *RmaStatusEmailTemplate) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *RmaStatusEmailTemplate) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *RmaStatusEmailTemplate) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *RmaStatusEmailTemplate) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetSubject

`func (o *RmaStatusEmailTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RmaStatusEmailTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RmaStatusEmailTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetBody

`func (o *RmaStatusEmailTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RmaStatusEmailTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RmaStatusEmailTemplate) SetBody(v string)`

SetBody sets Body field to given value.


### GetCopySenderFlag

`func (o *RmaStatusEmailTemplate) GetCopySenderFlag() bool`

GetCopySenderFlag returns the CopySenderFlag field if non-nil, zero value otherwise.

### GetCopySenderFlagOk

`func (o *RmaStatusEmailTemplate) GetCopySenderFlagOk() (*bool, bool)`

GetCopySenderFlagOk returns a tuple with the CopySenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySenderFlag

`func (o *RmaStatusEmailTemplate) SetCopySenderFlag(v bool)`

SetCopySenderFlag sets CopySenderFlag field to given value.

### HasCopySenderFlag

`func (o *RmaStatusEmailTemplate) HasCopySenderFlag() bool`

HasCopySenderFlag returns a boolean if a field has been set.

### SetCopySenderFlagNil

`func (o *RmaStatusEmailTemplate) SetCopySenderFlagNil(b bool)`

 SetCopySenderFlagNil sets the value for CopySenderFlag to be an explicit nil

### UnsetCopySenderFlag
`func (o *RmaStatusEmailTemplate) UnsetCopySenderFlag()`

UnsetCopySenderFlag ensures that no value is present for CopySenderFlag, not even an explicit nil
### GetInfo

`func (o *RmaStatusEmailTemplate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *RmaStatusEmailTemplate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *RmaStatusEmailTemplate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *RmaStatusEmailTemplate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


