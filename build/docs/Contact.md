# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**AddressLine1** | Pointer to **string** |  | [optional] 
**AddressLine2** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 
**Country** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**Relationship** | Pointer to [**RelationshipReference**](RelationshipReference.md) |  | [optional] 
**RelationshipOverride** | Pointer to **string** |  | [optional] 
**Department** | Pointer to [**ContactDepartmentReference**](ContactDepartmentReference.md) |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultMergeContactId** | Pointer to **NullableInt32** |  | [optional] 
**SecurityIdentifier** | Pointer to **string** |  | [optional] 
**ManagerContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**AssistantContact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**School** | Pointer to **string** |  | [optional] 
**NickName** | Pointer to **string** |  | [optional] 
**MarriedFlag** | Pointer to **NullableBool** |  | [optional] 
**ChildrenFlag** | Pointer to **NullableBool** |  | [optional] 
**Children** | Pointer to **string** |  | [optional] 
**SignificantOther** | Pointer to **string** |  | [optional] 
**PortalPassword** | Pointer to **string** |  | [optional] 
**PortalSecurityLevel** | Pointer to **NullableInt32** |  | [optional] 
**DisablePortalLoginFlag** | Pointer to **NullableBool** |  | [optional] 
**UnsubscribeFlag** | Pointer to **NullableBool** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**BirthDay** | Pointer to **string** |  | [optional] 
**Anniversary** | Pointer to **string** |  | [optional] 
**Presence** | Pointer to **NullableString** |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**FacebookUrl** | Pointer to **string** |  | [optional] 
**TwitterUrl** | Pointer to **string** |  | [optional] 
**LinkedInUrl** | Pointer to **string** |  | [optional] 
**DefaultPhoneType** | Pointer to **string** |  | [optional] 
**DefaultPhoneNbr** | Pointer to **string** |  | [optional] 
**DefaultPhoneExtension** | Pointer to **string** |  | [optional] 
**DefaultBillingFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**UserDefinedField1** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField2** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField3** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField4** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField5** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField6** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField7** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField8** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField9** | Pointer to **string** |  Max length: 50; | [optional] 
**UserDefinedField10** | Pointer to **string** |  Max length: 50; | [optional] 
**CompanyLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**CommunicationItems** | Pointer to [**[]ContactCommunicationItem**](ContactCommunicationItem.md) |  | [optional] 
**Types** | Pointer to [**[]ContactTypeReference**](ContactTypeReference.md) |  | [optional] 
**IntegratorTags** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 
**Photo** | Pointer to [**DocumentReference**](DocumentReference.md) |  | [optional] 
**IgnoreDuplicates** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**TypeIds** | Pointer to **[]int32** | Gets or sets integrer array of Contact_Type_Recids to be assigned to contact that can be passed in only during new contact creation (post)             To update existing contacts type, use the /company/contactTypeAssociations or /company/contacts/{ID}/typeAssociations endpoints. | [optional] 

## Methods

### NewContact

`func NewContact() *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Contact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Contact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Contact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Contact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFirstName

`func (o *Contact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Contact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Contact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Contact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Contact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Contact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Contact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Contact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetCompany

`func (o *Contact) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Contact) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Contact) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Contact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetSite

`func (o *Contact) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Contact) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Contact) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *Contact) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAddressLine1

`func (o *Contact) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Contact) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Contact) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *Contact) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *Contact) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Contact) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Contact) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Contact) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *Contact) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Contact) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Contact) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Contact) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *Contact) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Contact) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Contact) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Contact) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *Contact) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *Contact) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *Contact) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *Contact) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetCountry

`func (o *Contact) GetCountry() CountryReference`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Contact) GetCountryOk() (*CountryReference, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Contact) SetCountry(v CountryReference)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Contact) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRelationship

`func (o *Contact) GetRelationship() RelationshipReference`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *Contact) GetRelationshipOk() (*RelationshipReference, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *Contact) SetRelationship(v RelationshipReference)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *Contact) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetRelationshipOverride

`func (o *Contact) GetRelationshipOverride() string`

GetRelationshipOverride returns the RelationshipOverride field if non-nil, zero value otherwise.

### GetRelationshipOverrideOk

`func (o *Contact) GetRelationshipOverrideOk() (*string, bool)`

GetRelationshipOverrideOk returns a tuple with the RelationshipOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationshipOverride

`func (o *Contact) SetRelationshipOverride(v string)`

SetRelationshipOverride sets RelationshipOverride field to given value.

### HasRelationshipOverride

`func (o *Contact) HasRelationshipOverride() bool`

HasRelationshipOverride returns a boolean if a field has been set.

### GetDepartment

`func (o *Contact) GetDepartment() ContactDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Contact) GetDepartmentOk() (*ContactDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Contact) SetDepartment(v ContactDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Contact) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *Contact) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *Contact) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *Contact) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *Contact) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *Contact) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *Contact) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDefaultMergeContactId

