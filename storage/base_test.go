package storage

import (
	"testing"

	"github.com/huacnlee/gobackup/config"
	"github.com/longbridgeapp/assert"
)

func TestBase_newBase(t *testing.T) {
	model := config.ModelConfig{}
	storageConfig := config.SubConfig{}
	archivePath := "/tmp/gobackup/test-storeage/foo.zip"
	s := newBase(model, archivePath, storageConfig)

	assert.Equal(t, s.archivePath, archivePath)
	assert.Equal(t, s.model, model)
	assert.Equal(t, s.viper, model.Viper)
	assert.Equal(t, s.keep, 0)
}
