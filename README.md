# bitrot_scanner

[Bitrot](https://arstechnica.com/information-technology/2014/01/bitrot-and-atomic-cows-inside-next-gen-filesystems/) is a huge issue if you care about your data.
This program aims to make the silent corruption a little less silent for those irreplaceable memories.

# Install

## Packages

 * [CentOS / RHEL 6 64 bit](https://packager.io/gh/kormoc/bitrot_scanner/install#centos-6-bitrot-scanner)
 * [ebian 7 Wheezy 64 bit](https://packager.io/gh/kormoc/bitrot_scanner/install#debian-7-bitrot-scanner)
 * [Ubuntu 14.04 Trusty 64bit](https://packager.io/gh/kormoc/bitrot_scanner/install#ubuntu-14-04-bitrot-scanner)
 * [Ubuntu 16.04 Xenial 64bit](https://packager.io/gh/kormoc/bitrot_scanner/install#ubuntu-16-04-bitrot-scanner)

## From source

`go get github.com/kormoc/bitrot_scanner`

# Usage

Generally the best way to handle this is to run `bitrot_scanner` via cron on a cadence that matches the importance of your data.

## Simplistic Mode

This mode will create checksums for any files missing them and validate any checksums that exist. Useful if you want to create checksums and validate on the same schedule.

`bitrot_scanner -progressBar /path/to/directory/1 /path/to/directory/2 ...`

## Advanced Mode

Splitting the creation of the checksums from the validation works best on large datasets that often have new files being created.
This allows a rapid creation of new checksums, but allows the validation to happen on a different schedule.

### Create checksums for files missing them

`bitrot_scanner -lockfile /var/run/bitrot_scanner.pid -skipValidation /path/to/directory/1 /path/to/directory/2 ...`

### Validate

`bitrot_scanner -lockfile /var/run/bitrot_scanner.pid -skipCreate /path/to/directory/1 /path/to/directory/2 ...`

# Hash Functions

By default, we use `sha512`, however you can use multiple hashes (`md5,sha512`). This allows you to migrate between hash types if you desire.

## Valid hashes

 * md5
 * sha1
 * sha256
 * sha512

# Developer information

## Regenerate vendor after changes

```
rm -rvf Godeps vendor
godep save ./...
```
