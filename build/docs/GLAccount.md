# GLAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**GlType** | **NullableString** |  | 
**MappedType** | [**MappedTypeReference**](MappedTypeReference.md) |  | 
**MappedRecord** | [**MappedRecordReference**](MappedRecordReference.md) |  | 
**Segment1** | Pointer to **string** |  Max length: 255; | [optional] 
**Segment2** | Pointer to **string** |  Max length: 255; | [optional] 
**Segment3** | Pointer to **string** |  Max length: 255; | [optional] 
**Segment4** | Pointer to **string** |  Max length: 255; | [optional] 
**Segment5** | Pointer to **string** |  Max length: 255; | [optional] 
**Segment6** | Pointer to **string** |  Max length: 255; | [optional] 
**Segment7** | Pointer to **string** |  Max length: 255; | [optional] 
**Segment8** | Pointer to **string** |  Max length: 255; | [optional] 
**Segment9** | Pointer to **string** |  Max length: 255; | [optional] 
**Segment10** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs1** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs2** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs3** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs4** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs5** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs6** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs7** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs8** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs9** | Pointer to **string** |  Max length: 255; | [optional] 
**Cogs10** | Pointer to **string** |  Max length: 255; | [optional] 
**ProductId** | Pointer to **string** |  Max length: 255; | [optional] 
**Inventory** | Pointer to **string** |  Max length: 255; | [optional] 
**SalesCode** | Pointer to **string** |  Max length: 255; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGLAccount

`func NewGLAccount(glType NullableString, mappedType MappedTypeReference, mappedRecord MappedRecordReference, ) *GLAccount`

NewGLAccount instantiates a new GLAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLAccountWithDefaults

`func NewGLAccountWithDefaults() *GLAccount`

NewGLAccountWithDefaults instantiates a new GLAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGlType

`func (o *GLAccount) GetGlType() string`

GetGlType returns the GlType field if non-nil, zero value otherwise.

### GetGlTypeOk

`func (o *GLAccount) GetGlTypeOk() (*string, bool)`

GetGlTypeOk returns a tuple with the GlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlType

`func (o *GLAccount) SetGlType(v string)`

SetGlType sets GlType field to given value.


### SetGlTypeNil

`func (o *GLAccount) SetGlTypeNil(b bool)`

 SetGlTypeNil sets the value for GlType to be an explicit nil

### UnsetGlType
`func (o *GLAccount) UnsetGlType()`

UnsetGlType ensures that no value is present for GlType, not even an explicit nil
### GetMappedType

`func (o *GLAccount) GetMappedType() MappedTypeReference`

GetMappedType returns the MappedType field if non-nil, zero value otherwise.

### GetMappedTypeOk

`func (o *GLAccount) GetMappedTypeOk() (*MappedTypeReference, bool)`

GetMappedTypeOk returns a tuple with the MappedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedType

`func (o *GLAccount) SetMappedType(v MappedTypeReference)`

SetMappedType sets MappedType field to given value.


### GetMappedRecord

`func (o *GLAccount) GetMappedRecord() MappedRecordReference`

GetMappedRecord returns the MappedRecord field if non-nil, zero value otherwise.

### GetMappedRecordOk

`func (o *GLAccount) GetMappedRecordOk() (*MappedRecordReference, bool)`

GetMappedRecordOk returns a tuple with the MappedRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedRecord

`func (o *GLAccount) SetMappedRecord(v MappedRecordReference)`

SetMappedRecord sets MappedRecord field to given value.


### GetSegment1

`func (o *GLAccount) GetSegment1() string`

GetSegment1 returns the Segment1 field if non-nil, zero value otherwise.

### GetSegment1Ok

`func (o *GLAccount) GetSegment1Ok() (*string, bool)`

GetSegment1Ok returns a tuple with the Segment1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1

`func (o *GLAccount) SetSegment1(v string)`

SetSegment1 sets Segment1 field to given value.

