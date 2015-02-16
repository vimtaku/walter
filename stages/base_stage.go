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
package stages

import (
	"container/list"

	"github.com/recruit-tech/walter/log"
)

type BaseStage struct {
	Runner
	InputCh     *chan Mediator
	OutputCh    *chan Mediator
	ChildStages list.List
	StageName   string `config:"stage_name"`
	Message   string `config:"message"`
}

func (b *BaseStage) Run() bool {
	if b.Runner == nil {
		panic("Mast have a child class assigned")
	}
	return b.Runner.Run()
}

func (b *BaseStage) AddChildStage(stage Stage) {
	log.Debugf("added childstage: %v", stage)
	b.ChildStages.PushBack(stage)
}

func (b *BaseStage) GetChildStages() list.List {
	return b.ChildStages
}

func (b *BaseStage) GetStageName() string {
	return b.StageName
}

func (b *BaseStage) GetMessage() string {
	return b.Message
}

func (b *BaseStage) SetStageName(stageName string) {
	b.StageName = stageName
}
func (b *BaseStage) SetMessage(message string) {
	b.Message = message
}

func (b *BaseStage) SetInputCh(inputCh *chan Mediator) {
	b.InputCh = inputCh
}

func (b *BaseStage) GetInputCh() *chan Mediator {
	return b.InputCh
}

func (b *BaseStage) SetOutputCh(outputCh *chan Mediator) {
	b.OutputCh = outputCh
}

func (b *BaseStage) GetOutputCh() *chan Mediator {
	return b.OutputCh
}
