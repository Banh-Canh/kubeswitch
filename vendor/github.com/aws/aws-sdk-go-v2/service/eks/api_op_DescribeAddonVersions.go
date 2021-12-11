// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the Kubernetes versions that the add-on can be used with.
func (c *Client) DescribeAddonVersions(ctx context.Context, params *DescribeAddonVersionsInput, optFns ...func(*Options)) (*DescribeAddonVersionsOutput, error) {
	if params == nil {
		params = &DescribeAddonVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAddonVersions", params, optFns, c.addOperationDescribeAddonVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAddonVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAddonVersionsInput struct {

	// The name of the add-on. The name must match one of the names returned by
	// ListAddons
	// (https://docs.aws.amazon.com/eks/latest/APIReference/API_ListAddons.html).
	AddonName *string

	// The Kubernetes versions that the add-on can be used with.
	KubernetesVersion *string

	// The maximum number of results to return.
	MaxResults *int32

	// The nextToken value returned from a previous paginated
	// DescribeAddonVersionsRequest where maxResults was used and the results exceeded
	// the value of that parameter. Pagination continues from the end of the previous
	// results that returned the nextToken value. This token should be treated as an
	// opaque identifier that is used only to retrieve the next items in a list and not
	// for other programmatic purposes.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAddonVersionsOutput struct {

	// The list of available versions with Kubernetes version compatibility.
	Addons []types.AddonInfo

	// The nextToken value returned from a previous paginated
	// DescribeAddonVersionsResponse where maxResults was used and the results exceeded
	// the value of that parameter. Pagination continues from the end of the previous
	// results that returned the nextToken value. This token should be treated as an
	// opaque identifier that is used only to retrieve the next items in a list and not
	// for other programmatic purposes.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAddonVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAddonVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAddonVersions{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAddonVersions(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeAddonVersionsAPIClient is a client that implements the
// DescribeAddonVersions operation.
type DescribeAddonVersionsAPIClient interface {
	DescribeAddonVersions(context.Context, *DescribeAddonVersionsInput, ...func(*Options)) (*DescribeAddonVersionsOutput, error)
}

var _ DescribeAddonVersionsAPIClient = (*Client)(nil)

// DescribeAddonVersionsPaginatorOptions is the paginator options for
// DescribeAddonVersions
type DescribeAddonVersionsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAddonVersionsPaginator is a paginator for DescribeAddonVersions
type DescribeAddonVersionsPaginator struct {
	options   DescribeAddonVersionsPaginatorOptions
	client    DescribeAddonVersionsAPIClient
	params    *DescribeAddonVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeAddonVersionsPaginator returns a new DescribeAddonVersionsPaginator
func NewDescribeAddonVersionsPaginator(client DescribeAddonVersionsAPIClient, params *DescribeAddonVersionsInput, optFns ...func(*DescribeAddonVersionsPaginatorOptions)) *DescribeAddonVersionsPaginator {
	if params == nil {
		params = &DescribeAddonVersionsInput{}
	}

	options := DescribeAddonVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAddonVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAddonVersionsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeAddonVersions page.
func (p *DescribeAddonVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAddonVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeAddonVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeAddonVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "DescribeAddonVersions",
	}
}
