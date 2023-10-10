package yaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YAMLConfig struct {
	folder_location string
}

func (c *YAMLConfig) SetFolderLocation(folderpath string) error {
	c.folder_location = folderpath
	return nil
}

func GetApplicationsFromYAML(config YAMLConfig) (ApplicatonIndex, error) {
	var Applications ApplicatonIndex

	path := config.folder_location + "apps.yaml"

	// Open our yamlFile
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Applications, err
	}

	err = yaml.Unmarshal(data, &Applications)
	if err != nil {
		return Applications, err
	}

	return Applications, nil
}

func GetApplicationFromYAML(config YAMLConfig, applicationFolder string) (Appliciation, error) {
	var Application Appliciation

	path := config.folder_location + "apps/" + applicationFolder

	// Open our yamlFile
	data, err := ioutil.ReadFile(path + "/app_spec.yaml")
	if err != nil {
		return Application, err
	}

	err = yaml.Unmarshal(data, &Application)
	if err != nil {
		return Application, err
	}
	return Application, nil
}
