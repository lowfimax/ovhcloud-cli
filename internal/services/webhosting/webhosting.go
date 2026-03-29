// SPDX-FileCopyrightText: 2025 OVH SAS <opensource@ovh.net>
//
// SPDX-License-Identifier: Apache-2.0

package webhosting

import (
	_ "embed"
	"fmt"
	"net/url"

	"github.com/ovh/ovhcloud-cli/internal/assets"
	"github.com/ovh/ovhcloud-cli/internal/display"
	"github.com/ovh/ovhcloud-cli/internal/flags"
	"github.com/ovh/ovhcloud-cli/internal/services/common"
	"github.com/spf13/cobra"
)

var (
	webhostingColumnsToDisplay = []string{"serviceName", "displayName", "datacenter", "state"}

	//go:embed templates/webhosting.tmpl
	webhostingTemplate string

	//go:embed templates/abuse-state.tmpl
	abuseStateTemplate string

	//go:embed templates/service-name-domain.tmpl
	serviceNameDomainTemplate string

	//go:embed templates/dig-status-domain.tmpl
	digStatusDomainTemplate string

	//go:embed parameter-samples/domain-create.json
	AttachedDomainCreateExample string

	WebHostingSpec struct {
		DisplayName string `json:"displayName,omitempty"`
	}
	WebHostingServiceName string

	AttachedDomainCreateFlags struct {
		CDN                    string `json:"cdn,omitempty"`
		Domain                 string `json:"domain,omitempty"`
		Firewall               bool   `json:"firewall,omitempty"`
		OwnLog                 string `json:"ownLog,omitempty"`
		Path                   string `json:"path,omitempty"`
		SSL                    bool   `json:"ssl,omitempty"`
		BypassDNSConfiguration bool   `json:"bypassDNSConfiguration,omitempty"`
	}
)

func ListWebHosting(_ *cobra.Command, _ []string) {
	common.ManageListRequest("/v1/hosting/web", "", webhostingColumnsToDisplay, flags.GenericFilters)
}

func GetWebHosting(_ *cobra.Command, args []string) {
	common.ManageObjectRequest("/v1/hosting/web", args[0], webhostingTemplate)
}

func EditWebHosting(cmd *cobra.Command, args []string) {
	if err := common.EditResource(
		cmd,
		"/hosting/web/{serviceName}",
		fmt.Sprintf("/v1/hosting/web/%s", url.PathEscape(args[0])),
		WebHostingSpec,
		assets.WebhostingOpenapiSchema,
	); err != nil {
		display.OutputError(&flags.OutputFormatConfig, "%s", err)
		return
	}
}

func GetAbuseStatus(_ *cobra.Command, args []string) {
	common.ManageObjectRequestUntouchedURL(
		fmt.Sprintf("/v1/hosting/web/%s/abuseState", url.PathEscape(args[0])), args[0], abuseStateTemplate,
	)
}

func ListAttachedDomains(_ *cobra.Command, _ []string) {
	if WebHostingServiceName == "" {
		display.OutputError(&flags.OutputFormatConfig, "service name is required")
		return
	}
	common.ManageListRequest(
		fmt.Sprintf("/v1/hosting/web/%s/attachedDomain", url.PathEscape(WebHostingServiceName)), "", []string{"domain"}, flags.GenericFilters,
	)
}

func GetAttachedDomainDetails(_ *cobra.Command, args []string) {
	if WebHostingServiceName == "" {
		display.OutputError(&flags.OutputFormatConfig, "service name is required")
		return
	}
	common.ManageObjectRequestUntouchedURL(
		fmt.Sprintf("/v1/hosting/web/%s/attachedDomain/%s", url.PathEscape(WebHostingServiceName), url.PathEscape(WebHostingServiceName)), args[0], serviceNameDomainTemplate,
	)
}

func GetAttachedDomainDNSStatus(_ *cobra.Command, args []string) {
	if WebHostingServiceName == "" {
		display.OutputError(&flags.OutputFormatConfig, "service name is required")
		return
	}
	common.ManageObjectRequestUntouchedURL(
		fmt.Sprintf("/v1/hosting/web/%s/attachedDomain/%s/digStatus", url.PathEscape(WebHostingServiceName), url.PathEscape(args[0])), args[0], digStatusDomainTemplate,
	)
}

func CreateAttachedDomain(cmd *cobra.Command, args []string) {
	createdDomain, err := common.CreateResource(
		cmd,
		"/hosting/web/{serviceName}/attachedDomain",
		fmt.Sprintf("/v1/hosting/web/%s/attachedDomain", url.PathEscape(WebHostingServiceName)),
		AttachedDomainCreateExample,
		AttachedDomainCreateFlags,
		assets.WebhostingOpenapiSchema,
		[]string{"domain"},
	)
	if err != nil {
		display.OutputError(&flags.OutputFormatConfig, "%s", err)
		return
	}
	display.OutputInfo(&flags.OutputFormatConfig, createdDomain,
		fmt.Sprintf("✅ Attached domain '%s' to service '%s' in progress...", AttachedDomainCreateFlags.Domain, WebHostingServiceName))
}
