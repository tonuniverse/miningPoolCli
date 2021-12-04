## miningPoolCli (beta version)

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

  -auth-key string

	Example: -auth-key=gwhUUnLp0F5YZu6qanJHRl3SzoTrBq1
	Key for authorization in the mining pool.

  ---------------------------------------------------

  -url string

	Example: -url=http://192.0.0.1:8000
	Mining pool API url. (default "https://pool.tonuniverse.com")

  ---------------------------------------------------