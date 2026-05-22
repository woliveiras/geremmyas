package cli

import "testing"

func TestLoadCatalogAndResolveDependencies(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}

	packs, err := catalog.Resolve([]string{"afk"})
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	got := make([]string, 0, len(packs))
	for _, pack := range packs {
		got = append(got, pack.Name)
	}
	want := []string{"core", "sdd", "afk"}
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

func TestResolveRejectsUnknownPack(t *testing.T) {
	catalog, err := loadCatalog()
	if err != nil {
		t.Fatalf("loadCatalog returned error: %v", err)
	}
	if _, err := catalog.Resolve([]string{"missing-pack"}); err == nil {
		t.Fatal("Resolve succeeded, want error")
	}
}
