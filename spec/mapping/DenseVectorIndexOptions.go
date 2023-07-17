// This is implemented by the official document.
package mapping

type denseVectorSimilarity string

const (
	L2Norm     denseVectorSimilarity = "l2_norm"
	DotProduct denseVectorSimilarity = "dot_product"
	Cosine     denseVectorSimilarity = "cosine"
)

type DenseVectorIndexOptions struct {
	// Currently only hnsw is supported.
	Type string `json:"type"`
	// Defaults to 16.
	M int `json:"m"`
	// Defaults to 100.
	EfConstruction int `json:"ef_construction"`
}
