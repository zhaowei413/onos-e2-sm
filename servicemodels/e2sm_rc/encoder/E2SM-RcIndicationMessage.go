// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/choiceOptions"
	e2smrcv1 "github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"

	"github.com/onosproject/onos-lib-go/pkg/asn1/aper"
	"github.com/onosproject/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmRcIndicationMessage(msg *e2smrcv1.E2SmRcIndicationMessage) ([]byte, error) {

	log.Debugf("Obtained E2SM-RcIndicationMessage message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcIndicationMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RcIndicationMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcIndicationMessage(per []byte) (*e2smrcv1.E2SmRcIndicationMessage, error) {

	log.Debugf("Obtained E2SM-RcIndicationMessage PER bytes are\n%v", hex.Dump(per))

	result := e2smrcv1.E2SmRcIndicationMessage{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RcIndicationMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcIndicationMessage PDU %s", err.Error())
	}

	return &result, nil
}

func PerEncodeE2SmRcIndicationULDCCHMessage(msg *e2smrcv1.ULDCCHMessageItem) ([]byte, error) {

	log.Debugf("Obtained E2SM-RcIndicationULDCCHMessage message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcIndicationULDCCHMessage PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RcIndicationULDCCHMessage PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcIndicationULDCCMessage(per []byte) (*e2smrcv1.ULDCCHMessageItem, error) {

	log.Debugf("Obtained E2SM-RcIndicationULDCCHMessage PER bytes are\n%v", hex.Dump(per))

	result := e2smrcv1.ULDCCHMessageItem{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RcIndicationULDCCHMessage from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcIndicationULDCCHMessage PDU %s", err.Error())
	}

	return &result, nil
}

