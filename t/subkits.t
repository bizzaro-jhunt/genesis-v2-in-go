#!perl
use strict;
use warnings;

use lib 't';
use helper;

my $tmp = workdir;
ok -d "t/repos/subkit-test", "subkit-test repo exists" or die;
chdir "t/repos/subkit-test" or die;

output_ok "genesis lookup params.blobstore_type use-webdav nil", "webdav";
output_ok "genesis lookup params.blobstore_type use-s3 nil", "s3";
output_ok "genesis lookup params.blobstore_type use nil", "nil";

runs_ok "genesis manifest -c cloud.yml use-s3 >$tmp/manifest.yml";
is get_file("$tmp/manifest.yml"), <<EOF, "manifest generated with s3 subkit";
properties:
  blobstore:
    config:
      aki: yup, we got one
      secret: haha
    type: s3

EOF

runs_ok "genesis manifest -c cloud.yml use-webdav >$tmp/manifest.yml";
is get_file("$tmp/manifest.yml"), <<EOF, "manifest generated with webdav subkit";
properties:
  blobstore:
    config:
      url: https://blobstore.internal
    type: webdav

EOF

run_fails "genesis manifest -c cloud.yml use-the-wrong-thing >$tmp/errors", 67;
is get_file("$tmp/errors"), <<EOF, "manifest generate fails with an invalid blobstore_type param";
'magic' does not look like a valid blobstore_type.

I need to know what blobstore type you wish to use
for this deployment.  Valid values are 's3' for an
Amazon-backed blobstore solution, or 'webdav', for
a locally-hosted HTTP/DAV solution.
EOF

run_fails "genesis manifest -c cloud.yml use-nothing >$tmp/errors", 66;
is get_file("$tmp/errors"), <<EOF, "manifest generate fails without a valid blobstore_type param";
I could not find params.blobstore_type in use-nothing
(or in any of its predecessor files).

I need to know what blobstore type you wish to use
for this deployment.  Valid values are 's3' for an
Amazon-backed blobstore solution, or 'webdav', for
a locally-hosted HTTP/DAV solution.
EOF

done_testing;
