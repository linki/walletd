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

package mock

import (
	"bytes"
	"errors"

	"github.com/google/uuid"
	e2types "github.com/wealdtech/go-eth2-types/v2"
)

type Account struct {
	id         uuid.UUID
	name       string
	privateKey *e2types.BLSPrivateKey
	unlocked   bool
	passphrase []byte
}

// NewAccount creates a new account.
func NewAccount(name string, passphrase []byte) *Account {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	privateKey, err := e2types.GenerateBLSPrivateKey()
	if err != nil {
		panic(err)
	}
	return &Account{
		id:         uuid,
		privateKey: privateKey,
		name:       name,
		passphrase: passphrase,
	}
}

// ID provides the ID for the account.
func (a *Account) ID() uuid.UUID {
	return a.id
}

// Name provides the name for the account.
func (a *Account) Name() string {
	return a.name
}

// PublicKey provides the public key for the account.
func (a *Account) PublicKey() e2types.PublicKey {
	return a.privateKey.PublicKey()
}

// Path provides the path for the account.
// Can be empty if the account is not derived from a path.
func (a *Account) Path() string {
	return ""
}

// Lock locks the account.  A locked account cannot sign.
func (a *Account) Lock() {
	a.unlocked = false
}

// Unlock unlocks the account.  An unlocked account can sign.
func (a *Account) Unlock(passphrase []byte) error {
	if bytes.Equal(a.passphrase, passphrase) {
		a.unlocked = true
		return nil
	}
	return errors.New("invalid passphrase")
}

// IsUnlocked returns true if the account is unlocked.
func (a *Account) IsUnlocked() bool {
	return a.unlocked
}

// Sign signs data with the account.
func (a *Account) Sign(data []byte) (e2types.Signature, error) {
	if !a.IsUnlocked() {
		return nil, errors.New("account is locked")
	}
	return a.Sign(data)
}
