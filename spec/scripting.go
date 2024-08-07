// https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/Scripting.ts
/*
 * Licensed to Elasticsearch B.V. under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch B.V. licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package spec

import "github.com/ngicks/und/sliceund"

type ScriptLanguage string

const (
	Painless   ScriptLanguage = "painless"
	Expression ScriptLanguage = "expression"
	Mustache   ScriptLanguage = "mustache"
	Java       ScriptLanguage = "java"
)

type StoredScript struct {
	Lang    ScriptLanguage                  `json:"lang"`
	Options sliceund.Und[map[string]string] `json:"options,omitempty"`
	Source  string                          `json:"source"`
}

type ScriptBase struct {
	Params sliceund.Und[map[string]any] `json:"params,omitempty"`
}

/** @shortcut_property source */
type InlineScript struct {
	ScriptBase
	Lang    sliceund.Und[ScriptLanguage]    `json:"lang,omitempty"`
	Options sliceund.Und[map[string]string] `json:"options,omitempty"`
	Source  string                          `json:"source"`
}

type StoredScriptId struct {
	ScriptBase
	Id string `json:"id"`
}

// Script = InlineScript | StoredScriptId
type Script struct {
	ScriptBase
	Lang    sliceund.Und[ScriptLanguage]    `json:"lang,omitempty"`
	Options sliceund.Und[map[string]string] `json:"options,omitempty"`
	Source  sliceund.Und[string]            `json:"source,omitempty"`
	Id      sliceund.Und[string]            `json:"id,omitempty"`
}

func (s Script) IsInlineScript() bool {
	return s.Id.IsDefined()
}

func (s Script) IsStoredScriptId() bool {
	return s.Source.IsDefined()
}

func (s Script) InlineScript() InlineScript {
	return InlineScript{
		ScriptBase: s.ScriptBase,
		Lang:       s.Lang,
		Options:    s.Options,
		Source:     s.Source.Value(),
	}
}

func (s Script) StoredScriptId() StoredScriptId {
	return StoredScriptId{
		ScriptBase: s.ScriptBase,
		Id:         s.Id.Value(),
	}
}

type ScriptField struct {
	Script        Script             `json:"script"`
	IgnoreFailure sliceund.Und[bool] `json:"ignore_failure,omitempty"`
}
