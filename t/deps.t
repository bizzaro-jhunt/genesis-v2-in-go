#!perl
use strict;
use warnings;

use lib 't';
use helper;

my $tmp = workdir;
ok -d "t/repos/kit-deps-test", "kit-deps-test repo exists" or die;
chdir "t/repos/kit-deps-test" or die;
qx(rm -f *.yml);

$ENV{SHOULD_FAIL} = '';
runs_ok "genesis new successful-env";
ok -f "successful-env.yml", "Environment file should be created, when prereqs passes";

$ENV{SHOULD_FAIL} = 'yes';
run_fails "genesis new failed-env", 1;
ok ! -f "failed-env.yml", "Environment file should not be created, when prereqs fails";

done_testing;
