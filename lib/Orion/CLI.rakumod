use LWP::Simple;
use Digest::SHA;

multi sub MAIN(
    Bool :$verbose, #=increase verbosity
) is export {
    my Str $api = "https://api.pwnedpasswords.com";

    # Gather all `.gpg' files from the store.
    my @stack = "$*HOME/.password-store".IO;
    my @files = gather while @stack {
        with @stack.pop {
            when :d { @stack.append: .dir }
            .take when .extension.lc eq 'gpg'
        }
    }

    for @files -> $file {
        print "." if $verbose;

        my $proc = run <gpg2>, '--decrypt', '--quiet', $file, :out;
        # Read the password, convert it to hex & uppercase.
        my Str $hash = uc buf_to_hex sha1 $proc.out.slurp(:close).lines.head;

        # Send the prefix to HIBP API.
        for LWP::Simple.get(
            "$api/range/{ $hash.substr(0, 5) }",
            {
                Add-Padding => "true",
                User-Agent => "Andinus / Orion - https://andinus.nand.sh/orion",
            }
        ).lines.map(*.split(":")).grep(*[1] > 0) -> ($entry, $freq) {
            # Compare all the suffixes with our hash suffix.
            if $hash.substr(5, *) eq $entry {
                # Get the password name.
                my Str $pass = $file.Str.split('.password-store/').tail.split('.gpg').first;
                # Print the compromised password entry.
                print "\n" if $verbose;
                put $pass, " " x (72 - $pass.chars - $freq.chars), $freq;
            }
        }
    }
    print "\n" if $verbose;
}

sub buf_to_hex { [~] $^buf.list».fmt: "%02x" }

multi sub MAIN(
    Bool :$version #= print version
) { say "Orion v" ~ $?DISTRIBUTION.meta<version>; }
