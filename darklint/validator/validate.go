package validator

import (
	"github.com/darklab8/fl-configs/configs/configs_export"
	"github.com/darklab8/fl-configs/configs/configs_mapped"
)

type Validator struct {
}

func NewValidator(configs *configs_mapped.MappedConfigs, export *configs_export.Exporter) *Validator {
	v := &Validator{}

	return v
}

func (v *Validator) Run() {

}
