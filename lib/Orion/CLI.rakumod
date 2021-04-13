proto MAIN(|) is export { unless so @*ARGS { say $*USAGE; exit }; {*} }

multi sub MAIN(
) { }

multi sub MAIN(
    Bool :$version #= print version
) { say "Orion v" ~ $?DISTRIBUTION.meta<version>; }
