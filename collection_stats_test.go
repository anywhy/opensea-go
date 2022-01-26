package opensea

import "testing"

func TestClient_CollectionStats(t *testing.T) {
	var cli = newClient()
	var collection, err = cli.CollectionStats(ctx, &CollectionRequest{
		CollectionSlug: "doodles-official",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(collection)
}
