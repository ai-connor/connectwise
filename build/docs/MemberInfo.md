# MemberInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**MiddleInitial** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**DefaultEmail** | Pointer to **string** |  | [optional] 
**Photo** | Pointer to [**DocumentReference**](DocumentReference.md) |  | [optional] 
**LicenseClass** | Pointer to **NullableString** | F &#x3D; Full Member, A &#x3D; API Member, C &#x3D; StreamlineIT Member, X &#x3D; Subcontractor Member | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMemberInfo

`func NewMemberInfo() *MemberInfo`

NewMemberInfo instantiates a new MemberInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberInfoWithDefaults

`func NewMemberInfoWithDefaults() *MemberInfo`

NewMemberInfoWithDefaults instantiates a new MemberInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *MemberInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MemberInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MemberInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MemberInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetFirstName

`func (o *MemberInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MemberInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MemberInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MemberInfo) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleInitial

`func (o *MemberInfo) GetMiddleInitial() string`

GetMiddleInitial returns the MiddleInitial field if non-nil, zero value otherwise.

### GetMiddleInitialOk

`func (o *MemberInfo) GetMiddleInitialOk() (*string, bool)`

GetMiddleInitialOk returns a tuple with the MiddleInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleInitial

`func (o *MemberInfo) SetMiddleInitial(v string)`

SetMiddleInitial sets MiddleInitial field to given value.

### HasMiddleInitial

`func (o *MemberInfo) HasMiddleInitial() bool`

HasMiddleInitial returns a boolean if a field has been set.

### GetLastName

`func (o *MemberInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MemberInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MemberInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MemberInfo) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFullName

`func (o *MemberInfo) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *MemberInfo) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *MemberInfo) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *MemberInfo) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDefaultEmail

`func (o *MemberInfo) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *MemberInfo) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *MemberInfo) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.

### HasDefaultEmail

`func (o *MemberInfo) HasDefaultEmail() bool`

HasDefaultEmail returns a boolean if a field has been set.

### GetPhoto

`func (o *MemberInfo) GetPhoto() DocumentReference`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MemberInfo) GetPhotoOk() (*DocumentReference, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *MemberInfo) SetPhoto(v DocumentReference)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *MemberInfo) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetLicenseClass

`func (o *MemberInfo) GetLicenseClass() string`

GetLicenseClass returns the LicenseClass field if non-nil, zero value otherwise.

### GetLicenseClassOk

`func (o *MemberInfo) GetLicenseClassOk() (*string, bool)`

GetLicenseClassOk returns a tuple with the LicenseClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseClass

`func (o *MemberInfo) SetLicenseClass(v string)`

SetLicenseClass sets LicenseClass field to given value.

### HasLicenseClass

`func (o *MemberInfo) HasLicenseClass() bool`

HasLicenseClass returns a boolean if a field has been set.

### SetLicenseClassNil

`func (o *MemberInfo) SetLicenseClassNil(b bool)`

 SetLicenseClassNil sets the value for LicenseClass to be an explicit nil

### UnsetLicenseClass
`func (o *MemberInfo) UnsetLicenseClass()`

UnsetLicenseClass ensures that no value is present for LicenseClass, not even an explicit nil
### GetInactiveFlag

`func (o *MemberInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *MemberInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *MemberInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *MemberInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *MemberInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *MemberInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *MemberInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


