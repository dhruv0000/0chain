package datastore

import (
  "context"
  "io"
 "fmt"
 "bytes"
 "encoding/json"
 "0chain.net/common"
)

var (
	/*EntityNotFound code should be used to check whether an entity is found or not */
	EntityNotFound = "entity_not_found"
	/*EntityDuplicate codee should be used to check if an entity is already present */
	EntityDuplicate = "duplicate_entity"
)


/*Key - a type for the entity key */
type Key = string

/*ToString - return string representation of the key */
func ToString(key Key) string {
	return string(key)
}

func IsEmpty(key Key) bool {
	return len(key) == 0
}

func IsEqual(key1 Key, key2 Key) bool {
   return key1 == key2
}

/*EmptyKey - Represents an empty key */
var EmptyKey = Key("")

/*ToKey - takes an interface and returns a Key */
func ToKey(key interface{}) Key {
	switch v := key.(type) {
	case string:
		return Key(v)
	case []byte:
		return Key(v)
	default:
		return Key(fmt.Sprintf("%v", v))
	}
}

/*Entity - interface that reads and writes any implementing structure as JSON into the store */
type Entity interface {
	GetEntityName() string
	SetKey(key Key)
	GetKey() Key
	ComputeProperties()
	Validate(ctx context.Context) error
}

/*IDField - Useful to embed this into all the entities and get consistent behavior */
type IDField struct {
	ID Key `json:"id"`
}

/*SetKey sets the key */
func (k *IDField) SetKey(key Key) {
	k.ID = key
}

/*GetKey returns the key for the entity */
func (k *IDField) GetKey() Key {
	return k.ID
}

/*Validate - just an abstract implementation */
func (k *IDField) Validate(ctx context.Context) error {
	return nil
}

/*ComputeProperties - default dummy implementation so only entities that need this can implement */
func (k *IDField) ComputeProperties() {

}

/*Read - abstract method for memory store read */
func (k *IDField) Read(ctx context.Context, key string) error {
	return common.NewError("abstract_read", "Calling entity.Read() requires implementing the method")
}

/*Write - abstract method for memory store write */
func (k *IDField) Write(ctx context.Context) error {
	return common.NewError("abstract_write", "Calling entity.Write() requires implementing the method")
}

/*Delete - abstract method for memory store delete */
func (k *IDField) Delete(ctx context.Context) error {
	return common.NewError("abstract_delete", "Calling entity.Delete() requires implementing the method")
}

/*CreationTrackable - an interface that supports tracking the creation time */
type CreationTrackable interface {
	GetCreationTime() common.Timestamp
}

/*CreationDateField - Can be used to add a creation date functionality to an entity */
type CreationDateField struct {
	CreationDate common.Timestamp `json:"creation_date"`
}

/*InitializeCreationDate sets the creation date to current time */
func (cd *CreationDateField) InitializeCreationDate() {
	cd.CreationDate = common.Now()
}

/*GetCreationTime - Get the creation time */
func (cd *CreationDateField) GetCreationTime() common.Timestamp {
	return cd.CreationDate
}

/*ToJSON - given an entity, get the json of that entity as a buffer */
func ToJSON(entity Entity) *bytes.Buffer {
  buffer := bytes.NewBuffer(make([]byte, 0, 256))
	json.NewEncoder(buffer).Encode(entity)
  return buffer
}

/*FromJSON - read data into an entity */
func FromJSON(data interface{}, entity Entity) error {
  switch jsondata := data.(type) {
  case []byte:
    return json.Unmarshal(jsondata, entity)
  case string:
    return json.Unmarshal([]byte(jsondata), entity)
  case io.Reader:
    return json.NewDecoder(jsondata).Decode(entity)
  default:
    return common.NewError("unknown_data_type",fmt.Sprintf("unknown data type for reading entity from json: %T, %v\n",data,data))
  }
}
