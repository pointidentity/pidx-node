package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	ssitypes "github.com/pointidentity/pidx-node/x/types"
)

// isAuthzExecMsg checks if the message is of authz.MsgExec type
func isAuthzExecMsg(msg sdk.Msg) bool {
	switch msg.(type) {
	case *authz.MsgExec:
		return true
	default:
		return false
	}
}

// isSSIMsg checks if the message is of SSI type
func isSSIMsg(msg sdk.Msg) bool {
	switch msg.(type) {
	case *ssitypes.MsgRegisterDID:
		return true
	case *ssitypes.MsgUpdateDID:
		return true
	case *ssitypes.MsgDeactivateDID:
		return true
	case *ssitypes.MsgRegisterCredentialSchema:
		return true
	case *ssitypes.MsgUpdateCredentialSchema:
		return true
	case *ssitypes.MsgRegisterCredentialStatus:
		return true
	case *ssitypes.MsgUpdateCredentialStatus:
		return true
	default:
		return false
	}
}
