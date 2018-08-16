// Copyright (C) Extensible Service Proxy Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////
//
#include "src/api_manager/utils/marshalling.h"

#include "google/protobuf/io/zero_copy_stream_impl_lite.h"
#include "google/protobuf/util/json_util.h"
#include "google/protobuf/util/type_resolver.h"
#include "google/protobuf/util/type_resolver_util.h"

using ::google::protobuf::Message;
using ::google::protobuf::util::TypeResolver;
using ::google::protobuf::util::error::Code;

namespace pb = ::google::protobuf;

namespace google {
namespace api_manager {
namespace utils {

namespace {
const char kTypeUrlPrefix[] = "type.googleapis.com";

// Dual resolver accepts one or two resolvers. The first is the first
// resolver used to attempt to resolve a type url. If that fails, it
// will attempt to use the second resolver, if available.
class DualResolver : public pb::util::TypeResolver {
 public:
  DualResolver(pb::util::TypeResolver *a, pb::util::TypeResolver *b) :
    a(a), b(b) {}

  virtual pb::util::Status ResolveMessageType(const std::string& type_url,
                                            pb::Type* type) override {
    auto status = a->ResolveMessageType(type_url, type);
    if (b != nullptr && !status.ok()) {
      return b->ResolveMessageType(type_url, type);
    }
    return status;
  }

  virtual pb::util::Status ResolveEnumType(const std::string& type_url,
                                         pb::Enum* type) override {
    auto status = a->ResolveEnumType(type_url, type);
    if (b != nullptr && !status.ok()) {
      return b->ResolveEnumType(type_url, type);
    }
    return status;
  }
 private:
  pb::util::TypeResolver *a, *b;
};

// Creation function used by static lazy init.
TypeResolver* CreateTypeResolver() {
  return ::google::protobuf::util::NewTypeResolverForDescriptorPool(
      kTypeUrlPrefix, ::google::protobuf::DescriptorPool::generated_pool());
}

// Returns the singleton type resolver, creating it on first call.
TypeResolver* GetTypeResolver() {
  static TypeResolver* resolver = CreateTypeResolver();
  return resolver;
}
}  // namespace

std::string GetTypeUrl(const Message& message) {
  return std::string(kTypeUrlPrefix) + "/" +
         message.GetDescriptor()->full_name();
}

Status ProtoToJson(const Message& message, std::string* result, int options) {
  return ProtoToJson(message, result, options, *GetTypeResolver());
}

Status ProtoToJson(const Message& message,
                   std::string* result, int options,
                   ::google::protobuf::util::TypeResolver& resolver) {

  ::google::protobuf::util::JsonPrintOptions json_options;
  DualResolver dualResolver(&resolver, GetTypeResolver());
  if (options & JsonOptions::PRETTY_PRINT) {
    json_options.add_whitespace = true;
  }
  if (options & JsonOptions::OUTPUT_DEFAULTS) {
    json_options.always_print_primitive_fields = true;
  }
  // TODO: Skip going to bytes and use ProtoObjectSource directly.
  ::google::protobuf::util::Status status =
      ::google::protobuf::util::BinaryToJsonString(
          &dualResolver, GetTypeUrl(message), message.SerializeAsString(),
          result, json_options);
  return Status::FromProto(status);
}

Status ProtoToJson(const Message& message,
                   ::google::protobuf::io::ZeroCopyOutputStream* json,
                   int options) {
  ::google::protobuf::util::JsonPrintOptions json_options;
  if (options & JsonOptions::PRETTY_PRINT) {
    json_options.add_whitespace = true;
  }
  if (options & JsonOptions::OUTPUT_DEFAULTS) {
    json_options.always_print_primitive_fields = true;
  }
  // TODO: Skip going to bytes and use ProtoObjectSource directly.
  std::string binary = message.SerializeAsString();
  ::google::protobuf::io::ArrayInputStream binary_stream(binary.data(),
                                                         binary.size());
  ::google::protobuf::util::Status status =
      ::google::protobuf::util::BinaryToJsonStream(
          GetTypeResolver(), GetTypeUrl(message), &binary_stream, json,
          json_options);
  return Status::FromProto(status);
}

Status JsonToProto(const std::string& json, Message* message) {
  ::google::protobuf::util::JsonParseOptions options;
  options.ignore_unknown_fields = true;
  std::string binary;
  ::google::protobuf::util::Status status =
      ::google::protobuf::util::JsonToBinaryString(
          GetTypeResolver(), GetTypeUrl(*message), json, &binary, options);
  if (!status.ok()) {
    return Status::FromProto(status);
  }
  if (message->ParseFromString(binary)) {
    return Status::OK;
  }
  return Status(
      Code::INTERNAL,
      "Unable to parse bytes generated from JsonToBinaryString as proto.");
}

Status JsonToProto(::google::protobuf::io::ZeroCopyInputStream* json,
                   ::google::protobuf::Message* message) {
  ::google::protobuf::util::JsonParseOptions options;
  options.ignore_unknown_fields = true;
  std::string binary;
  ::google::protobuf::io::StringOutputStream output(&binary);
  ::google::protobuf::util::Status status =
      ::google::protobuf::util::JsonToBinaryStream(
          GetTypeResolver(), GetTypeUrl(*message), json, &output, options);

  if (!status.ok()) {
    return Status::FromProto(status);
  }
  if (message->ParseFromString(binary)) {
    return Status::OK;
  }
  return Status(
      Code::INTERNAL,
      "Unable to parse bytes generated from JsonToBinaryString as proto.");
}

}  // namespace utils
}  // namespace api_manager
}  // namespace google
