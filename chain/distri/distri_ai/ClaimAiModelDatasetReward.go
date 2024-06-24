// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package distri_ai

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ClaimAiModelDatasetReward is the `claimAiModelDatasetReward` instruction.
type ClaimAiModelDatasetReward struct {

	// [0] = [WRITE, SIGNER] owner
	//
	// [1] = [WRITE] ownerAta
	//
	// [2] = [WRITE] statisticsOwner
	//
	// [3] = [WRITE] rewardPool
	//
	// [4] = [] mint
	//
	// [5] = [] tokenProgram
	//
	// [6] = [] associatedTokenProgram
	//
	// [7] = [] systemProgram
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClaimAiModelDatasetRewardInstructionBuilder creates a new `ClaimAiModelDatasetReward` instruction builder.
func NewClaimAiModelDatasetRewardInstructionBuilder() *ClaimAiModelDatasetReward {
	nd := &ClaimAiModelDatasetReward{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetOwnerAccount sets the "owner" account.
func (inst *ClaimAiModelDatasetReward) SetOwnerAccount(owner ag_solanago.PublicKey) *ClaimAiModelDatasetReward {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(owner).WRITE().SIGNER()
	return inst
}

// GetOwnerAccount gets the "owner" account.
func (inst *ClaimAiModelDatasetReward) GetOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetOwnerAtaAccount sets the "ownerAta" account.
func (inst *ClaimAiModelDatasetReward) SetOwnerAtaAccount(ownerAta ag_solanago.PublicKey) *ClaimAiModelDatasetReward {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(ownerAta).WRITE()
	return inst
}

// GetOwnerAtaAccount gets the "ownerAta" account.
func (inst *ClaimAiModelDatasetReward) GetOwnerAtaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetStatisticsOwnerAccount sets the "statisticsOwner" account.
func (inst *ClaimAiModelDatasetReward) SetStatisticsOwnerAccount(statisticsOwner ag_solanago.PublicKey) *ClaimAiModelDatasetReward {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(statisticsOwner).WRITE()
	return inst
}

// GetStatisticsOwnerAccount gets the "statisticsOwner" account.
func (inst *ClaimAiModelDatasetReward) GetStatisticsOwnerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetRewardPoolAccount sets the "rewardPool" account.
func (inst *ClaimAiModelDatasetReward) SetRewardPoolAccount(rewardPool ag_solanago.PublicKey) *ClaimAiModelDatasetReward {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(rewardPool).WRITE()
	return inst
}

// GetRewardPoolAccount gets the "rewardPool" account.
func (inst *ClaimAiModelDatasetReward) GetRewardPoolAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMintAccount sets the "mint" account.
func (inst *ClaimAiModelDatasetReward) SetMintAccount(mint ag_solanago.PublicKey) *ClaimAiModelDatasetReward {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(mint)
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *ClaimAiModelDatasetReward) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *ClaimAiModelDatasetReward) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ClaimAiModelDatasetReward {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *ClaimAiModelDatasetReward) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *ClaimAiModelDatasetReward) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *ClaimAiModelDatasetReward {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *ClaimAiModelDatasetReward) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *ClaimAiModelDatasetReward) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ClaimAiModelDatasetReward {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *ClaimAiModelDatasetReward) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst ClaimAiModelDatasetReward) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClaimAiModelDatasetReward,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClaimAiModelDatasetReward) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClaimAiModelDatasetReward) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Owner is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.OwnerAta is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.StatisticsOwner is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.RewardPool is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
	}
	return nil
}

func (inst *ClaimAiModelDatasetReward) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClaimAiModelDatasetReward")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                 owner", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("              ownerAta", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("       statisticsOwner", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("            rewardPool", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj ClaimAiModelDatasetReward) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ClaimAiModelDatasetReward) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewClaimAiModelDatasetRewardInstruction declares a new ClaimAiModelDatasetReward instruction with the provided parameters and accounts.
func NewClaimAiModelDatasetRewardInstruction(
	// Accounts:
	owner ag_solanago.PublicKey,
	ownerAta ag_solanago.PublicKey,
	statisticsOwner ag_solanago.PublicKey,
	rewardPool ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey) *ClaimAiModelDatasetReward {
	return NewClaimAiModelDatasetRewardInstructionBuilder().
		SetOwnerAccount(owner).
		SetOwnerAtaAccount(ownerAta).
		SetStatisticsOwnerAccount(statisticsOwner).
		SetRewardPoolAccount(rewardPool).
		SetMintAccount(mint).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetSystemProgramAccount(systemProgram)
}
