/*
Copyright 2014 Rohith All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package endpoints

import (
	"github.com/gambol99/embassy/proxy/services"
)

type EndpointsProvider interface {
	/* get a list of the endpoints from the backend */
	List(*services.Service) ([]Endpoint, error)
	/* watch for changes on the backend */
	Watch(*services.Service) (EndpointEventChannel, error)
	/* shutdown and clean up the provider */
	Close()
}
