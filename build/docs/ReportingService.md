# ReportingService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ReportingUserName** | Pointer to **string** |  Max length: 50; | [optional] 
**ReportingPassword** | Pointer to **string** | To blank out the password, enter an empty string here. Max length: 50; | [optional] 
**ReportingDomain** | Pointer to **string** |  Max length: 50; | [optional] 
**ReportingUrl** | Pointer to **string** |  Max length: 100; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewReportingService

`func NewReportingService() *ReportingService`

NewReportingService instantiates a new ReportingService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportingServiceWithDefaults

`func NewReportingServiceWithDefaults() *ReportingService`

NewReportingServiceWithDefaults instantiates a new ReportingService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportingService) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportingService) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportingService) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReportingService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReportingUserName

`func (o *ReportingService) GetReportingUserName() string`

GetReportingUserName returns the ReportingUserName field if non-nil, zero value otherwise.

### GetReportingUserNameOk

`func (o *ReportingService) GetReportingUserNameOk() (*string, bool)`

GetReportingUserNameOk returns a tuple with the ReportingUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingUserName

`func (o *ReportingService) SetReportingUserName(v string)`

SetReportingUserName sets ReportingUserName field to given value.

### HasReportingUserName

`func (o *ReportingService) HasReportingUserName() bool`

HasReportingUserName returns a boolean if a field has been set.

### GetReportingPassword

`func (o *ReportingService) GetReportingPassword() string`

GetReportingPassword returns the ReportingPassword field if non-nil, zero value otherwise.

### GetReportingPasswordOk

`func (o *ReportingService) GetReportingPasswordOk() (*string, bool)`

GetReportingPasswordOk returns a tuple with the ReportingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingPassword

`func (o *ReportingService) SetReportingPassword(v string)`

SetReportingPassword sets ReportingPassword field to given value.

### HasReportingPassword

`func (o *ReportingService) HasReportingPassword() bool`

HasReportingPassword returns a boolean if a field has been set.

### GetReportingDomain

`func (o *ReportingService) GetReportingDomain() string`

GetReportingDomain returns the ReportingDomain field if non-nil, zero value otherwise.

### GetReportingDomainOk

`func (o *ReportingService) GetReportingDomainOk() (*string, bool)`

GetReportingDomainOk returns a tuple with the ReportingDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingDomain

`func (o *ReportingService) SetReportingDomain(v string)`

SetReportingDomain sets ReportingDomain field to given value.

### HasReportingDomain

`func (o *ReportingService) HasReportingDomain() bool`

HasReportingDomain returns a boolean if a field has been set.

### GetReportingUrl

`func (o *ReportingService) GetReportingUrl() string`

GetReportingUrl returns the ReportingUrl field if non-nil, zero value otherwise.

### GetReportingUrlOk

`func (o *ReportingService) GetReportingUrlOk() (*string, bool)`

GetReportingUrlOk returns a tuple with the ReportingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingUrl

`func (o *ReportingService) SetReportingUrl(v string)`

SetReportingUrl sets ReportingUrl field to given value.

### HasReportingUrl

`func (o *ReportingService) HasReportingUrl() bool`

HasReportingUrl returns a boolean if a field has been set.

### GetInfo

`func (o *ReportingService) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ReportingService) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ReportingService) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ReportingService) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


