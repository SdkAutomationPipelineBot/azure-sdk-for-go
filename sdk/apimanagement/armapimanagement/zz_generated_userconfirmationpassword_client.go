// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// UserConfirmationPasswordClient contains the methods for the UserConfirmationPassword group.
// Don't use this type directly, use NewUserConfirmationPasswordClient() instead.
type UserConfirmationPasswordClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewUserConfirmationPasswordClient creates a new instance of UserConfirmationPasswordClient with the specified values.
func NewUserConfirmationPasswordClient(con *armcore.Connection, subscriptionID string) *UserConfirmationPasswordClient {
	return &UserConfirmationPasswordClient{con: con, subscriptionID: subscriptionID}
}

// Send - Sends confirmation
// If the operation fails it returns the *ErrorResponse error type.
func (client *UserConfirmationPasswordClient) Send(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserConfirmationPasswordSendOptions) (UserConfirmationPasswordSendResponse, error) {
	req, err := client.sendCreateRequest(ctx, resourceGroupName, serviceName, userID, options)
	if err != nil {
		return UserConfirmationPasswordSendResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return UserConfirmationPasswordSendResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return UserConfirmationPasswordSendResponse{}, client.sendHandleError(resp)
	}
	return UserConfirmationPasswordSendResponse{RawResponse: resp.Response}, nil
}

// sendCreateRequest creates the Send request.
func (client *UserConfirmationPasswordClient) sendCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, userID string, options *UserConfirmationPasswordSendOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/users/{userId}/confirmations/password/send"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if userID == "" {
		return nil, errors.New("parameter userID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{userId}", url.PathEscape(userID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	if options != nil && options.AppType != nil {
		reqQP.Set("appType", string(*options.AppType))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// sendHandleError handles the Send error response.
func (client *UserConfirmationPasswordClient) sendHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
