package helper;
use Test::More;
use Cwd ();
use Config;
use File::Temp qw/tempdir/;

sub import {
	my ($class, @args) = @_;
	(my $perlpath = $Config{perlpath}) =~ s|/perl$||;
	$ENV{PATH} = Cwd::getcwd().":".$perlpath.":$ENV{PATH}";

	my $caller = caller;
	for my $glob (keys %helper::) {
		next if not defined *{$helper::{$glob}}{CODE}
		         or $glob eq 'import';
		*{$caller . "::$glob"} = \&{"helper::$glob"};
	}
}

our $WORKDIR = undef;
sub workdir() {
	$WORKDIR ||= tempdir(CLEANUP => 1);
}

sub put_file($$) {
	my ($file, $contents) = @_;
	open my $fh, ">", $file
		or fail "failed to open '$file' for writing: $!";

	print $fh $contents;
	close $fh;
}

sub get_file($) {
	my ($file) = @_;
	open my $fh, "<", $file
		or fail "failed to open '$file' for reading: $!";

	my $contents = do { local $/; <$fh> };
	close $fh;
	return $contents;
}

sub spruce_fmt($$) {
	my ($yaml, $file) = @_;
	open my $fh, "|-", "spruce merge - >$file"
		or die "Failed to reformat YAML via spruce: $!\n";
	print $fh $yaml;
	close $fh;
}

sub runs_ok($;$) {
	my ($cmd, $msg) = @_;
	$msg ||= "running `$cmd`";

	my $err = qx($cmd 2>&1);
	if ($? != 0) {
		my $exit = $? >> 8;
		fail $msg;
		diag "`$cmd` exited $exit";
		diag "";
		diag "----[ output from failing command: ]-----------";
		diag $err ? $err : '(no output)';
		diag "-----------------------------------------------";
		diag "";
		return 0;
	}
	pass $msg;
	return 1;
}

sub run_fails($$;$) {
	my ($cmd, $rc, $msg) = @_;
	$msg ||= "running `$cmd` (expecting exit code $rc)";
	if (defined $rc && $rc !~ m/^\d+$/) {
		$msg = $rc;
		$rc = undef;
	}

	my $err = qx($cmd 2>&1);
	my $exit = $? >> 8;
	if (defined $rc  && $exit != $rc) {
		fail $msg;
		diag "`$cmd` exited $exit (instead of $rc)";
		diag $err;
		return 0;
	} elsif (!defined $rc && $exit == 0) {
		fail $msg;
		diag "`$cmd` exited $exit (instead of non-zero)";
		diag $err;
		return 0;
	}
	pass $msg;
	return 1;
}

sub output_ok($$;$) {
	my ($cmd, $expect, $msg) = @_;
	$msg ||= "`$cmd` ≅ '$expect'";

	my $dir = workdir;
	my $got = qx($cmd 2>&1);
	if ($? != 0) {
		my $exit = $? >> 8;
		fail $msg;
		diag "`$cmd` exited $exit";
		diag $got;
		return 0;
	}
	$got =~ s/\s+$//mg;
	$expect =~ s/\s+$//mg;

	if ($got ne $expect) {
		fail "$msg: output was different.";
		put_file "$dir/got",    "$got\n";
		put_file "$dir/expect", "$expect\n";
		diag qx(cd $dir/; diff -u expect got);
		return 0;
	}
	pass $msg;
	return 0;
}

my $VAULT_PID;
sub vault_ok {
	if (defined $VAULT_PID) {
		pass "vault already running.";
		return 1;
	}

	$ENV{HOME} = "$ENV{PWD}/t/tmp/home";
	my $pid = qx(./t/bin/vault) or do {
		fail "failed to spin a vault server in (-dev) mode.";
		return 0;
	};

	chomp($pid);
	$VAULT_PID = $pid;
	kill -0, $pid or do {
		fail "failed to spin a vault server in (-dev) mode: couldn't signal pid $pid.";
		return 0;
	};
	pass "vault running [pid $pid]";
	return 1;
}

sub teardown_vault {
	if (defined $VAULT_PID) {
		kill 'TERM', $VAULT_PID;
	}
}

sub no_secret($;$) {
	my ($secret, $msg) = @_;
	$msg ||= "secret '$secret' should not exist";
	qx(safe exists $secret);
	if ($? == 0) {
		fail $msg;
		diag "    (safe exited $?)";
		return 0;
	}

	pass $msg;
	return 1;
}

sub secret_exists($;$) {
	my ($secret, $msg) = @_;
	$msg ||= "secret '$secret' should exist";
	qx(safe exists $secret);
	if ($? != 0) {
		fail $msg;
		diag "    (safe exited $?)";
		return 0;
	}

	pass $msg;
	return 1;
}

sub secret($) {
	chomp(my $secret = qx(safe read $_[0]));
	return $secret;
}

sub yaml_is($$$) {
	my ($got, $expect, $msg) = @_;
	my $dir = workdir;
	spruce_fmt $got,    "$dir/got.yml";
	spruce_fmt $expect, "$dir/expect.yml";

	$got    = get_file "$dir/got.yml";
	$expect = get_file "$dir/expect.yml";

	if ($got eq $expect) {
		pass $msg;
		return 1;
	}
	fail "$msg: strings were different.";
	diag qx(cd $dir/; diff -u expect.yml got.yml);
	return 0;
}

1;
