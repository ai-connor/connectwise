# ApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** |  | [optional] 
**GrandParentId** | Pointer to **NullableInt32** |  | [optional] 
**Entity** | Pointer to [**IRestIdentifiedItem**](IRestIdentifiedItem.md) |  | [optional] 
**Filters** | Pointer to [**FilterValues**](FilterValues.md) |  | [optional] 
**Page** | Pointer to [**PageValues**](PageValues.md) |  | [optional] 
**Fields** | Pointer to **string** |  | [optional] 
**MiscProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**MemberContext** | Pointer to **string** |  | [optional] 
**UpdateOnlyCesProperties** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiRequest

`func NewApiRequest() *ApiRequest`

NewApiRequest instantiates a new ApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRequestWithDefaults

`func NewApiRequestWithDefaults() *ApiRequest`

NewApiRequestWithDefaults instantiates a new ApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApiRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExternalId

`func (o *ApiRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ApiRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ApiRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ApiRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetParentId

`func (o *ApiRequest) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ApiRequest) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ApiRequest) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ApiRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ApiRequest) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ApiRequest) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetGrandParentId

`func (o *ApiRequest) GetGrandParentId() int32`

GetGrandParentId returns the GrandParentId field if non-nil, zero value otherwise.

### GetGrandParentIdOk

`func (o *ApiRequest) GetGrandParentIdOk() (*int32, bool)`

GetGrandParentIdOk returns a tuple with the GrandParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandParentId

`func (o *ApiRequest) SetGrandParentId(v int32)`

SetGrandParentId sets GrandParentId field to given value.

### HasGrandParentId

`func (o *ApiRequest) HasGrandParentId() bool`

HasGrandParentId returns a boolean if a field has been set.

### SetGrandParentIdNil

`func (o *ApiRequest) SetGrandParentIdNil(b bool)`

 SetGrandParentIdNil sets the value for GrandParentId to be an explicit nil

### UnsetGrandParentId
`func (o *ApiRequest) UnsetGrandParentId()`

UnsetGrandParentId ensures that no value is present for GrandParentId, not even an explicit nil
### GetEntity

`func (o *ApiRequest) GetEntity() IRestIdentifiedItem`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ApiRequest) GetEntityOk() (*IRestIdentifiedItem, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ApiRequest) SetEntity(v IRestIdentifiedItem)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *ApiRequest) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetFilters

`func (o *ApiRequest) GetFilters() FilterValues`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ApiRequest) GetFiltersOk() (*FilterValues, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ApiRequest) SetFilters(v FilterValues)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ApiRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPage

`func (o *ApiRequest) GetPage() PageValues`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ApiRequest) GetPageOk() (*PageValues, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ApiRequest) SetPage(v PageValues)`

SetPage sets Page field to given value.

### HasPage

`func (o *ApiRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetFields

`func (o *ApiRequest) GetFields() string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ApiRequest) GetFieldsOk() (*string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ApiRequest) SetFields(v string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ApiRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMiscProperties

`func (o *ApiRequest) GetMiscProperties() map[string]map[string]interface{}`

GetMiscProperties returns the MiscProperties field if non-nil, zero value otherwise.

### GetMiscPropertiesOk

`func (o *ApiRequest) GetMiscPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetMiscPropertiesOk returns a tuple with the MiscProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscProperties

`func (o *ApiRequest) SetMiscProperties(v map[string]map[string]interface{})`

SetMiscProperties sets MiscProperties field to given value.

### HasMiscProperties

`func (o *ApiRequest) HasMiscProperties() bool`

HasMiscProperties returns a boolean if a field has been set.

### GetMemberContext

`func (o *ApiRequest) GetMemberContext() string`

GetMemberContext returns the MemberContext field if non-nil, zero value otherwise.

### GetMemberContextOk

`func (o *ApiRequest) GetMemberContextOk() (*string, bool)`

GetMemberContextOk returns a tuple with the MemberContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberContext

`func (o *ApiRequest) SetMemberContext(v string)`

SetMemberContext sets MemberContext field to given value.

### HasMemberContext

`func (o *ApiRequest) HasMemberContext() bool`

HasMemberContext returns a boolean if a field has been set.

### GetUpdateOnlyCesProperties

`func (o *ApiRequest) GetUpdateOnlyCesProperties() bool`

GetUpdateOnlyCesProperties returns the UpdateOnlyCesProperties field if non-nil, zero value otherwise.

### GetUpdateOnlyCesPropertiesOk

`func (o *ApiRequest) GetUpdateOnlyCesPropertiesOk() (*bool, bool)`

GetUpdateOnlyCesPropertiesOk returns a tuple with the UpdateOnlyCesProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOnlyCesProperties

`func (o *ApiRequest) SetUpdateOnlyCesProperties(v bool)`

SetUpdateOnlyCesProperties sets UpdateOnlyCesProperties field to given value.

### HasUpdateOnlyCesProperties

`func (o *ApiRequest) HasUpdateOnlyCesProperties() bool`

HasUpdateOnlyCesProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


