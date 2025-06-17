# SsoConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier of the SSO Configuration | [optional] 
**Name** | **string** | Descriptor of the SSO Configuration Max length: 100; | 
**SsoType** | **NullableString** | Type of SSO Configuration | 
**InactiveFlag** | Pointer to **NullableBool** | Whether the SSO configuration is not active | [optional] 
**SamlEntityId** | Pointer to **string** | SAML Identity Provider Id Max length: 1000; | [optional] 
**SamlSignInUrl** | Pointer to **string** | Sign in url for the SAML Identity Provider Max length: 1000; | [optional] 
**SamlIdpCertificate** | Pointer to **string** | Public certificate for Identity Provider signatures | [optional] 
**SamlCertificateName** | Pointer to **string** | Name of the SAML certificate. Metadata on SAML_Idp_Certificate | [optional] 
**SamlCertificateIssuedTo** | Pointer to **string** | Who the SAML certificate was issued to. Metadata on SAML_Idp_Certificate | [optional] 
**SamlCertificateThumbprint** | Pointer to **string** | Thumbprint of the SAML certificate. Metadata on SAML_Idp_Certificate | [optional] 
**SamlCertificateValidFrom** | Pointer to **time.Time** | Date when the SAML certificate becomes valid. Metadata on SAML_Idp_Certificate | [optional] 
**SamlCertificateValidTo** | Pointer to **time.Time** | Date when the SAML certificate is no longer valid. Metadata on SAML_Idp_Certificate | [optional] 
**LocationIds** | **[]int32** | The locations where the SAML Idp Configuration is used | 
**ClientId** | Pointer to **string** | Client identity for this configuration of ConnectWise SSO Max length: 1000; | [optional] 
**StsBaseUrl** | Pointer to **string** | Sign in URL for ConnectWise SSO | [optional] 
**StsUserAdminUrl** | Pointer to **string** | User Admin Url for ConnectWise SSO | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**SubmittedMemberCount** | Pointer to **int32** |  | [optional] 
**AllMembersSubmitted** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**IsSsoOnByDefault** | Pointer to **bool** |  | [optional] 

## Methods

### NewSsoConfiguration

`func NewSsoConfiguration(name string, ssoType NullableString, locationIds []int32, ) *SsoConfiguration`

NewSsoConfiguration instantiates a new SsoConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigurationWithDefaults

`func NewSsoConfigurationWithDefaults() *SsoConfiguration`

NewSsoConfigurationWithDefaults instantiates a new SsoConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SsoConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SsoConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SsoConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SsoConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SsoConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SsoConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SsoConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetSsoType

`func (o *SsoConfiguration) GetSsoType() string`

GetSsoType returns the SsoType field if non-nil, zero value otherwise.

### GetSsoTypeOk

`func (o *SsoConfiguration) GetSsoTypeOk() (*string, bool)`

GetSsoTypeOk returns a tuple with the SsoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoType

`func (o *SsoConfiguration) SetSsoType(v string)`

SetSsoType sets SsoType field to given value.


### SetSsoTypeNil

`func (o *SsoConfiguration) SetSsoTypeNil(b bool)`

 SetSsoTypeNil sets the value for SsoType to be an explicit nil

### UnsetSsoType
`func (o *SsoConfiguration) UnsetSsoType()`

UnsetSsoType ensures that no value is present for SsoType, not even an explicit nil
### GetInactiveFlag

