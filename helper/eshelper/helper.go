package eshelper

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"github.com/elastic/go-elasticsearch/v8"
)

type Helper struct {
	Client *elasticsearch.Client
}

type IndexHelper struct {
	*Helper
	IndexName string
}

func (h *Helper) CreateRandomIndex(settingsJson []byte) (*IndexHelper, error) {
	buf := new(bytes.Buffer)
	_, err := io.CopyN(buf, rand.Reader, 16)
	if err != nil {
		return nil, err
	}
	randStr := hex.EncodeToString(buf.Bytes())

	opt := h.Client.Indices.Create.WithBody(bytes.NewReader(settingsJson))
	res, err := h.Client.Indices.Create(randStr, opt)
	if err != nil {
		return nil, err
	} else if res.IsError() {
		return nil, fmt.Errorf("%s", res.String())
	}

	return &IndexHelper{
		Helper:    h,
		IndexName: randStr,
	}, nil
}

func (h *Helper) IsHealthy(ctx context.Context) (bool, error) {
	res, err := h.Client.Cluster.Health(h.Client.Cluster.Health.WithContext(ctx))
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	if res.IsError() {
		return false, fmt.Errorf("response error: %s", res.String())
	}

	// body must be like.
	// {
	//		"cluster_name":"docker-cluster",
	//		"status":"green",
	//		"timed_out":false,
	//		"number_of_nodes":1,
	//		"number_of_data_nodes":1,
	//		"active_primary_shards":1,
	//		"active_shards":1,
	//		"relocating_shards":0,
	//		"initializing_shards":0,
	//		"unassigned_shards":0,
	//		"delayed_unassigned_shards":0,
	//		"number_of_pending_tasks":0,
	//		"number_of_in_flight_fetch":0,
	//		"task_max_waiting_in_queue_millis":0,
	//		"active_shards_percent_as_number":100.0
	// }

	type clusterHealthPartial struct {
		Status string `json:"status"`
	}

	var health clusterHealthPartial

	err = json.NewDecoder(res.Body).Decode(&health)
	if err != nil {
		return false, err
	}

	return health.Status == "green", nil
}

func (h *IndexHelper) PostDoc(doc any) (docId string, err error) {
	bin, err := json.Marshal(doc)
	if err != nil {
		return "", err
	}

	res, err := h.Client.Index(h.IndexName, bytes.NewReader(bin))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.IsError() {
		return "", fmt.Errorf("response error: %s", res.String())
	}

	var result IndexResult
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Id_, nil
}

func (h *IndexHelper) GetDoc(id string, v any) error {
	res, err := h.Client.Get(h.IndexName, id)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("%s", res.String())
	}

	dec := json.NewDecoder(res.Body)
	dec.DisallowUnknownFields()
	var f FetchDocResult
	if err := dec.Decode(&f); err != nil {
		return err
	}

	err = json.Unmarshal(f.Source_, v)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes the index pointed by h.
func (h *IndexHelper) Delete() error {
	res, err := h.Client.Indices.Delete([]string{h.IndexName})
	if err != nil {
		return err
	} else if res.IsError() {
		return fmt.Errorf("%s", res.String())
	}
	return nil
}
