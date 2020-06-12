// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package signout

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/exposure-notifications-verification-server/pkg/config"
	"github.com/google/exposure-notifications-verification-server/pkg/controller"
	"github.com/google/exposure-notifications-verification-server/pkg/database"
)

type signoutController struct {
	config  *config.Config
	db      database.Database
	session *controller.SessionHelper
}

// New creates a new signout controller. When run, clears the session cookie.
func New(config *config.Config, db database.Database, session *controller.SessionHelper) controller.Controller {
	return &signoutController{config, db, session}
}

func (soc *signoutController) Execute(c *gin.Context) {
	soc.session.DestroySession(c)
	m := controller.NewTemplateMap(soc.config)
	if reason := c.Query("reason"); reason != "" {
		m["reason"] = reason
	}
	m["firebase"] = soc.config.Firebase
	c.HTML(http.StatusOK, "signout", m)
}