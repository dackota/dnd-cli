/*
Copyright © 2022 Dackota Johnson dackota.j@gmail.com

*/
package cmd

import (
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_fetchCharacterJson(t *testing.T) {
	resp, err := http.Get("https://character-service.dndbeyond.com/character/v3/character/69377870")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	assert.Equal(t, resp.StatusCode, 200)
}
