package main

import (
    "fmt"
    "os"

    "github.com/lightningnetwork/lnd/channeldb"
    "github.com/lightningnetwork/lnd/lnwire"
)

func fatal(err error) {
    fmt.Printf("fatal err: %v", err)
    os.Exit(1)
}

func recoverDB() error {
    db, err := channeldb.Open(".")
    if err != nil {
        return err
    }
    defer db.Close()

    channels, err := db.FetchAllOpenChannels()
    if err != nil {
        return err
    }

    var localTotal lnwire.MilliSatoshi
    for _, channel := range channels {
        fmt.Printf("chanpoint: %v\n", channel.FundingOutpoint)
        localTotal += channel.LocalCommitment.LocalBalance
    }
    fmt.Printf("local_balance: %v\n", localTotal.ToSatoshis())

    return nil
}

func main() {
    err := recoverDB()
    if err != nil {
        fatal(err)
    }
    os.Exit(0)
}
