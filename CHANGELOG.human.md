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

## **0.3.0** <sub><sup>2022-12-04 ([cefe3f0...48b57c0](https://github.com/darklab8/darklab_freelancer_darktool/compare/cefe3f0...48b57c0?diff=split))</sup></sub>

### Features
*  dry run option ([f9efb7b](https://github.com/darklab8/darklab_freelancer_darktool/commit/f9efb7b))
*  allow --dry flag from CI ([40f9742](https://github.com/darklab8/darklab_freelancer_darktool/commit/40f9742))
*  add  command version ([967bfce](https://github.com/darklab8/darklab_freelancer_darktool/commit/967bfce))
*  validate \-\-search 'absolute\_path' flag\. So u could run at chosen freelancer folder location ([58b2944](https://github.com/darklab8/darklab_freelancer_darktool/commit/58b2944))
*  safeguard against deleting your OS\. check for DATA and SERVICE folders ([239ab19](https://github.com/darklab8/darklab_freelancer_darktool/commit/239ab19))


### Bug Fixes
*  ignore case sensetive values ([b555608](https://github.com/darklab8/darklab_freelancer_darktool/commit/b555608))
*  test discovery ([48b57c0](https://github.com/darklab8/darklab_freelancer_darktool/commit/48b57c0))

## **0.4.0** <sub><sup>2022-12-05 ([3a5d122...54376f1](https://github.com/darklab8/darklab_freelancer_darktool/compare/3a5d122...54376f1?diff=split))</sup></sub>

### Features
 - semantic parse of universe.ini and writing it back

 ## **0.4.1** <sub><sup>2022-12-05 ([6898c81...6898c81](https://github.com/darklab8/darklab_freelancer_darktool/compare/6898c81...6898c81?diff=split))</sup></sub>

### Bug Fixes
*  missing to lowercase in file of universe ([6898c81](https://github.com/darklab8/darklab_freelancer_darktool/commit/6898c81))

## **0.5.0** <sub><sup>2022-12-05 ([2363a46...adf20f1](https://github.com/darklab8/darklab_freelancer_darktool/compare/2363a46...adf20f1?diff=split))</sup></sub>

### Features
*  setup systems parsing preparations ([968fad1](https://github.com/darklab8/darklab_freelancer_darktool/commit/968fad1))
*  iniread systems ([7512e0c](https://github.com/darklab8/darklab_freelancer_darktool/commit/7512e0c))
*  semantic parsed bases in systems ([c44b43e](https://github.com/darklab8/darklab_freelancer_darktool/commit/c44b43e))
*  rendered recycling bases ([adf20f1](https://github.com/darklab8/darklab_freelancer_darktool/commit/adf20f1), [#7](https://github.com/darklab8/darklab_freelancer_darktool/issues/#7))