`func (o *Contact) GetDefaultMergeContactId() int32`

GetDefaultMergeContactId returns the DefaultMergeContactId field if non-nil, zero value otherwise.

### GetDefaultMergeContactIdOk

`func (o *Contact) GetDefaultMergeContactIdOk() (*int32, bool)`

GetDefaultMergeContactIdOk returns a tuple with the DefaultMergeContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMergeContactId

`func (o *Contact) SetDefaultMergeContactId(v int32)`

SetDefaultMergeContactId sets DefaultMergeContactId field to given value.

### HasDefaultMergeContactId

`func (o *Contact) HasDefaultMergeContactId() bool`

HasDefaultMergeContactId returns a boolean if a field has been set.

### SetDefaultMergeContactIdNil

`func (o *Contact) SetDefaultMergeContactIdNil(b bool)`

 SetDefaultMergeContactIdNil sets the value for DefaultMergeContactId to be an explicit nil

### UnsetDefaultMergeContactId
`func (o *Contact) UnsetDefaultMergeContactId()`

UnsetDefaultMergeContactId ensures that no value is present for DefaultMergeContactId, not even an explicit nil
### GetSecurityIdentifier

`func (o *Contact) GetSecurityIdentifier() string`

GetSecurityIdentifier returns the SecurityIdentifier field if non-nil, zero value otherwise.

### GetSecurityIdentifierOk

`func (o *Contact) GetSecurityIdentifierOk() (*string, bool)`

GetSecurityIdentifierOk returns a tuple with the SecurityIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIdentifier

`func (o *Contact) SetSecurityIdentifier(v string)`

SetSecurityIdentifier sets SecurityIdentifier field to given value.

### HasSecurityIdentifier

`func (o *Contact) HasSecurityIdentifier() bool`

HasSecurityIdentifier returns a boolean if a field has been set.

### GetManagerContact

`func (o *Contact) GetManagerContact() ContactReference`

GetManagerContact returns the ManagerContact field if non-nil, zero value otherwise.

### GetManagerContactOk

`func (o *Contact) GetManagerContactOk() (*ContactReference, bool)`

GetManagerContactOk returns a tuple with the ManagerContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerContact

`func (o *Contact) SetManagerContact(v ContactReference)`

SetManagerContact sets ManagerContact field to given value.

### HasManagerContact

`func (o *Contact) HasManagerContact() bool`

HasManagerContact returns a boolean if a field has been set.

### GetAssistantContact

`func (o *Contact) GetAssistantContact() ContactReference`

GetAssistantContact returns the AssistantContact field if non-nil, zero value otherwise.

### GetAssistantContactOk

`func (o *Contact) GetAssistantContactOk() (*ContactReference, bool)`

GetAssistantContactOk returns a tuple with the AssistantContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssistantContact

`func (o *Contact) SetAssistantContact(v ContactReference)`

SetAssistantContact sets AssistantContact field to given value.

### HasAssistantContact

`func (o *Contact) HasAssistantContact() bool`

HasAssistantContact returns a boolean if a field has been set.

### GetTitle

`func (o *Contact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Contact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Contact) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Contact) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSchool

`func (o *Contact) GetSchool() string`

GetSchool returns the School field if non-nil, zero value otherwise.

### GetSchoolOk

`func (o *Contact) GetSchoolOk() (*string, bool)`

GetSchoolOk returns a tuple with the School field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchool

`func (o *Contact) SetSchool(v string)`

SetSchool sets School field to given value.

### HasSchool

`func (o *Contact) HasSchool() bool`

HasSchool returns a boolean if a field has been set.

### GetNickName

`func (o *Contact) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *Contact) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *Contact) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *Contact) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetMarriedFlag

