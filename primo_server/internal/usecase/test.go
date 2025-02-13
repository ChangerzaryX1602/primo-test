package usecase

import "test/internal/repository"

type TestUsecase interface {
	Merge([]int64, []int64, []int64) ([]int64, error)
}
type testUsecase struct {
	repository repository.TestRepository
}

func NewTestUsecase(repository repository.TestRepository) TestUsecase {
	return &testUsecase{repository: repository}
}
func (u testUsecase) Merge(a []int64, b []int64, c []int64) ([]int64, error) {
	mergedAsc := MergeTwo(a, b)
	collection3Asc := Reverse(c)
	return MergeTwo(mergedAsc, collection3Asc), nil
}
func MergeTwo(a, b []int64) []int64 {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	if a[0] <= b[0] {
		return append([]int64{a[0]}, MergeTwo(a[1:], b)...)
	}
	return append([]int64{b[0]}, MergeTwo(a, b[1:])...)
}
func Reverse(arr []int64) []int64 {
	if len(arr) == 0 {
		return arr
	}
	return append([]int64{arr[len(arr)-1]}, Reverse(arr[:len(arr)-1])...)
}
