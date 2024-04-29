## Presrequisites

1. You would need an lnd, lightningd and bitcoind binary in your PATH to run 
   this test.

Recommended to build the binary with the project versions linked below, to ensure compatibility:

- [`lightningd`](https://github.com/ElementsProject/lightning/tree/95a92b6e4bf521456c1188ac8ecea2a49fa5f22f)
- [`lnd`](https://github.com/Chinwendu20/lnd/tree/peer-backup)

## Steps

- Ensure you have the lnd, lightningd and bitcoind binary on the PATH of your system.
- Run go test on the peerstoragetest package.

## Notes

This directory currently contains interoperability test for the peer backup 
feature between core lightning and lnd. The test is contained in the cln_test.go file.
