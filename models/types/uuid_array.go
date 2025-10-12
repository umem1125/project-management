package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UUIDArray []uuid.UUID

// function 
func (a *UUIDArray) Scan(value interface{}) error {
	var str string

	// logic 1: check the tipe data based on value
	switch v := value.(type) {
	case []byte :
		str = string(v)
	case string :
		str = v
		default :
		return errors.New("failed to parse UUIDArray: unsupported data type")
	}

	// buang kurawal depan - belakang
	str = strings.TrimPrefix(str,"{")
	str = strings.TrimPrefix(str,"}")
	parts := strings.Split(str,",")

	// make([]T, length, capacity)
	*a = make(UUIDArray, 0, len((parts)))
	for _ ,s := range parts {
		s = strings.TrimSpace(strings.Trim(s, `"`)) // menghapus spasi dan "
		if s == "" {
			continue
		}
		// parsing to UUID
		u, err := uuid.Parse(s)
		if err != nil {
			return fmt.Errorf("invalid UUID in Array : %v", err)
		}
		*a = append(*a, u);
	}
	return nil
}

// mengubah value UUIDArray to Postgre SQL
func (a UUIDArray) value()(driver.Value, error){
	// logic 1: pastikan receiver sudah punya nilai
	if len(a) == 0 {
		return "{}", nil
	}
	// menyusun agar UUID sesuai format Postgre SQL
	postgreFormat := make([]string,0,len(a))
	for _ , value := range a {
		postgreFormat = append(postgreFormat, fmt.Sprintf(`"%s`, value.String()))
	}

	return "{"+ strings.Join(postgreFormat, ",") +"}", nil
	
}

// gorm bisa mengenali
func (UUIDArray) GormDataType() string {
	return "uuid[]"
}