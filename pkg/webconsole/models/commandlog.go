// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"context"
	"time"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/sqlchemy"

	api "yunion.io/x/onecloud/pkg/apis/webconsole"
	"yunion.io/x/onecloud/pkg/cloudcommon/consts"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/mcclient"
)

var commandLogManager *SCommandLogManager

type CommandType string

const (
	CommandTypeSSH = "ssh"
)

func InitCommandLog() {
	commandLogManager = GetCommandLogManager()
}

func GetCommandLogManager() *SCommandLogManager {
	if commandLogManager != nil {
		return commandLogManager
	}
	commandLogManager = &SCommandLogManager{
		SOpsLogManager: db.SOpsLogManager{
			SModelBaseManager: db.NewModelBaseManagerWithSplitable(
				SCommandLog{},
				"command_log_tbl",
				"commandlog",
				"commandlogs",
				"id",
				"start_time",
				consts.SplitableMaxDuration(),
				consts.SplitableMaxKeepMonths(),
			),
		},
	}
	commandLogManager.SetVirtualObject(commandLogManager)
	return commandLogManager
}

type SCommandLogManager struct {
	db.SOpsLogManager
}

type SCommandLog struct {
	db.SOpsLog

	SessionId  string      `width:"128" charset:"ascii" list:"user"`
	AccessedAt time.Time   `nullable:"false" list:"user" create:"required"`
	Type       string      `width:"32" charset:"utf8" nullable:"true" list:"user" create:"required"`
	LoginUser  string      `charset:"utf8" list:"user" create:"required"`
	StartTime  time.Time   `list:"user" create:"required"`
	Ps1        string      `charset:"utf8" list:"user" create:"optional" json:"ps1"`
	Command    CommandType `charset:"utf8" list:"user" create:"required"`
}

type CommandLogCreateInput struct {
	ObjId           string
	ObjName         string
	ObjType         string
	Action          string
	UserId          string
	User            string
	TenantId        string
	Tenant          string
	DomainId        string
	Domain          string
	ProjectDomainId string
	ProjectDomain   string
	Roles           string
	SessionId       string
	AccessedAt      time.Time
	Type            CommandType
	LoginUser       string
	StartTime       time.Time
	Ps1             string `json:"ps1"`
	Command         string
	Notes           jsonutils.JSONObject
}

func (m *SCommandLogManager) Create(ctx context.Context, userCred mcclient.TokenCredential, input *CommandLogCreateInput) (*SCommandLog, error) {
	data := jsonutils.Marshal(input)
	obj, err := db.DoCreate(GetCommandLogManager(), ctx, userCred, jsonutils.NewDict(), data, userCred)
	if err != nil {
		return nil, errors.Wrap(err, "Create CommandLog")
	}
	return obj.(*SCommandLog), nil
}

func (m *SCommandLogManager) ListItemFilter(
	ctx context.Context,
	q *sqlchemy.SQuery,
	userCred mcclient.TokenCredential,
	input api.CommandLogListInput,
) (*sqlchemy.SQuery, error) {
	q, err := m.SOpsLogManager.ListItemFilter(ctx, q, userCred, input.OpsLogListInput)
	if err != nil {
		return nil, errors.Wrap(err, "SOpsLogManager.ListItemFilter")
	}

	return q, nil
}
