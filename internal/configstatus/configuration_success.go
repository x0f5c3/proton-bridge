// Copyright (c) 2023 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package configstatus

import (
	"strconv"
)

type ConfigSuccessValues struct {
	Duration int `json:"duration"`
}

type ConfigSuccessDimensions struct {
	Autoconf    string `json:"autoconf"`
	ReportClick string `json:"report_click"`
	ReportSent  string `json:"report_sent"`
	ClickedLink string `json:"clicked_link"`
}

type ConfigSuccessData struct {
	MeasurementGroup string
	Event            string
	Values           ConfigSuccessValues
	Dimensions       ConfigSuccessDimensions
}

type ConfigSuccessBuilder struct{}

func (*ConfigSuccessBuilder) New(config *ConfigurationStatus) ConfigSuccessData {
	config.DataLock.RLock()
	defer config.DataLock.RUnlock()

	return ConfigSuccessData{
		MeasurementGroup: "bridge.any.configuration",
		Event:            "bridge_config_success",
		Values: ConfigSuccessValues{
			Duration: config.isPendingSinceMin(),
		},
		Dimensions: ConfigSuccessDimensions{
			Autoconf:    config.Data.DataV1.Autoconf,
			ReportClick: strconv.FormatBool(config.Data.DataV1.ReportClick),
			ReportSent:  strconv.FormatBool(config.Data.DataV1.ReportSent),
			ClickedLink: config.Data.clickedLinkToString(),
		},
	}
}
