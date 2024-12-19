package scope

import (
	"github.com/c2micro/mlan/pkg/engine/object"
)

// GlobalScope глобальный скоуп
var GlobalScope *Scope

// CurrentScope текущий скоуп
var CurrentScope *Scope

type Scope struct {
	// parent родительский скоуп
	parent *Scope
	// depth текущая вложенность скоупов
	depth int
	// isFunc скоуп внутри функции
	isFunc bool
	// isLoop скоуп внутри цикла
	isLoop bool
	// isLoopBreak необходим ли выход из цикла
	isLoopBreak bool
	// isLoopContinue необходимо ли пропрыгнуть до конца блока в цикле
	isLoopContinue bool
	// objects объекты, объявленные в этом скоупе
	objects map[string]object.Object
}

// NewScope создание нового скоупа
func NewScope(
	parent *Scope,
	depth int,
	isFunc bool,
	isLoop bool,
	objects map[string]object.Object,
) *Scope {
	return &Scope{
		parent:  parent,
		depth:   depth,
		isFunc:  isFunc,
		isLoop:  isLoop,
		objects: objects,
	}
}

// IsGlobal является ли скоуп глобальным
func (s *Scope) IsGlobal() bool {
	return s.parent == nil
}

// Parent получение родительского скоупа
func (s *Scope) Parent() *Scope {
	return s.parent
}

// IsFunc является ли скоуп внутри функции
func (s *Scope) IsFunc() bool {
	return s.isFunc
}

// IsInFunc проверяем, что один из вышележащих скоупов - функция
func (s *Scope) IsInFunc() bool {
	t := s
	for {
		if t == nil {
			return false
		}
		if t.isFunc {
			return true
		}
		t = t.parent
	}
}

// IsInLoop проверяем, что один из вышележащих скоупов - цикл
func (s *Scope) IsInLoop() bool {
	t := s
	for {
		if t == nil {
			return false
		}
		if t.isLoop {
			return true
		}
		t = t.parent
	}
}

// IsLoop является ли скоуп внутри цикла
func (s *Scope) IsLoop() bool {
	return s.isLoop
}

// IsBreak необходимо ли выйти из цикла
func (s *Scope) IsBreak() bool {
	return s.isLoopBreak
}

// IsContinue необходимо ли пропрыгнуть в конец блока цикла
func (s *Scope) IsContinue() bool {
	return s.isLoopContinue
}

// установка индикатора необходимости выхода из цикла
func (s *Scope) SetLoopBreak() {
	t := s
	for {
		if t.isLoop {
			// set break on first loop context
			t.isLoopBreak = true
			break
		}
		t = t.parent
	}
}

// снятие индикатора необходимости выхода из цикла
func (s *Scope) UnsetLoopBreak() {
	t := s
	for {
		if t.isLoop {
			// set break on first loop context
			t.isLoopBreak = false
			break
		}
		t = t.parent
	}
}

// установка индикатора необходимости пропыга блока цикла
func (s *Scope) SetLoopContinue() {
	t := s
	for {
		if t.isLoop {
			// set continue on first loop context
			t.isLoopContinue = true
			break
		}
		t = t.parent
	}
}

// снятие индикатора необходимости пропыга блока цикла
func (s *Scope) UnsetLoopContinue() {
	t := s
	for {
		if t.isLoop {
			// set continue on first loop context
			t.isLoopContinue = false
			break
		}
		t = t.parent
	}
}

// Depth возвращение текущей вложенности
func (s *Scope) Depth() int {
	return s.depth
}

// Get получение объекта из скоупа (в том числе с возможностью поиска в родительском скоупе)
func (s *Scope) Get(n string, p bool) object.Object {
	val, ok := s.objects[n]
	if ok {
		return val
	}
	if p && !s.IsGlobal() {
		// искать в родительском скоупе
		return s.Parent().Get(n, p)
	} else {
		// ничего не найдено
		return nil
	}
}

// Put сохранение объекта в скоуп
func (s *Scope) Put(n string, o object.Object) {
	val := s.Get(n, !s.IsFunc())
	if val == nil {
		s.objects[n] = o
	} else {
		s.reAssign(n, o)
	}
}

// reAssign переприсвоение значения в мапе
func (s *Scope) reAssign(n string, o object.Object) {
	val := s.objects[n]
	if val == nil {
		s.Parent().reAssign(n, o)
	} else {
		s.objects[n] = o
	}
}
