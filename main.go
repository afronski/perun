// Copyright 2017 Appliscale
//
// Maintainers and contributors are listed in README file inside repository.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// A tool for CloudFormation template validation and conversion.
package main

import (
	"os"

	"github.com/Appliscale/perun/cliparser"
	"github.com/Appliscale/perun/configuration"
	"github.com/Appliscale/perun/configurator"
	"github.com/Appliscale/perun/context"
	"github.com/Appliscale/perun/converter"
	"github.com/Appliscale/perun/mysession"
	"github.com/Appliscale/perun/offlinevalidator"
	"github.com/Appliscale/perun/onlinevalidator"
	"github.com/Appliscale/perun/parameters"
	"github.com/Appliscale/perun/progress"
	"github.com/Appliscale/perun/stack"
	"github.com/Appliscale/perun/utilities"
)

func main() {
	context, err := context.GetContext(cliparser.ParseCliArguments, configuration.GetConfiguration, configuration.ReadInconsistencyConfiguration)
	if err != nil {
		os.Exit(1)
	}

	if *context.CliArguments.Mode == cliparser.ValidateMode {
		utilities.CheckFlagAndExit(onlinevalidator.ValidateAndEstimateCosts(&context))

	}

	if *context.CliArguments.Mode == cliparser.ConvertMode {
		utilities.CheckErrorCodeAndExit(converter.Convert(&context))

	}

	if *context.CliArguments.Mode == cliparser.OfflineValidateMode {
		utilities.CheckFlagAndExit(offlinevalidator.Validate(&context))

	}

	if *context.CliArguments.Mode == cliparser.ConfigureMode {
		configurator.FileName(&context)
		os.Exit(0)
	}

	if *context.CliArguments.Mode == cliparser.CreateStackMode {
		utilities.CheckErrorCodeAndExit(stack.NewStack(&context))

	}

	if *context.CliArguments.Mode == cliparser.DestroyStackMode {
		utilities.CheckErrorCodeAndExit(stack.DestroyStack(&context))

	}

	if *context.CliArguments.Mode == cliparser.MfaMode {
		err := mysession.UpdateSessionToken(context.Config.DefaultProfile, context.Config.DefaultRegion, context.Config.DefaultDurationForMFA, &context)
		if err == nil {
			os.Exit(0)
		} else {
			context.Logger.Error(err.Error())
			os.Exit(1)
		}
	}

	if *context.CliArguments.Mode == cliparser.UpdateStackMode {
		utilities.CheckErrorCodeAndExit(stack.UpdateStack(&context))
	}

	if *context.CliArguments.Mode == cliparser.SetupSinkMode {
		progress.ConfigureRemoteSink(&context)
		os.Exit(0)
	}

	if *context.CliArguments.Mode == cliparser.DestroySinkMode {
		progress.DestroyRemoteSink(&context)
		os.Exit(0)
	}

	if *context.CliArguments.Mode == cliparser.CreateParametersMode {
		parameters.ConfigureParameters(&context)
		os.Exit(0)
	}

	if *context.CliArguments.Mode == cliparser.SetStackPolicyMode {
		if *context.CliArguments.DisableStackTermination || *context.CliArguments.EnableStackTermination {
			utilities.CheckErrorCodeAndExit(stack.SetTerminationProtection(&context))
		} else {
			utilities.CheckErrorCodeAndExit(stack.ApplyStackPolicy(&context))
		}
	}
}
