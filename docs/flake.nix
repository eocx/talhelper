{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
  };

  outputs = inputs:
    inputs.flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [
        "aarch64-linux"
        "x86_64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      perSystem = {
        system,
        pkgs,
        ...
      }: {
        packages.default = pkgs.callPackage ./default.nix {};
        devShells.default = with pkgs; mkShell {
          name = "talhelper-docs-dev";
          packages = [
            python311Packages.mkdocs
            python311Packages.pymdown-extensions
            python311Packages.pygments
            python311Packages.mike
          ];
        };
      };
    };
}
