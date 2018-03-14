// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package jobopts contains shared options for job submission. These options
// are exposed to allow user code to inspect and modify them.
package jobopts

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/apache/beam/sdks/go/pkg/beam/log"
)

var (
	// Endpoint is the job service endpoint.
	Endpoint = flag.String("endpoint", "", "Job service endpoint (required).")

	// JobName is the name of the job.
	JobName = flag.String("job_name", "", "Job name (optional).")

	// ContainerImage is the location of the SDK harness container image.
	ContainerImage = flag.String("container_image", "", "Container image")

	// Experiments toggle experimental features in the runner.
	Experiments = flag.String("experiments", "", "Comma-separated list of experiments (optional).")

	// Async determines whether to wait for job completion.
	Async = flag.Bool("async", false, "Do not wait for job completion.")

	// InternalJavaRunner is the java class needed at this time for Java runners.
	// To be removed.
	InternalJavaRunner = flag.String("internal_java_runner", "", "Internal java runner class.")
)

// GetEndpoint returns the endpoint, if non empty and exits otherwise. Runners
// such as Dataflow set a reasonable default. Convenience function.
func GetEndpoint() (string, error) {
	if *Endpoint == "" {
		return "", fmt.Errorf("no job service endpoint specified. Use --endpoint=<endpoint>")
	}
	return *Endpoint, nil
}

// GetJobName returns the specified job name or, if not present, an
// autogenerated name. Convenience function.
func GetJobName() string {
	if *JobName == "" {
		*JobName = fmt.Sprintf("go-job-%v", time.Now().UnixNano())
	}
	return *JobName
}

// GetContainerImage returns the specified SDK harness container image or,
// if not present, the default development container for the current user.
// Convenience function.
func GetContainerImage(ctx context.Context) string {
	if *ContainerImage == "" {
		*ContainerImage = os.ExpandEnv("$USER-docker-apache.bintray.io/beam/go:latest")
		log.Infof(ctx, "No container image specified. Using dev image: '%v'", *ContainerImage)
	}
	return *ContainerImage
}

// GetExperiments returns the experiments.
func GetExperiments() []string {
	if *Experiments == "" {
		return nil
	}
	return strings.Split(*Experiments, ",")
}
