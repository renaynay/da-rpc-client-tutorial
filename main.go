package main

import (
	"context"
	"fmt"

	// please reference the go.mod file in this repository in order to correctly import
	// this package
	"github.com/celestiaorg/celestia-node/api/rpc/client"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create a new client by dialing the celestia-node's RPC endpoint --
	// by default, celestia-nodes run RPC on port 26658
	rpc, err := client.NewClient(ctx, "<insert RPC addr:port here>", "<insert JWT token here>")
	if err != nil {
		panic(err)
	}

	// call the GetByHeight method on the `HeaderModule` that returns a header to you
	// by the given height (20)
	header, err := rpc.Header.GetByHeight(context.Background(), 20)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got header: %v", header)

	// check whether the "shares" (a term used by celestia to refer to the block's
	// data) is available in the network
	err = rpc.Share.SharesAvailable(ctx, header.DAH)
	if err != nil {
		fmt.Printf("shares not available: %s", err.Error())
		panic(err)
	}
	fmt.Printf("Shares avaialble for header at height %d", header.Height())

	// close the client when you are finished :)
	rpc.Close()
}
