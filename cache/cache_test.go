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
    cacheHandler.Cache.Del("test1")
    cacheHandler.Cache.Del("test2")
	cacheHandler.Cache.Quit()
}

func (suite *CacheSuite) TestRegister() {
	require.NotNil(suite.T(), cacheHandler.Cache)

	version, err := cacheHandler.Cache.Version()
	require.NoError(suite.T(), err)
	server := map[string]string{"localhost:11211": "1.6.18"}
	require.Equal(suite.T(), server, version)
}

func (suite *CacheSuite) TestMemory() {
	cacheHandler.RememberMe("test1", "username")
	name := cacheHandler.WhatsMyName("test1")
	require.Equal(suite.T(), name, "username")
}

func (suite *CacheSuite) TestUnknownID() {
	name := cacheHandler.WhatsMyName("test2")
	require.Equal(suite.T(), "test2", name)
}

// 'go test' starts here
func TestCache(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip integration test for cache")
	}
	suite.Run(t, new(CacheSuite))
}
