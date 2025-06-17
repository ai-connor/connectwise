# DocumentSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**UploadAsLinkFlag** | Pointer to **NullableBool** |  | [optional] 
**IsPublicFlag** | Pointer to **NullableBool** |  | [optional] 
**DocPath** | Pointer to **string** |  Max length: 100; | [optional] 
**TemplatePath** | Pointer to **string** |  Max length: 200; | [optional] 
**TemplateOutputPath** | Pointer to **string** |  Max length: 200; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDocumentSetup

`func NewDocumentSetup() *DocumentSetup`

NewDocumentSetup instantiates a new DocumentSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSetupWithDefaults

`func NewDocumentSetupWithDefaults() *DocumentSetup`

NewDocumentSetupWithDefaults instantiates a new DocumentSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUploadAsLinkFlag

`func (o *DocumentSetup) GetUploadAsLinkFlag() bool`

GetUploadAsLinkFlag returns the UploadAsLinkFlag field if non-nil, zero value otherwise.

### GetUploadAsLinkFlagOk

`func (o *DocumentSetup) GetUploadAsLinkFlagOk() (*bool, bool)`

GetUploadAsLinkFlagOk returns a tuple with the UploadAsLinkFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadAsLinkFlag

`func (o *DocumentSetup) SetUploadAsLinkFlag(v bool)`

SetUploadAsLinkFlag sets UploadAsLinkFlag field to given value.

### HasUploadAsLinkFlag

`func (o *DocumentSetup) HasUploadAsLinkFlag() bool`

HasUploadAsLinkFlag returns a boolean if a field has been set.

### SetUploadAsLinkFlagNil

`func (o *DocumentSetup) SetUploadAsLinkFlagNil(b bool)`

 SetUploadAsLinkFlagNil sets the value for UploadAsLinkFlag to be an explicit nil

### UnsetUploadAsLinkFlag
`func (o *DocumentSetup) UnsetUploadAsLinkFlag()`

UnsetUploadAsLinkFlag ensures that no value is present for UploadAsLinkFlag, not even an explicit nil
### GetIsPublicFlag

`func (o *DocumentSetup) GetIsPublicFlag() bool`

GetIsPublicFlag returns the IsPublicFlag field if non-nil, zero value otherwise.

### GetIsPublicFlagOk

`func (o *DocumentSetup) GetIsPublicFlagOk() (*bool, bool)`

GetIsPublicFlagOk returns a tuple with the IsPublicFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublicFlag

`func (o *DocumentSetup) SetIsPublicFlag(v bool)`

SetIsPublicFlag sets IsPublicFlag field to given value.

### HasIsPublicFlag

`func (o *DocumentSetup) HasIsPublicFlag() bool`

HasIsPublicFlag returns a boolean if a field has been set.

### SetIsPublicFlagNil

`func (o *DocumentSetup) SetIsPublicFlagNil(b bool)`

 SetIsPublicFlagNil sets the value for IsPublicFlag to be an explicit nil

### UnsetIsPublicFlag
`func (o *DocumentSetup) UnsetIsPublicFlag()`

UnsetIsPublicFlag ensures that no value is present for IsPublicFlag, not even an explicit nil
### GetDocPath

`func (o *DocumentSetup) GetDocPath() string`

GetDocPath returns the DocPath field if non-nil, zero value otherwise.

### GetDocPathOk

`func (o *DocumentSetup) GetDocPathOk() (*string, bool)`

GetDocPathOk returns a tuple with the DocPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocPath

`func (o *DocumentSetup) SetDocPath(v string)`

SetDocPath sets DocPath field to given value.

### HasDocPath

`func (o *DocumentSetup) HasDocPath() bool`

HasDocPath returns a boolean if a field has been set.

### GetTemplatePath

`func (o *DocumentSetup) GetTemplatePath() string`

GetTemplatePath returns the TemplatePath field if non-nil, zero value otherwise.

### GetTemplatePathOk

`func (o *DocumentSetup) GetTemplatePathOk() (*string, bool)`

GetTemplatePathOk returns a tuple with the TemplatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePath

`func (o *DocumentSetup) SetTemplatePath(v string)`

SetTemplatePath sets TemplatePath field to given value.

### HasTemplatePath

`func (o *DocumentSetup) HasTemplatePath() bool`

HasTemplatePath returns a boolean if a field has been set.

### GetTemplateOutputPath

`func (o *DocumentSetup) GetTemplateOutputPath() string`

GetTemplateOutputPath returns the TemplateOutputPath field if non-nil, zero value otherwise.

### GetTemplateOutputPathOk

`func (o *DocumentSetup) GetTemplateOutputPathOk() (*string, bool)`

GetTemplateOutputPathOk returns a tuple with the TemplateOutputPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateOutputPath

`func (o *DocumentSetup) SetTemplateOutputPath(v string)`

SetTemplateOutputPath sets TemplateOutputPath field to given value.

### HasTemplateOutputPath

`func (o *DocumentSetup) HasTemplateOutputPath() bool`

HasTemplateOutputPath returns a boolean if a field has been set.

### GetInfo

`func (o *DocumentSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DocumentSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DocumentSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DocumentSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


