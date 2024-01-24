/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */

#ifndef NR_UE_SECURITY_CAPABILITY_SEEN
#define NR_UE_SECURITY_CAPABILITY_SEEN

#include <stdint.h>

#include "3gpp_24.301.h"

int encode_nr_ue_security_capability(
    nr_ue_security_capability_t* nruesecuritycapability, uint8_t iei,
    uint8_t* buffer, uint32_t len);

int decode_nr_ue_security_capability(
    nr_ue_security_capability_t* nruesecuritycapability, uint8_t iei,
    uint8_t* buffer, uint32_t len);

#endif /* NR UE SECURITY CAPABILITY_SEEN */
