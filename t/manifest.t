#!perl
use strict;
use warnings;

use lib 't';
use helper;

my $tmp = workdir;
ok -d "t/repos/manifest-test", "manifest-test repo exists" or die;
chdir "t/repos/manifest-test" or die;

runs_ok "genesis manifest -c cloud.yml us-east-1-sandbox >$tmp/manifest.yml";
is get_file("$tmp/manifest.yml"), <<EOF, "manifest generated for us-east-1/sandbox";
jobs:
- name: thing
  properties:
    domain: sb.us-east-1.example.com
    endpoint: https://sb.us-east-1.example.com:8443
  templates:
  - name: bar
    release: foo
releases:
- name: foo
  version: 1.2.3-rc.1

EOF

runs_ok "genesis manifest -c cloud.yml us-west-1-sandbox >$tmp/manifest.yml";
is get_file("$tmp/manifest.yml"), <<EOF, "manifest generated for us-west-1/sandbox";
jobs:
- name: thing
  properties:
    domain: sandbox.us-west-1.example.com
    endpoint: https://sandbox.us-west-1.example.com:8443
  templates:
  - name: bar
    release: foo
releases:
- name: foo
  version: 1.2.3-rc.1

EOF


done_testing;
