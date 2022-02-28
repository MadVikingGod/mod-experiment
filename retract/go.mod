module github.com/madvikinggod/mod-experiment/retract

go 1.17

replace github.com/madvikinggod/mod-experiment => ../

require github.com/madvikinggod/mod-experiment v0.2.0

retract [v0.1.0, v0.4.9]
