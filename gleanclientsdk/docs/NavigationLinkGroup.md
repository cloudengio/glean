# NavigationLinkGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**DefaultLink** | Pointer to [**NavigationLink**](NavigationLink.md) |  | [optional] 
**Links** | [**[]NavigationLink**](NavigationLink.md) | Links within the group | 
**Behaviors** | Pointer to [**[]NavigationLinkGroupBehavior**](NavigationLinkGroupBehavior.md) | Behavior hints such as whether the group should be separated.  May or may not be used by the client during rendering. | [optional] 

## Methods

### NewNavigationLinkGroup

`func NewNavigationLinkGroup(links []NavigationLink, ) *NavigationLinkGroup`

NewNavigationLinkGroup instantiates a new NavigationLinkGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigationLinkGroupWithDefaults

`func NewNavigationLinkGroupWithDefaults() *NavigationLinkGroup`

NewNavigationLinkGroupWithDefaults instantiates a new NavigationLinkGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NavigationLinkGroup) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NavigationLinkGroup) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NavigationLinkGroup) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NavigationLinkGroup) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDefaultLink

`func (o *NavigationLinkGroup) GetDefaultLink() NavigationLink`

GetDefaultLink returns the DefaultLink field if non-nil, zero value otherwise.

### GetDefaultLinkOk

`func (o *NavigationLinkGroup) GetDefaultLinkOk() (*NavigationLink, bool)`

GetDefaultLinkOk returns a tuple with the DefaultLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLink

`func (o *NavigationLinkGroup) SetDefaultLink(v NavigationLink)`

SetDefaultLink sets DefaultLink field to given value.

### HasDefaultLink

`func (o *NavigationLinkGroup) HasDefaultLink() bool`

HasDefaultLink returns a boolean if a field has been set.

### GetLinks

`func (o *NavigationLinkGroup) GetLinks() []NavigationLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NavigationLinkGroup) GetLinksOk() (*[]NavigationLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NavigationLinkGroup) SetLinks(v []NavigationLink)`

SetLinks sets Links field to given value.


### GetBehaviors

`func (o *NavigationLinkGroup) GetBehaviors() []NavigationLinkGroupBehavior`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *NavigationLinkGroup) GetBehaviorsOk() (*[]NavigationLinkGroupBehavior, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *NavigationLinkGroup) SetBehaviors(v []NavigationLinkGroupBehavior)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *NavigationLinkGroup) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


