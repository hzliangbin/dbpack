/*
 * Copyright 2022 CECTC, Inc.
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

package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

var applicationIDs = make([]string, 0)

func RegisterRoutes() (http.Handler, error) {
	router := mux.NewRouter().SkipClean(true).UseEncodedPath()
	// Add healthcheck router
	registerHealthCheckRouter(router)

	// Add server metrics router
	registerMetricsRouter(router)

	// Add status router
	registerStatusRouter(router)

	// Add branch session router
	registerBranchSessionsRouter(router)

	return router, nil
}

func AppendApplicationID(applicationID string) {
	applicationIDs = append(applicationIDs, applicationID)
}
