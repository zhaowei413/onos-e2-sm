//  SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
//  SPDX-License-Identifier: Apache-2.0

package pdubuilder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/encoder"
	e2smcommonies "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-common-ies"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"
	"github.com/onosproject/onos-lib-go/api/asn1/v1/asn1"
	"gotest.tools/assert"
	"testing"
)

func TestE2SmRcIndicationMessageFormat1(t *testing.T) {

	ranPReportedList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat1Item, 0)

	ranParameterValue, err := CreateRanparameterValueBoolean(true)
	assert.NilError(t, err)

	ranParameterValueType, err := CreateRanparameterValueTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	item, err := CreateE2SmRcIndicationMessageFormat1Item(1, ranParameterValueType)
	assert.NilError(t, err)

	ranPReportedList = append(ranPReportedList, item)

	msg, err := CreateE2SmRcIndicationMessageFormat1(ranPReportedList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcIndicationMessage(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcIndicationMessage(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
	assert.Equal(t, msg.String(), result.String())
}

func TestE2SmRcIndicationMessageFormat2(t *testing.T) {

	ueParameterList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat2Item, 0)

	ueID1, err := CreateUeIDGNb(1, []byte{0xAB, 0xCD, 0xEF}, []byte{0x21}, []byte{0xFF, 0xC0}, []byte{0xFC})
	assert.NilError(t, err)
	ueID1.GetGNbUeid().SetRanUeID([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}).SetMNgRanUeXnApID(1)

	ueID2, err := CreateUeIDNgEnb(1, []byte{0x1B, 0xD9, 0xEF}, []byte{0x21}, []byte{0xFF, 0xC0}, []byte{0xFC})
	assert.NilError(t, err)

	ranPList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat2RanparameterItem, 0)

	ranParameterValue, err := CreateRanparameterValueBoolean(true)
	assert.NilError(t, err)

	ranParameterValueType, err := CreateRanparameterValueTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	ranPItem, err := CreateE2SmRcIndicationMessageFormat2RanparameterItem(1, ranParameterValueType)
	assert.NilError(t, err)
	ranPList = append(ranPList, ranPItem)

	item1, err := CreateE2SmRcIndicationMessageFormat2Item(ueID1, ranPList)
	assert.NilError(t, err)
	item2, err := CreateE2SmRcIndicationMessageFormat2Item(ueID2, ranPList)
	assert.NilError(t, err)

	ueParameterList = append(ueParameterList, item1)
	ueParameterList = append(ueParameterList, item2)

	msg, err := CreateE2SmRcIndicationMessageFormat2(ueParameterList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcIndicationMessage(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcIndicationMessage(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
	assert.Equal(t, msg.String(), result.String())
}

func TestE2SmRcIndicationMessageFormat3(t *testing.T) {

	cellInfoList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat3Item, 0)

	nrCgi, err := CreateNrCgi([]byte{0xFF, 0xCC, 0xFF}, &asn1.BitString{
		Value: []byte{0xFF, 0xAA, 0xBB, 0xCC, 0xF0},
		Len:   36,
	})
	assert.NilError(t, err)
	cgi1, err := CreateCgiNRCgi(nrCgi)
	assert.NilError(t, err)

	servingCellPci, err := CreateServingCellPciNR(1)
	assert.NilError(t, err)
	servingCellArfcn, err := CreateServingCellArfcnNR(1)
	assert.NilError(t, err)

	neighborCellList := make([]*e2smrcv1.NeighborCellItem, 0)
	nrArfcn, err := CreateNrArfcn(1)
	assert.NilError(t, err)

	nrfrequencyBandList := &e2smcommonies.NrfrequencyBandList{
		Value: make([]*e2smcommonies.NrfrequencyBandItem, 0),
	}

	supportedSulbandList := make([]*e2smcommonies.SupportedSulfreqBandItem, 0)
	supportedSulbandItem, err := CreateSupportedSulfreqBandItem(1024)
	assert.NilError(t, err)
	supportedSulbandList = append(supportedSulbandList, supportedSulbandItem)

	frequencyBandItem, err := CreateNrfrequencyBandItem(1, &e2smcommonies.SupportedSulbandList{
		Value: supportedSulbandList,
	})
	assert.NilError(t, err)

	nrfrequencyBandList.Value = append(nrfrequencyBandList.Value, frequencyBandItem)

	nrFrequencyInfo, err := CreateNrfrequencyInfo(nrArfcn, nrfrequencyBandList)
	assert.NilError(t, err)
	nrFrequencyInfo.SetFrequencyShift7P5Khz(CreateNrfrequencyShift7P5KhzTrue())

	neighborCellItem, err := CreateNeighborCellItemRanTypeChoiceNr(nrCgi, 1, []byte{0xFF, 0xFF, 0xFF}, CreateNRModeInfoFDD(),
		nrFrequencyInfo, CreateX2XNEstablishedTrue(), CreateHOValidatedTrue(), 1)
	assert.NilError(t, err)
	neighborCellList = append(neighborCellList, neighborCellItem)

	neighborRelationTable, err := CreateNeighborRelationInfo(servingCellPci, servingCellArfcn, &e2smrcv1.NeighborCellList{
		Value: neighborCellList,
	})
	assert.NilError(t, err)

	item1, err := CreateE2SmRcIndicationMessageFormat3Item(cgi1)
	assert.NilError(t, err)
	item1.SetCellDeleted(false).SetCellContextInfo([]byte{0xFF, 0xFF, 0xFF, 0xFF}).SetNeighborRelationTable(neighborRelationTable)

	item2, err := CreateE2SmRcIndicationMessageFormat3Item(cgi1)
	assert.NilError(t, err)
	item2.SetCellDeleted(true).SetCellContextInfo([]byte{0x00})

	cellInfoList = append(cellInfoList, item1)
	cellInfoList = append(cellInfoList, item2)

	msg, err := CreateE2SmRcIndicationMessageFormat3(cellInfoList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcIndicationMessage(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcIndicationMessage(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
	assert.Equal(t, msg.String(), result.String())
}

func TestE2SmRcIndicationMessageFormat4(t *testing.T) {

	ueInfoList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat4ItemUe, 0)
	cellInfoList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat4ItemCell, 0)

	ueID, err := CreateUeIDNgEnbDu(1)
	assert.NilError(t, err)

	eutraCgi, err := CreateEutraCgi([]byte{0xAC, 0xFE, 0xDD}, &asn1.BitString{
		Value: []byte{0xFF, 0xAA, 0xBB, 0xCC, 0xF0},
		Len:   36,
	})
	assert.NilError(t, err)

	cgi, err := CreateCgiEutraCgi(eutraCgi)
	assert.NilError(t, err)

	ueInfoItem, err := CreateE2SmRcIndicationMessageFormat4ItemUe(ueID, cgi)
	assert.NilError(t, err)
	ueInfoItem.SetUeContextInfo([]byte{0xFF, 0x11, 0xFF})

	cellInfoItem, err := CreateE2SmRcIndicationMessageFormat4ItemCell(cgi)
	assert.NilError(t, err)
	cellInfoItem.SetCellContextInfo([]byte{0xFF, 0xAA})

	msg, err := CreateE2SmRcIndicationMessageFormat4(ueInfoList, cellInfoList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcIndicationMessage(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcIndicationMessage(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcIndicationMessageFormat5(t *testing.T) {

	ranPRequestedList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat5Item, 0)

	ranParameterValue, err := CreateRanparameterValueBoolean(true)
	assert.NilError(t, err)

	ranParameterValueType, err := CreateRanparameterValueTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	cellInfoItem, err := CreateE2SmRcIndicationMessageFormat5Item(1, ranParameterValueType)
	assert.NilError(t, err)

	ranPRequestedList = append(ranPRequestedList, cellInfoItem)

	msg, err := CreateE2SmRcIndicationMessageFormat5(ranPRequestedList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcIndicationMessage(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcIndicationMessage(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcIndicationMessageFormat6(t *testing.T) {

	ricInsertStyleList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat6StyleItem, 0)

	ricInsertIndicationList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat6IndicationItem, 0)

	ranPInsertIndicationList := make([]*e2smrcv1.E2SmRcIndicationMessageFormat6RanpItem, 0)

	ranParameterValue, err := CreateRanparameterValueBoolean(true)
	assert.NilError(t, err)

	ranParameterValueType, err := CreateRanparameterValueTypeChoiceElementTrue(ranParameterValue)
	assert.NilError(t, err)

	ranPIndicationItem, err := CreateE2SmRcIndicationMessageFormat6RanpItem(1, ranParameterValueType)
	assert.NilError(t, err)

	ranPInsertIndicationList = append(ranPInsertIndicationList, ranPIndicationItem)

	indicationItem, err := CreateE2SmRcIndicationMessageFormat6IndicationItem(1, ranPInsertIndicationList)
	assert.NilError(t, err)

	ricInsertIndicationList = append(ricInsertIndicationList, indicationItem)

	styleItem, err := CreateE2SmRcIndicationMessageFormat6StyleItem(1, ricInsertIndicationList)
	assert.NilError(t, err)

	ricInsertStyleList = append(ricInsertStyleList, styleItem)

	msg, err := CreateE2SmRcIndicationMessageFormat6(ricInsertStyleList)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcIndicationMessage(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcIndicationMessage(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
}

func TestE2SmRcIndicationMessageULDCCH(t *testing.T) {

	measId := int32(1)
	servCellId := int32(5)

	ServingCell_physCellId := int32(5)
	ServingCell_rsrp := int32(50)
	ServingCell_rsrq := int32(67)
	ServingCell_sinr := int32(75)

	NeighCells_physCellId1 := int32(6)
	NeighCells_rsrp1 := int32(50)
	NeighCells_rsrq1 := int32(67)
	NeighCells_sinr1 := int32(75)

	NeighCells_physCellId2 := int32(7)
	NeighCells_rsrp2 := int32(50)
	NeighCells_rsrq2 := int32(67)
	NeighCells_sinr2 := int32(75)

	PhysCellId, err := CreatePhysCellId(ServingCell_physCellId)
	RSRPRange, err := CreateRSRPRange(ServingCell_rsrp)
	RSRQRange, err := CreateRSRQRange(ServingCell_rsrq)
	SINRRange, err := CreateSINRRange(ServingCell_sinr)
	ResultsSsbCell, err := CreateMeasQuantityResults(RSRPRange, RSRQRange, SINRRange)
	CellResults, err := CreateCellResults(ResultsSsbCell, nil)
	MeasResult, err := CreateMeasResult(CellResults, nil)
	MeasResultNR, err := CreateMeasResultNR(PhysCellId, MeasResult, nil)
	ServCellIndex, err := CreateServCellIndex(servCellId)
	MeasResultServMO, err := CreateMeasResultServMO(ServCellIndex, MeasResultNR, nil)
	List := make([]*e2smrcv1.MeasResultServMO, 0)
	List = append(List, MeasResultServMO)
	MeasResultServMOList, err := CreateMeasResultServMOList(List)
	assert.NilError(t, err)

	PhysCellId, err = CreatePhysCellId(NeighCells_physCellId1)
	RSRPRange, err = CreateRSRPRange(NeighCells_rsrp1)
	RSRQRange, err = CreateRSRQRange(NeighCells_rsrq1)
	SINRRange, err = CreateSINRRange(NeighCells_sinr1)
	ResultsSsbCell, err = CreateMeasQuantityResults(RSRPRange, RSRQRange, SINRRange)
	CellResults, err = CreateCellResults(ResultsSsbCell, nil)
	MeasResult, err = CreateMeasResult(CellResults, nil)
	MeasResultNR, err = CreateMeasResultNR(PhysCellId, MeasResult, nil)
	assert.NilError(t, err)

	NRList := make([]*e2smrcv1.MeasResultNR, 0)
	NRList = append(NRList, MeasResultNR)

	PhysCellId, err = CreatePhysCellId(NeighCells_physCellId2)
	RSRPRange, err = CreateRSRPRange(NeighCells_rsrp2)
	RSRQRange, err = CreateRSRQRange(NeighCells_rsrq2)
	SINRRange, err = CreateSINRRange(NeighCells_sinr2)
	ResultsSsbCell, err = CreateMeasQuantityResults(RSRPRange, RSRQRange, SINRRange)
	CellResults, err = CreateCellResults(ResultsSsbCell, nil)
	MeasResult, err = CreateMeasResult(CellResults, nil)
	MeasResultNR, err = CreateMeasResultNR(PhysCellId, MeasResult, nil)
	assert.NilError(t, err)

	NRList = append(NRList, MeasResultNR)

	MeasResultListNR, err := CreateMeasResultListNR(NRList)
	MeasResultNeighCells, err := CreateMeasResultNeighCells_MeasResultListNr(MeasResultListNR)
	assert.NilError(t, err)

	MeasId, err := CreateMeasId(measId)
	MeasResults, err := CreateMeasResults(MeasId, MeasResultServMOList, MeasResultNeighCells)
	assert.NilError(t, err)

	MeasurementReportitem, err := CreateMeasurementReportitem(MeasResults, nil, nil)
	MeasurementReport, err := CreateMeasurementReport_MeasurementReport(MeasurementReportitem)
	C1, err := CreateC1_MeasurementReport(MeasurementReport)
	ULDCCHMessageTypeItem, err := CreateULDCCHMessageTypeItem_C1(C1)
	msg, err := CreateULDCCHMessageItem(ULDCCHMessageTypeItem)
	assert.NilError(t, err)

	aper, err := encoder.PerEncodeE2SmRcIndicationULDCCHMessage(msg)
	assert.NilError(t, err)
	t.Logf("APER bytes are\n%v", hex.Dump(aper))

	result, err := encoder.PerDecodeE2SmRcIndicationULDCCMessage(aper)
	assert.NilError(t, err)
	t.Logf("Decoded message is\n%v", result)
	assert.Equal(t, msg.String(), result.String())
}
