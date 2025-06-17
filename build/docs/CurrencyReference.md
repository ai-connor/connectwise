# CurrencyReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**CurrencyCode** | Pointer to **string** |  | [optional] 
**DecimalSeparator** | Pointer to **string** |  | [optional] 
**NumberOfDecimals** | Pointer to **int32** |  | [optional] 
**ThousandsSeparator** | Pointer to **string** |  | [optional] 
**NegativeParenthesesFlag** | Pointer to **bool** |  | [optional] 
**DisplaySymbolFlag** | Pointer to **bool** |  | [optional] 
**CurrencyIdentifier** | Pointer to **string** |  | [optional] 
**DisplayIdFlag** | Pointer to **bool** |  | [optional] 
**RightAlign** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCurrencyReference

`func NewCurrencyReference() *CurrencyReference`

NewCurrencyReference instantiates a new CurrencyReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyReferenceWithDefaults

`func NewCurrencyReferenceWithDefaults() *CurrencyReference`

NewCurrencyReferenceWithDefaults instantiates a new CurrencyReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrencyReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrencyReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrencyReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CurrencyReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CurrencyReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CurrencyReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSymbol

`func (o *CurrencyReference) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CurrencyReference) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CurrencyReference) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CurrencyReference) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *CurrencyReference) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CurrencyReference) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CurrencyReference) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CurrencyReference) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetDecimalSeparator

`func (o *CurrencyReference) GetDecimalSeparator() string`

GetDecimalSeparator returns the DecimalSeparator field if non-nil, zero value otherwise.

### GetDecimalSeparatorOk

`func (o *CurrencyReference) GetDecimalSeparatorOk() (*string, bool)`

GetDecimalSeparatorOk returns a tuple with the DecimalSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalSeparator

`func (o *CurrencyReference) SetDecimalSeparator(v string)`

SetDecimalSeparator sets DecimalSeparator field to given value.

### HasDecimalSeparator

`func (o *CurrencyReference) HasDecimalSeparator() bool`

HasDecimalSeparator returns a boolean if a field has been set.

### GetNumberOfDecimals

`func (o *CurrencyReference) GetNumberOfDecimals() int32`

GetNumberOfDecimals returns the NumberOfDecimals field if non-nil, zero value otherwise.

### GetNumberOfDecimalsOk

`func (o *CurrencyReference) GetNumberOfDecimalsOk() (*int32, bool)`

GetNumberOfDecimalsOk returns a tuple with the NumberOfDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDecimals

`func (o *CurrencyReference) SetNumberOfDecimals(v int32)`

SetNumberOfDecimals sets NumberOfDecimals field to given value.

### HasNumberOfDecimals

`func (o *CurrencyReference) HasNumberOfDecimals() bool`

HasNumberOfDecimals returns a boolean if a field has been set.

### GetThousandsSeparator

`func (o *CurrencyReference) GetThousandsSeparator() string`

GetThousandsSeparator returns the ThousandsSeparator field if non-nil, zero value otherwise.

### GetThousandsSeparatorOk

`func (o *CurrencyReference) GetThousandsSeparatorOk() (*string, bool)`

GetThousandsSeparatorOk returns a tuple with the ThousandsSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThousandsSeparator

`func (o *CurrencyReference) SetThousandsSeparator(v string)`

SetThousandsSeparator sets ThousandsSeparator field to given value.

### HasThousandsSeparator

`func (o *CurrencyReference) HasThousandsSeparator() bool`

HasThousandsSeparator returns a boolean if a field has been set.

### GetNegativeParenthesesFlag

`func (o *CurrencyReference) GetNegativeParenthesesFlag() bool`

GetNegativeParenthesesFlag returns the NegativeParenthesesFlag field if non-nil, zero value otherwise.

### GetNegativeParenthesesFlagOk

`func (o *CurrencyReference) GetNegativeParenthesesFlagOk() (*bool, bool)`

GetNegativeParenthesesFlagOk returns a tuple with the NegativeParenthesesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeParenthesesFlag

`func (o *CurrencyReference) SetNegativeParenthesesFlag(v bool)`

SetNegativeParenthesesFlag sets NegativeParenthesesFlag field to given value.

### HasNegativeParenthesesFlag

`func (o *CurrencyReference) HasNegativeParenthesesFlag() bool`

HasNegativeParenthesesFlag returns a boolean if a field has been set.

### GetDisplaySymbolFlag

`func (o *CurrencyReference) GetDisplaySymbolFlag() bool`

GetDisplaySymbolFlag returns the DisplaySymbolFlag field if non-nil, zero value otherwise.

### GetDisplaySymbolFlagOk

`func (o *CurrencyReference) GetDisplaySymbolFlagOk() (*bool, bool)`

GetDisplaySymbolFlagOk returns a tuple with the DisplaySymbolFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySymbolFlag

`func (o *CurrencyReference) SetDisplaySymbolFlag(v bool)`

SetDisplaySymbolFlag sets DisplaySymbolFlag field to given value.

### HasDisplaySymbolFlag

`func (o *CurrencyReference) HasDisplaySymbolFlag() bool`

HasDisplaySymbolFlag returns a boolean if a field has been set.

### GetCurrencyIdentifier

`func (o *CurrencyReference) GetCurrencyIdentifier() string`

GetCurrencyIdentifier returns the CurrencyIdentifier field if non-nil, zero value otherwise.

### GetCurrencyIdentifierOk

`func (o *CurrencyReference) GetCurrencyIdentifierOk() (*string, bool)`

GetCurrencyIdentifierOk returns a tuple with the CurrencyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIdentifier

`func (o *CurrencyReference) SetCurrencyIdentifier(v string)`

SetCurrencyIdentifier sets CurrencyIdentifier field to given value.

### HasCurrencyIdentifier

`func (o *CurrencyReference) HasCurrencyIdentifier() bool`

HasCurrencyIdentifier returns a boolean if a field has been set.

### GetDisplayIdFlag

`func (o *CurrencyReference) GetDisplayIdFlag() bool`

GetDisplayIdFlag returns the DisplayIdFlag field if non-nil, zero value otherwise.

### GetDisplayIdFlagOk

`func (o *CurrencyReference) GetDisplayIdFlagOk() (*bool, bool)`

GetDisplayIdFlagOk returns a tuple with the DisplayIdFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIdFlag

`func (o *CurrencyReference) SetDisplayIdFlag(v bool)`

SetDisplayIdFlag sets DisplayIdFlag field to given value.

### HasDisplayIdFlag

`func (o *CurrencyReference) HasDisplayIdFlag() bool`

HasDisplayIdFlag returns a boolean if a field has been set.

### GetRightAlign

`func (o *CurrencyReference) GetRightAlign() bool`

GetRightAlign returns the RightAlign field if non-nil, zero value otherwise.

### GetRightAlignOk

`func (o *CurrencyReference) GetRightAlignOk() (*bool, bool)`

GetRightAlignOk returns a tuple with the RightAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightAlign

`func (o *CurrencyReference) SetRightAlign(v bool)`

SetRightAlign sets RightAlign field to given value.

### HasRightAlign

`func (o *CurrencyReference) HasRightAlign() bool`

HasRightAlign returns a boolean if a field has been set.

### GetName

`func (o *CurrencyReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrencyReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrencyReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrencyReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInfo

`func (o *CurrencyReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CurrencyReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CurrencyReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CurrencyReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


