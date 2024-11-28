{
  description = "Demo Nix dev environment";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
    nixpkgs.url = "github:NixOS/nixpkgs/release-24.05";
  };
  outputs = { self, flake-utils, nixpkgs }@inputs:
    flake-utils.lib.eachDefaultSystem (system:
      let
        project_dep = inputs.nixpkgs.${system};
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = project_dep.mkShell {
          packages = with pkgs; [ podman podman-compose ];

          shellHook = ''
            nix --version
          '';
        };
      });
}

