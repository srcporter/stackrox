package clair

import (
	"testing"

	clairV1 "github.com/coreos/clair/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/clair/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertVulnerability(t *testing.T) {
	clairVulns, protoVulns := mock.GetTestVulns()
	for i, vuln := range clairVulns {
		assert.Equal(t, protoVulns[i], ConvertVulnerability(vuln))
	}
}

func TestConvertFeatures(t *testing.T) {
	clairFeatures, protoComponents := mock.GetTestFeatures()
	assert.Equal(t, protoComponents, ConvertFeatures(nil, clairFeatures))
}

func componentWithLayerIndex(name string, idx int32) *storage.ImageScanComponent {
	c := &storage.ImageScanComponent{
		Name: name,

		Vulns: []*storage.Vulnerability{},
	}
	if idx != -1 {
		c.HasLayerIndex = &storage.ImageScanComponent_LayerIndex{
			LayerIndex: idx,
		}
	}
	return c
}

func TestConvertFeaturesWithLayerIndexes(t *testing.T) {
	var cases = []struct {
		name               string
		metadata           *storage.ImageMetadata
		features           []clairV1.Feature
		expectedComponents []*storage.ImageScanComponent
	}{
		{
			name: "Nil metadata",
		},
		{
			name:     "Empty metadata",
			metadata: &storage.ImageMetadata{},
		},
		{
			name: "v1 metadata with equal vulns and layers - no empty",
			metadata: &storage.ImageMetadata{
				V1: &storage.V1Metadata{
					Layers: []*storage.ImageLayer{{}, {}},
				},
				LayerShas: []string{"A", "B"},
			},
			features: []clairV1.Feature{
				{
					Name:    "a-name",
					AddedBy: "A",
				},
				{
					Name:    "b-name",
					AddedBy: "B",
				},
			},
			expectedComponents: []*storage.ImageScanComponent{
				componentWithLayerIndex("a-name", 0),
				componentWithLayerIndex("b-name", 1),
			},
		},
		{
			name: "v1 metadata with fewer vulns than layers - no empty",
			metadata: &storage.ImageMetadata{
				V1: &storage.V1Metadata{
					Layers: []*storage.ImageLayer{{}, {}},
				},
				LayerShas: []string{"A", "B"},
			},
			features: []clairV1.Feature{
				{
					Name:    "b-name",
					AddedBy: "B",
				},
			},
			expectedComponents: []*storage.ImageScanComponent{
				componentWithLayerIndex("b-name", 1),
			},
		},
		{
			name: "v2 metadata with fewer vulns than layers - no empty",
			metadata: &storage.ImageMetadata{
				V1: &storage.V1Metadata{
					Layers: []*storage.ImageLayer{{}, {}},
				},
				V2:        &storage.V2Metadata{},
				LayerShas: []string{"A", "B"},
			},
			features: []clairV1.Feature{
				{
					Name:    "b-name",
					AddedBy: "B",
				},
			},
			expectedComponents: []*storage.ImageScanComponent{
				componentWithLayerIndex("b-name", 1),
			},
		},
		{
			name: "v2 metadata with empty layers",
			metadata: &storage.ImageMetadata{
				V1: &storage.V1Metadata{
					Layers: []*storage.ImageLayer{{Empty: true}, {}, {}},
				},
				V2:        &storage.V2Metadata{},
				LayerShas: []string{"A", "B"},
			},
			features: []clairV1.Feature{
				{
					Name:    "b-name",
					AddedBy: "B",
				},
			},
			expectedComponents: []*storage.ImageScanComponent{
				componentWithLayerIndex("b-name", 2),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			img := &storage.Image{
				Metadata: c.metadata,
			}
			convertedComponents := ConvertFeatures(img, c.features)
			require.Equal(t, len(c.expectedComponents), len(convertedComponents))
			for i := range convertedComponents {
				assert.Equal(t, c.expectedComponents[i], convertedComponents[i])
			}
		})
	}
}
