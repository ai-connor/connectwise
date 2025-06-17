# PortalConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Gets or sets and Sets             An existing Portal Configuration id is required when copying a Portal Configuration. | [optional] 
**Name** | **string** |  Max length: 150; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**LoginBackgroundColor** | Pointer to **string** |  Max length: 7; | [optional] 
**PortalBackgroundColor** | Pointer to **string** |  Max length: 7; | [optional] 
**MenuColor** | Pointer to **string** |  Max length: 7; | [optional] 
**ButtonColor** | Pointer to **string** |  Max length: 7; | [optional] 
**HeaderColor** | Pointer to **string** |  Max length: 7; | [optional] 
**Url** | Pointer to **string** |  Max length: 1000; | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**WelcomeText** | Pointer to **string** |  Max length: 4000; | [optional] 
**BoardIds** | Pointer to **[]int32** |  | [optional] 
**AgreementTypeIds** | Pointer to **[]int32** |  | [optional] 
**ConfigTypeIds** | Pointer to **[]int32** |  | [optional] 
**LocationIds** | Pointer to **[]int32** |  | [optional] 
**PortalImageCopySuccessFlag** | Pointer to **NullableBool** |  | [optional] 
**DisplayVendorFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalConfiguration

`func NewPortalConfiguration(name string, ) *PortalConfiguration`

NewPortalConfiguration instantiates a new PortalConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalConfigurationWithDefaults

`func NewPortalConfigurationWithDefaults() *PortalConfiguration`

NewPortalConfigurationWithDefaults instantiates a new PortalConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PortalConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortalConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortalConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *PortalConfiguration) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *PortalConfiguration) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *PortalConfiguration) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *PortalConfiguration) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *PortalConfiguration) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *PortalConfiguration) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetCompany

`func (o *PortalConfiguration) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *PortalConfiguration) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *PortalConfiguration) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *PortalConfiguration) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLoginBackgroundColor

`func (o *PortalConfiguration) GetLoginBackgroundColor() string`

GetLoginBackgroundColor returns the LoginBackgroundColor field if non-nil, zero value otherwise.

### GetLoginBackgroundColorOk

`func (o *PortalConfiguration) GetLoginBackgroundColorOk() (*string, bool)`

GetLoginBackgroundColorOk returns a tuple with the LoginBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBackgroundColor

`func (o *PortalConfiguration) SetLoginBackgroundColor(v string)`

SetLoginBackgroundColor sets LoginBackgroundColor field to given value.

### HasLoginBackgroundColor

`func (o *PortalConfiguration) HasLoginBackgroundColor() bool`

HasLoginBackgroundColor returns a boolean if a field has been set.

### GetPortalBackgroundColor

`func (o *PortalConfiguration) GetPortalBackgroundColor() string`

GetPortalBackgroundColor returns the PortalBackgroundColor field if non-nil, zero value otherwise.

### GetPortalBackgroundColorOk

`func (o *PortalConfiguration) GetPortalBackgroundColorOk() (*string, bool)`

GetPortalBackgroundColorOk returns a tuple with the PortalBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalBackgroundColor

`func (o *PortalConfiguration) SetPortalBackgroundColor(v string)`

SetPortalBackgroundColor sets PortalBackgroundColor field to given value.

### HasPortalBackgroundColor

`func (o *PortalConfiguration) HasPortalBackgroundColor() bool`

HasPortalBackgroundColor returns a boolean if a field has been set.

### GetMenuColor

`func (o *PortalConfiguration) GetMenuColor() string`

GetMenuColor returns the MenuColor field if non-nil, zero value otherwise.

### GetMenuColorOk

`func (o *PortalConfiguration) GetMenuColorOk() (*string, bool)`

GetMenuColorOk returns a tuple with the MenuColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuColor

`func (o *PortalConfiguration) SetMenuColor(v string)`

SetMenuColor sets MenuColor field to given value.

### HasMenuColor

`func (o *PortalConfiguration) HasMenuColor() bool`

HasMenuColor returns a boolean if a field has been set.

### GetButtonColor

`func (o *PortalConfiguration) GetButtonColor() string`

GetButtonColor returns the ButtonColor field if non-nil, zero value otherwise.

### GetButtonColorOk

`func (o *PortalConfiguration) GetButtonColorOk() (*string, bool)`

GetButtonColorOk returns a tuple with the ButtonColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonColor

`func (o *PortalConfiguration) SetButtonColor(v string)`

SetButtonColor sets ButtonColor field to given value.

### HasButtonColor

`func (o *PortalConfiguration) HasButtonColor() bool`

HasButtonColor returns a boolean if a field has been set.

### GetHeaderColor

`func (o *PortalConfiguration) GetHeaderColor() string`

GetHeaderColor returns the HeaderColor field if non-nil, zero value otherwise.

### GetHeaderColorOk

`func (o *PortalConfiguration) GetHeaderColorOk() (*string, bool)`

GetHeaderColorOk returns a tuple with the HeaderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderColor

`func (o *PortalConfiguration) SetHeaderColor(v string)`

SetHeaderColor sets HeaderColor field to given value.

### HasHeaderColor

`func (o *PortalConfiguration) HasHeaderColor() bool`

HasHeaderColor returns a boolean if a field has been set.

### GetUrl

`func (o *PortalConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PortalConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PortalConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PortalConfiguration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLanguage

`func (o *PortalConfiguration) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *PortalConfiguration) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *PortalConfiguration) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *PortalConfiguration) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *PortalConfiguration) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *PortalConfiguration) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetWelcomeText

`func (o *PortalConfiguration) GetWelcomeText() string`

GetWelcomeText returns the WelcomeText field if non-nil, zero value otherwise.

### GetWelcomeTextOk

`func (o *PortalConfiguration) GetWelcomeTextOk() (*string, bool)`

GetWelcomeTextOk returns a tuple with the WelcomeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeText

`func (o *PortalConfiguration) SetWelcomeText(v string)`

SetWelcomeText sets WelcomeText field to given value.

### HasWelcomeText

`func (o *PortalConfiguration) HasWelcomeText() bool`

HasWelcomeText returns a boolean if a field has been set.

### GetBoardIds

`func (o *PortalConfiguration) GetBoardIds() []int32`

GetBoardIds returns the BoardIds field if non-nil, zero value otherwise.

### GetBoardIdsOk

`func (o *PortalConfiguration) GetBoardIdsOk() (*[]int32, bool)`

GetBoardIdsOk returns a tuple with the BoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardIds

`func (o *PortalConfiguration) SetBoardIds(v []int32)`

SetBoardIds sets BoardIds field to given value.

### HasBoardIds

`func (o *PortalConfiguration) HasBoardIds() bool`

HasBoardIds returns a boolean if a field has been set.

### GetAgreementTypeIds

`func (o *PortalConfiguration) GetAgreementTypeIds() []int32`

GetAgreementTypeIds returns the AgreementTypeIds field if non-nil, zero value otherwise.

### GetAgreementTypeIdsOk

`func (o *PortalConfiguration) GetAgreementTypeIdsOk() (*[]int32, bool)`

GetAgreementTypeIdsOk returns a tuple with the AgreementTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementTypeIds

`func (o *PortalConfiguration) SetAgreementTypeIds(v []int32)`

SetAgreementTypeIds sets AgreementTypeIds field to given value.

### HasAgreementTypeIds

`func (o *PortalConfiguration) HasAgreementTypeIds() bool`

HasAgreementTypeIds returns a boolean if a field has been set.

### GetConfigTypeIds

`func (o *PortalConfiguration) GetConfigTypeIds() []int32`

GetConfigTypeIds returns the ConfigTypeIds field if non-nil, zero value otherwise.

### GetConfigTypeIdsOk

`func (o *PortalConfiguration) GetConfigTypeIdsOk() (*[]int32, bool)`

GetConfigTypeIdsOk returns a tuple with the ConfigTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTypeIds

`func (o *PortalConfiguration) SetConfigTypeIds(v []int32)`

SetConfigTypeIds sets ConfigTypeIds field to given value.

### HasConfigTypeIds

`func (o *PortalConfiguration) HasConfigTypeIds() bool`

HasConfigTypeIds returns a boolean if a field has been set.

### GetLocationIds

`func (o *PortalConfiguration) GetLocationIds() []int32`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *PortalConfiguration) GetLocationIdsOk() (*[]int32, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *PortalConfiguration) SetLocationIds(v []int32)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *PortalConfiguration) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetPortalImageCopySuccessFlag

`func (o *PortalConfiguration) GetPortalImageCopySuccessFlag() bool`

GetPortalImageCopySuccessFlag returns the PortalImageCopySuccessFlag field if non-nil, zero value otherwise.

### GetPortalImageCopySuccessFlagOk

`func (o *PortalConfiguration) GetPortalImageCopySuccessFlagOk() (*bool, bool)`

GetPortalImageCopySuccessFlagOk returns a tuple with the PortalImageCopySuccessFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalImageCopySuccessFlag

`func (o *PortalConfiguration) SetPortalImageCopySuccessFlag(v bool)`

SetPortalImageCopySuccessFlag sets PortalImageCopySuccessFlag field to given value.

### HasPortalImageCopySuccessFlag

`func (o *PortalConfiguration) HasPortalImageCopySuccessFlag() bool`

HasPortalImageCopySuccessFlag returns a boolean if a field has been set.

### SetPortalImageCopySuccessFlagNil

`func (o *PortalConfiguration) SetPortalImageCopySuccessFlagNil(b bool)`

 SetPortalImageCopySuccessFlagNil sets the value for PortalImageCopySuccessFlag to be an explicit nil

### UnsetPortalImageCopySuccessFlag
`func (o *PortalConfiguration) UnsetPortalImageCopySuccessFlag()`

UnsetPortalImageCopySuccessFlag ensures that no value is present for PortalImageCopySuccessFlag, not even an explicit nil
### GetDisplayVendorFlag

`func (o *PortalConfiguration) GetDisplayVendorFlag() bool`

GetDisplayVendorFlag returns the DisplayVendorFlag field if non-nil, zero value otherwise.

### GetDisplayVendorFlagOk

`func (o *PortalConfiguration) GetDisplayVendorFlagOk() (*bool, bool)`

GetDisplayVendorFlagOk returns a tuple with the DisplayVendorFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayVendorFlag

`func (o *PortalConfiguration) SetDisplayVendorFlag(v bool)`

SetDisplayVendorFlag sets DisplayVendorFlag field to given value.

### HasDisplayVendorFlag

`func (o *PortalConfiguration) HasDisplayVendorFlag() bool`

HasDisplayVendorFlag returns a boolean if a field has been set.

### SetDisplayVendorFlagNil

`func (o *PortalConfiguration) SetDisplayVendorFlagNil(b bool)`

 SetDisplayVendorFlagNil sets the value for DisplayVendorFlag to be an explicit nil

### UnsetDisplayVendorFlag
`func (o *PortalConfiguration) UnsetDisplayVendorFlag()`

UnsetDisplayVendorFlag ensures that no value is present for DisplayVendorFlag, not even an explicit nil
### GetInfo

`func (o *PortalConfiguration) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalConfiguration) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalConfiguration) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalConfiguration) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


