# ManagementItSolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**ManagementItSolutionType** | **NullableString** |  | 
**ManagementSolutionName** | Pointer to **string** | Gets or sets             this is only required when managementItSolutionType is Custom. Max length: 30; | [optional] 
**ManagementServerUrl** | Pointer to **string** | Gets or sets             this is only required for Level Platforms. Max length: 200; | [optional] 
**WebserviceOverrideUrl** | Pointer to **string** | Gets or sets             this is only required for Level Platforms when overrideWebServiceLocationFlag is true. Max length: 200; | [optional] 
**PortalOverrideLoginUrl** | Pointer to **string** | Gets or sets             this is only required for Level Platforms when overrideLoginLocationFlag is true. Max length: 200; | [optional] 
**GlobalLoginFlag** | Pointer to **NullableBool** |  | [optional] 
**GlobalLoginUsername** | Pointer to **string** | Gets or sets             this is only required when globalLoginFlag &#x3D; true. Max length: 50; | [optional] 
**GlobalLoginPassword** | Pointer to **string** | Gets or sets             this is only required when globalLoginFlag &#x3D; true. Max length: 50; | [optional] 
**UsingSslFlag** | Pointer to **NullableBool** |  | [optional] 
**NAbleUsername** | Pointer to **string** | Gets or sets             this is only required for N-Able solution. Max length: 50; | [optional] 
**NAblePassword** | Pointer to **string** | Gets or sets             this is only required for N-Able solution. Max length: 50; | [optional] 
**OverrideWebServiceLocationFlag** | Pointer to **NullableBool** |  | [optional] 
**OverrideLoginLocationFlag** | Pointer to **NullableBool** |  | [optional] 
**ContinuumApiUsername** | Pointer to **string** | Gets or sets             this is only required for Continuum solution. Max length: 100; | [optional] 
**ContinuumApiPassword** | Pointer to **string** | Gets or sets             this is only required for Continuum solution. Max length: 100; | [optional] 
**LevelApiUsername** | Pointer to **string** | Gets or sets             this is only required for Level Platforms solution. Max length: 100; | [optional] 
**LevelApiPassword** | Pointer to **string** | Gets or sets             this is only required for Level Platforms solution. Max length: 100; | [optional] 
**LevelVarDomain** | Pointer to **string** | Gets or sets             this is only required for Level Platforms solution. Max length: 100; | [optional] 
**NoDisplayFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagementItSolution

`func NewManagementItSolution(name string, managementItSolutionType NullableString, ) *ManagementItSolution`

NewManagementItSolution instantiates a new ManagementItSolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementItSolutionWithDefaults

`func NewManagementItSolutionWithDefaults() *ManagementItSolution`

