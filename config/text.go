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

type texts struct {
	Logo, GlobalHelpText, AuthKeyHelp, PoolUrlHelp,
	WelcomeAdditionalMsg, AuthKeyFlagError, GPL3 string
}

var Texts texts

func configureTexts() {
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

	Texts.GlobalHelpText = `Usage of ./miningPoolCli (Read more at tonuniverse.com):

-pool-id string

	Example: -pool-id=904f935185ef96c1ab4daf11e5d84b22
	A unique identifier of a pool participant.

-url string
  
	Mining pool API url. (default "https://pool.tonuniverse.com")

-stats bool
  
	If this flag is set, a "stats.json" file will be created 
	with automatically updated statistics. (Hive OS support)

-serve-stat bool

	If this flag is set, the local server serving "/stat" is started. 
	Accepts GET and POST methods. Returns the miner's statistics in 
	JSON format. The HTTP port is automatically selected and will be 
	printed in the terminal and written to the "` + NetSrv.HostFileName + `" file
`
}
