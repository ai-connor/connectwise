# DocumentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**ServerFileName** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**LinkFlag** | Pointer to **NullableBool** |  | [optional] 
**ImageFlag** | Pointer to **NullableBool** |  | [optional] 
**PublicFlag** | Pointer to **NullableBool** |  | [optional] 
**HtmlTemplateFlag** | Pointer to **NullableBool** |  | [optional] 
**ReadOnlyFlag** | Pointer to **NullableBool** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**UrlFlag** | Pointer to **NullableBool** |  | [optional] 
**CreatedOnDate** | Pointer to **string** |  | [optional] 
**DocumentType** | Pointer to [**DocumentTypeReference**](DocumentTypeReference.md) |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDocumentInfo

`func NewDocumentInfo() *DocumentInfo`

NewDocumentInfo instantiates a new DocumentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentInfoWithDefaults

`func NewDocumentInfoWithDefaults() *DocumentInfo`

NewDocumentInfoWithDefaults instantiates a new DocumentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *DocumentInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFileName

`func (o *DocumentInfo) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DocumentInfo) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DocumentInfo) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DocumentInfo) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetServerFileName

`func (o *DocumentInfo) GetServerFileName() string`

GetServerFileName returns the ServerFileName field if non-nil, zero value otherwise.

### GetServerFileNameOk

`func (o *DocumentInfo) GetServerFileNameOk() (*string, bool)`

GetServerFileNameOk returns a tuple with the ServerFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFileName

`func (o *DocumentInfo) SetServerFileName(v string)`

SetServerFileName sets ServerFileName field to given value.

### HasServerFileName

`func (o *DocumentInfo) HasServerFileName() bool`

HasServerFileName returns a boolean if a field has been set.

### GetOwner

`func (o *DocumentInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DocumentInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DocumentInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DocumentInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetLinkFlag

`func (o *DocumentInfo) GetLinkFlag() bool`

GetLinkFlag returns the LinkFlag field if non-nil, zero value otherwise.

### GetLinkFlagOk

`func (o *DocumentInfo) GetLinkFlagOk() (*bool, bool)`

GetLinkFlagOk returns a tuple with the LinkFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkFlag

`func (o *DocumentInfo) SetLinkFlag(v bool)`

SetLinkFlag sets LinkFlag field to given value.

### HasLinkFlag

`func (o *DocumentInfo) HasLinkFlag() bool`

HasLinkFlag returns a boolean if a field has been set.

### SetLinkFlagNil

`func (o *DocumentInfo) SetLinkFlagNil(b bool)`

 SetLinkFlagNil sets the value for LinkFlag to be an explicit nil

### UnsetLinkFlag
`func (o *DocumentInfo) UnsetLinkFlag()`

UnsetLinkFlag ensures that no value is present for LinkFlag, not even an explicit nil
### GetImageFlag

`func (o *DocumentInfo) GetImageFlag() bool`

GetImageFlag returns the ImageFlag field if non-nil, zero value otherwise.

### GetImageFlagOk

`func (o *DocumentInfo) GetImageFlagOk() (*bool, bool)`

GetImageFlagOk returns a tuple with the ImageFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFlag

`func (o *DocumentInfo) SetImageFlag(v bool)`

SetImageFlag sets ImageFlag field to given value.

### HasImageFlag

`func (o *DocumentInfo) HasImageFlag() bool`

HasImageFlag returns a boolean if a field has been set.

### SetImageFlagNil

`func (o *DocumentInfo) SetImageFlagNil(b bool)`

 SetImageFlagNil sets the value for ImageFlag to be an explicit nil

### UnsetImageFlag
`func (o *DocumentInfo) UnsetImageFlag()`

UnsetImageFlag ensures that no value is present for ImageFlag, not even an explicit nil
### GetPublicFlag

`func (o *DocumentInfo) GetPublicFlag() bool`

GetPublicFlag returns the PublicFlag field if non-nil, zero value otherwise.

### GetPublicFlagOk

`func (o *DocumentInfo) GetPublicFlagOk() (*bool, bool)`

GetPublicFlagOk returns a tuple with the PublicFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicFlag

`func (o *DocumentInfo) SetPublicFlag(v bool)`

SetPublicFlag sets PublicFlag field to given value.

### HasPublicFlag

`func (o *DocumentInfo) HasPublicFlag() bool`

HasPublicFlag returns a boolean if a field has been set.

### SetPublicFlagNil

`func (o *DocumentInfo) SetPublicFlagNil(b bool)`

 SetPublicFlagNil sets the value for PublicFlag to be an explicit nil

### UnsetPublicFlag
`func (o *DocumentInfo) UnsetPublicFlag()`

UnsetPublicFlag ensures that no value is present for PublicFlag, not even an explicit nil
### GetHtmlTemplateFlag

`func (o *DocumentInfo) GetHtmlTemplateFlag() bool`

GetHtmlTemplateFlag returns the HtmlTemplateFlag field if non-nil, zero value otherwise.

### GetHtmlTemplateFlagOk

`func (o *DocumentInfo) GetHtmlTemplateFlagOk() (*bool, bool)`

GetHtmlTemplateFlagOk returns a tuple with the HtmlTemplateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlTemplateFlag

`func (o *DocumentInfo) SetHtmlTemplateFlag(v bool)`

SetHtmlTemplateFlag sets HtmlTemplateFlag field to given value.

### HasHtmlTemplateFlag

`func (o *DocumentInfo) HasHtmlTemplateFlag() bool`

HasHtmlTemplateFlag returns a boolean if a field has been set.

### SetHtmlTemplateFlagNil

`func (o *DocumentInfo) SetHtmlTemplateFlagNil(b bool)`

 SetHtmlTemplateFlagNil sets the value for HtmlTemplateFlag to be an explicit nil

### UnsetHtmlTemplateFlag
`func (o *DocumentInfo) UnsetHtmlTemplateFlag()`

UnsetHtmlTemplateFlag ensures that no value is present for HtmlTemplateFlag, not even an explicit nil
### GetReadOnlyFlag

`func (o *DocumentInfo) GetReadOnlyFlag() bool`

GetReadOnlyFlag returns the ReadOnlyFlag field if non-nil, zero value otherwise.

### GetReadOnlyFlagOk

`func (o *DocumentInfo) GetReadOnlyFlagOk() (*bool, bool)`

GetReadOnlyFlagOk returns a tuple with the ReadOnlyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyFlag

`func (o *DocumentInfo) SetReadOnlyFlag(v bool)`

SetReadOnlyFlag sets ReadOnlyFlag field to given value.

### HasReadOnlyFlag

`func (o *DocumentInfo) HasReadOnlyFlag() bool`

HasReadOnlyFlag returns a boolean if a field has been set.

### SetReadOnlyFlagNil

`func (o *DocumentInfo) SetReadOnlyFlagNil(b bool)`

 SetReadOnlyFlagNil sets the value for ReadOnlyFlag to be an explicit nil

### UnsetReadOnlyFlag
`func (o *DocumentInfo) UnsetReadOnlyFlag()`

UnsetReadOnlyFlag ensures that no value is present for ReadOnlyFlag, not even an explicit nil
### GetSize

`func (o *DocumentInfo) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DocumentInfo) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DocumentInfo) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *DocumentInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *DocumentInfo) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *DocumentInfo) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetUrlFlag

`func (o *DocumentInfo) GetUrlFlag() bool`

GetUrlFlag returns the UrlFlag field if non-nil, zero value otherwise.

### GetUrlFlagOk

`func (o *DocumentInfo) GetUrlFlagOk() (*bool, bool)`

GetUrlFlagOk returns a tuple with the UrlFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlFlag

`func (o *DocumentInfo) SetUrlFlag(v bool)`

SetUrlFlag sets UrlFlag field to given value.

### HasUrlFlag

`func (o *DocumentInfo) HasUrlFlag() bool`

HasUrlFlag returns a boolean if a field has been set.

### SetUrlFlagNil

`func (o *DocumentInfo) SetUrlFlagNil(b bool)`

 SetUrlFlagNil sets the value for UrlFlag to be an explicit nil

### UnsetUrlFlag
`func (o *DocumentInfo) UnsetUrlFlag()`

UnsetUrlFlag ensures that no value is present for UrlFlag, not even an explicit nil
### GetCreatedOnDate

`func (o *DocumentInfo) GetCreatedOnDate() string`

GetCreatedOnDate returns the CreatedOnDate field if non-nil, zero value otherwise.

### GetCreatedOnDateOk

`func (o *DocumentInfo) GetCreatedOnDateOk() (*string, bool)`

GetCreatedOnDateOk returns a tuple with the CreatedOnDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedOnDate

`func (o *DocumentInfo) SetCreatedOnDate(v string)`

SetCreatedOnDate sets CreatedOnDate field to given value.

### HasCreatedOnDate

`func (o *DocumentInfo) HasCreatedOnDate() bool`

HasCreatedOnDate returns a boolean if a field has been set.

### GetDocumentType

`func (o *DocumentInfo) GetDocumentType() DocumentTypeReference`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *DocumentInfo) GetDocumentTypeOk() (*DocumentTypeReference, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *DocumentInfo) SetDocumentType(v DocumentTypeReference)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *DocumentInfo) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetGuid

`func (o *DocumentInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *DocumentInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *DocumentInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *DocumentInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetInfo

`func (o *DocumentInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DocumentInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DocumentInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DocumentInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


