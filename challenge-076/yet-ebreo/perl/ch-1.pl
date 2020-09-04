#! /usr/bin/perl
use strict;
use feature 'say';

my $N   = $ARGV[0] || 9;

for (2..$N/2) {
    (1x$_) !~ /^1?$|^(11+?)\1+$/ && (1x($N - $_)) !~ /^1?$|^(11+?)\1+$/ && say "$_ + ".($N - $_)." = $N";
}
=begin
perl ch-1.pl 90
7 + 83 = 90
11 + 79 = 90
17 + 73 = 90
19 + 71 = 90
23 + 67 = 90
29 + 61 = 90
31 + 59 = 90
37 + 53 = 90
43 + 47 = 90
=cut