`func (o *Contact) GetMarriedFlag() bool`

GetMarriedFlag returns the MarriedFlag field if non-nil, zero value otherwise.

### GetMarriedFlagOk

`func (o *Contact) GetMarriedFlagOk() (*bool, bool)`

GetMarriedFlagOk returns a tuple with the MarriedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarriedFlag

`func (o *Contact) SetMarriedFlag(v bool)`

SetMarriedFlag sets MarriedFlag field to given value.

### HasMarriedFlag

`func (o *Contact) HasMarriedFlag() bool`

HasMarriedFlag returns a boolean if a field has been set.

### SetMarriedFlagNil

`func (o *Contact) SetMarriedFlagNil(b bool)`

 SetMarriedFlagNil sets the value for MarriedFlag to be an explicit nil

### UnsetMarriedFlag
`func (o *Contact) UnsetMarriedFlag()`

UnsetMarriedFlag ensures that no value is present for MarriedFlag, not even an explicit nil
### GetChildrenFlag

`func (o *Contact) GetChildrenFlag() bool`

GetChildrenFlag returns the ChildrenFlag field if non-nil, zero value otherwise.

### GetChildrenFlagOk

`func (o *Contact) GetChildrenFlagOk() (*bool, bool)`

GetChildrenFlagOk returns a tuple with the ChildrenFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenFlag

`func (o *Contact) SetChildrenFlag(v bool)`

SetChildrenFlag sets ChildrenFlag field to given value.

### HasChildrenFlag

`func (o *Contact) HasChildrenFlag() bool`

HasChildrenFlag returns a boolean if a field has been set.

### SetChildrenFlagNil

`func (o *Contact) SetChildrenFlagNil(b bool)`

 SetChildrenFlagNil sets the value for ChildrenFlag to be an explicit nil

### UnsetChildrenFlag
`func (o *Contact) UnsetChildrenFlag()`

UnsetChildrenFlag ensures that no value is present for ChildrenFlag, not even an explicit nil
### GetChildren

`func (o *Contact) GetChildren() string`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Contact) GetChildrenOk() (*string, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Contact) SetChildren(v string)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Contact) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetSignificantOther

`func (o *Contact) GetSignificantOther() string`

GetSignificantOther returns the SignificantOther field if non-nil, zero value otherwise.

### GetSignificantOtherOk

`func (o *Contact) GetSignificantOtherOk() (*string, bool)`

GetSignificantOtherOk returns a tuple with the SignificantOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificantOther

`func (o *Contact) SetSignificantOther(v string)`

SetSignificantOther sets SignificantOther field to given value.

### HasSignificantOther

`func (o *Contact) HasSignificantOther() bool`

HasSignificantOther returns a boolean if a field has been set.

### GetPortalPassword

`func (o *Contact) GetPortalPassword() string`

GetPortalPassword returns the PortalPassword field if non-nil, zero value otherwise.

### GetPortalPasswordOk

`func (o *Contact) GetPortalPasswordOk() (*string, bool)`

GetPortalPasswordOk returns a tuple with the PortalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalPassword

`func (o *Contact) SetPortalPassword(v string)`

SetPortalPassword sets PortalPassword field to given value.

### HasPortalPassword

`func (o *Contact) HasPortalPassword() bool`

HasPortalPassword returns a boolean if a field has been set.

### GetPortalSecurityLevel

`func (o *Contact) GetPortalSecurityLevel() int32`

GetPortalSecurityLevel returns the PortalSecurityLevel field if non-nil, zero value otherwise.

### GetPortalSecurityLevelOk

`func (o *Contact) GetPortalSecurityLevelOk() (*int32, bool)`

GetPortalSecurityLevelOk returns a tuple with the PortalSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalSecurityLevel

`func (o *Contact) SetPortalSecurityLevel(v int32)`

SetPortalSecurityLevel sets PortalSecurityLevel field to given value.

