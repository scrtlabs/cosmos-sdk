package scrt

import (
	"encoding/binary"
)

const MAGIC_SIG uint64 = 0x607190D67FB720A
const LEN_SECTOR uint64 = 8

func Flatten(in [][]byte) []byte {
	if in == nil {
		return nil
	}
	var flat_data []byte
	sig_chunk := make([]byte, LEN_SECTOR)
	binary.BigEndian.PutUint64(sig_chunk, MAGIC_SIG)
	flat_data = append(flat_data, sig_chunk...)
	data_len_chunk := make([]byte, 8)
	binary.BigEndian.PutUint64(data_len_chunk, uint64(len(in)))
	flat_data = append(flat_data, data_len_chunk...)
	for _, a := range in {
		a_len := make([]byte, 8)
		binary.BigEndian.PutUint64(a_len, uint64(len(a)))
		flat_data = append(flat_data, a_len...)
		flat_data = append(flat_data, a...)
	}
	return flat_data
}

func UnFlatten(in []byte) [][]byte {

	offset := uint64(LEN_SECTOR)
	sig := binary.BigEndian.Uint64(in[:offset])
	if sig != MAGIC_SIG {
		return nil
	}
	num_of_arrs := uint64(binary.BigEndian.Uint64(in[offset : offset+LEN_SECTOR]))
	offset += LEN_SECTOR
	var rec_data [][]byte = make([][]byte, num_of_arrs)
	for i := uint64(0); i < num_of_arrs; i++ {
		start := offset
		offset = start + LEN_SECTOR
		next_len := uint64(binary.BigEndian.Uint64(in[start:offset]))
		if next_len == 0 {
			rec_data[i] = []byte{}
			continue
		}
		start = offset
		offset = start + next_len
		next_data_chunk := in[start:offset]
		rec_data[i] = append(rec_data[i], next_data_chunk...)
	}
	return rec_data
}
