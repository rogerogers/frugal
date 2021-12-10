/*
 * Copyright 2021 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package decoder

import (
    `github.com/cloudwego/frugal/internal/atm`
    `github.com/cloudwego/frugal/internal/utils`
)

type Linker interface {
    Link(p atm.Program) Decoder
}

var (
    linker   Linker
    F_decode atm.CallHandle
)

func init() {
    F_decode = atm.RegisterGCall(decode, emu_gcall_decode)
}

func Link(p atm.Program) Decoder {
    if linker == nil || utils.ForceEmulator {
        return link_emu(p)
    } else {
        return linker.Link(p)
    }
}

func SetLinker(v Linker) {
    linker = v
}