### HasPortalSecurityLevel

`func (o *Contact) HasPortalSecurityLevel() bool`

HasPortalSecurityLevel returns a boolean if a field has been set.

### SetPortalSecurityLevelNil

`func (o *Contact) SetPortalSecurityLevelNil(b bool)`

 SetPortalSecurityLevelNil sets the value for PortalSecurityLevel to be an explicit nil

### UnsetPortalSecurityLevel
`func (o *Contact) UnsetPortalSecurityLevel()`

UnsetPortalSecurityLevel ensures that no value is present for PortalSecurityLevel, not even an explicit nil
### GetDisablePortalLoginFlag

`func (o *Contact) GetDisablePortalLoginFlag() bool`

GetDisablePortalLoginFlag returns the DisablePortalLoginFlag field if non-nil, zero value otherwise.

### GetDisablePortalLoginFlagOk

`func (o *Contact) GetDisablePortalLoginFlagOk() (*bool, bool)`

GetDisablePortalLoginFlagOk returns a tuple with the DisablePortalLoginFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePortalLoginFlag

`func (o *Contact) SetDisablePortalLoginFlag(v bool)`

SetDisablePortalLoginFlag sets DisablePortalLoginFlag field to given value.

### HasDisablePortalLoginFlag

`func (o *Contact) HasDisablePortalLoginFlag() bool`

HasDisablePortalLoginFlag returns a boolean if a field has been set.

### SetDisablePortalLoginFlagNil

`func (o *Contact) SetDisablePortalLoginFlagNil(b bool)`

 SetDisablePortalLoginFlagNil sets the value for DisablePortalLoginFlag to be an explicit nil

### UnsetDisablePortalLoginFlag
`func (o *Contact) UnsetDisablePortalLoginFlag()`

UnsetDisablePortalLoginFlag ensures that no value is present for DisablePortalLoginFlag, not even an explicit nil
### GetUnsubscribeFlag

`func (o *Contact) GetUnsubscribeFlag() bool`

GetUnsubscribeFlag returns the UnsubscribeFlag field if non-nil, zero value otherwise.

### GetUnsubscribeFlagOk

`func (o *Contact) GetUnsubscribeFlagOk() (*bool, bool)`

GetUnsubscribeFlagOk returns a tuple with the UnsubscribeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeFlag

`func (o *Contact) SetUnsubscribeFlag(v bool)`

SetUnsubscribeFlag sets UnsubscribeFlag field to given value.

### HasUnsubscribeFlag

`func (o *Contact) HasUnsubscribeFlag() bool`

HasUnsubscribeFlag returns a boolean if a field has been set.

### SetUnsubscribeFlagNil

`func (o *Contact) SetUnsubscribeFlagNil(b bool)`

 SetUnsubscribeFlagNil sets the value for UnsubscribeFlag to be an explicit nil

### UnsetUnsubscribeFlag
`func (o *Contact) UnsetUnsubscribeFlag()`

UnsetUnsubscribeFlag ensures that no value is present for UnsubscribeFlag, not even an explicit nil
### GetGender

`func (o *Contact) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *Contact) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *Contact) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *Contact) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *Contact) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *Contact) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetBirthDay

`func (o *Contact) GetBirthDay() string`

GetBirthDay returns the BirthDay field if non-nil, zero value otherwise.

### GetBirthDayOk

`func (o *Contact) GetBirthDayOk() (*string, bool)`

GetBirthDayOk returns a tuple with the BirthDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDay

`func (o *Contact) SetBirthDay(v string)`

SetBirthDay sets BirthDay field to given value.

### HasBirthDay

`func (o *Contact) HasBirthDay() bool`

HasBirthDay returns a boolean if a field has been set.

### GetAnniversary

`func (o *Contact) GetAnniversary() string`

GetAnniversary returns the Anniversary field if non-nil, zero value otherwise.

### GetAnniversaryOk

`func (o *Contact) GetAnniversaryOk() (*string, bool)`

GetAnniversaryOk returns a tuple with the Anniversary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnniversary

`func (o *Contact) SetAnniversary(v string)`

SetAnniversary sets Anniversary field to given value.

### HasAnniversary

`func (o *Contact) HasAnniversary() bool`

HasAnniversary returns a boolean if a field has been set.

### GetPresence

`func (o *Contact) GetPresence() string`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *Contact) GetPresenceOk() (*string, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *Contact) SetPresence(v string)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *Contact) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### SetPresenceNil

`func (o *Contact) SetPresenceNil(b bool)`

 SetPresenceNil sets the value for Presence to be an explicit nil

### UnsetPresence
`func (o *Contact) UnsetPresence()`

UnsetPresence ensures that no value is present for Presence, not even an explicit nil
### GetMobileGuid

`func (o *Contact) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *Contact) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *Contact) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *Contact) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *Contact) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *Contact) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetFacebookUrl

