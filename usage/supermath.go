package main

import (
	"plugin"
)


/**
 * Содержимое библиотеки
 */
type SuperMath struct {
	Add func(int, int) int
	Sub func(int, int) int
	Mult func(int, int) int
}


/**
 * Загружает функции из плагина
 */
func LoadSuperMath(pluginPath string) (*SuperMath, error) {
	
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, err
	}

	Add, err := p.Lookup("Add")
	if err != nil {
		return nil, err
	}

	Sub, err := p.Lookup("Sub")
	if err != nil {
		return nil, err
	}

	Mult, err := p.Lookup("Mult")
	if err != nil {
		return nil, err
	}

	supermath := &SuperMath{
		Add: Add.(func(int, int) int),
		Sub: Sub.(func(int, int) int),
		Mult: Mult.(func(int, int) int),
	}
	
	return supermath, nil
}

