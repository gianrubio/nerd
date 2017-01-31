// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package servicecatalog_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleServiceCatalog_AcceptPortfolioShare() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.AcceptPortfolioShareInput{
		PortfolioId:    aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.AcceptPortfolioShare(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_AssociatePrincipalWithPortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.AssociatePrincipalWithPortfolioInput{
		PortfolioId:    aws.String("Id"),            // Required
		PrincipalARN:   aws.String("PrincipalARN"),  // Required
		PrincipalType:  aws.String("PrincipalType"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.AssociatePrincipalWithPortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_AssociateProductWithPortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.AssociateProductWithPortfolioInput{
		PortfolioId:       aws.String("Id"), // Required
		ProductId:         aws.String("Id"), // Required
		AcceptLanguage:    aws.String("AcceptLanguage"),
		SourcePortfolioId: aws.String("Id"),
	}
	resp, err := svc.AssociateProductWithPortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_CreateConstraint() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.CreateConstraintInput{
		IdempotencyToken: aws.String("IdempotencyToken"),     // Required
		Parameters:       aws.String("ConstraintParameters"), // Required
		PortfolioId:      aws.String("Id"),                   // Required
		ProductId:        aws.String("Id"),                   // Required
		Type:             aws.String("ConstraintType"),       // Required
		AcceptLanguage:   aws.String("AcceptLanguage"),
		Description:      aws.String("ConstraintDescription"),
	}
	resp, err := svc.CreateConstraint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_CreatePortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.CreatePortfolioInput{
		DisplayName:      aws.String("PortfolioDisplayName"), // Required
		IdempotencyToken: aws.String("IdempotencyToken"),     // Required
		ProviderName:     aws.String("ProviderName"),         // Required
		AcceptLanguage:   aws.String("AcceptLanguage"),
		Description:      aws.String("PortfolioDescription"),
		Tags: []*servicecatalog.Tag{
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.CreatePortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_CreatePortfolioShare() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.CreatePortfolioShareInput{
		AccountId:      aws.String("AccountId"), // Required
		PortfolioId:    aws.String("Id"),        // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.CreatePortfolioShare(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_CreateProduct() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.CreateProductInput{
		IdempotencyToken: aws.String("IdempotencyToken"), // Required
		Name:             aws.String("ProductViewName"),  // Required
		Owner:            aws.String("ProductViewOwner"), // Required
		ProductType:      aws.String("ProductType"),      // Required
		ProvisioningArtifactParameters: &servicecatalog.ProvisioningArtifactProperties{ // Required
			Info: map[string]*string{ // Required
				"Key": aws.String("ProvisioningArtifactInfoValue"), // Required
				// More values...
			},
			Description: aws.String("ProvisioningArtifactDescription"),
			Name:        aws.String("ProvisioningArtifactName"),
			Type:        aws.String("ProvisioningArtifactType"),
		},
		AcceptLanguage:     aws.String("AcceptLanguage"),
		Description:        aws.String("ProductViewShortDescription"),
		Distributor:        aws.String("ProductViewOwner"),
		SupportDescription: aws.String("SupportDescription"),
		SupportEmail:       aws.String("SupportEmail"),
		SupportUrl:         aws.String("SupportUrl"),
		Tags: []*servicecatalog.Tag{
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.CreateProduct(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_CreateProvisioningArtifact() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.CreateProvisioningArtifactInput{
		IdempotencyToken: aws.String("IdempotencyToken"), // Required
		Parameters: &servicecatalog.ProvisioningArtifactProperties{ // Required
			Info: map[string]*string{ // Required
				"Key": aws.String("ProvisioningArtifactInfoValue"), // Required
				// More values...
			},
			Description: aws.String("ProvisioningArtifactDescription"),
			Name:        aws.String("ProvisioningArtifactName"),
			Type:        aws.String("ProvisioningArtifactType"),
		},
		ProductId:      aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.CreateProvisioningArtifact(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DeleteConstraint() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DeleteConstraintInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DeleteConstraint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DeletePortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DeletePortfolioInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DeletePortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DeletePortfolioShare() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DeletePortfolioShareInput{
		AccountId:      aws.String("AccountId"), // Required
		PortfolioId:    aws.String("Id"),        // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DeletePortfolioShare(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DeleteProduct() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DeleteProductInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DeleteProduct(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DeleteProvisioningArtifact() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DeleteProvisioningArtifactInput{
		ProductId:              aws.String("Id"), // Required
		ProvisioningArtifactId: aws.String("Id"), // Required
		AcceptLanguage:         aws.String("AcceptLanguage"),
	}
	resp, err := svc.DeleteProvisioningArtifact(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DescribeConstraint() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DescribeConstraintInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DescribeConstraint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DescribePortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DescribePortfolioInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DescribePortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DescribeProduct() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DescribeProductInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DescribeProduct(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DescribeProductAsAdmin() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DescribeProductAsAdminInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DescribeProductAsAdmin(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DescribeProductView() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DescribeProductViewInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DescribeProductView(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DescribeProvisioningArtifact() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DescribeProvisioningArtifactInput{
		ProductId:              aws.String("Id"), // Required
		ProvisioningArtifactId: aws.String("Id"), // Required
		AcceptLanguage:         aws.String("AcceptLanguage"),
	}
	resp, err := svc.DescribeProvisioningArtifact(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DescribeProvisioningParameters() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DescribeProvisioningParametersInput{
		ProductId:              aws.String("Id"), // Required
		ProvisioningArtifactId: aws.String("Id"), // Required
		AcceptLanguage:         aws.String("AcceptLanguage"),
		PathId:                 aws.String("Id"),
	}
	resp, err := svc.DescribeProvisioningParameters(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DescribeRecord() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DescribeRecordInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
		PageSize:       aws.Int64(1),
		PageToken:      aws.String("PageToken"),
	}
	resp, err := svc.DescribeRecord(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DisassociatePrincipalFromPortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DisassociatePrincipalFromPortfolioInput{
		PortfolioId:    aws.String("Id"),           // Required
		PrincipalARN:   aws.String("PrincipalARN"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DisassociatePrincipalFromPortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_DisassociateProductFromPortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.DisassociateProductFromPortfolioInput{
		PortfolioId:    aws.String("Id"), // Required
		ProductId:      aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.DisassociateProductFromPortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ListAcceptedPortfolioShares() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ListAcceptedPortfolioSharesInput{
		AcceptLanguage: aws.String("AcceptLanguage"),
		PageSize:       aws.Int64(1),
		PageToken:      aws.String("PageToken"),
	}
	resp, err := svc.ListAcceptedPortfolioShares(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ListConstraintsForPortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ListConstraintsForPortfolioInput{
		PortfolioId:    aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
		PageSize:       aws.Int64(1),
		PageToken:      aws.String("PageToken"),
		ProductId:      aws.String("Id"),
	}
	resp, err := svc.ListConstraintsForPortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ListLaunchPaths() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ListLaunchPathsInput{
		ProductId:      aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
		PageSize:       aws.Int64(1),
		PageToken:      aws.String("PageToken"),
	}
	resp, err := svc.ListLaunchPaths(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ListPortfolioAccess() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ListPortfolioAccessInput{
		PortfolioId:    aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.ListPortfolioAccess(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ListPortfolios() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ListPortfoliosInput{
		AcceptLanguage: aws.String("AcceptLanguage"),
		PageSize:       aws.Int64(1),
		PageToken:      aws.String("PageToken"),
	}
	resp, err := svc.ListPortfolios(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ListPortfoliosForProduct() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ListPortfoliosForProductInput{
		ProductId:      aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
		PageSize:       aws.Int64(1),
		PageToken:      aws.String("PageToken"),
	}
	resp, err := svc.ListPortfoliosForProduct(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ListPrincipalsForPortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ListPrincipalsForPortfolioInput{
		PortfolioId:    aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
		PageSize:       aws.Int64(1),
		PageToken:      aws.String("PageToken"),
	}
	resp, err := svc.ListPrincipalsForPortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ListProvisioningArtifacts() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ListProvisioningArtifactsInput{
		ProductId:      aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.ListProvisioningArtifacts(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ListRecordHistory() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ListRecordHistoryInput{
		AcceptLanguage: aws.String("AcceptLanguage"),
		AccessLevelFilter: &servicecatalog.AccessLevelFilter{
			Key:   aws.String("AccessLevelFilterKey"),
			Value: aws.String("AccessLevelFilterValue"),
		},
		PageSize:  aws.Int64(1),
		PageToken: aws.String("PageToken"),
		SearchFilter: &servicecatalog.ListRecordHistorySearchFilter{
			Key:   aws.String("SearchFilterKey"),
			Value: aws.String("SearchFilterValue"),
		},
	}
	resp, err := svc.ListRecordHistory(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ProvisionProduct() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ProvisionProductInput{
		ProductId:              aws.String("Id"),                     // Required
		ProvisionToken:         aws.String("IdempotencyToken"),       // Required
		ProvisionedProductName: aws.String("ProvisionedProductName"), // Required
		ProvisioningArtifactId: aws.String("Id"),                     // Required
		AcceptLanguage:         aws.String("AcceptLanguage"),
		NotificationArns: []*string{
			aws.String("NotificationArn"), // Required
			// More values...
		},
		PathId: aws.String("Id"),
		ProvisioningParameters: []*servicecatalog.ProvisioningParameter{
			{ // Required
				Key:   aws.String("ParameterKey"),
				Value: aws.String("ParameterValue"),
			},
			// More values...
		},
		Tags: []*servicecatalog.Tag{
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.ProvisionProduct(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_RejectPortfolioShare() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.RejectPortfolioShareInput{
		PortfolioId:    aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
	}
	resp, err := svc.RejectPortfolioShare(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_ScanProvisionedProducts() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.ScanProvisionedProductsInput{
		AcceptLanguage: aws.String("AcceptLanguage"),
		AccessLevelFilter: &servicecatalog.AccessLevelFilter{
			Key:   aws.String("AccessLevelFilterKey"),
			Value: aws.String("AccessLevelFilterValue"),
		},
		PageSize:  aws.Int64(1),
		PageToken: aws.String("PageToken"),
	}
	resp, err := svc.ScanProvisionedProducts(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_SearchProducts() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.SearchProductsInput{
		AcceptLanguage: aws.String("AcceptLanguage"),
		Filters: map[string][]*string{
			"Key": { // Required
				aws.String("ProductViewFilterValue"), // Required
				// More values...
			},
			// More values...
		},
		PageSize:  aws.Int64(1),
		PageToken: aws.String("PageToken"),
		SortBy:    aws.String("ProductViewSortBy"),
		SortOrder: aws.String("SortOrder"),
	}
	resp, err := svc.SearchProducts(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_SearchProductsAsAdmin() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.SearchProductsAsAdminInput{
		AcceptLanguage: aws.String("AcceptLanguage"),
		Filters: map[string][]*string{
			"Key": { // Required
				aws.String("ProductViewFilterValue"), // Required
				// More values...
			},
			// More values...
		},
		PageSize:      aws.Int64(1),
		PageToken:     aws.String("PageToken"),
		PortfolioId:   aws.String("Id"),
		ProductSource: aws.String("ProductSource"),
		SortBy:        aws.String("ProductViewSortBy"),
		SortOrder:     aws.String("SortOrder"),
	}
	resp, err := svc.SearchProductsAsAdmin(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_TerminateProvisionedProduct() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.TerminateProvisionedProductInput{
		TerminateToken:         aws.String("IdempotencyToken"), // Required
		AcceptLanguage:         aws.String("AcceptLanguage"),
		IgnoreErrors:           aws.Bool(true),
		ProvisionedProductId:   aws.String("Id"),
		ProvisionedProductName: aws.String("ProvisionedProductNameOrArn"),
	}
	resp, err := svc.TerminateProvisionedProduct(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_UpdateConstraint() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.UpdateConstraintInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
		Description:    aws.String("ConstraintDescription"),
	}
	resp, err := svc.UpdateConstraint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_UpdatePortfolio() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.UpdatePortfolioInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
		AddTags: []*servicecatalog.Tag{
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
		Description:  aws.String("PortfolioDescription"),
		DisplayName:  aws.String("PortfolioDisplayName"),
		ProviderName: aws.String("ProviderName"),
		RemoveTags: []*string{
			aws.String("TagKey"), // Required
			// More values...
		},
	}
	resp, err := svc.UpdatePortfolio(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_UpdateProduct() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.UpdateProductInput{
		Id:             aws.String("Id"), // Required
		AcceptLanguage: aws.String("AcceptLanguage"),
		AddTags: []*servicecatalog.Tag{
			{ // Required
				Key:   aws.String("TagKey"),   // Required
				Value: aws.String("TagValue"), // Required
			},
			// More values...
		},
		Description: aws.String("ProductViewShortDescription"),
		Distributor: aws.String("ProductViewOwner"),
		Name:        aws.String("ProductViewName"),
		Owner:       aws.String("ProductViewOwner"),
		RemoveTags: []*string{
			aws.String("TagKey"), // Required
			// More values...
		},
		SupportDescription: aws.String("SupportDescription"),
		SupportEmail:       aws.String("SupportEmail"),
		SupportUrl:         aws.String("SupportUrl"),
	}
	resp, err := svc.UpdateProduct(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_UpdateProvisionedProduct() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.UpdateProvisionedProductInput{
		UpdateToken:            aws.String("IdempotencyToken"), // Required
		AcceptLanguage:         aws.String("AcceptLanguage"),
		PathId:                 aws.String("Id"),
		ProductId:              aws.String("Id"),
		ProvisionedProductId:   aws.String("Id"),
		ProvisionedProductName: aws.String("ProvisionedProductNameOrArn"),
		ProvisioningArtifactId: aws.String("Id"),
		ProvisioningParameters: []*servicecatalog.UpdateProvisioningParameter{
			{ // Required
				Key:              aws.String("ParameterKey"),
				UsePreviousValue: aws.Bool(true),
				Value:            aws.String("ParameterValue"),
			},
			// More values...
		},
	}
	resp, err := svc.UpdateProvisionedProduct(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleServiceCatalog_UpdateProvisioningArtifact() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := servicecatalog.New(sess)

	params := &servicecatalog.UpdateProvisioningArtifactInput{
		ProductId:              aws.String("Id"), // Required
		ProvisioningArtifactId: aws.String("Id"), // Required
		AcceptLanguage:         aws.String("AcceptLanguage"),
		Description:            aws.String("ProvisioningArtifactDescription"),
		Name:                   aws.String("ProvisioningArtifactName"),
	}
	resp, err := svc.UpdateProvisioningArtifact(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
