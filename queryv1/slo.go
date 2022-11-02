package queryv1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/urn"
	v1 "github.com/idiomatic-go/core-types/corev1"
	"github.com/idiomatic-go/postgresql-adapter/pgxsql"
)

func SLOEntry(ctx context.Context, urn *urn.QbeURN) (*v1.SLOEntry, error) {
	var err = validateUrn(urn)
	if err != nil {
		return nil, err
	}
	entry := new(v1.SLOEntry)
	switch urn.Scheme() {
	case pgxsql.Scheme, "":
		break
	case fse.Scheme:
		err = processFse(ctx, entry)
		break
	default:
		err = errors.New(fmt.Sprintf("invalid scheme : %v", urn.Scheme()))
	}
	return entry, err
}

func SLOEntryList(ctx context.Context, urn *urn.QbeURN) ([]v1.SLOEntry, error) {
	var err = validateUrn(urn)
	if err != nil {
		return nil, err
	}
	var entry []v1.SLOEntry
	switch urn.Scheme() {
	case pgxsql.Scheme, "":

		break
	case fse.Scheme:
		err = processFse(ctx, entry)
		break
	default:
		err = errors.New(fmt.Sprintf("invalid scheme : %v", urn.Scheme()))
	}
	return entry, err
}

func validateUrn(urn *urn.QbeURN) error {
	if urn == nil {
		return errors.New("invalid QbeURN is nil")
	}
	return urn.Err
}

func processFse(ctx context.Context, entry any) error {
	fs := fse.ContextContent(ctx)
	if fs == nil {
		return errors.New(fmt.Sprintf("no content available for scheme : %v", fse.Scheme))
	} else {
		if fs.Content != nil {
			return json.Unmarshal(fs.Content, entry)
		}
	}
	return nil
}
