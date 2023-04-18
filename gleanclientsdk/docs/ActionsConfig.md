# ActionsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to [**AnswerActionsConfig**](AnswerActionsConfig.md) |  | [optional] 
**DisableActions** | Pointer to **bool** |  | [optional] 
**DisableCopyLink** | Pointer to **bool** |  | [optional] 

## Methods

### NewActionsConfig

`func NewActionsConfig() *ActionsConfig`

NewActionsConfig instantiates a new ActionsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsConfigWithDefaults

`func NewActionsConfigWithDefaults() *ActionsConfig`

NewActionsConfigWithDefaults instantiates a new ActionsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *ActionsConfig) GetAnswer() AnswerActionsConfig`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *ActionsConfig) GetAnswerOk() (*AnswerActionsConfig, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *ActionsConfig) SetAnswer(v AnswerActionsConfig)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *ActionsConfig) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetDisableActions

`func (o *ActionsConfig) GetDisableActions() bool`

GetDisableActions returns the DisableActions field if non-nil, zero value otherwise.

### GetDisableActionsOk

`func (o *ActionsConfig) GetDisableActionsOk() (*bool, bool)`

GetDisableActionsOk returns a tuple with the DisableActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableActions

`func (o *ActionsConfig) SetDisableActions(v bool)`

SetDisableActions sets DisableActions field to given value.

### HasDisableActions

`func (o *ActionsConfig) HasDisableActions() bool`

HasDisableActions returns a boolean if a field has been set.

### GetDisableCopyLink

`func (o *ActionsConfig) GetDisableCopyLink() bool`

GetDisableCopyLink returns the DisableCopyLink field if non-nil, zero value otherwise.

### GetDisableCopyLinkOk

`func (o *ActionsConfig) GetDisableCopyLinkOk() (*bool, bool)`

GetDisableCopyLinkOk returns a tuple with the DisableCopyLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCopyLink

`func (o *ActionsConfig) SetDisableCopyLink(v bool)`

SetDisableCopyLink sets DisableCopyLink field to given value.

### HasDisableCopyLink

`func (o *ActionsConfig) HasDisableCopyLink() bool`

HasDisableCopyLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


