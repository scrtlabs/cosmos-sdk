package scrt_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/scrt"
	"github.com/magiconair/properties/assert"
)

func TestFlattenEmpty(t *testing.T) {
	test_data := [][]byte{}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 0})
}

func TestUnFlattenEmpty(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 0}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{})
}

func TestFlattenOne_Empty(t *testing.T) {
	test_data := [][]byte{{}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0})
}

func TestUnFlattenOne_Empty(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{}})
}

func TestFlattenOne_0(t *testing.T) {
	test_data := [][]byte{{0}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0})
}

func TestUnFlattenOne_0(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{0}})
}

func TestFlattenOne_1(t *testing.T) {
	test_data := [][]byte{{1}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1})
}

func TestUnFlattenOne_1(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{1}})
}

func TestFlattenOne_2(t *testing.T) {
	test_data := [][]byte{{2}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 2})
}

func TestUnFlattenOne_2(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 2}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{2}})
}

func TestFlattenOne_0F(t *testing.T) {
	test_data := [][]byte{{0x0F}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0xF})
}

func TestUnFlattenOne_0F(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0xF}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{0x0F}})
}

func TestFlattenOne_FF(t *testing.T) {
	test_data := [][]byte{{0xFF}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0xFF})
}

func TestUnFlattenOne_FF(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0xFF}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{0xFF}})
}

func TestFlattenTwo_Empty(t *testing.T) {
	test_data := [][]byte{{}, {}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

func TestUnFlattenTwo_Empty(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{}, {}})
}

func TestFlattenTwo_1Empty(t *testing.T) {
	test_data := [][]byte{{1}, {}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0})
}

func TestUnFlattenTwo_1Empty(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{1}, {}})
}

func TestFlattenTwo_Empty1(t *testing.T) {
	test_data := [][]byte{{}, {1}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1})
}

func TestUnFlattenTwo_Empty1(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{}, {1}})
}

func TestFlattenTwo_1_1(t *testing.T) {
	test_data := [][]byte{{1}, {1}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1})
}

func TestUnFlattenTwo_1_1(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 1}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{1}, {1}})
}

func TestFlattenTwo_F_F(t *testing.T) {
	test_data := [][]byte{{0xF}, {0xF}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 1, 0xF, 0, 0, 0, 0, 0, 0, 0, 1, 0xF})
}

func TestUnFlattenTwo_F_F(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 1, 0xF, 0, 0, 0, 0, 0, 0, 0, 1, 0xF}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{0xF}, {0xF}})
}

func TestFlattenThree_F_F_F(t *testing.T) {
	test_data := [][]byte{{0xF}, {0xF}, {0xF}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 1, 0xF, 0, 0, 0, 0, 0, 0, 0, 1, 0xF, 0, 0, 0, 0, 0, 0, 0, 1, 0xF})
}

func TestUnFlattenTree_F_F_F(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 1, 0xF, 0, 0, 0, 0, 0, 0, 0, 1, 0xF, 0, 0, 0, 0, 0, 0, 0, 1, 0xF}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{0xF}, {0xF}, {0xF}})
}

func TestFlattenThree_FF_FFF_FFFF(t *testing.T) {
	test_data := [][]byte{{0xF, 0xF}, {0xF, 0xF, 0xF}, {0xF, 0xF, 0xF, 0xF}}
	flattened_data := scrt.Flatten(test_data)
	assert.Equal(t, flattened_data, []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0xF, 0xF, 0, 0, 0, 0, 0, 0, 0, 3, 0xF, 0xF, 0xF, 0, 0, 0, 0, 0, 0, 0, 4, 0xF, 0xF, 0xF, 0xF})
}

func TestUnFlattenTree_FF_FFF_FFFF(t *testing.T) {
	flat_data := []byte{6, 7, 25, 13, 103, 251, 114, 10, 0, 0, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0xF, 0xF, 0, 0, 0, 0, 0, 0, 0, 3, 0xF, 0xF, 0xF, 0, 0, 0, 0, 0, 0, 0, 4, 0xF, 0xF, 0xF, 0xF}
	x2_data := scrt.UnFlatten(flat_data)
	assert.Equal(t, x2_data, [][]byte{{0xF, 0xF}, {0xF, 0xF, 0xF}, {0xF, 0xF, 0xF, 0xF}})
}
