package stellarnet

import (
	"errors"
	"fmt"

	"github.com/stellar/go/clients/horizon"
)

// HorizonStatus returns the root status information from the global horizon
// server.
func HorizonStatus() (horizon.Root, error) {
	return HorizonStatusForClient(Client())
}

// HorizonStatusForClient returns the root status information from client's horizon
// server.
func HorizonStatusForClient(client *horizon.Client) (horizon.Root, error) {
	if client == nil {
		return horizon.Root{}, errors.New("nil horizon client")
	}

	return client.Root()
}

// Ping returns a formatted string of info about the horizon server
// that stellarnet is connected to.
func Ping() (string, error) {
	status, err := HorizonStatus()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("horizon ver: %s, stellar-core ver: %s, horizon seqno: %d, core seqno: %d", status.HorizonVersion, status.StellarCoreVersion, status.HorizonSequence, status.CoreSequence), nil
}
