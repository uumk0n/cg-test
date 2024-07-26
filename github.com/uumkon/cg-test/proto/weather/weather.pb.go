// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/weather.proto

package weather

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CurrentWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiToken string `protobuf:"bytes,1,opt,name=api_token,json=apiToken,proto3" json:"api_token,omitempty"`
	City     string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *CurrentWeatherRequest) Reset() {
	*x = CurrentWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentWeatherRequest) ProtoMessage() {}

func (x *CurrentWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentWeatherRequest.ProtoReflect.Descriptor instead.
func (*CurrentWeatherRequest) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{0}
}

func (x *CurrentWeatherRequest) GetApiToken() string {
	if x != nil {
		return x.ApiToken
	}
	return ""
}

func (x *CurrentWeatherRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CurrentWeatherRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type CurrentWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	TempC      float32 `protobuf:"fixed32,2,opt,name=temp_c,json=tempC,proto3" json:"temp_c,omitempty"`
	Data       string  `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CurrentWeatherResponse) Reset() {
	*x = CurrentWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentWeatherResponse) ProtoMessage() {}

func (x *CurrentWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentWeatherResponse.ProtoReflect.Descriptor instead.
func (*CurrentWeatherResponse) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{1}
}

func (x *CurrentWeatherResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CurrentWeatherResponse) GetTempC() float32 {
	if x != nil {
		return x.TempC
	}
	return 0
}

func (x *CurrentWeatherResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_proto_weather_proto protoreflect.FileDescriptor

var file_proto_weather_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x22, 0x64,
	0x0a, 0x15, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x16, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x15, 0x0a, 0x06, 0x74, 0x65, 0x6d, 0x70, 0x5f, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x74, 0x65, 0x6d, 0x70, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x66, 0x0a, 0x0e, 0x57, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x75, 0x75, 0x6d, 0x6b, 0x6f, 0x6e, 0x2f, 0x63, 0x67, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_weather_proto_rawDescOnce sync.Once
	file_proto_weather_proto_rawDescData = file_proto_weather_proto_rawDesc
)

func file_proto_weather_proto_rawDescGZIP() []byte {
	file_proto_weather_proto_rawDescOnce.Do(func() {
		file_proto_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_weather_proto_rawDescData)
	})
	return file_proto_weather_proto_rawDescData
}

var file_proto_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_weather_proto_goTypes = []any{
	(*CurrentWeatherRequest)(nil),  // 0: weather.CurrentWeatherRequest
	(*CurrentWeatherResponse)(nil), // 1: weather.CurrentWeatherResponse
}
var file_proto_weather_proto_depIdxs = []int32{
	0, // 0: weather.WeatherService.GetCurrentWeather:input_type -> weather.CurrentWeatherRequest
	1, // 1: weather.WeatherService.GetCurrentWeather:output_type -> weather.CurrentWeatherResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_weather_proto_init() }
func file_proto_weather_proto_init() {
	if File_proto_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_weather_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CurrentWeatherRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_weather_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CurrentWeatherResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_weather_proto_goTypes,
		DependencyIndexes: file_proto_weather_proto_depIdxs,
		MessageInfos:      file_proto_weather_proto_msgTypes,
	}.Build()
	File_proto_weather_proto = out.File
	file_proto_weather_proto_rawDesc = nil
	file_proto_weather_proto_goTypes = nil
	file_proto_weather_proto_depIdxs = nil
}
