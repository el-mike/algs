package sorting

import "testing"

func benchmarkMergeSort(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		MergeSort(data)
	}
}

func BenchmarkMergeSortSmall(b *testing.B) {
	benchmarkMergeSort(TestDataProviderFactory(SmallTestData), b)
}

func BenchmarkMergeSortBig(b *testing.B) {
	benchmarkMergeSort(TestDataProviderFactory(TestData), b)
}

// =============================================

func benchmarkMergeSort2(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		MergeSort2(data, 0, len(data)-1)
	}
}

// func BenchmarkMergeSort2Small(b *testing.B) {
// 	benchmarkMergeSort2(GetTestDataCopy(SmallTestData), b)
// }

// func BenchmarkMergeSort2Big(b *testing.B) {
// 	benchmarkMergeSort2(GetTestDataCopy, b)
// }

// =============================================

func benchmarkMergeSortFromNet(dataProvider func() []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := dataProvider()
		b.StartTimer()

		MergeSortFromNet(data)
	}
}

// func BenchmarkMergeSortFromNetSmall(b *testing.B) {
// 	benchmarkMergeSortFromNet(GetTestDataCopy(SmallTestData), b)
// }

// func BenchmarkMergeSortFromNetBig(b *testing.B) {
// 	benchmarkMergeSortFromNet(GetTestDataCopy, b)
// }
