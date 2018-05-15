/*
 * Copyright 2017-2018 Dgraph Labs, Inc.
 *
 * This file is available under the Apache License, Version 2.0,
 * with the Commons Clause restriction.
 */

package bulk

import "expvar"

var (
	NumBadgerWrites     = expvar.NewInt("dgraph_bulk_loader_badger_writes_pending")
	NumReducers         = expvar.NewInt("dgraph_bulk_loader_num_reducers_total")
	NumQueuedReduceJobs = expvar.NewInt("dgraph_bulk_loader_reduce_queue_size")
)
