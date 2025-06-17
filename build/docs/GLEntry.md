# GLEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Description** | Pointer to **string** |  Max length: 100; | [optional] 
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
**ProductId** | Pointer to **string** |  Max length: 255; | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**SalesCode** | Pointer to **string** |  Max length: 255; | [optional] 
**Inventory** | Pointer to **string** |  Max length: 255; | [optional] 
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
**IsBatched** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGLEntry

`func NewGLEntry() *GLEntry`

NewGLEntry instantiates a new GLEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLEntryWithDefaults

`func NewGLEntryWithDefaults() *GLEntry`

NewGLEntryWithDefaults instantiates a new GLEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GLEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GLEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GLEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GLEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAmount

`func (o *GLEntry) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GLEntry) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GLEntry) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *GLEntry) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *GLEntry) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *GLEntry) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetDescription

`func (o *GLEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GLEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GLEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GLEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSegment1

`func (o *GLEntry) GetSegment1() string`

GetSegment1 returns the Segment1 field if non-nil, zero value otherwise.

### GetSegment1Ok

`func (o *GLEntry) GetSegment1Ok() (*string, bool)`

GetSegment1Ok returns a tuple with the Segment1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment1

`func (o *GLEntry) SetSegment1(v string)`

SetSegment1 sets Segment1 field to given value.

### HasSegment1

`func (o *GLEntry) HasSegment1() bool`

HasSegment1 returns a boolean if a field has been set.

### GetSegment2

`func (o *GLEntry) GetSegment2() string`

GetSegment2 returns the Segment2 field if non-nil, zero value otherwise.

### GetSegment2Ok

`func (o *GLEntry) GetSegment2Ok() (*string, bool)`

GetSegment2Ok returns a tuple with the Segment2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment2

`func (o *GLEntry) SetSegment2(v string)`

SetSegment2 sets Segment2 field to given value.

### HasSegment2

`func (o *GLEntry) HasSegment2() bool`

HasSegment2 returns a boolean if a field has been set.

### GetSegment3

`func (o *GLEntry) GetSegment3() string`

GetSegment3 returns the Segment3 field if non-nil, zero value otherwise.

### GetSegment3Ok

`func (o *GLEntry) GetSegment3Ok() (*string, bool)`

GetSegment3Ok returns a tuple with the Segment3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment3

`func (o *GLEntry) SetSegment3(v string)`

SetSegment3 sets Segment3 field to given value.

### HasSegment3

`func (o *GLEntry) HasSegment3() bool`

HasSegment3 returns a boolean if a field has been set.

### GetSegment4

`func (o *GLEntry) GetSegment4() string`

GetSegment4 returns the Segment4 field if non-nil, zero value otherwise.

### GetSegment4Ok

`func (o *GLEntry) GetSegment4Ok() (*string, bool)`

GetSegment4Ok returns a tuple with the Segment4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment4

`func (o *GLEntry) SetSegment4(v string)`

SetSegment4 sets Segment4 field to given value.

### HasSegment4

`func (o *GLEntry) HasSegment4() bool`

HasSegment4 returns a boolean if a field has been set.

### GetSegment5

`func (o *GLEntry) GetSegment5() string`

GetSegment5 returns the Segment5 field if non-nil, zero value otherwise.

### GetSegment5Ok

`func (o *GLEntry) GetSegment5Ok() (*string, bool)`

GetSegment5Ok returns a tuple with the Segment5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment5

`func (o *GLEntry) SetSegment5(v string)`

SetSegment5 sets Segment5 field to given value.

### HasSegment5

`func (o *GLEntry) HasSegment5() bool`

HasSegment5 returns a boolean if a field has been set.

### GetSegment6

`func (o *GLEntry) GetSegment6() string`

GetSegment6 returns the Segment6 field if non-nil, zero value otherwise.

### GetSegment6Ok

`func (o *GLEntry) GetSegment6Ok() (*string, bool)`

GetSegment6Ok returns a tuple with the Segment6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment6

`func (o *GLEntry) SetSegment6(v string)`

SetSegment6 sets Segment6 field to given value.

### HasSegment6

`func (o *GLEntry) HasSegment6() bool`

HasSegment6 returns a boolean if a field has been set.

### GetSegment7

`func (o *GLEntry) GetSegment7() string`

GetSegment7 returns the Segment7 field if non-nil, zero value otherwise.

### GetSegment7Ok

`func (o *GLEntry) GetSegment7Ok() (*string, bool)`

GetSegment7Ok returns a tuple with the Segment7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment7

`func (o *GLEntry) SetSegment7(v string)`

SetSegment7 sets Segment7 field to given value.

### HasSegment7

`func (o *GLEntry) HasSegment7() bool`

HasSegment7 returns a boolean if a field has been set.

### GetSegment8

`func (o *GLEntry) GetSegment8() string`

GetSegment8 returns the Segment8 field if non-nil, zero value otherwise.

### GetSegment8Ok

`func (o *GLEntry) GetSegment8Ok() (*string, bool)`

GetSegment8Ok returns a tuple with the Segment8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment8

`func (o *GLEntry) SetSegment8(v string)`

SetSegment8 sets Segment8 field to given value.

### HasSegment8

`func (o *GLEntry) HasSegment8() bool`

HasSegment8 returns a boolean if a field has been set.

### GetSegment9

`func (o *GLEntry) GetSegment9() string`

GetSegment9 returns the Segment9 field if non-nil, zero value otherwise.

### GetSegment9Ok

`func (o *GLEntry) GetSegment9Ok() (*string, bool)`

GetSegment9Ok returns a tuple with the Segment9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment9

`func (o *GLEntry) SetSegment9(v string)`

SetSegment9 sets Segment9 field to given value.

### HasSegment9

`func (o *GLEntry) HasSegment9() bool`

HasSegment9 returns a boolean if a field has been set.

### GetSegment10

`func (o *GLEntry) GetSegment10() string`

GetSegment10 returns the Segment10 field if non-nil, zero value otherwise.

### GetSegment10Ok

`func (o *GLEntry) GetSegment10Ok() (*string, bool)`

GetSegment10Ok returns a tuple with the Segment10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment10

`func (o *GLEntry) SetSegment10(v string)`

SetSegment10 sets Segment10 field to given value.

### HasSegment10

`func (o *GLEntry) HasSegment10() bool`

HasSegment10 returns a boolean if a field has been set.

### GetProductId

`func (o *GLEntry) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *GLEntry) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *GLEntry) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *GLEntry) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetCost

`func (o *GLEntry) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GLEntry) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GLEntry) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GLEntry) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *GLEntry) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *GLEntry) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetSalesCode

