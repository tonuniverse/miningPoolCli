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

package config

import (
	"runtime"
	"time"
)

type colors struct {
	ColorRed, ColorGreen,
	ColorYellow, ColorBlue,
	ColorReset string
}

type texts struct {
	Logo, GlobalHelpText, AuthKeyHelp, PoolUrlHelp, WelcomeAdditionalMsg, AuthKeyFlagError, GPL3 string
}

type ubuntuMiner struct {
	ReleaseURL, FileName, ExecutableName string
}

type winMiner struct {
	ReleaseURL, FileName, ExecutableName string
}

type minerGetter struct {
	MinerDirectory string
	UbuntuSettings ubuntuMiner
	WinSettings    winMiner
}

type os struct {
	OperatingSystem, Architecture string
}

type serverSettings struct {
	MiningPoolServerURL, AuthKey string
}

type staticBeforeMinerSettings struct {
	NumCPUForWFlag int
	BoostFactor    int
	TimeoutT       int
	Iterations     string
	PoolAddress    string
}

type osType struct {
	Linux, Win, Macos string
}

var Colors colors
var Texts texts
var MinerGetter minerGetter
var OS os
var ServerSettings serverSettings
var StaticBeforeMinerSettings staticBeforeMinerSettings
var OSType osType

var UpdateStatsFile bool
var StartProgramTimestamp int

func Configure() {
	StartProgramTimestamp = int(time.Now().Unix())

	Colors = colors{
		"\u001b[31m", "\u001b[32m",
		"\u001b[33m", "\u001b[34m",
		"\u001b[0m",
	}

	// -------- OS Types
	OSType.Linux = "linux"
	OSType.Win = "windows"
	OSType.Macos = "darwin"
	// --------

	MinerGetter.MinerDirectory = "__miner__"

	// -------- set Release for Ubuntu
	MinerGetter.UbuntuSettings.FileName = "miner-opencl-ubuntu-20.04-x86-64.tar.gz"
	MinerGetter.UbuntuSettings.ReleaseURL = "https://github.com/tonuniverse/pow-miner-gpu/releases/download/v0.0.1/" +
		MinerGetter.UbuntuSettings.FileName
	MinerGetter.UbuntuSettings.ExecutableName = "pow-miner-opencl"
	// --------

	// -------- set Release for Win
	MinerGetter.WinSettings.FileName = "..."
	MinerGetter.WinSettings.ReleaseURL = "..." +
		MinerGetter.WinSettings.FileName
	MinerGetter.WinSettings.ExecutableName = "pow-miner-opencl.exe"
	// --------

	// -------- StaticBeforeMinerSettings
	StaticBeforeMinerSettings.NumCPUForWFlag = runtime.NumCPU()
	StaticBeforeMinerSettings.BoostFactor = 256
	StaticBeforeMinerSettings.Iterations = "100000000000"
	StaticBeforeMinerSettings.TimeoutT = 256
	// --------

	// BuildVersion = "1.0.1"

	Texts.Logo = `
 _                                _                               
| |                              (_)                              
| |_   ___   _ __   _   _  _ __   _ __   __  ___  _ __  ___   ___ 
| __| / _ \ | '_ \ | | | || '_ \ | |\ \ / / / _ \| '__|/ __| / _ \
| |_ | (_) || | | || |_| || | | || | \ V / |  __/| |   \__ \|  __/
 \__| \___/ |_| |_| \__,_||_| |_||_|  \_/   \___||_|   |___/ \___|
	`

	Texts.GPL3 = `miningPoolCli (v` + BuildVersion + `) – open-source tonuniverse mining pool client

Copyright (C) 2021 tonuniverse.com

At this time, the authors can be contacted by this email:
contact@tonuniverse.com

This program comes with ABSOLUTELY NO WARRANTY; for details read LICENSE.
This is free software, and you are welcome to redistribute it
under certain conditions; read LICENSE for details.`

	Texts.WelcomeAdditionalMsg = "- - - - - - - - - - - - - - - - - - - -\n" + Texts.GPL3 +
		"\n\nofficial website: tonuniverse.com \n" +
		"source code: github.com/tonuniverse/miningPoolCli \n" +
		"- - - - - - - - - - - - - - - - - - - -\n"

	Texts.GlobalHelpText = `
Usage of ./miningPoolCli (Read more at tonuniverse.com):
  ---------------------------------------------------

  -pool-id string

	Example: -pool-id=904f935185ef96c1ab4daf11e5d84b22
	A unique identifier of a pool participant.

  ---------------------------------------------------

  -url string
  
	Example: -url=https://pool.tonuniverse.com
	Mining pool API url. (default "https://pool.tonuniverse.com")

  ---------------------------------------------------

  -stats bool
  
	Example: -stats
	If this flag is set, a "stats.json" file will be created 
	with automatically updated statistics.

  ---------------------------------------------------
`
}
