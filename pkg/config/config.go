package config

import (
	"github.com/golang/glog"
	toml "github.com/pelletier/go-toml"

	"github.com/ReflecBeatCustom/haereticus/pkg/types"
)

// NewServerConfig ...
func NewServerConfig(file string) (*types.ServerConfig, error) {
	tomlFile, err := toml.LoadFile(file)
	if err != nil {
		glog.Errorf("Load toml config error: %v", err)
		return nil, err
	}
	conf := types.ServerConfig{}
	err = tomlFile.Unmarshal(&conf)
	if err != nil {
		glog.Errorf("Unmarshal config error: %v", err)
		return nil, err
	}
	return &conf, nil
}