`func (o *GLEntry) GetSalesCode() string`

GetSalesCode returns the SalesCode field if non-nil, zero value otherwise.

### GetSalesCodeOk

`func (o *GLEntry) GetSalesCodeOk() (*string, bool)`

GetSalesCodeOk returns a tuple with the SalesCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesCode

`func (o *GLEntry) SetSalesCode(v string)`

SetSalesCode sets SalesCode field to given value.

### HasSalesCode

`func (o *GLEntry) HasSalesCode() bool`

HasSalesCode returns a boolean if a field has been set.

### GetInventory

`func (o *GLEntry) GetInventory() string`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *GLEntry) GetInventoryOk() (*string, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *GLEntry) SetInventory(v string)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *GLEntry) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetCogs1

`func (o *GLEntry) GetCogs1() string`

GetCogs1 returns the Cogs1 field if non-nil, zero value otherwise.

### GetCogs1Ok

`func (o *GLEntry) GetCogs1Ok() (*string, bool)`

GetCogs1Ok returns a tuple with the Cogs1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs1

`func (o *GLEntry) SetCogs1(v string)`

SetCogs1 sets Cogs1 field to given value.

### HasCogs1

`func (o *GLEntry) HasCogs1() bool`

HasCogs1 returns a boolean if a field has been set.

### GetCogs2

`func (o *GLEntry) GetCogs2() string`

GetCogs2 returns the Cogs2 field if non-nil, zero value otherwise.

### GetCogs2Ok

`func (o *GLEntry) GetCogs2Ok() (*string, bool)`

GetCogs2Ok returns a tuple with the Cogs2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs2

`func (o *GLEntry) SetCogs2(v string)`

SetCogs2 sets Cogs2 field to given value.

### HasCogs2

`func (o *GLEntry) HasCogs2() bool`

HasCogs2 returns a boolean if a field has been set.

### GetCogs3

`func (o *GLEntry) GetCogs3() string`

GetCogs3 returns the Cogs3 field if non-nil, zero value otherwise.

### GetCogs3Ok

`func (o *GLEntry) GetCogs3Ok() (*string, bool)`

GetCogs3Ok returns a tuple with the Cogs3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs3

`func (o *GLEntry) SetCogs3(v string)`

SetCogs3 sets Cogs3 field to given value.

### HasCogs3

`func (o *GLEntry) HasCogs3() bool`

HasCogs3 returns a boolean if a field has been set.

### GetCogs4

