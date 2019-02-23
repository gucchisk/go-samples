#include <napi.h>
#include "libsub.h"

Napi::Number Sub(const Napi::CallbackInfo& info) {
  Napi::Env env = info.Env();
  int32_t a = info[0].As<Napi::Number>().Int32Value();
  int32_t b = info[1].As<Napi::Number>().Int32Value();
  return Napi::Number::New(env, sub(a, b));
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
  exports.Set(Napi::String::New(env, "sub"), Napi::Function::New(env, Sub));
  return exports;
}

NODE_API_MODULE(sub, Init);
