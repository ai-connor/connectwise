# OrderStatusEmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**OrderStatusReference**](OrderStatusReference.md) |  | [optional] 
**UseSenderFlag** | Pointer to **NullableBool** |  | [optional] 
**FirstName** | Pointer to **string** |  Max length: 100; | [optional] 
**LastName** | Pointer to **string** |  Max length: 100; | [optional] 
**EmailAddress** | Pointer to **string** |  Max length: 100; | [optional] 
**Subject** | **string** |  Max length: 200; | 
**Body** | **string** |  | 
**CopySenderFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOrderStatusEmailTemplate

`func NewOrderStatusEmailTemplate(subject string, body string, ) *OrderStatusEmailTemplate`

NewOrderStatusEmailTemplate instantiates a new OrderStatusEmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusEmailTemplateWithDefaults

`func NewOrderStatusEmailTemplateWithDefaults() *OrderStatusEmailTemplate`

NewOrderStatusEmailTemplateWithDefaults instantiates a new OrderStatusEmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderStatusEmailTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderStatusEmailTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderStatusEmailTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrderStatusEmailTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *OrderStatusEmailTemplate) GetStatus() OrderStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderStatusEmailTemplate) GetStatusOk() (*OrderStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderStatusEmailTemplate) SetStatus(v OrderStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderStatusEmailTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUseSenderFlag

`func (o *OrderStatusEmailTemplate) GetUseSenderFlag() bool`

GetUseSenderFlag returns the UseSenderFlag field if non-nil, zero value otherwise.

### GetUseSenderFlagOk

`func (o *OrderStatusEmailTemplate) GetUseSenderFlagOk() (*bool, bool)`

GetUseSenderFlagOk returns a tuple with the UseSenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSenderFlag

`func (o *OrderStatusEmailTemplate) SetUseSenderFlag(v bool)`

SetUseSenderFlag sets UseSenderFlag field to given value.

### HasUseSenderFlag

`func (o *OrderStatusEmailTemplate) HasUseSenderFlag() bool`

HasUseSenderFlag returns a boolean if a field has been set.

### SetUseSenderFlagNil

`func (o *OrderStatusEmailTemplate) SetUseSenderFlagNil(b bool)`

 SetUseSenderFlagNil sets the value for UseSenderFlag to be an explicit nil

### UnsetUseSenderFlag
`func (o *OrderStatusEmailTemplate) UnsetUseSenderFlag()`

UnsetUseSenderFlag ensures that no value is present for UseSenderFlag, not even an explicit nil
### GetFirstName

`func (o *OrderStatusEmailTemplate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrderStatusEmailTemplate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrderStatusEmailTemplate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *OrderStatusEmailTemplate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *OrderStatusEmailTemplate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrderStatusEmailTemplate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrderStatusEmailTemplate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *OrderStatusEmailTemplate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *OrderStatusEmailTemplate) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *OrderStatusEmailTemplate) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *OrderStatusEmailTemplate) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *OrderStatusEmailTemplate) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetSubject

`func (o *OrderStatusEmailTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OrderStatusEmailTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OrderStatusEmailTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetBody

`func (o *OrderStatusEmailTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *OrderStatusEmailTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *OrderStatusEmailTemplate) SetBody(v string)`

SetBody sets Body field to given value.


### GetCopySenderFlag

`func (o *OrderStatusEmailTemplate) GetCopySenderFlag() bool`

GetCopySenderFlag returns the CopySenderFlag field if non-nil, zero value otherwise.

### GetCopySenderFlagOk

`func (o *OrderStatusEmailTemplate) GetCopySenderFlagOk() (*bool, bool)`

GetCopySenderFlagOk returns a tuple with the CopySenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySenderFlag

`func (o *OrderStatusEmailTemplate) SetCopySenderFlag(v bool)`

SetCopySenderFlag sets CopySenderFlag field to given value.

### HasCopySenderFlag

`func (o *OrderStatusEmailTemplate) HasCopySenderFlag() bool`

HasCopySenderFlag returns a boolean if a field has been set.

### SetCopySenderFlagNil

`func (o *OrderStatusEmailTemplate) SetCopySenderFlagNil(b bool)`

 SetCopySenderFlagNil sets the value for CopySenderFlag to be an explicit nil

### UnsetCopySenderFlag
`func (o *OrderStatusEmailTemplate) UnsetCopySenderFlag()`

UnsetCopySenderFlag ensures that no value is present for CopySenderFlag, not even an explicit nil
### GetInfo

`func (o *OrderStatusEmailTemplate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OrderStatusEmailTemplate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OrderStatusEmailTemplate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OrderStatusEmailTemplate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


