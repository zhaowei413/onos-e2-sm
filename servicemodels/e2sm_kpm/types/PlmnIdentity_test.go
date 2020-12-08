// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package types

import (
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestPlmnID_NewPlmnId(t *testing.T) {

	value := []byte{0x22, 0x21, 0x20}
	plmnID, err := NewPlmnID(value)
	assert.NilError(t, err)

	assert.Assert(t, reflect.TypeOf(PlmnIdentity{}) == reflect.TypeOf(*plmnID), "plmnID{} types are mismatched")
	assert.DeepEqual(t, plmnID.Value, value)
}

func TestPlmnID_GetValue(t *testing.T) {

	value := []byte{0x22, 0x21, 0x20}
	plmnID, err := NewPlmnID(value)
	assert.NilError(t, err)

	assert.DeepEqual(t, plmnID.GetValue(), value)
}

func TestPlmnID_GetPlmnID(t *testing.T) {

	value := []byte{0x22, 0x21, 0x21}
	plmnID1, err := NewPlmnID(value)
	assert.NilError(t, err)
	plmnID2 := plmnID1.GetPlmnID()

	assert.DeepEqual(t, plmnID1.GetValue(), plmnID2.GetValue())
}