NewManagementItSolutionWithDefaults instantiates a new ManagementItSolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagementItSolution) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementItSolution) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementItSolution) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagementItSolution) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManagementItSolution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagementItSolution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagementItSolution) SetName(v string)`

SetName sets Name field to given value.


### GetManagementItSolutionType

`func (o *ManagementItSolution) GetManagementItSolutionType() string`

GetManagementItSolutionType returns the ManagementItSolutionType field if non-nil, zero value otherwise.

### GetManagementItSolutionTypeOk

`func (o *ManagementItSolution) GetManagementItSolutionTypeOk() (*string, bool)`

GetManagementItSolutionTypeOk returns a tuple with the ManagementItSolutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementItSolutionType

`func (o *ManagementItSolution) SetManagementItSolutionType(v string)`

SetManagementItSolutionType sets ManagementItSolutionType field to given value.


### SetManagementItSolutionTypeNil

`func (o *ManagementItSolution) SetManagementItSolutionTypeNil(b bool)`

 SetManagementItSolutionTypeNil sets the value for ManagementItSolutionType to be an explicit nil

### UnsetManagementItSolutionType
`func (o *ManagementItSolution) UnsetManagementItSolutionType()`

UnsetManagementItSolutionType ensures that no value is present for ManagementItSolutionType, not even an explicit nil
### GetManagementSolutionName

`func (o *ManagementItSolution) GetManagementSolutionName() string`

GetManagementSolutionName returns the ManagementSolutionName field if non-nil, zero value otherwise.

### GetManagementSolutionNameOk

`func (o *ManagementItSolution) GetManagementSolutionNameOk() (*string, bool)`

GetManagementSolutionNameOk returns a tuple with the ManagementSolutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementSolutionName

`func (o *ManagementItSolution) SetManagementSolutionName(v string)`

SetManagementSolutionName sets ManagementSolutionName field to given value.

### HasManagementSolutionName

`func (o *ManagementItSolution) HasManagementSolutionName() bool`

HasManagementSolutionName returns a boolean if a field has been set.

### GetManagementServerUrl

`func (o *ManagementItSolution) GetManagementServerUrl() string`

GetManagementServerUrl returns the ManagementServerUrl field if non-nil, zero value otherwise.

### GetManagementServerUrlOk

`func (o *ManagementItSolution) GetManagementServerUrlOk() (*string, bool)`

GetManagementServerUrlOk returns a tuple with the ManagementServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementServerUrl

`func (o *ManagementItSolution) SetManagementServerUrl(v string)`

SetManagementServerUrl sets ManagementServerUrl field to given value.

### HasManagementServerUrl

`func (o *ManagementItSolution) HasManagementServerUrl() bool`

HasManagementServerUrl returns a boolean if a field has been set.

### GetWebserviceOverrideUrl

`func (o *ManagementItSolution) GetWebserviceOverrideUrl() string`

GetWebserviceOverrideUrl returns the WebserviceOverrideUrl field if non-nil, zero value otherwise.

### GetWebserviceOverrideUrlOk

`func (o *ManagementItSolution) GetWebserviceOverrideUrlOk() (*string, bool)`

GetWebserviceOverrideUrlOk returns a tuple with the WebserviceOverrideUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebserviceOverrideUrl

`func (o *ManagementItSolution) SetWebserviceOverrideUrl(v string)`

SetWebserviceOverrideUrl sets WebserviceOverrideUrl field to given value.

### HasWebserviceOverrideUrl

`func (o *ManagementItSolution) HasWebserviceOverrideUrl() bool`

HasWebserviceOverrideUrl returns a boolean if a field has been set.

### GetPortalOverrideLoginUrl

`func (o *ManagementItSolution) GetPortalOverrideLoginUrl() string`

GetPortalOverrideLoginUrl returns the PortalOverrideLoginUrl field if non-nil, zero value otherwise.

### GetPortalOverrideLoginUrlOk

`func (o *ManagementItSolution) GetPortalOverrideLoginUrlOk() (*string, bool)`

GetPortalOverrideLoginUrlOk returns a tuple with the PortalOverrideLoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalOverrideLoginUrl

`func (o *ManagementItSolution) SetPortalOverrideLoginUrl(v string)`

SetPortalOverrideLoginUrl sets PortalOverrideLoginUrl field to given value.

### HasPortalOverrideLoginUrl

`func (o *ManagementItSolution) HasPortalOverrideLoginUrl() bool`

HasPortalOverrideLoginUrl returns a boolean if a field has been set.

### GetGlobalLoginFlag

`func (o *ManagementItSolution) GetGlobalLoginFlag() bool`

GetGlobalLoginFlag returns the GlobalLoginFlag field if non-nil, zero value otherwise.

### GetGlobalLoginFlagOk

`func (o *ManagementItSolution) GetGlobalLoginFlagOk() (*bool, bool)`

GetGlobalLoginFlagOk returns a tuple with the GlobalLoginFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalLoginFlag

`func (o *ManagementItSolution) SetGlobalLoginFlag(v bool)`

SetGlobalLoginFlag sets GlobalLoginFlag field to given value.

### HasGlobalLoginFlag

`func (o *ManagementItSolution) HasGlobalLoginFlag() bool`

HasGlobalLoginFlag returns a boolean if a field has been set.

### SetGlobalLoginFlagNil

`func (o *ManagementItSolution) SetGlobalLoginFlagNil(b bool)`

 SetGlobalLoginFlagNil sets the value for GlobalLoginFlag to be an explicit nil

### UnsetGlobalLoginFlag
`func (o *ManagementItSolution) UnsetGlobalLoginFlag()`

UnsetGlobalLoginFlag ensures that no value is present for GlobalLoginFlag, not even an explicit nil
### GetGlobalLoginUsername

`func (o *ManagementItSolution) GetGlobalLoginUsername() string`

GetGlobalLoginUsername returns the GlobalLoginUsername field if non-nil, zero value otherwise.

### GetGlobalLoginUsernameOk

`func (o *ManagementItSolution) GetGlobalLoginUsernameOk() (*string, bool)`

GetGlobalLoginUsernameOk returns a tuple with the GlobalLoginUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalLoginUsername

`func (o *ManagementItSolution) SetGlobalLoginUsername(v string)`

SetGlobalLoginUsername sets GlobalLoginUsername field to given value.

### HasGlobalLoginUsername

`func (o *ManagementItSolution) HasGlobalLoginUsername() bool`

HasGlobalLoginUsername returns a boolean if a field has been set.

### GetGlobalLoginPassword

`func (o *ManagementItSolution) GetGlobalLoginPassword() string`

GetGlobalLoginPassword returns the GlobalLoginPassword field if non-nil, zero value otherwise.

### GetGlobalLoginPasswordOk

`func (o *ManagementItSolution) GetGlobalLoginPasswordOk() (*string, bool)`

GetGlobalLoginPasswordOk returns a tuple with the GlobalLoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalLoginPassword

`func (o *ManagementItSolution) SetGlobalLoginPassword(v string)`

SetGlobalLoginPassword sets GlobalLoginPassword field to given value.

### HasGlobalLoginPassword

`func (o *ManagementItSolution) HasGlobalLoginPassword() bool`

HasGlobalLoginPassword returns a boolean if a field has been set.

### GetUsingSslFlag

`func (o *ManagementItSolution) GetUsingSslFlag() bool`

GetUsingSslFlag returns the UsingSslFlag field if non-nil, zero value otherwise.

### GetUsingSslFlagOk

`func (o *ManagementItSolution) GetUsingSslFlagOk() (*bool, bool)`

GetUsingSslFlagOk returns a tuple with the UsingSslFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingSslFlag

`func (o *ManagementItSolution) SetUsingSslFlag(v bool)`

SetUsingSslFlag sets UsingSslFlag field to given value.

### HasUsingSslFlag

`func (o *ManagementItSolution) HasUsingSslFlag() bool`

HasUsingSslFlag returns a boolean if a field has been set.

### SetUsingSslFlagNil

`func (o *ManagementItSolution) SetUsingSslFlagNil(b bool)`

 SetUsingSslFlagNil sets the value for UsingSslFlag to be an explicit nil

### UnsetUsingSslFlag
`func (o *ManagementItSolution) UnsetUsingSslFlag()`

UnsetUsingSslFlag ensures that no value is present for UsingSslFlag, not even an explicit nil
### GetNAbleUsername

`func (o *ManagementItSolution) GetNAbleUsername() string`

GetNAbleUsername returns the NAbleUsername field if non-nil, zero value otherwise.

### GetNAbleUsernameOk

`func (o *ManagementItSolution) GetNAbleUsernameOk() (*string, bool)`

GetNAbleUsernameOk returns a tuple with the NAbleUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAbleUsername

`func (o *ManagementItSolution) SetNAbleUsername(v string)`

SetNAbleUsername sets NAbleUsername field to given value.

### HasNAbleUsername

`func (o *ManagementItSolution) HasNAbleUsername() bool`

HasNAbleUsername returns a boolean if a field has been set.

### GetNAblePassword

`func (o *ManagementItSolution) GetNAblePassword() string`

GetNAblePassword returns the NAblePassword field if non-nil, zero value otherwise.

### GetNAblePasswordOk

`func (o *ManagementItSolution) GetNAblePasswordOk() (*string, bool)`

GetNAblePasswordOk returns a tuple with the NAblePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNAblePassword

`func (o *ManagementItSolution) SetNAblePassword(v string)`

SetNAblePassword sets NAblePassword field to given value.

### HasNAblePassword

`func (o *ManagementItSolution) HasNAblePassword() bool`

HasNAblePassword returns a boolean if a field has been set.

### GetOverrideWebServiceLocationFlag

`func (o *ManagementItSolution) GetOverrideWebServiceLocationFlag() bool`

GetOverrideWebServiceLocationFlag returns the OverrideWebServiceLocationFlag field if non-nil, zero value otherwise.

### GetOverrideWebServiceLocationFlagOk

`func (o *ManagementItSolution) GetOverrideWebServiceLocationFlagOk() (*bool, bool)`

GetOverrideWebServiceLocationFlagOk returns a tuple with the OverrideWebServiceLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideWebServiceLocationFlag

`func (o *ManagementItSolution) SetOverrideWebServiceLocationFlag(v bool)`

SetOverrideWebServiceLocationFlag sets OverrideWebServiceLocationFlag field to given value.

### HasOverrideWebServiceLocationFlag

`func (o *ManagementItSolution) HasOverrideWebServiceLocationFlag() bool`

HasOverrideWebServiceLocationFlag returns a boolean if a field has been set.

### SetOverrideWebServiceLocationFlagNil

`func (o *ManagementItSolution) SetOverrideWebServiceLocationFlagNil(b bool)`

 SetOverrideWebServiceLocationFlagNil sets the value for OverrideWebServiceLocationFlag to be an explicit nil

### UnsetOverrideWebServiceLocationFlag
`func (o *ManagementItSolution) UnsetOverrideWebServiceLocationFlag()`

UnsetOverrideWebServiceLocationFlag ensures that no value is present for OverrideWebServiceLocationFlag, not even an explicit nil
### GetOverrideLoginLocationFlag

`func (o *ManagementItSolution) GetOverrideLoginLocationFlag() bool`

GetOverrideLoginLocationFlag returns the OverrideLoginLocationFlag field if non-nil, zero value otherwise.

### GetOverrideLoginLocationFlagOk

`func (o *ManagementItSolution) GetOverrideLoginLocationFlagOk() (*bool, bool)`

GetOverrideLoginLocationFlagOk returns a tuple with the OverrideLoginLocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLoginLocationFlag

`func (o *ManagementItSolution) SetOverrideLoginLocationFlag(v bool)`

SetOverrideLoginLocationFlag sets OverrideLoginLocationFlag field to given value.

### HasOverrideLoginLocationFlag

`func (o *ManagementItSolution) HasOverrideLoginLocationFlag() bool`

HasOverrideLoginLocationFlag returns a boolean if a field has been set.

### SetOverrideLoginLocationFlagNil

`func (o *ManagementItSolution) SetOverrideLoginLocationFlagNil(b bool)`

 SetOverrideLoginLocationFlagNil sets the value for OverrideLoginLocationFlag to be an explicit nil

### UnsetOverrideLoginLocationFlag
`func (o *ManagementItSolution) UnsetOverrideLoginLocationFlag()`

UnsetOverrideLoginLocationFlag ensures that no value is present for OverrideLoginLocationFlag, not even an explicit nil
### GetContinuumApiUsername

`func (o *ManagementItSolution) GetContinuumApiUsername() string`

GetContinuumApiUsername returns the ContinuumApiUsername field if non-nil, zero value otherwise.

### GetContinuumApiUsernameOk

`func (o *ManagementItSolution) GetContinuumApiUsernameOk() (*string, bool)`

GetContinuumApiUsernameOk returns a tuple with the ContinuumApiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuumApiUsername

`func (o *ManagementItSolution) SetContinuumApiUsername(v string)`

SetContinuumApiUsername sets ContinuumApiUsername field to given value.

### HasContinuumApiUsername

`func (o *ManagementItSolution) HasContinuumApiUsername() bool`

HasContinuumApiUsername returns a boolean if a field has been set.

### GetContinuumApiPassword

`func (o *ManagementItSolution) GetContinuumApiPassword() string`

GetContinuumApiPassword returns the ContinuumApiPassword field if non-nil, zero value otherwise.

### GetContinuumApiPasswordOk

`func (o *ManagementItSolution) GetContinuumApiPasswordOk() (*string, bool)`

GetContinuumApiPasswordOk returns a tuple with the ContinuumApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuumApiPassword

`func (o *ManagementItSolution) SetContinuumApiPassword(v string)`

SetContinuumApiPassword sets ContinuumApiPassword field to given value.

### HasContinuumApiPassword

`func (o *ManagementItSolution) HasContinuumApiPassword() bool`

HasContinuumApiPassword returns a boolean if a field has been set.

### GetLevelApiUsername

`func (o *ManagementItSolution) GetLevelApiUsername() string`

GetLevelApiUsername returns the LevelApiUsername field if non-nil, zero value otherwise.

### GetLevelApiUsernameOk

`func (o *ManagementItSolution) GetLevelApiUsernameOk() (*string, bool)`

GetLevelApiUsernameOk returns a tuple with the LevelApiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelApiUsername

`func (o *ManagementItSolution) SetLevelApiUsername(v string)`

SetLevelApiUsername sets LevelApiUsername field to given value.

### HasLevelApiUsername

`func (o *ManagementItSolution) HasLevelApiUsername() bool`

HasLevelApiUsername returns a boolean if a field has been set.

### GetLevelApiPassword

`func (o *ManagementItSolution) GetLevelApiPassword() string`

GetLevelApiPassword returns the LevelApiPassword field if non-nil, zero value otherwise.

### GetLevelApiPasswordOk

`func (o *ManagementItSolution) GetLevelApiPasswordOk() (*string, bool)`

GetLevelApiPasswordOk returns a tuple with the LevelApiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelApiPassword

`func (o *ManagementItSolution) SetLevelApiPassword(v string)`

SetLevelApiPassword sets LevelApiPassword field to given value.

### HasLevelApiPassword

`func (o *ManagementItSolution) HasLevelApiPassword() bool`

HasLevelApiPassword returns a boolean if a field has been set.

### GetLevelVarDomain

`func (o *ManagementItSolution) GetLevelVarDomain() string`

GetLevelVarDomain returns the LevelVarDomain field if non-nil, zero value otherwise.

### GetLevelVarDomainOk

`func (o *ManagementItSolution) GetLevelVarDomainOk() (*string, bool)`

GetLevelVarDomainOk returns a tuple with the LevelVarDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelVarDomain

`func (o *ManagementItSolution) SetLevelVarDomain(v string)`

SetLevelVarDomain sets LevelVarDomain field to given value.

### HasLevelVarDomain

`func (o *ManagementItSolution) HasLevelVarDomain() bool`

HasLevelVarDomain returns a boolean if a field has been set.

### GetNoDisplayFlag

`func (o *ManagementItSolution) GetNoDisplayFlag() bool`

GetNoDisplayFlag returns the NoDisplayFlag field if non-nil, zero value otherwise.

### GetNoDisplayFlagOk

`func (o *ManagementItSolution) GetNoDisplayFlagOk() (*bool, bool)`

GetNoDisplayFlagOk returns a tuple with the NoDisplayFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDisplayFlag

`func (o *ManagementItSolution) SetNoDisplayFlag(v bool)`

SetNoDisplayFlag sets NoDisplayFlag field to given value.

### HasNoDisplayFlag

`func (o *ManagementItSolution) HasNoDisplayFlag() bool`

HasNoDisplayFlag returns a boolean if a field has been set.

### SetNoDisplayFlagNil

`func (o *ManagementItSolution) SetNoDisplayFlagNil(b bool)`

 SetNoDisplayFlagNil sets the value for NoDisplayFlag to be an explicit nil

### UnsetNoDisplayFlag
`func (o *ManagementItSolution) UnsetNoDisplayFlag()`

UnsetNoDisplayFlag ensures that no value is present for NoDisplayFlag, not even an explicit nil
### GetInfo

`func (o *ManagementItSolution) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagementItSolution) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagementItSolution) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagementItSolution) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


