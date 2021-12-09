## miningPoolCli

Open-source mining pool client

To use with tonuniverse mining pool, follow the instructions at https://tonuniverse.com

The text of the license can be obtained in the file "LICENSE".

## Source code

You can always get the source code from the github repository page:
https://github.com/tonuniverse/miningPoolCli/

## Build

```
go build -o miningPoolCli main.go
```

## Usage

! When the software starts up, it downloads the miner executable 
file from the release of the given github repository: 
https://github.com/tonuniverse/pow-miner-gpu

Run `./miningPoolCli` with flags:

  ---------------------------------------------------

  -pool-id string

	Example: -pool-id=904f935185ef96c1ab4daf11e5d84b22
	Key for authorization in the mining pool.

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