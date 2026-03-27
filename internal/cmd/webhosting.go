// SPDX-FileCopyrightText: 2025 OVH SAS <opensource@ovh.net>
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/ovh/ovhcloud-cli/internal/services/webhosting"
	"github.com/spf13/cobra"
)

func init() {
	webhostingCmd := &cobra.Command{
		Use:   "webhosting",
		Short: "Retrieve information and manage your WebHosting services",
	}

	// Command to list WebHosting services
	webhostingListCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List your WebHosting services",
		Run:     webhosting.ListWebHosting,
	}
	webhostingCmd.AddCommand(withFilterFlag(webhostingListCmd))

	// Command to get a single WebHosting
	webhostingCmd.AddCommand(&cobra.Command{
		Use:   "get <service_name>",
		Short: "Retrieve information of a specific WebHosting",
		Args:  cobra.ExactArgs(1),
		Run:   webhosting.GetWebHosting,
	})

	// Command to update a single WebHosting
	webhostingEditCmd := &cobra.Command{
		Use:   "edit <service_name>",
		Short: "Edit the given WebHosting",
		Args:  cobra.ExactArgs(1),
		Run:   webhosting.EditWebHosting,
	}
	webhostingEditCmd.Flags().StringVar(&webhosting.WebHostingSpec.DisplayName, "display-name", "", "Display name of the WebHosting")
	addInteractiveEditorFlag(webhostingEditCmd)
	webhostingCmd.AddCommand(webhostingEditCmd)

	// Command to get abuse status of a single WebHosting
	webhostingCmd.AddCommand(&cobra.Command{
		Use:   "abuse <service_name>",
		Short: "get abuse status of a single WebHosting",
		Args:  cobra.ExactArgs(1),
		Run:   webhosting.GetAbuseStatus,
	})

	// Command to manage attached domains of a single WebHosting
	webhostingDomainCmd := &cobra.Command{
		Use:     "domain",
		Aliases: []string{"attached-domains"},
		Short:   "Manage attached domains of a WebHosting",
	}
	// Command to list attached domains of a single WebHosting
	webhostingDomainCmd.AddCommand(&cobra.Command{
		Use:   "list <service_name>",
		Short: "List attached domains of a WebHosting",
		Run:   webhosting.ListAttachedDomains,
	})
	webhostingDomainCmd.PersistentFlags().StringVarP(&webhosting.WebHostingServiceName, "service-name", "s", "", "Filter by service name (only one string allowed)")
	webhostingDomainCmd.MarkPersistentFlagRequired("service-name")

	// Command to get details about a specific attached domain of a WebHosting
	webhostingDomainCmd.AddCommand(&cobra.Command{
		Use:   "get <domain>",
		Short: "Get details about attached domains of a WebHosting service-name",
		Args:  cobra.ExactArgs(1),
		Run:   webhosting.GetAttachedDomainDetails,
	})

	// Command to get DNS dig status of a specific attached domain of a WebHosting
	webhostingDomainCmd.AddCommand(&cobra.Command{
		Use:   "dig-status <domain>",
		Short: "Display detailed DNS dig status for an attached domain",
		Args:  cobra.ExactArgs(1),
		Run:   webhosting.GetAttachedDomainDNSStatus,
	})

	// Command to create a new attached domain of a WebHosting
	webhostingDomainCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new attached domain",
		Long:  `TODO: Add long description`,
		Run:   webhosting.CreateAttachedDomain,
		Args:  cobra.NoArgs,
	}
	// common flags for other means to define parameters
	// addParameterFileFlags(webhostingDomainCreateCmd, false, assets.WebhostingOpenapiSchema, "/hosting/web/{serviceName}/attachedDomain", "post", webhosting.AttachedDomainCreateExample, nil)
	// addInteractiveEditorFlag(webhostingDomainCreateCmd)
	// markFlagsMutuallyExclusive(webhostingDomainCreateCmd, "from-file", "editor")
	webhostingDomainCmd.AddCommand(webhostingDomainCreateCmd)

	webhostingCmd.AddCommand(webhostingDomainCmd)

	rootCmd.AddCommand(webhostingCmd)
}
