package cfghelper

import "testing"

func TestConfigLoad(t *testing.T) {
	err := LoadConfig()

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("cfghelper: Default load config is ok")
}

func TestConfigLoadFromEnv(t *testing.T) {
	err := LoadConfig(Source("env"))

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("cfghelper: Load config from env is ok")
}

func TestConfigLoadFromFile(t *testing.T) {
	err := LoadConfig(Source("file:config/config.json"))

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("cfghelper: Load config from file is ok")
}

func TestConfigLoadFromEtcd(t *testing.T) {
	err := LoadConfig(Source("etcd:localhost:8500"))

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("cfghelper: Load config from etcd is ok")
}
