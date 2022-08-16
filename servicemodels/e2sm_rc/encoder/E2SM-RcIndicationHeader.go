// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package encoder

import (
	"encoding/hex"
	"github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc/v1/choiceOptions"
	e2smrcv1 "github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_rc/v1/e2sm-rc-ies"

	"github.com/wangxn2015/onos-lib-go/pkg/asn1/aper"
	"github.com/wangxn2015/onos-lib-go/pkg/errors"
)

func PerEncodeE2SmRcIndicationHeader(msg *e2smrcv1.E2SmRcIndicationHeader) ([]byte, error) {

	log.Debugf("Obtained E2SM-RcIndicationHeader message is\n%v", msg)
	if err := msg.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcIndicationHeader PDU %s", err.Error())
	}

	per, err := aper.MarshalWithParams(msg, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}
	log.Debugf("Encoded E2SM-RcIndicationHeader PER bytes are\n%v", hex.Dump(per))

	return per, nil
}

func PerDecodeE2SmRcIndicationHeader(per []byte) (*e2smrcv1.E2SmRcIndicationHeader, error) {

	log.Debugf("Obtained E2SM-RcIndicationHeader PER bytes are\n%v", hex.Dump(per))

	result := e2smrcv1.E2SmRcIndicationHeader{}
	err := aper.UnmarshalWithParams(per, &result, "valueExt", choiceOptions.E2smRcChoicemap, nil)
	if err != nil {
		return nil, err
	}

	log.Debugf("Decoded E2SM-RcIndicationHeader from PER is\n%v", &result)
	if err = result.Validate(); err != nil {
		return nil, errors.NewInvalid("error validating E2SM-RcIndicationHeader PDU %s", err.Error())
	}

	return &result, nil
}