### HasSegment1

`func (o *GLAccount) HasSegment1() bool`

HasSegment1 returns a boolean if a field has been set.

### GetSegment2

`func (o *GLAccount) GetSegment2() string`

GetSegment2 returns the Segment2 field if non-nil, zero value otherwise.

### GetSegment2Ok

`func (o *GLAccount) GetSegment2Ok() (*string, bool)`

GetSegment2Ok returns a tuple with the Segment2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2

`func (o *GLAccount) SetSegment2(v string)`

SetSegment2 sets Segment2 field to given value.

### HasSegment2

`func (o *GLAccount) HasSegment2() bool`

HasSegment2 returns a boolean if a field has been set.

### GetSegment3

`func (o *GLAccount) GetSegment3() string`

GetSegment3 returns the Segment3 field if non-nil, zero value otherwise.

### GetSegment3Ok

`func (o *GLAccount) GetSegment3Ok() (*string, bool)`

GetSegment3Ok returns a tuple with the Segment3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3

`func (o *GLAccount) SetSegment3(v string)`

SetSegment3 sets Segment3 field to given value.

### HasSegment3

`func (o *GLAccount) HasSegment3() bool`

HasSegment3 returns a boolean if a field has been set.

### GetSegment4

`func (o *GLAccount) GetSegment4() string`

GetSegment4 returns the Segment4 field if non-nil, zero value otherwise.

### GetSegment4Ok

`func (o *GLAccount) GetSegment4Ok() (*string, bool)`

GetSegment4Ok returns a tuple with the Segment4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment4

`func (o *GLAccount) SetSegment4(v string)`

SetSegment4 sets Segment4 field to given value.

### HasSegment4

`func (o *GLAccount) HasSegment4() bool`

HasSegment4 returns a boolean if a field has been set.

### GetSegment5

`func (o *GLAccount) GetSegment5() string`

GetSegment5 returns the Segment5 field if non-nil, zero value otherwise.

### GetSegment5Ok

`func (o *GLAccount) GetSegment5Ok() (*string, bool)`

GetSegment5Ok returns a tuple with the Segment5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment5

`func (o *GLAccount) SetSegment5(v string)`

SetSegment5 sets Segment5 field to given value.

### HasSegment5

`func (o *GLAccount) HasSegment5() bool`

HasSegment5 returns a boolean if a field has been set.

### GetSegment6

`func (o *GLAccount) GetSegment6() string`

GetSegment6 returns the Segment6 field if non-nil, zero value otherwise.

### GetSegment6Ok

`func (o *GLAccount) GetSegment6Ok() (*string, bool)`

GetSegment6Ok returns a tuple with the Segment6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment6

`func (o *GLAccount) SetSegment6(v string)`

SetSegment6 sets Segment6 field to given value.

### HasSegment6

`func (o *GLAccount) HasSegment6() bool`

HasSegment6 returns a boolean if a field has been set.

### GetSegment7

`func (o *GLAccount) GetSegment7() string`

GetSegment7 returns the Segment7 field if non-nil, zero value otherwise.

### GetSegment7Ok

`func (o *GLAccount) GetSegment7Ok() (*string, bool)`

GetSegment7Ok returns a tuple with the Segment7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment7

`func (o *GLAccount) SetSegment7(v string)`

SetSegment7 sets Segment7 field to given value.

### HasSegment7

`func (o *GLAccount) HasSegment7() bool`

HasSegment7 returns a boolean if a field has been set.

### GetSegment8

`func (o *GLAccount) GetSegment8() string`

GetSegment8 returns the Segment8 field if non-nil, zero value otherwise.

### GetSegment8Ok

`func (o *GLAccount) GetSegment8Ok() (*string, bool)`

GetSegment8Ok returns a tuple with the Segment8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment8

`func (o *GLAccount) SetSegment8(v string)`

SetSegment8 sets Segment8 field to given value.

