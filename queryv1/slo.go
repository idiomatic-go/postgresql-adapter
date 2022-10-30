package queryv1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/util"
	v1 "github.com/idiomatic-go/core-types/corev1"
	"github.com/idiomatic-go/postgresql-adapter/pgxsql"
)

func SLOEntry(ctx context.Context, urn string) (*v1.SLOEntry, error) {
	var err error
	entry := new(v1.SLOEntry)
	nsid, nss := util.UrnParse(urn)
	switch nsid {
	case pgxsql.Scheme:
		if nss != "" {

		}
		break
	case fse.Scheme:
		fs := fse.ContextContent(ctx)
		if fs == nil || fs.Content == nil {
			errors.New(fmt.Sprintf("no content available for urn scheme : %v", nsid))
		}
		err = json.Unmarshal(fs.Content, entry)
		break
	default:
		return nil, errors.New(fmt.Sprintf("invalid urn scheme : %v", nsid))
	}
	return entry, err
}

func SLOEntryList(ctx context.Context, urn string) ([]v1.SLOEntry, error) {
	var err error
	var entry []v1.SLOEntry
	nsid, nss := util.UrnParse(urn)
	switch nsid {
	case pgxsql.Scheme:
		if nss != "" {

		}
		break
	case fse.Scheme:
		fs := fse.ContextContent(ctx)
		if fs == nil || fs.Content == nil {
			errors.New(fmt.Sprintf("no content available for urn scheme : %v", nsid))
		}
		err = json.Unmarshal(fs.Content, entry)
		break
	default:
		return nil, errors.New(fmt.Sprintf("invalid urn scheme : %v", nsid))
	}
	return entry, err
}
