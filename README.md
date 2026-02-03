
Works with rules_go 0.57 but breaks in 0.58.2.
Culprit bisected to https://github.com/bazel-contrib/rules_go/pull/4438.

```
% bazel build //:world_go_test
INFO: Analyzed target //:world_go_test (129 packages loaded, 11094 targets configured).
ERROR: HOME/Development/experiments/reproduction_rules_go_0_59/BUILD:18:8: GoLink world_go_test_/world_go_test failed: (Exit 1): builder failed: error executing GoLink command (from target //:world_go_test) bazel-out/k8-opt-exec-ST-d57f47055a04/bin/external/rules_go++go_sdk+go_default_sdk/builder_reset/builder link -sdk external/rules_go++go_sdk+go_default_sdk -goroot ... (remaining 39 arguments skipped)

Use --sandbox_debug to see verbose messages from the sandbox and retain the sandbox build root for debugging
external/rules_go++go_sdk+go_default_sdk/pkg/tool/linux_amd64/link: running /tmp/3430629485/builder-cc failed: exit status 1
/tmp/3430629485/builder-cc -m64 -s -Wl,--build-id=0x498919aff001d8366249403a2544fba2d833084f -o bazel-out/k8-fastbuild/bin/world_go_test_/world_go_test -Wl,--export-dynamic-symbol=_cgo_panic -Wl,--export-dynamic-symbol=_cgo_topofstack -Wl,--export-dynamic-symbol=crosscall2 -Qunused-arguments -Wl,--compress-debug-sections=zlib /tmp/go-link-2449248489/go.o /tmp/go-link-2449248489/000000.o /tmp/go-link-2449248489/000001.o /tmp/go-link-2449248489/000002.o /tmp/go-link-2449248489/000003.o /tmp/go-link-2449248489/000004.o /tmp/go-link-2449248489/000005.o /tmp/go-link-2449248489/000006.o /tmp/go-link-2449248489/000007.o /tmp/go-link-2449248489/000008.o /tmp/go-link-2449248489/000009.o /tmp/go-link-2449248489/000010.o /tmp/go-link-2449248489/000011.o /tmp/go-link-2449248489/000012.o /tmp/go-link-2449248489/000013.o /tmp/go-link-2449248489/000014.o /tmp/go-link-2449248489/000015.o /tmp/go-link-2449248489/000016.o /tmp/go-link-2449248489/000017.o -Wl,-whole-archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo -Wl,-no-whole-archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libglobals.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libflag_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libmarshalling.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libreflection.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libconfig.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libprogram_name.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libprivate_handle_accessor.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libcommandlineflag.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libcommandlineflag_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/container/libraw_hash_set.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcord.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcordz_info.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcord_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcordz_functions.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcordz_handle.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/crc/libcrc_cord_state.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/crc/libcrc32c.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/crc/libcrc_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/crc/libcpu_detect.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libstr_format_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/container/libhashtablez_sampler.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/profiling/libexponential_biased.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/hash/libhash.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/hash/libcity.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/internal/libvlog_config.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/internal/libfnmatch.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/synchronization/libsynchronization.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/synchronization/libgraphcycles_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/synchronization/libkernel_timeout_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libtracing_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libstacktrace.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libsymbolize.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libdebugging_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libdemangle_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libdemangle_rust.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libdecode_rust_punycode.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libutf8_for_code_point.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libmalloc_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/time/libtime.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/time/internal/cctz/libtime_zone.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/time/internal/cctz/libcivil_time.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libstrings.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libinternal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libstring_view.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libthrow_delegate.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/numeric/libint128.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libbase.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libraw_logging_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/liblog_severity.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libspinlock_wait.a --target=x86_64-unknown-linux-gnu -no-canonical-prefixes -fuse-ld=lld -lm -Wl,--build-id=md5 -Wl,--hash-style=gnu -Wl,-z,relro,-z,now -l:libunwind.a -lpthread -ldl -rtlib=compiler-rt -l:libc++.a -l:libc++abi.a -pthread -pthread -pthread -pthread --target=x86_64-unknown-linux-gnu -no-canonical-prefixes -fuse-ld=lld -lm -Wl,--build-id=md5 -Wl,--hash-style=gnu -Wl,-z,relro,-z,now -l:libunwind.a -lpthread -ldl -rtlib=compiler-rt -l:libc++.a -l:libc++abi.a -pthread -lpthread -Wl,-whole-archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo -Wl,-no-whole-archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libglobals.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libflag_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libmarshalling.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libreflection.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libconfig.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libprogram_name.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libprivate_handle_accessor.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libcommandlineflag.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/flags/libcommandlineflag_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/container/libraw_hash_set.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcord.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcordz_info.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcord_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcordz_functions.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libcordz_handle.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/crc/libcrc_cord_state.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/crc/libcrc32c.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/crc/libcrc_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/crc/libcpu_detect.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libstr_format_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/container/libhashtablez_sampler.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/profiling/libexponential_biased.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/hash/libhash.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/hash/libcity.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/internal/libvlog_config.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/internal/libfnmatch.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/synchronization/libsynchronization.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/synchronization/libgraphcycles_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/synchronization/libkernel_timeout_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libtracing_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libstacktrace.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libsymbolize.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libdebugging_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libdemangle_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libdemangle_rust.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libdecode_rust_punycode.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/debugging/libutf8_for_code_point.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libmalloc_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/time/libtime.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/time/internal/cctz/libtime_zone.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/time/internal/cctz/libcivil_time.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libstrings.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libinternal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/strings/libstring_view.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libthrow_delegate.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/numeric/libint128.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libbase.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libraw_logging_internal.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/liblog_severity.a bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/base/libspinlock_wait.a --target=x86_64-unknown-linux-gnu -no-canonical-prefixes -fuse-ld=lld -lm -Wl,--build-id=md5 -Wl,--hash-style=gnu -Wl,-z,relro,-z,now -l:libunwind.a -lpthread -ldl -rtlib=compiler-rt -l:libc++.a -l:libc++abi.a -pthread -pthread -pthread -pthread -no-pie --target=x86_64-unknown-linux-gnu -no-canonical-prefixes -fuse-ld=lld -lm -Wl,--build-id=md5 -Wl,--hash-style=gnu -Wl,-z,relro,-z,now -l:libunwind.a -lpthread -ldl -rtlib=compiler-rt -l:libc++.a -l:libc++abi.a
ld.lld: error: duplicate symbol: FLAGS_stderrthreshold
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_stderrthreshold) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.data+0x0) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_minloglevel
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_minloglevel) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.data+0x60) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_log_backtrace_at
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_log_backtrace_at) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.data+0xc0) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_log_prefix
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_log_prefix) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.data+0x138) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_v
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_v) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.data+0x198) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_vmodule
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_vmodule) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.data+0x1f8) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_nostderrthreshold
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_nostderrthreshold) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.bss+0x1) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_nominloglevel
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_nominloglevel) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.bss+0x2) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_nolog_backtrace_at
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_nolog_backtrace_at) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.bss+0x3) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_nolog_prefix
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_nolog_prefix) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.bss+0x4) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_nov
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_nov) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.bss+0x5) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo

ld.lld: error: duplicate symbol: FLAGS_novmodule
>>> defined at flags.cc
>>>            flags.pic.o:(FLAGS_novmodule) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
>>> defined at flags.cc
>>>            flags.pic.o:(.bss+0x6) in archive bazel-out/k8-fastbuild/bin/external/abseil-cpp+/absl/log/libflags.lo
clang: error: linker command failed with exit code 1 (use -v to see invocation)

link: error running subcommand external/rules_go++go_sdk+go_default_sdk/pkg/tool/linux_amd64/link: exit status 2
Target //:world_go_test failed to build
Use --verbose_failures to see the command lines of failed build steps.
INFO: Elapsed time: 8.819s, Critical Path: 0.89s
INFO: 3 processes: 277 action cache hit, 3 internal.
ERROR: Build did NOT complete successfully
```
