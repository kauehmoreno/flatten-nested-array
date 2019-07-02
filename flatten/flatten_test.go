package flatten

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testFlattenSuiteCase struct {
	suite.Suite
}

func TestFlattenSuite(t *testing.T) {
	suite.Run(t, new(testFlattenSuiteCase))
}

func (suite testFlattenSuiteCase) TestIsNumShouldReturnOk() {
	data, ok := isANum(10)
	suite.Require().True(ok, "should return true to num")
	suite.Require().Equal(int64(10), data, "should be 10 as result")
}

func (suite testFlattenSuiteCase) TestIsNumShouldReturnOkForInt64Num() {
	data, ok := isANum(int64(10))
	suite.Require().True(ok, "should return true to num")
	suite.Require().Equal(int64(10), data, "should be 10 as result")
}

func (suite testFlattenSuiteCase) TestIsNumShouldReturnOkForInt32Num() {
	data, ok := isANum(int32(10))
	suite.Require().True(ok, "should return true to num")
	suite.Require().Equal(int64(10), data, "should be 10 as result")
}

func (suite testFlattenSuiteCase) TestIsNumShouldReturnOkForInt16Num() {
	data, ok := isANum(int16(10))
	suite.Require().True(ok, "should return true to num")
	suite.Require().Equal(int64(10), data, "should be 10 as result")
}

func (suite testFlattenSuiteCase) TestIsNumShouldReturnOkForInt6Num() {
	data, ok := isANum(int8(12))
	suite.Require().True(ok, "should return true to num")
	suite.Require().Equal(int64(12), data, "should be 10 as result")
}

func (suite testFlattenSuiteCase) TestIsNumInCaseOFArrayOfIntShouldReturnFalse() {
	_, ok := isANum([]int64{10, 12, 31})
	suite.Require().False(ok, "should not return true for non number cases")
}

func (suite testFlattenSuiteCase) TestNestedArray() {
	lala := []interface{}{[]interface{}{1, 2, []int64{3}}, 4}
	result, err := Flatten(lala)
	suite.Require().NoError(err, "should fail on flatten array")
	suite.Require().Len(result, 4, "should contains 4 elements")
	expected := []int64{1, 2, 3, 4}
	suite.Require().Equal(expected, result, "elements are no equal")
}

func (suite testFlattenSuiteCase) TestExtractElemFromNestArray() {
	nested := []interface{}{[]interface{}{1, 2, []interface{}{[]int64{3, 4, 5}}}, 6}
	expected := []int64{1, 2, 3, 4, 5, 6}
	var result []int64
	err := extractElem(nested, &result)
	suite.Require().NoError(err, "should not fail over extract elem from nestd array")
	suite.Require().Equal(expected, result, "Array are not equal")
}

func (suite testFlattenSuiteCase) TestFlattenArray() {
	nested := []interface{}{[]interface{}{1, 2, []interface{}{[]int64{3, 4, 5}}}, 6}
	expected := []int64{1, 2, 3, 4, 5, 6}
	result, err := Flatten(nested)
	suite.Require().NoError(err, "should not fail over extract elem from nestd array")
	suite.Require().Equal(expected, result, "Array are not equal")
}
