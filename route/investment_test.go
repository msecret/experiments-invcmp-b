package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RouteInvestmentTests struct {
	suite.Suite
}

func (suite *RouteInvestmentTests) SetupTest() {

}

func (suite *RouteInvestmentTests) TestCreateOneWithInput() {
}

func TestRouteInvestmentTests(t *testing.T) {
	suite.Run(t, new(RouteInvestmentTests))

}
