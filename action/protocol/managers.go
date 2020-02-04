package protocol

import (
	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-core/db"
	"github.com/iotexproject/iotex-core/db/batch"
)

type (
	// StateConfig is the config for accessing stateDB
	StateConfig struct {
		Namespace string // namespace used by state's storage
	}

	// StateOption sets parameter for access state
	StateOption func(*StateConfig) error

	// StateReader defines an interface to read stateDB
	StateReader interface {
		Height() (uint64, error)
		State(hash.Hash160, interface{}, ...StateOption) error
		StateAtHeight(uint64, hash.Hash160, interface{}) error
	}

	// StateManager defines the stateDB interface atop IoTeX blockchain
	StateManager interface {
		StateReader
		// Accounts
		Snapshot() int
		Revert(int) error
		// General state
		PutState(hash.Hash160, interface{}, ...StateOption) error
		DelState(hash.Hash160, ...StateOption) error
		GetDB() db.KVStore
		GetCachedBatch() batch.CachedBatch
	}
)
