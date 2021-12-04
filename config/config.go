/*
miningPoolCli – open-source tonuniverse mining pool client

Copyright (C) 2021 Alexander Gapak
Copyright (C) 2021 Kirill Glushakov
Copyright (C) 2021 Roman Klimov

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

import "runtime"

type colors struct {
	ColorRed, ColorGreen,
	ColorYellow, ColorBlue,
	ColorReset string
}

type texts struct {
	Logo, GlobalHelpText, AuthKeyHelp, PoolUrlHelp, WelcomeAdditionalMsg, AuthKeyFlagError, GPL3 string
}
type minerGetter struct {
	MinerDirectory, UbuntuMinerRelUrl, UbubntuFileName string
}

type os struct {
	OperatingSystem, Architecture string
}

type serverSettings struct {
	MiningPoolServerURL, AuthKey string
}

type staticBeforeMinerSettings struct {
	NumCPUForWFlag   int
	PlatformID       int
	BoostFactor      int
	TimeoutT         int
	Iterations       string
	PoolAddress      string
	CheckActualInMin int
}

var Colors colors
var Texts texts
var MinerGetter minerGetter
var OS os
var ServerSettings serverSettings
var StaticBeforeMinerSettings staticBeforeMinerSettings

var MiningPoolServerURL string
var OperatingSystem string

func Configure() {
	Colors = colors{
		"\u001b[31m", "\u001b[32m",
		"\u001b[33m", "\u001b[34m",
		"\u001b[0m",
	}

	// ServerSettings.MiningPoolServerURL = ""

	MinerGetter.MinerDirectory = "__miner__"
	MinerGetter.UbubntuFileName = "miner-opencl-ubuntu-20.04-x86-64.tar.gz"
	MinerGetter.UbuntuMinerRelUrl = "https://github.com/tonuniverse/pow-miner-gpu/releases/download/v0.0.1/" +
		MinerGetter.UbubntuFileName

	// -------- StaticBeforeMinerSettings
	StaticBeforeMinerSettings.NumCPUForWFlag = runtime.NumCPU()
	StaticBeforeMinerSettings.PlatformID = 0
	StaticBeforeMinerSettings.BoostFactor = 32
	StaticBeforeMinerSettings.Iterations = "100000000000"
	StaticBeforeMinerSettings.TimeoutT = 256
	StaticBeforeMinerSettings.CheckActualInMin = 4 // How often to check the relevance of tasks in minutes
	// --------

	Texts.Logo = `
 _                                _                               
| |                              (_)                              
| |_   ___   _ __   _   _  _ __   _ __   __  ___  _ __  ___   ___ 
| __| / _ \ | '_ \ | | | || '_ \ | |\ \ / / / _ \| '__|/ __| / _ \
| |_ | (_) || | | || |_| || | | || | \ V / |  __/| |   \__ \|  __/
 \__| \___/ |_| |_| \__,_||_| |_||_|  \_/   \___||_|   |___/ \___|
	`

	Texts.GPL3 = `miningPoolCli (v0.0.1-alpha) – open-source tonuniverse mining pool client

Copyright (C) 2021 Alexander Gapak
Copyright (C) 2021 Kirill Glushakov
Copyright (C) 2021 Roman Klimov

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

  -auth-key string

	Example: -auth-key=gwhUUnLp0F5YZu6qanJHRl3SzoTrBq1
	Key for authorization in the mining pool.

  ---------------------------------------------------

  -url string
  
	Example: -url=http://192.0.0.1:8000
	Mining pool API url. (default "https://pool.tonuniverse.com")

  ---------------------------------------------------
`
}
