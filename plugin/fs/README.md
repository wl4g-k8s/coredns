# fs

## Name

*fs* - provide a file system abstraction.

## Description

Take into account *root*? Only for zone data??

## Syntax

~~~ txt
fs TYPE MOUNTPOINT
~~~

**TYPE** is the "file system" type. Defined are `disk` and `mem`. **MOUNTPOINT** is the location
where is mounted. By default the plugin mount "/" to point to `disk`. But for instance if you can
mount `s3` (if implemented) under `/s3`; then a path used in any other plugin that starts with
`/s3` will use S3 storage underneath. This plugin may be used multiple times to register multiple
paths using (different) "file system" implementations.

Or use a longer syntax if the file system needs more options.

~~~ corefile {
. {
    fs disk / {
      //  option_1
    }
}

## Examples

The following is the compiled in default:

~~~ corefile
. {
    fs disk /
}
~~~

Here we have 2 mountpoints:

~~~ txt
. {
    fs disk /
    fs s3 /etc
}
~~~

All accesses to /etc will use a s3 backed storage and everything else will go to disk (/).

## Notes

The *tls* plugin (as all default plugins) uses the *fs* plugin for storage abstraction, except when
finding the root CAs bundled by the OS. This is a std pkg library call which we can not intercept.
