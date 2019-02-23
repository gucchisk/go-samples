#include <napi.h>
#include "libadd.h"

Napi::Number Add(const Napi::CallbackInfo& info) {
  Napi::Env env = info.Env();
  int32_t a = info[0].As<Napi::Number>().Int32Value();
  int32_t b = info[1].As<Napi::Number>().Int32Value();
  return Napi::Number::New(env, add(a, b));
}
Napi::Object Init(Napi::Env env, Napi::Object exports) {
  exports.Set(Napi::String::New(env, "add"), Napi::Function::New(env, Add));
  return exports;
}
    
NODE_API_MODULE(add, Init);
