#!perl
use strict;
use warnings;

use lib 't';
use helper;

qx(rm -rf t/tmp; mkdir -p t/tmp);
chdir "t/tmp";

runs_ok "genesis init new";
ok -d "new-deployments", "created initial deployments directory";
chdir "new-deployments";

mkdir "dev";
put_file "dev/kit.yml", <<EOF;
---
name: Sample Dev Kit
EOF

ok ! -f "x-y-z.yml",            "x-y-z.yml doesn't exist before we test `genesis new`";
runs_ok "genesis new x-y-z";
ok   -f "x-y-z.yml",            "x-y-z.yml created by `genesis new`";
run_fails "genesis new x-y-z",  "`genesis new` refuses to overwrite existing files";
ok   -f "x-y-z.yml",             "x-y-z.yml wasn't clobbered by a bad `genesis new` command";

run_fails "genesis new *best*", "`genesis new` validates environment names";
ok ! -f "*best*.yml",           "`genesis new` refused to create env file when name validation failed";
run_fails "genesis new a--b",   "`genesis new` doesn't allow multi-dash environment names";
ok ! -f "a--b.yml",             "`genesis new` refused to create env file when name validation failed";

qx(rm -rf t/tmp);
done_testing;