`func (o *GLEntry) GetCogs4() string`

GetCogs4 returns the Cogs4 field if non-nil, zero value otherwise.

### GetCogs4Ok

`func (o *GLEntry) GetCogs4Ok() (*string, bool)`

GetCogs4Ok returns a tuple with the Cogs4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs4

`func (o *GLEntry) SetCogs4(v string)`

SetCogs4 sets Cogs4 field to given value.

### HasCogs4

`func (o *GLEntry) HasCogs4() bool`

HasCogs4 returns a boolean if a field has been set.

### GetCogs5

`func (o *GLEntry) GetCogs5() string`

GetCogs5 returns the Cogs5 field if non-nil, zero value otherwise.

### GetCogs5Ok

`func (o *GLEntry) GetCogs5Ok() (*string, bool)`

GetCogs5Ok returns a tuple with the Cogs5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs5

`func (o *GLEntry) SetCogs5(v string)`

SetCogs5 sets Cogs5 field to given value.

### HasCogs5

`func (o *GLEntry) HasCogs5() bool`

HasCogs5 returns a boolean if a field has been set.

### GetCogs6

`func (o *GLEntry) GetCogs6() string`

GetCogs6 returns the Cogs6 field if non-nil, zero value otherwise.

### GetCogs6Ok

`func (o *GLEntry) GetCogs6Ok() (*string, bool)`

GetCogs6Ok returns a tuple with the Cogs6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs6

`func (o *GLEntry) SetCogs6(v string)`

SetCogs6 sets Cogs6 field to given value.

### HasCogs6

`func (o *GLEntry) HasCogs6() bool`

HasCogs6 returns a boolean if a field has been set.

### GetCogs7

`func (o *GLEntry) GetCogs7() string`

GetCogs7 returns the Cogs7 field if non-nil, zero value otherwise.

### GetCogs7Ok

`func (o *GLEntry) GetCogs7Ok() (*string, bool)`

GetCogs7Ok returns a tuple with the Cogs7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs7

`func (o *GLEntry) SetCogs7(v string)`

SetCogs7 sets Cogs7 field to given value.

### HasCogs7

`func (o *GLEntry) HasCogs7() bool`

HasCogs7 returns a boolean if a field has been set.

### GetCogs8

`func (o *GLEntry) GetCogs8() string`

GetCogs8 returns the Cogs8 field if non-nil, zero value otherwise.

### GetCogs8Ok

`func (o *GLEntry) GetCogs8Ok() (*string, bool)`

GetCogs8Ok returns a tuple with the Cogs8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs8

`func (o *GLEntry) SetCogs8(v string)`

SetCogs8 sets Cogs8 field to given value.

### HasCogs8

`func (o *GLEntry) HasCogs8() bool`

HasCogs8 returns a boolean if a field has been set.

### GetCogs9

`func (o *GLEntry) GetCogs9() string`

GetCogs9 returns the Cogs9 field if non-nil, zero value otherwise.

### GetCogs9Ok

`func (o *GLEntry) GetCogs9Ok() (*string, bool)`

GetCogs9Ok returns a tuple with the Cogs9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs9

`func (o *GLEntry) SetCogs9(v string)`

SetCogs9 sets Cogs9 field to given value.

### HasCogs9

`func (o *GLEntry) HasCogs9() bool`

HasCogs9 returns a boolean if a field has been set.

### GetCogs10

`func (o *GLEntry) GetCogs10() string`

GetCogs10 returns the Cogs10 field if non-nil, zero value otherwise.

### GetCogs10Ok

`func (o *GLEntry) GetCogs10Ok() (*string, bool)`

GetCogs10Ok returns a tuple with the Cogs10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCogs10

`func (o *GLEntry) SetCogs10(v string)`

SetCogs10 sets Cogs10 field to given value.

### HasCogs10

`func (o *GLEntry) HasCogs10() bool`

HasCogs10 returns a boolean if a field has been set.

### GetIsBatched

`func (o *GLEntry) GetIsBatched() bool`

GetIsBatched returns the IsBatched field if non-nil, zero value otherwise.

### GetIsBatchedOk

`func (o *GLEntry) GetIsBatchedOk() (*bool, bool)`

GetIsBatchedOk returns a tuple with the IsBatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBatched

`func (o *GLEntry) SetIsBatched(v bool)`

SetIsBatched sets IsBatched field to given value.

### HasIsBatched

`func (o *GLEntry) HasIsBatched() bool`

HasIsBatched returns a boolean if a field has been set.

### GetInfo

`func (o *GLEntry) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GLEntry) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GLEntry) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GLEntry) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