`func (o *SsoConfiguration) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *SsoConfiguration) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *SsoConfiguration) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *SsoConfiguration) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *SsoConfiguration) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *SsoConfiguration) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetSamlEntityId

`func (o *SsoConfiguration) GetSamlEntityId() string`

GetSamlEntityId returns the SamlEntityId field if non-nil, zero value otherwise.

### GetSamlEntityIdOk

`func (o *SsoConfiguration) GetSamlEntityIdOk() (*string, bool)`

GetSamlEntityIdOk returns a tuple with the SamlEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlEntityId

`func (o *SsoConfiguration) SetSamlEntityId(v string)`

SetSamlEntityId sets SamlEntityId field to given value.

### HasSamlEntityId

`func (o *SsoConfiguration) HasSamlEntityId() bool`

HasSamlEntityId returns a boolean if a field has been set.

### GetSamlSignInUrl

`func (o *SsoConfiguration) GetSamlSignInUrl() string`

GetSamlSignInUrl returns the SamlSignInUrl field if non-nil, zero value otherwise.

### GetSamlSignInUrlOk

`func (o *SsoConfiguration) GetSamlSignInUrlOk() (*string, bool)`

GetSamlSignInUrlOk returns a tuple with the SamlSignInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSignInUrl

`func (o *SsoConfiguration) SetSamlSignInUrl(v string)`

SetSamlSignInUrl sets SamlSignInUrl field to given value.

### HasSamlSignInUrl

`func (o *SsoConfiguration) HasSamlSignInUrl() bool`

HasSamlSignInUrl returns a boolean if a field has been set.

### GetSamlIdpCertificate

`func (o *SsoConfiguration) GetSamlIdpCertificate() string`

GetSamlIdpCertificate returns the SamlIdpCertificate field if non-nil, zero value otherwise.

### GetSamlIdpCertificateOk

`func (o *SsoConfiguration) GetSamlIdpCertificateOk() (*string, bool)`

GetSamlIdpCertificateOk returns a tuple with the SamlIdpCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlIdpCertificate

`func (o *SsoConfiguration) SetSamlIdpCertificate(v string)`

SetSamlIdpCertificate sets SamlIdpCertificate field to given value.

### HasSamlIdpCertificate

`func (o *SsoConfiguration) HasSamlIdpCertificate() bool`

HasSamlIdpCertificate returns a boolean if a field has been set.

### GetSamlCertificateName

`func (o *SsoConfiguration) GetSamlCertificateName() string`

GetSamlCertificateName returns the SamlCertificateName field if non-nil, zero value otherwise.

### GetSamlCertificateNameOk

`func (o *SsoConfiguration) GetSamlCertificateNameOk() (*string, bool)`

GetSamlCertificateNameOk returns a tuple with the SamlCertificateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertificateName

`func (o *SsoConfiguration) SetSamlCertificateName(v string)`

SetSamlCertificateName sets SamlCertificateName field to given value.

### HasSamlCertificateName

`func (o *SsoConfiguration) HasSamlCertificateName() bool`

HasSamlCertificateName returns a boolean if a field has been set.

### GetSamlCertificateIssuedTo

`func (o *SsoConfiguration) GetSamlCertificateIssuedTo() string`

GetSamlCertificateIssuedTo returns the SamlCertificateIssuedTo field if non-nil, zero value otherwise.

### GetSamlCertificateIssuedToOk

`func (o *SsoConfiguration) GetSamlCertificateIssuedToOk() (*string, bool)`

GetSamlCertificateIssuedToOk returns a tuple with the SamlCertificateIssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertificateIssuedTo

`func (o *SsoConfiguration) SetSamlCertificateIssuedTo(v string)`

SetSamlCertificateIssuedTo sets SamlCertificateIssuedTo field to given value.

### HasSamlCertificateIssuedTo

`func (o *SsoConfiguration) HasSamlCertificateIssuedTo() bool`

HasSamlCertificateIssuedTo returns a boolean if a field has been set.

### GetSamlCertificateThumbprint

`func (o *SsoConfiguration) GetSamlCertificateThumbprint() string`

GetSamlCertificateThumbprint returns the SamlCertificateThumbprint field if non-nil, zero value otherwise.

### GetSamlCertificateThumbprintOk

`func (o *SsoConfiguration) GetSamlCertificateThumbprintOk() (*string, bool)`

GetSamlCertificateThumbprintOk returns a tuple with the SamlCertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertificateThumbprint

`func (o *SsoConfiguration) SetSamlCertificateThumbprint(v string)`

SetSamlCertificateThumbprint sets SamlCertificateThumbprint field to given value.

### HasSamlCertificateThumbprint

`func (o *SsoConfiguration) HasSamlCertificateThumbprint() bool`

HasSamlCertificateThumbprint returns a boolean if a field has been set.

### GetSamlCertificateValidFrom

`func (o *SsoConfiguration) GetSamlCertificateValidFrom() time.Time`

GetSamlCertificateValidFrom returns the SamlCertificateValidFrom field if non-nil, zero value otherwise.

### GetSamlCertificateValidFromOk

`func (o *SsoConfiguration) GetSamlCertificateValidFromOk() (*time.Time, bool)`

GetSamlCertificateValidFromOk returns a tuple with the SamlCertificateValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertificateValidFrom

`func (o *SsoConfiguration) SetSamlCertificateValidFrom(v time.Time)`

SetSamlCertificateValidFrom sets SamlCertificateValidFrom field to given value.

### HasSamlCertificateValidFrom

`func (o *SsoConfiguration) HasSamlCertificateValidFrom() bool`

HasSamlCertificateValidFrom returns a boolean if a field has been set.

### GetSamlCertificateValidTo

`func (o *SsoConfiguration) GetSamlCertificateValidTo() time.Time`

GetSamlCertificateValidTo returns the SamlCertificateValidTo field if non-nil, zero value otherwise.

### GetSamlCertificateValidToOk

`func (o *SsoConfiguration) GetSamlCertificateValidToOk() (*time.Time, bool)`

GetSamlCertificateValidToOk returns a tuple with the SamlCertificateValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlCertificateValidTo

`func (o *SsoConfiguration) SetSamlCertificateValidTo(v time.Time)`

SetSamlCertificateValidTo sets SamlCertificateValidTo field to given value.

### HasSamlCertificateValidTo

`func (o *SsoConfiguration) HasSamlCertificateValidTo() bool`

HasSamlCertificateValidTo returns a boolean if a field has been set.

### GetLocationIds

`func (o *SsoConfiguration) GetLocationIds() []int32`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *SsoConfiguration) GetLocationIdsOk() (*[]int32, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *SsoConfiguration) SetLocationIds(v []int32)`

