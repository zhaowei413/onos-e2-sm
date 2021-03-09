// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package kpmv2ctypes

import (
	e2sm_kpm_v2 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2/v2/e2sm-kpm-ies"
	"gotest.tools/assert"
	"testing"
)

func createLabelInfoList() *e2sm_kpm_v2.LabelInfoList{

	labelInfoList := &e2sm_kpm_v2.LabelInfoList{
		Value: make([]*e2sm_kpm_v2.LabelInfoItem, 0),
	}

	item := &e2sm_kpm_v2.LabelInfoItem{
		MeasLabel: &e2sm_kpm_v2.MeasurementLabel{
			PlmnId: &e2sm_kpm_v2.PlmnIdentity{
				Value: []byte{0x1, 0x2, 0x3},
			},
			SliceId: &e2sm_kpm_v2.Snssai{
				SD:  []byte{0x01, 0x02, 0x03},
				SSt: []byte{0x01},
			},
			FiveQi: &e2sm_kpm_v2.FiveQi{
				Value: 23,
			},
			QCi: &e2sm_kpm_v2.Qci{
				Value: 24,
			},
			QCimax: &e2sm_kpm_v2.Qci{
				Value: 30,
			},
			QCimin: &e2sm_kpm_v2.Qci{
				Value: 1,
			},
			ARpmax: &e2sm_kpm_v2.Arp{
				Value: 50,
			},
			ARpmin: &e2sm_kpm_v2.Arp{
				Value: 1,
			},
			BitrateRange:     25,
			LayerMuMimo:      1,
			SUm:              0,
			DistBinX:         123,
			DistBinY:         456,
			DistBinZ:         789,
			PreLabelOverride: 0,
			StartEndInd:      1,
		},
	}
	labelInfoList.Value = append(labelInfoList.Value, item)

	return labelInfoList
}

func Test_xerEncodeLabelInfoLabel(t *testing.T) {

	lil := createLabelInfoList()

	xer, err := xerEncodeLabelInfoList(lil)
	assert.NilError(t, err)
	assert.Equal(t, 778, len(xer))
	t.Logf("LabelInfoList XER\n%s", string(xer))
}

func Test_xerDecodeLabelInfoList(t *testing.T) {

	lil := createLabelInfoList()

	xer, err := xerEncodeLabelInfoList(lil)
	assert.NilError(t, err)
	assert.Equal(t, 778, len(xer))
	t.Logf("LabelInfoList XER\n%s", string(xer))

	result, err := xerDecodeLabelInfoList(xer)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("LabelInfoList XER - decoded\n%s", result)
}

func Test_perEncodeLabelInfoList(t *testing.T) {

	lil := createLabelInfoList()

	per, err := perEncodeLabelInfoList(lil)
	assert.NilError(t, err)
	assert.Equal(t, 39, len(per))
	t.Logf("LabelInfoList PER\n%s", string(per))
}

func Test_perDecodeLabelInfoList(t *testing.T) {

	lil := createLabelInfoList()

	per, err := perEncodeLabelInfoList(lil)
	assert.NilError(t, err)
	assert.Equal(t, 39, len(per))
	t.Logf("LabelInfoList PER\n%s", string(per))

	result, err := perDecodeLabelInfoList(per)
	assert.NilError(t, err)
	assert.Assert(t, result != nil)
	t.Logf("LabelInfoList PER - decoded\n%v", result)
}