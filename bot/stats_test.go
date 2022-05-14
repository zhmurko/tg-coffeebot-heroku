package bot

import (
	"strings"
	"testing"

	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	db "github.com/zhmurko/tg-coffeebot-heroku/db"
)

type StatsSuite struct {
	suite.Suite
	coffee [3]string
	orders []db.Order
}

func (suite *StatsSuite) SetupSuite() {
	suite.coffee[1] = "Latte"
	suite.coffee[2] = "Espresso"
	suite.orders = []db.Order{
		db.Order{1, "Latte"},
		db.Order{1, "Espresso"},
		db.Order{2, "Latte"},
		db.Order{2, "Latte"},
		db.Order{2, "Espresso"},
	}
}

func (suite *StatsSuite) TestStatsMessage() {
	var expected [2]string
	expected[0] = "Latte: 3"
	expected[1] = "Espresso: 2"
	expected_line := strings.Join(expected[:], "\n")
	result := stats(suite.orders)
	require.Equal(suite.T(), expected_line, result)
}

func TestStatsSuite(t *testing.T) {
	suite.Run(t, new(StatsSuite))
}
