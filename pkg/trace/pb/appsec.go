// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package pb

// AppSecRule represents an AppSec rule.
type AppSecRule struct {
	ID   string            `msg:"id"`
	Name string            `msg:"name"`
	Tags map[string]string `msg:"tags"`
}

// AppSecRuleMatchParameter represents the data matched by an AppSec rule.
type AppSecRuleMatchParameter struct {
	Address   string        `msg:"address"`
	KeyPath   []interface{} `msg:"key_path"`
	Value     string        `msg:"value"`
	Highlight []string      `msg:"highlight"`
}

// AppSecRuleMatch represents an AppSec rule match.
type AppSecRuleMatch struct {
	Operator      string                     `msg:"operator"`
	OperatorValue string                     `msg:"operator_value"`
	Parameters    []AppSecRuleMatchParameter `msg:"parameters"`
}

// AppSecTrigger associates an AppSec rule and the data it matched.
type AppSecTrigger struct {
	Rule        AppSecRule        `msg:"rule"`
	RuleMatches []AppSecRuleMatch `msg:"rule_matches"`
}

// AppSecStruct is a container for AppSec data sent by the tracers.
type AppSecStruct struct {
	Triggers []AppSecTrigger `msg:"triggers"`
}
