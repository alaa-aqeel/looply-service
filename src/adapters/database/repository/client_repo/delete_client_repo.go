package client_repo

import "context"

func (r *ClientRepo) Delete(ctx context.Context, id string) error {

	return r.base.Delete(ctx, "clients", id)
}
