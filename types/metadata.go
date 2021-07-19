package types

type MetadataModules struct {
	Name        string                 `json:"name"`
	Prefix      string                 `json:"prefix"`
	Storage     []MetadataStorage      `json:"storage"`
	Calls       []MetadataCalls        `json:"calls,omitempty"`
	CallsValue  map[string]interface{} `json:"calls_value,omitempty"`
	Events      []MetadataEvents       `json:"events,omitempty"`
	EventsValue map[string]interface{} `json:"events_value,omitempty"`
	Constants   []MetadataConstants    `json:"constants,omitempty"`
	Errors      []MetadataModuleError  `json:"errors"`
	Index       int                    `json:"index"`
}

type MetadataStorage struct {
	Name     string      `json:"name"`
	Modifier string      `json:"modifier"`
	Type     StorageType `json:"type"`
	Fallback string      `json:"fallback"`
	Docs     []string    `json:"docs"`
	Hasher   string      `json:"hasher,omitempty"`
}

type StorageType struct {
	Origin             string         `json:"origin"`
	PlainType          *string        `json:"plain_type,omitempty"`
	PlainTypeValue     *int           `json:"PlainTypeValue,omitempty"`
	MapType            *MapType       `json:"map_type,omitempty"`
	MapTypeValue       *MapTypeValue  `json:"map_type_value,omitempty"`
	DoubleMapType      *MapType       `json:"double_map_type,omitempty"`
	DoubleMapTypeValue *MapTypeValue  `json:"double_map_type_value,omitempty"`
	NMapType           *NMapType      `json:"n_map_type,omitempty"`
	NMapTypeValue      *NMapTypeValue `json:"n_map_type_value,omitempty"`
}

type MapType struct {
	Key        string `json:"key"`
	Key2       string `json:"key2,omitempty"`
	Hasher     string `json:"hasher"`
	Key2Hasher string `json:"key2Hasher,omitempty"`
	Value      string `json:"value"`
	IsLinked   bool   `json:"isLinked"`
}

type MapTypeValue struct {
	Key        int    `json:"key"`
	Key2       int    `json:"key2,omitempty"`
	Hasher     string `json:"hasher"`
	Key2Hasher string `json:"key2Hasher,omitempty"`
	Value      int    `json:"value"`
}

type NMapType struct {
	Hashers []string `json:"hashers"`
	KeyVec  []string `json:"key_vec"`
	Value   string   `json:"value"`
}

type NMapTypeValue struct {
	Hashers []string `json:"hashers"`
	KeyVec  []int    `json:"key_vec"`
	Value   int      `json:"value"`
}

type MetadataCalls struct {
	Lookup string                       `json:"lookup"`
	Name   string                       `json:"name"`
	Docs   []string                     `json:"docs"`
	Args   []MetadataModuleCallArgument `json:"args"`
}

type MetadataEvents struct {
	Lookup string   `json:"lookup"`
	Name   string   `json:"name"`
	Docs   []string `json:"docs"`
	Args   []string `json:"args"`
}

type MetadataStruct struct {
	MetadataVersion int                   `json:"metadata_version"`
	Metadata        MetadataTag           `json:"metadata"`
	Lookup          interface{}           `json:"lookup,omitempty"`
	CallIndex       map[string]CallIndex  `json:"call_index"`
	EventIndex      map[string]EventIndex `json:"event_index"`
	Extrinsic       *ExtrinsicMetadata    `json:"extrinsic"`
}

type CallIndex struct {
	Module MetadataModules `json:"module"`
	Call   MetadataCalls   `json:"call"`
}

type EventIndex struct {
	Module MetadataModules `json:"module"`
	Call   MetadataEvents  `json:"call"`
}

type MetadataTag struct {
	Modules []MetadataModules `json:"modules"`
}

type MetadataConstants struct {
	Name           string   `json:"name"`
	Type           string   `json:"type"`
	ConstantsValue string   `json:"constants_value"`
	Docs           []string `json:"docs"`
}

type ExtrinsicParam struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}
