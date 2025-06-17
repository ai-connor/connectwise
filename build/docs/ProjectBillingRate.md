# ProjectBillingRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProjectRecId** | Pointer to **int32** |  | [optional] 
**HourlyRate** | Pointer to **float64** |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**ActivityClassRecId** | Pointer to **int32** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**MemberRecId** | Pointer to **int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectBillingRate

`func NewProjectBillingRate() *ProjectBillingRate`

NewProjectBillingRate instantiates a new ProjectBillingRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectBillingRateWithDefaults

`func NewProjectBillingRateWithDefaults() *ProjectBillingRate`

NewProjectBillingRateWithDefaults instantiates a new ProjectBillingRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectBillingRate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectBillingRate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectBillingRate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectBillingRate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProjectRecId

`func (o *ProjectBillingRate) GetProjectRecId() int32`

GetProjectRecId returns the ProjectRecId field if non-nil, zero value otherwise.

### GetProjectRecIdOk

`func (o *ProjectBillingRate) GetProjectRecIdOk() (*int32, bool)`

GetProjectRecIdOk returns a tuple with the ProjectRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRecId

`func (o *ProjectBillingRate) SetProjectRecId(v int32)`

SetProjectRecId sets ProjectRecId field to given value.

### HasProjectRecId

`func (o *ProjectBillingRate) HasProjectRecId() bool`

HasProjectRecId returns a boolean if a field has been set.

### GetHourlyRate

`func (o *ProjectBillingRate) GetHourlyRate() float64`

GetHourlyRate returns the HourlyRate field if non-nil, zero value otherwise.

### GetHourlyRateOk

`func (o *ProjectBillingRate) GetHourlyRateOk() (*float64, bool)`

GetHourlyRateOk returns a tuple with the HourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyRate

`func (o *ProjectBillingRate) SetHourlyRate(v float64)`

SetHourlyRate sets HourlyRate field to given value.

### HasHourlyRate

`func (o *ProjectBillingRate) HasHourlyRate() bool`

HasHourlyRate returns a boolean if a field has been set.

### GetWorkRole

`func (o *ProjectBillingRate) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *ProjectBillingRate) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *ProjectBillingRate) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *ProjectBillingRate) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetActivityClassRecId

`func (o *ProjectBillingRate) GetActivityClassRecId() int32`

GetActivityClassRecId returns the ActivityClassRecId field if non-nil, zero value otherwise.

### GetActivityClassRecIdOk

`func (o *ProjectBillingRate) GetActivityClassRecIdOk() (*int32, bool)`

GetActivityClassRecIdOk returns a tuple with the ActivityClassRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityClassRecId

`func (o *ProjectBillingRate) SetActivityClassRecId(v int32)`

SetActivityClassRecId sets ActivityClassRecId field to given value.

### HasActivityClassRecId

`func (o *ProjectBillingRate) HasActivityClassRecId() bool`

HasActivityClassRecId returns a boolean if a field has been set.

### GetMember

`func (o *ProjectBillingRate) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ProjectBillingRate) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ProjectBillingRate) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ProjectBillingRate) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMemberRecId

`func (o *ProjectBillingRate) GetMemberRecId() int32`

GetMemberRecId returns the MemberRecId field if non-nil, zero value otherwise.

### GetMemberRecIdOk

`func (o *ProjectBillingRate) GetMemberRecIdOk() (*int32, bool)`

GetMemberRecIdOk returns a tuple with the MemberRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberRecId

`func (o *ProjectBillingRate) SetMemberRecId(v int32)`

SetMemberRecId sets MemberRecId field to given value.

### HasMemberRecId

`func (o *ProjectBillingRate) HasMemberRecId() bool`

HasMemberRecId returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectBillingRate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectBillingRate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectBillingRate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectBillingRate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


