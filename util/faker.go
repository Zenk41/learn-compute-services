package util

import "github.com/go-faker/faker/v4"

func CreateFaker[T any]() (T, error) {
	var fakeData *T = new(T)
	if err := faker.FakeData(fakeData); err != nil {
		return *fakeData, err
	}
	return *fakeData, nil

}
