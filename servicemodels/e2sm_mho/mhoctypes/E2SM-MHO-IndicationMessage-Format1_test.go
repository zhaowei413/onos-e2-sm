// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package mhoctypes

import (
	"encoding/hex"
	"fmt"
	//pdubuilder "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/pdubuilder"
	e2sm_mho "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho/v1/e2sm-mho"
	"gotest.tools/assert"
	"testing"
)

func createE2SmMhoIndicationMessageFormat1Msg() (*e2sm_mho.E2SmMhoIndicationMessageFormat1, error) {

	// e2SmMhoIndicationMessageFormat1 := pdubuilder.CreateE2SmMhoIndicationMessageFormat1() //ToDo - fill in arguments here(if this function exists

	e2SmMhoIndicationMessageFormat1 := e2sm_mho.E2SmMhoIndicationMessageFormat1{
		UeId: &e2sm_mho.UeIdentity{Value: "1234"},
		Rsrp: &e2sm_mho.Rsrp{Value: 1234},
	}

	if err := e2SmMhoIndicationMessageFormat1.Validate(); err != nil {
		return nil, fmt.Errorf("error validating E2SmMhoIndicationMessageFormat1 %s", err.Error())
	}
	return &e2SmMhoIndicationMessageFormat1, nil
}

func Test_XerEncodingE2SmMhoIndicationMessageFormat1(t *testing.T) {

	e2SmMhoIndicationMessageFormat1, err := createE2SmMhoIndicationMessageFormat1Msg()
	assert.NilError(t, err, "Error creating E2SmMhoIndicationMessageFormat1 PDU")

	xer, err := xerEncodeE2SmMhoIndicationMessageFormat1(e2SmMhoIndicationMessageFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 126, len(xer)) //ToDo - adjust length of the XER encoded message
	t.Logf("E2SmMhoIndicationMessageFormat1 XER\n%s", string(xer))

	result, err := xerDecodeE2SmMhoIndicationMessageFormat1(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoIndicationMessageFormat1 XER - decoded\n%v", result)
	//ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoIndicationMessageFormat1.GetUeId(), result.GetUeId())
	//assert.Equal(t, e2SmMhoIndicationMessageFormat1.GetRsrp(), result.GetRsrp())

}

func Test_PerEncodingE2SmMhoIndicationMessageFormat1(t *testing.T) {

	e2SmMhoIndicationMessageFormat1, err := createE2SmMhoIndicationMessageFormat1Msg()
	assert.NilError(t, err, "Error creating E2SmMhoIndicationMessageFormat1 PDU")

	per, err := perEncodeE2SmMhoIndicationMessageFormat1(e2SmMhoIndicationMessageFormat1)
	assert.NilError(t, err)
	assert.Equal(t, 9, len(per)) // ToDo - adjust length of the PER encoded message
	t.Logf("E2SmMhoIndicationMessageFormat1 PER\n%v", hex.Dump(per))

	result, err := perDecodeE2SmMhoIndicationMessageFormat1(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("E2SmMhoIndicationMessageFormat1 PER - decoded\n%v", result)
	//ToDo - adjust field's verification
	//assert.Equal(t, e2SmMhoIndicationMessageFormat1.GetUeId(), result.GetUeId())
	//assert.Equal(t, e2SmMhoIndicationMessageFormat1.GetRsrp(), result.GetRsrp())

}