`func (o *Contact) GetFacebookUrl() string`

GetFacebookUrl returns the FacebookUrl field if non-nil, zero value otherwise.

### GetFacebookUrlOk

`func (o *Contact) GetFacebookUrlOk() (*string, bool)`

GetFacebookUrlOk returns a tuple with the FacebookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookUrl

`func (o *Contact) SetFacebookUrl(v string)`

SetFacebookUrl sets FacebookUrl field to given value.

### HasFacebookUrl

`func (o *Contact) HasFacebookUrl() bool`

HasFacebookUrl returns a boolean if a field has been set.

### GetTwitterUrl

`func (o *Contact) GetTwitterUrl() string`

GetTwitterUrl returns the TwitterUrl field if non-nil, zero value otherwise.

### GetTwitterUrlOk

`func (o *Contact) GetTwitterUrlOk() (*string, bool)`

GetTwitterUrlOk returns a tuple with the TwitterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUrl

`func (o *Contact) SetTwitterUrl(v string)`

SetTwitterUrl sets TwitterUrl field to given value.

### HasTwitterUrl

`func (o *Contact) HasTwitterUrl() bool`

HasTwitterUrl returns a boolean if a field has been set.

### GetLinkedInUrl

`func (o *Contact) GetLinkedInUrl() string`

GetLinkedInUrl returns the LinkedInUrl field if non-nil, zero value otherwise.

### GetLinkedInUrlOk

`func (o *Contact) GetLinkedInUrlOk() (*string, bool)`

GetLinkedInUrlOk returns a tuple with the LinkedInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedInUrl

`func (o *Contact) SetLinkedInUrl(v string)`

SetLinkedInUrl sets LinkedInUrl field to given value.

### HasLinkedInUrl

`func (o *Contact) HasLinkedInUrl() bool`

HasLinkedInUrl returns a boolean if a field has been set.

### GetDefaultPhoneType

`func (o *Contact) GetDefaultPhoneType() string`

GetDefaultPhoneType returns the DefaultPhoneType field if non-nil, zero value otherwise.

### GetDefaultPhoneTypeOk

`func (o *Contact) GetDefaultPhoneTypeOk() (*string, bool)`

GetDefaultPhoneTypeOk returns a tuple with the DefaultPhoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhoneType

`func (o *Contact) SetDefaultPhoneType(v string)`

SetDefaultPhoneType sets DefaultPhoneType field to given value.

### HasDefaultPhoneType

`func (o *Contact) HasDefaultPhoneType() bool`

HasDefaultPhoneType returns a boolean if a field has been set.

### GetDefaultPhoneNbr

`func (o *Contact) GetDefaultPhoneNbr() string`

GetDefaultPhoneNbr returns the DefaultPhoneNbr field if non-nil, zero value otherwise.

### GetDefaultPhoneNbrOk

`func (o *Contact) GetDefaultPhoneNbrOk() (*string, bool)`

GetDefaultPhoneNbrOk returns a tuple with the DefaultPhoneNbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhoneNbr

`func (o *Contact) SetDefaultPhoneNbr(v string)`

SetDefaultPhoneNbr sets DefaultPhoneNbr field to given value.

### HasDefaultPhoneNbr

`func (o *Contact) HasDefaultPhoneNbr() bool`

HasDefaultPhoneNbr returns a boolean if a field has been set.

### GetDefaultPhoneExtension

`func (o *Contact) GetDefaultPhoneExtension() string`

