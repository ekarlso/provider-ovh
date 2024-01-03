/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this ProjectDatabase.
func (mg *ProjectDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabase.
func (mg *ProjectDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabase.
func (mg *ProjectDatabase) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabase.
func (mg *ProjectDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabase.
func (mg *ProjectDatabase) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabase.
func (mg *ProjectDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabase.
func (mg *ProjectDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabase.
func (mg *ProjectDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabase.
func (mg *ProjectDatabase) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabase.
func (mg *ProjectDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabase.
func (mg *ProjectDatabase) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabase.
func (mg *ProjectDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseDatabase.
func (mg *ProjectDatabaseDatabase) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseIPRestriction.
func (mg *ProjectDatabaseIPRestriction) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseIntegration.
func (mg *ProjectDatabaseIntegration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseKafkaACL.
func (mg *ProjectDatabaseKafkaACL) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseKafkaSchemaregistryacl.
func (mg *ProjectDatabaseKafkaSchemaregistryacl) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseKafkaTopic.
func (mg *ProjectDatabaseKafkaTopic) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseM3DbNamespace.
func (mg *ProjectDatabaseM3DbNamespace) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseM3DbUser.
func (mg *ProjectDatabaseM3DbUser) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseMongodbUser.
func (mg *ProjectDatabaseMongodbUser) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseOpensearchPattern.
func (mg *ProjectDatabaseOpensearchPattern) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseOpensearchUser.
func (mg *ProjectDatabaseOpensearchUser) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabasePostgresqlUser.
func (mg *ProjectDatabasePostgresqlUser) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseRedisUser.
func (mg *ProjectDatabaseRedisUser) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicies of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) GetManagementPolicies() xpv1.ManagementPolicies {
	return mg.Spec.ManagementPolicies
}

// GetProviderConfigReference of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

// GetPublishConnectionDetailsTo of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicies of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) SetManagementPolicies(r xpv1.ManagementPolicies) {
	mg.Spec.ManagementPolicies = r
}

// SetProviderConfigReference of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

// SetPublishConnectionDetailsTo of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this ProjectDatabaseUser.
func (mg *ProjectDatabaseUser) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
