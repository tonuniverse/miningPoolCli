## miningPoolCli

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tonuniverse/miningPoolCli)
![GitHub](https://img.shields.io/github/license/tonuniverse/miningPoolCli)

Open-source mining pool client

To use with tonuniverse mining pool, follow the instructions at https://tonuniverse.com

## Source code

You can always get the source code from the github repository page:
https://github.com/tonuniverse/miningPoolCli/

## Build

```
go build -o miningPoolCli main.go
```

## Usage

When the software starts up, it downloads the miner executable 
file from the release of the given github repository: 
https://github.com/tontechio/pow-miner-gpu/

Run `./miningPoolCli` with flags:

`-pool-id` string

	Example: -pool-id=904f935185ef96c1ab4daf11e5d84b22
	A unique identifier of a pool participant.

`-url` string
  
	Mining pool API url. (default "https://pool.tonuniverse.com")

`-stats` bool
  
	If this flag is set, a "stats.json" file will be created 
	with automatically updated statistics. (Hive OS support)

`-serve-stat` bool

	If this flag is set, the local server serving "/stat" is started. 
	Accepts GET and POST methods. Returns the miner's statistics in 
	JSON format. The HTTP port is automatically selected and will be 
	printed in the terminal and written to the "serveraddr.txt" file

`-handle-kill` bool

	Allows server to process HTTP requests to "/kill" to 
	force the client to terminate. The flag is applied 
	only if -serve-stat was previously specified

## Do release

To generate a new release, use `do-release.sh`.

This script will automatically compile the source code for the required OS, set the version from the `config/version.go` and create an archive

Usage: 

	./do-release.sh {linux|windows}

## LICENSE

GPL-3.0 License

The text of the license can be obtained in the file "LICENSE"

