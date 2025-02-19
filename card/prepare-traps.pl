#!/usr/bin/perl

use strict;
use warnings;
use String::Util qw(collapse unquote);

my $fn = 'Upsmate.mib';
my $ent = '.1.3.6.1.4.1.935.0.';

################################################################################
# создание переменных
#
if ( open(my $fh, $fn) ) {

    while (my $line = <$fh>) {

        if ( $line =~ m/^(\w+)\s+TRAP-TYPE$/ ) {

            my $name = 'Trap' . ucfirst($1);
            my $oid;

            while (my $sub = <$fh>) {

                if ( $sub =~ m/^ ::= (\d+)$/ ) {
                    $oid = $ent . $1;
                    last;
                }
            }

            print "\t$name = \"$oid\"\n";
        }

    }

} else {
    print "open($fn): ", $!, "\n";
}

print "\n";


################################################################################
# создание map
#
if ( open(my $fh, $fn) ) {

    while (my $line = <$fh>) {

        if ( $line =~ m/^(\w+)\s+TRAP-TYPE$/ ) {

            my $name  = $1;
            my $namek = 'Trap' . ucfirst($1);
            my $oid;
            my $desc;

            while (my $sub = <$fh>) {
                if ( $sub =~ m/\s+DESCRIPTION/ ) {
                    $desc = readline($fh);
                    last;
                }
            }

            while (my $sub = <$fh>) {

                if ( $sub =~ m/SUMMARY | ::= / ) {

                    if ( $sub =~ m/SUMMARY (.+)$/ ) {
                        $desc = $1;
                    }

                    if ( $sub =~ m/^ ::= (\d+)$/ ) {
                        $oid = $ent . $1;
                        last;
                    }
                }
            }

            $desc = unquote(collapse($desc));
            if ( $desc =~ m/(?:INFORMATION: |WARNING: )?([^\.]+)/ ) {
                $desc = $1;
            }

            print "\t$namek: trapSpec{name: \"$name\", desc: \"$desc\", send: false,},\n";
        }

    }

} else {
    print "open($fn): ", $!, "\n";
}

exit;
