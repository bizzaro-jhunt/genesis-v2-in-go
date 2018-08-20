#!perl
use strict;
use warnings;

use lib 't';
use helper;

my $tmp = workdir();

ok -d "t/repos/compile-test", "compile-test repo exists" or die;
chdir "t/repos/compile-test" or die;
qx(rm -f *.tar.gz *.tgz); # just to be safe

runs_ok "genesis compile-kit --name test-kit --version 1.0.4";
ok -f "test-kit-1.0.4.tar.gz", "genesis compile-kit should create the tarball";
output_ok "tar -tzvf test-kit-1.0.4.tar.gz | awk '{print \$1, \$2, \$5, \$9}'", <<EOF, "tarball contents are correct";
drwxr-xr-x 0 0 test-kit-1.0.4/
drwxr-xr-x 0 0 test-kit-1.0.4/base/
-rw-r--r-- 0 28 test-kit-1.0.4/kit.yml
-rw-r--r-- 0 15 test-kit-1.0.4/base/params.yml
-rw-r--r-- 0 46 test-kit-1.0.4/base/stuff.yml
EOF

done_testing;
