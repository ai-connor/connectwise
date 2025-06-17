# ServiceEmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | **NullableString** |  | 
**ServiceSurvey** | Pointer to [**ServiceSurveyReference**](ServiceSurveyReference.md) |  | [optional] 
**ServiceBoard** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**UseSenderFlag** | Pointer to **NullableBool** |  | [optional] 
**FirstName** | Pointer to **string** | From fields (first name, last name, email address) are required if useSenderFlag is false. Max length: 100; | [optional] 
**LastName** | Pointer to **string** | From fields (first name, last name, email address) are required if useSenderFlag is false. Max length: 100; | [optional] 
**EmailAddress** | Pointer to **string** | From fields (first name, last name, email address) are required if useSenderFlag is false. Max length: 100; | [optional] 
**Subject** | Pointer to **string** |  Max length: 200; | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**CopySenderFlag** | Pointer to **NullableBool** |  | [optional] 
**TasksFlag** | Pointer to **NullableBool** |  | [optional] 
**ResourceRecordsFlag** | Pointer to **NullableBool** |  | [optional] 
**ExternalContactNotifications** | Pointer to **NullableBool** |  | [optional] 
**InternalContactNotifications** | Pointer to **NullableBool** |  | [optional] 
**ServiceStatus** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceEmailTemplate

`func NewServiceEmailTemplate(type_ NullableString, ) *ServiceEmailTemplate`

NewServiceEmailTemplate instantiates a new ServiceEmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceEmailTemplateWithDefaults

`func NewServiceEmailTemplateWithDefaults() *ServiceEmailTemplate`

NewServiceEmailTemplateWithDefaults instantiates a new ServiceEmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceEmailTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceEmailTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceEmailTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceEmailTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ServiceEmailTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceEmailTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceEmailTemplate) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ServiceEmailTemplate) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ServiceEmailTemplate) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetServiceSurvey

`func (o *ServiceEmailTemplate) GetServiceSurvey() ServiceSurveyReference`

GetServiceSurvey returns the ServiceSurvey field if non-nil, zero value otherwise.

### GetServiceSurveyOk

`func (o *ServiceEmailTemplate) GetServiceSurveyOk() (*ServiceSurveyReference, bool)`

GetServiceSurveyOk returns a tuple with the ServiceSurvey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSurvey

`func (o *ServiceEmailTemplate) SetServiceSurvey(v ServiceSurveyReference)`

SetServiceSurvey sets ServiceSurvey field to given value.

### HasServiceSurvey

`func (o *ServiceEmailTemplate) HasServiceSurvey() bool`

HasServiceSurvey returns a boolean if a field has been set.

### GetServiceBoard

`func (o *ServiceEmailTemplate) GetServiceBoard() BoardReference`

GetServiceBoard returns the ServiceBoard field if non-nil, zero value otherwise.

### GetServiceBoardOk

`func (o *ServiceEmailTemplate) GetServiceBoardOk() (*BoardReference, bool)`

GetServiceBoardOk returns a tuple with the ServiceBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoard

`func (o *ServiceEmailTemplate) SetServiceBoard(v BoardReference)`

SetServiceBoard sets ServiceBoard field to given value.

### HasServiceBoard

`func (o *ServiceEmailTemplate) HasServiceBoard() bool`

HasServiceBoard returns a boolean if a field has been set.

### GetUseSenderFlag

`func (o *ServiceEmailTemplate) GetUseSenderFlag() bool`

GetUseSenderFlag returns the UseSenderFlag field if non-nil, zero value otherwise.

### GetUseSenderFlagOk

`func (o *ServiceEmailTemplate) GetUseSenderFlagOk() (*bool, bool)`

GetUseSenderFlagOk returns a tuple with the UseSenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSenderFlag

`func (o *ServiceEmailTemplate) SetUseSenderFlag(v bool)`

SetUseSenderFlag sets UseSenderFlag field to given value.

