//go:build (all || resource_workitem) && !resource_workitem
// +build all resource_workitem
// +build !resource_workitem

package workitemtracking

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestWorkItem_GetWorkItem(t *testing.T) {
	r := ResourceWorkItem()
	d := schema.TestResourceDataRaw(t, r.Schema, map[string]interface{}{})
	input := map[string]interface{}{
		"System.State":      "To Do",
		"System.Title":      "TestTitle",
		"instance_protocol": "http",
		"lb_protocol":       "http",
	}
	mapSystemFields(d, &input)

	require.Equal(t, "TestTitle", d.Get("title").(string))
	//if !reflect.DeepEqual(out, c.expected) {
	//   t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
	//}
}
