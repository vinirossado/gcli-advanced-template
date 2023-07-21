package repository

import (
    "basic/source/model"
    "github.com/pkg/errors"
)

type V2Repository interface {
    GetV2ById(id uint) (*model.V2, error)
    GetAllV2() (*[]model.V2, error)
    CreateV2(v2 *model.V2) (uint, error)
    UpdateV2(v2 *model.V2) (uint, error)
    DeleteV2(id uint) (uint, error)
}

type v2Repository struct {
    *Repository
}

func NewV2Repository(repository *Repository) V2Repository {
    return &v2Repository{
         Repository: repository,
    }
}

func (r *v2Repository) GetAllV2() (*[]model.V2, error) {
    //TODO implement me
    panic("implement me")
}


func (r *v2Repository) GetV2ById(id uint) (*model.V2, error) {
    var v2 model.V2

    if err := r.db.Where("id = ?", id).First(&v2).Error; err != nil {
        return nil, errors.Wrap(err, "failed to get user by ID")
    }

    return &v2, nil
}

func (r *v2Repository) CreateV2(v2 *model.V2) (uint, error) {
    //TODO implement me
    panic("implement me")
}

func (r *v2Repository) UpdateV2(v2 *model.V2) (uint, error) {
    if err := r.db.Save(v2).Error; err != nil {
         return v2.ID, errors.Wrap(err, "failed to update user")
    }

    return v2.ID, nil
}

func (r *v2Repository) DeleteV2(id uint) (uint, error) {
    //TODO implement me
    panic("implement me")
}






















