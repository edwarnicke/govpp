// Copyright (c) 2020 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"time"

	"github.com/edwarnicke/govpp/binapi/vpe"
	"github.com/edwarnicke/log"

	"github.com/edwarnicke/vpphelper"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Connect to VPP with a 1 second timeout
	connectCtx, _ := context.WithTimeout(ctx, time.Second)
	conn, vppErrCh := vpphelper.StartAndDialContext(connectCtx)
	exitOnErrCh(ctx, cancel, vppErrCh)

	// Create a RPC client for the vpe api
	c := vpe.NewServiceClient(conn)
	version, err := c.ShowVersion(ctx, &vpe.ShowVersion{})
	if err != nil {
		log.Entry(ctx).Fatalln("ERROR: ShowVersion failed:", err)
	}
	log.Entry(ctx).Infof("VPP Version: %v\nCompile Date: %v\nCompile Location: %v", version.Version, version.BuildDate, version.BuildDirectory)

	// Cancel the context governing vpp's lifecycle and wait for it to exit
	cancel()
	<-vppErrCh
}

func exitOnErrCh(ctx context.Context, cancel context.CancelFunc, errCh <-chan error) {
	// If we already have an error, log it and exit
	select {
	case err := <-errCh:
		log.Entry(ctx).Fatal(err)
	default:
	}
	// Otherwise wait for an error in the background to log and cancel
	go func(ctx context.Context, errCh <-chan error) {
		err := <-errCh
		log.Entry(ctx).Error(err)
		cancel()
	}(ctx, errCh)
}