GetDefaultPhoneExtension returns the DefaultPhoneExtension field if non-nil, zero value otherwise.

### GetDefaultPhoneExtensionOk

`func (o *Contact) GetDefaultPhoneExtensionOk() (*string, bool)`

GetDefaultPhoneExtensionOk returns a tuple with the DefaultPhoneExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPhoneExtension

`func (o *Contact) SetDefaultPhoneExtension(v string)`

SetDefaultPhoneExtension sets DefaultPhoneExtension field to given value.

### HasDefaultPhoneExtension

`func (o *Contact) HasDefaultPhoneExtension() bool`

HasDefaultPhoneExtension returns a boolean if a field has been set.

### GetDefaultBillingFlag

`func (o *Contact) GetDefaultBillingFlag() bool`

GetDefaultBillingFlag returns the DefaultBillingFlag field if non-nil, zero value otherwise.

### GetDefaultBillingFlagOk

`func (o *Contact) GetDefaultBillingFlagOk() (*bool, bool)`

GetDefaultBillingFlagOk returns a tuple with the DefaultBillingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBillingFlag

`func (o *Contact) SetDefaultBillingFlag(v bool)`

SetDefaultBillingFlag sets DefaultBillingFlag field to given value.

### HasDefaultBillingFlag

`func (o *Contact) HasDefaultBillingFlag() bool`

HasDefaultBillingFlag returns a boolean if a field has been set.

### SetDefaultBillingFlagNil

`func (o *Contact) SetDefaultBillingFlagNil(b bool)`

 SetDefaultBillingFlagNil sets the value for DefaultBillingFlag to be an explicit nil

### UnsetDefaultBillingFlag
`func (o *Contact) UnsetDefaultBillingFlag()`

UnsetDefaultBillingFlag ensures that no value is present for DefaultBillingFlag, not even an explicit nil
### GetDefaultFlag

