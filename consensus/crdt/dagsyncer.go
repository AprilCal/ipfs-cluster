package crdt

import (
	ipfslite "github.com/hsanjuan/ipfs-lite"
	cid "github.com/ipfs/go-cid"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
)

// A DAGSyncer component implementation using ipfs-lite.
type liteDAGSyncer struct {
	*ipfslite.Peer
	blockstore blockstore.Blockstore
}

func (lds *liteDAGSyncer) IsKnownBlock(c cid.Cid) (bool, error) {
	return lds.blockstore.Has(c)
}
