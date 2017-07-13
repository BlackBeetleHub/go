package build

import (
	"github.com/BlackBeetleHub/go/support/errors"
	"github.com/BlackBeetleHub/go/xdr"
)

// CreateAccount groups the creation of a new CreateAccountBuilder with a call
// to Mutate.
func CreateAlias(muts ...interface{}) (result CreateAliasBuilder) {
	result.Mutate(muts...)
	return
}

// CreateAccountMutator is a interface that wraps the
// MutateCreateAccount operation.  types may implement this interface to
// specify how they modify an xdr.PaymentOp object
type CreateAliasMutator interface {
	MutateCreateAlias(*xdr.CreateAliasOp) error
}

// CreateAccountBuilder helps to build CreateAccountOp structs.
type CreateAliasBuilder struct {
	O   xdr.Operation
	CAL  xdr.CreateAliasOp
	Err error
}

// Mutate applies the provided mutators to this builder's payment or operation.
func (b *CreateAliasBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error
		switch mut := m.(type) {
		case CreateAliasMutator:
			err = mut.MutateCreateAlias(&b.CAL)
		case OperationMutator:
			err = mut.MutateOperation(&b.O)
		default:
			err = errors.New("Mutator type not allowed")
		}

		if err != nil {
			b.Err = err
			return
		}
	}
}

// MutateCreateAccount for Destination sets the CreateAccountOp's Destination
// field
func (m AliasID) MutateCreateAlias(o *xdr.CreateAliasOp) error {
	return setAccountId(m.Address, &o.AccountId)
}

// MutateCreateAccount for NativeAmount sets the CreateAccountOp's
// StartingBalance field
func (m Source) MutateCreateAlias(o *xdr.CreateAliasOp) (err error) {
	return setAccountId(m.Address, &o.SourceId)
}
