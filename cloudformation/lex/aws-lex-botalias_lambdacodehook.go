package lex

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// BotAlias_LambdaCodeHook AWS CloudFormation Resource (AWS::Lex::BotAlias.LambdaCodeHook)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-lambdacodehook.html
type BotAlias_LambdaCodeHook struct {

	// CodeHookInterfaceVersion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-lambdacodehook.html#cfn-lex-botalias-lambdacodehook-codehookinterfaceversion
	CodeHookInterfaceVersion string `json:"CodeHookInterfaceVersion,omitempty"`

	// LambdaArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-botalias-lambdacodehook.html#cfn-lex-botalias-lambdacodehook-lambdaarn
	LambdaArn string `json:"LambdaArn,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *BotAlias_LambdaCodeHook) AWSCloudFormationType() string {
	return "AWS::Lex::BotAlias.LambdaCodeHook"
}
