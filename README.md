# ethereum-spikes

My random collection of experiments for understanding things in the
ethereum ecosystem.


## Collaborating

I'm not really expecting collaboration on this project. Reach out if
there's something here you find interesting and useful.


### Project Conventions & Organization

I find little command line utilities helpful. This project contains
several within the ./cmd/... package structure and they're generally
implemented with https://cobra.dev/

Running `make` without arguments will build the entire project.

Running `make watch` will start a build loop. It watches the source tree
for updates and builds the project, providing a fast feedback loop on
your edits.

I also like to keep my project dependencies isolated (contained within
the project, instead of across my system) and rely on `direnv` to help
me do that. I consider this tool optional, but recommended. You can read
about it here: https://direnv.net/


## Things I want to understand

- [ ] discv5
- [ ] libp2p
- [ ] where's the boundary between discv5 & libp2p