`func (o *Contact) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *Contact) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *Contact) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *Contact) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *Contact) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *Contact) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetUserDefinedField1

`func (o *Contact) GetUserDefinedField1() string`

GetUserDefinedField1 returns the UserDefinedField1 field if non-nil, zero value otherwise.

### GetUserDefinedField1Ok

`func (o *Contact) GetUserDefinedField1Ok() (*string, bool)`

GetUserDefinedField1Ok returns a tuple with the UserDefinedField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField1

`func (o *Contact) SetUserDefinedField1(v string)`

SetUserDefinedField1 sets UserDefinedField1 field to given value.

### HasUserDefinedField1

`func (o *Contact) HasUserDefinedField1() bool`

HasUserDefinedField1 returns a boolean if a field has been set.

### GetUserDefinedField2

`func (o *Contact) GetUserDefinedField2() string`

GetUserDefinedField2 returns the UserDefinedField2 field if non-nil, zero value otherwise.

### GetUserDefinedField2Ok

`func (o *Contact) GetUserDefinedField2Ok() (*string, bool)`

GetUserDefinedField2Ok returns a tuple with the UserDefinedField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField2

`func (o *Contact) SetUserDefinedField2(v string)`

SetUserDefinedField2 sets UserDefinedField2 field to given value.

### HasUserDefinedField2

`func (o *Contact) HasUserDefinedField2() bool`

HasUserDefinedField2 returns a boolean if a field has been set.

### GetUserDefinedField3

`func (o *Contact) GetUserDefinedField3() string`

GetUserDefinedField3 returns the UserDefinedField3 field if non-nil, zero value otherwise.

### GetUserDefinedField3Ok

`func (o *Contact) GetUserDefinedField3Ok() (*string, bool)`

GetUserDefinedField3Ok returns a tuple with the UserDefinedField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField3

`func (o *Contact) SetUserDefinedField3(v string)`

SetUserDefinedField3 sets UserDefinedField3 field to given value.

### HasUserDefinedField3

`func (o *Contact) HasUserDefinedField3() bool`

HasUserDefinedField3 returns a boolean if a field has been set.

### GetUserDefinedField4

`func (o *Contact) GetUserDefinedField4() string`

GetUserDefinedField4 returns the UserDefinedField4 field if non-nil, zero value otherwise.

### GetUserDefinedField4Ok

`func (o *Contact) GetUserDefinedField4Ok() (*string, bool)`

GetUserDefinedField4Ok returns a tuple with the UserDefinedField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField4

`func (o *Contact) SetUserDefinedField4(v string)`

SetUserDefinedField4 sets UserDefinedField4 field to given value.

### HasUserDefinedField4

`func (o *Contact) HasUserDefinedField4() bool`

HasUserDefinedField4 returns a boolean if a field has been set.

### GetUserDefinedField5

`func (o *Contact) GetUserDefinedField5() string`

GetUserDefinedField5 returns the UserDefinedField5 field if non-nil, zero value otherwise.

### GetUserDefinedField5Ok

`func (o *Contact) GetUserDefinedField5Ok() (*string, bool)`

GetUserDefinedField5Ok returns a tuple with the UserDefinedField5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField5

`func (o *Contact) SetUserDefinedField5(v string)`

SetUserDefinedField5 sets UserDefinedField5 field to given value.

### HasUserDefinedField5

`func (o *Contact) HasUserDefinedField5() bool`

HasUserDefinedField5 returns a boolean if a field has been set.

### GetUserDefinedField6

`func (o *Contact) GetUserDefinedField6() string`

GetUserDefinedField6 returns the UserDefinedField6 field if non-nil, zero value otherwise.

### GetUserDefinedField6Ok

`func (o *Contact) GetUserDefinedField6Ok() (*string, bool)`

GetUserDefinedField6Ok returns a tuple with the UserDefinedField6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField6

`func (o *Contact) SetUserDefinedField6(v string)`

SetUserDefinedField6 sets UserDefinedField6 field to given value.

### HasUserDefinedField6

`func (o *Contact) HasUserDefinedField6() bool`

HasUserDefinedField6 returns a boolean if a field has been set.

### GetUserDefinedField7

`func (o *Contact) GetUserDefinedField7() string`

GetUserDefinedField7 returns the UserDefinedField7 field if non-nil, zero value otherwise.

### GetUserDefinedField7Ok

`func (o *Contact) GetUserDefinedField7Ok() (*string, bool)`

GetUserDefinedField7Ok returns a tuple with the UserDefinedField7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField7

`func (o *Contact) SetUserDefinedField7(v string)`

SetUserDefinedField7 sets UserDefinedField7 field to given value.

### HasUserDefinedField7

`func (o *Contact) HasUserDefinedField7() bool`

HasUserDefinedField7 returns a boolean if a field has been set.

### GetUserDefinedField8

`func (o *Contact) GetUserDefinedField8() string`

GetUserDefinedField8 returns the UserDefinedField8 field if non-nil, zero value otherwise.

### GetUserDefinedField8Ok

`func (o *Contact) GetUserDefinedField8Ok() (*string, bool)`

GetUserDefinedField8Ok returns a tuple with the UserDefinedField8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField8

`func (o *Contact) SetUserDefinedField8(v string)`

SetUserDefinedField8 sets UserDefinedField8 field to given value.

### HasUserDefinedField8

`func (o *Contact) HasUserDefinedField8() bool`

HasUserDefinedField8 returns a boolean if a field has been set.

### GetUserDefinedField9

`func (o *Contact) GetUserDefinedField9() string`

GetUserDefinedField9 returns the UserDefinedField9 field if non-nil, zero value otherwise.

### GetUserDefinedField9Ok

`func (o *Contact) GetUserDefinedField9Ok() (*string, bool)`

GetUserDefinedField9Ok returns a tuple with the UserDefinedField9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField9

`func (o *Contact) SetUserDefinedField9(v string)`

SetUserDefinedField9 sets UserDefinedField9 field to given value.

### HasUserDefinedField9

`func (o *Contact) HasUserDefinedField9() bool`

HasUserDefinedField9 returns a boolean if a field has been set.

### GetUserDefinedField10

`func (o *Contact) GetUserDefinedField10() string`

GetUserDefinedField10 returns the UserDefinedField10 field if non-nil, zero value otherwise.

### GetUserDefinedField10Ok

`func (o *Contact) GetUserDefinedField10Ok() (*string, bool)`

GetUserDefinedField10Ok returns a tuple with the UserDefinedField10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField10

`func (o *Contact) SetUserDefinedField10(v string)`

SetUserDefinedField10 sets UserDefinedField10 field to given value.

### HasUserDefinedField10

`func (o *Contact) HasUserDefinedField10() bool`

HasUserDefinedField10 returns a boolean if a field has been set.

### GetCompanyLocation

`func (o *Contact) GetCompanyLocation() SystemLocationReference`

GetCompanyLocation returns the CompanyLocation field if non-nil, zero value otherwise.

### GetCompanyLocationOk

`func (o *Contact) GetCompanyLocationOk() (*SystemLocationReference, bool)`

GetCompanyLocationOk returns a tuple with the CompanyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLocation

`func (o *Contact) SetCompanyLocation(v SystemLocationReference)`

SetCompanyLocation sets CompanyLocation field to given value.

### HasCompanyLocation

`func (o *Contact) HasCompanyLocation() bool`

HasCompanyLocation returns a boolean if a field has been set.

### GetCommunicationItems

`func (o *Contact) GetCommunicationItems() []ContactCommunicationItem`

GetCommunicationItems returns the CommunicationItems field if non-nil, zero value otherwise.

### GetCommunicationItemsOk

`func (o *Contact) GetCommunicationItemsOk() (*[]ContactCommunicationItem, bool)`

GetCommunicationItemsOk returns a tuple with the CommunicationItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationItems

`func (o *Contact) SetCommunicationItems(v []ContactCommunicationItem)`

SetCommunicationItems sets CommunicationItems field to given value.

### HasCommunicationItems

`func (o *Contact) HasCommunicationItems() bool`

HasCommunicationItems returns a boolean if a field has been set.

### GetTypes

`func (o *Contact) GetTypes() []ContactTypeReference`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *Contact) GetTypesOk() (*[]ContactTypeReference, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *Contact) SetTypes(v []ContactTypeReference)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *Contact) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetIntegratorTags

`func (o *Contact) GetIntegratorTags() []string`

GetIntegratorTags returns the IntegratorTags field if non-nil, zero value otherwise.

### GetIntegratorTagsOk

`func (o *Contact) GetIntegratorTagsOk() (*[]string, bool)`

GetIntegratorTagsOk returns a tuple with the IntegratorTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorTags

`func (o *Contact) SetIntegratorTags(v []string)`

SetIntegratorTags sets IntegratorTags field to given value.

### HasIntegratorTags

`func (o *Contact) HasIntegratorTags() bool`

HasIntegratorTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *Contact) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Contact) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Contact) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Contact) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetPhoto

`func (o *Contact) GetPhoto() DocumentReference`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *Contact) GetPhotoOk() (*DocumentReference, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *Contact) SetPhoto(v DocumentReference)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *Contact) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetIgnoreDuplicates

`func (o *Contact) GetIgnoreDuplicates() bool`

GetIgnoreDuplicates returns the IgnoreDuplicates field if non-nil, zero value otherwise.

### GetIgnoreDuplicatesOk

`func (o *Contact) GetIgnoreDuplicatesOk() (*bool, bool)`

GetIgnoreDuplicatesOk returns a tuple with the IgnoreDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDuplicates

`func (o *Contact) SetIgnoreDuplicates(v bool)`

SetIgnoreDuplicates sets IgnoreDuplicates field to given value.

### HasIgnoreDuplicates

`func (o *Contact) HasIgnoreDuplicates() bool`

HasIgnoreDuplicates returns a boolean if a field has been set.

### GetInfo

`func (o *Contact) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Contact) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Contact) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Contact) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetTypeIds

`func (o *Contact) GetTypeIds() []int32`

GetTypeIds returns the TypeIds field if non-nil, zero value otherwise.

### GetTypeIdsOk

`func (o *Contact) GetTypeIdsOk() (*[]int32, bool)`

GetTypeIdsOk returns a tuple with the TypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeIds

`func (o *Contact) SetTypeIds(v []int32)`

SetTypeIds sets TypeIds field to given value.

### HasTypeIds

`func (o *Contact) HasTypeIds() bool`

HasTypeIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