### HasUseSenderFlag

`func (o *ServiceEmailTemplate) HasUseSenderFlag() bool`

HasUseSenderFlag returns a boolean if a field has been set.

### SetUseSenderFlagNil

`func (o *ServiceEmailTemplate) SetUseSenderFlagNil(b bool)`

 SetUseSenderFlagNil sets the value for UseSenderFlag to be an explicit nil

### UnsetUseSenderFlag
`func (o *ServiceEmailTemplate) UnsetUseSenderFlag()`

UnsetUseSenderFlag ensures that no value is present for UseSenderFlag, not even an explicit nil
### GetFirstName

`func (o *ServiceEmailTemplate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ServiceEmailTemplate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ServiceEmailTemplate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ServiceEmailTemplate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ServiceEmailTemplate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ServiceEmailTemplate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ServiceEmailTemplate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ServiceEmailTemplate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *ServiceEmailTemplate) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ServiceEmailTemplate) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ServiceEmailTemplate) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ServiceEmailTemplate) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetSubject

`func (o *ServiceEmailTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ServiceEmailTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ServiceEmailTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ServiceEmailTemplate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetBody

`func (o *ServiceEmailTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ServiceEmailTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ServiceEmailTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ServiceEmailTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCopySenderFlag

`func (o *ServiceEmailTemplate) GetCopySenderFlag() bool`

GetCopySenderFlag returns the CopySenderFlag field if non-nil, zero value otherwise.

### GetCopySenderFlagOk

`func (o *ServiceEmailTemplate) GetCopySenderFlagOk() (*bool, bool)`

GetCopySenderFlagOk returns a tuple with the CopySenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySenderFlag

`func (o *ServiceEmailTemplate) SetCopySenderFlag(v bool)`

SetCopySenderFlag sets CopySenderFlag field to given value.

### HasCopySenderFlag

`func (o *ServiceEmailTemplate) HasCopySenderFlag() bool`

HasCopySenderFlag returns a boolean if a field has been set.

### SetCopySenderFlagNil

`func (o *ServiceEmailTemplate) SetCopySenderFlagNil(b bool)`

 SetCopySenderFlagNil sets the value for CopySenderFlag to be an explicit nil

### UnsetCopySenderFlag
`func (o *ServiceEmailTemplate) UnsetCopySenderFlag()`

UnsetCopySenderFlag ensures that no value is present for CopySenderFlag, not even an explicit nil
### GetTasksFlag

`func (o *ServiceEmailTemplate) GetTasksFlag() bool`

GetTasksFlag returns the TasksFlag field if non-nil, zero value otherwise.

### GetTasksFlagOk

`func (o *ServiceEmailTemplate) GetTasksFlagOk() (*bool, bool)`

GetTasksFlagOk returns a tuple with the TasksFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasksFlag

`func (o *ServiceEmailTemplate) SetTasksFlag(v bool)`

SetTasksFlag sets TasksFlag field to given value.

### HasTasksFlag

`func (o *ServiceEmailTemplate) HasTasksFlag() bool`

HasTasksFlag returns a boolean if a field has been set.

### SetTasksFlagNil

`func (o *ServiceEmailTemplate) SetTasksFlagNil(b bool)`

 SetTasksFlagNil sets the value for TasksFlag to be an explicit nil

### UnsetTasksFlag
`func (o *ServiceEmailTemplate) UnsetTasksFlag()`

UnsetTasksFlag ensures that no value is present for TasksFlag, not even an explicit nil
### GetResourceRecordsFlag

`func (o *ServiceEmailTemplate) GetResourceRecordsFlag() bool`

GetResourceRecordsFlag returns the ResourceRecordsFlag field if non-nil, zero value otherwise.

### GetResourceRecordsFlagOk

`func (o *ServiceEmailTemplate) GetResourceRecordsFlagOk() (*bool, bool)`

GetResourceRecordsFlagOk returns a tuple with the ResourceRecordsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRecordsFlag

`func (o *ServiceEmailTemplate) SetResourceRecordsFlag(v bool)`

SetResourceRecordsFlag sets ResourceRecordsFlag field to given value.

### HasResourceRecordsFlag

`func (o *ServiceEmailTemplate) HasResourceRecordsFlag() bool`

HasResourceRecordsFlag returns a boolean if a field has been set.

### SetResourceRecordsFlagNil

`func (o *ServiceEmailTemplate) SetResourceRecordsFlagNil(b bool)`

 SetResourceRecordsFlagNil sets the value for ResourceRecordsFlag to be an explicit nil

### UnsetResourceRecordsFlag
`func (o *ServiceEmailTemplate) UnsetResourceRecordsFlag()`

UnsetResourceRecordsFlag ensures that no value is present for ResourceRecordsFlag, not even an explicit nil
### GetExternalContactNotifications

`func (o *ServiceEmailTemplate) GetExternalContactNotifications() bool`

GetExternalContactNotifications returns the ExternalContactNotifications field if non-nil, zero value otherwise.

### GetExternalContactNotificationsOk

`func (o *ServiceEmailTemplate) GetExternalContactNotificationsOk() (*bool, bool)`

GetExternalContactNotificationsOk returns a tuple with the ExternalContactNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalContactNotifications

`func (o *ServiceEmailTemplate) SetExternalContactNotifications(v bool)`

SetExternalContactNotifications sets ExternalContactNotifications field to given value.

### HasExternalContactNotifications

`func (o *ServiceEmailTemplate) HasExternalContactNotifications() bool`

HasExternalContactNotifications returns a boolean if a field has been set.

### SetExternalContactNotificationsNil

`func (o *ServiceEmailTemplate) SetExternalContactNotificationsNil(b bool)`

 SetExternalContactNotificationsNil sets the value for ExternalContactNotifications to be an explicit nil

### UnsetExternalContactNotifications
`func (o *ServiceEmailTemplate) UnsetExternalContactNotifications()`

UnsetExternalContactNotifications ensures that no value is present for ExternalContactNotifications, not even an explicit nil
### GetInternalContactNotifications

`func (o *ServiceEmailTemplate) GetInternalContactNotifications() bool`

GetInternalContactNotifications returns the InternalContactNotifications field if non-nil, zero value otherwise.

### GetInternalContactNotificationsOk

`func (o *ServiceEmailTemplate) GetInternalContactNotificationsOk() (*bool, bool)`

GetInternalContactNotificationsOk returns a tuple with the InternalContactNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalContactNotifications

`func (o *ServiceEmailTemplate) SetInternalContactNotifications(v bool)`

SetInternalContactNotifications sets InternalContactNotifications field to given value.

### HasInternalContactNotifications

`func (o *ServiceEmailTemplate) HasInternalContactNotifications() bool`

HasInternalContactNotifications returns a boolean if a field has been set.

### SetInternalContactNotificationsNil

`func (o *ServiceEmailTemplate) SetInternalContactNotificationsNil(b bool)`

 SetInternalContactNotificationsNil sets the value for InternalContactNotifications to be an explicit nil

### UnsetInternalContactNotifications
`func (o *ServiceEmailTemplate) UnsetInternalContactNotifications()`

UnsetInternalContactNotifications ensures that no value is present for InternalContactNotifications, not even an explicit nil
### GetServiceStatus

`func (o *ServiceEmailTemplate) GetServiceStatus() ServiceStatusReference`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *ServiceEmailTemplate) GetServiceStatusOk() (*ServiceStatusReference, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *ServiceEmailTemplate) SetServiceStatus(v ServiceStatusReference)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *ServiceEmailTemplate) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.

### GetInfo

`func (o *ServiceEmailTemplate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceEmailTemplate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceEmailTemplate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceEmailTemplate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


