/* walter: a deployment pipeline template
 * Copyright (C) 2014 Recruit Technologies Co., Ltd. and contributors
 * (see CONTRIBUTORS.md)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package messengers

import (
	"fmt"
)

type Messenger interface {
	Post(string) bool
}

func InitMessenger(mtype string) (Messenger, error) {
	var messenger Messenger
	switch mtype {
	case "hipchat":
		messenger = new(HipChat)
	case "hipchat2":
		messenger = new(HipChat2)
	case "slack":
		messenger = new(Slack)
	case "fake":
		messenger = new(FakeMessenger)
	default:
		err := fmt.Errorf("no messenger type: %s", mtype)
		return nil, err
	}
	return messenger, nil
}
