package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/litmuschaos/litmus/chaoscenter/graphql/server/graph/model"
	"github.com/litmuschaos/litmus/chaoscenter/graphql/server/pkg/authorization"
	"github.com/sirupsen/logrus"
)

func (r *mutationResolver) CreateImageRegistry(ctx context.Context, projectID string, imageRegistryInfo model.ImageRegistryInput) (*model.ImageRegistryResponse, error) {
	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.CreateImageRegistry],
		model.InvitationAccepted.String())
	if err != nil {
		return nil, err
	}

	ciResponse, err := r.imageRegistryService.CreateImageRegistry(ctx, projectID, imageRegistryInfo)
	if err != nil {
		logrus.Error(err)
	}
	return ciResponse, err
}

func (r *mutationResolver) UpdateImageRegistry(ctx context.Context, imageRegistryID string, projectID string, imageRegistryInfo model.ImageRegistryInput) (*model.ImageRegistryResponse, error) {
	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.UpdateImageRegistry],
		model.InvitationAccepted.String())
	if err != nil {
		return nil, err
	}

	uiRegistry, err := r.imageRegistryService.UpdateImageRegistry(ctx, imageRegistryID, projectID, imageRegistryInfo)
	if err != nil {
		logrus.Error(err)
	}

	return uiRegistry, err
}

func (r *mutationResolver) DeleteImageRegistry(ctx context.Context, imageRegistryID string, projectID string) (string, error) {
	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.DeleteImageRegistry],
		model.InvitationAccepted.String())
	if err != nil {
		return "", err
	}

	diRegistry, err := r.imageRegistryService.DeleteImageRegistry(ctx, imageRegistryID, projectID)
	if err != nil {
		logrus.Error(err)
	}

	return diRegistry, err
}

func (r *queryResolver) ListImageRegistry(ctx context.Context, projectID string) ([]*model.ImageRegistryResponse, error) {
	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.ListImageRegistry],
		model.InvitationAccepted.String())
	if err != nil {
		return nil, err
	}

	imageRegistries, err := r.imageRegistryService.ListImageRegistries(ctx, projectID)
	if err != nil {
		logrus.Error(err)
	}

	return imageRegistries, err
}

func (r *queryResolver) GetImageRegistry(ctx context.Context, projectID string) (*model.ImageRegistryResponse, error) {
	err := authorization.ValidateRole(ctx, projectID,
		authorization.MutationRbacRules[authorization.GetImageRegistry],
		model.InvitationAccepted.String())
	if err != nil {
		return nil, err
	}

	imageRegistry, err := r.imageRegistryService.GetImageRegistry(ctx, projectID)
	if err != nil {
		logrus.Error(err)
	}

	return imageRegistry, err
}
