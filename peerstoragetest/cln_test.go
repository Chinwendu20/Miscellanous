package peerstoragetest

import (
	"github.com/elementsproject/peerswap/testframework"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
	"time"
)

type fundingNode string

const (
	FUNDER_LND fundingNode = "lnd"
	FUNDER_CLN fundingNode = "cln"
)

type setUpCtx struct {
	bitcoinNode *testframework.BitcoinNode
	clnNode     *testframework.CLightningNode
	lnNode      *testframework.LndNode
	scid        string
}

func setUp(t *testing.T, fundAmt uint64, funder fundingNode) *setUpCtx {

	//  Intentionally left the directory uncleaned so that it can be
	//  examined after the test.

	testDir, err := os.MkdirTemp("", "pop")
	if err != nil {
		t.Fatalf("could not create temp dir %v", err)
	}
	testDir = testDir + "/001"
	// fmt.Println(testDir)

	// Setup nodes (1 bitcoind, 1 cln, 1 lnd)
	bitcoind, err := testframework.NewBitcoinNode(testDir, 1)
	if err != nil {
		t.Fatalf("could not create bitcoind %v", err)
	}
	t.Cleanup(bitcoind.Kill)

	// cln
	cln, err := testframework.NewCLightningNode(testDir, bitcoind, 1)
	if err != nil {
		t.Fatalf("could not create cln %v", err)
	}
	t.Cleanup(cln.Kill)

	// Use lightningd with --developer turned on
	cln.WithCmd("lightningd")

	// Add to cmd line options
	cln.AppendCmdLine([]string{
		"--experimental-peer-storage",
	})

	// lnd
	lnd, err := testframework.NewLndNode(testDir, bitcoind, 1, nil)
	if err != nil {
		t.Fatalf("could not create lnd %v", err)
	}

	lnd.AppendCmdLine([]string{
		"--protocol.peer-storage",
	})

	t.Cleanup(lnd.Kill)

	// Write the logs for core lightning when the test ends.
	t.Cleanup(func() {

		r := strings.NewReader(cln.StdOut.String())
		b := make([]byte, r.Len())
		_, err := r.Read(b)
		require.NoError(t, err)
		filePath := cln.DataDir + "/pop.log"

		// Open the file for writing; create it if it does not exist, or truncate it if it does
		file, err := os.Create(filePath)
		require.NoError(t, err, "Failed to open file for writing stdout logs")

		// Ensure the file is closed after writing
		defer func() {
			require.NoError(t, file.Close())
		}()

		// Write the buffer to the file
		_, err = file.Write(b)
		require.NoError(t, err, "Failed to write stdout logs to file")

	})

	// Start nodes
	err = bitcoind.Run(true)
	if err != nil {
		t.Fatalf("bitcoind.Run() got err %v", err)
	}

	err = cln.Run(true, true)
	if err != nil {
		t.Fatalf("cln.Run() got err %v", err)
	}

	err = lnd.Run(true, true)
	if err != nil {
		t.Fatalf("lnd.Run() got err %v", err)
	}

	var lightningds []testframework.LightningNode
	switch funder {
	case FUNDER_CLN:
		lightningds = append(lightningds, cln)
		lightningds = append(lightningds, lnd)

	case FUNDER_LND:
		lightningds = append(lightningds, lnd)
		lightningds = append(lightningds, cln)
	default:
		t.Fatalf("unknown fundingNode %s", funder)
	}

	scid, err := lightningds[0].OpenChannel(
		lightningds[1], fundAmt, 0, true, true, true,
	)
	if err != nil {
		t.Fatalf("lightningds[0].OpenChannel() %v", err)
	}

	return &setUpCtx{
		bitcoinNode: bitcoind,
		lnNode:      lnd,
		clnNode:     cln,
		scid:        scid,
	}
}

// Check reference for test here:
// https://github.com/ElementsProject/lightning/blob/2d0778ec38334f5d471b7f2bc54636e04aedeb44/tests/test_misc.py#L2883-L2928
func TestStoringPeerBackup(t *testing.T) {
	// Start nodes and open a channel between both LN nodes.
	harness := setUp(t, uint64(90000), FUNDER_CLN)

	// Check that core lightning sends its storage.
	err := harness.clnNode.WaitForLog(
		"Peer storage sent!", 180*time.Second,
	)
	require.NoError(t, err)

	// Stop core lightning
	require.NoError(t, harness.clnNode.Stop())

	// Start core lightning,
	// this should reconnect it with LndNode since it is now a persistence
	// peer as they have a channel.
	require.NoError(t, harness.clnNode.Run(true, true))

	err = harness.clnNode.WaitForLog(
		"peer_in WIRE_YOUR_PEER_STORAGE", 180*time.Second,
	)
	require.NoError(t, err)
}
