use LWP::Simple;
use Digest::SHA;

#proto MAIN(|) is export { unless so @*ARGS { say $*USAGE; exit }; {*} }

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
        my $proc = run <gpg2>, '--decrypt', '--quiet', $file, :out;
        my Str $hash = uc buf_to_hex sha1 $proc.out.slurp(:close).lines.head;

        for LWP::Simple.get(
            "$api/range/{ $hash.substr(0, 5) }",
            {
                Add-Padding => "true",
                User-Agent => "Andinus / Orion - https://andinus.nand.sh/orion",
            }
        ).lines.map(*.split(":")) -> ($entry, $freq) {
            next if $freq == 0;
            if $hash.substr(5, *) eq $entry.split(":").head {
                say $file.Str.split('.password-store/').tail.split('.gpg').first, " ",
                "has been seen $freq times before.";
            };
        }
    }
}

sub buf_to_hex { [~] $^buf.list».fmt: "%02x" }

multi sub MAIN(
    Bool :$version #= print version
) { say "Orion v" ~ $?DISTRIBUTION.meta<version>; }
