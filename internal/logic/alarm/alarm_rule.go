package alarm

import (
	"sviwo/internal/service"
)

type sAlarmRule struct{}

func init() {
	service.RegisterAlarmRule(alarmRuleNew())
}

func alarmRuleNew() *sAlarmRule {
	return &sAlarmRule{}
}
