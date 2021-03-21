// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-swagger/go-swagger/examples/cli/client/todos"
	"github.com/go-swagger/go-swagger/examples/cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTodosAddOneCmd returns a cmd to handle operation addOne
func makeOperationTodosAddOneCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addOne",
		Short: ``,
		RunE:  runOperationTodosAddOne,
	}

	if err := registerOperationTodosAddOneParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTodosAddOne uses cmd flags to call endpoint api
func runOperationTodosAddOne(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := todos.NewAddOneParams()
	if err, _ := retrieveOperationTodosAddOneBodyFlag(params, "", cmd); err != nil {
		return err
	}
	// make request and then print result
	if err := printOperationTodosAddOneResult(appCli.Todos.AddOne(params, nil)); err != nil {
		return err
	}
	return nil
}

func retrieveOperationTodosAddOneBodyFlag(m *todos.AddOneParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.Item{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.Item: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.Item{}
	}
	err, added := retrieveItemFlags(bodyValueModel, "item", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	bodyValueDebugBytes, err := json.Marshal(m.Body)
	if err != nil {
		return err, false
	}
	logDebugf("Body payload: %v", string(bodyValueDebugBytes))
	return nil, retAdded
}

// printOperationTodosAddOneResult prints output to stdout
func printOperationTodosAddOneResult(resp0 *todos.AddOneCreated, respErr error) error {
	if respErr != nil {

		var iResp interface{} = respErr
		defaultResp, ok := iResp.(*todos.AddOneDefault)
		if !ok {
			return respErr
		}
		if defaultResp.Payload != nil {
			msgStr, err := json.Marshal(defaultResp.Payload)
			if err != nil {
				return err
			}
			fmt.Println(string(msgStr))
			return nil
		}

		return respErr
	}

	if resp0.Payload != nil {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return err
		}
		fmt.Println(string(msgStr))
	}

	return nil
}

// registerOperationTodosAddOneParamFlags registers all flags needed to fill params
func registerOperationTodosAddOneParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTodosAddOneBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTodosAddOneBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	exampleBodyStr, err := json.Marshal(&models.Item{})
	if err != nil {
		return err
	}
	_ = cmd.PersistentFlags().String(bodyFlagName, "", fmt.Sprintf("Optional json string for [body] of form %v.", string(exampleBodyStr)))

	// add flags for body
	if err := registerItemFlags("item", cmd); err != nil {
		return err
	}

	return nil
}