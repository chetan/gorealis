/**
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

package realis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadClusters(t *testing.T) {

	clusters, err := LoadClusters("examples/clusters.json")
	if err != nil {
		fmt.Print(err)
	}

	assert.Equal(t, clusters[0].Name, "devcluster")
	assert.Equal(t, clusters[0].ZK, "192.168.33.7")
	assert.Equal(t, clusters[0].SchedZKPath, "/aurora/scheduler")
	assert.Equal(t, clusters[0].AuthMechanism, "UNAUTHENTICATED")
	assert.Equal(t, clusters[0].AgentRunDir, "latest")
	assert.Equal(t, clusters[0].AgentRoot, "/var/lib/mesos")
}
