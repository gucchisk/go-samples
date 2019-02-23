{
  "targets": [
    {
      "target_name": "add",
      "cflags!": [ "-fno-exceptions" ],
      "cflags_cc!": [ "-fno-exceptions" ],
      "sources": [
	"add.cc"
      ],
      "libraries": [
	"-ladd",
	"-L<(module_root_dir)"
      ],
      "include_dirs": [
	"<!@(node -p \"require('node-addon-api').include\")"
      ],
      "defines": [ "NAPI_DISABLE_CPP_EXCEPTIONS" ]
    }
  ]
}