SetLocationIds sets LocationIds field to given value.


### GetClientId

`func (o *SsoConfiguration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SsoConfiguration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SsoConfiguration) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SsoConfiguration) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetStsBaseUrl

`func (o *SsoConfiguration) GetStsBaseUrl() string`

GetStsBaseUrl returns the StsBaseUrl field if non-nil, zero value otherwise.

### GetStsBaseUrlOk

`func (o *SsoConfiguration) GetStsBaseUrlOk() (*string, bool)`

GetStsBaseUrlOk returns a tuple with the StsBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsBaseUrl

`func (o *SsoConfiguration) SetStsBaseUrl(v string)`

SetStsBaseUrl sets StsBaseUrl field to given value.

### HasStsBaseUrl

`func (o *SsoConfiguration) HasStsBaseUrl() bool`

HasStsBaseUrl returns a boolean if a field has been set.

### GetStsUserAdminUrl

`func (o *SsoConfiguration) GetStsUserAdminUrl() string`

GetStsUserAdminUrl returns the StsUserAdminUrl field if non-nil, zero value otherwise.

### GetStsUserAdminUrlOk

`func (o *SsoConfiguration) GetStsUserAdminUrlOk() (*string, bool)`

GetStsUserAdminUrlOk returns a tuple with the StsUserAdminUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsUserAdminUrl

`func (o *SsoConfiguration) SetStsUserAdminUrl(v string)`

SetStsUserAdminUrl sets StsUserAdminUrl field to given value.

### HasStsUserAdminUrl

`func (o *SsoConfiguration) HasStsUserAdminUrl() bool`

HasStsUserAdminUrl returns a boolean if a field has been set.

### GetToken

`func (o *SsoConfiguration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SsoConfiguration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SsoConfiguration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SsoConfiguration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetSubmittedMemberCount

`func (o *SsoConfiguration) GetSubmittedMemberCount() int32`

GetSubmittedMemberCount returns the SubmittedMemberCount field if non-nil, zero value otherwise.

### GetSubmittedMemberCountOk

`func (o *SsoConfiguration) GetSubmittedMemberCountOk() (*int32, bool)`

GetSubmittedMemberCountOk returns a tuple with the SubmittedMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedMemberCount

`func (o *SsoConfiguration) SetSubmittedMemberCount(v int32)`

SetSubmittedMemberCount sets SubmittedMemberCount field to given value.

### HasSubmittedMemberCount

`func (o *SsoConfiguration) HasSubmittedMemberCount() bool`

HasSubmittedMemberCount returns a boolean if a field has been set.

### GetAllMembersSubmitted

`func (o *SsoConfiguration) GetAllMembersSubmitted() bool`

GetAllMembersSubmitted returns the AllMembersSubmitted field if non-nil, zero value otherwise.

### GetAllMembersSubmittedOk

`func (o *SsoConfiguration) GetAllMembersSubmittedOk() (*bool, bool)`

GetAllMembersSubmittedOk returns a tuple with the AllMembersSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMembersSubmitted

`func (o *SsoConfiguration) SetAllMembersSubmitted(v bool)`

SetAllMembersSubmitted sets AllMembersSubmitted field to given value.

### HasAllMembersSubmitted

`func (o *SsoConfiguration) HasAllMembersSubmitted() bool`

HasAllMembersSubmitted returns a boolean if a field has been set.

### GetInfo

`func (o *SsoConfiguration) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SsoConfiguration) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SsoConfiguration) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SsoConfiguration) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetIsSsoOnByDefault

`func (o *SsoConfiguration) GetIsSsoOnByDefault() bool`

GetIsSsoOnByDefault returns the IsSsoOnByDefault field if non-nil, zero value otherwise.

### GetIsSsoOnByDefaultOk

`func (o *SsoConfiguration) GetIsSsoOnByDefaultOk() (*bool, bool)`

GetIsSsoOnByDefaultOk returns a tuple with the IsSsoOnByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoOnByDefault

`func (o *SsoConfiguration) SetIsSsoOnByDefault(v bool)`

SetIsSsoOnByDefault sets IsSsoOnByDefault field to given value.

### HasIsSsoOnByDefault

`func (o *SsoConfiguration) HasIsSsoOnByDefault() bool`

HasIsSsoOnByDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


