package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/litmuschaos/litmus/chaoscenter/graphql/server/graph/model"
	"github.com/litmuschaos/litmus/chaoscenter/graphql/server/pkg/authorization"
	"github.com/sirupsen/logrus"
)

func (r *mutationResolver) AddProbe(ctx context.Context, request model.ProbeRequest, projectID string) (*model.Probe, error) {
	logFields := logrus.Fields{
		"projectId": projectID,
	}

	logrus.WithFields(logFields).Info("request received to create a probe")

	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.AddProbe],
		model.InvitationAccepted.String())
	if err != nil {
		return nil, err
	}

	// Validate if probe type and probe properties match for the selected probe
	// HTTP Probe type and Property validation
	if request.Type == model.ProbeTypeHTTPProbe && request.KubernetesHTTPProperties == nil {
		err := "probe type and properties don't match, selected http Probe but http properties are empty"
		logrus.WithFields(logFields).Error(err)
		return nil, errors.New(err)
	}
	// PROM Probe type and Property validation
	if request.Type == model.ProbeTypePromProbe && request.PromProperties == nil {
		err := "probe type and properties don't match, selected prom Probe but prom properties are empty"
		logrus.WithFields(logFields).Error(err)
		return nil, errors.New(err)
	}
	// CMD Probe type and Property validation
	if request.Type == model.ProbeTypeCmdProbe && request.KubernetesCMDProperties == nil {
		err := "probe type and properties don't match, selected cmd Probe but cmd properties are empty"
		logrus.WithFields(logFields).Error(err)
		return nil, errors.New(err)
	}
	// K8S Probe type and Property validation
	if request.Type == model.ProbeTypeK8sProbe && request.K8sProperties == nil {
		err := "probe type and properties don't match, selected k8s Probe but k8s properties are empty"
		logrus.WithFields(logFields).Error(err)
		return nil, errors.New(err)
	}

	response, err := r.probeService.AddProbe(ctx, request, projectID)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return nil, err
	}

	return response, err
}

func (r *mutationResolver) UpdateProbe(ctx context.Context, request model.ProbeRequest, projectID string) (string, error) {
	logFields := logrus.Fields{
		"projectId": projectID,
		"probeName": request.Name,
	}

	logrus.WithFields(logFields).WithField("probeID", request.Name).Info("request received to update probe")

	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.UpdateProbe],
		model.InvitationAccepted.String())
	if err != nil {
		return "", err
	}

	response, err := r.probeService.UpdateProbe(ctx, request, projectID)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return "", err
	}

	return response, err
}

func (r *mutationResolver) DeleteProbe(ctx context.Context, probeName string, projectID string) (bool, error) {
	logFields := logrus.Fields{
		"projectId": projectID,
		"probeName": probeName,
	}

	logrus.WithFields(logFields).WithField("probeID", probeName).Info("request received to delete a probe")

	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.DeleteProbe],
		model.InvitationAccepted.String())
	if err != nil {
		return false, err
	}

	response, err := r.probeService.DeleteProbe(ctx, probeName, projectID)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return false, err
	}

	return response, err
}

func (r *queryResolver) ListProbes(ctx context.Context, projectID string, infrastructureType *model.InfrastructureType, probeNames []string, filter *model.ProbeFilterInput) ([]*model.Probe, error) {
	logFields := logrus.Fields{
		"projectId":  projectID,
		"probeNames": probeNames,
	}

	logrus.WithFields(logFields).Info("request received to get probes")
	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.ListProbes],
		model.InvitationAccepted.String())
	if err != nil {
		return nil, err
	}

	response, err := r.probeService.ListProbes(ctx, probeNames, infrastructureType, filter, projectID)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return nil, err
	}

	return response, err
}

func (r *queryResolver) GetProbe(ctx context.Context, projectID string, probeName string) (*model.Probe, error) {
	logFields := logrus.Fields{
		"projectId": projectID,
		"probeName": probeName,
	}

	logrus.WithFields(logFields).Info("request received to get probe")

	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.GetProbe],
		model.InvitationAccepted.String())
	if err != nil {
		return nil, err
	}

	response, err := r.probeService.GetProbe(ctx, probeName, projectID)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return nil, err
	}

	return response, err
}

func (r *queryResolver) GetProbeYaml(ctx context.Context, projectID string, request model.GetProbeYAMLRequest) (string, error) {
	logFields := logrus.Fields{
		"projectId": projectID,
		"probeName": request.ProbeName,
	}

	logrus.WithFields(logFields).Info("request received to get probe YAML")

	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.GetProbe],
		model.InvitationAccepted.String())
	if err != nil {
		return "", err
	}

	response, err := r.probeService.GetProbeYAMLData(ctx, request, projectID)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return "", err
	}

	return response, err
}

func (r *queryResolver) GetProbeReference(ctx context.Context, projectID string, probeName string) (*model.GetProbeReferenceResponse, error) {
	logFields := logrus.Fields{
		"projectId": projectID,
		"probeName": probeName,
	}

	logrus.WithFields(logFields).Info("request received to get probe references")

	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.GetProbe],
		model.InvitationAccepted.String())
	if err != nil {
		return nil, err
	}

	response, err := r.probeService.GetProbeReference(ctx, probeName, projectID)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return nil, err
	}

	return response, err
}

func (r *queryResolver) GetProbesInExperimentRun(ctx context.Context, projectID string, experimentRunID string, faultName string) ([]*model.GetProbesInExperimentRunResponse, error) {
	logFields := logrus.Fields{
		"projectId":       projectID,
		"experimentRunId": experimentRunID,
	}

	logrus.WithFields(logFields).Info("request received to get probes of the experiment run")

	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.GetProbe],
		model.InvitationAccepted.String())
	if err != nil {
		return nil, err
	}

	response, err := r.chaosExperimentHandler.GetProbesInExperimentRun(ctx, projectID, experimentRunID, faultName)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return nil, err
	}

	return response, err
}

func (r *queryResolver) ValidateUniqueProbe(ctx context.Context, projectID string, probeName string) (bool, error) {
	logFields := logrus.Fields{
		"projectId": projectID,
		"probeName": probeName,
	}

	logrus.WithFields(logFields).Info("request received to validate probe uniqueness")

	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.GetProbe],
		model.InvitationAccepted.String())
	if err != nil {
		return false, err
	}

	response, err := r.probeService.ValidateUniqueProbe(ctx, probeName, projectID)
	if err != nil {
		logrus.WithFields(logFields).Error(err)
		return false, err
	}

	return response, err
}
