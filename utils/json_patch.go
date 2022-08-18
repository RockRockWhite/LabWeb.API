package utils

import (
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch/v5"
)

// ApplyJsonPatch 对entity应用JsonPatch
func ApplyJsonPatch(v interface{}, patchJson []byte) {
	// 反序列化v
	vJson, marshalErr := json.Marshal(v)
	if marshalErr != nil {
		panic(fmt.Sprintf("failed to marshel %v", v))
	}

	// 应用patch
	patch, decodePatchErr := jsonpatch.DecodePatch(patchJson)
	if decodePatchErr != nil {
		panic(fmt.Sprintf("PatchJson it not Correct"))
	}
	vJson, applyErr := patch.Apply(vJson)
	if applyErr != nil {
		panic(fmt.Sprintf("Failed to appliy patch %s to %s : %s", patchJson, vJson, applyErr))
	}

	// 反序列化为v
	if err := json.Unmarshal(vJson, v); err != nil {
		panic(fmt.Sprintf("Failed to Unmarshal json %s : %s", vJson, err))
	}
}
