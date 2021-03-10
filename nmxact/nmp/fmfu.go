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

package nmp

const SMP_PACKET_MTU = 256

type FmfuUploadReq struct {
	NmpBase `codec:"-"`
	Data    []byte `codec:"data"`
	Length  uint32 `codec:"len"`
	Offset  uint32 `codec:"off"`
	Sha     []byte `codec:"sha"`
	Address uint32 `codec:"addr"`
}

type FmfuUploadRsp struct {
	NmpBase `codec:"-"`
	Rc      int  `codec:"rc"`
	Offset  uint `codec:"off"`
}

func NewFmfuUploadReq() *FmfuUploadReq {
	r := &FmfuUploadReq{}
	fillNmpReq(r, NMP_OP_WRITE, NMP_GROUP_FMFU, NMP_ID_FMFU_UPLOAD)
	return r
}

func (r *FmfuUploadReq) Msg() *NmpMsg { return MsgFromReq(r) }

func NewFmfuUploadRsp() *FmfuUploadRsp {
	return &FmfuUploadRsp{}
}

func (r *FmfuUploadRsp) Msg() *NmpMsg { return MsgFromReq(r) }
