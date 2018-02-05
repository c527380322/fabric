/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mgmt

import (
	"testing"

	"github.com/hyperledger/fabric/msp"
	"github.com/stretchr/testify/assert"
)

func TestGetManagerForChains(t *testing.T) {
	// MSPManager for channel does not exist prior to this call
	mspMgr1 := GetManagerForChain("test")
	// ensure MSPManager is set
	if mspMgr1 == nil {
		t.FailNow()
	}

	// MSPManager for channel now exists
	mspMgr2 := GetManagerForChain("test")
	// ensure MSPManager returned matches the first result
	if mspMgr2 != mspMgr1 {
		t.FailNow()
	}
}

func TestGetManagerForChains_usingMSPConfigHandlers(t *testing.T) {
	XXXSetMSPManager("foo", msp.NewMSPManager())
	msp2 := GetManagerForChain("foo")
	// return value should be set because the MSPManager was initialized
	if msp2 == nil {
		t.FailNow()
	}
}

func TestGetIdentityDeserializer(t *testing.T) {
	XXXSetMSPManager("baz", msp.NewMSPManager())
	ids := GetIdentityDeserializer("baz")
	assert.NotNil(t, ids)
	ids = GetIdentityDeserializer("")
	assert.NotNil(t, ids)
}

func TestGetLocalSigningIdentityOrPanic(t *testing.T) {
	sid := GetLocalSigningIdentityOrPanic()
	assert.NotNil(t, sid)
}

func TestUpdateLocalMspCache(t *testing.T) {
	// reset localMsp to force it to be initialized on the first call
	localMsp = nil

	// first call should initialize local MSP and returned the cached version
	firstMsp := GetLocalMSP()
	// second call should return the same
	secondMsp := GetLocalMSP()
	// third call should return the same
	thirdMsp := GetLocalMSP()

	// the same (non-cached if not patched) instance
	if thirdMsp != secondMsp {
		t.Fatalf("thirdMsp != secondMsp")
	}
	// first (cached) and second (non-cached) different unless patched
	if firstMsp != secondMsp {
		t.Fatalf("firstMsp != secondMsp")
	}
}
