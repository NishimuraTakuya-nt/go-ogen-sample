package main

import (
	"context"
	"sync"

	petstore "github.com/NishimuraTakuya-nt/go-ogen-sample/petstore"
)

type handler struct {
	pets map[int64]petstore.Pet
	id   int64
	mux  sync.Mutex
}

func (p *handler) AddPet(ctx context.Context, req *petstore.Pet) (*petstore.Pet, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.pets[p.id] = *req
	p.id++
	return &petstore.Pet{
		ID: petstore.OptInt64{
			Value: p.id,
			Set:   true,
		},
		Name: req.Name,
	}, nil
}

func (p *handler) DeletePet(ctx context.Context, params petstore.DeletePetParams) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	delete(p.pets, params.PetId)
	return nil
}

func (p *handler) GetPetById(ctx context.Context, params petstore.GetPetByIdParams) (petstore.GetPetByIdRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet, ok := p.pets[params.PetId]
	if !ok {
		// Return Not Found.
		return &petstore.GetPetByIdNotFound{}, nil
	}
	return &pet, nil
}

func (p *handler) UpdatePet(ctx context.Context, params petstore.UpdatePetParams) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	pet := p.pets[params.PetId]
	pet.Status = params.Status
	if val, ok := params.Name.Get(); ok {
		pet.Name = val
	}
	p.pets[params.PetId] = pet

	return nil
}
