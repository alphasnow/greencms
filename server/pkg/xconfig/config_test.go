// @author AlphaSnow

package xconfig

import (
	"os"
	"path"
	"testing"
)

func TestReadConfig(t *testing.T) {
	p, _ := os.Getwd()
	conf := path.Join(p, "stub/dev.yaml")

	cnf, _ := ReadConfig(conf)
	if cnf.GetString("app.name") != "viper-demo" {
		t.Errorf("want = %v, got %v", "viper-demo", cnf.GetString("app.name"))
		return
	}
	if cnf.GetBool("app.debug") != true {
		t.Errorf("want = %v, got %v", true, cnf.GetBool("app.debug"))
		return
	}
	if cnf.GetString("app.url") != "http://localhost" {
		t.Errorf("want = %v, got %v", "http://localhost", cnf.GetString("app.url"))
		return
	}
}

func TestReadConfigAndEnv(t *testing.T) {
	p, _ := os.Getwd()
	env := path.Join(p, "stub/.env")
	conf := path.Join(p, "stub/dev.yaml")

	cnf, _ := ReadConfigAndEnv(conf, env)
	if cnf.GetString("app.name") != "ViperDemo" {
		t.Errorf("want = %v, got %v", "ViperDemo", cnf.GetString("app.name"))
		return
	}
	if cnf.GetBool("app.debug") != false {
		t.Errorf("want = %v, got %v", false, cnf.GetBool("app.debug"))
		return
	}
	if cnf.GetString("app.url") != "http://localhost" {
		t.Errorf("want = %v, got %v", "http://localhost", cnf.GetString("app.url"))
		return
	}
}
