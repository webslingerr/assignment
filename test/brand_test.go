package test

import (
	"app/api/models"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

var c int64

func TestActor(t *testing.T) {
	c = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := createBrand(t)
			deleteBrand(t, id)
		}()


	}

	wg.Wait()
	c++
	fmt.Println("c: ", c)
}

func createBrand(t *testing.T) int {
	response := &models.Brand{}

	request := &models.CreateBrand{
		BrandName: faker.FirstName(),
	}

	resp, err := PerformRequest(http.MethodPost, "/brand", &request, &response)

	assert.NoError(t, err)

	// a := object{} check whether the object is nil or not
	// 1 way
	assert.NotNil(t, resp)
	// another
	// b := object{}
	// reflect.DeepEqual(a, b)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.BrandId
}

func updateBrand(t *testing.T, id string) int {
	response := &models.Brand{}
	request := &models.UpdateBrand{
		BrandName: faker.FirstName(),
	}

	resp, err := PerformRequest(http.MethodPut, "/brand/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.BrandId
}

func deleteBrand(t *testing.T, id int) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/brand/%s", strconv.Itoa(id)),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
