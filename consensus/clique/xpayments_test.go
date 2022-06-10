package clique

import (
	"math/big"
	"testing"

	"github.com/xpaymentsorg/go-xpayments/params"
)

// BlockReward is the reward in wei distributed each block.
var BlockRewardTest = big.NewInt(1e+18)
var signerRewardTest = new(big.Int).Rsh(BlockRewardTest, 1)
var stakeRewardTest = new(big.Int).Sub(BlockRewardTest, signerRewardTest)

func TestBlockReward(t *testing.T) {

	t.Errorf("Block reward: %s ", BlockRewardTest)
	t.Errorf("Signer reward: %s ", signerRewardTest)
	t.Errorf("Stake address reward: %s ", stakeRewardTest)
	t.Errorf("Int: %s ", new(big.Int).Mul(BlockRewardTest, big.NewInt(params.Xps)))

}
