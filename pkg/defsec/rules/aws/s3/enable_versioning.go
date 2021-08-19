package s3

import (
	"fmt"

	"github.com/aquasecurity/tfsec/pkg/defsec/infra"
	"github.com/aquasecurity/tfsec/pkg/result"
)

func CheckVersioningIsEnabled(context *infra.Context) *result.Result {

	for _, bucket := range context.AWS.S3.Buckets {
		if !bucket.Versioning.Enabled.IsTrue() {
			return &result.Result{
				Description: fmt.Sprintf("Resource '%s' does not have versioning enabled", bucket.Reference),
				Location:    bucket.Versioning.Enabled.Range,
			}
		}
	}
	return nil
}