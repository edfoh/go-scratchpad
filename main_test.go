package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc                string
		data                [][]string
		topK                int
		wantTotalFileSize   int64
		wantTopKCollections map[string]int64
	}{
		{
			desc: "test 1",
			data: [][]string{
				{"file1.txt", "100"},
				{"file2.txt", "200", "collection1"},
				{"file3.txt", "200", "collection1"},
				{"file4.txt", "300", "collection2"},
				{"file5.txt", "100"},
			},
			topK:              2,
			wantTotalFileSize: 900,
			wantTopKCollections: map[string]int64{
				"collection1": 400,
				"collection2": 300,
			},
		},
		{
			desc: "test 1",
			data: [][]string{
				{"file1.txt", "100"},
				{"file2.txt", "200", "collection1"},
				{"file3.txt", "200", "collection3"},
				{"file4.txt", "300", "collection2"},
				{"file5.txt", "100", "collection4"},
				{"file6.txt", "1200", "collection1"},
				{"file7.txt", "600", "collection3"},
				{"file8.txt", "900", "collection5"},
				{"file9.txt", "1100", "collection5"},
				{"file10.txt", "2300", "collection6"},
				{"file11.txt", "5600", "collection6"},
				{"file12.txt", "900", "collection6"},
				{"file13.txt", "900", "collection7"},
			},
			topK:              5,
			wantTotalFileSize: 14400,
			// 1 - 1400, 2-300, 3-800, 4-100, 5-2000, 6-8800, 7-900
			wantTopKCollections: map[string]int64{
				"collection6": 8800,
				"collection5": 2000,
				"collection1": 1400,
				"collection7": 900,
				"collection3": 800,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			catalog := NewCatalog()
			catalog.Add(tC.data)

			assert.Equal(t, tC.wantTotalFileSize, catalog.TotalSizeOfFiles())
			assert.Equal(t, tC.wantTopKCollections, catalog.TopKCollections(tC.topK))
		})
	}
}
