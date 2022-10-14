package perconaservermongodb

import (
	"context"

	"github.com/pkg/errors"

	api "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
	"github.com/percona/percona-server-mongodb-operator/pkg/psmdb"
	"github.com/percona/percona-server-mongodb-operator/pkg/psmdb/mongo"
)

func (r *ReconcilePerconaServerMongoDB) mongoClientWithRole(ctx context.Context, cr *api.PerconaServerMongoDB, rs api.ReplsetSpec, role UserRole) (mongo.Client, error) {
	c, err := r.getInternalCredentials(ctx, cr, role)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credentials")
	}

	return psmdb.MongoClient(ctx, r.client, cr, rs, c)
}

func (r *ReconcilePerconaServerMongoDB) mongosClientWithRole(ctx context.Context, cr *api.PerconaServerMongoDB, role UserRole) (mongo.Client, error) {
	c, err := r.getInternalCredentials(ctx, cr, role)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credentials")
	}

	return psmdb.MongosClient(ctx, r.client, cr, c)
}

func (r *ReconcilePerconaServerMongoDB) standaloneClientWithRole(ctx context.Context, cr *api.PerconaServerMongoDB, role UserRole, host string) (mongo.Client, error) {
	c, err := r.getInternalCredentials(ctx, cr, role)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get credentials")
	}

	return psmdb.StandaloneClient(ctx, r.client, cr, c, host)
}
