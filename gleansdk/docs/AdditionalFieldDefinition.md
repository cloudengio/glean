# AdditionalFieldDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key to reference this field, e.g. \&quot;languages\&quot;. | [optional] 
**Value** | Pointer to **[]map[string]interface{}** | List of type string or HypertextField.  HypertextField is defined as &#x60;&#x60;&#x60; {   anchor: string,    // Anchor text for the hypertext field.   hyperlink: string, // URL for the hypertext field. } &#x60;&#x60;&#x60; Example: &#x60;&#x60;&#x60;{\&quot;anchor\&quot;:\&quot;Glean\&quot;,\&quot;hyperlink\&quot;:\&quot;https://glean.com\&quot;}&#x60;&#x60;&#x60;  When OpenAPI Generator supports oneOf, we will semantically enforce this in the docs.  | [optional] 

## Methods

### NewAdditionalFieldDefinition

`func NewAdditionalFieldDefinition() *AdditionalFieldDefinition`

NewAdditionalFieldDefinition instantiates a new AdditionalFieldDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionalFieldDefinitionWithDefaults

`func NewAdditionalFieldDefinitionWithDefaults() *AdditionalFieldDefinition`

NewAdditionalFieldDefinitionWithDefaults instantiates a new AdditionalFieldDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *AdditionalFieldDefinition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AdditionalFieldDefinition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AdditionalFieldDefinition) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AdditionalFieldDefinition) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *AdditionalFieldDefinition) GetValue() []map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AdditionalFieldDefinition) GetValueOk() (*[]map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AdditionalFieldDefinition) SetValue(v []map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *AdditionalFieldDefinition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


