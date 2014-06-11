package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/msecret/invcmp-b/util/clock"
)

type BaseTests struct {
	suite.Suite
}

func (suite *BaseTests) SetupTest() {
	clock.NowForce(1)
}

func (suite *BaseTests) TestBaseUpdate() {
	expected := clock.Now()
	testBase := Base{}
	testBase.Update()

	assert.Equal(suite.T(), testBase.UpdatedAt, expected)
}

func (suite *BaseTests) TestBaseUpdate_ShouldNotUpdateCreatedAt() {
	expected := clock.Now()
	testBase := Base{CreatedAt: expected}
	clock.NowForce(2)

	testBase.Update()

	assert.Equal(suite.T(), testBase.CreatedAt, expected)
}

func (suite *BaseTests) TestBaseCreated() {
	expected := clock.Now()
	testBase := Base{}
	testBase.Create()

	assert.Equal(suite.T(), testBase.UpdatedAt, expected)
	assert.Equal(suite.T(), testBase.CreatedAt, expected)
}

func TestBaseTests(t *testing.T) {
	suite.Run(t, new(BaseTests))

}
