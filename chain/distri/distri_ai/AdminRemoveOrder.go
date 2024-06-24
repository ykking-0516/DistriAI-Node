// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package distri_ai

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// AdminRemoveOrder is the `adminRemoveOrder` instruction.
type AdminRemoveOrder struct {

	// [0] = [WRITE] order
	//
	// [1] = [WRITE, SIGNER] admin
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAdminRemoveOrderInstructionBuilder creates a new `AdminRemoveOrder` instruction builder.
func NewAdminRemoveOrderInstructionBuilder() *AdminRemoveOrder {
	nd := &AdminRemoveOrder{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 2),
	}
	return nd
}

// SetOrderAccount sets the "order" account.
func (inst *AdminRemoveOrder) SetOrderAccount(order ag_solanago.PublicKey) *AdminRemoveOrder {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(order).WRITE()
	return inst
}

// GetOrderAccount gets the "order" account.
func (inst *AdminRemoveOrder) GetOrderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *AdminRemoveOrder) SetAdminAccount(admin ag_solanago.PublicKey) *AdminRemoveOrder {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *AdminRemoveOrder) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

func (inst AdminRemoveOrder) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AdminRemoveOrder,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AdminRemoveOrder) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AdminRemoveOrder) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Order is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
	}
	return nil
}

func (inst *AdminRemoveOrder) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AdminRemoveOrder")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("order", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("admin", inst.AccountMetaSlice.Get(1)))
					})
				})
		})
}

func (obj AdminRemoveOrder) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *AdminRemoveOrder) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewAdminRemoveOrderInstruction declares a new AdminRemoveOrder instruction with the provided parameters and accounts.
func NewAdminRemoveOrderInstruction(
	// Accounts:
	order ag_solanago.PublicKey,
	admin ag_solanago.PublicKey) *AdminRemoveOrder {
	return NewAdminRemoveOrderInstructionBuilder().
		SetOrderAccount(order).
		SetAdminAccount(admin)
}
