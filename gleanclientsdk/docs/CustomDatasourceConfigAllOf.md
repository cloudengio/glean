# CustomDatasourceConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityDatasourceName** | Pointer to **string** | If the datasource uses another datasource for identity info, then the name of the datasource. The identity datasource must exist already. | [optional] 
**ProductAccessGroup** | Pointer to **string** | If the datasource uses a specific product access group, then the name of that group. | [optional] 
**IsUserReferencedByEmail** | Pointer to **bool** | whether email is used to reference users in document ACLs and in group memberships. | [optional] 
**IsEntityDatasource** | Pointer to **bool** | True if this datasource is used to push custom entities. | [optional] [default to false]
**IsTestDatasource** | Pointer to **bool** | True if this datasource will be used for testing purpose only. Documents from such a datasource wouldn&#39;t have any effect on search rankings. | [optional] [default to false]

## Methods

### NewCustomDatasourceConfigAllOf

`func NewCustomDatasourceConfigAllOf() *CustomDatasourceConfigAllOf`

NewCustomDatasourceConfigAllOf instantiates a new CustomDatasourceConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDatasourceConfigAllOfWithDefaults

`func NewCustomDatasourceConfigAllOfWithDefaults() *CustomDatasourceConfigAllOf`

NewCustomDatasourceConfigAllOfWithDefaults instantiates a new CustomDatasourceConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityDatasourceName

`func (o *CustomDatasourceConfigAllOf) GetIdentityDatasourceName() string`

GetIdentityDatasourceName returns the IdentityDatasourceName field if non-nil, zero value otherwise.

### GetIdentityDatasourceNameOk

`func (o *CustomDatasourceConfigAllOf) GetIdentityDatasourceNameOk() (*string, bool)`

GetIdentityDatasourceNameOk returns a tuple with the IdentityDatasourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityDatasourceName

`func (o *CustomDatasourceConfigAllOf) SetIdentityDatasourceName(v string)`

SetIdentityDatasourceName sets IdentityDatasourceName field to given value.

### HasIdentityDatasourceName

`func (o *CustomDatasourceConfigAllOf) HasIdentityDatasourceName() bool`

HasIdentityDatasourceName returns a boolean if a field has been set.

### GetProductAccessGroup

`func (o *CustomDatasourceConfigAllOf) GetProductAccessGroup() string`

GetProductAccessGroup returns the ProductAccessGroup field if non-nil, zero value otherwise.

### GetProductAccessGroupOk

`func (o *CustomDatasourceConfigAllOf) GetProductAccessGroupOk() (*string, bool)`

GetProductAccessGroupOk returns a tuple with the ProductAccessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAccessGroup

`func (o *CustomDatasourceConfigAllOf) SetProductAccessGroup(v string)`

SetProductAccessGroup sets ProductAccessGroup field to given value.

### HasProductAccessGroup

`func (o *CustomDatasourceConfigAllOf) HasProductAccessGroup() bool`

HasProductAccessGroup returns a boolean if a field has been set.

### GetIsUserReferencedByEmail

`func (o *CustomDatasourceConfigAllOf) GetIsUserReferencedByEmail() bool`

GetIsUserReferencedByEmail returns the IsUserReferencedByEmail field if non-nil, zero value otherwise.

### GetIsUserReferencedByEmailOk

`func (o *CustomDatasourceConfigAllOf) GetIsUserReferencedByEmailOk() (*bool, bool)`

GetIsUserReferencedByEmailOk returns a tuple with the IsUserReferencedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserReferencedByEmail

`func (o *CustomDatasourceConfigAllOf) SetIsUserReferencedByEmail(v bool)`

SetIsUserReferencedByEmail sets IsUserReferencedByEmail field to given value.

### HasIsUserReferencedByEmail

`func (o *CustomDatasourceConfigAllOf) HasIsUserReferencedByEmail() bool`

HasIsUserReferencedByEmail returns a boolean if a field has been set.

### GetIsEntityDatasource

`func (o *CustomDatasourceConfigAllOf) GetIsEntityDatasource() bool`

GetIsEntityDatasource returns the IsEntityDatasource field if non-nil, zero value otherwise.

### GetIsEntityDatasourceOk

`func (o *CustomDatasourceConfigAllOf) GetIsEntityDatasourceOk() (*bool, bool)`

GetIsEntityDatasourceOk returns a tuple with the IsEntityDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEntityDatasource

`func (o *CustomDatasourceConfigAllOf) SetIsEntityDatasource(v bool)`

SetIsEntityDatasource sets IsEntityDatasource field to given value.

### HasIsEntityDatasource

`func (o *CustomDatasourceConfigAllOf) HasIsEntityDatasource() bool`

HasIsEntityDatasource returns a boolean if a field has been set.

### GetIsTestDatasource

`func (o *CustomDatasourceConfigAllOf) GetIsTestDatasource() bool`

GetIsTestDatasource returns the IsTestDatasource field if non-nil, zero value otherwise.

### GetIsTestDatasourceOk

`func (o *CustomDatasourceConfigAllOf) GetIsTestDatasourceOk() (*bool, bool)`

GetIsTestDatasourceOk returns a tuple with the IsTestDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestDatasource

`func (o *CustomDatasourceConfigAllOf) SetIsTestDatasource(v bool)`

SetIsTestDatasource sets IsTestDatasource field to given value.

### HasIsTestDatasource

`func (o *CustomDatasourceConfigAllOf) HasIsTestDatasource() bool`

HasIsTestDatasource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


