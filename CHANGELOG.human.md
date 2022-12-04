# Changelog

## v0.1.0

- Added parsing and validating feature to darktool on 'darktool validate'
- Parses market_ships.ini, universe.ini and infocards.txt
- Base names from infocards forwards to market_ships.ini as a commented out parameter
- Rewrites market_ships.ini as a completely enforced in style file
- Supports having comments at the file beginning and as commented out parameters on top of ini syntax
- Implements unique INI reader capable to parse Freelancer INI configs and write them back

## v0.1.1

- Fix bug with float parsing in inireader

## **0.2.0** <sub><sup>2022-12-04 ([8d52846...33ad0b6](https://github.com/darklab8/darklab_freelancer_darktool/compare/8d52846...33ad0b6?diff=split))</sup></sub>

### Features
*  max precision of floats is 1 ([8d52846](https://github.com/darklab8/darklab_freelancer_darktool/commit/8d52846))
*  all param keys to lowercase ([e658fa8](https://github.com/darklab8/darklab_freelancer_darktool/commit/e658fa8))
*  enforce case sensetive section types ([fcddc5c](https://github.com/darklab8/darklab_freelancer_darktool/commit/fcddc5c))
*  autofix unnecessary space in values ([385aaa1](https://github.com/darklab8/darklab_freelancer_darktool/commit/385aaa1))
*  market\_commodities\.txt is reformated too ([5ca7608](https://github.com/darklab8/darklab_freelancer_darktool/commit/5ca7608))
*  reformating style of universe\.ini ([33ad0b6](https://github.com/darklab8/darklab_freelancer_darktool/commit/33ad0b6))