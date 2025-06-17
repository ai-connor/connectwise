# FinanceCurrency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CurrencyIdentifier** | **string** |  Max length: 10; | 
**Name** | **string** |  Max length: 50; | 
**Symbol** | Pointer to **string** |  Max length: 10; | [optional] 
**DisplayIdFlag** | Pointer to **NullableBool** |  | [optional] 
**DisplaySymbolFlag** | Pointer to **NullableBool** |  | [optional] 
**CurrencyCode** | Pointer to [**CurrencyCodeReference**](CurrencyCodeReference.md) |  | [optional] 
**ThousandsSeparator** | Pointer to **string** |  Max length: 1; | [optional] 
**DecimalSeparator** | Pointer to **string** |  Max length: 1; | [optional] 
**NegativeParenthesesFlag** | Pointer to **NullableBool** |  | [optional] 
**RightAlign** | Pointer to **NullableBool** |  | [optional] 
**NumberOfDecimals** | Pointer to **NullableInt32** |  | [optional] 
**ReportFormat** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewFinanceCurrency

`func NewFinanceCurrency(currencyIdentifier string, name string, ) *FinanceCurrency`

NewFinanceCurrency instantiates a new FinanceCurrency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinanceCurrencyWithDefaults

`func NewFinanceCurrencyWithDefaults() *FinanceCurrency`

NewFinanceCurrencyWithDefaults instantiates a new FinanceCurrency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FinanceCurrency) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FinanceCurrency) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FinanceCurrency) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FinanceCurrency) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCurrencyIdentifier

`func (o *FinanceCurrency) GetCurrencyIdentifier() string`

GetCurrencyIdentifier returns the CurrencyIdentifier field if non-nil, zero value otherwise.

### GetCurrencyIdentifierOk

`func (o *FinanceCurrency) GetCurrencyIdentifierOk() (*string, bool)`

GetCurrencyIdentifierOk returns a tuple with the CurrencyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyIdentifier

`func (o *FinanceCurrency) SetCurrencyIdentifier(v string)`

SetCurrencyIdentifier sets CurrencyIdentifier field to given value.


### GetName

