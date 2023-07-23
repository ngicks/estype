package test

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/ngicks/estype/helper/eshelper"
	"github.com/ngicks/estype/spec/indices/indexstate"
	"github.com/ngicks/estype/spec/mapping"
	serde "github.com/ngicks/und/serde"
	"github.com/stretchr/testify/require"
)

var client *elasticsearch.Client

func init() {
	var err error
	client, err = elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
}

var (
	//go:embed testdata/all.json
	allMapping []byte
	//go:embed testdata/conversion.json
	conversionMapping []byte
	//go:embed testdata/dynamic.json
	dynamicMapping []byte
)

type settings struct {
	Settings map[string]any      `json:"settings"`
	Mappings mapping.TypeMapping `json:"mappings"`
}

type elasticsearchAcceptanceTestCase struct {
	mappings    []byte
	sampleInput any
	emptyValue  any // special value for testing stroing empty value.
	toRaw       func(v any) any
}

func toRaw[T interface{ ToRaw() U }, U any](v any) any {
	return v.(T).ToRaw()
}

// This test checks that generated types are accepted with an index created with source mapping.
// Both plain and raw types are tested.
func TestElasticsearchAcceptance(t *testing.T) {
	require := require.New(t)

	h := eshelper.Helper{Client: client}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if healthy, err := h.IsHealthy(ctx); err != nil || !healthy {
		t.Skipf("test is skipped because an elasticsearch is not accessible nor healthy." +
			" You must set the ELASTICSEARCH_URL environment variable to get this test working.")
		return
	}

	for _, tc := range []elasticsearchAcceptanceTestCase{
		{
			mappings:    allMapping,
			sampleInput: sampleAll,
			emptyValue:  sampleEmptyAllRaw,
			toRaw:       toRaw[All, AllRaw],
		},
		{
			mappings:    allMapping,
			sampleInput: sampleAllOptional,
			toRaw:       toRaw[AllOptional, AllOptionalRaw],
		},
		{
			mappings:    allMapping,
			sampleInput: sampleAllOptionalZero,
			toRaw:       toRaw[AllOptional, AllOptionalRaw],
		},
		{
			mappings:    conversionMapping,
			sampleInput: sampleConversion,
			toRaw:       toRaw[Conversion, ConversionRaw],
		}, {
			mappings:    dynamicMapping,
			sampleInput: sampleDynamic,
			toRaw:       toRaw[Dynamic, DynamicRaw],
		},
	} {
		var mapping map[string]indexstate.IndexState
		err := json.Unmarshal(tc.mappings, &mapping)
		require.NoError(err)

		indexSettings := settings{
			Settings: map[string]any{
				"number_of_replicas": 0, // This prevents es from being yellow after creation of index. Only needed if es is single-node.
			},
			Mappings: getMapping(mapping),
		}

		settingsJson, err := serde.Marshal(indexSettings)
		require.NoError(err)

		indexHelper, err := h.CreateRandomIndex(settingsJson)
		require.NoError(err)
		defer func() {
			_ = indexHelper.Delete()
		}()

		printJson(t, tc.sampleInput)
		docId, err := indexHelper.PostDoc(tc.sampleInput)
		require.NoError(err)
		t.Logf("docId = %s\n", docId)

		raw := tc.toRaw(tc.sampleInput)
		printJson(t, raw)
		docId, err = indexHelper.PostDoc(raw)
		require.NoError(err)
		t.Logf("docId = %s\n", docId)

		if tc.emptyValue != nil {
			printJson(t, tc.emptyValue)
			docId, err = indexHelper.PostDoc(tc.emptyValue)
			require.NoError(err)
			t.Logf("docId = %s\n", docId)
		}
	}
}

func getMapping(m map[string]indexstate.IndexState) mapping.TypeMapping {
	for _, v := range m {
		// return first. The input is assumed to only have an entry.
		return v.Mappings.Value()
	}
	panic(fmt.Errorf("unknown input = %+#v", m))
}

func printJson(t *testing.T, v any) {
	bin, err := serde.Marshal(v)
	if err != nil {
		panic(err)
	}
	t.Logf("%s\n", bin)
}
