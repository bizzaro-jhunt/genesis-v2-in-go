#!perl
use strict;
use warnings;

use lib 't';
use helper;

vault_ok;

runs_ok "genesis ping";
ok -d "t/repos/vault-test", "vault-test repo exists" or die;
chdir "t/repos/vault-test";

qx(rm -f *.yml);
runs_ok "genesis new test-env --vault unit-tests";
diag "connecting to the local vault (this may take a while)...";
no_secret "secret/test/env/random:password", "'control' secret doesn't exist before we run `genesis secrets`";

put_file "test-env.yml", <<EOF;
---
params:
  env:   test-env
  vault: test/env
EOF
runs_ok "genesis secrets test-env --vault unit-tests";

my $secret;
diag "connecting to the local vault (this may take a while)...";
secret_exists "secret/test/env/random:password" or die;
$secret = secret "secret/test/env/random:password";
is length($secret), 42, "random password length parameter honored";

secret_exists "secret/test/env/ssh:private" or die;
$secret = secret "secret/test/env/ssh:private";
ok $secret =~ m/--BEGIN RSA PRIVATE KEY--.*--END RSA PRIVATE KEY--/s,
	"SSH private key is in PEM format";

secret_exists "secret/test/env/ssh:public" or die;
$secret = secret "secret/test/env/ssh:public";
ok $secret =~ m/^ssh-rsa A/, "SSH public key is in OpenSSH format";

secret_exists "secret/test/env/rsa:private" or die;
$secret = secret "secret/test/env/rsa:private";
ok $secret =~ m/--BEGIN RSA PRIVATE KEY--.*--END RSA PRIVATE KEY--/s,
	"RSA private key is in PEM format";

secret_exists "secret/test/env/rsa:public" or die;
$secret = secret "secret/test/env/rsa:public";
ok $secret =~ m/--BEGIN PUBLIC KEY--.*--END PUBLIC KEY--/s,
	"RSA public key is in PEM format";

sub cache_creds($) {
	my ($ref) = @_;
	my %check = (
	  random => 'secret/test/env/random:password',
	  fixed  => 'secret/test/env/fixed:password',
	  sshpub => 'secret/test/env/ssh:public',
	  sshkey => 'secret/test/env/ssh:private',
	  rsapub => 'secret/test/env/rsa:public',
	  rsakey => 'secret/test/env/rsa:private',

	  ext_sshpub => 'secret/test/env/ext_ssh:public',
	  ext_sshkey => 'secret/test/env/ext_ssh:private',
	  ext_rsapub => 'secret/test/env/ext_rsa:public',
	  ext_rsakey => 'secret/test/env/ext_rsa:private',
	);
	for my $k (keys %check) {
		$ref->{$k} = secret $check{$k};
	}
}

sub creds_changed($$@) {
	my ($before, $after, @keys) = @_;
	my (%all, %diff);
	$all{$_} = 1 for keys %$before;
	$all{$_} = 1 for keys %$after;

	$diff{$_} = 1 for @keys;
	if (@keys == 0) {
		$diff{$_} = 1 for keys %$before;
		$diff{$_} = 1 for keys %$after;
	}

	for (keys %all) {
		if ($diff{$_}) {
			isnt $before->{$_}, $after->{$_}, "$_ should have been changed.";
		} else {
			is $before->{$_}, $after->{$_}, "$_ should not have changed.";
		}
	}
}

my (%before, %regen, %rotated);
cache_creds \%before;
runs_ok "genesis secrets test-env --vault unit-tests";
cache_creds \%regen;
runs_ok "genesis secrets --rotate test-env --vault unit-tests";
cache_creds \%rotated;

creds_changed \%before, \%regen;
creds_changed \%regen,  \%rotated, qw(random sshpub sshkey rsapub rsakey);

qx(rm -f a-b-c.yml);
diag "connecting to the local vault (this may take a while)...";
runs_ok "genesis new a-b-c --vault unit-tests";
yaml_is get_file("a-b-c.yml"), <<EOF, "`genesis new` creates the correct starting YAML";
---
params:
  env:   a-b-c
  vault: a/b/c/vault/test
EOF
secret_exists "secret/a/b/c/vault/test/random:password";

# what if we have no Vault targets in ~/.saferc?
unlink "$ENV{HOME}/.saferc";
run_fails "genesis new no-vault-here", 1;
ok ! -f "no-vault-here.yml";

teardown_vault;
done_testing;