`func (o *FinanceCurrency) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FinanceCurrency) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FinanceCurrency) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *FinanceCurrency) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FinanceCurrency) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FinanceCurrency) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FinanceCurrency) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDisplayIdFlag

`func (o *FinanceCurrency) GetDisplayIdFlag() bool`

GetDisplayIdFlag returns the DisplayIdFlag field if non-nil, zero value otherwise.

### GetDisplayIdFlagOk

`func (o *FinanceCurrency) GetDisplayIdFlagOk() (*bool, bool)`

GetDisplayIdFlagOk returns a tuple with the DisplayIdFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIdFlag

`func (o *FinanceCurrency) SetDisplayIdFlag(v bool)`

SetDisplayIdFlag sets DisplayIdFlag field to given value.

### HasDisplayIdFlag

`func (o *FinanceCurrency) HasDisplayIdFlag() bool`

HasDisplayIdFlag returns a boolean if a field has been set.

### SetDisplayIdFlagNil

`func (o *FinanceCurrency) SetDisplayIdFlagNil(b bool)`

 SetDisplayIdFlagNil sets the value for DisplayIdFlag to be an explicit nil

### UnsetDisplayIdFlag
`func (o *FinanceCurrency) UnsetDisplayIdFlag()`

UnsetDisplayIdFlag ensures that no value is present for DisplayIdFlag, not even an explicit nil
### GetDisplaySymbolFlag

`func (o *FinanceCurrency) GetDisplaySymbolFlag() bool`

GetDisplaySymbolFlag returns the DisplaySymbolFlag field if non-nil, zero value otherwise.

### GetDisplaySymbolFlagOk

`func (o *FinanceCurrency) GetDisplaySymbolFlagOk() (*bool, bool)`

GetDisplaySymbolFlagOk returns a tuple with the DisplaySymbolFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySymbolFlag

`func (o *FinanceCurrency) SetDisplaySymbolFlag(v bool)`

SetDisplaySymbolFlag sets DisplaySymbolFlag field to given value.

### HasDisplaySymbolFlag

`func (o *FinanceCurrency) HasDisplaySymbolFlag() bool`

HasDisplaySymbolFlag returns a boolean if a field has been set.

### SetDisplaySymbolFlagNil

`func (o *FinanceCurrency) SetDisplaySymbolFlagNil(b bool)`

 SetDisplaySymbolFlagNil sets the value for DisplaySymbolFlag to be an explicit nil

### UnsetDisplaySymbolFlag
`func (o *FinanceCurrency) UnsetDisplaySymbolFlag()`

UnsetDisplaySymbolFlag ensures that no value is present for DisplaySymbolFlag, not even an explicit nil
### GetCurrencyCode

`func (o *FinanceCurrency) GetCurrencyCode() CurrencyCodeReference`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *FinanceCurrency) GetCurrencyCodeOk() (*CurrencyCodeReference, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *FinanceCurrency) SetCurrencyCode(v CurrencyCodeReference)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *FinanceCurrency) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetThousandsSeparator

`func (o *FinanceCurrency) GetThousandsSeparator() string`

GetThousandsSeparator returns the ThousandsSeparator field if non-nil, zero value otherwise.

### GetThousandsSeparatorOk

`func (o *FinanceCurrency) GetThousandsSeparatorOk() (*string, bool)`

GetThousandsSeparatorOk returns a tuple with the ThousandsSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThousandsSeparator

`func (o *FinanceCurrency) SetThousandsSeparator(v string)`

SetThousandsSeparator sets ThousandsSeparator field to given value.

### HasThousandsSeparator

`func (o *FinanceCurrency) HasThousandsSeparator() bool`

HasThousandsSeparator returns a boolean if a field has been set.

### GetDecimalSeparator

`func (o *FinanceCurrency) GetDecimalSeparator() string`

GetDecimalSeparator returns the DecimalSeparator field if non-nil, zero value otherwise.

### GetDecimalSeparatorOk

`func (o *FinanceCurrency) GetDecimalSeparatorOk() (*string, bool)`

GetDecimalSeparatorOk returns a tuple with the DecimalSeparator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimalSeparator

`func (o *FinanceCurrency) SetDecimalSeparator(v string)`

SetDecimalSeparator sets DecimalSeparator field to given value.

### HasDecimalSeparator

`func (o *FinanceCurrency) HasDecimalSeparator() bool`

HasDecimalSeparator returns a boolean if a field has been set.

### GetNegativeParenthesesFlag

`func (o *FinanceCurrency) GetNegativeParenthesesFlag() bool`

GetNegativeParenthesesFlag returns the NegativeParenthesesFlag field if non-nil, zero value otherwise.

### GetNegativeParenthesesFlagOk

`func (o *FinanceCurrency) GetNegativeParenthesesFlagOk() (*bool, bool)`

GetNegativeParenthesesFlagOk returns a tuple with the NegativeParenthesesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeParenthesesFlag

`func (o *FinanceCurrency) SetNegativeParenthesesFlag(v bool)`

SetNegativeParenthesesFlag sets NegativeParenthesesFlag field to given value.

### HasNegativeParenthesesFlag

`func (o *FinanceCurrency) HasNegativeParenthesesFlag() bool`

HasNegativeParenthesesFlag returns a boolean if a field has been set.

### SetNegativeParenthesesFlagNil

`func (o *FinanceCurrency) SetNegativeParenthesesFlagNil(b bool)`

 SetNegativeParenthesesFlagNil sets the value for NegativeParenthesesFlag to be an explicit nil

### UnsetNegativeParenthesesFlag
`func (o *FinanceCurrency) UnsetNegativeParenthesesFlag()`

UnsetNegativeParenthesesFlag ensures that no value is present for NegativeParenthesesFlag, not even an explicit nil
### GetRightAlign

`func (o *FinanceCurrency) GetRightAlign() bool`

GetRightAlign returns the RightAlign field if non-nil, zero value otherwise.

### GetRightAlignOk

`func (o *FinanceCurrency) GetRightAlignOk() (*bool, bool)`

GetRightAlignOk returns a tuple with the RightAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightAlign

`func (o *FinanceCurrency) SetRightAlign(v bool)`

SetRightAlign sets RightAlign field to given value.

### HasRightAlign

`func (o *FinanceCurrency) HasRightAlign() bool`

HasRightAlign returns a boolean if a field has been set.

### SetRightAlignNil

`func (o *FinanceCurrency) SetRightAlignNil(b bool)`

 SetRightAlignNil sets the value for RightAlign to be an explicit nil

### UnsetRightAlign
`func (o *FinanceCurrency) UnsetRightAlign()`

UnsetRightAlign ensures that no value is present for RightAlign, not even an explicit nil
### GetNumberOfDecimals

`func (o *FinanceCurrency) GetNumberOfDecimals() int32`

GetNumberOfDecimals returns the NumberOfDecimals field if non-nil, zero value otherwise.

### GetNumberOfDecimalsOk

`func (o *FinanceCurrency) GetNumberOfDecimalsOk() (*int32, bool)`

GetNumberOfDecimalsOk returns a tuple with the NumberOfDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDecimals

`func (o *FinanceCurrency) SetNumberOfDecimals(v int32)`

SetNumberOfDecimals sets NumberOfDecimals field to given value.

### HasNumberOfDecimals

`func (o *FinanceCurrency) HasNumberOfDecimals() bool`

HasNumberOfDecimals returns a boolean if a field has been set.

### SetNumberOfDecimalsNil

`func (o *FinanceCurrency) SetNumberOfDecimalsNil(b bool)`

 SetNumberOfDecimalsNil sets the value for NumberOfDecimals to be an explicit nil

### UnsetNumberOfDecimals
`func (o *FinanceCurrency) UnsetNumberOfDecimals()`

UnsetNumberOfDecimals ensures that no value is present for NumberOfDecimals, not even an explicit nil
### GetReportFormat

`func (o *FinanceCurrency) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *FinanceCurrency) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *FinanceCurrency) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### HasReportFormat

`func (o *FinanceCurrency) HasReportFormat() bool`

HasReportFormat returns a boolean if a field has been set.

### GetInfo

`func (o *FinanceCurrency) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *FinanceCurrency) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *FinanceCurrency) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *FinanceCurrency) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


