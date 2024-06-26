package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"dario.cat/mergo"
	"github.com/PICT-LibraryAutomation/granthpal/graph"
	"github.com/PICT-LibraryAutomation/granthpal/remote/models"
	"github.com/PICT-LibraryAutomation/granthpal/utils"
)

// CreatePublication is the resolver for the createPublication field.
func (r *mutationResolver) CreatePublication(ctx context.Context, inp graph.CreatePublicationInp) (*graph.Publication, error) {
	p := models.Publication{
		ID:   inp.ID,
		Name: inp.Name,
	}

	if err := r.Remote.Save(&p).Error; err != nil {
		return nil, err
	}

	return p.ToGraphQL(), nil
}

// DeletePublication is the resolver for the deletePublication field.
func (r *mutationResolver) DeletePublication(ctx context.Context, id string) (*string, error) {
	if err := r.Remote.Delete(&models.Publication{ID: id}).Error; err != nil {
		return nil, err
	}

	return &id, nil
}

// UpdatePublication is the resolver for the updatePublication field.
func (r *mutationResolver) UpdatePublication(ctx context.Context, inp graph.UpdatePublicationInp) (*graph.Publication, error) {
	var p models.Publication
	if err := r.Remote.First(&p, "id = ?", inp.ID).Error; err != nil {
		return nil, err
	}

	if err := mergo.Merge(&p, inp, mergo.WithOverride); err != nil {
		return nil, err
	}

	if err := r.Remote.Save(&p).Error; err != nil {
		return nil, err
	}

	return p.ToGraphQL(), nil
}

// Books is the resolver for the books field.
func (r *publicationResolver) Books(ctx context.Context, obj *graph.Publication) ([]*graph.BookMetadata, error) {
	var bs []models.BookMetadata
	if err := r.Remote.First(&models.Publication{ID: obj.ID}).Association("Books").Find(&bs); err != nil {
		return nil, err
	}

	return utils.Map(bs, func(b models.BookMetadata) *graph.BookMetadata {
		return b.ToGraphQL()
	}), nil
}

// Publication is the resolver for the publication field.
func (r *queryResolver) Publication(ctx context.Context, id string) (*graph.Publication, error) {
	var p models.Publication
	if err := r.Remote.First(&p, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return p.ToGraphQL(), nil
}

// Publications is the resolver for the publications field.
func (r *queryResolver) Publications(ctx context.Context) ([]*graph.Publication, error) {
	var ps []models.Publication
	if err := r.Remote.Find(&ps).Error; err != nil {
		return nil, err
	}

	return utils.Map(ps, func(p models.Publication) *graph.Publication {
		return p.ToGraphQL()
	}), nil
}

// Publication returns graph.PublicationResolver implementation.
func (r *Resolver) Publication() graph.PublicationResolver { return &publicationResolver{r} }

type publicationResolver struct{ *Resolver }
