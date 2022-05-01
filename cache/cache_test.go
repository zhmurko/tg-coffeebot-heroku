package cache_test

import (
	"testing"

	_ "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	// "github.com/memcachier/mc/v3"
	cacheHandler "github.com/zhmurko/tg-coffeebot-heroku/cache"
)

type CacheSuite struct {
	suite.Suite
	// Cache *mc.Client
}

func (suite *CacheSuite) SetupSuite() {
	cacheHandler.Cache = cacheHandler.Register()
}

func (suite *CacheSuite) TearDownSuite() {
	cacheHandler.Cache.Quit()
}

func (suite *CacheSuite) TestRegister() {
	require.NotNil(suite.T(), cacheHandler.Cache)

	version, err := cacheHandler.Cache.Version()
	require.NoError(suite.T(), err)
	server := map[string]string{"localhost:11211": "1.6.15"}
	require.Equal(suite.T(), version, server)
}

func (suite *CacheSuite) TestMemory() {
	cacheHandler.RememberMe("1", "username")
	name := cacheHandler.WhatsMyName("1")
	require.Equal(suite.T(), name, "username")
}

func (suite *CacheSuite) TestUnknownID() {
	name := cacheHandler.WhatsMyName("newid")
	require.Equal(suite.T(), "newid", name)
}

// 'go test' starts here
func TestCache(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip integration test for cache")
	}
	suite.Run(t, new(CacheSuite))
}
