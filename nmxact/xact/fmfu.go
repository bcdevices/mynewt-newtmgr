/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package xact

import (
	"mynewt.apache.org/newtmgr/nmxact/nmp"
	"mynewt.apache.org/newtmgr/nmxact/sesn"
)

// const SMP_PACKET_MTU = 256

type FmfuUploadCmd struct {
	CmdBase
	Data    []byte
	Length  uint32
	Offset  uint32
	Sha     []byte
	Address uint32
}

func NewFmfuUploadCmd() *FmfuUploadCmd {
	return &FmfuUploadCmd{
		CmdBase: NewCmdBase(),
	}
}

type FmfuResult struct {
	Rsp *nmp.FmfuUploadRsp
}

func newFmfuUploadResult() *FmfuResult {
	return &FmfuResult{}
}

func (r *FmfuResult) Status() int {
	return r.Rsp.Rc
}

func (c *FmfuUploadCmd) Run(s sesn.Sesn) (Result, error) {
	r := nmp.NewFmfuUploadReq()
	r.Data = c.Data
	r.Length = c.Length
	r.Offset = c.Offset
	r.Sha = c.Sha
	r.Address = c.Address

	rsp, err := txReq(s, r.Msg(), &c.CmdBase)
	if err != nil {
		return nil, err
	}
	srsp := rsp.(*nmp.FmfuUploadRsp)

	res := newFmfuUploadResult()
	res.Rsp = srsp
	return res, nil
}
