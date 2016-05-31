package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/trillian/log"
)

var batchLimitFlag = flag.Int("batch_limit", 50, "Max number of leaves to process")

// This just runs a one shot sequencing operation. Use queue_leaves to prepare work to
// and then run this.
func main() {
	flag.Parse()

	treeId := getLogIdFromFlagsOrDie()
	storage := getStorageFromFlagsOrDie(treeId)

	sequencer := log.NewSequencer(storage)

	err := sequencer.SequenceBatch(*batchLimitFlag)

	if err != nil {
		panic(err)
	}
}
