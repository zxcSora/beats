// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "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"
}
