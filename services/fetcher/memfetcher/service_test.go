// Copyright © 2020 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memfetcher_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	e2wallet "github.com/wealdtech/go-eth2-wallet"
	scratch "github.com/wealdtech/go-eth2-wallet-store-scratch"
	e2wtypes "github.com/wealdtech/go-eth2-wallet-types/v2"
	"github.com/wealdtech/walletd/services/fetcher/memfetcher"
)

func TestMain(m *testing.M) {
	if err := e2types.InitBLS(); err != nil {
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		stores []e2wtypes.Store
		err    string
	}{
		{
			name: "Nil",
			err:  "no stores provided",
		},
		{
			name:   "Empty",
			stores: make([]e2wtypes.Store, 0),
			err:    "no stores provided",
		},
		{
			name:   "Good",
			stores: []e2wtypes.Store{scratch.New()},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := memfetcher.New(test.stores)
			if test.err == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.EqualError(t, err, test.err)
			}
		})
	}
}

func TestFetchWallet(t *testing.T) {
	stores, err := createTestStores()
	require.Nil(t, err)
	fetcher, err := memfetcher.New(stores)
	require.Nil(t, err)

	tests := []struct {
		name string
		path string
		err  string
	}{
		{
			name: "Nil",
			err:  "invalid account format",
		},
		{
			name: "Unknown",
			path: "unknown wallet",
			err:  "wallet not found",
		},
		{
			name: "Good",
			path: "Test wallet",
		},
		{
			name: "HDGood",
			path: "Test HD wallet",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := fetcher.FetchWallet(test.path)
			if test.err == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.EqualError(t, err, test.err)
			}
			// Fetch again to check cache path.
			_, err = fetcher.FetchWallet(test.path)
			if test.err == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.EqualError(t, err, test.err)
			}
		})
	}
}

func TestFetchAccount(t *testing.T) {
	stores, err := createTestStores()
	require.Nil(t, err)
	fetcher, err := memfetcher.New(stores)
	require.Nil(t, err)

	tests := []struct {
		name string
		path string
		err  string
	}{
		{
			name: "Nil",
			err:  "invalid account format",
		},
		{
			name: "UnknownWallet",
			path: "unknown wallet",
			err:  "wallet not found",
		},
		{
			name: "UnknownAccount",
			path: "Test wallet/unknown account",
			err:  "no account with name \"unknown account\"",
		},
		{
			name: "Good",
			path: "Test wallet/Test account",
		},
		{
			name: "HDGood",
			path: "Test HD wallet/Test account",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, _, err := fetcher.FetchAccount(test.path)
			if test.err == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.EqualError(t, err, test.err)
			}
			// Fetch again to test cache path.
			_, _, err = fetcher.FetchAccount(test.path)
			if test.err == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.EqualError(t, err, test.err)
			}
		})
	}
}

func TestFetchAccountByKey(t *testing.T) {
	stores, err := createTestStores()
	require.Nil(t, err)
	fetcher, err := memfetcher.New(stores)
	require.Nil(t, err)

	wallet, err := e2wallet.OpenWallet("Test wallet", e2wallet.WithStore(stores[0]))
	require.Nil(t, err)
	account, err := wallet.AccountByName("Test account")
	require.Nil(t, err)

	hdWallet, err := e2wallet.OpenWallet("Test HD wallet", e2wallet.WithStore(stores[0]))
	require.Nil(t, err)
	hdAccount, err := hdWallet.AccountByName("Test account")
	require.Nil(t, err)

	tests := []struct {
		name string
		key  []byte
		err  string
	}{
		{
			name: "Nil",
			err:  "account not found",
		},
		{
			name: "Empty",
			key:  []byte{},
			err:  "account not found",
		},
		{
			name: "UnknownAccount",
			key:  []byte{0x00},
			err:  "account not found",
		},
		{
			name: "Good",
			key:  account.PublicKey().Marshal(),
		},
		{
			name: "HDGood",
			key:  hdAccount.PublicKey().Marshal(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, _, err := fetcher.FetchAccountByKey(test.key)
			if test.err == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.EqualError(t, err, test.err)
			}
			// Fetch again to test cache path.
			_, _, err = fetcher.FetchAccountByKey(test.key)
			if test.err == "" {
				require.Nil(t, err)
			} else {
				require.NotNil(t, err)
				require.EqualError(t, err, test.err)
			}
		})
	}
}

func TestWalletLocking(t *testing.T) {
	stores, err := createTestStores()
	require.Nil(t, err)
	fetcher, err := memfetcher.New(stores)
	require.Nil(t, err)

	// Kick off 16 goroutines each retrieving the wallet 1024 times.
	for i := 0; i < 16; i++ {
		go func() {
			for i := 0; i < 1024; i++ {
				wallet, err := fetcher.FetchWallet("Test wallet")
				assert.Nil(t, err)
				assert.NotNil(t, wallet)
			}
		}()
	}
}

func TestAccountLocking(t *testing.T) {
	stores, err := createTestStores()
	require.Nil(t, err)
	fetcher, err := memfetcher.New(stores)
	require.Nil(t, err)

	// Kick off 16 goroutines each retrieving the account 1024 times.
	for i := 0; i < 16; i++ {
		go func() {
			for i := 0; i < 1024; i++ {
				wallet, account, err := fetcher.FetchAccount("Test wallet/Test account")
				assert.Nil(t, err)
				assert.NotNil(t, wallet)
				assert.NotNil(t, account)
			}
		}()
	}
}

// createTestStores is a helper to create and populate some stores for testing.
func createTestStores() ([]e2wtypes.Store, error) {
	store := scratch.New()
	walletID := uuid.New()
	err := store.StoreWallet(walletID, "Test wallet", []byte(fmt.Sprintf(`{"uuid":"%s","version":1,"name":"Test wallet","type":"non-deterministic"}`, walletID.String())))
	if err != nil {
		return nil, err
	}
	wallet, err := e2wallet.OpenWallet("Test wallet", e2wallet.WithStore(store))
	if err != nil {
		return nil, err
	}
	err = wallet.Unlock(nil)
	if err != nil {
		return nil, err
	}
	_, err = wallet.CreateAccount("Test account", []byte{})
	if err != nil {
		return nil, err
	}

	walletID = uuid.New()
	err = store.StoreWallet(walletID, "Test HD wallet", []byte(fmt.Sprintf(`{"crypto":{"checksum":{"function":"sha256","message":"2f56b0b5338c5afed07f88698686bb214c6fb42b9c1a812dc2f15e58aedec18c","params":{}},"cipher":{"function":"aes-128-ctr","message":"172821d450d10846b1f033d3955c500f1e2ae01304363abb6ad7fcc7e5020b37","params":{"iv":"66cafd4641b91f3e31d5dcc768ddb641"}},"kdf":{"function":"pbkdf2","message":"","params":{"c":16,"dklen":32,"prf":"hmac-sha256","salt":"45e396e8375da8a410ad992bee063bfe68c8084aa290b58e8408de05bd01ca82"}}},"name":"Test HD wallet","nextaccount":0,"type":"hierarchical deterministic","uuid":"%s","version":1}`, walletID.String())))
	if err != nil {
		return nil, err
	}
	wallet, err = e2wallet.OpenWallet("Test HD wallet", e2wallet.WithStore(store))
	if err != nil {
		return nil, err
	}
	err = wallet.Unlock([]byte("secret"))
	if err != nil {
		return nil, err
	}
	_, err = wallet.CreateAccount("Test account", []byte{})
	if err != nil {
		return nil, err
	}

	return []e2wtypes.Store{store}, nil
}
