// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package kpmv2

import (
	"encoding/hex"
	e2smkpmv2 "github.com/wangxn2015/onos-e2-sm/servicemodels/e2sm_kpm_v2_go/v2/e2sm-kpm-v2-go"
	"github.com/wangxn2015/onos-lib-go/pkg/asn1/aper"
	hexlib "github.com/wangxn2015/onos-lib-go/pkg/hex"
	"gotest.tools/assert"
	"testing"
)

var refPerMeasCI = "00000000  20 00 7a 00 01 42 10 01  15 1f ff f0 01 02 03 40  | .z..B.........@|\n" +
	"00000010  40 01 02 03 00 17 68 18  00 1e 00 01 70 00 00 18  |@.....h.....p...|\n" +
	"00000020  00 00 00 00 00 7a 00 01  c7 00 03 14 20           |.....z...... |"

func createMeasurementCondItem() (*e2smkpmv2.MeasurementCondItem, error) {

	measCondItem := &e2smkpmv2.MeasurementCondItem{
		MeasType: &e2smkpmv2.MeasurementType{
			MeasurementType: &e2smkpmv2.MeasurementType_MeasId{
				MeasId: &e2smkpmv2.MeasurementTypeId{
					Value: 123,
				},
			},
		},
		MatchingCond: &e2smkpmv2.MatchingCondList{
			Value: make([]*e2smkpmv2.MatchingCondItem, 0),
		},
	}

	mci1 := &e2smkpmv2.MatchingCondItem{
		MatchingCondItem: &e2smkpmv2.MatchingCondItem_TestCondInfo{
			TestCondInfo: &e2smkpmv2.TestCondInfo{
				TestType: &e2smkpmv2.TestCondType{
					TestCondType: &e2smkpmv2.TestCondType_AMbr{
						AMbr: e2smkpmv2.AMBR_AMBR_TRUE,
					},
				},
				TestExpr: e2smkpmv2.TestCondExpression_TEST_COND_EXPRESSION_GREATERTHAN,
				TestValue: &e2smkpmv2.TestCondValue{
					TestCondValue: &e2smkpmv2.TestCondValue_ValueInt{
						ValueInt: 21,
					},
				},
			},
		},
	}
	measCondItem.MatchingCond.Value = append(measCondItem.MatchingCond.Value, mci1)

	var br int32 = 25
	var lmm int32 = 1
	var dbx int32 = 123
	var dby int32 = 456
	var dbz int32 = 789
	sum := e2smkpmv2.SUM_SUM_TRUE
	plo := e2smkpmv2.PreLabelOverride_PRE_LABEL_OVERRIDE_TRUE
	seind := e2smkpmv2.StartEndInd_START_END_IND_END

	mci2 := &e2smkpmv2.MatchingCondItem{
		MatchingCondItem: &e2smkpmv2.MatchingCondItem_MeasLabel{
			MeasLabel: &e2smkpmv2.MeasurementLabel{
				PlmnId: &e2smkpmv2.PlmnIdentity{
					Value: []byte{0x1, 0x2, 0x3},
				},
				SliceId: &e2smkpmv2.Snssai{
					SD:  []byte{0x01, 0x02, 0x03},
					SSt: []byte{0x01},
				},
				FiveQi: &e2smkpmv2.FiveQi{
					Value: 23,
				},
				QFi: &e2smkpmv2.Qfi{
					Value: 52,
				},
				QCi: &e2smkpmv2.Qci{
					Value: 24,
				},
				QCimax: &e2smkpmv2.Qci{
					Value: 30,
				},
				QCimin: &e2smkpmv2.Qci{
					Value: 1,
				},
				ARpmax: &e2smkpmv2.Arp{
					Value: 15,
				},
				ARpmin: &e2smkpmv2.Arp{
					Value: 1,
				},
				BitrateRange:     &br,
				LayerMuMimo:      &lmm,
				SUm:              &sum,
				DistBinX:         &dbx,
				DistBinY:         &dby,
				DistBinZ:         &dbz,
				PreLabelOverride: &plo,
				StartEndInd:      &seind,
			},
		},
	}
	measCondItem.MatchingCond.Value = append(measCondItem.MatchingCond.Value, mci2)

	//if err := measCondItem.Validate(); err != nil {
	//	return nil, err
	//}
	return measCondItem, nil
}

func Test_perEncodeMeasurementCondItem(t *testing.T) {

	mci, err := createMeasurementCondItem()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mci, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementCondItem PER\n%v", hex.Dump(per))

	result := e2smkpmv2.MeasurementCondItem{}
	err = aper.UnmarshalWithParams(per, &result, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	//assert.Assert(t, &result != nil)
	t.Logf("MeasurementCondItem PER - decoded\n%v", &result)
	assert.Equal(t, mci.GetMeasType().GetMeasId().GetValue(), result.GetMeasType().GetMeasId().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestValue().GetValueInt(), result.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestValue().GetValueInt())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestType().GetAMbr().Number(), result.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestType().GetAMbr().Number())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestExpr().Number(), result.GetMatchingCond().GetValue()[0].GetTestCondInfo().GetTestExpr().Number())
	assert.DeepEqual(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetPlmnId().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetPlmnId().GetValue())
	assert.DeepEqual(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSliceId().GetSD(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSliceId().GetSD())
	assert.DeepEqual(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSliceId().GetSSt(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSliceId().GetSSt())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetFiveQi().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetFiveQi().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQFi().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQFi().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCi().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCi().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCimax().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCimax().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCimin().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetQCimin().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetARpmax().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetARpmax().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetARpmin().GetValue(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetARpmin().GetValue())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetBitrateRange(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetBitrateRange())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetLayerMuMimo(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetLayerMuMimo())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSUm().Number(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetSUm().Number())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinX(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinX())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinY(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinY())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinZ(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetDistBinZ())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetPreLabelOverride().Number(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetPreLabelOverride().Number())
	assert.Equal(t, mci.GetMatchingCond().GetValue()[1].GetMeasLabel().GetStartEndInd().Number(), result.GetMatchingCond().GetValue()[1].GetMeasLabel().GetStartEndInd().Number())
}

func Test_perMeasurementCondItemCompareBytes(t *testing.T) {

	mci, err := createMeasurementCondItem()
	assert.NilError(t, err)

	//aper.ChoiceMap = e2smkpmv2.Choicemape2smKpm
	per, err := aper.MarshalWithParams(mci, "valueExt", e2smkpmv2.Choicemape2smKpm, nil)
	assert.NilError(t, err)
	t.Logf("MeasurementCondItem PER\n%v", hex.Dump(per))

	//Comparing with reference bytes
	perRefBytes, err := hexlib.DumpToByte(refPerMeasCI)
	assert.NilError(t, err)
	assert.DeepEqual(t, per, perRefBytes)
}
