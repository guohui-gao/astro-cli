package cmd

import (
	"github.com/spf13/cobra"
)

var (
	workspaceRootCmd = &cobra.Command{
		Use:     "workspace",
		Aliases: []string{"wo", "ws"},
		Short:   "Manage Astronomer workspaces",
		Long:    "Workspaces contain a group of Airflow Cluster Deployments. The creator of the workspace can invite other users into it",
	}

	workspaceListCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List astronomer workspaces",
		Long:    "List astronomer workspaces",
		RunE:    workspaceList,
	}

	workspaceCreateCmd = &cobra.Command{
		Use:   "create",
		Short: "Create an astronomer workspaces",
		Long:  "Create an astronomer workspaces",
		RunE:  workspaceCreate,
	}

	workspaceDeleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an astronomer workspace",
		Long:  "Delete an astronomer workspace",
		RunE:  workspaceDelete,
	}

	workspaceUpdateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an Astronomer workspace",
		Long:  "Update a workspace name, as well as users and roles assigned to a workspace",
		RunE:  workspaceUpdate,
	}
)

func init() {
	// workspace root
	RootCmd.AddCommand(workspaceRootCmd)

	// workspace list
	workspaceRootCmd.AddCommand(workspaceListCmd)

	// workspace create
	workspaceRootCmd.AddCommand(workspaceCreateCmd)

	// workspace delete
	workspaceRootCmd.AddCommand(workspaceDeleteCmd)

	// workspace update
	workspaceRootCmd.AddCommand(workspaceUpdateCmd)
}

func workspaceCreate(cmd *cobra.Command, args []string) error {
	return nil
}

func workspaceList(cmd *cobra.Command, args []string) error {
	return nil
}

// TODO
func workspaceDelete(cmd *cobra.Command, args []string) error {
	return nil
}

// TODO
func workspaceUpdate(cmd *cobra.Command, args []string) error {
	return nil
}
