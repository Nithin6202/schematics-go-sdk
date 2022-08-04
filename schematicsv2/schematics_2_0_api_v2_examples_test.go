//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package schematics20apiv2_test

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

const externalConfigFile = "../schematics20_api_v2.env"

var (
	schematics20ApiService *schematics20apiv2.Schematics20ApiV2
	config                 map[string]string
	configLoaded           bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`Schematics20ApiV2 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(schematics20apiv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			schematics20ApiServiceOptions := &schematics20apiv2.Schematics20ApiV2Options{}

			schematics20ApiService, err = schematics20apiv2.NewSchematics20ApiV2UsingExternalConfig(schematics20ApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(schematics20ApiService).ToNot(BeNil())
		})
	})

	Describe(`Schematics20ApiV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSchematicsInfo request example`, func() {
			// begin-get_schematics_info

			getSchematicsInfoOptions := schematics20ApiService.NewGetSchematicsInfoOptions()

			schematicsInfo, response, err := schematics20ApiService.GetSchematicsInfo(getSchematicsInfoOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(schematicsInfo, "", "  ")
			fmt.Println(string(b))

			// end-get_schematics_info

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schematicsInfo).ToNot(BeNil())

		})
		It(`ListSchematicsAccessRules request example`, func() {
			// begin-list_schematics_access_rules

			listSchematicsAccessRulesOptions := schematics20ApiService.NewListSchematicsAccessRulesOptions()

			schematicsAccessRuleList, response, err := schematics20ApiService.ListSchematicsAccessRules(listSchematicsAccessRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(schematicsAccessRuleList, "", "  ")
			fmt.Println(string(b))

			// end-list_schematics_access_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schematicsAccessRuleList).ToNot(BeNil())

		})
		It(`GetSchematicsAccessRulesForResource request example`, func() {
			// begin-get_schematics_access_rules_for_resource

			getSchematicsAccessRulesForResourceOptions := schematics20ApiService.NewGetSchematicsAccessRulesForResourceOptions(
				"testString",
			)

			schematicsAccessRuleList, response, err := schematics20ApiService.GetSchematicsAccessRulesForResource(getSchematicsAccessRulesForResourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(schematicsAccessRuleList, "", "  ")
			fmt.Println(string(b))

			// end-get_schematics_access_rules_for_resource

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schematicsAccessRuleList).ToNot(BeNil())

		})
		It(`GetSchematicsCommandNames request example`, func() {
			// begin-get_schematics_command_names

			getSchematicsCommandNamesOptions := schematics20ApiService.NewGetSchematicsCommandNamesOptions(
				"testString",
			)

			commandList, response, err := schematics20ApiService.GetSchematicsCommandNames(getSchematicsCommandNamesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(commandList, "", "  ")
			fmt.Println(string(b))

			// end-get_schematics_command_names

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(commandList).ToNot(BeNil())

		})
		It(`ListWorkspaces request example`, func() {
			// begin-list_workspaces

			listWorkspacesOptions := schematics20ApiService.NewListWorkspacesOptions()

			workspaceList, response, err := schematics20ApiService.ListWorkspaces(listWorkspacesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspaceList, "", "  ")
			fmt.Println(string(b))

			// end-list_workspaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspaceList).ToNot(BeNil())

		})
		It(`CreateWorkspace request example`, func() {
			// begin-create_workspace

			workspaceModel := &schematics20apiv2.Workspace{
				Name:          core.StringPtr("My Workspace"),
				ResourceGroup: core.StringPtr("testString"),
				Location:      core.StringPtr("us-south"),
			}

			createWorkspaceOptions := schematics20ApiService.NewCreateWorkspaceOptions(
				workspaceModel,
			)

			workspace, response, err := schematics20ApiService.CreateWorkspace(createWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspace, "", "  ")
			fmt.Println(string(b))

			// end-create_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspace).ToNot(BeNil())

		})
		It(`GetWorkspace request example`, func() {
			// begin-get_workspace

			getWorkspaceOptions := schematics20ApiService.NewGetWorkspaceOptions(
				"testString",
			)

			workspace, response, err := schematics20ApiService.GetWorkspace(getWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspace, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspace).ToNot(BeNil())

		})
		It(`ReplaceWorkspace request example`, func() {
			// begin-replace_workspace

			workspaceModel := &schematics20apiv2.Workspace{
				Name:          core.StringPtr("My Workspace"),
				ResourceGroup: core.StringPtr("testString"),
				Location:      core.StringPtr("us-south"),
			}

			replaceWorkspaceOptions := schematics20ApiService.NewReplaceWorkspaceOptions(
				"testString",
				workspaceModel,
			)

			workspace, response, err := schematics20ApiService.ReplaceWorkspace(replaceWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspace, "", "  ")
			fmt.Println(string(b))

			// end-replace_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspace).ToNot(BeNil())

		})
		It(`UpdateWorkspace request example`, func() {
			// begin-update_workspace

			workspaceModel := &schematics20apiv2.Workspace{
				Name:          core.StringPtr("My Workspace"),
				ResourceGroup: core.StringPtr("testString"),
				Location:      core.StringPtr("us-south"),
			}

			updateWorkspaceOptions := schematics20ApiService.NewUpdateWorkspaceOptions(
				"testString",
				workspaceModel,
			)

			workspace, response, err := schematics20ApiService.UpdateWorkspace(updateWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspace, "", "  ")
			fmt.Println(string(b))

			// end-update_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspace).ToNot(BeNil())

		})
		It(`ListWorkspaceInputs request example`, func() {
			// begin-list_workspace_inputs

			listWorkspaceInputsOptions := schematics20ApiService.NewListWorkspaceInputsOptions(
				"testString",
			)

			inputVariableDataList, response, err := schematics20ApiService.ListWorkspaceInputs(listWorkspaceInputsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(inputVariableDataList, "", "  ")
			fmt.Println(string(b))

			// end-list_workspace_inputs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inputVariableDataList).ToNot(BeNil())

		})
		It(`ListWorkspaceOutputs request example`, func() {
			// begin-list_workspace_outputs

			listWorkspaceOutputsOptions := schematics20ApiService.NewListWorkspaceOutputsOptions(
				"testString",
			)

			outputVariableDataList, response, err := schematics20ApiService.ListWorkspaceOutputs(listWorkspaceOutputsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(outputVariableDataList, "", "  ")
			fmt.Println(string(b))

			// end-list_workspace_outputs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(outputVariableDataList).ToNot(BeNil())

		})
		It(`CreateWorkspaceTemplates request example`, func() {
			// begin-create_workspace_templates

			templateListRequestModel := &schematics20apiv2.TemplateListRequest{}

			createWorkspaceTemplatesOptions := schematics20ApiService.NewCreateWorkspaceTemplatesOptions(
				templateListRequestModel,
			)

			workspace, response, err := schematics20ApiService.CreateWorkspaceTemplates(createWorkspaceTemplatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(workspace, "", "  ")
			fmt.Println(string(b))

			// end-create_workspace_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(workspace).ToNot(BeNil())

		})
		It(`ListTemplatesInWorkspace request example`, func() {
			// begin-list_templates_in_workspace

			listTemplatesInWorkspaceOptions := schematics20ApiService.NewListTemplatesInWorkspaceOptions(
				"testString",
			)

			templateList, response, err := schematics20ApiService.ListTemplatesInWorkspace(listTemplatesInWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateList, "", "  ")
			fmt.Println(string(b))

			// end-list_templates_in_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateList).ToNot(BeNil())

		})
		It(`AddTemplatesToWorkspace request example`, func() {
			// begin-add_templates_to_workspace

			templateModel := &schematics20apiv2.Template{}

			addTemplatesToWorkspaceOptions := schematics20ApiService.NewAddTemplatesToWorkspaceOptions(
				"testString",
				templateModel,
			)

			template, response, err := schematics20ApiService.AddTemplatesToWorkspace(addTemplatesToWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(template, "", "  ")
			fmt.Println(string(b))

			// end-add_templates_to_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(template).ToNot(BeNil())

		})
		It(`GetTemplateInWorkspace request example`, func() {
			// begin-get_template_in_workspace

			getTemplateInWorkspaceOptions := schematics20ApiService.NewGetTemplateInWorkspaceOptions(
				"testString",
				"testString",
			)

			template, response, err := schematics20ApiService.GetTemplateInWorkspace(getTemplateInWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(template, "", "  ")
			fmt.Println(string(b))

			// end-get_template_in_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(template).ToNot(BeNil())

		})
		It(`ReplaceTemplateInWorkspace request example`, func() {
			// begin-replace_template_in_workspace

			replaceTemplateInWorkspaceOptions := schematics20ApiService.NewReplaceTemplateInWorkspaceOptions(
				"testString",
				"testString",
			)

			template, response, err := schematics20ApiService.ReplaceTemplateInWorkspace(replaceTemplateInWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(template, "", "  ")
			fmt.Println(string(b))

			// end-replace_template_in_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(template).ToNot(BeNil())

		})
		It(`UpdateTemplateInWorkspace request example`, func() {
			// begin-update_template_in_workspace

			updateTemplateInWorkspaceOptions := schematics20ApiService.NewUpdateTemplateInWorkspaceOptions(
				"testString",
				"testString",
			)

			template, response, err := schematics20ApiService.UpdateTemplateInWorkspace(updateTemplateInWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(template, "", "  ")
			fmt.Println(string(b))

			// end-update_template_in_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(template).ToNot(BeNil())

		})
		It(`UploadTemplateTarInWorkspace request example`, func() {
			// begin-upload_template_tar_in_workspace

			uploadTemplateTarInWorkspaceOptions := schematics20ApiService.NewUploadTemplateTarInWorkspaceOptions(
				"testString",
				"testString",
			)

			template, response, err := schematics20ApiService.UploadTemplateTarInWorkspace(uploadTemplateTarInWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(template, "", "  ")
			fmt.Println(string(b))

			// end-upload_template_tar_in_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(template).ToNot(BeNil())

		})
		It(`GetTemplateReadme request example`, func() {
			// begin-get_template_readme

			getTemplateReadmeOptions := schematics20ApiService.NewGetTemplateReadmeOptions(
				"testString",
				"testString",
			)

			result, response, err := schematics20ApiService.GetTemplateReadme(getTemplateReadmeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(result, "", "  ")
			fmt.Println(string(b))

			// end-get_template_readme

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
		It(`GetTemplateSource request example`, func() {
			// begin-get_template_source

			getTemplateSourceOptions := schematics20ApiService.NewGetTemplateSourceOptions(
				"testString",
				"testString",
			)

			externalSource, response, err := schematics20ApiService.GetTemplateSource(getTemplateSourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(externalSource, "", "  ")
			fmt.Println(string(b))

			// end-get_template_source

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(externalSource).ToNot(BeNil())

		})
		It(`ListTemplateInputs request example`, func() {
			// begin-list_template_inputs

			listTemplateInputsOptions := schematics20ApiService.NewListTemplateInputsOptions(
				"testString",
				"testString",
			)

			variableDataList, response, err := schematics20ApiService.ListTemplateInputs(listTemplateInputsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(variableDataList, "", "  ")
			fmt.Println(string(b))

			// end-list_template_inputs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(variableDataList).ToNot(BeNil())

		})
		It(`ListTemplateOutputs request example`, func() {
			// begin-list_template_outputs

			listTemplateOutputsOptions := schematics20ApiService.NewListTemplateOutputsOptions(
				"testString",
				"testString",
			)

			variableDataList, response, err := schematics20ApiService.ListTemplateOutputs(listTemplateOutputsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(variableDataList, "", "  ")
			fmt.Println(string(b))

			// end-list_template_outputs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(variableDataList).ToNot(BeNil())

		})
		It(`ListFlowsInWorkspace request example`, func() {
			// begin-list_flows_in_workspace

			listFlowsInWorkspaceOptions := schematics20ApiService.NewListFlowsInWorkspaceOptions(
				"testString",
			)

			templateFlowList, response, err := schematics20ApiService.ListFlowsInWorkspace(listFlowsInWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateFlowList, "", "  ")
			fmt.Println(string(b))

			// end-list_flows_in_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateFlowList).ToNot(BeNil())

		})
		It(`AddFlowToWorkspace request example`, func() {
			// begin-add_flow_to_workspace

			templateFlowModel := &schematics20apiv2.TemplateFlow{}

			addFlowToWorkspaceOptions := schematics20ApiService.NewAddFlowToWorkspaceOptions(
				"testString",
				templateFlowModel,
			)

			templateFlow, response, err := schematics20ApiService.AddFlowToWorkspace(addFlowToWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateFlow, "", "  ")
			fmt.Println(string(b))

			// end-add_flow_to_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateFlow).ToNot(BeNil())

		})
		It(`GetFlowInWorkspace request example`, func() {
			// begin-get_flow_in_workspace

			getFlowInWorkspaceOptions := schematics20ApiService.NewGetFlowInWorkspaceOptions(
				"testString",
				"testString",
			)

			templateFlow, response, err := schematics20ApiService.GetFlowInWorkspace(getFlowInWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateFlow, "", "  ")
			fmt.Println(string(b))

			// end-get_flow_in_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateFlow).ToNot(BeNil())

		})
		It(`ReplaceFlowInWorkspace request example`, func() {
			// begin-replace_flow_in_workspace

			replaceFlowInWorkspaceOptions := schematics20ApiService.NewReplaceFlowInWorkspaceOptions(
				"testString",
				"testString",
			)

			templateFlow, response, err := schematics20ApiService.ReplaceFlowInWorkspace(replaceFlowInWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateFlow, "", "  ")
			fmt.Println(string(b))

			// end-replace_flow_in_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateFlow).ToNot(BeNil())

		})
		It(`UpdateFlowInWorkspace request example`, func() {
			// begin-update_flow_in_workspace

			updateFlowInWorkspaceOptions := schematics20ApiService.NewUpdateFlowInWorkspaceOptions(
				"testString",
				"testString",
			)

			templateFlow, response, err := schematics20ApiService.UpdateFlowInWorkspace(updateFlowInWorkspaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateFlow, "", "  ")
			fmt.Println(string(b))

			// end-update_flow_in_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateFlow).ToNot(BeNil())

		})
		It(`ListActions request example`, func() {
			// begin-list_actions

			listActionsOptions := schematics20ApiService.NewListActionsOptions()

			actionList, response, err := schematics20ApiService.ListActions(listActionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(actionList, "", "  ")
			fmt.Println(string(b))

			// end-list_actions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(actionList).ToNot(BeNil())

		})
		It(`CreateAction request example`, func() {
			// begin-create_action

			actionModel := &schematics20apiv2.Action{}

			createActionOptions := schematics20ApiService.NewCreateActionOptions(
				actionModel,
			)

			action, response, err := schematics20ApiService.CreateAction(createActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(action, "", "  ")
			fmt.Println(string(b))

			// end-create_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(action).ToNot(BeNil())

		})
		It(`GetAction request example`, func() {
			// begin-get_action

			getActionOptions := schematics20ApiService.NewGetActionOptions(
				"testString",
			)

			action, response, err := schematics20ApiService.GetAction(getActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(action, "", "  ")
			fmt.Println(string(b))

			// end-get_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(action).ToNot(BeNil())

		})
		It(`UpdateAction request example`, func() {
			// begin-update_action

			actionModel := &schematics20apiv2.Action{}

			updateActionOptions := schematics20ApiService.NewUpdateActionOptions(
				"testString",
				actionModel,
			)

			action, response, err := schematics20ApiService.UpdateAction(updateActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(action, "", "  ")
			fmt.Println(string(b))

			// end-update_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(action).ToNot(BeNil())

		})
		It(`UploadTemplateTarAction request example`, func() {
			// begin-upload_template_tar_action

			uploadTemplateTarActionOptions := schematics20ApiService.NewUploadTemplateTarActionOptions(
				"testString",
			)

			templateRepoTarUploadResponse, response, err := schematics20ApiService.UploadTemplateTarAction(uploadTemplateTarActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateRepoTarUploadResponse, "", "  ")
			fmt.Println(string(b))

			// end-upload_template_tar_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateRepoTarUploadResponse).ToNot(BeNil())

		})
		It(`ListJobs request example`, func() {
			// begin-list_jobs

			listJobsOptions := schematics20ApiService.NewListJobsOptions()

			jobList, response, err := schematics20ApiService.ListJobs(listJobsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobList, "", "  ")
			fmt.Println(string(b))

			// end-list_jobs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobList).ToNot(BeNil())

		})
		It(`CreateJob request example`, func() {
			// begin-create_job

			jobModel := &schematics20apiv2.Job{}

			createJobOptions := schematics20ApiService.NewCreateJobOptions(
				"testString",
				jobModel,
			)

			job, response, err := schematics20ApiService.CreateJob(createJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(job, "", "  ")
			fmt.Println(string(b))

			// end-create_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(job).ToNot(BeNil())

		})
		It(`GetJob request example`, func() {
			// begin-get_job

			getJobOptions := schematics20ApiService.NewGetJobOptions(
				"testString",
			)

			job, response, err := schematics20ApiService.GetJob(getJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(job, "", "  ")
			fmt.Println(string(b))

			// end-get_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())

		})
		It(`UpdateJob request example`, func() {
			// begin-update_job

			jobModel := &schematics20apiv2.Job{}

			updateJobOptions := schematics20ApiService.NewUpdateJobOptions(
				"testString",
				"testString",
				jobModel,
			)

			job, response, err := schematics20ApiService.UpdateJob(updateJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(job, "", "  ")
			fmt.Println(string(b))

			// end-update_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(job).ToNot(BeNil())

		})
		It(`ListJobLogs request example`, func() {
			// begin-list_job_logs

			listJobLogsOptions := schematics20ApiService.NewListJobLogsOptions(
				"testString",
			)

			jobLog, response, err := schematics20ApiService.ListJobLogs(listJobLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobLog, "", "  ")
			fmt.Println(string(b))

			// end-list_job_logs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobLog).ToNot(BeNil())

		})
		It(`GetJobFiles request example`, func() {
			// begin-get_job_files

			getJobFilesOptions := schematics20ApiService.NewGetJobFilesOptions(
				"testString",
				"state_file",
			)

			jobFileData, response, err := schematics20ApiService.GetJobFiles(getJobFilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobFileData, "", "  ")
			fmt.Println(string(b))

			// end-get_job_files

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobFileData).ToNot(BeNil())

		})
		It(`ListControls request example`, func() {
			// begin-list_controls

			listControlsOptions := schematics20ApiService.NewListControlsOptions()

			controlsList, response, err := schematics20ApiService.ListControls(listControlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controlsList, "", "  ")
			fmt.Println(string(b))

			// end-list_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controlsList).ToNot(BeNil())

		})
		It(`CreateControls request example`, func() {
			// begin-create_controls

			controlsModel := &schematics20apiv2.Controls{}

			createControlsOptions := schematics20ApiService.NewCreateControlsOptions(
				controlsModel,
			)

			controls, response, err := schematics20ApiService.CreateControls(createControlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controls, "", "  ")
			fmt.Println(string(b))

			// end-create_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controls).ToNot(BeNil())

		})
		It(`GetControls request example`, func() {
			// begin-get_controls

			getControlsOptions := schematics20ApiService.NewGetControlsOptions(
				"testString",
			)

			controls, response, err := schematics20ApiService.GetControls(getControlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controls, "", "  ")
			fmt.Println(string(b))

			// end-get_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controls).ToNot(BeNil())

		})
		It(`UpdateControls request example`, func() {
			// begin-update_controls

			controlsModel := &schematics20apiv2.Controls{}

			updateControlsOptions := schematics20ApiService.NewUpdateControlsOptions(
				"testString",
				controlsModel,
			)

			controls, response, err := schematics20ApiService.UpdateControls(updateControlsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(controls, "", "  ")
			fmt.Println(string(b))

			// end-update_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(controls).ToNot(BeNil())

		})
		It(`ListCapsules request example`, func() {
			// begin-list_capsules

			listCapsulesOptions := schematics20ApiService.NewListCapsulesOptions(
				"testString",
			)

			capsuleList, response, err := schematics20ApiService.ListCapsules(listCapsulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(capsuleList, "", "  ")
			fmt.Println(string(b))

			// end-list_capsules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(capsuleList).ToNot(BeNil())

		})
		It(`CreateCapsule request example`, func() {
			// begin-create_capsule

			capsuleModel := &schematics20apiv2.Capsule{}

			createCapsuleOptions := schematics20ApiService.NewCreateCapsuleOptions(
				"testString",
				capsuleModel,
			)

			capsuleList, response, err := schematics20ApiService.CreateCapsule(createCapsuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(capsuleList, "", "  ")
			fmt.Println(string(b))

			// end-create_capsule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(capsuleList).ToNot(BeNil())

		})
		It(`GetCapsule request example`, func() {
			// begin-get_capsule

			getCapsuleOptions := schematics20ApiService.NewGetCapsuleOptions(
				"testString",
				"testString",
			)

			capsule, response, err := schematics20ApiService.GetCapsule(getCapsuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(capsule, "", "  ")
			fmt.Println(string(b))

			// end-get_capsule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(capsule).ToNot(BeNil())

		})
		It(`UpdateCapsule request example`, func() {
			// begin-update_capsule

			capsuleModel := &schematics20apiv2.Capsule{}

			updateCapsuleOptions := schematics20ApiService.NewUpdateCapsuleOptions(
				"testString",
				"testString",
				capsuleModel,
			)

			capsule, response, err := schematics20ApiService.UpdateCapsule(updateCapsuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(capsule, "", "  ")
			fmt.Println(string(b))

			// end-update_capsule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(capsule).ToNot(BeNil())

		})
		It(`ListDatasets request example`, func() {
			// begin-list_datasets

			listDatasetsOptions := schematics20ApiService.NewListDatasetsOptions()

			datasetList, response, err := schematics20ApiService.ListDatasets(listDatasetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(datasetList, "", "  ")
			fmt.Println(string(b))

			// end-list_datasets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(datasetList).ToNot(BeNil())

		})
		It(`CreateDataset request example`, func() {
			// begin-create_dataset

			datasetModel := &schematics20apiv2.Dataset{}

			createDatasetOptions := schematics20ApiService.NewCreateDatasetOptions(
				datasetModel,
			)

			dataset, response, err := schematics20ApiService.CreateDataset(createDatasetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataset, "", "  ")
			fmt.Println(string(b))

			// end-create_dataset

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataset).ToNot(BeNil())

		})
		It(`GetDataset request example`, func() {
			// begin-get_dataset

			getDatasetOptions := schematics20ApiService.NewGetDatasetOptions(
				"testString",
			)

			dataset, response, err := schematics20ApiService.GetDataset(getDatasetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataset, "", "  ")
			fmt.Println(string(b))

			// end-get_dataset

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataset).ToNot(BeNil())

		})
		It(`ReplaceDataset request example`, func() {
			// begin-replace_dataset

			datasetModel := &schematics20apiv2.Dataset{}

			replaceDatasetOptions := schematics20ApiService.NewReplaceDatasetOptions(
				"testString",
				datasetModel,
			)

			dataset, response, err := schematics20ApiService.ReplaceDataset(replaceDatasetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataset, "", "  ")
			fmt.Println(string(b))

			// end-replace_dataset

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataset).ToNot(BeNil())

		})
		It(`UpdateAtaset request example`, func() {
			// begin-update_ataset

			datasetModel := &schematics20apiv2.Dataset{}

			updateAtasetOptions := schematics20ApiService.NewUpdateAtasetOptions(
				"testString",
				datasetModel,
			)

			dataset, response, err := schematics20ApiService.UpdateAtaset(updateAtasetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dataset, "", "  ")
			fmt.Println(string(b))

			// end-update_ataset

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dataset).ToNot(BeNil())

		})
		It(`ListDatasetVariables request example`, func() {
			// begin-list_dataset_variables

			listDatasetVariablesOptions := schematics20ApiService.NewListDatasetVariablesOptions(
				"testString",
			)

			variableDataList, response, err := schematics20ApiService.ListDatasetVariables(listDatasetVariablesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(variableDataList, "", "  ")
			fmt.Println(string(b))

			// end-list_dataset_variables

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(variableDataList).ToNot(BeNil())

		})
		It(`GetDatasetVariable request example`, func() {
			// begin-get_dataset_variable

			getDatasetVariableOptions := schematics20ApiService.NewGetDatasetVariableOptions(
				"testString",
				"testString",
			)

			variableData, response, err := schematics20ApiService.GetDatasetVariable(getDatasetVariableOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(variableData, "", "  ")
			fmt.Println(string(b))

			// end-get_dataset_variable

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(variableData).ToNot(BeNil())

		})
		It(`ListInventories request example`, func() {
			// begin-list_inventories

			listInventoriesOptions := schematics20ApiService.NewListInventoriesOptions()

			inventoryResourceRecordList, response, err := schematics20ApiService.ListInventories(listInventoriesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(inventoryResourceRecordList, "", "  ")
			fmt.Println(string(b))

			// end-list_inventories

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecordList).ToNot(BeNil())

		})
		It(`CreateInventory request example`, func() {
			// begin-create_inventory

			inventoryResourceDefinitionModel := &schematics20apiv2.InventoryResourceDefinition{}

			createInventoryOptions := schematics20ApiService.NewCreateInventoryOptions(
				inventoryResourceDefinitionModel,
			)

			inventoryResourceRecord, response, err := schematics20ApiService.CreateInventory(createInventoryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(inventoryResourceRecord, "", "  ")
			fmt.Println(string(b))

			// end-create_inventory

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecord).ToNot(BeNil())

		})
		It(`GetInventory request example`, func() {
			// begin-get_inventory

			getInventoryOptions := schematics20ApiService.NewGetInventoryOptions(
				"testString",
			)

			inventoryResourceRecord, response, err := schematics20ApiService.GetInventory(getInventoryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(inventoryResourceRecord, "", "  ")
			fmt.Println(string(b))

			// end-get_inventory

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecord).ToNot(BeNil())

		})
		It(`ReplaceInventory request example`, func() {
			// begin-replace_inventory

			inventoryResourceDefinitionModel := &schematics20apiv2.InventoryResourceDefinition{}

			replaceInventoryOptions := schematics20ApiService.NewReplaceInventoryOptions(
				"testString",
				inventoryResourceDefinitionModel,
			)

			inventoryResourceRecord, response, err := schematics20ApiService.ReplaceInventory(replaceInventoryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(inventoryResourceRecord, "", "  ")
			fmt.Println(string(b))

			// end-replace_inventory

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(inventoryResourceRecord).ToNot(BeNil())

		})
		It(`ListResourceQuery request example`, func() {
			// begin-list_resource_query

			listResourceQueryOptions := schematics20ApiService.NewListResourceQueryOptions()

			resourceQueryRecordList, response, err := schematics20ApiService.ListResourceQuery(listResourceQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryRecordList, "", "  ")
			fmt.Println(string(b))

			// end-list_resource_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecordList).ToNot(BeNil())

		})
		It(`CreateResourceQuery request example`, func() {
			// begin-create_resource_query

			resourceQueryDefinitionModel := &schematics20apiv2.ResourceQueryDefinition{}

			createResourceQueryOptions := schematics20ApiService.NewCreateResourceQueryOptions(
				resourceQueryDefinitionModel,
			)

			resourceQueryRecord, response, err := schematics20ApiService.CreateResourceQuery(createResourceQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryRecord, "", "  ")
			fmt.Println(string(b))

			// end-create_resource_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecord).ToNot(BeNil())

		})
		It(`GetResourcesQuery request example`, func() {
			// begin-get_resources_query

			getResourcesQueryOptions := schematics20ApiService.NewGetResourcesQueryOptions(
				"testString",
			)

			resourceQueryRecord, response, err := schematics20ApiService.GetResourcesQuery(getResourcesQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryRecord, "", "  ")
			fmt.Println(string(b))

			// end-get_resources_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecord).ToNot(BeNil())

		})
		It(`ReplaceResourcesQuery request example`, func() {
			// begin-replace_resources_query

			resourceQueryDefinitionModel := &schematics20apiv2.ResourceQueryDefinition{}

			replaceResourcesQueryOptions := schematics20ApiService.NewReplaceResourcesQueryOptions(
				"testString",
				resourceQueryDefinitionModel,
			)

			resourceQueryRecord, response, err := schematics20ApiService.ReplaceResourcesQuery(replaceResourcesQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryRecord, "", "  ")
			fmt.Println(string(b))

			// end-replace_resources_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryRecord).ToNot(BeNil())

		})
		It(`ExecuteResourceQuery request example`, func() {
			// begin-execute_resource_query

			executeResourceQueryOptions := schematics20ApiService.NewExecuteResourceQueryOptions(
				"testString",
			)

			resourceQueryResponseRecord, response, err := schematics20ApiService.ExecuteResourceQuery(executeResourceQueryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceQueryResponseRecord, "", "  ")
			fmt.Println(string(b))

			// end-execute_resource_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceQueryResponseRecord).ToNot(BeNil())

		})
		It(`ListAgent request example`, func() {
			// begin-list_agent

			listAgentOptions := schematics20ApiService.NewListAgentOptions()

			agentList, response, err := schematics20ApiService.ListAgent(listAgentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agentList, "", "  ")
			fmt.Println(string(b))

			// end-list_agent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agentList).ToNot(BeNil())

		})
		It(`RegisterAgent request example`, func() {
			// begin-register_agent

			agentModel := &schematics20apiv2.Agent{
				Name:      core.StringPtr("MyDevAgent"),
				ProfileID: core.StringPtr("testString"),
				AgentID:   core.StringPtr("testString"),
			}

			registerAgentOptions := schematics20ApiService.NewRegisterAgentOptions(
				agentModel,
			)

			agent, response, err := schematics20ApiService.RegisterAgent(registerAgentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agent, "", "  ")
			fmt.Println(string(b))

			// end-register_agent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agent).ToNot(BeNil())

		})
		It(`GetAgent request example`, func() {
			// begin-get_agent

			getAgentOptions := schematics20ApiService.NewGetAgentOptions(
				"testString",
			)

			agent, response, err := schematics20ApiService.GetAgent(getAgentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agent, "", "  ")
			fmt.Println(string(b))

			// end-get_agent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(agent).ToNot(BeNil())

		})
		It(`UpdateAgentRegistration request example`, func() {
			// begin-update_agent_registration

			agentModel := &schematics20apiv2.Agent{
				Name:      core.StringPtr("MyDevAgent"),
				ProfileID: core.StringPtr("testString"),
				AgentID:   core.StringPtr("testString"),
			}

			updateAgentRegistrationOptions := schematics20ApiService.NewUpdateAgentRegistrationOptions(
				"testString",
				agentModel,
			)

			agent, response, err := schematics20ApiService.UpdateAgentRegistration(updateAgentRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(agent, "", "  ")
			fmt.Println(string(b))

			// end-update_agent_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(agent).ToNot(BeNil())

		})
		It(`GetDatasetVariableValue request example`, func() {
			// begin-get_dataset_variable_value

			getDatasetVariableValueOptions := schematics20ApiService.NewGetDatasetVariableValueOptions(
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetDatasetVariableValue(getDatasetVariableValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_dataset_variable_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetCredentialVariableValue request example`, func() {
			// begin-get_credential_variable_value

			getCredentialVariableValueOptions := schematics20ApiService.NewGetCredentialVariableValueOptions(
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetCredentialVariableValue(getCredentialVariableValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_credential_variable_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetInventoryVariableValue request example`, func() {
			// begin-get_inventory_variable_value

			getInventoryVariableValueOptions := schematics20ApiService.NewGetInventoryVariableValueOptions(
				"testString",
				"testString",
			)

			valueListResponse, response, err := schematics20ApiService.GetInventoryVariableValue(getInventoryVariableValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueListResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_inventory_variable_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueListResponse).ToNot(BeNil())

		})
		It(`GetWorkspaceInputValue request example`, func() {
			// begin-get_workspace_input_value

			getWorkspaceInputValueOptions := schematics20ApiService.NewGetWorkspaceInputValueOptions(
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetWorkspaceInputValue(getWorkspaceInputValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_input_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetWorkspaceOutputValue request example`, func() {
			// begin-get_workspace_output_value

			getWorkspaceOutputValueOptions := schematics20ApiService.NewGetWorkspaceOutputValueOptions(
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetWorkspaceOutputValue(getWorkspaceOutputValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_output_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetWorkspaceSettingsValue request example`, func() {
			// begin-get_workspace_settings_value

			getWorkspaceSettingsValueOptions := schematics20ApiService.NewGetWorkspaceSettingsValueOptions(
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetWorkspaceSettingsValue(getWorkspaceSettingsValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_settings_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetWorkspaceTemplateInputValue request example`, func() {
			// begin-get_workspace_template_input_value

			getWorkspaceTemplateInputValueOptions := schematics20ApiService.NewGetWorkspaceTemplateInputValueOptions(
				"testString",
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetWorkspaceTemplateInputValue(getWorkspaceTemplateInputValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_template_input_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetWorkspaceTemplateOutputValue request example`, func() {
			// begin-get_workspace_template_output_value

			getWorkspaceTemplateOutputValueOptions := schematics20ApiService.NewGetWorkspaceTemplateOutputValueOptions(
				"testString",
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetWorkspaceTemplateOutputValue(getWorkspaceTemplateOutputValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_template_output_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetWorkspaceTemplateSettingValue request example`, func() {
			// begin-get_workspace_template_setting_value

			getWorkspaceTemplateSettingValueOptions := schematics20ApiService.NewGetWorkspaceTemplateSettingValueOptions(
				"testString",
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetWorkspaceTemplateSettingValue(getWorkspaceTemplateSettingValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_workspace_template_setting_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetJobWorkspaceInputValue request example`, func() {
			// begin-get_job_workspace_input_value

			getJobWorkspaceInputValueOptions := schematics20ApiService.NewGetJobWorkspaceInputValueOptions(
				"testString",
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetJobWorkspaceInputValue(getJobWorkspaceInputValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_job_workspace_input_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetJobWorkspaceOutputValue request example`, func() {
			// begin-get_job_workspace_output_value

			getJobWorkspaceOutputValueOptions := schematics20ApiService.NewGetJobWorkspaceOutputValueOptions(
				"testString",
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetJobWorkspaceOutputValue(getJobWorkspaceOutputValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_job_workspace_output_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetJobWorkspaceSettingValue request example`, func() {
			// begin-get_job_workspace_setting_value

			getJobWorkspaceSettingValueOptions := schematics20ApiService.NewGetJobWorkspaceSettingValueOptions(
				"testString",
				"testString",
				"testString",
			)

			valueResponse, response, err := schematics20ApiService.GetJobWorkspaceSettingValue(getJobWorkspaceSettingValueOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(valueResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_job_workspace_setting_value

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(valueResponse).ToNot(BeNil())

		})
		It(`GetKmsSettings request example`, func() {
			// begin-get_kms_settings

			getKmsSettingsOptions := schematics20ApiService.NewGetKmsSettingsOptions(
				"testString",
			)

			kmsSettings, response, err := schematics20ApiService.GetKmsSettings(getKmsSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(kmsSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_kms_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(kmsSettings).ToNot(BeNil())

		})
		It(`UpdateKmsSettings request example`, func() {
			// begin-update_kms_settings

			kmsSettingsModel := &schematics20apiv2.KMSSettings{}

			updateKmsSettingsOptions := schematics20ApiService.NewUpdateKmsSettingsOptions(
				kmsSettingsModel,
			)

			kmsSettings, response, err := schematics20ApiService.UpdateKmsSettings(updateKmsSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(kmsSettings, "", "  ")
			fmt.Println(string(b))

			// end-update_kms_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(kmsSettings).ToNot(BeNil())

		})
		It(`ListKms request example`, func() {
			// begin-list_kms

			listKmsOptions := schematics20ApiService.NewListKmsOptions(
				"testString",
				"testString",
			)

			kmsDiscovery, response, err := schematics20ApiService.ListKms(listKmsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(kmsDiscovery, "", "  ")
			fmt.Println(string(b))

			// end-list_kms

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(kmsDiscovery).ToNot(BeNil())

		})
		It(`ListTriggers request example`, func() {
			// begin-list_triggers

			listTriggersOptions := schematics20ApiService.NewListTriggersOptions()

			triggerList, response, err := schematics20ApiService.ListTriggers(listTriggersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(triggerList, "", "  ")
			fmt.Println(string(b))

			// end-list_triggers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(triggerList).ToNot(BeNil())

		})
		It(`CreateTrigger request example`, func() {
			// begin-create_trigger

			triggerModel := &schematics20apiv2.Trigger{}

			createTriggerOptions := schematics20ApiService.NewCreateTriggerOptions(
				triggerModel,
			)

			trigger, response, err := schematics20ApiService.CreateTrigger(createTriggerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trigger, "", "  ")
			fmt.Println(string(b))

			// end-create_trigger

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trigger).ToNot(BeNil())

		})
		It(`GetTrigger request example`, func() {
			// begin-get_trigger

			getTriggerOptions := schematics20ApiService.NewGetTriggerOptions(
				"testString",
			)

			trigger, response, err := schematics20ApiService.GetTrigger(getTriggerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trigger, "", "  ")
			fmt.Println(string(b))

			// end-get_trigger

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trigger).ToNot(BeNil())

		})
		It(`ReplaceTrigger request example`, func() {
			// begin-replace_trigger

			triggerModel := &schematics20apiv2.Trigger{}

			replaceTriggerOptions := schematics20ApiService.NewReplaceTriggerOptions(
				"testString",
				triggerModel,
			)

			trigger, response, err := schematics20ApiService.ReplaceTrigger(replaceTriggerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(trigger, "", "  ")
			fmt.Println(string(b))

			// end-replace_trigger

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(trigger).ToNot(BeNil())

		})
		It(`ListWebhook request example`, func() {
			// begin-list_webhook

			listWebhookOptions := schematics20ApiService.NewListWebhookOptions()

			webHookList, response, err := schematics20ApiService.ListWebhook(listWebhookOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(webHookList, "", "  ")
			fmt.Println(string(b))

			// end-list_webhook

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(webHookList).ToNot(BeNil())

		})
		It(`CreateWebhook request example`, func() {
			// begin-create_webhook

			webHookModel := &schematics20apiv2.WebHook{}

			createWebhookOptions := schematics20ApiService.NewCreateWebhookOptions(
				webHookModel,
			)

			webHook, response, err := schematics20ApiService.CreateWebhook(createWebhookOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(webHook, "", "  ")
			fmt.Println(string(b))

			// end-create_webhook

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(webHook).ToNot(BeNil())

		})
		It(`GetWebhook request example`, func() {
			// begin-get_webhook

			getWebhookOptions := schematics20ApiService.NewGetWebhookOptions(
				"testString",
			)

			webHook, response, err := schematics20ApiService.GetWebhook(getWebhookOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(webHook, "", "  ")
			fmt.Println(string(b))

			// end-get_webhook

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(webHook).ToNot(BeNil())

		})
		It(`ReplaceWebhook request example`, func() {
			// begin-replace_webhook

			webHookModel := &schematics20apiv2.WebHook{}

			replaceWebhookOptions := schematics20ApiService.NewReplaceWebhookOptions(
				"testString",
				webHookModel,
			)

			webHook, response, err := schematics20ApiService.ReplaceWebhook(replaceWebhookOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(webHook, "", "  ")
			fmt.Println(string(b))

			// end-replace_webhook

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(webHook).ToNot(BeNil())

		})
		It(`GetPrivateCluster request example`, func() {
			// begin-get_private_cluster

			getPrivateClusterOptions := schematics20ApiService.NewGetPrivateClusterOptions()

			privateClusterList, response, err := schematics20ApiService.GetPrivateCluster(getPrivateClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(privateClusterList, "", "  ")
			fmt.Println(string(b))

			// end-get_private_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(privateClusterList).ToNot(BeNil())

		})
		It(`CreatePrivateCluster request example`, func() {
			// begin-create_private_cluster

			privateClusterModel := &schematics20apiv2.PrivateCluster{}

			createPrivateClusterOptions := schematics20ApiService.NewCreatePrivateClusterOptions(
				privateClusterModel,
			)

			privateCluster, response, err := schematics20ApiService.CreatePrivateCluster(createPrivateClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(privateCluster, "", "  ")
			fmt.Println(string(b))

			// end-create_private_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(privateCluster).ToNot(BeNil())

		})
		It(`ListAdapter request example`, func() {
			// begin-list_adapter

			listAdapterOptions := schematics20ApiService.NewListAdapterOptions()

			adapterList, response, err := schematics20ApiService.ListAdapter(listAdapterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(adapterList, "", "  ")
			fmt.Println(string(b))

			// end-list_adapter

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(adapterList).ToNot(BeNil())

		})
		It(`CreateAdapter request example`, func() {
			// begin-create_adapter

			adapterModel := &schematics20apiv2.Adapter{}

			createAdapterOptions := schematics20ApiService.NewCreateAdapterOptions(
				adapterModel,
			)

			adapter, response, err := schematics20ApiService.CreateAdapter(createAdapterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(adapter, "", "  ")
			fmt.Println(string(b))

			// end-create_adapter

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(adapter).ToNot(BeNil())

		})
		It(`GetAdapter request example`, func() {
			// begin-get_adapter

			getAdapterOptions := schematics20ApiService.NewGetAdapterOptions(
				"testString",
			)

			adapter, response, err := schematics20ApiService.GetAdapter(getAdapterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(adapter, "", "  ")
			fmt.Println(string(b))

			// end-get_adapter

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(adapter).ToNot(BeNil())

		})
		It(`ReplaceAdapter request example`, func() {
			// begin-replace_adapter

			adapterModel := &schematics20apiv2.Adapter{}

			replaceAdapterOptions := schematics20ApiService.NewReplaceAdapterOptions(
				"testString",
				adapterModel,
			)

			adapter, response, err := schematics20ApiService.ReplaceAdapter(replaceAdapterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(adapter, "", "  ")
			fmt.Println(string(b))

			// end-replace_adapter

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(adapter).ToNot(BeNil())

		})
		It(`ListConnection request example`, func() {
			// begin-list_connection

			listConnectionOptions := schematics20ApiService.NewListConnectionOptions()

			connectionList, response, err := schematics20ApiService.ListConnection(listConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(connectionList, "", "  ")
			fmt.Println(string(b))

			// end-list_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connectionList).ToNot(BeNil())

		})
		It(`CreateConnection request example`, func() {
			// begin-create_connection

			connectionModel := &schematics20apiv2.Connection{}

			createConnectionOptions := schematics20ApiService.NewCreateConnectionOptions(
				connectionModel,
			)

			connection, response, err := schematics20ApiService.CreateConnection(createConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(connection, "", "  ")
			fmt.Println(string(b))

			// end-create_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connection).ToNot(BeNil())

		})
		It(`GetConnection request example`, func() {
			// begin-get_connection

			getConnectionOptions := schematics20ApiService.NewGetConnectionOptions(
				"testString",
			)

			connectionLite, response, err := schematics20ApiService.GetConnection(getConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(connectionLite, "", "  ")
			fmt.Println(string(b))

			// end-get_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connectionLite).ToNot(BeNil())

		})
		It(`ListDatasources request example`, func() {
			// begin-list_datasources

			listDatasourcesOptions := schematics20ApiService.NewListDatasourcesOptions()

			datasourceList, response, err := schematics20ApiService.ListDatasources(listDatasourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(datasourceList, "", "  ")
			fmt.Println(string(b))

			// end-list_datasources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(datasourceList).ToNot(BeNil())

		})
		It(`CreateDatasource request example`, func() {
			// begin-create_datasource

			datasourceModel := &schematics20apiv2.Datasource{}

			createDatasourceOptions := schematics20ApiService.NewCreateDatasourceOptions(
				datasourceModel,
			)

			datasource, response, err := schematics20ApiService.CreateDatasource(createDatasourceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(datasource, "", "  ")
			fmt.Println(string(b))

			// end-create_datasource

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(datasource).ToNot(BeNil())

		})
		It(`GetData request example`, func() {
			// begin-get_data

			getDataOptions := schematics20ApiService.NewGetDataOptions(
				"testString",
			)

			previewData, response, err := schematics20ApiService.GetData(getDataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(previewData, "", "  ")
			fmt.Println(string(b))

			// end-get_data

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(previewData).ToNot(BeNil())

		})
		It(`ListCart request example`, func() {
			// begin-list_cart

			listCartOptions := schematics20ApiService.NewListCartOptions()

			cartOrderList, response, err := schematics20ApiService.ListCart(listCartOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cartOrderList, "", "  ")
			fmt.Println(string(b))

			// end-list_cart

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cartOrderList).ToNot(BeNil())

		})
		It(`CreateCartOrder request example`, func() {
			// begin-create_cart_order

			cartOrderModel := &schematics20apiv2.CartOrder{
				Name: core.StringPtr("MySatelliteLocation"),
			}

			createCartOrderOptions := schematics20ApiService.NewCreateCartOrderOptions(
				cartOrderModel,
			)

			cartOrder, response, err := schematics20ApiService.CreateCartOrder(createCartOrderOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cartOrder, "", "  ")
			fmt.Println(string(b))

			// end-create_cart_order

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cartOrder).ToNot(BeNil())

		})
		It(`GetCartOrder request example`, func() {
			// begin-get_cart_order

			getCartOrderOptions := schematics20ApiService.NewGetCartOrderOptions(
				"testString",
			)

			cartOrder, response, err := schematics20ApiService.GetCartOrder(getCartOrderOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cartOrder, "", "  ")
			fmt.Println(string(b))

			// end-get_cart_order

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cartOrder).ToNot(BeNil())

		})
		It(`ReplaceCartOrder request example`, func() {
			// begin-replace_cart_order

			cartOrderModel := &schematics20apiv2.CartOrder{
				Name: core.StringPtr("MySatelliteLocation"),
			}

			replaceCartOrderOptions := schematics20ApiService.NewReplaceCartOrderOptions(
				"testString",
				"testString",
				cartOrderModel,
			)

			cartOrder, response, err := schematics20ApiService.ReplaceCartOrder(replaceCartOrderOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cartOrder, "", "  ")
			fmt.Println(string(b))

			// end-replace_cart_order

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cartOrder).ToNot(BeNil())

		})
		It(`RunFulfilmentOperation request example`, func() {
			// begin-run_fulfilment_operation

			runFulfilmentOperationOptions := schematics20ApiService.NewRunFulfilmentOperationOptions(
				"testString",
				"testString",
			)

			response, err := schematics20ApiService.RunFulfilmentOperation(runFulfilmentOperationOptions)
			if err != nil {
				panic(err)
			}

			// end-run_fulfilment_operation

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
		It(`UpdateCartOrder request example`, func() {
			// begin-update_cart_order

			updateCartOrderModel := &schematics20apiv2.UpdateCartOrder{}

			updateCartOrderOptions := schematics20ApiService.NewUpdateCartOrderOptions(
				"testString",
				"testString",
				updateCartOrderModel,
			)

			cartOrder, response, err := schematics20ApiService.UpdateCartOrder(updateCartOrderOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cartOrder, "", "  ")
			fmt.Println(string(b))

			// end-update_cart_order

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cartOrder).ToNot(BeNil())

		})
		It(`GetCartMetadata request example`, func() {
			// begin-get_cart_metadata

			getCartMetadataOptions := schematics20ApiService.NewGetCartMetadataOptions(
				"testString",
			)

			itemMetadata, response, err := schematics20ApiService.GetCartMetadata(getCartMetadataOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(itemMetadata, "", "  ")
			fmt.Println(string(b))

			// end-get_cart_metadata

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(itemMetadata).ToNot(BeNil())

		})
		It(`ListCartResources request example`, func() {
			// begin-list_cart_resources

			listCartResourcesOptions := schematics20ApiService.NewListCartResourcesOptions(
				"testString",
			)

			cloudResourceList, response, err := schematics20ApiService.ListCartResources(listCartResourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cloudResourceList, "", "  ")
			fmt.Println(string(b))

			// end-list_cart_resources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudResourceList).ToNot(BeNil())

		})
		It(`ListCartOrderJobs request example`, func() {
			// begin-list_cart_order_jobs

			listCartOrderJobsOptions := schematics20ApiService.NewListCartOrderJobsOptions(
				"testString",
			)

			jobList, response, err := schematics20ApiService.ListCartOrderJobs(listCartOrderJobsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobList, "", "  ")
			fmt.Println(string(b))

			// end-list_cart_order_jobs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobList).ToNot(BeNil())

		})
		It(`GetCartOrderJob request example`, func() {
			// begin-get_cart_order_job

			getCartOrderJobOptions := schematics20ApiService.NewGetCartOrderJobOptions(
				"testString",
				"testString",
			)

			jobLite, response, err := schematics20ApiService.GetCartOrderJob(getCartOrderJobOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobLite, "", "  ")
			fmt.Println(string(b))

			// end-get_cart_order_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobLite).ToNot(BeNil())

		})
		It(`GetCartOrderJobLog request example`, func() {
			// begin-get_cart_order_job_log

			getCartOrderJobLogOptions := schematics20ApiService.NewGetCartOrderJobLogOptions(
				"testString",
				"testString",
			)

			jobLog, response, err := schematics20ApiService.GetCartOrderJobLog(getCartOrderJobLogOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(jobLog, "", "  ")
			fmt.Println(string(b))

			// end-get_cart_order_job_log

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(jobLog).ToNot(BeNil())

		})
		It(`GetCartOrderJobResources request example`, func() {
			// begin-get_cart_order_job_resources

			getCartOrderJobResourcesOptions := schematics20ApiService.NewGetCartOrderJobResourcesOptions(
				"testString",
				"testString",
			)

			cloudResourceList, response, err := schematics20ApiService.GetCartOrderJobResources(getCartOrderJobResourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cloudResourceList, "", "  ")
			fmt.Println(string(b))

			// end-get_cart_order_job_resources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cloudResourceList).ToNot(BeNil())

		})
		It(`ListBlueprint request example`, func() {
			// begin-list_blueprint

			listBlueprintOptions := schematics20ApiService.NewListBlueprintOptions()

			blueprintList, response, err := schematics20ApiService.ListBlueprint(listBlueprintOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(blueprintList, "", "  ")
			fmt.Println(string(b))

			// end-list_blueprint

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprintList).ToNot(BeNil())

		})
		It(`CreateBlueprint request example`, func() {
			// begin-create_blueprint

			blueprintModel := &schematics20apiv2.Blueprint{
				Name: core.StringPtr("Toronto Dev Environtment"),
			}

			createBlueprintOptions := schematics20ApiService.NewCreateBlueprintOptions(
				blueprintModel,
			)

			blueprint, response, err := schematics20ApiService.CreateBlueprint(createBlueprintOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(blueprint, "", "  ")
			fmt.Println(string(b))

			// end-create_blueprint

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprint).ToNot(BeNil())

		})
		It(`GetBlueprint request example`, func() {
			// begin-get_blueprint

			getBlueprintOptions := schematics20ApiService.NewGetBlueprintOptions(
				"testString",
			)

			blueprint, response, err := schematics20ApiService.GetBlueprint(getBlueprintOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(blueprint, "", "  ")
			fmt.Println(string(b))

			// end-get_blueprint

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprint).ToNot(BeNil())

		})
		It(`ReplaceBlueprint request example`, func() {
			// begin-replace_blueprint

			blueprintModel := &schematics20apiv2.Blueprint{
				Name: core.StringPtr("Toronto Dev Environtment"),
			}

			replaceBlueprintOptions := schematics20ApiService.NewReplaceBlueprintOptions(
				"testString",
				blueprintModel,
			)

			blueprint, response, err := schematics20ApiService.ReplaceBlueprint(replaceBlueprintOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(blueprint, "", "  ")
			fmt.Println(string(b))

			// end-replace_blueprint

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprint).ToNot(BeNil())

		})
		It(`UpdateBlueprint request example`, func() {
			// begin-update_blueprint

			blueprintModel := &schematics20apiv2.Blueprint{
				Name: core.StringPtr("Toronto Dev Environtment"),
			}

			updateBlueprintOptions := schematics20ApiService.NewUpdateBlueprintOptions(
				"testString",
				blueprintModel,
			)

			blueprint, response, err := schematics20ApiService.UpdateBlueprint(updateBlueprintOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(blueprint, "", "  ")
			fmt.Println(string(b))

			// end-update_blueprint

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(blueprint).ToNot(BeNil())

		})
		It(`UploadTemplateTarBlueprint request example`, func() {
			// begin-upload_template_tar_blueprint

			uploadTemplateTarBlueprintOptions := schematics20ApiService.NewUploadTemplateTarBlueprintOptions(
				"testString",
			)

			templateRepoTarUploadResponse, response, err := schematics20ApiService.UploadTemplateTarBlueprint(uploadTemplateTarBlueprintOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateRepoTarUploadResponse, "", "  ")
			fmt.Println(string(b))

			// end-upload_template_tar_blueprint

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templateRepoTarUploadResponse).ToNot(BeNil())

		})
		It(`ListCatalogItems request example`, func() {
			// begin-list_catalog_items

			listCatalogItemsOptions := schematics20ApiService.NewListCatalogItemsOptions()

			catalogOfferingItem, response, err := schematics20ApiService.ListCatalogItems(listCatalogItemsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogOfferingItem, "", "  ")
			fmt.Println(string(b))

			// end-list_catalog_items

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogOfferingItem).ToNot(BeNil())

		})
		It(`GetCatalogItem request example`, func() {
			// begin-get_catalog_item

			getCatalogItemOptions := schematics20ApiService.NewGetCatalogItemOptions(
				"testString",
			)

			catalogOfferingItem, response, err := schematics20ApiService.GetCatalogItem(getCatalogItemOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(catalogOfferingItem, "", "  ")
			fmt.Println(string(b))

			// end-get_catalog_item

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(catalogOfferingItem).ToNot(BeNil())

		})
		It(`ListLocations request example`, func() {
			// begin-list_locations

			listLocationsOptions := schematics20ApiService.NewListLocationsOptions()

			schematicsLocationsList, response, err := schematics20ApiService.ListLocations(listLocationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(schematicsLocationsList, "", "  ")
			fmt.Println(string(b))

			// end-list_locations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(schematicsLocationsList).ToNot(BeNil())

		})
		It(`DeleteWorkspace request example`, func() {
			// begin-delete_workspace

			deleteWorkspaceOptions := schematics20ApiService.NewDeleteWorkspaceOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteWorkspace(deleteWorkspaceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteWebhook request example`, func() {
			// begin-delete_webhook

			deleteWebhookOptions := schematics20ApiService.NewDeleteWebhookOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteWebhook(deleteWebhookOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_webhook

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTrigger request example`, func() {
			// begin-delete_trigger

			deleteTriggerOptions := schematics20ApiService.NewDeleteTriggerOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteTrigger(deleteTriggerOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_trigger

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteTemplateFromWorkspace request example`, func() {
			// begin-delete_template_from_workspace

			deleteTemplateFromWorkspaceOptions := schematics20ApiService.NewDeleteTemplateFromWorkspaceOptions(
				"testString",
				"testString",
			)

			response, err := schematics20ApiService.DeleteTemplateFromWorkspace(deleteTemplateFromWorkspaceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_template_from_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteResourcesQuery request example`, func() {
			// begin-delete_resources_query

			deleteResourcesQueryOptions := schematics20ApiService.NewDeleteResourcesQueryOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteResourcesQuery(deleteResourcesQueryOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_resources_query

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeletePrivateCluster request example`, func() {
			// begin-delete_private_cluster

			deletePrivateClusterOptions := schematics20ApiService.NewDeletePrivateClusterOptions()

			response, err := schematics20ApiService.DeletePrivateCluster(deletePrivateClusterOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_private_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteJob request example`, func() {
			// begin-delete_job

			deleteJobOptions := schematics20ApiService.NewDeleteJobOptions(
				"testString",
				"testString",
			)

			response, err := schematics20ApiService.DeleteJob(deleteJobOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInventory request example`, func() {
			// begin-delete_inventory

			deleteInventoryOptions := schematics20ApiService.NewDeleteInventoryOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteInventory(deleteInventoryOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_inventory

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteFlowFromWorkspace request example`, func() {
			// begin-delete_flow_from_workspace

			deleteFlowFromWorkspaceOptions := schematics20ApiService.NewDeleteFlowFromWorkspaceOptions(
				"testString",
				"testString",
			)

			response, err := schematics20ApiService.DeleteFlowFromWorkspace(deleteFlowFromWorkspaceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_flow_from_workspace

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteDataset request example`, func() {
			// begin-delete_dataset

			deleteDatasetOptions := schematics20ApiService.NewDeleteDatasetOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteDataset(deleteDatasetOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_dataset

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteControls request example`, func() {
			// begin-delete_controls

			deleteControlsOptions := schematics20ApiService.NewDeleteControlsOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteControls(deleteControlsOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_controls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteConnection request example`, func() {
			// begin-delete_connection

			deleteConnectionOptions := schematics20ApiService.NewDeleteConnectionOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteConnection(deleteConnectionOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteCartOrder request example`, func() {
			// begin-delete_cart_order

			deleteCartOrderOptions := schematics20ApiService.NewDeleteCartOrderOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteCartOrder(deleteCartOrderOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_cart_order

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteCapsule request example`, func() {
			// begin-delete_capsule

			deleteCapsuleOptions := schematics20ApiService.NewDeleteCapsuleOptions(
				"testString",
				"testString",
			)

			response, err := schematics20ApiService.DeleteCapsule(deleteCapsuleOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_capsule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteBlueprint request example`, func() {
			// begin-delete_blueprint

			deleteBlueprintOptions := schematics20ApiService.NewDeleteBlueprintOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteBlueprint(deleteBlueprintOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_blueprint

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteAgent request example`, func() {
			// begin-delete_agent

			deleteAgentOptions := schematics20ApiService.NewDeleteAgentOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteAgent(deleteAgentOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_agent

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteAdapter request example`, func() {
			// begin-delete_adapter

			deleteAdapterOptions := schematics20ApiService.NewDeleteAdapterOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteAdapter(deleteAdapterOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_adapter

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteAction request example`, func() {
			// begin-delete_action

			deleteActionOptions := schematics20ApiService.NewDeleteActionOptions(
				"testString",
			)

			response, err := schematics20ApiService.DeleteAction(deleteActionOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
