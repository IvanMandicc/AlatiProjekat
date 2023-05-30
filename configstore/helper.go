package configstore

import (
	"fmt"
)

const (
	allConfigs                = "config"
	allGroups                 = "group"
	config                    = "config/%s"
	configVersion             = "config/%s/%s"
	group                     = "group/%s"
	groupVersion              = "group/%s/%s"
	groupVersionConfigLabelId = "gconfig/%s/%s/%s/%s/%s"
	groupVersionConfigLabel   = "gconfig/%s/%s/%s/%s"
)

func generateConfigKey(id string, version string) string {
	if version != "" {
		return fmt.Sprintf(configVersion, id, version)
	} else {
		return fmt.Sprintf(config, id)
	}
}
func generateGroupKey(id string, version string, configId string, configLabels string) string {
	if configLabels != "" && configId != "" {
		return fmt.Sprintf(groupVersionConfigLabelId, id, version, "config", configLabels, configId)
	} else if configLabels != "" && configId == "" {
		return fmt.Sprintf(groupVersionConfigLabel, id, version, "config", configLabels)
	} else if configLabels == "" && configId == "" && version != "" {
		return fmt.Sprintf(groupVersion, id, version)
	} else {
		return fmt.Sprintf(group, id)
	}
}
