package hdkey

import (
	"testing"
)

func BenchmarkDerivePublicFromSecret(b *testing.B) {
	seed, _ := GenerateSeed(RecommendedSeedSize)
	master, _ := NewMasterHDKey(seed, testMasterKey, testSecretVersion)

	b.ResetTimer()
	for i := uint32(0); i < uint32(b.N); i++ {
		_, _ = master.Child(i)
	}
}

func BenchmarkDeriveSecretFromSecret(b *testing.B) {
	seed, _ := GenerateSeed(RecommendedSeedSize)
	master, _ := NewMasterHDKey(seed, testMasterKey, testSecretVersion)

	b.ResetTimer()
	for i := uint32(0); i < uint32(b.N); i++ {
		_, _ = master.Child(HardenedKeyStart + i)
	}
}

func BenchmarkDerivePublicFromPublic(b *testing.B) {
	seed, _ := GenerateSeed(RecommendedSeedSize)
	master, _ := NewMasterHDKey(seed, testMasterKey, testSecretVersion)
	masterPub, _ := master.Neuter(testVMap)

	b.ResetTimer()
	for i := uint32(0); i < uint32(b.N); i++ {
		_, _ = masterPub.Child(i)
	}
}
