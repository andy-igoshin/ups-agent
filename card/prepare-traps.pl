#!/usr/bin/perl

use strict;
use warnings;
use String::Util qw(collapse unquote);

my $fn = 'Upsmate.mib';
my $ent = '.1.3.6.1.4.1.935.0.';

################################################################################
# создание map
#
if ( open(my $fh, $fn) ) {

    while (my $line = <$fh>) {

        if ( $line =~ m/^(\w+)\s+TRAP-TYPE$/ ) {

            my $name  = $1;
            my $oid;

            while (my $sub = <$fh>) {

                if ( $sub =~ m/^ ::= (\d+)$/ ) {
                    $oid = $ent . $1;
                    last;
                }
            }

            print "\t\"$oid\": trapSpec{name: \"$name\"},\n";
        }

    }

} else {
    print "open($fn): ", $!, "\n";
}

exit;
