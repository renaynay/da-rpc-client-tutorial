package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/celestiaorg/celestia-app/pkg/appconsts"
	"github.com/celestiaorg/celestia-node/blob"
	"github.com/celestiaorg/celestia-node/share"
)

const (
	perByteGasTolerance = 2
	pfbGasFixedCost     = 80000
)

// some helper functions copied from internal Celestia app repo
func ParseV0Namespace(param string) (share.Namespace, error) {
	userBytes, err := DecodeToBytes(param)
	if err != nil {
		return nil, err
	}

	// if the namespace ID is <= 10 bytes, left pad it with 0s
	return share.NewBlobNamespaceV0(userBytes)
}

func DecodeToBytes(param string) ([]byte, error) {
	if strings.HasPrefix(param, "0x") {
		decoded, err := hex.DecodeString(param[2:])
		if err != nil {
			return nil, fmt.Errorf("error decoding namespace ID: %w", err)
		}
		return decoded, nil
	}
	// otherwise, it's just a base64 string
	decoded, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		return nil, fmt.Errorf("error decoding namespace ID: %w", err)
	}
	return decoded, nil
}

func EstimateGas(blobs ...*blob.Blob) uint64 {
	totalByteCount := 0
	for _, blob := range blobs {
		totalByteCount += len(blob.Data) + appconsts.NamespaceSize
	}
	variableGasAmount := (appconsts.DefaultGasPerBlobByte + perByteGasTolerance) * totalByteCount

	return uint64(variableGasAmount + pfbGasFixedCost)
}
