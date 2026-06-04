package cli

import "testing"

func TestLoadCatalogAndResolveDependencies(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	packs, err := catalog.Resolve([]string{"sdd"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	got := make([]string, 0, len(packs))
	for _, pack := range packs {
		got = append(got, pack.Name)
	}
	want := []string{"core", "sdd"}
	if len(got) != len(want) {
		t.Fatalf("resolved packs = %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("resolved packs = %v, want %v", got, want)
		}
	}
}

func TestCatalogSourcesExist(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	if err := catalog.ValidateSources(); err != nil {
		t.Fatalf("ValidateSources returned error: %v", err)
	}
}

func TestResolveBragMePack(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	packs, err := catalog.Resolve([]string{"brag-me"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	if len(packs) != 1 {
		t.Fatalf("resolved packs = %d, want 1", len(packs))
	}
	if packs[0].Name != "brag-me" {
		t.Fatalf("resolved pack = %q, want brag-me", packs[0].Name)
	}
}

func TestResolveRejectsUnknownPack(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	if _, err := catalog.Resolve([]string{"missing-pack"}); err == nil {
		t.Fatal("Resolve succeeded, want error")
	}
}
