# Recover funds off a LND node that has a corrupted channel.db.
See this issue for reference on the situation. https://github.com/lightningnetwork/lnd/issues/2825

## Inital Info & Disclaimer. 
* Use this repo **completely at your own risk**. Hopefully you will never have a courrpted channel.db, if you do please let me know of your success stories. This repo has proven effective for my situation, your milage may vary. 
* This is in no-way officially part of the LND repo. LND very well may have a robust recovery feature in the future. (This is neither official or robust).
* If you are working with a corrupted harddrive you will first need to run commands like `fsck` or similar to try and fix the corrupted drive. This has worked for me. 
* If you are using a lightning node / Raspberry Pi / etc / you will need to shutdown LND and retrieve the channel.db
  * I then recommend spinning up your LND instance as a recovery environment.
* Please create a branch for your unique situation and push if you find success outside of the code in `master`.
  * `go get -d github.com/miketwenty1/lightningrecovery`
  * `cd ~/gocode/src/github.com/miketwenty1/lightningrecovery`
  * `git checkout -b [Descriptive Branch]`

## When using this repo I recommend making a complete copy of your .lnd directory.
`mkdir ~/recoverybkp` <br>
`cp -R ~/.lnd ~/recoverybkp` <br>

## Check corrupted channel.db with channelcheck
First thing we will want to do is see if we can parse the `channel.db` and find funds.<br>
`cp ~/.lnd/data/graph/mainnet/channel.db ~/gocode/src/github.com/miketwenty1/lightningrecovery/channelbackup/`<br>
_Path to your channel.db may vary based on your setup_<br>
`cd ~/gocode/src/github.com/miketwenty1/lightningrecovery/channelbackup/`<br>
`ls` make sure your `channel.db` is in this directory
`./main`<br>
Hopefully you will see Channel information along with a summation totaled at the bottom. <br>
You should only expect to recover this amount of less.<br>
_If you don't see any channel or amount..._ (I'm sorry, but your situation is much different than mine) - Feel free to continue at your own risk. 

## LND Conf
`~/.lnd/lnd.conf`
```
[Application Options]
nolisten=1
nobootstrap=1
debuglevel=CNCT=debug,SWPR=debug,UTXN=debug

[Bitcoin]
bitcoin.mainnet=1
bitcoin.active=1
...
...
...
```
It's recommended to enable nolisten and nobootstrap options, it might also help to enable extra debugging.

## Running LND
When you run lnd, hopefully you will be able to close channels, see pending channels, see sweep transactions. You might run into a situation where pending channels do not get earsed. This is ok, as channels might still be closed.<br>
 `lncli closedchannel -h` you will want to close all channels.<br>
 In my case I needed to force close all channels then wait 2 weeks before continuing. I then needed to start LND a couple times and retrieve some blocks. Shutdown LND, restart, retrieve a few blocks. I eventually swept all the funds that were locked up. <br>
Example command to send coins off.<br>
`lncli sendcoins --addr=[BTC address] --sweepall`<br>
Hopefully you never are in this situation, but if you are.. hopefully you see some numbers in `lncli walletbalance`. 
Good luck!