### HasSegment8

`func (o *GLAccount) HasSegment8() bool`

HasSegment8 returns a boolean if a field has been set.

### GetSegment9

`func (o *GLAccount) GetSegment9() string`

GetSegment9 returns the Segment9 field if non-nil, zero value otherwise.

### GetSegment9Ok

`func (o *GLAccount) GetSegment9Ok() (*string, bool)`

GetSegment9Ok returns a tuple with the Segment9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment9

`func (o *GLAccount) SetSegment9(v string)`

SetSegment9 sets Segment9 field to given value.

### HasSegment9

`func (o *GLAccount) HasSegment9() bool`

HasSegment9 returns a boolean if a field has been set.

### GetSegment10

`func (o *GLAccount) GetSegment10() string`

GetSegment10 returns the Segment10 field if non-nil, zero value otherwise.

### GetSegment10Ok

`func (o *GLAccount) GetSegment10Ok() (*string, bool)`

GetSegment10Ok returns a tuple with the Segment10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment10

`func (o *GLAccount) SetSegment10(v string)`

SetSegment10 sets Segment10 field to given value.

### HasSegment10

`func (o *GLAccount) HasSegment10() bool`

HasSegment10 returns a boolean if a field has been set.

### GetCogs1

`func (o *GLAccount) GetCogs1() string`

GetCogs1 returns the Cogs1 field if non-nil, zero value otherwise.

### GetCogs1Ok

`func (o *GLAccount) GetCogs1Ok() (*string, bool)`

GetCogs1Ok returns a tuple with the Cogs1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs1

`func (o *GLAccount) SetCogs1(v string)`

SetCogs1 sets Cogs1 field to given value.

### HasCogs1

`func (o *GLAccount) HasCogs1() bool`

HasCogs1 returns a boolean if a field has been set.

### GetCogs2

`func (o *GLAccount) GetCogs2() string`

GetCogs2 returns the Cogs2 field if non-nil, zero value otherwise.

### GetCogs2Ok

`func (o *GLAccount) GetCogs2Ok() (*string, bool)`

GetCogs2Ok returns a tuple with the Cogs2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs2

`func (o *GLAccount) SetCogs2(v string)`

SetCogs2 sets Cogs2 field to given value.

### HasCogs2

`func (o *GLAccount) HasCogs2() bool`

HasCogs2 returns a boolean if a field has been set.

### GetCogs3

`func (o *GLAccount) GetCogs3() string`

GetCogs3 returns the Cogs3 field if non-nil, zero value otherwise.

### GetCogs3Ok

`func (o *GLAccount) GetCogs3Ok() (*string, bool)`

GetCogs3Ok returns a tuple with the Cogs3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs3

`func (o *GLAccount) SetCogs3(v string)`

SetCogs3 sets Cogs3 field to given value.

### HasCogs3

`func (o *GLAccount) HasCogs3() bool`

HasCogs3 returns a boolean if a field has been set.

### GetCogs4

`func (o *GLAccount) GetCogs4() string`

GetCogs4 returns the Cogs4 field if non-nil, zero value otherwise.

### GetCogs4Ok

`func (o *GLAccount) GetCogs4Ok() (*string, bool)`

GetCogs4Ok returns a tuple with the Cogs4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs4

`func (o *GLAccount) SetCogs4(v string)`

SetCogs4 sets Cogs4 field to given value.

### HasCogs4

`func (o *GLAccount) HasCogs4() bool`

HasCogs4 returns a boolean if a field has been set.

### GetCogs5

`func (o *GLAccount) GetCogs5() string`

GetCogs5 returns the Cogs5 field if non-nil, zero value otherwise.

### GetCogs5Ok

`func (o *GLAccount) GetCogs5Ok() (*string, bool)`

GetCogs5Ok returns a tuple with the Cogs5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs5

`func (o *GLAccount) SetCogs5(v string)`

SetCogs5 sets Cogs5 field to given value.

