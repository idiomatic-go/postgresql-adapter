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

func SLOEntry(ctx context.Context, uri string) (entry *v1.SLOEntry, err error) {
	u := urn.Parse(uri)
	if u.Err != nil {
		return nil, u.Err
	}
	entry = new(v1.SLOEntry)
	switch u.Nid {
	case urn.QbeNid, "", pgxsql.Scheme:
		break
	case fse.Scheme:
		err = processFse(ctx, entry)
		break
	default:
		err = errors.New(fmt.Sprintf("invalid Nid : %v", u.Nid))
	}
	return entry, err
}

func SLOEntryList(ctx context.Context, uri string) (entry []v1.SLOEntry, err error) {
	u := urn.Parse(uri)
	if u.Err != nil {
		return nil, u.Err
	}
	switch u.Nid {
	case urn.QbeNid, "":

		break
	case fse.Scheme:
		err = processFse(ctx, entry)
		break
	default:
		err = errors.New(fmt.Sprintf("invalid Nid : %v", u.Nid))
	}
	return entry, err
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
