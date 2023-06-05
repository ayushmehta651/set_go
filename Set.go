package main

import "fmt"

type Set[T comparable] struct {
	//using map here as set uses maps internally
	mp map[T]bool
}

//initialize the map
func (s *Set[T]) CreateSet() {
	s.mp = make(map[T]bool, 0)
}

func (s *Set[T]) Add(ele T) {
	s.mp[ele] = true
}

func (set *Set[T]) PrintSet() {
	fmt.Print("[ ")
	for key, _ := range set.mp {
		fmt.Printf("%v ", key)
	}
	fmt.Println("]")
}

func (s *Set[T]) Union(set1, set2 *Set[T]) map[T]bool {
	s.CreateSet()
	for key, _ := range set1.mp {
		s.mp[key] = true
	}
	for key, _ := range set2.mp {
		s.mp[key] = true
	}
	return s.mp
}

func (s *Set[T]) Intersection(set1, set2 *Set[T]) map[T]bool {
	s.CreateSet()
	if len(set1.mp) > len(set2.mp) {
		set1, set2 = set2, set1
	}
	for key, _ := range set1.mp {
		if set2.mp[key] {
			s.mp[key] = true
		}
	}
	return s.mp
}

func main() {
	s1 := Set[int]{} // Instantiation of Generic type
	s1.CreateSet()
	s1.Add(22)
	s1.Add(23)
	s1.Add(35)
	s1.Add(35)
	s1.Add(43)
	s1.Add(35)
	s1.Add(390)
	fmt.Println("Set1: ")
	s1.PrintSet()

	s2 := Set[int]{}
	s2.CreateSet()
	s2.Add(22)
	s2.Add(23)
	s2.Add(35)
	s2.Add(35)
	fmt.Println("Set2: ")
	s2.PrintSet()

	union_set := Set[int]{}
	union_set.Union(&s1, &s2)
	fmt.Println("Union set: ")
	union_set.PrintSet()

	intersection_set := Set[int]{}
	intersection_set.Intersection(&s1, &s2)
	fmt.Println("Intersection set: ")
	intersection_set.PrintSet()
}