### HasCogs5

`func (o *GLAccount) HasCogs5() bool`

HasCogs5 returns a boolean if a field has been set.

### GetCogs6

`func (o *GLAccount) GetCogs6() string`

GetCogs6 returns the Cogs6 field if non-nil, zero value otherwise.

### GetCogs6Ok

`func (o *GLAccount) GetCogs6Ok() (*string, bool)`

GetCogs6Ok returns a tuple with the Cogs6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs6

`func (o *GLAccount) SetCogs6(v string)`

SetCogs6 sets Cogs6 field to given value.

### HasCogs6

`func (o *GLAccount) HasCogs6() bool`

HasCogs6 returns a boolean if a field has been set.

### GetCogs7

`func (o *GLAccount) GetCogs7() string`

GetCogs7 returns the Cogs7 field if non-nil, zero value otherwise.

### GetCogs7Ok

`func (o *GLAccount) GetCogs7Ok() (*string, bool)`

GetCogs7Ok returns a tuple with the Cogs7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs7

`func (o *GLAccount) SetCogs7(v string)`

SetCogs7 sets Cogs7 field to given value.

### HasCogs7

`func (o *GLAccount) HasCogs7() bool`

HasCogs7 returns a boolean if a field has been set.

### GetCogs8

`func (o *GLAccount) GetCogs8() string`

GetCogs8 returns the Cogs8 field if non-nil, zero value otherwise.

### GetCogs8Ok

`func (o *GLAccount) GetCogs8Ok() (*string, bool)`

GetCogs8Ok returns a tuple with the Cogs8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs8

`func (o *GLAccount) SetCogs8(v string)`

SetCogs8 sets Cogs8 field to given value.

### HasCogs8

`func (o *GLAccount) HasCogs8() bool`

HasCogs8 returns a boolean if a field has been set.

### GetCogs9

`func (o *GLAccount) GetCogs9() string`

GetCogs9 returns the Cogs9 field if non-nil, zero value otherwise.

### GetCogs9Ok

`func (o *GLAccount) GetCogs9Ok() (*string, bool)`

GetCogs9Ok returns a tuple with the Cogs9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs9

`func (o *GLAccount) SetCogs9(v string)`

SetCogs9 sets Cogs9 field to given value.

### HasCogs9

`func (o *GLAccount) HasCogs9() bool`

HasCogs9 returns a boolean if a field has been set.

### GetCogs10

`func (o *GLAccount) GetCogs10() string`

GetCogs10 returns the Cogs10 field if non-nil, zero value otherwise.

### GetCogs10Ok

`func (o *GLAccount) GetCogs10Ok() (*string, bool)`

GetCogs10Ok returns a tuple with the Cogs10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs10

`func (o *GLAccount) SetCogs10(v string)`

SetCogs10 sets Cogs10 field to given value.

### HasCogs10

`func (o *GLAccount) HasCogs10() bool`

HasCogs10 returns a boolean if a field has been set.

### GetProductId

`func (o *GLAccount) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GLAccount) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GLAccount) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GLAccount) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetInventory

`func (o *GLAccount) GetInventory() string`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *GLAccount) GetInventoryOk() (*string, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *GLAccount) SetInventory(v string)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *GLAccount) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetSalesCode

`func (o *GLAccount) GetSalesCode() string`

GetSalesCode returns the SalesCode field if non-nil, zero value otherwise.

### GetSalesCodeOk

`func (o *GLAccount) GetSalesCodeOk() (*string, bool)`

GetSalesCodeOk returns a tuple with the SalesCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCode

`func (o *GLAccount) SetSalesCode(v string)`

SetSalesCode sets SalesCode field to given value.

### HasSalesCode

`func (o *GLAccount) HasSalesCode() bool`

HasSalesCode returns a boolean if a field has been set.

### GetInfo

`func (o *GLAccount) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GLAccount) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GLAccount) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GLAccount) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


