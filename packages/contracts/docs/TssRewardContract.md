# TssRewardContract



> TssRewardContract



*Release to batch roll up tss members.*

## Methods

### bestBlockID

```solidity
function bestBlockID() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### bvmGasPriceOracleAddress

```solidity
function bvmGasPriceOracleAddress() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### claimReward

```solidity
function claimReward(uint256 _blockStartHeight, uint32 _length, uint256 _batchTime, address[] _tssMembers) external nonpayable
```



*claimReward distribute reward to tss member.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockStartHeight | uint256 | undefined
| _length | uint32 | The distribute batch block number
| _batchTime | uint256 | Batch corresponds to L1 Block Timestamp
| _tssMembers | address[] | The address array of tss group members

### deadAddress

```solidity
function deadAddress() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### dust

```solidity
function dust() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### lastBatchTime

```solidity
function lastBatchTime() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### ledger

```solidity
function ledger(uint256) external view returns (uint256)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### messenger

```solidity
function messenger() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### owner

```solidity
function owner() external view returns (address)
```



*Returns the address of the current owner.*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### queryReward

```solidity
function queryReward() external view returns (uint256)
```



*return the total undistributed amount*


#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### querySendAmountPerSecond

```solidity
function querySendAmountPerSecond() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### renounceOwnership

```solidity
function renounceOwnership() external nonpayable
```



*Leaves the contract without owner. It will not be possible to call `onlyOwner` functions anymore. Can only be called by the current owner. NOTE: Renouncing ownership will leave the contract without an owner, thereby removing any functionality that is only available to the owner.*


### sccAddress

```solidity
function sccAddress() external view returns (address)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | address | undefined

### sendAmountPerYear

```solidity
function sendAmountPerYear() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### setSendAmountPerYear

```solidity
function setSendAmountPerYear(uint256 _sendAmountPerYear) external nonpayable
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| _sendAmountPerYear | uint256 | undefined

### totalAmount

```solidity
function totalAmount() external view returns (uint256)
```






#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | uint256 | undefined

### transferOwnership

```solidity
function transferOwnership(address newOwner) external nonpayable
```



*Transfers ownership of the contract to a new account (`newOwner`). Can only be called by the current owner.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| newOwner | address | undefined

### updateReward

```solidity
function updateReward(uint256 _blockID, uint256 _amount) external nonpayable returns (bool)
```



*update tss member gas reward by every block.*

#### Parameters

| Name | Type | Description |
|---|---|---|
| _blockID | uint256 | The block height at L2 which needs to distribute profits
| _amount | uint256 | undefined

#### Returns

| Name | Type | Description |
|---|---|---|
| _0 | bool | _tssMembers Address array of tss group members

### withdraw

```solidity
function withdraw() external nonpayable
```



*clear balance*


### withdrawDust

```solidity
function withdrawDust() external nonpayable
```



*withdraw div dust*




## Events

### DistributeTssReward

```solidity
event DistributeTssReward(uint256 lastBatchTime, uint256 batchTime, uint256 amount, address[] tssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| lastBatchTime  | uint256 | undefined |
| batchTime  | uint256 | undefined |
| amount  | uint256 | undefined |
| tssMembers  | address[] | undefined |

### DistributeTssRewardByBlock

```solidity
event DistributeTssRewardByBlock(uint256 blockStartHeight, uint32 length, uint256 amount, address[] tssMembers)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| blockStartHeight  | uint256 | undefined |
| length  | uint32 | undefined |
| amount  | uint256 | undefined |
| tssMembers  | address[] | undefined |

### OwnershipTransferred

```solidity
event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
```





#### Parameters

| Name | Type | Description |
|---|---|---|
| previousOwner `indexed` | address | undefined |
| newOwner `indexed` | address | undefined |



