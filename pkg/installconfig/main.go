package installconfig

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
)

// This should come from openshift-installer but because i'm a go noob
// redefine this here
// the installer uses pkg/assets for parsing and storing asses. look there
// struct Networking struct
// struct InstallConfig struct {
// 	ApiVersion string `json:"apiVersion"`
// 	BaseDomain string `json:"baseDomain"`
// }

const installConfigFile = "install-config.yaml"

func Load(dir string) (err error) {
	var ic map[string]interface{}
	y, err := ioutil.ReadFile(filepath.Join(dir, installConfigFile))

	err = yaml.Unmarshal(y, &ic)
	if err != nil {
		logrus.Errorf("bla %v", err)
	}
	logrus.Errorf("bla %v", err)
	fmt.Printf("%v\n", ic["compute"])

	return err
}
