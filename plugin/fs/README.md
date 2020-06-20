# fs

## Name

*fs* - provide a file system abstraction.

## Description

This plugin provides a file system abstraction for other plugins to use. This makes it possible that
data is hosted on local disk, memory or cloud storage like S3.

## Syntax

~~~ txt
fs TYPE MOUNTPOINT
~~~

**TYPE** is the "file system" type. Currently only `disk` is defined. **MOUNTPOINT** is the location
where is mounted. By default the plugin mounts "/" to point to `disk`. But for instance if you can
mount `s3` (if implemented) under `/s3`; then a path used by another plugin that starts with `/s3`
will use S3 storage underneath.

This plugin may be used multiple times to register multiple paths using (different) "file system"
implementations.

Supported **TYPE**s:

* `disk`: use local (disk) storage as provided by the OS

If any option are needed for the file system implementation they may be given with this extended
syntax:

~~~ txt
fs TYPE MOUNTPOINT {
    disk_ro
}
~~~

Where **TYPE** and **MOUNTPOINT** are as described about, each `<TYPE>_option` indicates an option
for that **TYPE** of file system.

## Examples

The following is the default configuration:

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

## Use in Plugins

Using this plugin in other plugins is done by querying the registry and using the returned `FileSystem`
implementation. `Lookup` will fallback on `disk` if no other file systems have been registered.

~~~ go
path := /my/path/to/a/file"

disk := fs.Registry.Lookup(path)
buf, err := disk.ReadFile(path)
if err != nil {
    return err
}
~~~

The returned `FileSystem` will not be changed at run-time and may be cached by the plugin.

### Notes on Specific Plugins

The *tls* plugin (as all default plugins) uses the *fs* plugin for storage abstraction, except when
finding the root CAs bundled by the OS. This is a std pkg library call which we can not intercept.
