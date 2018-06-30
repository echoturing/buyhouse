package etc

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"path/filepath"
)

type ConfigTestSuit struct {
	suite.Suite
}

func (c *ConfigTestSuit) TestLoadConfigFromFile() {
	filePath, err := filepath.Abs("config.yaml")
	c.Nil(err)
	cfg, err := LoadConfigFromFile(filePath)
	c.Nil(err)
	c.Equal("root", cfg.Mysql.User)
	c.Equal("test", cfg.Mysql.Password)
}

func TestConnectionTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuit))
}
