/*
miningPoolCli – open-source tonuniverse mining pool client

Copyright (C) 2021 tonuniverse.com

This file is part of miningPoolCli.

miningPoolCli is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

miningPoolCli is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with miningPoolCli.  If not, see <https://www.gnu.org/licenses/>.
*/

package mlog

import (
	"fmt"
	"miningPoolCli/config"
	"os"
	"time"

	"github.com/go-errors/errors"
)

func colorize(message string, color string) {
	switch config.OS.OperatingSystem {
	case config.OSType.Linux:
		fmt.Print(color, message, config.Colors.ColorReset+"\n")
	case config.OSType.Win:
		// windows not support unix colorize system
		fmt.Print(message + "\n")
	default:
		fmt.Print(message + "\n")
	}
}

func getNowTimeAsString() string {
	dt := time.Now()
	return "[" + dt.Format("15:04:05") + "]"
}

func LogOk(okMsg string) {
	colorize(getNowTimeAsString()+"[--&--] "+"OK: "+okMsg, config.Colors.ColorGreen)
}

func LogText(txtMsg string) {
	colorize(txtMsg, config.Colors.ColorBlue)
}

func LogInfo(infMsg string) {
	colorize(getNowTimeAsString()+"[--*--] "+"INFO: "+infMsg, config.Colors.ColorYellow)
}

func LogError(errMsg string) {
	colorize(getNowTimeAsString()+"[--!--] "+"ERROR: "+errMsg, config.Colors.ColorRed)
}

func LogFatal(errMsg string) {
	colorize(getNowTimeAsString()+"[--!--] "+"FATAL ERROR: "+errMsg, config.Colors.ColorRed)
	os.Exit(1)
}

func LogFatalStackError(err error) {
	colorize(
		getNowTimeAsString()+"[--!--] "+"FATAL ERROR: "+errors.Wrap(err, 1).ErrorStack(),
		config.Colors.ColorRed,
	)
	os.Exit(1)
}

func LogPass() {
	fmt.Println("")
}
