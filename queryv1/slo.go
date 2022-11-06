package queryv1

import (
	"context"
	"errors"
	"fmt"
	"github.com/idiomatic-go/common-lib/fse"
	"github.com/idiomatic-go/common-lib/util"
	v1 "github.com/idiomatic-go/core-types/corev1"
)

func SLOEntry(ctx context.Context, urn util.URN) (entry []v1.SLOEntry, err error) {
	switch urn.Nid {
	case util.QbeNid, "":

		break
	case fse.Scheme:
		return fse.ProcessContent[[]v1.SLOEntry](ctx)
		break
	default:
		err = errors.New(fmt.Sprintf("invalid request: Nid not supported %v", urn.Nid))
	}
	return entry, err
}
